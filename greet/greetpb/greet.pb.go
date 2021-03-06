// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/greetpb/greet.proto

package greetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetManyTimesRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetManyTimesRequest) Reset()         { *m = GreetManyTimesRequest{} }
func (m *GreetManyTimesRequest) String() string { return proto.CompactTextString(m) }
func (*GreetManyTimesRequest) ProtoMessage()    {}
func (*GreetManyTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{3}
}

func (m *GreetManyTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManyTimesRequest.Unmarshal(m, b)
}
func (m *GreetManyTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManyTimesRequest.Marshal(b, m, deterministic)
}
func (m *GreetManyTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManyTimesRequest.Merge(m, src)
}
func (m *GreetManyTimesRequest) XXX_Size() int {
	return xxx_messageInfo_GreetManyTimesRequest.Size(m)
}
func (m *GreetManyTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManyTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManyTimesRequest proto.InternalMessageInfo

func (m *GreetManyTimesRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetmanytimesResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetmanytimesResponse) Reset()         { *m = GreetmanytimesResponse{} }
func (m *GreetmanytimesResponse) String() string { return proto.CompactTextString(m) }
func (*GreetmanytimesResponse) ProtoMessage()    {}
func (*GreetmanytimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{4}
}

func (m *GreetmanytimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetmanytimesResponse.Unmarshal(m, b)
}
func (m *GreetmanytimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetmanytimesResponse.Marshal(b, m, deterministic)
}
func (m *GreetmanytimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetmanytimesResponse.Merge(m, src)
}
func (m *GreetmanytimesResponse) XXX_Size() int {
	return xxx_messageInfo_GreetmanytimesResponse.Size(m)
}
func (m *GreetmanytimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetmanytimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetmanytimesResponse proto.InternalMessageInfo

func (m *GreetmanytimesResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type LongGreatRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LongGreatRequest) Reset()         { *m = LongGreatRequest{} }
func (m *LongGreatRequest) String() string { return proto.CompactTextString(m) }
func (*LongGreatRequest) ProtoMessage()    {}
func (*LongGreatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{5}
}

func (m *LongGreatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreatRequest.Unmarshal(m, b)
}
func (m *LongGreatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreatRequest.Marshal(b, m, deterministic)
}
func (m *LongGreatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreatRequest.Merge(m, src)
}
func (m *LongGreatRequest) XXX_Size() int {
	return xxx_messageInfo_LongGreatRequest.Size(m)
}
func (m *LongGreatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreatRequest proto.InternalMessageInfo

func (m *LongGreatRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type LongGreatResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongGreatResponse) Reset()         { *m = LongGreatResponse{} }
func (m *LongGreatResponse) String() string { return proto.CompactTextString(m) }
func (*LongGreatResponse) ProtoMessage()    {}
func (*LongGreatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{6}
}

func (m *LongGreatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreatResponse.Unmarshal(m, b)
}
func (m *LongGreatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreatResponse.Marshal(b, m, deterministic)
}
func (m *LongGreatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreatResponse.Merge(m, src)
}
func (m *LongGreatResponse) XXX_Size() int {
	return xxx_messageInfo_LongGreatResponse.Size(m)
}
func (m *LongGreatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreatResponse proto.InternalMessageInfo

func (m *LongGreatResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreatEveryoneRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreatEveryoneRequest) Reset()         { *m = GreatEveryoneRequest{} }
func (m *GreatEveryoneRequest) String() string { return proto.CompactTextString(m) }
func (*GreatEveryoneRequest) ProtoMessage()    {}
func (*GreatEveryoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{7}
}

func (m *GreatEveryoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreatEveryoneRequest.Unmarshal(m, b)
}
func (m *GreatEveryoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreatEveryoneRequest.Marshal(b, m, deterministic)
}
func (m *GreatEveryoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreatEveryoneRequest.Merge(m, src)
}
func (m *GreatEveryoneRequest) XXX_Size() int {
	return xxx_messageInfo_GreatEveryoneRequest.Size(m)
}
func (m *GreatEveryoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreatEveryoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreatEveryoneRequest proto.InternalMessageInfo

func (m *GreatEveryoneRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreatEveryoneResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreatEveryoneResponse) Reset()         { *m = GreatEveryoneResponse{} }
func (m *GreatEveryoneResponse) String() string { return proto.CompactTextString(m) }
func (*GreatEveryoneResponse) ProtoMessage()    {}
func (*GreatEveryoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{8}
}

func (m *GreatEveryoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreatEveryoneResponse.Unmarshal(m, b)
}
func (m *GreatEveryoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreatEveryoneResponse.Marshal(b, m, deterministic)
}
func (m *GreatEveryoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreatEveryoneResponse.Merge(m, src)
}
func (m *GreatEveryoneResponse) XXX_Size() int {
	return xxx_messageInfo_GreatEveryoneResponse.Size(m)
}
func (m *GreatEveryoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreatEveryoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreatEveryoneResponse proto.InternalMessageInfo

func (m *GreatEveryoneResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetWithDeadlineRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetWithDeadlineRequest) Reset()         { *m = GreetWithDeadlineRequest{} }
func (m *GreetWithDeadlineRequest) String() string { return proto.CompactTextString(m) }
func (*GreetWithDeadlineRequest) ProtoMessage()    {}
func (*GreetWithDeadlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{9}
}

func (m *GreetWithDeadlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDeadlineRequest.Unmarshal(m, b)
}
func (m *GreetWithDeadlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDeadlineRequest.Marshal(b, m, deterministic)
}
func (m *GreetWithDeadlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDeadlineRequest.Merge(m, src)
}
func (m *GreetWithDeadlineRequest) XXX_Size() int {
	return xxx_messageInfo_GreetWithDeadlineRequest.Size(m)
}
func (m *GreetWithDeadlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDeadlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDeadlineRequest proto.InternalMessageInfo

func (m *GreetWithDeadlineRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetWithDeadlineResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetWithDeadlineResponse) Reset()         { *m = GreetWithDeadlineResponse{} }
func (m *GreetWithDeadlineResponse) String() string { return proto.CompactTextString(m) }
func (*GreetWithDeadlineResponse) ProtoMessage()    {}
func (*GreetWithDeadlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{10}
}

func (m *GreetWithDeadlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDeadlineResponse.Unmarshal(m, b)
}
func (m *GreetWithDeadlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDeadlineResponse.Marshal(b, m, deterministic)
}
func (m *GreetWithDeadlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDeadlineResponse.Merge(m, src)
}
func (m *GreetWithDeadlineResponse) XXX_Size() int {
	return xxx_messageInfo_GreetWithDeadlineResponse.Size(m)
}
func (m *GreetWithDeadlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDeadlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDeadlineResponse proto.InternalMessageInfo

func (m *GreetWithDeadlineResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*GreetManyTimesRequest)(nil), "greet.GreetManyTimesRequest")
	proto.RegisterType((*GreetmanytimesResponse)(nil), "greet.GreetmanytimesResponse")
	proto.RegisterType((*LongGreatRequest)(nil), "greet.LongGreatRequest")
	proto.RegisterType((*LongGreatResponse)(nil), "greet.LongGreatResponse")
	proto.RegisterType((*GreatEveryoneRequest)(nil), "greet.GreatEveryoneRequest")
	proto.RegisterType((*GreatEveryoneResponse)(nil), "greet.GreatEveryoneResponse")
	proto.RegisterType((*GreetWithDeadlineRequest)(nil), "greet.GreetWithDeadlineRequest")
	proto.RegisterType((*GreetWithDeadlineResponse)(nil), "greet.GreetWithDeadlineResponse")
}

func init() {
	proto.RegisterFile("greet/greetpb/greet.proto", fileDescriptor_fe6f881da19a2871)
}

var fileDescriptor_fe6f881da19a2871 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xe9, 0xfb, 0x06, 0xa4, 0xe3, 0x5f, 0x56, 0x44, 0x28, 0x10, 0x49, 0x2f, 0x92, 0x90,
	0x00, 0x01, 0x6f, 0x1e, 0x4c, 0x10, 0xe5, 0xa2, 0x46, 0xd1, 0x44, 0xe3, 0xc5, 0x14, 0x1d, 0x6b,
	0x13, 0x68, 0xb1, 0x5d, 0x48, 0xf8, 0x04, 0x7e, 0x6d, 0xcb, 0x74, 0x2b, 0x85, 0xa2, 0x4d, 0x7a,
	0x81, 0xee, 0xcc, 0xec, 0xef, 0x99, 0xdd, 0x79, 0xb2, 0x50, 0xd0, 0x6d, 0x44, 0xde, 0xa0, 0xdf,
	0xf1, 0xc0, 0xfb, 0xaf, 0x8f, 0x6d, 0x8b, 0x5b, 0x2c, 0x49, 0x0b, 0xf5, 0x12, 0xd2, 0xbd, 0xf9,
	0x87, 0x61, 0xea, 0xac, 0x0c, 0xf0, 0x6e, 0xd8, 0x0e, 0x7f, 0x31, 0xb5, 0x11, 0xe6, 0xa5, 0x8a,
	0x54, 0x95, 0xfb, 0x32, 0x45, 0x6e, 0xdc, 0x00, 0x2b, 0x82, 0x3c, 0xd4, 0xfc, 0xec, 0x3f, 0xca,
	0xa6, 0xe7, 0x81, 0x79, 0x52, 0x3d, 0x85, 0x2d, 0xe2, 0xf4, 0xf1, 0x73, 0x82, 0x0e, 0x67, 0x35,
	0x48, 0xeb, 0x82, 0x4b, 0xa4, 0xcd, 0xd6, 0x6e, 0xdd, 0x93, 0xf7, 0xe5, 0xfa, 0x3f, 0x05, 0xea,
	0x31, 0x6c, 0x8b, 0xcd, 0xce, 0xd8, 0x32, 0x1d, 0x64, 0x39, 0x48, 0xd9, 0xe8, 0x4c, 0x86, 0x5c,
	0x74, 0x21, 0x56, 0x6a, 0x17, 0x0e, 0xa8, 0xf0, 0x5a, 0x33, 0x67, 0x0f, 0xc6, 0x08, 0x9d, 0x58,
	0x72, 0x4d, 0xc8, 0x51, 0x74, 0xe4, 0x52, 0xb8, 0x47, 0x89, 0xd0, 0x3d, 0x83, 0xbd, 0x2b, 0xcb,
	0xd4, 0xdd, 0x5d, 0x5a, 0xbc, 0x13, 0xd6, 0x20, 0x13, 0x00, 0x44, 0xa8, 0x9d, 0x43, 0x96, 0x0a,
	0x2f, 0xa6, 0x68, 0xcf, 0x2c, 0x13, 0x63, 0x29, 0x36, 0xe8, 0xaa, 0x82, 0x90, 0x08, 0xd5, 0x1e,
	0xe4, 0x09, 0xf3, 0x68, 0xf0, 0x8f, 0x2e, 0x6a, 0x6f, 0x43, 0x23, 0xa6, 0x72, 0x1b, 0x0a, 0x6b,
	0x40, 0x7f, 0xab, 0xb7, 0xbe, 0xfe, 0x0b, 0x03, 0xdd, 0xa3, 0x3d, 0x35, 0x5e, 0x91, 0x9d, 0x40,
	0x92, 0xd6, 0x6c, 0x3f, 0xa8, 0x24, 0x1a, 0x52, 0xb2, 0xcb, 0x41, 0x0f, 0xae, 0x26, 0xd8, 0x1d,
	0xec, 0x2c, 0x1b, 0x84, 0x95, 0x82, 0x95, 0xab, 0xbe, 0x51, 0xca, 0xc1, 0x6c, 0xc8, 0x0f, 0x6a,
	0xa2, 0x29, 0xb1, 0x0e, 0xc8, 0x62, 0x74, 0x6e, 0x33, 0x87, 0xa2, 0x7e, 0xd5, 0x0d, 0x4a, 0x3e,
	0x9c, 0xf0, 0x19, 0x55, 0x89, 0xdd, 0x92, 0xc1, 0x17, 0xc3, 0x60, 0xc5, 0x85, 0x6e, 0x68, 0xce,
	0x4a, 0x69, 0x7d, 0x72, 0xc1, 0x73, 0xbb, 0x7a, 0x82, 0x4c, 0xe8, 0x92, 0xd9, 0x51, 0xf0, 0x34,
	0x6b, 0xe6, 0xa8, 0x54, 0x7e, 0x2f, 0xf0, 0xe9, 0x1d, 0xf9, 0x79, 0x43, 0xbc, 0x17, 0x83, 0x14,
	0x3d, 0x15, 0xed, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0x49, 0xe9, 0xff, 0x47, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// Unary
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	// Server Streaming
	GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
	// Client Streaming
	LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error)
	// BiDi Streaming
	GreatEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreatEveryoneClient, error)
	// Unary
	GreetWithDeadline(ctx context.Context, in *GreetWithDeadlineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetManyTimesClient interface {
	Recv() (*GreetmanytimesResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyTimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyTimesClient) Recv() (*GreetmanytimesResponse, error) {
	m := new(GreetmanytimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[1], "/greet.GreetService/LongGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceLongGreetClient{stream}
	return x, nil
}

type GreetService_LongGreetClient interface {
	Send(*LongGreatRequest) error
	CloseAndRecv() (*LongGreatResponse, error)
	grpc.ClientStream
}

type greetServiceLongGreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceLongGreetClient) Send(m *LongGreatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceLongGreetClient) CloseAndRecv() (*LongGreatResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongGreatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreatEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreatEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[2], "/greet.GreetService/GreatEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreatEveryoneClient{stream}
	return x, nil
}

type GreetService_GreatEveryoneClient interface {
	Send(*GreatEveryoneRequest) error
	Recv() (*GreatEveryoneResponse, error)
	grpc.ClientStream
}

type greetServiceGreatEveryoneClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreatEveryoneClient) Send(m *GreatEveryoneRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreatEveryoneClient) Recv() (*GreatEveryoneResponse, error) {
	m := new(GreatEveryoneResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetWithDeadline(ctx context.Context, in *GreetWithDeadlineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error) {
	out := new(GreetWithDeadlineResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/GreetWithDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// Unary
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	// Server Streaming
	GreetManyTimes(*GreetManyTimesRequest, GreetService_GreetManyTimesServer) error
	// Client Streaming
	LongGreet(GreetService_LongGreetServer) error
	// BiDi Streaming
	GreatEveryone(GreetService_GreatEveryoneServer) error
	// Unary
	GreetWithDeadline(context.Context, *GreetWithDeadlineRequest) (*GreetWithDeadlineResponse, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedGreetServiceServer) GreetManyTimes(req *GreetManyTimesRequest, srv GreetService_GreetManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyTimes not implemented")
}
func (*UnimplementedGreetServiceServer) LongGreet(srv GreetService_LongGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method LongGreet not implemented")
}
func (*UnimplementedGreetServiceServer) GreatEveryone(srv GreetService_GreatEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreatEveryone not implemented")
}
func (*UnimplementedGreetServiceServer) GreetWithDeadline(ctx context.Context, req *GreetWithDeadlineRequest) (*GreetWithDeadlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetWithDeadline not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetManyTimes(m, &greetServiceGreetManyTimesServer{stream})
}

type GreetService_GreetManyTimesServer interface {
	Send(*GreetmanytimesResponse) error
	grpc.ServerStream
}

type greetServiceGreetManyTimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyTimesServer) Send(m *GreetmanytimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_LongGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).LongGreet(&greetServiceLongGreetServer{stream})
}

type GreetService_LongGreetServer interface {
	SendAndClose(*LongGreatResponse) error
	Recv() (*LongGreatRequest, error)
	grpc.ServerStream
}

type greetServiceLongGreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceLongGreetServer) SendAndClose(m *LongGreatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceLongGreetServer) Recv() (*LongGreatRequest, error) {
	m := new(LongGreatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreatEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreatEveryone(&greetServiceGreatEveryoneServer{stream})
}

type GreetService_GreatEveryoneServer interface {
	Send(*GreatEveryoneResponse) error
	Recv() (*GreatEveryoneRequest, error)
	grpc.ServerStream
}

type greetServiceGreatEveryoneServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreatEveryoneServer) Send(m *GreatEveryoneResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreatEveryoneServer) Recv() (*GreatEveryoneRequest, error) {
	m := new(GreatEveryoneRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreetWithDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetWithDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/GreetWithDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, req.(*GreetWithDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
		{
			MethodName: "GreetWithDeadline",
			Handler:    _GreetService_GreetWithDeadline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyTimes",
			Handler:       _GreetService_GreetManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LongGreet",
			Handler:       _GreetService_LongGreet_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreatEveryone",
			Handler:       _GreetService_GreatEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greetpb/greet.proto",
}
