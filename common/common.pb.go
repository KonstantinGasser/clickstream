// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: common/common.proto

package common

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

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello.ProtoReflect.Descriptor instead.
func (*Hello) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *Hello) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Hello) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type AppInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	URL         string    `protobuf:"bytes,3,opt,name=URL,proto3" json:"URL,omitempty"`
	Description string    `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Owner       *UserInfo `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Member      []string  `protobuf:"bytes,6,rep,name=member,proto3" json:"member,omitempty"`
}

func (x *AppInfo) Reset() {
	*x = AppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfo) ProtoMessage() {}

func (x *AppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInfo.ProtoReflect.Descriptor instead.
func (*AppInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *AppInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AppInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppInfo) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *AppInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AppInfo) GetOwner() *UserInfo {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *AppInfo) GetMember() []string {
	if x != nil {
		return x.Member
	}
	return nil
}

type AppMetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *AppMetaInfo) Reset() {
	*x = AppMetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppMetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMetaInfo) ProtoMessage() {}

func (x *AppMetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMetaInfo.ProtoReflect.Descriptor instead.
func (*AppMetaInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *AppMetaInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppMetaInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ConfigInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Funnel   []*Funnel   `protobuf:"bytes,1,rep,name=funnel,proto3" json:"funnel,omitempty"`
	Campaign []*Campaign `protobuf:"bytes,2,rep,name=campaign,proto3" json:"campaign,omitempty"`
	BtnTime  []*BtnTime  `protobuf:"bytes,3,rep,name=btn_time,json=btnTime,proto3" json:"btn_time,omitempty"`
}

func (x *ConfigInfo) Reset() {
	*x = ConfigInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigInfo) ProtoMessage() {}

func (x *ConfigInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigInfo.ProtoReflect.Descriptor instead.
func (*ConfigInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigInfo) GetFunnel() []*Funnel {
	if x != nil {
		return x.Funnel
	}
	return nil
}

func (x *ConfigInfo) GetCampaign() []*Campaign {
	if x != nil {
		return x.Campaign
	}
	return nil
}

func (x *ConfigInfo) GetBtnTime() []*BtnTime {
	if x != nil {
		return x.BtnTime
	}
	return nil
}

type Funnel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Transition string `protobuf:"bytes,3,opt,name=transition,proto3" json:"transition,omitempty"`
}

func (x *Funnel) Reset() {
	*x = Funnel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Funnel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Funnel) ProtoMessage() {}

func (x *Funnel) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Funnel.ProtoReflect.Descriptor instead.
func (*Funnel) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *Funnel) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Funnel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Funnel) GetTransition() string {
	if x != nil {
		return x.Transition
	}
	return ""
}

type Campaign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Prefix string `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *Campaign) Reset() {
	*x = Campaign{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Campaign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Campaign) ProtoMessage() {}

func (x *Campaign) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Campaign.ProtoReflect.Descriptor instead.
func (*Campaign) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{5}
}

func (x *Campaign) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Campaign) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Campaign) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type BtnTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BtnName string `protobuf:"bytes,3,opt,name=btn_name,json=btnName,proto3" json:"btn_name,omitempty"`
}

func (x *BtnTime) Reset() {
	*x = BtnTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BtnTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BtnTime) ProtoMessage() {}

func (x *BtnTime) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BtnTime.ProtoReflect.Descriptor instead.
func (*BtnTime) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{6}
}

func (x *BtnTime) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BtnTime) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BtnTime) GetBtnName() string {
	if x != nil {
		return x.BtnName
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Username      string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName     string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	OrgnDomain    string `protobuf:"bytes,5,opt,name=orgn_domain,json=orgnDomain,proto3" json:"orgn_domain,omitempty"`
	OrgnPosition  string `protobuf:"bytes,6,opt,name=orgn_position,json=orgnPosition,proto3" json:"orgn_position,omitempty"`
	ProfileImgUrl string `protobuf:"bytes,7,opt,name=profile_img_url,json=profileImgUrl,proto3" json:"profile_img_url,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{7}
}

func (x *UserInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserInfo) GetOrgnDomain() string {
	if x != nil {
		return x.OrgnDomain
	}
	return ""
}

func (x *UserInfo) GetOrgnPosition() string {
	if x != nil {
		return x.OrgnPosition
	}
	return ""
}

func (x *UserInfo) GetProfileImgUrl() string {
	if x != nil {
		return x.ProfileImgUrl
	}
	return ""
}

type UserMetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UserMetaInfo) Reset() {
	*x = UserMetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMetaInfo) ProtoMessage() {}

func (x *UserMetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMetaInfo.ProtoReflect.Descriptor instead.
func (*UserMetaInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{8}
}

func (x *UserMetaInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserMetaInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type AppTokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Exp   int64  `protobuf:"varint,2,opt,name=exp,proto3" json:"exp,omitempty"`
}

func (x *AppTokenInfo) Reset() {
	*x = AppTokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppTokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppTokenInfo) ProtoMessage() {}

func (x *AppTokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppTokenInfo.ProtoReflect.Descriptor instead.
func (*AppTokenInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{9}
}

func (x *AppTokenInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AppTokenInfo) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x2d, 0x0a,
	0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xa5, 0x01, 0x0a,
	0x07, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55,
	0x52, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x06, 0x66, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x12, 0x2a, 0x0a, 0x08, 0x62, 0x74, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x74, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x07, 0x62, 0x74, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x06,
	0x46, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x08, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x22, 0x48, 0x0a, 0x07, 0x42, 0x74, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x74, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x74, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xe4, 0x01, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x67, 0x6e, 0x5f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x67, 0x6e,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x6e, 0x5f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f,
	0x72, 0x67, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x67,
	0x55, 0x72, 0x6c, 0x22, 0x3e, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78, 0x70, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x69, 0x6e, 0x47, 0x61, 0x73, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6c,
	0x61, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_common_common_proto_goTypes = []interface{}{
	(*Hello)(nil),        // 0: common.Hello
	(*AppInfo)(nil),      // 1: common.AppInfo
	(*AppMetaInfo)(nil),  // 2: common.AppMetaInfo
	(*ConfigInfo)(nil),   // 3: common.ConfigInfo
	(*Funnel)(nil),       // 4: common.Funnel
	(*Campaign)(nil),     // 5: common.Campaign
	(*BtnTime)(nil),      // 6: common.BtnTime
	(*UserInfo)(nil),     // 7: common.UserInfo
	(*UserMetaInfo)(nil), // 8: common.UserMetaInfo
	(*AppTokenInfo)(nil), // 9: common.AppTokenInfo
}
var file_common_common_proto_depIdxs = []int32{
	7, // 0: common.AppInfo.owner:type_name -> common.UserInfo
	4, // 1: common.ConfigInfo.funnel:type_name -> common.Funnel
	5, // 2: common.ConfigInfo.campaign:type_name -> common.Campaign
	6, // 3: common.ConfigInfo.btn_time:type_name -> common.BtnTime
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello); i {
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
		file_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInfo); i {
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
		file_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppMetaInfo); i {
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
		file_common_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigInfo); i {
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
		file_common_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Funnel); i {
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
		file_common_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Campaign); i {
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
		file_common_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BtnTime); i {
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
		file_common_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_common_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMetaInfo); i {
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
		file_common_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppTokenInfo); i {
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
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
