// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/infra_proxy/request/databags.proto

package request

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

type DataBags struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Data bag name.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DataBags) Reset()         { *m = DataBags{} }
func (m *DataBags) String() string { return proto.CompactTextString(m) }
func (*DataBags) ProtoMessage()    {}
func (*DataBags) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a68b987aeeb8e42, []int{0}
}

func (m *DataBags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBags.Unmarshal(m, b)
}
func (m *DataBags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBags.Marshal(b, m, deterministic)
}
func (m *DataBags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBags.Merge(m, src)
}
func (m *DataBags) XXX_Size() int {
	return xxx_messageInfo_DataBags.Size(m)
}
func (m *DataBags) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBags.DiscardUnknown(m)
}

var xxx_messageInfo_DataBags proto.InternalMessageInfo

func (m *DataBags) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *DataBags) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *DataBags) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DataBag struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Data bag name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Data bag item name.
	Item                 string   `protobuf:"bytes,4,opt,name=item,proto3" json:"item,omitempty" toml:"item,omitempty" mapstructure:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DataBag) Reset()         { *m = DataBag{} }
func (m *DataBag) String() string { return proto.CompactTextString(m) }
func (*DataBag) ProtoMessage()    {}
func (*DataBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a68b987aeeb8e42, []int{1}
}

func (m *DataBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBag.Unmarshal(m, b)
}
func (m *DataBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBag.Marshal(b, m, deterministic)
}
func (m *DataBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBag.Merge(m, src)
}
func (m *DataBag) XXX_Size() int {
	return xxx_messageInfo_DataBag.Size(m)
}
func (m *DataBag) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBag.DiscardUnknown(m)
}

var xxx_messageInfo_DataBag proto.InternalMessageInfo

func (m *DataBag) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *DataBag) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *DataBag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataBag) GetItem() string {
	if m != nil {
		return m.Item
	}
	return ""
}

type CreateDataBag struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Data bag name.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateDataBag) Reset()         { *m = CreateDataBag{} }
func (m *CreateDataBag) String() string { return proto.CompactTextString(m) }
func (*CreateDataBag) ProtoMessage()    {}
func (*CreateDataBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a68b987aeeb8e42, []int{2}
}

func (m *CreateDataBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDataBag.Unmarshal(m, b)
}
func (m *CreateDataBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDataBag.Marshal(b, m, deterministic)
}
func (m *CreateDataBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDataBag.Merge(m, src)
}
func (m *CreateDataBag) XXX_Size() int {
	return xxx_messageInfo_CreateDataBag.Size(m)
}
func (m *CreateDataBag) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDataBag.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDataBag proto.InternalMessageInfo

func (m *CreateDataBag) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *CreateDataBag) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *CreateDataBag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateDataBagItem struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Data bag name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Data bag item stringify data.
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty" toml:"data,omitempty" mapstructure:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateDataBagItem) Reset()         { *m = CreateDataBagItem{} }
func (m *CreateDataBagItem) String() string { return proto.CompactTextString(m) }
func (*CreateDataBagItem) ProtoMessage()    {}
func (*CreateDataBagItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a68b987aeeb8e42, []int{3}
}

func (m *CreateDataBagItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDataBagItem.Unmarshal(m, b)
}
func (m *CreateDataBagItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDataBagItem.Marshal(b, m, deterministic)
}
func (m *CreateDataBagItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDataBagItem.Merge(m, src)
}
func (m *CreateDataBagItem) XXX_Size() int {
	return xxx_messageInfo_CreateDataBagItem.Size(m)
}
func (m *CreateDataBagItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDataBagItem.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDataBagItem proto.InternalMessageInfo

func (m *CreateDataBagItem) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *CreateDataBagItem) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *CreateDataBagItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateDataBagItem) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*DataBags)(nil), "chef.automate.domain.infra_proxy.request.DataBags")
	proto.RegisterType((*DataBag)(nil), "chef.automate.domain.infra_proxy.request.DataBag")
	proto.RegisterType((*CreateDataBag)(nil), "chef.automate.domain.infra_proxy.request.CreateDataBag")
	proto.RegisterType((*CreateDataBagItem)(nil), "chef.automate.domain.infra_proxy.request.CreateDataBagItem")
}

func init() {
	proto.RegisterFile("api/interservice/infra_proxy/request/databags.proto", fileDescriptor_8a68b987aeeb8e42)
}

var fileDescriptor_8a68b987aeeb8e42 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x31, 0x4b, 0x03, 0x41,
	0x10, 0x85, 0x89, 0xc6, 0x98, 0x2c, 0x58, 0xb8, 0x20, 0x1c, 0xd8, 0xc8, 0x55, 0xa9, 0x76, 0x8b,
	0xd4, 0x22, 0x44, 0x9b, 0x6b, 0xd3, 0x08, 0x36, 0x61, 0xee, 0x76, 0xb2, 0xd9, 0x62, 0x6f, 0xce,
	0xb9, 0x39, 0xd1, 0x7f, 0x2f, 0xbb, 0xc9, 0x81, 0x76, 0x29, 0xae, 0x7b, 0x7c, 0xc3, 0xbc, 0x61,
	0xf8, 0xd4, 0x06, 0xba, 0x60, 0x43, 0x2b, 0xc8, 0x3d, 0xf2, 0x57, 0x68, 0xd0, 0x86, 0xf6, 0xc0,
	0xb0, 0xef, 0x98, 0xbe, 0x7f, 0x2c, 0xe3, 0xe7, 0x80, 0xbd, 0x58, 0x07, 0x02, 0x35, 0xf8, 0xde,
	0x74, 0x4c, 0x42, 0x7a, 0xdd, 0x1c, 0xf1, 0x60, 0x60, 0x10, 0x8a, 0x20, 0x68, 0x1c, 0x45, 0x08,
	0xad, 0xf9, 0xb3, 0x68, 0xce, 0x8b, 0xe5, 0x4e, 0x2d, 0xdf, 0x40, 0x60, 0x0b, 0xbe, 0xd7, 0x0f,
	0x6a, 0x41, 0xec, 0xf7, 0xc1, 0x15, 0xb3, 0xa7, 0xd9, 0x7a, 0xb5, 0xbb, 0x21, 0xf6, 0x95, 0xd3,
	0x8f, 0x6a, 0x95, 0x4e, 0x23, 0xa7, 0xc9, 0x55, 0x9e, 0x2c, 0x4f, 0xa0, 0x72, 0x5a, 0xab, 0x79,
	0x0b, 0x11, 0x8b, 0xeb, 0xcc, 0x73, 0x2e, 0x51, 0xdd, 0x9e, 0x3b, 0xa7, 0xaa, 0x4c, 0x2c, 0x08,
	0xc6, 0x62, 0x7e, 0x62, 0x29, 0x97, 0xef, 0xea, 0xee, 0x95, 0x11, 0x04, 0x27, 0x3e, 0x56, 0x92,
	0xba, 0xff, 0x57, 0x5c, 0x09, 0xc6, 0x29, 0x3f, 0x49, 0xb2, 0xc6, 0x4f, 0x52, 0xde, 0xbe, 0x7c,
	0x3c, 0xfb, 0x20, 0xc7, 0xa1, 0x36, 0x0d, 0x45, 0x9b, 0xdc, 0xd9, 0xd1, 0x9d, 0xbd, 0x44, 0x7f,
	0xbd, 0xc8, 0xda, 0x37, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x06, 0x08, 0x69, 0x26, 0x2d, 0x02,
	0x00, 0x00,
}
