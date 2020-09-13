// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Product.proto

package services

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

type ProductRequest struct {
	ProdId               int32    `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductRequest) Reset()         { *m = ProductRequest{} }
func (m *ProductRequest) String() string { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()    {}
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a057a70116a426a, []int{0}
}

func (m *ProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRequest.Unmarshal(m, b)
}
func (m *ProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRequest.Marshal(b, m, deterministic)
}
func (m *ProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRequest.Merge(m, src)
}
func (m *ProductRequest) XXX_Size() int {
	return xxx_messageInfo_ProductRequest.Size(m)
}
func (m *ProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRequest proto.InternalMessageInfo

func (m *ProductRequest) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

type ProductResponse struct {
	ProdStock            int32    `protobuf:"varint,1,opt,name=prod_stock,json=prodStock,proto3" json:"prod_stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductResponse) Reset()         { *m = ProductResponse{} }
func (m *ProductResponse) String() string { return proto.CompactTextString(m) }
func (*ProductResponse) ProtoMessage()    {}
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a057a70116a426a, []int{1}
}

func (m *ProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductResponse.Unmarshal(m, b)
}
func (m *ProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductResponse.Marshal(b, m, deterministic)
}
func (m *ProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductResponse.Merge(m, src)
}
func (m *ProductResponse) XXX_Size() int {
	return xxx_messageInfo_ProductResponse.Size(m)
}
func (m *ProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductResponse proto.InternalMessageInfo

func (m *ProductResponse) GetProdStock() int32 {
	if m != nil {
		return m.ProdStock
	}
	return 0
}

func init() {
	proto.RegisterType((*ProductRequest)(nil), "services.ProductRequest")
	proto.RegisterType((*ProductResponse)(nil), "services.ProductResponse")
}

func init() { proto.RegisterFile("Product.proto", fileDescriptor_2a057a70116a426a) }

var fileDescriptor_2a057a70116a426a = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0d, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0x2d, 0x56, 0xd2, 0xe4, 0xe2, 0x83, 0x4a, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x89, 0x73, 0xb1, 0x17, 0x14, 0xe5, 0xa7, 0xc4, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xb0, 0x06, 0xb1, 0x81, 0xb8, 0x9e, 0x29, 0x4a, 0x06, 0x5c, 0xfc, 0x70, 0xa5, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0xa9, 0x42, 0xb2, 0x5c, 0x5c, 0x60, 0xb5, 0xc5, 0x25, 0xf9, 0xc9, 0xd9, 0x50, 0xe5,
	0x9c, 0x20, 0x91, 0x60, 0x90, 0x80, 0x51, 0x28, 0xdc, 0xf0, 0x60, 0x88, 0x7d, 0x42, 0xce, 0x5c,
	0x3c, 0xee, 0xa9, 0x25, 0x01, 0x30, 0x15, 0x42, 0x12, 0x7a, 0x30, 0x97, 0xe8, 0xa1, 0x3a, 0x43,
	0x4a, 0x12, 0x8b, 0x0c, 0xc4, 0xd6, 0x24, 0x36, 0xb0, 0x27, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x2b, 0xa4, 0x55, 0x29, 0xd5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetProdStock(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProdStock(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/services.ProductService/GetProdStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	GetProdStock(context.Context, *ProductRequest) (*ProductResponse, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) GetProdStock(ctx context.Context, req *ProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStock not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_GetProdStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProdStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProductService/GetProdStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProdStock(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProdStock",
			Handler:    _ProductService_GetProdStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Product.proto",
}
