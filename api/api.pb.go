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
	DeleteContactReq
	DeleteContactRes
	ShowContactReq
	ShowContactRes
	UpdateContactReq
	UpdateContactRes
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
func (PhoneNumber_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{11, 0} }

type CreateContactReq struct {
	Name         string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email        string         `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	PhoneNumbers []*PhoneNumber `protobuf:"bytes,3,rep,name=phone_numbers,json=phoneNumbers" json:"phone_numbers,omitempty"`
	Home         *PhoneNumber   `protobuf:"bytes,4,opt,name=home" json:"home,omitempty"`
	Mobile       *PhoneNumber   `protobuf:"bytes,5,opt,name=mobile" json:"mobile,omitempty"`
	Work         *PhoneNumber   `protobuf:"bytes,6,opt,name=work" json:"work,omitempty"`
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

func (m *CreateContactReq) GetHome() *PhoneNumber {
	if m != nil {
		return m.Home
	}
	return nil
}

func (m *CreateContactReq) GetMobile() *PhoneNumber {
	if m != nil {
		return m.Mobile
	}
	return nil
}

func (m *CreateContactReq) GetWork() *PhoneNumber {
	if m != nil {
		return m.Work
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

type DeleteContactReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteContactReq) Reset()                    { *m = DeleteContactReq{} }
func (m *DeleteContactReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteContactReq) ProtoMessage()               {}
func (*DeleteContactReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DeleteContactRes struct {
	Contact *Contact `protobuf:"bytes,1,opt,name=contact" json:"contact,omitempty"`
}

func (m *DeleteContactRes) Reset()                    { *m = DeleteContactRes{} }
func (m *DeleteContactRes) String() string            { return proto.CompactTextString(m) }
func (*DeleteContactRes) ProtoMessage()               {}
func (*DeleteContactRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteContactRes) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

type ShowContactReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ShowContactReq) Reset()                    { *m = ShowContactReq{} }
func (m *ShowContactReq) String() string            { return proto.CompactTextString(m) }
func (*ShowContactReq) ProtoMessage()               {}
func (*ShowContactReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ShowContactRes struct {
	Contact *Contact `protobuf:"bytes,1,opt,name=contact" json:"contact,omitempty"`
}

func (m *ShowContactRes) Reset()                    { *m = ShowContactRes{} }
func (m *ShowContactRes) String() string            { return proto.CompactTextString(m) }
func (*ShowContactRes) ProtoMessage()               {}
func (*ShowContactRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ShowContactRes) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

type UpdateContactReq struct {
	Name         string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email        string         `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	PhoneNumbers []*PhoneNumber `protobuf:"bytes,3,rep,name=phone_numbers,json=phoneNumbers" json:"phone_numbers,omitempty"`
	Home         *PhoneNumber   `protobuf:"bytes,4,opt,name=home" json:"home,omitempty"`
	Mobile       *PhoneNumber   `protobuf:"bytes,5,opt,name=mobile" json:"mobile,omitempty"`
	Work         *PhoneNumber   `protobuf:"bytes,6,opt,name=work" json:"work,omitempty"`
}

func (m *UpdateContactReq) Reset()                    { *m = UpdateContactReq{} }
func (m *UpdateContactReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateContactReq) ProtoMessage()               {}
func (*UpdateContactReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateContactReq) GetPhoneNumbers() []*PhoneNumber {
	if m != nil {
		return m.PhoneNumbers
	}
	return nil
}

func (m *UpdateContactReq) GetHome() *PhoneNumber {
	if m != nil {
		return m.Home
	}
	return nil
}

func (m *UpdateContactReq) GetMobile() *PhoneNumber {
	if m != nil {
		return m.Mobile
	}
	return nil
}

func (m *UpdateContactReq) GetWork() *PhoneNumber {
	if m != nil {
		return m.Work
	}
	return nil
}

type UpdateContactRes struct {
	Contact *Contact `protobuf:"bytes,1,opt,name=contact" json:"contact,omitempty"`
}

func (m *UpdateContactRes) Reset()                    { *m = UpdateContactRes{} }
func (m *UpdateContactRes) String() string            { return proto.CompactTextString(m) }
func (*UpdateContactRes) ProtoMessage()               {}
func (*UpdateContactRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UpdateContactRes) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

type Contact struct {
	Name         string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email        string         `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	PhoneNumbers []*PhoneNumber `protobuf:"bytes,3,rep,name=phone_numbers,json=phoneNumbers" json:"phone_numbers,omitempty"`
	Home         *PhoneNumber   `protobuf:"bytes,4,opt,name=home" json:"home,omitempty"`
	Mobile       *PhoneNumber   `protobuf:"bytes,5,opt,name=mobile" json:"mobile,omitempty"`
	Work         *PhoneNumber   `protobuf:"bytes,6,opt,name=work" json:"work,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Contact) GetPhoneNumbers() []*PhoneNumber {
	if m != nil {
		return m.PhoneNumbers
	}
	return nil
}

func (m *Contact) GetHome() *PhoneNumber {
	if m != nil {
		return m.Home
	}
	return nil
}

func (m *Contact) GetMobile() *PhoneNumber {
	if m != nil {
		return m.Mobile
	}
	return nil
}

func (m *Contact) GetWork() *PhoneNumber {
	if m != nil {
		return m.Work
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
func (*PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*CreateContactReq)(nil), "api.CreateContactReq")
	proto.RegisterType((*CreateContactRes)(nil), "api.CreateContactRes")
	proto.RegisterType((*ListContactsReq)(nil), "api.ListContactsReq")
	proto.RegisterType((*ListContactsRes)(nil), "api.ListContactsRes")
	proto.RegisterType((*DeleteContactReq)(nil), "api.DeleteContactReq")
	proto.RegisterType((*DeleteContactRes)(nil), "api.DeleteContactRes")
	proto.RegisterType((*ShowContactReq)(nil), "api.ShowContactReq")
	proto.RegisterType((*ShowContactRes)(nil), "api.ShowContactRes")
	proto.RegisterType((*UpdateContactReq)(nil), "api.UpdateContactReq")
	proto.RegisterType((*UpdateContactRes)(nil), "api.UpdateContactRes")
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
	DeleteContact(ctx context.Context, in *DeleteContactReq, opts ...grpc.CallOption) (*DeleteContactRes, error)
	ShowContact(ctx context.Context, in *ShowContactReq, opts ...grpc.CallOption) (*ShowContactRes, error)
	UpdateContact(ctx context.Context, in *UpdateContactReq, opts ...grpc.CallOption) (*UpdateContactRes, error)
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

func (c *phoneBookClient) DeleteContact(ctx context.Context, in *DeleteContactReq, opts ...grpc.CallOption) (*DeleteContactRes, error) {
	out := new(DeleteContactRes)
	err := grpc.Invoke(ctx, "/api.PhoneBook/DeleteContact", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phoneBookClient) ShowContact(ctx context.Context, in *ShowContactReq, opts ...grpc.CallOption) (*ShowContactRes, error) {
	out := new(ShowContactRes)
	err := grpc.Invoke(ctx, "/api.PhoneBook/ShowContact", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phoneBookClient) UpdateContact(ctx context.Context, in *UpdateContactReq, opts ...grpc.CallOption) (*UpdateContactRes, error) {
	out := new(UpdateContactRes)
	err := grpc.Invoke(ctx, "/api.PhoneBook/UpdateContact", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PhoneBook service

type PhoneBookServer interface {
	CreateContact(context.Context, *CreateContactReq) (*CreateContactRes, error)
	ListContacts(context.Context, *ListContactsReq) (*ListContactsRes, error)
	DeleteContact(context.Context, *DeleteContactReq) (*DeleteContactRes, error)
	ShowContact(context.Context, *ShowContactReq) (*ShowContactRes, error)
	UpdateContact(context.Context, *UpdateContactReq) (*UpdateContactRes, error)
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

func _PhoneBook_DeleteContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneBookServer).DeleteContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PhoneBook/DeleteContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneBookServer).DeleteContact(ctx, req.(*DeleteContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhoneBook_ShowContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneBookServer).ShowContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PhoneBook/ShowContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneBookServer).ShowContact(ctx, req.(*ShowContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhoneBook_UpdateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneBookServer).UpdateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PhoneBook/UpdateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneBookServer).UpdateContact(ctx, req.(*UpdateContactReq))
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
		{
			MethodName: "DeleteContact",
			Handler:    _PhoneBook_DeleteContact_Handler,
		},
		{
			MethodName: "ShowContact",
			Handler:    _PhoneBook_ShowContact_Handler,
		},
		{
			MethodName: "UpdateContact",
			Handler:    _PhoneBook_UpdateContact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x55, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xc5, 0x6d, 0x36, 0xbb, 0x9d, 0x74, 0x97, 0x30, 0x2c, 0x28, 0xca, 0xa9, 0x8a, 0x56, 0xab,
	0x48, 0x48, 0x3d, 0x04, 0x21, 0xd0, 0x22, 0x2e, 0xbb, 0xac, 0x04, 0xa2, 0xa5, 0x28, 0x80, 0x38,
	0xa2, 0xb4, 0x58, 0x6a, 0xd4, 0x26, 0x76, 0xe3, 0xa0, 0xaa, 0x57, 0xfe, 0x8b, 0x7f, 0xe1, 0xca,
	0x5f, 0x20, 0x3b, 0x6e, 0x69, 0x4c, 0xaa, 0xf6, 0xcc, 0xde, 0x3c, 0xef, 0x3d, 0xcf, 0x78, 0x26,
	0x7e, 0x0e, 0x74, 0x12, 0x9e, 0xf6, 0x79, 0xc1, 0x4a, 0x86, 0xed, 0x84, 0xa7, 0xc1, 0x6f, 0x02,
	0xee, 0x4d, 0x41, 0x93, 0x92, 0xde, 0xb0, 0xbc, 0x4c, 0x26, 0x65, 0x4c, 0x17, 0x88, 0x60, 0xe5,
	0x49, 0x46, 0x3d, 0xd2, 0x23, 0x61, 0x27, 0x56, 0x6b, 0x3c, 0x87, 0x23, 0x9a, 0x25, 0xe9, 0xdc,
	0x6b, 0x29, 0xb0, 0x0a, 0xf0, 0x19, 0x9c, 0xf2, 0x29, 0xcb, 0xe9, 0xd7, 0xfc, 0x7b, 0x36, 0xa6,
	0x85, 0xf0, 0xda, 0xbd, 0x76, 0xe8, 0x44, 0x6e, 0x5f, 0x96, 0xf9, 0x20, 0x99, 0xf7, 0x8a, 0x88,
	0xbb, 0xfc, 0x6f, 0x20, 0xf0, 0x02, 0xac, 0x29, 0xcb, 0xa8, 0x67, 0xf5, 0x48, 0xa3, 0x5a, 0xb1,
	0x18, 0x82, 0x9d, 0xb1, 0x71, 0x3a, 0xa7, 0xde, 0xd1, 0x0e, 0x9d, 0xe6, 0x65, 0xbe, 0x25, 0x2b,
	0x66, 0x9e, 0xbd, 0x2b, 0x9f, 0x64, 0x83, 0xab, 0x7f, 0x5a, 0x15, 0x78, 0x09, 0xc7, 0x93, 0x2a,
	0x52, 0xdd, 0x3a, 0x51, 0x57, 0x6d, 0x5e, 0x2b, 0xd6, 0x64, 0xf0, 0x00, 0xee, 0x0f, 0x52, 0x51,
	0x6a, 0x5c, 0xc4, 0x74, 0x11, 0xbc, 0x34, 0x21, 0x81, 0x21, 0x9c, 0xe8, 0x0d, 0xc2, 0x23, 0x6a,
	0x12, 0xf5, 0x74, 0x1b, 0x36, 0xb8, 0x04, 0xf7, 0x35, 0x9d, 0xd3, 0x7d, 0x63, 0x97, 0x67, 0x36,
	0x74, 0x87, 0x9f, 0xf9, 0x02, 0xce, 0x3e, 0x4e, 0xd9, 0x72, 0x4f, 0x85, 0x17, 0x86, 0xea, 0xf0,
	0xfc, 0xf2, 0xee, 0x7c, 0xe6, 0xdf, 0xee, 0xca, 0xdd, 0x31, 0x5a, 0x3d, 0x7c, 0x4e, 0xbf, 0x08,
	0x1c, 0x6b, 0xf0, 0x7f, 0x1d, 0xcf, 0x0f, 0x02, 0xce, 0x16, 0x8a, 0x8f, 0xc1, 0xae, 0x8e, 0xad,
	0x1b, 0xd5, 0x11, 0xf6, 0xc1, 0x2a, 0x57, 0x9c, 0xaa, 0x4e, 0xcf, 0x22, 0xdf, 0xcc, 0x56, 0xad,
	0x3f, 0xad, 0x38, 0x8d, 0x95, 0x2e, 0x78, 0x02, 0x9d, 0x0d, 0x84, 0x00, 0xf6, 0x70, 0x74, 0xfd,
	0x76, 0x70, 0xeb, 0xde, 0xc3, 0x13, 0xb0, 0xde, 0x8c, 0x86, 0xb7, 0x2e, 0x91, 0xab, 0x2f, 0xa3,
	0xf8, 0x9d, 0xdb, 0x8a, 0x7e, 0xb6, 0xb4, 0xfa, 0x9a, 0xb1, 0x19, 0xbe, 0x82, 0xd3, 0x9a, 0xdb,
	0xf1, 0x51, 0xf5, 0x75, 0x8c, 0xc7, 0xce, 0x6f, 0x84, 0x05, 0x5e, 0x41, 0x77, 0xdb, 0xdd, 0x78,
	0xae, 0x64, 0xc6, 0x1b, 0xe0, 0x37, 0xa1, 0x42, 0x96, 0xae, 0x99, 0x56, 0x97, 0x36, 0x0d, 0xef,
	0x37, 0xc2, 0x02, 0x9f, 0x83, 0xb3, 0xe5, 0x48, 0x7c, 0xa8, 0x54, 0x75, 0x27, 0xfb, 0x0d, 0xa0,
	0xaa, 0x5b, 0xbb, 0xa4, 0xba, 0xae, 0xe9, 0x51, 0xbf, 0x11, 0x16, 0x63, 0x5b, 0xfd, 0x17, 0x9e,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x95, 0x56, 0x36, 0x2c, 0x24, 0x06, 0x00, 0x00,
}
