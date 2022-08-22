// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: audience/v1/audience.proto

package audiencev1

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

// AudienceServiceClient is the client API for AudienceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudienceServiceClient interface {
	BulkCreateAudiences(ctx context.Context, in *CreateAudienceRequest, opts ...grpc.CallOption) (*CreateAudienceResponse, error)
	CreateAudience(ctx context.Context, in *CreateAudienceRequest, opts ...grpc.CallOption) (*CreateAudienceResponse, error)
	UpdateAudience(ctx context.Context, in *UpdateAudienceRequest, opts ...grpc.CallOption) (*UpdateAudienceResponse, error)
	DeleteAudience(ctx context.Context, in *DeleteAudienceRequest, opts ...grpc.CallOption) (*DeleteAudienceResponse, error)
	ListAudiences(ctx context.Context, in *ListAudiencesRequest, opts ...grpc.CallOption) (*ListAudiencesResponse, error)
	GetAudience(ctx context.Context, in *GetAudienceRequest, opts ...grpc.CallOption) (*GetAudienceResponse, error)
}

type audienceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAudienceServiceClient(cc grpc.ClientConnInterface) AudienceServiceClient {
	return &audienceServiceClient{cc}
}

func (c *audienceServiceClient) BulkCreateAudiences(ctx context.Context, in *CreateAudienceRequest, opts ...grpc.CallOption) (*CreateAudienceResponse, error) {
	out := new(CreateAudienceResponse)
	err := c.cc.Invoke(ctx, "/audience.v1.AudienceService/BulkCreateAudiences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceServiceClient) CreateAudience(ctx context.Context, in *CreateAudienceRequest, opts ...grpc.CallOption) (*CreateAudienceResponse, error) {
	out := new(CreateAudienceResponse)
	err := c.cc.Invoke(ctx, "/audience.v1.AudienceService/CreateAudience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceServiceClient) UpdateAudience(ctx context.Context, in *UpdateAudienceRequest, opts ...grpc.CallOption) (*UpdateAudienceResponse, error) {
	out := new(UpdateAudienceResponse)
	err := c.cc.Invoke(ctx, "/audience.v1.AudienceService/UpdateAudience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceServiceClient) DeleteAudience(ctx context.Context, in *DeleteAudienceRequest, opts ...grpc.CallOption) (*DeleteAudienceResponse, error) {
	out := new(DeleteAudienceResponse)
	err := c.cc.Invoke(ctx, "/audience.v1.AudienceService/DeleteAudience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceServiceClient) ListAudiences(ctx context.Context, in *ListAudiencesRequest, opts ...grpc.CallOption) (*ListAudiencesResponse, error) {
	out := new(ListAudiencesResponse)
	err := c.cc.Invoke(ctx, "/audience.v1.AudienceService/ListAudiences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceServiceClient) GetAudience(ctx context.Context, in *GetAudienceRequest, opts ...grpc.CallOption) (*GetAudienceResponse, error) {
	out := new(GetAudienceResponse)
	err := c.cc.Invoke(ctx, "/audience.v1.AudienceService/GetAudience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudienceServiceServer is the server API for AudienceService service.
// All implementations should embed UnimplementedAudienceServiceServer
// for forward compatibility
type AudienceServiceServer interface {
	BulkCreateAudiences(context.Context, *CreateAudienceRequest) (*CreateAudienceResponse, error)
	CreateAudience(context.Context, *CreateAudienceRequest) (*CreateAudienceResponse, error)
	UpdateAudience(context.Context, *UpdateAudienceRequest) (*UpdateAudienceResponse, error)
	DeleteAudience(context.Context, *DeleteAudienceRequest) (*DeleteAudienceResponse, error)
	ListAudiences(context.Context, *ListAudiencesRequest) (*ListAudiencesResponse, error)
	GetAudience(context.Context, *GetAudienceRequest) (*GetAudienceResponse, error)
}

// UnimplementedAudienceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAudienceServiceServer struct {
}

func (UnimplementedAudienceServiceServer) BulkCreateAudiences(context.Context, *CreateAudienceRequest) (*CreateAudienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkCreateAudiences not implemented")
}
func (UnimplementedAudienceServiceServer) CreateAudience(context.Context, *CreateAudienceRequest) (*CreateAudienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAudience not implemented")
}
func (UnimplementedAudienceServiceServer) UpdateAudience(context.Context, *UpdateAudienceRequest) (*UpdateAudienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAudience not implemented")
}
func (UnimplementedAudienceServiceServer) DeleteAudience(context.Context, *DeleteAudienceRequest) (*DeleteAudienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAudience not implemented")
}
func (UnimplementedAudienceServiceServer) ListAudiences(context.Context, *ListAudiencesRequest) (*ListAudiencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAudiences not implemented")
}
func (UnimplementedAudienceServiceServer) GetAudience(context.Context, *GetAudienceRequest) (*GetAudienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAudience not implemented")
}

// UnsafeAudienceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudienceServiceServer will
// result in compilation errors.
type UnsafeAudienceServiceServer interface {
	mustEmbedUnimplementedAudienceServiceServer()
}

func RegisterAudienceServiceServer(s grpc.ServiceRegistrar, srv AudienceServiceServer) {
	s.RegisterService(&AudienceService_ServiceDesc, srv)
}

func _AudienceService_BulkCreateAudiences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAudienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceServiceServer).BulkCreateAudiences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audience.v1.AudienceService/BulkCreateAudiences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceServiceServer).BulkCreateAudiences(ctx, req.(*CreateAudienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceService_CreateAudience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAudienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceServiceServer).CreateAudience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audience.v1.AudienceService/CreateAudience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceServiceServer).CreateAudience(ctx, req.(*CreateAudienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceService_UpdateAudience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAudienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceServiceServer).UpdateAudience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audience.v1.AudienceService/UpdateAudience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceServiceServer).UpdateAudience(ctx, req.(*UpdateAudienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceService_DeleteAudience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAudienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceServiceServer).DeleteAudience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audience.v1.AudienceService/DeleteAudience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceServiceServer).DeleteAudience(ctx, req.(*DeleteAudienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceService_ListAudiences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAudiencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceServiceServer).ListAudiences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audience.v1.AudienceService/ListAudiences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceServiceServer).ListAudiences(ctx, req.(*ListAudiencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceService_GetAudience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAudienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceServiceServer).GetAudience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audience.v1.AudienceService/GetAudience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceServiceServer).GetAudience(ctx, req.(*GetAudienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AudienceService_ServiceDesc is the grpc.ServiceDesc for AudienceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudienceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "audience.v1.AudienceService",
	HandlerType: (*AudienceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BulkCreateAudiences",
			Handler:    _AudienceService_BulkCreateAudiences_Handler,
		},
		{
			MethodName: "CreateAudience",
			Handler:    _AudienceService_CreateAudience_Handler,
		},
		{
			MethodName: "UpdateAudience",
			Handler:    _AudienceService_UpdateAudience_Handler,
		},
		{
			MethodName: "DeleteAudience",
			Handler:    _AudienceService_DeleteAudience_Handler,
		},
		{
			MethodName: "ListAudiences",
			Handler:    _AudienceService_ListAudiences_Handler,
		},
		{
			MethodName: "GetAudience",
			Handler:    _AudienceService_GetAudience_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audience/v1/audience.proto",
}