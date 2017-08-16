// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	CreateContactReq
	CreateContactRes
	ListContactsReq
	ListContactsRes
	Contact
	PhoneNumber
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PhoneNumber_PhoneType int32

const (
	PhoneNumber_MOBILE PhoneNumber_PhoneType = 0
	PhoneNumber_HOME   PhoneNumber_PhoneType = 1
	PhoneNumber_WORK   PhoneNumber_PhoneType = 2
)

var PhoneNumber_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var PhoneNumber_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x PhoneNumber_PhoneType) String() string {
	return proto.EnumName(PhoneNumber_PhoneType_name, int32(x))
}
func (PhoneNumber_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type CreateContactReq struct {
	Name         string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email        string         `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	PhoneNumbers []*PhoneNumber `protobuf:"bytes,3,rep,name=phone_numbers,json=phoneNumbers" json:"phone_numbers,omitempty"`
}

func (m *CreateContactReq) Reset()                    { *m = CreateContactReq{} }
func (m *CreateContactReq) String() string            { return proto.CompactTextString(m) }
func (*CreateContactReq) ProtoMessage()               {}
func (*CreateContactReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateContactReq) GetPhoneNumbers() []*PhoneNumber {
	if m != nil {
		return m.PhoneNumbers
	}
	return nil
}

type CreateContactRes struct {
	Contact *Contact `protobuf:"bytes,1,opt,name=contact" json:"contact,omitempty"`
}

func (m *CreateContactRes) Reset()                    { *m = CreateContactRes{} }
func (m *CreateContactRes) String() string            { return proto.CompactTextString(m) }
func (*CreateContactRes) ProtoMessage()               {}
func (*CreateContactRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateContactRes) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

type ListContactsReq struct {
}

func (m *ListContactsReq) Reset()                    { *m = ListContactsReq{} }
func (m *ListContactsReq) String() string            { return proto.CompactTextString(m) }
func (*ListContactsReq) ProtoMessage()               {}
func (*ListContactsReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ListContactsRes struct {
	Contacts []*Contact `protobuf:"bytes,1,rep,name=contacts" json:"contacts,omitempty"`
}

func (m *ListContactsRes) Reset()                    { *m = ListContactsRes{} }
func (m *ListContactsRes) String() string            { return proto.CompactTextString(m) }
func (*ListContactsRes) ProtoMessage()               {}
func (*ListContactsRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListContactsRes) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type Contact struct {
	Name         string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email        string         `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	PhoneNumbers []*PhoneNumber `protobuf:"bytes,3,rep,name=phone_numbers,json=phoneNumbers" json:"phone_numbers,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Contact) GetPhoneNumbers() []*PhoneNumber {
	if m != nil {
		return m.PhoneNumbers
	}
	return nil
}

type PhoneNumber struct {
	Number string                `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   PhoneNumber_PhoneType `protobuf:"varint,2,opt,name=type,enum=api.PhoneNumber_PhoneType" json:"type,omitempty"`
}

func (m *PhoneNumber) Reset()                    { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()               {}
func (*PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*CreateContactReq)(nil), "api.CreateContactReq")
	proto.RegisterType((*CreateContactRes)(nil), "api.CreateContactRes")
	proto.RegisterType((*ListContactsReq)(nil), "api.ListContactsReq")
	proto.RegisterType((*ListContactsRes)(nil), "api.ListContactsRes")
	proto.RegisterType((*Contact)(nil), "api.Contact")
	proto.RegisterType((*PhoneNumber)(nil), "api.PhoneNumber")
	proto.RegisterEnum("api.PhoneNumber_PhoneType", PhoneNumber_PhoneType_name, PhoneNumber_PhoneType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PhoneBook service

type PhoneBookClient interface {
	CreateContact(ctx context.Context, in *CreateContactReq, opts ...grpc.CallOption) (*CreateContactRes, error)
	ListContacts(ctx context.Context, in *ListContactsReq, opts ...grpc.CallOption) (*ListContactsRes, error)
}

type phoneBookClient struct {
	cc *grpc.ClientConn
}

func NewPhoneBookClient(cc *grpc.ClientConn) PhoneBookClient {
	return &phoneBookClient{cc}
}

func (c *phoneBookClient) CreateContact(ctx context.Context, in *CreateContactReq, opts ...grpc.CallOption) (*CreateContactRes, error) {
	out := new(CreateContactRes)
	err := grpc.Invoke(ctx, "/api.PhoneBook/CreateContact", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phoneBookClient) ListContacts(ctx context.Context, in *ListContactsReq, opts ...grpc.CallOption) (*ListContactsRes, error) {
	out := new(ListContactsRes)
	err := grpc.Invoke(ctx, "/api.PhoneBook/ListContacts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PhoneBook service

type PhoneBookServer interface {
	CreateContact(context.Context, *CreateContactReq) (*CreateContactRes, error)
	ListContacts(context.Context, *ListContactsReq) (*ListContactsRes, error)
}

func RegisterPhoneBookServer(s *grpc.Server, srv PhoneBookServer) {
	s.RegisterService(&_PhoneBook_serviceDesc, srv)
}

func _PhoneBook_CreateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneBookServer).CreateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PhoneBook/CreateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneBookServer).CreateContact(ctx, req.(*CreateContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhoneBook_ListContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContactsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneBookServer).ListContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PhoneBook/ListContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneBookServer).ListContacts(ctx, req.(*ListContactsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PhoneBook_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.PhoneBook",
	HandlerType: (*PhoneBookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContact",
			Handler:    _PhoneBook_CreateContact_Handler,
		},
		{
			MethodName: "ListContacts",
			Handler:    _PhoneBook_ListContacts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcf, 0x4f, 0xc2, 0x30,
	0x14, 0xb6, 0x80, 0xfc, 0x78, 0x80, 0xce, 0x17, 0x34, 0x0b, 0x27, 0xd2, 0x83, 0x21, 0x31, 0xe1,
	0x30, 0xe3, 0x05, 0xe3, 0x05, 0x42, 0xa2, 0x11, 0x9c, 0x69, 0x4c, 0x3c, 0x9a, 0x42, 0x9a, 0x38,
	0x75, 0x6b, 0x59, 0xeb, 0x81, 0xab, 0x07, 0xff, 0x6e, 0xb3, 0xae, 0x20, 0x4c, 0xae, 0xde, 0xde,
	0xf7, 0x63, 0xdf, 0xb7, 0xb6, 0x0f, 0x1a, 0x5c, 0x45, 0x03, 0x95, 0x4a, 0x23, 0xb1, 0xcc, 0x55,
	0x44, 0x35, 0x78, 0xe3, 0x54, 0x70, 0x23, 0xc6, 0x32, 0x31, 0x7c, 0x61, 0x98, 0x58, 0x22, 0x42,
	0x25, 0xe1, 0xb1, 0xf0, 0x49, 0x8f, 0xf4, 0x1b, 0xcc, 0xce, 0xd8, 0x81, 0x43, 0x11, 0xf3, 0xe8,
	0xc3, 0x2f, 0x59, 0x32, 0x07, 0x78, 0x05, 0x6d, 0xf5, 0x2a, 0x13, 0xf1, 0x92, 0x7c, 0xc6, 0x73,
	0x91, 0x6a, 0xbf, 0xdc, 0x2b, 0xf7, 0x9b, 0x81, 0x37, 0xc8, 0x5a, 0x1e, 0x33, 0xe5, 0xc1, 0x0a,
	0xac, 0xa5, 0x7e, 0x81, 0xa6, 0xc3, 0x3f, 0xa5, 0x1a, 0xcf, 0xa1, 0xb6, 0xc8, 0x91, 0xed, 0x6d,
	0x06, 0x2d, 0x1b, 0xb2, 0x76, 0xac, 0x45, 0x7a, 0x02, 0xc7, 0xd3, 0x48, 0x1b, 0xc7, 0x6b, 0x26,
	0x96, 0xf4, 0xba, 0x48, 0x69, 0xec, 0x43, 0xdd, 0x7d, 0xa0, 0x7d, 0x62, 0xff, 0x69, 0x37, 0x6e,
	0xa3, 0xd2, 0x37, 0xa8, 0x39, 0xf2, 0xff, 0xcf, 0xfd, 0x45, 0xa0, 0xb9, 0xa5, 0xe2, 0x19, 0x54,
	0xf3, 0x00, 0x57, 0xe9, 0x10, 0x0e, 0xa0, 0x62, 0x56, 0x4a, 0xd8, 0xce, 0xa3, 0xa0, 0x5b, 0x4c,
	0xcd, 0xe7, 0xa7, 0x95, 0x12, 0xcc, 0xfa, 0xe8, 0x05, 0x34, 0x36, 0x14, 0x02, 0x54, 0x67, 0xe1,
	0xe8, 0x6e, 0x3a, 0xf1, 0x0e, 0xb0, 0x0e, 0x95, 0xdb, 0x70, 0x36, 0xf1, 0x48, 0x36, 0x3d, 0x87,
	0xec, 0xde, 0x2b, 0x05, 0xdf, 0xc4, 0xb9, 0x47, 0x52, 0xbe, 0xe3, 0x0d, 0xb4, 0x77, 0x9e, 0x02,
	0x4f, 0xf3, 0x7b, 0x2a, 0xec, 0x44, 0x77, 0x2f, 0xad, 0x71, 0x08, 0xad, 0xed, 0xab, 0xc7, 0x8e,
	0xb5, 0x15, 0x1e, 0xa8, 0xbb, 0x8f, 0xd5, 0xf3, 0xaa, 0x5d, 0xc3, 0xcb, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5b, 0x0e, 0xcb, 0x8a, 0x93, 0x02, 0x00, 0x00,
}
