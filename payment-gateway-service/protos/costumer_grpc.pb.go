// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: protos/costumer.proto

package master_service_tokogua

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CostumerDataServer_FindCostumerById_FullMethodName = "/CostumerDataServer/FindCostumerById"
)

// CostumerDataServerClient is the client API for CostumerDataServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CostumerDataServerClient interface {
	FindCostumerById(ctx context.Context, in *IdCostumer, opts ...grpc.CallOption) (*Costumer, error)
}

type costumerDataServerClient struct {
	cc grpc.ClientConnInterface
}

func NewCostumerDataServerClient(cc grpc.ClientConnInterface) CostumerDataServerClient {
	return &costumerDataServerClient{cc}
}

func (c *costumerDataServerClient) FindCostumerById(ctx context.Context, in *IdCostumer, opts ...grpc.CallOption) (*Costumer, error) {
	out := new(Costumer)
	err := c.cc.Invoke(ctx, CostumerDataServer_FindCostumerById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CostumerDataServerServer is the client API for CostumerDataServer service.
// All implementations must embed UnimplementedCostumerDataServerServer
// for forward compatibility
type CostumerDataServerServer interface {
	FindCostumerById(context.Context, *IdCostumer) (*Costumer, error)
	mustEmbedUnimplementedCostumerDataServerServer()
}

// UnimplementedCostumerDataServerServer must be embedded to have forward compatible implementations.
type UnimplementedCostumerDataServerServer struct {
}

func (UnimplementedCostumerDataServerServer) FindCostumerById(context.Context, *IdCostumer) (*Costumer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCostumerById not implemented")
}
func (UnimplementedCostumerDataServerServer) mustEmbedUnimplementedCostumerDataServerServer() {}

// UnsafeCostumerDataServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CostumerDataServerServer will
// result in compilation errors.
type UnsafeCostumerDataServerServer interface {
	mustEmbedUnimplementedCostumerDataServerServer()
}

func RegisterCostumerDataServerServer(s grpc.ServiceRegistrar, srv CostumerDataServerServer) {
	s.RegisterService(&CostumerDataServer_ServiceDesc, srv)
}

func _CostumerDataServer_FindCostumerById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdCostumer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostumerDataServerServer).FindCostumerById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostumerDataServer_FindCostumerById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostumerDataServerServer).FindCostumerById(ctx, req.(*IdCostumer))
	}
	return interceptor(ctx, in, info, handler)
}

// CostumerDataServer_ServiceDesc is the grpc.ServiceDesc for CostumerDataServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CostumerDataServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CostumerDataServer",
	HandlerType: (*CostumerDataServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindCostumerById",
			Handler:    _CostumerDataServer_FindCostumerById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/costumer.proto",
}
