// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: pb_auth/auth.proto

package pb_auth

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
	Auth_SignUp_FullMethodName               = "/pb_auth.Auth/SignUp"
	Auth_SignIn_FullMethodName               = "/pb_auth.Auth/SignIn"
	Auth_RefreshToken_FullMethodName         = "/pb_auth.Auth/RefreshToken"
	Auth_SignOut_FullMethodName              = "/pb_auth.Auth/SignOut"
	Auth_GithubOAuth2Callback_FullMethodName = "/pb_auth.Auth/GithubOAuth2Callback"
	Auth_GoogleOAuth2Callback_FullMethodName = "/pb_auth.Auth/GoogleOAuth2Callback"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	SignUp(ctx context.Context, in *SignUpReq, opts ...grpc.CallOption) (*SignUpResp, error)
	SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error)
	RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error)
	SignOut(ctx context.Context, in *SignOutReq, opts ...grpc.CallOption) (*SignOutResp, error)
	GithubOAuth2Callback(ctx context.Context, in *GithubOAuth2CallbackReq, opts ...grpc.CallOption) (*GithubOAuth2CallbackResp, error)
	GoogleOAuth2Callback(ctx context.Context, in *GoogleOAuth2CallbackReq, opts ...grpc.CallOption) (*GoogleOAuth2CallbackResp, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) SignUp(ctx context.Context, in *SignUpReq, opts ...grpc.CallOption) (*SignUpResp, error) {
	out := new(SignUpResp)
	err := c.cc.Invoke(ctx, Auth_SignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error) {
	out := new(SignInResp)
	err := c.cc.Invoke(ctx, Auth_SignIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error) {
	out := new(RefreshTokenResp)
	err := c.cc.Invoke(ctx, Auth_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SignOut(ctx context.Context, in *SignOutReq, opts ...grpc.CallOption) (*SignOutResp, error) {
	out := new(SignOutResp)
	err := c.cc.Invoke(ctx, Auth_SignOut_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GithubOAuth2Callback(ctx context.Context, in *GithubOAuth2CallbackReq, opts ...grpc.CallOption) (*GithubOAuth2CallbackResp, error) {
	out := new(GithubOAuth2CallbackResp)
	err := c.cc.Invoke(ctx, Auth_GithubOAuth2Callback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GoogleOAuth2Callback(ctx context.Context, in *GoogleOAuth2CallbackReq, opts ...grpc.CallOption) (*GoogleOAuth2CallbackResp, error) {
	out := new(GoogleOAuth2CallbackResp)
	err := c.cc.Invoke(ctx, Auth_GoogleOAuth2Callback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	SignUp(context.Context, *SignUpReq) (*SignUpResp, error)
	SignIn(context.Context, *SignInReq) (*SignInResp, error)
	RefreshToken(context.Context, *RefreshTokenReq) (*RefreshTokenResp, error)
	SignOut(context.Context, *SignOutReq) (*SignOutResp, error)
	GithubOAuth2Callback(context.Context, *GithubOAuth2CallbackReq) (*GithubOAuth2CallbackResp, error)
	GoogleOAuth2Callback(context.Context, *GoogleOAuth2CallbackReq) (*GoogleOAuth2CallbackResp, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) SignUp(context.Context, *SignUpReq) (*SignUpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthServer) SignIn(context.Context, *SignInReq) (*SignInResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthServer) RefreshToken(context.Context, *RefreshTokenReq) (*RefreshTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthServer) SignOut(context.Context, *SignOutReq) (*SignOutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (UnimplementedAuthServer) GithubOAuth2Callback(context.Context, *GithubOAuth2CallbackReq) (*GithubOAuth2CallbackResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GithubOAuth2Callback not implemented")
}
func (UnimplementedAuthServer) GoogleOAuth2Callback(context.Context, *GoogleOAuth2CallbackReq) (*GoogleOAuth2CallbackResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoogleOAuth2Callback not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignUp(ctx, req.(*SignUpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignIn(ctx, req.(*SignInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RefreshToken(ctx, req.(*RefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_SignOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignOut(ctx, req.(*SignOutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GithubOAuth2Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GithubOAuth2CallbackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GithubOAuth2Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GithubOAuth2Callback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GithubOAuth2Callback(ctx, req.(*GithubOAuth2CallbackReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GoogleOAuth2Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoogleOAuth2CallbackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GoogleOAuth2Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GoogleOAuth2Callback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GoogleOAuth2Callback(ctx, req.(*GoogleOAuth2CallbackReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Auth_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Auth_SignIn_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Auth_RefreshToken_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _Auth_SignOut_Handler,
		},
		{
			MethodName: "GithubOAuth2Callback",
			Handler:    _Auth_GithubOAuth2Callback_Handler,
		},
		{
			MethodName: "GoogleOAuth2Callback",
			Handler:    _Auth_GoogleOAuth2Callback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_auth/auth.proto",
}