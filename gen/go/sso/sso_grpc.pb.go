// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: sso/sso.proto

package sso

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
	UserService_GetOwnUser_FullMethodName               = "/auth.UserService/GetOwnUser"
	UserService_GoogleLogin_FullMethodName              = "/auth.UserService/GoogleLogin"
	UserService_GoogleCallback_FullMethodName           = "/auth.UserService/GoogleCallback"
	UserService_Logout_FullMethodName                   = "/auth.UserService/Logout"
	UserService_RefreshToken_FullMethodName             = "/auth.UserService/RefreshToken"
	UserService_GetOwnUserSession_FullMethodName        = "/auth.UserService/GetOwnUserSession"
	UserService_GetOwnUserSessions_FullMethodName       = "/auth.UserService/GetOwnUserSessions"
	UserService_DeleteOwnUserSession_FullMethodName     = "/auth.UserService/DeleteOwnUserSession"
	UserService_TerminateOwnUserSessions_FullMethodName = "/auth.UserService/TerminateOwnUserSessions"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetOwnUser(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserResponse, error)
	// Google OAuth operations
	GoogleLogin(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GoogleLoginResponse, error)
	GoogleCallback(ctx context.Context, in *GoogleCallbackRequest, opts ...grpc.CallOption) (*TokensPairResponse, error)
	Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*TokensPairResponse, error)
	GetOwnUserSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SessionResponse, error)
	GetOwnUserSessions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SessionsListResponse, error)
	DeleteOwnUserSession(ctx context.Context, in *DeleteOwnUserSessionRequest, opts ...grpc.CallOption) (*Empty, error)
	TerminateOwnUserSessions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetOwnUser(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_GetOwnUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GoogleLogin(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GoogleLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GoogleLoginResponse)
	err := c.cc.Invoke(ctx, UserService_GoogleLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GoogleCallback(ctx context.Context, in *GoogleCallbackRequest, opts ...grpc.CallOption) (*TokensPairResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokensPairResponse)
	err := c.cc.Invoke(ctx, UserService_GoogleCallback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*TokensPairResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokensPairResponse)
	err := c.cc.Invoke(ctx, UserService_RefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetOwnUserSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, UserService_GetOwnUserSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetOwnUserSessions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SessionsListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionsListResponse)
	err := c.cc.Invoke(ctx, UserService_GetOwnUserSessions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteOwnUserSession(ctx context.Context, in *DeleteOwnUserSessionRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_DeleteOwnUserSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) TerminateOwnUserSessions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_TerminateOwnUserSessions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	GetOwnUser(context.Context, *Empty) (*UserResponse, error)
	// Google OAuth operations
	GoogleLogin(context.Context, *Empty) (*GoogleLoginResponse, error)
	GoogleCallback(context.Context, *GoogleCallbackRequest) (*TokensPairResponse, error)
	Logout(context.Context, *Empty) (*Empty, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*TokensPairResponse, error)
	GetOwnUserSession(context.Context, *Empty) (*SessionResponse, error)
	GetOwnUserSessions(context.Context, *Empty) (*SessionsListResponse, error)
	DeleteOwnUserSession(context.Context, *DeleteOwnUserSessionRequest) (*Empty, error)
	TerminateOwnUserSessions(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetOwnUser(context.Context, *Empty) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwnUser not implemented")
}
func (UnimplementedUserServiceServer) GoogleLogin(context.Context, *Empty) (*GoogleLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoogleLogin not implemented")
}
func (UnimplementedUserServiceServer) GoogleCallback(context.Context, *GoogleCallbackRequest) (*TokensPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoogleCallback not implemented")
}
func (UnimplementedUserServiceServer) Logout(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*TokensPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedUserServiceServer) GetOwnUserSession(context.Context, *Empty) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwnUserSession not implemented")
}
func (UnimplementedUserServiceServer) GetOwnUserSessions(context.Context, *Empty) (*SessionsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwnUserSessions not implemented")
}
func (UnimplementedUserServiceServer) DeleteOwnUserSession(context.Context, *DeleteOwnUserSessionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOwnUserSession not implemented")
}
func (UnimplementedUserServiceServer) TerminateOwnUserSessions(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerminateOwnUserSessions not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetOwnUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetOwnUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetOwnUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetOwnUser(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GoogleLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GoogleLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GoogleLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GoogleLogin(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GoogleCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoogleCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GoogleCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GoogleCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GoogleCallback(ctx, req.(*GoogleCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Logout(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetOwnUserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetOwnUserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetOwnUserSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetOwnUserSession(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetOwnUserSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetOwnUserSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetOwnUserSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetOwnUserSessions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteOwnUserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOwnUserSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteOwnUserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteOwnUserSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteOwnUserSession(ctx, req.(*DeleteOwnUserSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_TerminateOwnUserSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).TerminateOwnUserSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_TerminateOwnUserSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).TerminateOwnUserSessions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOwnUser",
			Handler:    _UserService_GetOwnUser_Handler,
		},
		{
			MethodName: "GoogleLogin",
			Handler:    _UserService_GoogleLogin_Handler,
		},
		{
			MethodName: "GoogleCallback",
			Handler:    _UserService_GoogleCallback_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserService_Logout_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _UserService_RefreshToken_Handler,
		},
		{
			MethodName: "GetOwnUserSession",
			Handler:    _UserService_GetOwnUserSession_Handler,
		},
		{
			MethodName: "GetOwnUserSessions",
			Handler:    _UserService_GetOwnUserSessions_Handler,
		},
		{
			MethodName: "DeleteOwnUserSession",
			Handler:    _UserService_DeleteOwnUserSession_Handler,
		},
		{
			MethodName: "TerminateOwnUserSessions",
			Handler:    _UserService_TerminateOwnUserSessions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}

const (
	AdminService_AdminGetUser_FullMethodName               = "/auth.AdminService/AdminGetUser"
	AdminService_AdminUpdateUserRole_FullMethodName        = "/auth.AdminService/AdminUpdateUserRole"
	AdminService_AdminGetUserSessions_FullMethodName       = "/auth.AdminService/AdminGetUserSessions"
	AdminService_AdminGetUserSession_FullMethodName        = "/auth.AdminService/AdminGetUserSession"
	AdminService_AdminTerminateUserSessions_FullMethodName = "/auth.AdminService/AdminTerminateUserSessions"
	AdminService_AdminDeleteUserSession_FullMethodName     = "/auth.AdminService/AdminDeleteUserSession"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	AdminGetUser(ctx context.Context, in *AdminGetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	AdminUpdateUserRole(ctx context.Context, in *AdminUpdateUserRoleRequest, opts ...grpc.CallOption) (*UserResponse, error)
	AdminGetUserSessions(ctx context.Context, in *AdminGetUserSessionsRequest, opts ...grpc.CallOption) (*SessionsListResponse, error)
	AdminGetUserSession(ctx context.Context, in *AdminGetUserSessionRequest, opts ...grpc.CallOption) (*SessionResponse, error)
	AdminTerminateUserSessions(ctx context.Context, in *AdminTerminateUserSessionsRequest, opts ...grpc.CallOption) (*Empty, error)
	AdminDeleteUserSession(ctx context.Context, in *AdminDeleteUserSessionRequest, opts ...grpc.CallOption) (*Empty, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) AdminGetUser(ctx context.Context, in *AdminGetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, AdminService_AdminGetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminUpdateUserRole(ctx context.Context, in *AdminUpdateUserRoleRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, AdminService_AdminUpdateUserRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminGetUserSessions(ctx context.Context, in *AdminGetUserSessionsRequest, opts ...grpc.CallOption) (*SessionsListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionsListResponse)
	err := c.cc.Invoke(ctx, AdminService_AdminGetUserSessions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminGetUserSession(ctx context.Context, in *AdminGetUserSessionRequest, opts ...grpc.CallOption) (*SessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, AdminService_AdminGetUserSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminTerminateUserSessions(ctx context.Context, in *AdminTerminateUserSessionsRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AdminService_AdminTerminateUserSessions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminDeleteUserSession(ctx context.Context, in *AdminDeleteUserSessionRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AdminService_AdminDeleteUserSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility.
type AdminServiceServer interface {
	AdminGetUser(context.Context, *AdminGetUserRequest) (*UserResponse, error)
	AdminUpdateUserRole(context.Context, *AdminUpdateUserRoleRequest) (*UserResponse, error)
	AdminGetUserSessions(context.Context, *AdminGetUserSessionsRequest) (*SessionsListResponse, error)
	AdminGetUserSession(context.Context, *AdminGetUserSessionRequest) (*SessionResponse, error)
	AdminTerminateUserSessions(context.Context, *AdminTerminateUserSessionsRequest) (*Empty, error)
	AdminDeleteUserSession(context.Context, *AdminDeleteUserSessionRequest) (*Empty, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdminServiceServer struct{}

func (UnimplementedAdminServiceServer) AdminGetUser(context.Context, *AdminGetUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetUser not implemented")
}
func (UnimplementedAdminServiceServer) AdminUpdateUserRole(context.Context, *AdminUpdateUserRoleRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateUserRole not implemented")
}
func (UnimplementedAdminServiceServer) AdminGetUserSessions(context.Context, *AdminGetUserSessionsRequest) (*SessionsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetUserSessions not implemented")
}
func (UnimplementedAdminServiceServer) AdminGetUserSession(context.Context, *AdminGetUserSessionRequest) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetUserSession not implemented")
}
func (UnimplementedAdminServiceServer) AdminTerminateUserSessions(context.Context, *AdminTerminateUserSessionsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminTerminateUserSessions not implemented")
}
func (UnimplementedAdminServiceServer) AdminDeleteUserSession(context.Context, *AdminDeleteUserSessionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteUserSession not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}
func (UnimplementedAdminServiceServer) testEmbeddedByValue()                      {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdminServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_AdminGetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminGetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AdminGetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminGetUser(ctx, req.(*AdminGetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminUpdateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminUpdateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AdminUpdateUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminUpdateUserRole(ctx, req.(*AdminUpdateUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminGetUserSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetUserSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminGetUserSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AdminGetUserSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminGetUserSessions(ctx, req.(*AdminGetUserSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminGetUserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetUserSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminGetUserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AdminGetUserSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminGetUserSession(ctx, req.(*AdminGetUserSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminTerminateUserSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminTerminateUserSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminTerminateUserSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AdminTerminateUserSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminTerminateUserSessions(ctx, req.(*AdminTerminateUserSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminDeleteUserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteUserSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminDeleteUserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AdminDeleteUserSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminDeleteUserSession(ctx, req.(*AdminDeleteUserSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminGetUser",
			Handler:    _AdminService_AdminGetUser_Handler,
		},
		{
			MethodName: "AdminUpdateUserRole",
			Handler:    _AdminService_AdminUpdateUserRole_Handler,
		},
		{
			MethodName: "AdminGetUserSessions",
			Handler:    _AdminService_AdminGetUserSessions_Handler,
		},
		{
			MethodName: "AdminGetUserSession",
			Handler:    _AdminService_AdminGetUserSession_Handler,
		},
		{
			MethodName: "AdminTerminateUserSessions",
			Handler:    _AdminService_AdminTerminateUserSessions_Handler,
		},
		{
			MethodName: "AdminDeleteUserSession",
			Handler:    _AdminService_AdminDeleteUserSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}
