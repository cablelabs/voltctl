help:
	@echo "release      - build binaries using cross compliing for the support architectures"
	@echo "build        - build the binary as a local executable"
	@echo "install      - build and install the binary into \$$GOPATH/bin"
	@echo "run          - runs voltctl using the command specified as \$$CMD"
	@echo "lint-style   - Verify code is properly gofmt-ed"
	@echo "lint-sanity  - Verify that 'go vet' doesn't report any issues"
	@echo "lint-mod     - Verify the integrity of the 'mod' files"
	@echo "lint         - run static code analysis"
	@echo "sca          - Runs various SCA through golangci-lint tool"
	@echo "test         - run unity tests"
	@echo "check        - runs targets that should be run before a commit"
	@echo "clean        - remove temporary and generated files"

SHELL=bash -e -o pipefail

VERSION=$(shell cat ./VERSION)
GITCOMMIT=$(shell git rev-parse HEAD)
ifeq ($(shell git ls-files --others --modified --exclude-standard 2>/dev/null | wc -l | sed -e 's/ //g'),0)
GITDIRTY=false
else
GITDIRTY=true
endif
GOVERSION=$(shell go version 2>&1 | sed -E  's/.*(go[0-9]+\.[0-9]+\.[0-9]+).*/\1/g')
HOST_OS=$(shell uname -s | tr A-Z a-z)
ifeq ($(shell uname -m),x86_64)
	HOST_ARCH ?= amd64
else
	HOST_ARCH ?= $(shell uname -m)
endif
BUILDTIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS=-ldflags \
	"-X \"github.com/opencord/voltctl/internal/pkg/cli/version.Version=$(VERSION)\"  \
	 -X \"github.com/opencord/voltctl/internal/pkg/cli/version.VcsRef=$(GITCOMMIT)\"  \
	 -X \"github.com/opencord/voltctl/internal/pkg/cli/version.VcsDirty=$(GITDIRTY)\"  \
	 -X \"github.com/opencord/voltctl/internal/pkg/cli/version.GoVersion=$(GOVERSION)\"  \
	 -X \"github.com/opencord/voltctl/internal/pkg/cli/version.Os=$(HOST_OS)\" \
	 -X \"github.com/opencord/voltctl/internal/pkg/cli/version.Arch=$(HOST_ARCH)\" \
	 -X \"github.com/opencord/voltctl/internal/pkg/cli/version.BuildTime=$(BUILDTIME)\""

# Release related items
# Generates binaries in $RELEASE_DIR with name $RELEASE_NAME-$RELEASE_OS_ARCH
# Inspired by: https://github.com/kubernetes/minikube/releases
RELEASE_DIR     ?= release
RELEASE_NAME    ?= voltctl
RELEASE_OS_ARCH ?= linux-amd64 linux-arm64 windows-amd64 darwin-amd64

# tool containers
VOLTHA_TOOLS_VERSION ?= 2.4.0

GO                = docker run --rm --user $$(id -u):$$(id -g) -v ${CURDIR}:/app $(shell test -t 0 && echo "-it") -v gocache:/.cache -v gocache-${VOLTHA_TOOLS_VERSION}:/go/pkg voltha/voltha-ci-tools:${VOLTHA_TOOLS_VERSION}-golang go
GO_SH             = docker run --rm --user $$(id -u):$$(id -g) -v ${CURDIR}:/app $(shell test -t 0 && echo "-it") -v gocache:/.cache -v gocache-${VOLTHA_TOOLS_VERSION}:/go/pkg voltha/voltha-ci-tools:${VOLTHA_TOOLS_VERSION}-golang sh -c
GO_JUNIT_REPORT   = docker run --rm --user $$(id -u):$$(id -g) -v ${CURDIR}:/appecho  -i voltha/voltha-ci-tools:${VOLTHA_TOOLS_VERSION}-go-junit-report go-junit-report
GOCOVER_COBERTURA = docker run --rm --user $$(id -u):$$(id -g) -v ${CURDIR}:/app -i voltha/voltha-ci-tools:${VOLTHA_TOOLS_VERSION}-gocover-cobertura gocover-cobertura
GOLANGCI_LINT     = docker run --rm --user $$(id -u):$$(id -g) -v ${CURDIR}:/app $(shell test -t 0 && echo "-it") -v gocache:/.cache -v gocache-${VOLTHA_TOOLS_VERSION}:/go/pkg -e GOGC=10 voltha/voltha-ci-tools:${VOLTHA_TOOLS_VERSION}-golangci-lint golangci-lint

release:
	@mkdir -p $(RELEASE_DIR)
	@${GO_SH} ' \
	  set -e -o pipefail; \
	  for x in ${RELEASE_OS_ARCH}; do \
	    OUT_PATH="$(RELEASE_DIR)/$(RELEASE_NAME)-$(subst -dev,_dev,$(VERSION))-$$x"; \
	    echo "$$OUT_PATH"; \
	    GOOS=$${x%-*} GOARCH=$${x#*-} go build -mod=vendor -v $(LDFLAGS) -o "$$OUT_PATH" cmd/voltctl/voltctl.go; \
	  done'
## Local Development Helpers
local-lib-go:
ifdef LOCAL_LIB_GO
	rm -rf vendor/github.com/opencord/voltha-lib-go/v7/pkg
	mkdir -p vendor/github.com/opencord/voltha-lib-go/v7/pkg
	cp -r ${LOCAL_LIB_GO}/pkg/* vendor/github.com/opencord/voltha-lib-go/v7/pkg/
endif

build: local-lib-go
	go build -mod=vendor $(LDFLAGS) cmd/voltctl/voltctl.go

install:
	go install -mod=vendor $(LDFLAGS) cmd/voltctl/voltctl.go

run:
	go run -mod=vendor $(LDFLAGS) cmd/voltctl/voltctl.go $(CMD)

lint-mod:
	@echo "Running dependency check..."
	@${GO} mod verify
	@echo "Dependency check OK. Running vendor check..."
	@git status > /dev/null
	@git diff-index --quiet HEAD -- go.mod go.sum vendor || (echo "ERROR: Staged or modified files must be committed before running this test" && git status -- go.mod go.sum vendor && exit 1)
	@[[ `git ls-files --exclude-standard --others go.mod go.sum vendor` == "" ]] || (echo "ERROR: Untracked files must be cleaned up before running this test" && git status -- go.mod go.sum vendor && exit 1)
	${GO} mod tidy
	${GO} mod vendor
	@git status > /dev/null
	@git diff-index --quiet HEAD -- go.mod go.sum vendor || (echo "ERROR: Modified files detected after running go mod tidy / go mod vendor" && git status -- go.mod go.sum vendor && git checkout -- go.mod go.sum vendor && exit 1)
	@[[ `git ls-files --exclude-standard --others go.mod go.sum vendor` == "" ]] || (echo "ERROR: Untracked files detected after running go mod tidy / go mod vendor" && git status -- go.mod go.sum vendor && git checkout -- go.mod go.sum vendor && exit 1)
	@echo "Vendor check OK."

lint: lint-mod

sca:
	@rm -rf ./sca-report
	@mkdir -p ./sca-report
	@echo "Running static code analysis..."
	@${GOLANGCI_LINT} run --deadline=20m --out-format junit-xml ./... | tee ./sca-report/sca-report.xml
	@echo ""
	@echo "Static code analysis OK"

test:
	@mkdir -p ./tests/results
	@${GO} test -mod=vendor -v -coverprofile ./tests/results/go-test-coverage.out -covermode count ./... 2>&1 | tee ./tests/results/go-test-results.out ;\
	RETURN=$$? ;\
	${GO_JUNIT_REPORT} < ./tests/results/go-test-results.out > ./tests/results/go-test-results.xml ;\
	${GOCOVER_COBERTURA} < ./tests/results/go-test-coverage.out > ./tests/results/go-test-coverage.xml ;\
	exit $$RETURN

view-coverage:
	go tool cover -html ./tests/results/go-test-coverage.out

check: lint sca test

clean:
	rm -rf voltctl voltctl.cp release sca-report

mod-update:
	${GO} mod tidy
	${GO} mod vendor
