// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: featureflag/v1/featureflag.proto

package featureflagpb

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

// FeatureServiceClient is the client API for FeatureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeatureServiceClient interface {
	// CreateFeature
	//
	// Add new feature resource in your feature flags system.
	// Can be toggled on and off on [UpdateFeature](/#/UpdateFeature).
	CreateFeature(ctx context.Context, in *CreateFeatureRequest, opts ...grpc.CallOption) (*CreateFeatureResponse, error)
	// GetFeature
	//
	// Retrieve feature resource by it's name. It will return feature object,
	// and enabled or disabled state.
	GetFeature(ctx context.Context, in *GetFeatureRequest, opts ...grpc.CallOption) (*GetFeatureResponse, error)
	// ListFeatures
	//
	// Retrieve feature collections.
	ListFeatures(ctx context.Context, in *ListFeaturesRequest, opts ...grpc.CallOption) (*ListFeaturesResponse, error)
	// UpdateFeature
	//
	// Update a feature resource, usually being used to toggle on and off.
	UpdateFeature(ctx context.Context, in *UpdateFeatureRequest, opts ...grpc.CallOption) (*UpdateFeatureResponse, error)
	// DeleteFeature
	//
	// Delete a feature resource.
	DeleteFeature(ctx context.Context, in *DeleteFeatureRequest, opts ...grpc.CallOption) (*DeleteFeatureResponse, error)
}

type featureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureServiceClient(cc grpc.ClientConnInterface) FeatureServiceClient {
	return &featureServiceClient{cc}
}

func (c *featureServiceClient) CreateFeature(ctx context.Context, in *CreateFeatureRequest, opts ...grpc.CallOption) (*CreateFeatureResponse, error) {
	out := new(CreateFeatureResponse)
	err := c.cc.Invoke(ctx, "/featureflag.v1.FeatureService/CreateFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureServiceClient) GetFeature(ctx context.Context, in *GetFeatureRequest, opts ...grpc.CallOption) (*GetFeatureResponse, error) {
	out := new(GetFeatureResponse)
	err := c.cc.Invoke(ctx, "/featureflag.v1.FeatureService/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureServiceClient) ListFeatures(ctx context.Context, in *ListFeaturesRequest, opts ...grpc.CallOption) (*ListFeaturesResponse, error) {
	out := new(ListFeaturesResponse)
	err := c.cc.Invoke(ctx, "/featureflag.v1.FeatureService/ListFeatures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureServiceClient) UpdateFeature(ctx context.Context, in *UpdateFeatureRequest, opts ...grpc.CallOption) (*UpdateFeatureResponse, error) {
	out := new(UpdateFeatureResponse)
	err := c.cc.Invoke(ctx, "/featureflag.v1.FeatureService/UpdateFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureServiceClient) DeleteFeature(ctx context.Context, in *DeleteFeatureRequest, opts ...grpc.CallOption) (*DeleteFeatureResponse, error) {
	out := new(DeleteFeatureResponse)
	err := c.cc.Invoke(ctx, "/featureflag.v1.FeatureService/DeleteFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeatureServiceServer is the server API for FeatureService service.
// All implementations should embed UnimplementedFeatureServiceServer
// for forward compatibility
type FeatureServiceServer interface {
	// CreateFeature
	//
	// Add new feature resource in your feature flags system.
	// Can be toggled on and off on [UpdateFeature](/#/UpdateFeature).
	CreateFeature(context.Context, *CreateFeatureRequest) (*CreateFeatureResponse, error)
	// GetFeature
	//
	// Retrieve feature resource by it's name. It will return feature object,
	// and enabled or disabled state.
	GetFeature(context.Context, *GetFeatureRequest) (*GetFeatureResponse, error)
	// ListFeatures
	//
	// Retrieve feature collections.
	ListFeatures(context.Context, *ListFeaturesRequest) (*ListFeaturesResponse, error)
	// UpdateFeature
	//
	// Update a feature resource, usually being used to toggle on and off.
	UpdateFeature(context.Context, *UpdateFeatureRequest) (*UpdateFeatureResponse, error)
	// DeleteFeature
	//
	// Delete a feature resource.
	DeleteFeature(context.Context, *DeleteFeatureRequest) (*DeleteFeatureResponse, error)
}

// UnimplementedFeatureServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFeatureServiceServer struct {
}

func (UnimplementedFeatureServiceServer) CreateFeature(context.Context, *CreateFeatureRequest) (*CreateFeatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeature not implemented")
}
func (UnimplementedFeatureServiceServer) GetFeature(context.Context, *GetFeatureRequest) (*GetFeatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (UnimplementedFeatureServiceServer) ListFeatures(context.Context, *ListFeaturesRequest) (*ListFeaturesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeatures not implemented")
}
func (UnimplementedFeatureServiceServer) UpdateFeature(context.Context, *UpdateFeatureRequest) (*UpdateFeatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFeature not implemented")
}
func (UnimplementedFeatureServiceServer) DeleteFeature(context.Context, *DeleteFeatureRequest) (*DeleteFeatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeature not implemented")
}

// UnsafeFeatureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeatureServiceServer will
// result in compilation errors.
type UnsafeFeatureServiceServer interface {
	mustEmbedUnimplementedFeatureServiceServer()
}

func RegisterFeatureServiceServer(s grpc.ServiceRegistrar, srv FeatureServiceServer) {
	s.RegisterService(&FeatureService_ServiceDesc, srv)
}

func _FeatureService_CreateFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureServiceServer).CreateFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/featureflag.v1.FeatureService/CreateFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureServiceServer).CreateFeature(ctx, req.(*CreateFeatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureService_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureServiceServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/featureflag.v1.FeatureService/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureServiceServer).GetFeature(ctx, req.(*GetFeatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureService_ListFeatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureServiceServer).ListFeatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/featureflag.v1.FeatureService/ListFeatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureServiceServer).ListFeatures(ctx, req.(*ListFeaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureService_UpdateFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFeatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureServiceServer).UpdateFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/featureflag.v1.FeatureService/UpdateFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureServiceServer).UpdateFeature(ctx, req.(*UpdateFeatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureService_DeleteFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureServiceServer).DeleteFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/featureflag.v1.FeatureService/DeleteFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureServiceServer).DeleteFeature(ctx, req.(*DeleteFeatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeatureService_ServiceDesc is the grpc.ServiceDesc for FeatureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeatureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "featureflag.v1.FeatureService",
	HandlerType: (*FeatureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeature",
			Handler:    _FeatureService_CreateFeature_Handler,
		},
		{
			MethodName: "GetFeature",
			Handler:    _FeatureService_GetFeature_Handler,
		},
		{
			MethodName: "ListFeatures",
			Handler:    _FeatureService_ListFeatures_Handler,
		},
		{
			MethodName: "UpdateFeature",
			Handler:    _FeatureService_UpdateFeature_Handler,
		},
		{
			MethodName: "DeleteFeature",
			Handler:    _FeatureService_DeleteFeature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "featureflag/v1/featureflag.proto",
}
