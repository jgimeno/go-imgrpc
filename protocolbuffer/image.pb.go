// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image.proto

/*
Package protocolbuffer is a generated protocol buffer package.

It is generated from these files:
	image.proto

It has these top-level messages:
	ImageId
	Image
	FileType
*/
package protocolbuffer

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

type ImageId struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *ImageId) Reset()                    { *m = ImageId{} }
func (m *ImageId) String() string            { return proto.CompactTextString(m) }
func (*ImageId) ProtoMessage()               {}
func (*ImageId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ImageId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ImageId) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Image struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Image) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type FileType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *FileType) Reset()                    { *m = FileType{} }
func (m *FileType) String() string            { return proto.CompactTextString(m) }
func (*FileType) ProtoMessage()               {}
func (*FileType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FileType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageId)(nil), "protocolbuffer.ImageId")
	proto.RegisterType((*Image)(nil), "protocolbuffer.Image")
	proto.RegisterType((*FileType)(nil), "protocolbuffer.FileType")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ImageService service

type ImageServiceClient interface {
	SaveImage(ctx context.Context, in *Image, opts ...grpc.CallOption) (*ImageId, error)
	GetImage(ctx context.Context, in *ImageId, opts ...grpc.CallOption) (*Image, error)
}

type imageServiceClient struct {
	cc *grpc.ClientConn
}

func NewImageServiceClient(cc *grpc.ClientConn) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) SaveImage(ctx context.Context, in *Image, opts ...grpc.CallOption) (*ImageId, error) {
	out := new(ImageId)
	err := grpc.Invoke(ctx, "/protocolbuffer.ImageService/SaveImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) GetImage(ctx context.Context, in *ImageId, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := grpc.Invoke(ctx, "/protocolbuffer.ImageService/GetImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImageService service

type ImageServiceServer interface {
	SaveImage(context.Context, *Image) (*ImageId, error)
	GetImage(context.Context, *ImageId) (*Image, error)
}

func RegisterImageServiceServer(s *grpc.Server, srv ImageServiceServer) {
	s.RegisterService(&_ImageService_serviceDesc, srv)
}

func _ImageService_SaveImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Image)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).SaveImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocolbuffer.ImageService/SaveImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).SaveImage(ctx, req.(*Image))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocolbuffer.ImageService/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).GetImage(ctx, req.(*ImageId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocolbuffer.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveImage",
			Handler:    _ImageService_SaveImage_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _ImageService_GetImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image.proto",
}

func init() { proto.RegisterFile("image.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcc, 0x4d, 0x4c,
	0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0x53, 0xc9, 0xf9, 0x39, 0x49, 0xa5,
	0x69, 0x69, 0xa9, 0x45, 0x4a, 0xba, 0x5c, 0xec, 0x9e, 0x20, 0x69, 0xcf, 0x14, 0x21, 0x3e, 0x2e,
	0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e,
	0x96, 0x92, 0xca, 0x82, 0x54, 0x09, 0x26, 0xb0, 0x08, 0x98, 0xad, 0x64, 0xcf, 0xc5, 0x0a, 0x56,
	0x4e, 0x8c, 0x62, 0x90, 0x58, 0x4a, 0x62, 0x49, 0xa2, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4f, 0x10,
	0x98, 0xad, 0x24, 0xc7, 0xc5, 0xe1, 0x96, 0x99, 0x93, 0x1a, 0x02, 0x95, 0x07, 0xeb, 0x61, 0x44,
	0xe8, 0x31, 0x6a, 0x67, 0xe4, 0xe2, 0x01, 0xdb, 0x10, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a,
	0x64, 0xcd, 0xc5, 0x19, 0x9c, 0x58, 0x96, 0x0a, 0xb1, 0x55, 0x54, 0x0f, 0xd5, 0xf9, 0x7a, 0x60,
	0x61, 0x29, 0x71, 0xac, 0xc2, 0x9e, 0x29, 0x42, 0x56, 0x5c, 0x1c, 0xee, 0xa9, 0x25, 0x10, 0xbd,
	0xb8, 0x14, 0x49, 0x61, 0x37, 0x34, 0x89, 0x0d, 0x2c, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x09, 0x4c, 0xa1, 0x9f, 0x3f, 0x01, 0x00, 0x00,
}
