// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coordinator/coordinator.proto

package coordinator

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
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

type Resource int32

const (
	Resource_WBLIST                          Resource = 0
	Resource_CONFIG                          Resource = 1
	Resource_CACHE                           Resource = 2
	Resource_POLICY_TCKFEE                   Resource = 3
	Resource_POLICY_MARKUP                   Resource = 4
	Resource_POLICY_INCENTIVE                Resource = 5
	Resource_POLICY_PLATFORM                 Resource = 6
	Resource_POLICY_VCC                      Resource = 7
	Resource_POLICY_COMMISSION               Resource = 8
	Resource_POLICY_BACKEND                  Resource = 9
	Resource_POLICY_UNIMARKUP                Resource = 10
	Resource_PLATFORM_CONFIG                 Resource = 11
	Resource_PROVIDER_CONFIG                 Resource = 12
	Resource_PLATFORM_PROVIDER_BINDING       Resource = 13
	Resource_PLATFORM_GROUP                  Resource = 14
	Resource_PROVIDER_GROUP                  Resource = 15
	Resource_PRICE_INTERCEPT                 Resource = 16
	Resource_POLICY_FARERULE                 Resource = 17
	Resource_CURRENCY_RULE                   Resource = 18
	Resource_FARE_FILTER                     Resource = 19
	Resource_POLICY_MANUAL_RULE              Resource = 20
	Resource_POLICY_AI                       Resource = 21
	Resource_POLICY_MAIRI_PLATFORM           Resource = 22
	Resource_POLICY_MAIRI_CURRENCY           Resource = 23
	Resource_POLICY_MAIRI_PAYMENT            Resource = 24
	Resource_AIRLINE_CONFIG                  Resource = 25
	Resource_POLICY_MAIRI_PAYMENT_METHOD_FEE Resource = 26
)

var Resource_name = map[int32]string{
	0:  "WBLIST",
	1:  "CONFIG",
	2:  "CACHE",
	3:  "POLICY_TCKFEE",
	4:  "POLICY_MARKUP",
	5:  "POLICY_INCENTIVE",
	6:  "POLICY_PLATFORM",
	7:  "POLICY_VCC",
	8:  "POLICY_COMMISSION",
	9:  "POLICY_BACKEND",
	10: "POLICY_UNIMARKUP",
	11: "PLATFORM_CONFIG",
	12: "PROVIDER_CONFIG",
	13: "PLATFORM_PROVIDER_BINDING",
	14: "PLATFORM_GROUP",
	15: "PROVIDER_GROUP",
	16: "PRICE_INTERCEPT",
	17: "POLICY_FARERULE",
	18: "CURRENCY_RULE",
	19: "FARE_FILTER",
	20: "POLICY_MANUAL_RULE",
	21: "POLICY_AI",
	22: "POLICY_MAIRI_PLATFORM",
	23: "POLICY_MAIRI_CURRENCY",
	24: "POLICY_MAIRI_PAYMENT",
	25: "AIRLINE_CONFIG",
	26: "POLICY_MAIRI_PAYMENT_METHOD_FEE",
}

var Resource_value = map[string]int32{
	"WBLIST":                          0,
	"CONFIG":                          1,
	"CACHE":                           2,
	"POLICY_TCKFEE":                   3,
	"POLICY_MARKUP":                   4,
	"POLICY_INCENTIVE":                5,
	"POLICY_PLATFORM":                 6,
	"POLICY_VCC":                      7,
	"POLICY_COMMISSION":               8,
	"POLICY_BACKEND":                  9,
	"POLICY_UNIMARKUP":                10,
	"PLATFORM_CONFIG":                 11,
	"PROVIDER_CONFIG":                 12,
	"PLATFORM_PROVIDER_BINDING":       13,
	"PLATFORM_GROUP":                  14,
	"PROVIDER_GROUP":                  15,
	"PRICE_INTERCEPT":                 16,
	"POLICY_FARERULE":                 17,
	"CURRENCY_RULE":                   18,
	"FARE_FILTER":                     19,
	"POLICY_MANUAL_RULE":              20,
	"POLICY_AI":                       21,
	"POLICY_MAIRI_PLATFORM":           22,
	"POLICY_MAIRI_CURRENCY":           23,
	"POLICY_MAIRI_PAYMENT":            24,
	"AIRLINE_CONFIG":                  25,
	"POLICY_MAIRI_PAYMENT_METHOD_FEE": 26,
}

func (x Resource) String() string {
	return proto.EnumName(Resource_name, int32(x))
}

func (Resource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b3b4a5a824774e2, []int{0}
}

type Action int32

const (
	Action_ADD    Action = 0
	Action_CHANGE Action = 1
	Action_DELETE Action = 2
	Action_ANY    Action = 3
)

var Action_name = map[int32]string{
	0: "ADD",
	1: "CHANGE",
	2: "DELETE",
	3: "ANY",
}

var Action_value = map[string]int32{
	"ADD":    0,
	"CHANGE": 1,
	"DELETE": 2,
	"ANY":    3,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b3b4a5a824774e2, []int{1}
}

type ResponseStatus int32

const (
	ResponseStatus_SUCCESS ResponseStatus = 0
	ResponseStatus_FAILED  ResponseStatus = 1
)

var ResponseStatus_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILED",
}

var ResponseStatus_value = map[string]int32{
	"SUCCESS": 0,
	"FAILED":  1,
}

func (x ResponseStatus) String() string {
	return proto.EnumName(ResponseStatus_name, int32(x))
}

func (ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b3b4a5a824774e2, []int{2}
}

type Event struct {
	Resource             Resource             `protobuf:"varint,1,opt,name=resource,proto3,enum=com.gloryholiday.grpc.coordinator.Resource" json:"resource,omitempty"`
	Action               Action               `protobuf:"varint,2,opt,name=action,proto3,enum=com.gloryholiday.grpc.coordinator.Action" json:"action,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TraceId              string               `protobuf:"bytes,4,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b3b4a5a824774e2, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetResource() Resource {
	if m != nil {
		return m.Resource
	}
	return Resource_WBLIST
}

func (m *Event) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_ADD
}

func (m *Event) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Event) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

type Subscriber struct {
	Event                *Event               `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Identifier           string               `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Subscriber) Reset()         { *m = Subscriber{} }
func (m *Subscriber) String() string { return proto.CompactTextString(m) }
func (*Subscriber) ProtoMessage()    {}
func (*Subscriber) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b3b4a5a824774e2, []int{1}
}

func (m *Subscriber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscriber.Unmarshal(m, b)
}
func (m *Subscriber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscriber.Marshal(b, m, deterministic)
}
func (m *Subscriber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscriber.Merge(m, src)
}
func (m *Subscriber) XXX_Size() int {
	return xxx_messageInfo_Subscriber.Size(m)
}
func (m *Subscriber) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscriber.DiscardUnknown(m)
}

var xxx_messageInfo_Subscriber proto.InternalMessageInfo

func (m *Subscriber) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Subscriber) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Subscriber) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type NotificationResponse struct {
	Status               ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=com.gloryholiday.grpc.coordinator.ResponseStatus" json:"status,omitempty"`
	Msg                  string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NotificationResponse) Reset()         { *m = NotificationResponse{} }
func (m *NotificationResponse) String() string { return proto.CompactTextString(m) }
func (*NotificationResponse) ProtoMessage()    {}
func (*NotificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b3b4a5a824774e2, []int{2}
}

func (m *NotificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationResponse.Unmarshal(m, b)
}
func (m *NotificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationResponse.Marshal(b, m, deterministic)
}
func (m *NotificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationResponse.Merge(m, src)
}
func (m *NotificationResponse) XXX_Size() int {
	return xxx_messageInfo_NotificationResponse.Size(m)
}
func (m *NotificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationResponse proto.InternalMessageInfo

func (m *NotificationResponse) GetStatus() ResponseStatus {
	if m != nil {
		return m.Status
	}
	return ResponseStatus_SUCCESS
}

func (m *NotificationResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("com.gloryholiday.grpc.coordinator.Resource", Resource_name, Resource_value)
	proto.RegisterEnum("com.gloryholiday.grpc.coordinator.Action", Action_name, Action_value)
	proto.RegisterEnum("com.gloryholiday.grpc.coordinator.ResponseStatus", ResponseStatus_name, ResponseStatus_value)
	proto.RegisterType((*Event)(nil), "com.gloryholiday.grpc.coordinator.Event")
	proto.RegisterType((*Subscriber)(nil), "com.gloryholiday.grpc.coordinator.Subscriber")
	proto.RegisterType((*NotificationResponse)(nil), "com.gloryholiday.grpc.coordinator.NotificationResponse")
}

func init() { proto.RegisterFile("coordinator/coordinator.proto", fileDescriptor_1b3b4a5a824774e2) }

var fileDescriptor_1b3b4a5a824774e2 = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xe3, 0x36,
	0x10, 0x8d, 0xd6, 0x89, 0x13, 0x8f, 0x1b, 0x87, 0x99, 0x38, 0x5b, 0xc7, 0xc0, 0x76, 0x53, 0xf7,
	0x92, 0x4d, 0x51, 0xa5, 0x75, 0x0f, 0xed, 0xa9, 0x80, 0x4c, 0xd1, 0x0e, 0x11, 0x99, 0x32, 0x28,
	0x39, 0x45, 0x7a, 0x11, 0x1c, 0x59, 0xeb, 0x15, 0x10, 0x5b, 0x86, 0xa4, 0x2c, 0x90, 0x1f, 0xea,
	0x67, 0xf4, 0x7b, 0x7a, 0xea, 0x37, 0x14, 0xa2, 0x24, 0x5b, 0x8b, 0x5d, 0xa0, 0x46, 0x6f, 0xe4,
	0xe3, 0x7b, 0x33, 0x7c, 0x43, 0xce, 0xc0, 0x1b, 0x3f, 0x8a, 0xe2, 0x79, 0xb8, 0x9a, 0xa5, 0x51,
	0x7c, 0x53, 0x59, 0xeb, 0xeb, 0x38, 0x4a, 0x23, 0xfc, 0xd6, 0x8f, 0x96, 0xfa, 0xe2, 0x29, 0x8a,
	0x5f, 0x3e, 0x44, 0x4f, 0xe1, 0x7c, 0xf6, 0xa2, 0x2f, 0xe2, 0xb5, 0xaf, 0x57, 0x88, 0xdd, 0xb7,
	0x8b, 0x28, 0x5a, 0x3c, 0x05, 0x37, 0x4a, 0xf0, 0xf8, 0xfc, 0xfe, 0x26, 0x0d, 0x97, 0x41, 0x92,
	0xce, 0x96, 0xeb, 0x3c, 0x46, 0xef, 0x6f, 0x0d, 0x0e, 0xd8, 0xc7, 0x60, 0x95, 0xe2, 0x08, 0x8e,
	0xe2, 0x20, 0x89, 0x9e, 0x63, 0x3f, 0xe8, 0x68, 0x97, 0xda, 0x55, 0xab, 0xff, 0xbd, 0xfe, 0x9f,
	0x09, 0x74, 0x59, 0x48, 0xe4, 0x46, 0x8c, 0x06, 0xd4, 0x67, 0x7e, 0x1a, 0x46, 0xab, 0xce, 0x2b,
	0x15, 0xe6, 0xdd, 0x0e, 0x61, 0x0c, 0x25, 0x90, 0x85, 0x10, 0x7f, 0x85, 0xc6, 0xe6, 0xa2, 0x9d,
	0xda, 0xa5, 0x76, 0xd5, 0xec, 0x77, 0xf5, 0xdc, 0x8a, 0x5e, 0x5a, 0xd1, 0xdd, 0x92, 0x21, 0xb7,
	0x64, 0xbc, 0x80, 0xa3, 0x34, 0x9e, 0xf9, 0x81, 0x17, 0xce, 0x3b, 0xfb, 0x97, 0xda, 0x55, 0x43,
	0x1e, 0xaa, 0x3d, 0x9f, 0xf7, 0xfe, 0xd4, 0x00, 0x9c, 0xe7, 0xc7, 0xc4, 0x8f, 0xc3, 0xc7, 0x20,
	0xc6, 0xdf, 0xe0, 0x20, 0xc8, 0x8c, 0x2b, 0xb3, 0xcd, 0xfe, 0xd5, 0x0e, 0xb7, 0x54, 0x85, 0x92,
	0xb9, 0x0c, 0xbf, 0x01, 0x08, 0xe7, 0xc1, 0x2a, 0x0d, 0xdf, 0x87, 0x41, 0xac, 0xac, 0x36, 0x64,
	0x05, 0xf9, 0xff, 0x1e, 0x7a, 0x09, 0xb4, 0x45, 0x94, 0x45, 0xf1, 0x67, 0xaa, 0x2a, 0x41, 0xb2,
	0x8e, 0x56, 0x49, 0x80, 0x1c, 0xea, 0x49, 0x3a, 0x4b, 0x9f, 0x93, 0xe2, 0x7d, 0x7e, 0xda, 0xed,
	0x7d, 0x94, 0xd8, 0x51, 0x42, 0x59, 0x04, 0x40, 0x02, 0xb5, 0x65, 0xb2, 0x28, 0x6e, 0x9d, 0x2d,
	0xaf, 0xff, 0xda, 0x87, 0xa3, 0xf2, 0x31, 0x11, 0xa0, 0xfe, 0xfb, 0xc0, 0xe2, 0x8e, 0x4b, 0xf6,
	0xb2, 0x35, 0xb5, 0xc5, 0x90, 0x8f, 0x88, 0x86, 0x0d, 0x38, 0xa0, 0x06, 0xbd, 0x65, 0xe4, 0x15,
	0x9e, 0xc2, 0xf1, 0xc4, 0xb6, 0x38, 0x7d, 0xf0, 0x5c, 0x7a, 0x37, 0x64, 0x8c, 0xd4, 0x2a, 0xd0,
	0xd8, 0x90, 0x77, 0xd3, 0x09, 0xd9, 0xc7, 0x36, 0x90, 0x02, 0xe2, 0x82, 0x32, 0xe1, 0xf2, 0x7b,
	0x46, 0x0e, 0xf0, 0x0c, 0x4e, 0x0a, 0x74, 0x62, 0x19, 0xee, 0xd0, 0x96, 0x63, 0x52, 0xc7, 0x16,
	0x40, 0x01, 0xde, 0x53, 0x4a, 0x0e, 0xf1, 0x1c, 0x4e, 0x8b, 0x3d, 0xb5, 0xc7, 0x63, 0xee, 0x38,
	0xdc, 0x16, 0xe4, 0x08, 0x11, 0x5a, 0x05, 0x3c, 0x30, 0xe8, 0x1d, 0x13, 0x26, 0x69, 0x54, 0xb2,
	0x4c, 0x05, 0x2f, 0x72, 0x83, 0xca, 0x52, 0x84, 0xf7, 0x0a, 0x07, 0x4d, 0x05, 0x4a, 0xfb, 0x9e,
	0x9b, 0x4c, 0x96, 0xe0, 0x57, 0xf8, 0x06, 0x2e, 0x36, 0xcc, 0xcd, 0xe9, 0x80, 0x0b, 0x93, 0x8b,
	0x11, 0x39, 0x56, 0x29, 0xcb, 0xe3, 0x91, 0xb4, 0xa7, 0x13, 0xd2, 0x52, 0x58, 0xc9, 0xcc, 0xb1,
	0x93, 0x3c, 0x36, 0xa7, 0xcc, 0xe3, 0xc2, 0x65, 0x92, 0xb2, 0x89, 0x4b, 0x48, 0xc5, 0xeb, 0xd0,
	0x90, 0x4c, 0x4e, 0x2d, 0x46, 0x4e, 0xb3, 0x4a, 0xd1, 0xa9, 0x94, 0x4c, 0xd0, 0x07, 0x4f, 0x41,
	0x88, 0x27, 0xd0, 0xcc, 0x08, 0xde, 0x90, 0x5b, 0x2e, 0x93, 0xe4, 0x0c, 0x5f, 0x03, 0x6e, 0xaa,
	0x29, 0xa6, 0x86, 0x95, 0x13, 0xdb, 0x78, 0x0c, 0x8d, 0x02, 0x37, 0x38, 0x39, 0xc7, 0x0b, 0x38,
	0xdf, 0xd0, 0xb8, 0xe4, 0xdb, 0x8a, 0xbe, 0xfe, 0xec, 0xa8, 0x4c, 0x49, 0xbe, 0xc6, 0x0e, 0xb4,
	0x3f, 0x55, 0x19, 0x0f, 0x63, 0x26, 0x5c, 0xd2, 0xc9, 0x8c, 0x19, 0x5c, 0x5a, 0x5c, 0xb0, 0xb2,
	0x3e, 0x17, 0xf8, 0x1d, 0xbc, 0xfd, 0x12, 0xdb, 0x1b, 0x33, 0xf7, 0xd6, 0x36, 0xbd, 0xec, 0xf5,
	0xbb, 0xd7, 0x7d, 0xa8, 0xe7, 0x5d, 0x8c, 0x87, 0x50, 0x33, 0x4c, 0xb3, 0xf8, 0x3a, 0xb7, 0x86,
	0x18, 0x31, 0xa2, 0x65, 0x6b, 0x93, 0x59, 0xcc, 0xcd, 0xfe, 0x4e, 0x46, 0x10, 0x0f, 0xa4, 0x76,
	0xfd, 0x0e, 0x5a, 0x9f, 0x7e, 0x50, 0x6c, 0xc2, 0xa1, 0x33, 0xa5, 0x94, 0x39, 0x4e, 0xae, 0x1f,
	0x1a, 0xdc, 0x62, 0x26, 0xd1, 0xfa, 0xff, 0x68, 0x70, 0x56, 0xed, 0x0a, 0x27, 0x88, 0x3f, 0x86,
	0x7e, 0x80, 0x1f, 0xa0, 0xb1, 0x69, 0x6a, 0xfc, 0x61, 0x87, 0x8e, 0xd8, 0x8e, 0x80, 0xee, 0xce,
	0x3d, 0xdf, 0xdb, 0xfb, 0x51, 0xc3, 0x25, 0xd4, 0xd5, 0x05, 0x5e, 0x70, 0x67, 0x5d, 0xf7, 0x97,
	0x1d, 0x98, 0x5f, 0xea, 0xf5, 0xde, 0xde, 0xe0, 0x16, 0x3e, 0x9f, 0xef, 0x7e, 0x14, 0x07, 0x55,
	0xed, 0xa0, 0x3d, 0x8a, 0xd7, 0x3e, 0xdd, 0x02, 0x93, 0x6c, 0xb0, 0x4c, 0xb4, 0x3f, 0x9a, 0x15,
	0xd2, 0x63, 0x5d, 0x8d, 0x9b, 0x9f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x9b, 0x0a, 0xf1,
	0x4f, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	Subscribe(ctx context.Context, in *Subscriber, opts ...grpc.CallOption) (NotificationService_SubscribeClient, error)
	Notify(ctx context.Context, in *Event, opts ...grpc.CallOption) (*NotificationResponse, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Subscribe(ctx context.Context, in *Subscriber, opts ...grpc.CallOption) (NotificationService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NotificationService_serviceDesc.Streams[0], "/com.gloryholiday.grpc.coordinator.NotificationService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_SubscribeClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type notificationServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *notificationServiceSubscribeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notificationServiceClient) Notify(ctx context.Context, in *Event, opts ...grpc.CallOption) (*NotificationResponse, error) {
	out := new(NotificationResponse)
	err := c.cc.Invoke(ctx, "/com.gloryholiday.grpc.coordinator.NotificationService/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	Subscribe(*Subscriber, NotificationService_SubscribeServer) error
	Notify(context.Context, *Event) (*NotificationResponse, error)
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Subscriber)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).Subscribe(m, &notificationServiceSubscribeServer{stream})
}

type NotificationService_SubscribeServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type notificationServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *notificationServiceSubscribeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _NotificationService_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.gloryholiday.grpc.coordinator.NotificationService/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Notify(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.gloryholiday.grpc.coordinator.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _NotificationService_Notify_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _NotificationService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "coordinator/coordinator.proto",
}