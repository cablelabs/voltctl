// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/omci_alarm_db.proto

package omci

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/opencord/voltha-protos/v5/go/common"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AlarmOpenOmciEventType_OpenOmciEventType int32

const (
	AlarmOpenOmciEventType_state_change AlarmOpenOmciEventType_OpenOmciEventType = 0
)

var AlarmOpenOmciEventType_OpenOmciEventType_name = map[int32]string{
	0: "state_change",
}

var AlarmOpenOmciEventType_OpenOmciEventType_value = map[string]int32{
	"state_change": 0,
}

func (x AlarmOpenOmciEventType_OpenOmciEventType) String() string {
	return proto.EnumName(AlarmOpenOmciEventType_OpenOmciEventType_name, int32(x))
}

func (AlarmOpenOmciEventType_OpenOmciEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d41f1e38aadb08d, []int{6, 0}
}

type AlarmAttributeData struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmAttributeData) Reset()         { *m = AlarmAttributeData{} }
func (m *AlarmAttributeData) String() string { return proto.CompactTextString(m) }
func (*AlarmAttributeData) ProtoMessage()    {}
func (*AlarmAttributeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d41f1e38aadb08d, []int{0}
}

func (m *AlarmAttributeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmAttributeData.Unmarshal(m, b)
}
func (m *AlarmAttributeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmAttributeData.Marshal(b, m, deterministic)
}
func (m *AlarmAttributeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmAttributeData.Merge(m, src)
}
func (m *AlarmAttributeData) XXX_Size() int {
	return xxx_messageInfo_AlarmAttributeData.Size(m)
}
func (m *AlarmAttributeData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmAttributeData.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmAttributeData proto.InternalMessageInfo

func (m *AlarmAttributeData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlarmAttributeData) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AlarmInstanceData struct {
	InstanceId           uint32                `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Created              string                `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Modified             string                `protobuf:"bytes,3,opt,name=modified,proto3" json:"modified,omitempty"`
	Attributes           []*AlarmAttributeData `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AlarmInstanceData) Reset()         { *m = AlarmInstanceData{} }
func (m *AlarmInstanceData) String() string { return proto.CompactTextString(m) }
func (*AlarmInstanceData) ProtoMessage()    {}
func (*AlarmInstanceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d41f1e38aadb08d, []int{1}
}

func (m *AlarmInstanceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmInstanceData.Unmarshal(m, b)
}
func (m *AlarmInstanceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmInstanceData.Marshal(b, m, deterministic)
}
func (m *AlarmInstanceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmInstanceData.Merge(m, src)
}
func (m *AlarmInstanceData) XXX_Size() int {
	return xxx_messageInfo_AlarmInstanceData.Size(m)
}
func (m *AlarmInstanceData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmInstanceData.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmInstanceData proto.InternalMessageInfo

func (m *AlarmInstanceData) GetInstanceId() uint32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *AlarmInstanceData) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *AlarmInstanceData) GetModified() string {
	if m != nil {
		return m.Modified
	}
	return ""
}

func (m *AlarmInstanceData) GetAttributes() []*AlarmAttributeData {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type AlarmClassData struct {
	ClassId              uint32               `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Instances            []*AlarmInstanceData `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AlarmClassData) Reset()         { *m = AlarmClassData{} }
func (m *AlarmClassData) String() string { return proto.CompactTextString(m) }
func (*AlarmClassData) ProtoMessage()    {}
func (*AlarmClassData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d41f1e38aadb08d, []int{2}
}

func (m *AlarmClassData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmClassData.Unmarshal(m, b)
}
func (m *AlarmClassData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmClassData.Marshal(b, m, deterministic)
}
func (m *AlarmClassData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmClassData.Merge(m, src)
}
func (m *AlarmClassData) XXX_Size() int {
	return xxx_messageInfo_AlarmClassData.Size(m)
}
func (m *AlarmClassData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmClassData.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmClassData proto.InternalMessageInfo

func (m *AlarmClassData) GetClassId() uint32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *AlarmClassData) GetInstances() []*AlarmInstanceData {
	if m != nil {
		return m.Instances
	}
	return nil
}

type AlarmManagedEntity struct {
	ClassId              uint32   `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmManagedEntity) Reset()         { *m = AlarmManagedEntity{} }
func (m *AlarmManagedEntity) String() string { return proto.CompactTextString(m) }
func (*AlarmManagedEntity) ProtoMessage()    {}
func (*AlarmManagedEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d41f1e38aadb08d, []int{3}
}

func (m *AlarmManagedEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmManagedEntity.Unmarshal(m, b)
}
func (m *AlarmManagedEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmManagedEntity.Marshal(b, m, deterministic)
}
func (m *AlarmManagedEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmManagedEntity.Merge(m, src)
}
func (m *AlarmManagedEntity) XXX_Size() int {
	return xxx_messageInfo_AlarmManagedEntity.Size(m)
}
func (m *AlarmManagedEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmManagedEntity.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmManagedEntity proto.InternalMessageInfo

func (m *AlarmManagedEntity) GetClassId() uint32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *AlarmManagedEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AlarmMessageType struct {
	MessageType          uint32   `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmMessageType) Reset()         { *m = AlarmMessageType{} }
func (m *AlarmMessageType) String() string { return proto.CompactTextString(m) }
func (*AlarmMessageType) ProtoMessage()    {}
func (*AlarmMessageType) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d41f1e38aadb08d, []int{4}
}

func (m *AlarmMessageType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMessageType.Unmarshal(m, b)
}
func (m *AlarmMessageType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMessageType.Marshal(b, m, deterministic)
}
func (m *AlarmMessageType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMessageType.Merge(m, src)
}
func (m *AlarmMessageType) XXX_Size() int {
	return xxx_messageInfo_AlarmMessageType.Size(m)
}
func (m *AlarmMessageType) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMessageType.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMessageType proto.InternalMessageInfo

func (m *AlarmMessageType) GetMessageType() uint32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

type AlarmDeviceData struct {
	DeviceId             string                `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Created              string                `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	LastAlarmSequence    uint32                `protobuf:"varint,3,opt,name=last_alarm_sequence,json=lastAlarmSequence,proto3" json:"last_alarm_sequence,omitempty"`
	LastSyncTime         string                `protobuf:"bytes,4,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
	Version              uint32                `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	Classes              []*AlarmClassData     `protobuf:"bytes,6,rep,name=classes,proto3" json:"classes,omitempty"`
	ManagedEntities      []*AlarmManagedEntity `protobuf:"bytes,7,rep,name=managed_entities,json=managedEntities,proto3" json:"managed_entities,omitempty"`
	MessageTypes         []*AlarmMessageType   `protobuf:"bytes,8,rep,name=message_types,json=messageTypes,proto3" json:"message_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AlarmDeviceData) Reset()         { *m = AlarmDeviceData{} }
func (m *AlarmDeviceData) String() string { return proto.CompactTextString(m) }
func (*AlarmDeviceData) ProtoMessage()    {}
func (*AlarmDeviceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d41f1e38aadb08d, []int{5}
}

func (m *AlarmDeviceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmDeviceData.Unmarshal(m, b)
}
func (m *AlarmDeviceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmDeviceData.Marshal(b, m, deterministic)
}
func (m *AlarmDeviceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmDeviceData.Merge(m, src)
}
func (m *AlarmDeviceData) XXX_Size() int {
	return xxx_messageInfo_AlarmDeviceData.Size(m)
}
func (m *AlarmDeviceData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmDeviceData.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmDeviceData proto.InternalMessageInfo

func (m *AlarmDeviceData) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *AlarmDeviceData) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *AlarmDeviceData) GetLastAlarmSequence() uint32 {
	if m != nil {
		return m.LastAlarmSequence
	}
	return 0
}

func (m *AlarmDeviceData) GetLastSyncTime() string {
	if m != nil {
		return m.LastSyncTime
	}
	return ""
}

func (m *AlarmDeviceData) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *AlarmDeviceData) GetClasses() []*AlarmClassData {
	if m != nil {
		return m.Classes
	}
	return nil
}

func (m *AlarmDeviceData) GetManagedEntities() []*AlarmManagedEntity {
	if m != nil {
		return m.ManagedEntities
	}
	return nil
}

func (m *AlarmDeviceData) GetMessageTypes() []*AlarmMessageType {
	if m != nil {
		return m.MessageTypes
	}
	return nil
}

type AlarmOpenOmciEventType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmOpenOmciEventType) Reset()         { *m = AlarmOpenOmciEventType{} }
func (m *AlarmOpenOmciEventType) String() string { return proto.CompactTextString(m) }
func (*AlarmOpenOmciEventType) ProtoMessage()    {}
func (*AlarmOpenOmciEventType) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d41f1e38aadb08d, []int{6}
}

func (m *AlarmOpenOmciEventType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmOpenOmciEventType.Unmarshal(m, b)
}
func (m *AlarmOpenOmciEventType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmOpenOmciEventType.Marshal(b, m, deterministic)
}
func (m *AlarmOpenOmciEventType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmOpenOmciEventType.Merge(m, src)
}
func (m *AlarmOpenOmciEventType) XXX_Size() int {
	return xxx_messageInfo_AlarmOpenOmciEventType.Size(m)
}
func (m *AlarmOpenOmciEventType) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmOpenOmciEventType.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmOpenOmciEventType proto.InternalMessageInfo

type AlarmOpenOmciEvent struct {
	Type                 AlarmOpenOmciEventType_OpenOmciEventType `protobuf:"varint,1,opt,name=type,proto3,enum=omci.AlarmOpenOmciEventType_OpenOmciEventType" json:"type,omitempty"`
	Data                 string                                   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *AlarmOpenOmciEvent) Reset()         { *m = AlarmOpenOmciEvent{} }
func (m *AlarmOpenOmciEvent) String() string { return proto.CompactTextString(m) }
func (*AlarmOpenOmciEvent) ProtoMessage()    {}
func (*AlarmOpenOmciEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d41f1e38aadb08d, []int{7}
}

func (m *AlarmOpenOmciEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmOpenOmciEvent.Unmarshal(m, b)
}
func (m *AlarmOpenOmciEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmOpenOmciEvent.Marshal(b, m, deterministic)
}
func (m *AlarmOpenOmciEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmOpenOmciEvent.Merge(m, src)
}
func (m *AlarmOpenOmciEvent) XXX_Size() int {
	return xxx_messageInfo_AlarmOpenOmciEvent.Size(m)
}
func (m *AlarmOpenOmciEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmOpenOmciEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmOpenOmciEvent proto.InternalMessageInfo

func (m *AlarmOpenOmciEvent) GetType() AlarmOpenOmciEventType_OpenOmciEventType {
	if m != nil {
		return m.Type
	}
	return AlarmOpenOmciEventType_state_change
}

func (m *AlarmOpenOmciEvent) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterEnum("omci.AlarmOpenOmciEventType_OpenOmciEventType", AlarmOpenOmciEventType_OpenOmciEventType_name, AlarmOpenOmciEventType_OpenOmciEventType_value)
	proto.RegisterType((*AlarmAttributeData)(nil), "omci.AlarmAttributeData")
	proto.RegisterType((*AlarmInstanceData)(nil), "omci.AlarmInstanceData")
	proto.RegisterType((*AlarmClassData)(nil), "omci.AlarmClassData")
	proto.RegisterType((*AlarmManagedEntity)(nil), "omci.AlarmManagedEntity")
	proto.RegisterType((*AlarmMessageType)(nil), "omci.AlarmMessageType")
	proto.RegisterType((*AlarmDeviceData)(nil), "omci.AlarmDeviceData")
	proto.RegisterType((*AlarmOpenOmciEventType)(nil), "omci.AlarmOpenOmciEventType")
	proto.RegisterType((*AlarmOpenOmciEvent)(nil), "omci.AlarmOpenOmciEvent")
}

func init() { proto.RegisterFile("voltha_protos/omci_alarm_db.proto", fileDescriptor_8d41f1e38aadb08d) }

var fileDescriptor_8d41f1e38aadb08d = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x25, 0x6d, 0xda, 0xa6, 0x93, 0xa4, 0x4d, 0x97, 0xaa, 0x2c, 0x95, 0x2a, 0x15, 0x0b, 0x50,
	0x0f, 0xe0, 0x48, 0x45, 0x9c, 0x40, 0xaa, 0x9a, 0x36, 0x12, 0x39, 0xa0, 0x0a, 0xb7, 0x27, 0x2e,
	0xd6, 0xc6, 0x1e, 0xdc, 0x95, 0xbc, 0xeb, 0xe0, 0xdd, 0x58, 0xca, 0x81, 0x0b, 0x5f, 0xc5, 0x4f,
	0xf4, 0x27, 0x38, 0xf1, 0x05, 0x3d, 0x23, 0x8f, 0xed, 0xc4, 0x91, 0x25, 0xc4, 0x6d, 0xdf, 0x9b,
	0x99, 0x37, 0xb3, 0xf3, 0x56, 0x0b, 0x2f, 0xb2, 0x24, 0xb6, 0xf7, 0xc2, 0x9f, 0xa5, 0x89, 0x4d,
	0xcc, 0x30, 0x51, 0x81, 0xf4, 0x45, 0x2c, 0x52, 0xe5, 0x87, 0x53, 0x97, 0x48, 0xd6, 0xce, 0xc9,
	0x63, 0xbe, 0x9e, 0xa8, 0xd0, 0x8a, 0x22, 0xee, 0x8c, 0x81, 0x5d, 0xe6, 0x15, 0x97, 0xd6, 0xa6,
	0x72, 0x3a, 0xb7, 0x78, 0x2d, 0xac, 0x60, 0xcf, 0xa1, 0xad, 0x85, 0x42, 0xde, 0x3a, 0x6d, 0x9d,
	0xed, 0x8e, 0xb6, 0xfe, 0x3c, 0x3e, 0x9c, 0xb4, 0x3c, 0xa2, 0xd8, 0x21, 0x6c, 0x65, 0x22, 0x9e,
	0x23, 0xdf, 0xc8, 0x63, 0x5e, 0x01, 0x9c, 0x5f, 0x2d, 0x38, 0x20, 0x9d, 0x89, 0x36, 0x56, 0xe8,
	0xa0, 0x90, 0x79, 0x0d, 0x5d, 0x59, 0x62, 0x5f, 0x86, 0xa4, 0xd6, 0xaf, 0xd4, 0xa0, 0x8a, 0x4c,
	0x42, 0xc6, 0x61, 0x27, 0x48, 0x51, 0x58, 0x0c, 0x4b, 0xd5, 0x0a, 0xb2, 0x63, 0xe8, 0xa8, 0x24,
	0x94, 0xdf, 0x24, 0x86, 0x7c, 0x93, 0x42, 0x4b, 0xcc, 0xc6, 0x00, 0xa2, 0x9a, 0xda, 0xf0, 0xf6,
	0xe9, 0xe6, 0x59, 0xf7, 0x9c, 0xbb, 0xf9, 0x7d, 0xdd, 0xe6, 0x95, 0x46, 0xdd, 0xdf, 0x8f, 0x0f,
	0x27, 0xdb, 0xc5, 0xbd, 0xbc, 0x5a, 0xa1, 0xf3, 0x03, 0xf6, 0x28, 0xfd, 0x2a, 0x16, 0xc6, 0xd0,
	0xd8, 0xa7, 0xd0, 0x09, 0x72, 0xd0, 0x98, 0x79, 0x87, 0xe8, 0x49, 0xc8, 0x26, 0xb0, 0x5b, 0x8d,
	0x6f, 0xf8, 0x06, 0x75, 0x7e, 0x56, 0xeb, 0x5c, 0x5f, 0xc2, 0x88, 0xe5, 0x8d, 0xfb, 0x6b, 0x9b,
	0xf0, 0x56, 0xd5, 0xce, 0x97, 0xd2, 0x80, 0xcf, 0x42, 0x8b, 0x08, 0xc3, 0xb1, 0xb6, 0xd2, 0x2e,
	0xfe, 0x63, 0x84, 0xca, 0xa2, 0x8d, 0x86, 0x45, 0xce, 0x47, 0x18, 0x14, 0x92, 0x68, 0x8c, 0x88,
	0xf0, 0x6e, 0x31, 0x43, 0x76, 0x06, 0x3d, 0x55, 0x40, 0xdf, 0x2e, 0x66, 0xb8, 0x2e, 0xda, 0x55,
	0xab, 0x4c, 0xe7, 0xe7, 0x26, 0xec, 0x53, 0xf9, 0x35, 0x66, 0xb2, 0x34, 0xd2, 0x81, 0xdd, 0x90,
	0x50, 0x35, 0xcf, 0xb2, 0x63, 0xa7, 0xe0, 0xff, 0x69, 0xa2, 0x0b, 0x4f, 0x63, 0x61, 0x6c, 0xf9,
	0x34, 0x0d, 0x7e, 0x9f, 0xa3, 0x0e, 0x90, 0xfc, 0xec, 0x7b, 0x07, 0x79, 0x88, 0xfa, 0xdd, 0x96,
	0x01, 0xf6, 0x12, 0xf6, 0x28, 0xdf, 0x2c, 0x74, 0xe0, 0x5b, 0xa9, 0x90, 0xb7, 0x49, 0xb0, 0x97,
	0xb3, 0xb7, 0x0b, 0x1d, 0xdc, 0x49, 0x85, 0x79, 0xbf, 0x0c, 0x53, 0x23, 0x13, 0xcd, 0xb7, 0x48,
	0xa9, 0x82, 0xec, 0x02, 0x8a, 0x2d, 0xa1, 0xe1, 0xdb, 0xe4, 0xcd, 0x61, 0xcd, 0x9b, 0xa5, 0xcd,
	0xa3, 0xfd, 0xdc, 0x18, 0x58, 0x2d, 0xda, 0xab, 0xaa, 0xd8, 0x15, 0x0c, 0x54, 0x61, 0x87, 0x8f,
	0xb9, 0x1f, 0x12, 0x0d, 0xdf, 0x69, 0xbc, 0xaf, 0x35, 0xc7, 0xbc, 0x7d, 0x55, 0x83, 0x12, 0x0d,
	0xfb, 0x00, 0xfd, 0xfa, 0xc6, 0x0d, 0xef, 0x90, 0xc2, 0x51, 0x5d, 0x61, 0xb5, 0x76, 0xaf, 0x57,
	0xf3, 0xc0, 0x38, 0x17, 0x70, 0x44, 0x19, 0x37, 0x33, 0xd4, 0x37, 0x2a, 0x90, 0xe3, 0x0c, 0xb5,
	0x25, 0x7b, 0x5e, 0xc1, 0x41, 0x83, 0x64, 0x03, 0xe8, 0x19, 0x2b, 0x2c, 0xfa, 0xc1, 0xbd, 0xd0,
	0x11, 0x0e, 0x9e, 0x38, 0x71, 0xf9, 0xac, 0xd6, 0x72, 0xd9, 0x08, 0xda, 0x4b, 0xf7, 0xf7, 0xce,
	0xdd, 0xda, 0x28, 0x0d, 0x4d, 0xb7, 0xc1, 0x78, 0x54, 0xcb, 0x18, 0xb4, 0x43, 0x61, 0x45, 0x69,
	0x32, 0x9d, 0x47, 0x9f, 0x80, 0x27, 0x69, 0xe4, 0x26, 0x33, 0xd4, 0x41, 0x92, 0x86, 0x6e, 0xf1,
	0xdd, 0x90, 0xfc, 0xd7, 0x37, 0x91, 0xb4, 0xf7, 0xf3, 0xa9, 0x1b, 0x24, 0x6a, 0x58, 0x25, 0x0c,
	0x8b, 0x84, 0xb7, 0xe5, 0x7f, 0x94, 0xbd, 0x1f, 0x46, 0x09, 0x7d, 0x5f, 0xd3, 0x6d, 0xa2, 0xde,
	0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x71, 0x53, 0xfa, 0xbd, 0xdb, 0x04, 0x00, 0x00,
}