// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: svc/sso/sso_user.proto

package sso

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent        string `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	mi := &file_svc_sso_sso_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_sso_sso_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_svc_sso_sso_user_proto_rawDescGZIP(), []int{0}
}

func (x *RefreshTokenRequest) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type DeleteOwnUserSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *DeleteOwnUserSessionRequest) Reset() {
	*x = DeleteOwnUserSessionRequest{}
	mi := &file_svc_sso_sso_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOwnUserSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOwnUserSessionRequest) ProtoMessage() {}

func (x *DeleteOwnUserSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_sso_sso_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOwnUserSessionRequest.ProtoReflect.Descriptor instead.
func (*DeleteOwnUserSessionRequest) Descriptor() ([]byte, []int) {
	return file_svc_sso_sso_user_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteOwnUserSessionRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type GoogleLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GoogleLoginResponse) Reset() {
	*x = GoogleLoginResponse{}
	mi := &file_svc_sso_sso_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoogleLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleLoginResponse) ProtoMessage() {}

func (x *GoogleLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_sso_sso_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleLoginResponse.ProtoReflect.Descriptor instead.
func (*GoogleLoginResponse) Descriptor() ([]byte, []int) {
	return file_svc_sso_sso_user_proto_rawDescGZIP(), []int{2}
}

func (x *GoogleLoginResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GoogleCallbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GoogleCallbackRequest) Reset() {
	*x = GoogleCallbackRequest{}
	mi := &file_svc_sso_sso_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoogleCallbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleCallbackRequest) ProtoMessage() {}

func (x *GoogleCallbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_sso_sso_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleCallbackRequest.ProtoReflect.Descriptor instead.
func (*GoogleCallbackRequest) Descriptor() ([]byte, []int) {
	return file_svc_sso_sso_user_proto_rawDescGZIP(), []int{3}
}

func (x *GoogleCallbackRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_svc_sso_sso_user_proto protoreflect.FileDescriptor

var file_svc_sso_sso_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x76, 0x63, 0x2f, 0x73, 0x73, 0x6f, 0x2f, 0x73, 0x73, 0x6f, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x73, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x76, 0x63, 0x2f,
	0x73, 0x73, 0x6f, 0x2f, 0x73, 0x73, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x13, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3c, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x2b,
	0x0a, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xc5, 0x04, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x4f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x09, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0b,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x0e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12,
	0x1a, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x73, 0x73,
	0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x61, 0x69, 0x72, 0x12, 0x38, 0x0a, 0x06,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x61, 0x69,
	0x72, 0x12, 0x39, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c,
	0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x73, 0x6f,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4b, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2d, 0x6c, 0x61, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x73, 0x73, 0x6f, 0x3b, 0x73, 0x73, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_sso_sso_user_proto_rawDescOnce sync.Once
	file_svc_sso_sso_user_proto_rawDescData = file_svc_sso_sso_user_proto_rawDesc
)

func file_svc_sso_sso_user_proto_rawDescGZIP() []byte {
	file_svc_sso_sso_user_proto_rawDescOnce.Do(func() {
		file_svc_sso_sso_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_sso_sso_user_proto_rawDescData)
	})
	return file_svc_sso_sso_user_proto_rawDescData
}

var file_svc_sso_sso_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_svc_sso_sso_user_proto_goTypes = []any{
	(*RefreshTokenRequest)(nil),         // 0: sso.RefreshTokenRequest
	(*DeleteOwnUserSessionRequest)(nil), // 1: sso.DeleteOwnUserSessionRequest
	(*GoogleLoginResponse)(nil),         // 2: sso.GoogleLoginResponse
	(*GoogleCallbackRequest)(nil),       // 3: sso.GoogleCallbackRequest
	(*emptypb.Empty)(nil),               // 4: google.protobuf.Empty
	(*User)(nil),                        // 5: sso.User
	(*TokensPair)(nil),                  // 6: sso.TokensPair
	(*Session)(nil),                     // 7: sso.Session
	(*SessionsList)(nil),                // 8: sso.SessionsList
}
var file_svc_sso_sso_user_proto_depIdxs = []int32{
	4, // 0: sso.UserService.GetOwnUser:input_type -> google.protobuf.Empty
	4, // 1: sso.UserService.GoogleLogin:input_type -> google.protobuf.Empty
	3, // 2: sso.UserService.GoogleCallback:input_type -> sso.GoogleCallbackRequest
	4, // 3: sso.UserService.Logout:input_type -> google.protobuf.Empty
	0, // 4: sso.UserService.RefreshToken:input_type -> sso.RefreshTokenRequest
	4, // 5: sso.UserService.GetOwnUserSession:input_type -> google.protobuf.Empty
	4, // 6: sso.UserService.GetOwnUserSessions:input_type -> google.protobuf.Empty
	1, // 7: sso.UserService.DeleteOwnUserSession:input_type -> sso.DeleteOwnUserSessionRequest
	4, // 8: sso.UserService.DeleteOwnUserSessions:input_type -> google.protobuf.Empty
	5, // 9: sso.UserService.GetOwnUser:output_type -> sso.User
	2, // 10: sso.UserService.GoogleLogin:output_type -> sso.GoogleLoginResponse
	6, // 11: sso.UserService.GoogleCallback:output_type -> sso.TokensPair
	4, // 12: sso.UserService.Logout:output_type -> google.protobuf.Empty
	6, // 13: sso.UserService.RefreshToken:output_type -> sso.TokensPair
	7, // 14: sso.UserService.GetOwnUserSession:output_type -> sso.Session
	8, // 15: sso.UserService.GetOwnUserSessions:output_type -> sso.SessionsList
	8, // 16: sso.UserService.DeleteOwnUserSession:output_type -> sso.SessionsList
	4, // 17: sso.UserService.DeleteOwnUserSessions:output_type -> google.protobuf.Empty
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svc_sso_sso_user_proto_init() }
func file_svc_sso_sso_user_proto_init() {
	if File_svc_sso_sso_user_proto != nil {
		return
	}
	file_svc_sso_sso_structs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_sso_sso_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_sso_sso_user_proto_goTypes,
		DependencyIndexes: file_svc_sso_sso_user_proto_depIdxs,
		MessageInfos:      file_svc_sso_sso_user_proto_msgTypes,
	}.Build()
	File_svc_sso_sso_user_proto = out.File
	file_svc_sso_sso_user_proto_rawDesc = nil
	file_svc_sso_sso_user_proto_goTypes = nil
	file_svc_sso_sso_user_proto_depIdxs = nil
}
