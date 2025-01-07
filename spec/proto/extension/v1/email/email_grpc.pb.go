// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: email.proto

package email

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

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailServiceClient interface {
	// Send an email with template
	SendEmailWithTemplate(ctx context.Context, in *SendEmailWithTemplateRequest, opts ...grpc.CallOption) (*SendEmailWithTemplateResponse, error)
	// Send an email with raw content instead of using templates.
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) SendEmailWithTemplate(ctx context.Context, in *SendEmailWithTemplateRequest, opts ...grpc.CallOption) (*SendEmailWithTemplateResponse, error) {
	out := new(SendEmailWithTemplateResponse)
	err := c.cc.Invoke(ctx, "/spec.proto.extension.v1.email.EmailService/SendEmailWithTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error) {
	out := new(SendEmailResponse)
	err := c.cc.Invoke(ctx, "/spec.proto.extension.v1.email.EmailService/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
// All implementations should embed UnimplementedEmailServiceServer
// for forward compatibility
type EmailServiceServer interface {
	// Send an email with template
	SendEmailWithTemplate(context.Context, *SendEmailWithTemplateRequest) (*SendEmailWithTemplateResponse, error)
	// Send an email with raw content instead of using templates.
	SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error)
}

// UnimplementedEmailServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (UnimplementedEmailServiceServer) SendEmailWithTemplate(context.Context, *SendEmailWithTemplateRequest) (*SendEmailWithTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailWithTemplate not implemented")
}
func (UnimplementedEmailServiceServer) SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}

// UnsafeEmailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServiceServer will
// result in compilation errors.
type UnsafeEmailServiceServer interface {
	mustEmbedUnimplementedEmailServiceServer()
}

func RegisterEmailServiceServer(s grpc.ServiceRegistrar, srv EmailServiceServer) {
	s.RegisterService(&EmailService_ServiceDesc, srv)
}

func _EmailService_SendEmailWithTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailWithTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendEmailWithTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.proto.extension.v1.email.EmailService/SendEmailWithTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendEmailWithTemplate(ctx, req.(*SendEmailWithTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.proto.extension.v1.email.EmailService/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailService_ServiceDesc is the grpc.ServiceDesc for EmailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spec.proto.extension.v1.email.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmailWithTemplate",
			Handler:    _EmailService_SendEmailWithTemplate_Handler,
		},
		{
			MethodName: "SendEmail",
			Handler:    _EmailService_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}
