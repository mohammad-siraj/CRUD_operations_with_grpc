// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: crud_proto/crud_proto.proto

package crud_proto

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

// CarInfoClient is the client API for CarInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarInfoClient interface {
	GetCarInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Car, error)
	UpdateInfo(ctx context.Context, in *CarComplete, opts ...grpc.CallOption) (*Car, error)
	DeleteInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Car, error)
	CreateInfo(ctx context.Context, in *Car, opts ...grpc.CallOption) (*Car, error)
}

type carInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewCarInfoClient(cc grpc.ClientConnInterface) CarInfoClient {
	return &carInfoClient{cc}
}

func (c *carInfoClient) GetCarInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Car, error) {
	out := new(Car)
	err := c.cc.Invoke(ctx, "/crud_proto.CarInfo/GetCarInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carInfoClient) UpdateInfo(ctx context.Context, in *CarComplete, opts ...grpc.CallOption) (*Car, error) {
	out := new(Car)
	err := c.cc.Invoke(ctx, "/crud_proto.CarInfo/UpdateInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carInfoClient) DeleteInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Car, error) {
	out := new(Car)
	err := c.cc.Invoke(ctx, "/crud_proto.CarInfo/DeleteInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carInfoClient) CreateInfo(ctx context.Context, in *Car, opts ...grpc.CallOption) (*Car, error) {
	out := new(Car)
	err := c.cc.Invoke(ctx, "/crud_proto.CarInfo/CreateInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarInfoServer is the server API for CarInfo service.
// All implementations must embed UnimplementedCarInfoServer
// for forward compatibility
type CarInfoServer interface {
	GetCarInfo(context.Context, *Id) (*Car, error)
	UpdateInfo(context.Context, *CarComplete) (*Car, error)
	DeleteInfo(context.Context, *Id) (*Car, error)
	CreateInfo(context.Context, *Car) (*Car, error)
	mustEmbedUnimplementedCarInfoServer()
}

// UnimplementedCarInfoServer must be embedded to have forward compatible implementations.
type UnimplementedCarInfoServer struct {
}

func (UnimplementedCarInfoServer) GetCarInfo(context.Context, *Id) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarInfo not implemented")
}
func (UnimplementedCarInfoServer) UpdateInfo(context.Context, *CarComplete) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInfo not implemented")
}
func (UnimplementedCarInfoServer) DeleteInfo(context.Context, *Id) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInfo not implemented")
}
func (UnimplementedCarInfoServer) CreateInfo(context.Context, *Car) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInfo not implemented")
}
func (UnimplementedCarInfoServer) mustEmbedUnimplementedCarInfoServer() {}

// UnsafeCarInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarInfoServer will
// result in compilation errors.
type UnsafeCarInfoServer interface {
	mustEmbedUnimplementedCarInfoServer()
}

func RegisterCarInfoServer(s grpc.ServiceRegistrar, srv CarInfoServer) {
	s.RegisterService(&CarInfo_ServiceDesc, srv)
}

func _CarInfo_GetCarInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarInfoServer).GetCarInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_proto.CarInfo/GetCarInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarInfoServer).GetCarInfo(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarInfo_UpdateInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarComplete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarInfoServer).UpdateInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_proto.CarInfo/UpdateInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarInfoServer).UpdateInfo(ctx, req.(*CarComplete))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarInfo_DeleteInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarInfoServer).DeleteInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_proto.CarInfo/DeleteInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarInfoServer).DeleteInfo(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarInfo_CreateInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Car)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarInfoServer).CreateInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_proto.CarInfo/CreateInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarInfoServer).CreateInfo(ctx, req.(*Car))
	}
	return interceptor(ctx, in, info, handler)
}

// CarInfo_ServiceDesc is the grpc.ServiceDesc for CarInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crud_proto.CarInfo",
	HandlerType: (*CarInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCarInfo",
			Handler:    _CarInfo_GetCarInfo_Handler,
		},
		{
			MethodName: "UpdateInfo",
			Handler:    _CarInfo_UpdateInfo_Handler,
		},
		{
			MethodName: "DeleteInfo",
			Handler:    _CarInfo_DeleteInfo_Handler,
		},
		{
			MethodName: "CreateInfo",
			Handler:    _CarInfo_CreateInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crud_proto/crud_proto.proto",
}
