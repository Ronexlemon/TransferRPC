// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: transfer.proto

package Transfer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TransferService_Transfer_FullMethodName = "/TransferService/Transfer"
)

// TransferServiceClient is the client API for TransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferServiceClient interface {
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
}

type transferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferServiceClient(cc grpc.ClientConnInterface) TransferServiceClient {
	return &transferServiceClient{cc}
}

func (c *transferServiceClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, TransferService_Transfer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferServiceServer is the server API for TransferService service.
// All implementations must embed UnimplementedTransferServiceServer
// for forward compatibility.
type TransferServiceServer interface {
	Transfer(context.Context, *TransferRequest) (*TransferResponse, error)
	mustEmbedUnimplementedTransferServiceServer()
}

// UnimplementedTransferServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransferServiceServer struct{}

func (UnimplementedTransferServiceServer) Transfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedTransferServiceServer) mustEmbedUnimplementedTransferServiceServer() {}
func (UnimplementedTransferServiceServer) testEmbeddedByValue()                         {}

// UnsafeTransferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferServiceServer will
// result in compilation errors.
type UnsafeTransferServiceServer interface {
	mustEmbedUnimplementedTransferServiceServer()
}

func RegisterTransferServiceServer(s grpc.ServiceRegistrar, srv TransferServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransferServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransferService_ServiceDesc, srv)
}

func _TransferService_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_Transfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransferService_ServiceDesc is the grpc.ServiceDesc for TransferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TransferService",
	HandlerType: (*TransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _TransferService_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfer.proto",
}
