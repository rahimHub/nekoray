// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: libcore.proto

package gen

import (
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

type TestMode int32

const (
	TestMode_TcpPing  TestMode = 0
	TestMode_UrlTest  TestMode = 1
	TestMode_FullTest TestMode = 2
)

// Enum value maps for TestMode.
var (
	TestMode_name = map[int32]string{
		0: "TcpPing",
		1: "UrlTest",
		2: "FullTest",
	}
	TestMode_value = map[string]int32{
		"TcpPing":  0,
		"UrlTest":  1,
		"FullTest": 2,
	}
)

func (x TestMode) Enum() *TestMode {
	p := new(TestMode)
	*p = x
	return p
}

func (x TestMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestMode) Descriptor() protoreflect.EnumDescriptor {
	return file_libcore_proto_enumTypes[0].Descriptor()
}

func (TestMode) Type() protoreflect.EnumType {
	return &file_libcore_proto_enumTypes[0]
}

func (x TestMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestMode.Descriptor instead.
func (TestMode) EnumDescriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{0}
}

type UpdateAction int32

const (
	UpdateAction_Check    UpdateAction = 0
	UpdateAction_Download UpdateAction = 1
)

// Enum value maps for UpdateAction.
var (
	UpdateAction_name = map[int32]string{
		0: "Check",
		1: "Download",
	}
	UpdateAction_value = map[string]int32{
		"Check":    0,
		"Download": 1,
	}
)

func (x UpdateAction) Enum() *UpdateAction {
	p := new(UpdateAction)
	*p = x
	return p
}

func (x UpdateAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateAction) Descriptor() protoreflect.EnumDescriptor {
	return file_libcore_proto_enumTypes[1].Descriptor()
}

func (UpdateAction) Type() protoreflect.EnumType {
	return &file_libcore_proto_enumTypes[1]
}

func (x UpdateAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateAction.Descriptor instead.
func (UpdateAction) EnumDescriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{1}
}

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_libcore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{0}
}

type EmptyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResp) Reset() {
	*x = EmptyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResp) ProtoMessage() {}

func (x *EmptyResp) ProtoReflect() protoreflect.Message {
	mi := &file_libcore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResp.ProtoReflect.Descriptor instead.
func (*EmptyResp) Descriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{1}
}

type ErrorResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorResp) Reset() {
	*x = ErrorResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResp) ProtoMessage() {}

func (x *ErrorResp) ProtoReflect() protoreflect.Message {
	mi := &file_libcore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResp.ProtoReflect.Descriptor instead.
func (*ErrorResp) Descriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type LoadConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoreConfig string `protobuf:"bytes,1,opt,name=coreConfig,proto3" json:"coreConfig,omitempty"`
	TryDomains string `protobuf:"bytes,2,opt,name=tryDomains,proto3" json:"tryDomains,omitempty"`
}

func (x *LoadConfigReq) Reset() {
	*x = LoadConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadConfigReq) ProtoMessage() {}

func (x *LoadConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_libcore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadConfigReq.ProtoReflect.Descriptor instead.
func (*LoadConfigReq) Descriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{3}
}

func (x *LoadConfigReq) GetCoreConfig() string {
	if x != nil {
		return x.CoreConfig
	}
	return ""
}

func (x *LoadConfigReq) GetTryDomains() string {
	if x != nil {
		return x.TryDomains
	}
	return ""
}

type TestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode    TestMode `protobuf:"varint,1,opt,name=mode,proto3,enum=libcore.TestMode" json:"mode,omitempty"`
	Timeout int32    `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// TcpPing
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// UrlTest
	Config  *LoadConfigReq `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	Inbound string         `protobuf:"bytes,4,opt,name=inbound,proto3" json:"inbound,omitempty"`
	Url     string         `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	// FullTest
	InAddress   string `protobuf:"bytes,7,opt,name=in_address,json=inAddress,proto3" json:"in_address,omitempty"`
	FullLatency bool   `protobuf:"varint,8,opt,name=full_latency,json=fullLatency,proto3" json:"full_latency,omitempty"`
	FullSpeed   bool   `protobuf:"varint,9,opt,name=full_speed,json=fullSpeed,proto3" json:"full_speed,omitempty"`
	FullInOut   bool   `protobuf:"varint,10,opt,name=full_in_out,json=fullInOut,proto3" json:"full_in_out,omitempty"`
	FullNat     bool   `protobuf:"varint,11,opt,name=full_nat,json=fullNat,proto3" json:"full_nat,omitempty"`
}

func (x *TestReq) Reset() {
	*x = TestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestReq) ProtoMessage() {}

func (x *TestReq) ProtoReflect() protoreflect.Message {
	mi := &file_libcore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestReq.ProtoReflect.Descriptor instead.
func (*TestReq) Descriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{4}
}

func (x *TestReq) GetMode() TestMode {
	if x != nil {
		return x.Mode
	}
	return TestMode_TcpPing
}

func (x *TestReq) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *TestReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TestReq) GetConfig() *LoadConfigReq {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *TestReq) GetInbound() string {
	if x != nil {
		return x.Inbound
	}
	return ""
}

func (x *TestReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *TestReq) GetInAddress() string {
	if x != nil {
		return x.InAddress
	}
	return ""
}

func (x *TestReq) GetFullLatency() bool {
	if x != nil {
		return x.FullLatency
	}
	return false
}

func (x *TestReq) GetFullSpeed() bool {
	if x != nil {
		return x.FullSpeed
	}
	return false
}

func (x *TestReq) GetFullInOut() bool {
	if x != nil {
		return x.FullInOut
	}
	return false
}

func (x *TestReq) GetFullNat() bool {
	if x != nil {
		return x.FullNat
	}
	return false
}

type TestResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error      string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Ms         int32  `protobuf:"varint,2,opt,name=ms,proto3" json:"ms,omitempty"`
	FullReport string `protobuf:"bytes,3,opt,name=full_report,json=fullReport,proto3" json:"full_report,omitempty"`
}

func (x *TestResp) Reset() {
	*x = TestResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResp) ProtoMessage() {}

func (x *TestResp) ProtoReflect() protoreflect.Message {
	mi := &file_libcore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResp.ProtoReflect.Descriptor instead.
func (*TestResp) Descriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{5}
}

func (x *TestResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *TestResp) GetMs() int32 {
	if x != nil {
		return x.Ms
	}
	return 0
}

func (x *TestResp) GetFullReport() string {
	if x != nil {
		return x.FullReport
	}
	return ""
}

type QueryStatsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag    string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Direct string `protobuf:"bytes,2,opt,name=direct,proto3" json:"direct,omitempty"`
}

func (x *QueryStatsReq) Reset() {
	*x = QueryStatsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryStatsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryStatsReq) ProtoMessage() {}

func (x *QueryStatsReq) ProtoReflect() protoreflect.Message {
	mi := &file_libcore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryStatsReq.ProtoReflect.Descriptor instead.
func (*QueryStatsReq) Descriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{6}
}

func (x *QueryStatsReq) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *QueryStatsReq) GetDirect() string {
	if x != nil {
		return x.Direct
	}
	return ""
}

type QueryStatsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Traffic int64 `protobuf:"varint,1,opt,name=traffic,proto3" json:"traffic,omitempty"`
}

func (x *QueryStatsResp) Reset() {
	*x = QueryStatsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryStatsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryStatsResp) ProtoMessage() {}

func (x *QueryStatsResp) ProtoReflect() protoreflect.Message {
	mi := &file_libcore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryStatsResp.ProtoReflect.Descriptor instead.
func (*QueryStatsResp) Descriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{7}
}

func (x *QueryStatsResp) GetTraffic() int64 {
	if x != nil {
		return x.Traffic
	}
	return 0
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action          UpdateAction `protobuf:"varint,1,opt,name=action,proto3,enum=libcore.UpdateAction" json:"action,omitempty"`
	CheckPreRelease bool         `protobuf:"varint,2,opt,name=check_pre_release,json=checkPreRelease,proto3" json:"check_pre_release,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_libcore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateReq) GetAction() UpdateAction {
	if x != nil {
		return x.Action
	}
	return UpdateAction_Check
}

func (x *UpdateReq) GetCheckPreRelease() bool {
	if x != nil {
		return x.CheckPreRelease
	}
	return false
}

type UpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error        string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	AssetsName   string `protobuf:"bytes,2,opt,name=assets_name,json=assetsName,proto3" json:"assets_name,omitempty"`
	DownloadUrl  string `protobuf:"bytes,3,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
	ReleaseUrl   string `protobuf:"bytes,4,opt,name=release_url,json=releaseUrl,proto3" json:"release_url,omitempty"`
	ReleaseNote  string `protobuf:"bytes,5,opt,name=release_note,json=releaseNote,proto3" json:"release_note,omitempty"`
	IsPreRelease bool   `protobuf:"varint,6,opt,name=is_pre_release,json=isPreRelease,proto3" json:"is_pre_release,omitempty"`
}

func (x *UpdateResp) Reset() {
	*x = UpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcore_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResp) ProtoMessage() {}

func (x *UpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_libcore_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResp.ProtoReflect.Descriptor instead.
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *UpdateResp) GetAssetsName() string {
	if x != nil {
		return x.AssetsName
	}
	return ""
}

func (x *UpdateResp) GetDownloadUrl() string {
	if x != nil {
		return x.DownloadUrl
	}
	return ""
}

func (x *UpdateResp) GetReleaseUrl() string {
	if x != nil {
		return x.ReleaseUrl
	}
	return ""
}

func (x *UpdateResp) GetReleaseNote() string {
	if x != nil {
		return x.ReleaseNote
	}
	return ""
}

func (x *UpdateResp) GetIsPreRelease() bool {
	if x != nil {
		return x.IsPreRelease
	}
	return false
}

type ListConnectionsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatsuriConnectionsJson string `protobuf:"bytes,1,opt,name=matsuri_connections_json,json=matsuriConnectionsJson,proto3" json:"matsuri_connections_json,omitempty"`
}

func (x *ListConnectionsResp) Reset() {
	*x = ListConnectionsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcore_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConnectionsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConnectionsResp) ProtoMessage() {}

func (x *ListConnectionsResp) ProtoReflect() protoreflect.Message {
	mi := &file_libcore_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConnectionsResp.ProtoReflect.Descriptor instead.
func (*ListConnectionsResp) Descriptor() ([]byte, []int) {
	return file_libcore_proto_rawDescGZIP(), []int{10}
}

func (x *ListConnectionsResp) GetMatsuriConnectionsJson() string {
	if x != nil {
		return x.MatsuriConnectionsJson
	}
	return ""
}

var File_libcore_proto protoreflect.FileDescriptor

var file_libcore_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x22, 0x0b, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x21, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x4f, 0x0a, 0x0d, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x79, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x79, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x73, 0x22, 0xdc, 0x02, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x25, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c,
	0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x66,
	0x75, 0x6c, 0x6c, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x66, 0x75, 0x6c, 0x6c, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x69, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x66, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x6e, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x75, 0x6c,
	0x6c, 0x4e, 0x61, 0x74, 0x22, 0x51, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x75, 0x6c,
	0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x39, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x22, 0x66,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6c, 0x69,
	0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x65, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x5f, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x50,
	0x72, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0x4f, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x38, 0x0a, 0x18, 0x6d, 0x61, 0x74, 0x73, 0x75, 0x72, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x16, 0x6d, 0x61, 0x74, 0x73, 0x75, 0x72, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x73, 0x6f, 0x6e, 0x2a, 0x32, 0x0a, 0x08, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e,
	0x67, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x72, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x10, 0x02, 0x2a, 0x27,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09,
	0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x10, 0x01, 0x32, 0xca, 0x03, 0x0a, 0x0e, 0x4c, 0x69, 0x62, 0x63,
	0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x45, 0x78,
	0x69, 0x74, 0x12, 0x11, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x4b,
	0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x11, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6c, 0x69,
	0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x6c, 0x69,
	0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x13, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x16, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2f, 0x0a,
	0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x11, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2d,
	0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x6c, 0x69,
	0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x11, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x6e, 0x65, 0x6b, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_libcore_proto_rawDescOnce sync.Once
	file_libcore_proto_rawDescData = file_libcore_proto_rawDesc
)

func file_libcore_proto_rawDescGZIP() []byte {
	file_libcore_proto_rawDescOnce.Do(func() {
		file_libcore_proto_rawDescData = protoimpl.X.CompressGZIP(file_libcore_proto_rawDescData)
	})
	return file_libcore_proto_rawDescData
}

var file_libcore_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_libcore_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_libcore_proto_goTypes = []interface{}{
	(TestMode)(0),               // 0: libcore.TestMode
	(UpdateAction)(0),           // 1: libcore.UpdateAction
	(*EmptyReq)(nil),            // 2: libcore.EmptyReq
	(*EmptyResp)(nil),           // 3: libcore.EmptyResp
	(*ErrorResp)(nil),           // 4: libcore.ErrorResp
	(*LoadConfigReq)(nil),       // 5: libcore.LoadConfigReq
	(*TestReq)(nil),             // 6: libcore.TestReq
	(*TestResp)(nil),            // 7: libcore.TestResp
	(*QueryStatsReq)(nil),       // 8: libcore.QueryStatsReq
	(*QueryStatsResp)(nil),      // 9: libcore.QueryStatsResp
	(*UpdateReq)(nil),           // 10: libcore.UpdateReq
	(*UpdateResp)(nil),          // 11: libcore.UpdateResp
	(*ListConnectionsResp)(nil), // 12: libcore.ListConnectionsResp
}
var file_libcore_proto_depIdxs = []int32{
	0,  // 0: libcore.TestReq.mode:type_name -> libcore.TestMode
	5,  // 1: libcore.TestReq.config:type_name -> libcore.LoadConfigReq
	1,  // 2: libcore.UpdateReq.action:type_name -> libcore.UpdateAction
	2,  // 3: libcore.LibcoreService.Exit:input_type -> libcore.EmptyReq
	2,  // 4: libcore.LibcoreService.KeepAlive:input_type -> libcore.EmptyReq
	10, // 5: libcore.LibcoreService.Update:input_type -> libcore.UpdateReq
	5,  // 6: libcore.LibcoreService.Start:input_type -> libcore.LoadConfigReq
	2,  // 7: libcore.LibcoreService.Stop:input_type -> libcore.EmptyReq
	6,  // 8: libcore.LibcoreService.Test:input_type -> libcore.TestReq
	8,  // 9: libcore.LibcoreService.QueryStats:input_type -> libcore.QueryStatsReq
	2,  // 10: libcore.LibcoreService.ListConnections:input_type -> libcore.EmptyReq
	3,  // 11: libcore.LibcoreService.Exit:output_type -> libcore.EmptyResp
	3,  // 12: libcore.LibcoreService.KeepAlive:output_type -> libcore.EmptyResp
	11, // 13: libcore.LibcoreService.Update:output_type -> libcore.UpdateResp
	4,  // 14: libcore.LibcoreService.Start:output_type -> libcore.ErrorResp
	4,  // 15: libcore.LibcoreService.Stop:output_type -> libcore.ErrorResp
	7,  // 16: libcore.LibcoreService.Test:output_type -> libcore.TestResp
	9,  // 17: libcore.LibcoreService.QueryStats:output_type -> libcore.QueryStatsResp
	12, // 18: libcore.LibcoreService.ListConnections:output_type -> libcore.ListConnectionsResp
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_libcore_proto_init() }
func file_libcore_proto_init() {
	if File_libcore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_libcore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
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
		file_libcore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResp); i {
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
		file_libcore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResp); i {
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
		file_libcore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadConfigReq); i {
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
		file_libcore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestReq); i {
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
		file_libcore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResp); i {
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
		file_libcore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryStatsReq); i {
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
		file_libcore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryStatsResp); i {
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
		file_libcore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_libcore_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResp); i {
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
		file_libcore_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConnectionsResp); i {
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
			RawDescriptor: file_libcore_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_libcore_proto_goTypes,
		DependencyIndexes: file_libcore_proto_depIdxs,
		EnumInfos:         file_libcore_proto_enumTypes,
		MessageInfos:      file_libcore_proto_msgTypes,
	}.Build()
	File_libcore_proto = out.File
	file_libcore_proto_rawDesc = nil
	file_libcore_proto_goTypes = nil
	file_libcore_proto_depIdxs = nil
}
