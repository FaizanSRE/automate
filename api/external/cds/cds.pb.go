// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cds/cds.proto

package cds

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/api/external/cds/request"
	response "github.com/chef/automate/api/external/cds/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() {
	proto.RegisterFile("api/external/cds/cds.proto", fileDescriptor_5676bc0e222533b1)
}

var fileDescriptor_5676bc0e222533b1 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x4e, 0xfb, 0x30,
	0x10, 0xc6, 0xd5, 0x7f, 0xa5, 0xff, 0x10, 0x09, 0x81, 0x2c, 0xaa, 0xb6, 0x11, 0x62, 0x28, 0x2c,
	0xd0, 0xd6, 0x16, 0x65, 0xeb, 0x84, 0xe8, 0x54, 0xc4, 0xc4, 0xc8, 0xe6, 0xb8, 0xd7, 0xd4, 0x52,
	0xe2, 0x33, 0xf1, 0x45, 0xc0, 0xca, 0xd8, 0xb5, 0xcf, 0xd2, 0x81, 0xa7, 0x60, 0xe0, 0x15, 0x78,
	0x10, 0x94, 0xa4, 0xa0, 0xb4, 0x09, 0xea, 0x10, 0x45, 0xb9, 0xfc, 0xbe, 0xf3, 0x7d, 0xdf, 0xd9,
	0xf3, 0xa5, 0xd5, 0x02, 0x5e, 0x08, 0x12, 0x23, 0x23, 0xa1, 0x66, 0x2e, 0x7b, 0xb8, 0x4d, 0x90,
	0x90, 0xb5, 0xd4, 0x02, 0xe6, 0x5c, 0xa6, 0x84, 0xb1, 0x24, 0xe0, 0xd2, 0x6a, 0xae, 0x66, 0xce,
	0x3f, 0xab, 0x48, 0x12, 0x78, 0x4a, 0xc1, 0x91, 0x48, 0x10, 0xa9, 0xd0, 0xfa, 0xe7, 0x35, 0x90,
	0xb3, 0x68, 0x1c, 0x94, 0xa9, 0x93, 0x10, 0x31, 0x8c, 0x40, 0xe4, 0x5f, 0x41, 0x3a, 0x17, 0x8e,
	0x92, 0x54, 0xed, 0xfe, 0xcd, 0x5a, 0x49, 0x63, 0x90, 0x24, 0x69, 0x34, 0x9b, 0xe9, 0xfc, 0x41,
	0xfe, 0x52, 0xc3, 0x10, 0xcc, 0xd0, 0x3d, 0xcb, 0x30, 0x84, 0x44, 0xa0, 0xcd, 0x89, 0x1a, 0xfa,
	0x46, 0x61, 0x6c, 0xd1, 0x80, 0x21, 0x27, 0x7e, 0x1c, 0x0d, 0xc3, 0xc4, 0x2a, 0x51, 0x6a, 0x63,
	0x31, 0xd2, 0xea, 0x55, 0x68, 0x19, 0x57, 0x3b, 0x8c, 0x56, 0x4d, 0xaf, 0x39, 0x99, 0x39, 0xf6,
	0xde, 0xf0, 0x8e, 0xee, 0xb5, 0xa3, 0x09, 0x1a, 0x02, 0x43, 0x53, 0x82, 0xd8, 0xb1, 0x3e, 0xaf,
	0xcd, 0x8a, 0x6f, 0x92, 0xe1, 0x65, 0xd8, 0x1f, 0xfc, 0x09, 0x17, 0x09, 0x6d, 0xd1, 0xbd, 0xbb,
	0xe5, 0xba, 0x73, 0xe8, 0x1d, 0xa8, 0xa2, 0x36, 0xd6, 0x59, 0x71, 0xb9, 0xee, 0x1c, 0x33, 0xb6,
	0x55, 0x1a, 0x47, 0xda, 0xd1, 0xdb, 0xe7, 0xd7, 0xea, 0x5f, 0x97, 0xb5, 0xf3, 0xc8, 0x02, 0x20,
	0x29, 0x36, 0x8c, 0xc8, 0x19, 0xf6, 0xd1, 0xf0, 0xd8, 0xd4, 0x38, 0x92, 0x51, 0x54, 0x3a, 0x83,
	0x5d, 0xed, 0x99, 0xbe, 0x2a, 0xf1, 0x47, 0xfb, 0x3c, 0x54, 0x35, 0xbd, 0x87, 0x7a, 0x27, 0x6d,
	0xd6, 0xda, 0x76, 0xa2, 0x0b, 0x69, 0x6e, 0xe6, 0xb4, 0xd7, 0xad, 0x31, 0x53, 0x00, 0xe3, 0xc6,
	0xe5, 0x6d, 0xff, 0xf1, 0x22, 0xd4, 0xb4, 0x48, 0x03, 0xae, 0x30, 0x16, 0xd9, 0x4c, 0xbf, 0xeb,
	0x15, 0xbb, 0x57, 0x30, 0xf8, 0x9f, 0x6f, 0xf2, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xba,
	0x8f, 0x21, 0xf5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CdsClient is the client API for Cds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CdsClient interface {
	//
	//ListContentItems
	//
	//Returns a list of metadata for each CDS content. Provides a description and current
	//state of each content item.
	//
	//Authorization Action:
	//```
	//content:items:list
	//```
	ListContentItems(ctx context.Context, in *request.ContentItems, opts ...grpc.CallOption) (*response.ContentItems, error)
	//
	//InstallContentItem
	//
	//Installs a content item from its ID
	//
	//Authorization Action:
	//```
	//content:items:install
	//```
	InstallContentItem(ctx context.Context, in *request.InstallContentItem, opts ...grpc.CallOption) (*response.InstallContentItem, error)
}

type cdsClient struct {
	cc grpc.ClientConnInterface
}

func NewCdsClient(cc grpc.ClientConnInterface) CdsClient {
	return &cdsClient{cc}
}

func (c *cdsClient) ListContentItems(ctx context.Context, in *request.ContentItems, opts ...grpc.CallOption) (*response.ContentItems, error) {
	out := new(response.ContentItems)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cds.Cds/ListContentItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cdsClient) InstallContentItem(ctx context.Context, in *request.InstallContentItem, opts ...grpc.CallOption) (*response.InstallContentItem, error) {
	out := new(response.InstallContentItem)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cds.Cds/InstallContentItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CdsServer is the server API for Cds service.
type CdsServer interface {
	//
	//ListContentItems
	//
	//Returns a list of metadata for each CDS content. Provides a description and current
	//state of each content item.
	//
	//Authorization Action:
	//```
	//content:items:list
	//```
	ListContentItems(context.Context, *request.ContentItems) (*response.ContentItems, error)
	//
	//InstallContentItem
	//
	//Installs a content item from its ID
	//
	//Authorization Action:
	//```
	//content:items:install
	//```
	InstallContentItem(context.Context, *request.InstallContentItem) (*response.InstallContentItem, error)
}

// UnimplementedCdsServer can be embedded to have forward compatible implementations.
type UnimplementedCdsServer struct {
}

func (*UnimplementedCdsServer) ListContentItems(ctx context.Context, req *request.ContentItems) (*response.ContentItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContentItems not implemented")
}
func (*UnimplementedCdsServer) InstallContentItem(ctx context.Context, req *request.InstallContentItem) (*response.InstallContentItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallContentItem not implemented")
}

func RegisterCdsServer(s *grpc.Server, srv CdsServer) {
	s.RegisterService(&_Cds_serviceDesc, srv)
}

func _Cds_ListContentItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ContentItems)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CdsServer).ListContentItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cds.Cds/ListContentItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CdsServer).ListContentItems(ctx, req.(*request.ContentItems))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cds_InstallContentItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.InstallContentItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CdsServer).InstallContentItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cds.Cds/InstallContentItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CdsServer).InstallContentItem(ctx, req.(*request.InstallContentItem))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.cds.Cds",
	HandlerType: (*CdsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListContentItems",
			Handler:    _Cds_ListContentItems_Handler,
		},
		{
			MethodName: "InstallContentItem",
			Handler:    _Cds_InstallContentItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/external/cds/cds.proto",
}
