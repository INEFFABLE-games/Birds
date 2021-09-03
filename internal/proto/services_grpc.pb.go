// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// BirdServiceClient is the client API for BirdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BirdServiceClient interface {
	CreateBird(ctx context.Context, in *CreateBirdRequest, opts ...grpc.CallOption) (*CreateBirdReply, error)
	GetBird(ctx context.Context, in *GetBirdRequest, opts ...grpc.CallOption) (*GetBirdReply, error)
	ChangeBird(ctx context.Context, in *ChangeBirdRequest, opts ...grpc.CallOption) (*ChangeBirdReply, error)
	DeleteBird(ctx context.Context, in *DeleteBirdRequest, opts ...grpc.CallOption) (*DeleteBirdReply, error)
}

type birdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBirdServiceClient(cc grpc.ClientConnInterface) BirdServiceClient {
	return &birdServiceClient{cc}
}

func (c *birdServiceClient) CreateBird(ctx context.Context, in *CreateBirdRequest, opts ...grpc.CallOption) (*CreateBirdReply, error) {
	out := new(CreateBirdReply)
	err := c.cc.Invoke(ctx, "/proto.BirdService/CreateBird", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdServiceClient) GetBird(ctx context.Context, in *GetBirdRequest, opts ...grpc.CallOption) (*GetBirdReply, error) {
	out := new(GetBirdReply)
	err := c.cc.Invoke(ctx, "/proto.BirdService/GetBird", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdServiceClient) ChangeBird(ctx context.Context, in *ChangeBirdRequest, opts ...grpc.CallOption) (*ChangeBirdReply, error) {
	out := new(ChangeBirdReply)
	err := c.cc.Invoke(ctx, "/proto.BirdService/ChangeBird", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birdServiceClient) DeleteBird(ctx context.Context, in *DeleteBirdRequest, opts ...grpc.CallOption) (*DeleteBirdReply, error) {
	out := new(DeleteBirdReply)
	err := c.cc.Invoke(ctx, "/proto.BirdService/DeleteBird", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BirdServiceServer is the server API for BirdService service.
// All implementations must embed UnimplementedBirdServiceServer
// for forward compatibility
type BirdServiceServer interface {
	CreateBird(context.Context, *CreateBirdRequest) (*CreateBirdReply, error)
	GetBird(context.Context, *GetBirdRequest) (*GetBirdReply, error)
	ChangeBird(context.Context, *ChangeBirdRequest) (*ChangeBirdReply, error)
	DeleteBird(context.Context, *DeleteBirdRequest) (*DeleteBirdReply, error)
	mustEmbedUnimplementedBirdServiceServer()
}

// UnimplementedBirdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBirdServiceServer struct {
}

func (UnimplementedBirdServiceServer) CreateBird(context.Context, *CreateBirdRequest) (*CreateBirdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBird not implemented")
}
func (UnimplementedBirdServiceServer) GetBird(context.Context, *GetBirdRequest) (*GetBirdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBird not implemented")
}
func (UnimplementedBirdServiceServer) ChangeBird(context.Context, *ChangeBirdRequest) (*ChangeBirdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeBird not implemented")
}
func (UnimplementedBirdServiceServer) DeleteBird(context.Context, *DeleteBirdRequest) (*DeleteBirdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBird not implemented")
}
func (UnimplementedBirdServiceServer) mustEmbedUnimplementedBirdServiceServer() {}

// UnsafeBirdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BirdServiceServer will
// result in compilation errors.
type UnsafeBirdServiceServer interface {
	mustEmbedUnimplementedBirdServiceServer()
}

func RegisterBirdServiceServer(s grpc.ServiceRegistrar, srv BirdServiceServer) {
	s.RegisterService(&BirdService_ServiceDesc, srv)
}

func _BirdService_CreateBird_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBirdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdServiceServer).CreateBird(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BirdService/CreateBird",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdServiceServer).CreateBird(ctx, req.(*CreateBirdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdService_GetBird_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBirdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdServiceServer).GetBird(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BirdService/GetBird",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdServiceServer).GetBird(ctx, req.(*GetBirdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdService_ChangeBird_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeBirdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdServiceServer).ChangeBird(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BirdService/ChangeBird",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdServiceServer).ChangeBird(ctx, req.(*ChangeBirdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirdService_DeleteBird_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBirdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirdServiceServer).DeleteBird(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BirdService/DeleteBird",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirdServiceServer).DeleteBird(ctx, req.(*DeleteBirdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BirdService_ServiceDesc is the grpc.ServiceDesc for BirdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BirdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BirdService",
	HandlerType: (*BirdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBird",
			Handler:    _BirdService_CreateBird_Handler,
		},
		{
			MethodName: "GetBird",
			Handler:    _BirdService_GetBird_Handler,
		},
		{
			MethodName: "ChangeBird",
			Handler:    _BirdService_ChangeBird_Handler,
		},
		{
			MethodName: "DeleteBird",
			Handler:    _BirdService_DeleteBird_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/services.proto",
}
