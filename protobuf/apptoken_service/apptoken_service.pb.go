// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: apptoken_service.proto

package apptoken_service

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracing_ID string `protobuf:"bytes,1,opt,name=tracing_ID,json=tracingID,proto3" json:"tracing_ID,omitempty"`
	CallerUuid string `protobuf:"bytes,2,opt,name=caller_uuid,json=callerUuid,proto3" json:"caller_uuid,omitempty"`
	AppUuid    string `protobuf:"bytes,3,opt,name=app_uuid,json=appUuid,proto3" json:"app_uuid,omitempty"`
	AppHash    string `protobuf:"bytes,4,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	AppOrigin  string `protobuf:"bytes,5,opt,name=app_origin,json=appOrigin,proto3" json:"app_origin,omitempty"`
}

func (x *IssueRequest) Reset() {
	*x = IssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apptoken_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueRequest) ProtoMessage() {}

func (x *IssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apptoken_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueRequest.ProtoReflect.Descriptor instead.
func (*IssueRequest) Descriptor() ([]byte, []int) {
	return file_apptoken_service_proto_rawDescGZIP(), []int{0}
}

func (x *IssueRequest) GetTracing_ID() string {
	if x != nil {
		return x.Tracing_ID
	}
	return ""
}

func (x *IssueRequest) GetCallerUuid() string {
	if x != nil {
		return x.CallerUuid
	}
	return ""
}

func (x *IssueRequest) GetAppUuid() string {
	if x != nil {
		return x.AppUuid
	}
	return ""
}

func (x *IssueRequest) GetAppHash() string {
	if x != nil {
		return x.AppHash
	}
	return ""
}

func (x *IssueRequest) GetAppOrigin() string {
	if x != nil {
		return x.AppOrigin
	}
	return ""
}

type IssueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token      *MetaToken `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *IssueResponse) Reset() {
	*x = IssueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apptoken_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueResponse) ProtoMessage() {}

func (x *IssueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apptoken_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueResponse.ProtoReflect.Descriptor instead.
func (*IssueResponse) Descriptor() ([]byte, []int) {
	return file_apptoken_service_proto_rawDescGZIP(), []int{1}
}

func (x *IssueResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *IssueResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *IssueResponse) GetToken() *MetaToken {
	if x != nil {
		return x.Token
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracing_ID string `protobuf:"bytes,1,opt,name=tracing_ID,json=tracingID,proto3" json:"tracing_ID,omitempty"`
	AppUuid    string `protobuf:"bytes,2,opt,name=app_uuid,json=appUuid,proto3" json:"app_uuid,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apptoken_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apptoken_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_apptoken_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetTracing_ID() string {
	if x != nil {
		return x.Tracing_ID
	}
	return ""
}

func (x *GetRequest) GetAppUuid() string {
	if x != nil {
		return x.AppUuid
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token      *MetaToken `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apptoken_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apptoken_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_apptoken_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetResponse) GetToken() *MetaToken {
	if x != nil {
		return x.Token
	}
	return nil
}

type MetaToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Exp   int64  `protobuf:"varint,2,opt,name=exp,proto3" json:"exp,omitempty"`
}

func (x *MetaToken) Reset() {
	*x = MetaToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apptoken_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaToken) ProtoMessage() {}

func (x *MetaToken) ProtoReflect() protoreflect.Message {
	mi := &file_apptoken_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaToken.ProtoReflect.Descriptor instead.
func (*MetaToken) Descriptor() ([]byte, []int) {
	return file_apptoken_service_proto_rawDescGZIP(), []int{4}
}

func (x *MetaToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MetaToken) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

var File_apptoken_service_proto protoreflect.FileDescriptor

var file_apptoken_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x0c, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x70, 0x70, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x70, 0x70, 0x55, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x22, 0x75, 0x0a, 0x0d, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x31, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67,
	0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x49, 0x44, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x55, 0x75, 0x69, 0x64, 0x22,
	0x73, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x31, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x33, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78, 0x70, 0x32, 0x9c, 0x01, 0x0a, 0x08, 0x41, 0x70,
	0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4a, 0x0a, 0x05, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12,
	0x1e, 0x2e, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apptoken_service_proto_rawDescOnce sync.Once
	file_apptoken_service_proto_rawDescData = file_apptoken_service_proto_rawDesc
)

func file_apptoken_service_proto_rawDescGZIP() []byte {
	file_apptoken_service_proto_rawDescOnce.Do(func() {
		file_apptoken_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_apptoken_service_proto_rawDescData)
	})
	return file_apptoken_service_proto_rawDescData
}

var file_apptoken_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apptoken_service_proto_goTypes = []interface{}{
	(*IssueRequest)(nil),  // 0: apptoken_service.IssueRequest
	(*IssueResponse)(nil), // 1: apptoken_service.IssueResponse
	(*GetRequest)(nil),    // 2: apptoken_service.GetRequest
	(*GetResponse)(nil),   // 3: apptoken_service.GetResponse
	(*MetaToken)(nil),     // 4: apptoken_service.MetaToken
}
var file_apptoken_service_proto_depIdxs = []int32{
	4, // 0: apptoken_service.IssueResponse.token:type_name -> apptoken_service.MetaToken
	4, // 1: apptoken_service.GetResponse.token:type_name -> apptoken_service.MetaToken
	0, // 2: apptoken_service.AppToken.Issue:input_type -> apptoken_service.IssueRequest
	2, // 3: apptoken_service.AppToken.Get:input_type -> apptoken_service.GetRequest
	1, // 4: apptoken_service.AppToken.Issue:output_type -> apptoken_service.IssueResponse
	3, // 5: apptoken_service.AppToken.Get:output_type -> apptoken_service.GetResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apptoken_service_proto_init() }
func file_apptoken_service_proto_init() {
	if File_apptoken_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apptoken_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apptoken_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apptoken_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apptoken_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apptoken_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apptoken_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apptoken_service_proto_goTypes,
		DependencyIndexes: file_apptoken_service_proto_depIdxs,
		MessageInfos:      file_apptoken_service_proto_msgTypes,
	}.Build()
	File_apptoken_service_proto = out.File
	file_apptoken_service_proto_rawDesc = nil
	file_apptoken_service_proto_goTypes = nil
	file_apptoken_service_proto_depIdxs = nil
}
