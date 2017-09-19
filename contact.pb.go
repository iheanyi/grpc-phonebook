// Code generated by protoc-gen-go.
// source: contact.proto
// DO NOT EDIT!

/*
Package contact is a generated protocol buffer package.

It is generated from these files:
	contact.proto

It has these top-level messages:
	Contact
	PhoneNumber
*/
package contact

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
func (PhoneNumber_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

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
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	Type   PhoneNumber_PhoneType `protobuf:"varint,2,opt,name=type,enum=contact.PhoneNumber_PhoneType" json:"type,omitempty"`
}

func (m *PhoneNumber) Reset()                    { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()               {}
func (*PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Contact)(nil), "contact.Contact")
	proto.RegisterType((*PhoneNumber)(nil), "contact.PhoneNumber")
	proto.RegisterEnum("contact.PhoneNumber_PhoneType", PhoneNumber_PhoneType_name, PhoneNumber_PhoneType_value)
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0x2b,
	0x49, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x3e, 0x30,
	0x72, 0xb1, 0x3b, 0x43, 0xd8, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x08, 0x17, 0x6b, 0x6a, 0x6e, 0x62, 0x66, 0x8e, 0x04,
	0x13, 0x58, 0x10, 0xc2, 0x11, 0xb2, 0xe4, 0xe2, 0x2d, 0xc8, 0xc8, 0xcf, 0x4b, 0x8d, 0xcf, 0x2b,
	0xcd, 0x4d, 0x4a, 0x2d, 0x2a, 0x96, 0x60, 0x56, 0x60, 0xd6, 0xe0, 0x36, 0x12, 0xd1, 0x83, 0xd9,
	0x12, 0x00, 0x92, 0xf5, 0x03, 0x4b, 0x06, 0xf1, 0x14, 0x20, 0x38, 0xc5, 0x42, 0x1a, 0x5c, 0x2c,
	0x19, 0xf9, 0xb9, 0xa9, 0x12, 0x2c, 0x0a, 0x8c, 0x38, 0x75, 0x80, 0x55, 0x08, 0xe9, 0x70, 0xb1,
	0xe5, 0xe6, 0x27, 0x65, 0xe6, 0xa4, 0x4a, 0xb0, 0xe2, 0x51, 0x0b, 0x55, 0x03, 0x32, 0xb7, 0x3c,
	0xbf, 0x28, 0x5b, 0x82, 0x0d, 0x9f, 0xb9, 0x20, 0x15, 0x4a, 0x6d, 0x8c, 0x5c, 0xdc, 0x48, 0xa2,
	0x42, 0x62, 0x5c, 0x6c, 0x10, 0x6f, 0x40, 0x3d, 0x0e, 0xe5, 0x09, 0x19, 0x71, 0xb1, 0x94, 0x54,
	0x16, 0xa4, 0x82, 0x7d, 0xce, 0x67, 0x24, 0x87, 0xcd, 0x44, 0x08, 0x3b, 0xa4, 0xb2, 0x20, 0x35,
	0x08, 0xac, 0x56, 0x49, 0x9b, 0x8b, 0x13, 0x2e, 0x24, 0xc4, 0xc5, 0xc5, 0xe6, 0xeb, 0xef, 0xe4,
	0xe9, 0xe3, 0x2a, 0xc0, 0x20, 0xc4, 0xc1, 0xc5, 0xe2, 0xe1, 0xef, 0xeb, 0x2a, 0xc0, 0x08, 0x62,
	0x85, 0xfb, 0x07, 0x79, 0x0b, 0x30, 0x25, 0xb1, 0x81, 0xe3, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0xec, 0x7d, 0xa3, 0x9c, 0x01, 0x00, 0x00,
}
