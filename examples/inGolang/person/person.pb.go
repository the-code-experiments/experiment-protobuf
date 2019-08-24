// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

package profile

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Person_ROLE int32

const (
	Person_ENGINEER Person_ROLE = 0
	Person_MANAGER  Person_ROLE = 1
)

var Person_ROLE_name = map[int32]string{
	0: "ENGINEER",
	1: "MANAGER",
}

var Person_ROLE_value = map[string]int32{
	"ENGINEER": 0,
	"MANAGER":  1,
}

func (x Person_ROLE) String() string {
	return proto.EnumName(Person_ROLE_name, int32(x))
}

func (Person_ROLE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0, 0}
}

type Person struct {
	Username  string      `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FirstName string      `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string      `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Role      Person_ROLE `protobuf:"varint,4,opt,name=role,proto3,enum=profile.Person_ROLE" json:"role,omitempty"`
	Email     string      `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Birthdate *Date       `protobuf:"bytes,6,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	// multiple addresses
	Addresses            []*Person_Address `protobuf:"bytes,7,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Followers            int32             `protobuf:"varint,8,opt,name=followers,proto3" json:"followers,omitempty"`
	Following            int32             `protobuf:"varint,9,opt,name=following,proto3" json:"following,omitempty"`
	Repositories         int32             `protobuf:"varint,10,opt,name=repositories,proto3" json:"repositories,omitempty"`
	Projects             int32             `protobuf:"varint,11,opt,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Person) GetRole() Person_ROLE {
	if m != nil {
		return m.Role
	}
	return Person_ENGINEER
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetBirthdate() *Date {
	if m != nil {
		return m.Birthdate
	}
	return nil
}

func (m *Person) GetAddresses() []*Person_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Person) GetFollowers() int32 {
	if m != nil {
		return m.Followers
	}
	return 0
}

func (m *Person) GetFollowing() int32 {
	if m != nil {
		return m.Following
	}
	return 0
}

func (m *Person) GetRepositories() int32 {
	if m != nil {
		return m.Repositories
	}
	return 0
}

func (m *Person) GetProjects() int32 {
	if m != nil {
		return m.Projects
	}
	return 0
}

// Nested messages
type Person_Address struct {
	AddressLine_1        string   `protobuf:"bytes,1,opt,name=address_line_1,json=addressLine1,proto3" json:"address_line_1,omitempty"`
	AddressLine_2        string   `protobuf:"bytes,2,opt,name=address_line_2,json=addressLine2,proto3" json:"address_line_2,omitempty"`
	ZipCode              string   `protobuf:"bytes,3,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person_Address) Reset()         { *m = Person_Address{} }
func (m *Person_Address) String() string { return proto.CompactTextString(m) }
func (*Person_Address) ProtoMessage()    {}
func (*Person_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0, 0}
}

func (m *Person_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_Address.Unmarshal(m, b)
}
func (m *Person_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_Address.Marshal(b, m, deterministic)
}
func (m *Person_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_Address.Merge(m, src)
}
func (m *Person_Address) XXX_Size() int {
	return xxx_messageInfo_Person_Address.Size(m)
}
func (m *Person_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Person_Address proto.InternalMessageInfo

func (m *Person_Address) GetAddressLine_1() string {
	if m != nil {
		return m.AddressLine_1
	}
	return ""
}

func (m *Person_Address) GetAddressLine_2() string {
	if m != nil {
		return m.AddressLine_2
	}
	return ""
}

func (m *Person_Address) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *Person_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Person_Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type Date struct {
	// Must be from 1 to 9999, or 0 if skip the year
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Must be from 1 to 12
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// Must be from 1 to 31
	Day                  int32    `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{1}
}

func (m *Date) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Date.Unmarshal(m, b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Date.Marshal(b, m, deterministic)
}
func (m *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(m, src)
}
func (m *Date) XXX_Size() int {
	return xxx_messageInfo_Date.Size(m)
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func init() {
	proto.RegisterEnum("profile.Person_ROLE", Person_ROLE_name, Person_ROLE_value)
	proto.RegisterType((*Person)(nil), "profile.Person")
	proto.RegisterType((*Person_Address)(nil), "profile.Person.Address")
	proto.RegisterType((*Date)(nil), "profile.Date")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xf1, 0x8a, 0xd3, 0x40,
	0x10, 0xc6, 0x8d, 0x4d, 0x9a, 0x66, 0x5a, 0x8f, 0x32, 0x1c, 0xb8, 0x9e, 0x0a, 0xb1, 0xf8, 0x47,
	0x40, 0x28, 0x5c, 0xc4, 0x07, 0xa8, 0x1a, 0x0e, 0xe1, 0xac, 0xb2, 0x2f, 0x50, 0x72, 0xcd, 0xf4,
	0x6e, 0x25, 0xcd, 0x86, 0xdd, 0x3d, 0xa4, 0xf7, 0x30, 0x3e, 0x8d, 0x0f, 0x26, 0x3b, 0x49, 0x9b,
	0x53, 0xff, 0x9b, 0x99, 0xdf, 0x97, 0x99, 0x8f, 0xfd, 0x02, 0xb3, 0x96, 0x8c, 0xd5, 0xcd, 0xb2,
	0x35, 0xda, 0x69, 0x8c, 0x5b, 0xa3, 0x77, 0xaa, 0xa6, 0xc5, 0xef, 0x10, 0xc6, 0xdf, 0x99, 0xe0,
	0x05, 0x4c, 0xee, 0x2d, 0x99, 0xa6, 0xdc, 0x93, 0x08, 0xd2, 0x20, 0x4b, 0xe4, 0xa9, 0xc7, 0xd7,
	0x00, 0x3b, 0x65, 0xac, 0xdb, 0x30, 0x7d, 0xca, 0x34, 0xe1, 0xc9, 0xda, 0xe3, 0x97, 0x90, 0xd4,
	0xe5, 0x91, 0x8e, 0xba, 0x6f, 0xfd, 0x80, 0x61, 0x06, 0xa1, 0xd1, 0x35, 0x89, 0x30, 0x0d, 0xb2,
	0xb3, 0xfc, 0x7c, 0xd9, 0x9f, 0x5e, 0x76, 0x67, 0x97, 0xf2, 0xdb, 0x75, 0x21, 0x59, 0x81, 0xe7,
	0x10, 0xd1, 0xbe, 0x54, 0xb5, 0x88, 0x78, 0x45, 0xd7, 0xe0, 0x3b, 0x48, 0x6e, 0x94, 0x71, 0x77,
	0x55, 0xe9, 0x48, 0x8c, 0xd3, 0x20, 0x9b, 0xe6, 0xcf, 0x4e, 0x4b, 0x3e, 0x97, 0x8e, 0xe4, 0xc0,
	0xf1, 0x03, 0x24, 0x65, 0x55, 0x19, 0xb2, 0x96, 0xac, 0x88, 0xd3, 0x51, 0x36, 0xcd, 0x9f, 0xff,
	0x7b, 0x71, 0xd5, 0x09, 0xe4, 0xa0, 0xc4, 0x57, 0x90, 0xec, 0x74, 0x5d, 0xeb, 0x9f, 0x64, 0xac,
	0x98, 0xa4, 0x41, 0x16, 0xc9, 0x61, 0x30, 0x50, 0xd5, 0xdc, 0x8a, 0xe4, 0x31, 0x55, 0xcd, 0x2d,
	0x2e, 0x60, 0x66, 0xa8, 0xd5, 0x56, 0x39, 0x6d, 0x14, 0x59, 0x01, 0x2c, 0xf8, 0x6b, 0xe6, 0xdf,
	0xb6, 0x35, 0xfa, 0x07, 0x6d, 0x9d, 0x15, 0x53, 0xe6, 0xa7, 0xfe, 0xe2, 0x57, 0x00, 0x71, 0x6f,
	0x09, 0xdf, 0xc2, 0x59, 0x6f, 0x6a, 0x53, 0xab, 0x86, 0x36, 0x97, 0x7d, 0x12, 0xb3, 0x7e, 0x7a,
	0xad, 0x1a, 0xba, 0xfc, 0x4f, 0x95, 0xf7, 0x89, 0x3c, 0x56, 0xe5, 0xf8, 0x02, 0x26, 0x0f, 0xaa,
	0xdd, 0x6c, 0x75, 0x75, 0xcc, 0x24, 0x7e, 0x50, 0xed, 0x27, 0x5d, 0x11, 0x22, 0x84, 0x5b, 0xe5,
	0x0e, 0x1c, 0x49, 0x22, 0xb9, 0x46, 0x01, 0xf1, 0x56, 0xdf, 0x37, 0xce, 0x1c, 0xfa, 0xe7, 0x3f,
	0xb6, 0x8b, 0x37, 0x10, 0xfa, 0x90, 0x70, 0x06, 0x93, 0x62, 0x7d, 0xf5, 0x65, 0x5d, 0x14, 0x72,
	0xfe, 0x04, 0xa7, 0x10, 0x7f, 0x5d, 0xad, 0x57, 0x57, 0x85, 0x9c, 0x07, 0x8b, 0x8f, 0x10, 0xfa,
	0x24, 0xfc, 0xe2, 0x03, 0x95, 0x86, 0x5d, 0x47, 0x92, 0x6b, 0x9f, 0xea, 0x5e, 0x37, 0xee, 0x8e,
	0x4d, 0x46, 0xb2, 0x6b, 0x70, 0x0e, 0xa3, 0xaa, 0x3c, 0xb0, 0xb1, 0x48, 0xfa, 0xf2, 0x66, 0xcc,
	0xbf, 0xe6, 0xfb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xd4, 0x3c, 0xf4, 0xaa, 0x02, 0x00,
	0x00,
}
