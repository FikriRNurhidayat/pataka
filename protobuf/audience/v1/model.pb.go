// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: audience/v1/model.proto

package audiencev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AudienceStatus int32

const (
	AudienceStatus_AUDIENCE_STATUS_UNSPECIFIED AudienceStatus = 0
	AudienceStatus_AUDIENCE_STATUS_ENABLED     AudienceStatus = 1
	AudienceStatus_AUDIENCE_STATUS_DISABLED    AudienceStatus = 2
)

// Enum value maps for AudienceStatus.
var (
	AudienceStatus_name = map[int32]string{
		0: "AUDIENCE_STATUS_UNSPECIFIED",
		1: "AUDIENCE_STATUS_ENABLED",
		2: "AUDIENCE_STATUS_DISABLED",
	}
	AudienceStatus_value = map[string]int32{
		"AUDIENCE_STATUS_UNSPECIFIED": 0,
		"AUDIENCE_STATUS_ENABLED":     1,
		"AUDIENCE_STATUS_DISABLED":    2,
	}
)

func (x AudienceStatus) Enum() *AudienceStatus {
	p := new(AudienceStatus)
	*p = x
	return p
}

func (x AudienceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AudienceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_audience_v1_model_proto_enumTypes[0].Descriptor()
}

func (AudienceStatus) Type() protoreflect.EnumType {
	return &file_audience_v1_model_proto_enumTypes[0]
}

func (x AudienceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AudienceStatus.Descriptor instead.
func (AudienceStatus) EnumDescriptor() ([]byte, []int) {
	return file_audience_v1_model_proto_rawDescGZIP(), []int{0}
}

type Audience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudienceId  string                 `protobuf:"bytes,1,opt,name=audience_id,json=audienceId,proto3" json:"audience_id,omitempty"`
	FeatureName string                 `protobuf:"bytes,2,opt,name=feature_name,json=featureName,proto3" json:"feature_name,omitempty"`
	Enabled     bool                   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	EnabledAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=enabled_at,json=enabledAt,proto3" json:"enabled_at,omitempty"`
}

func (x *Audience) Reset() {
	*x = Audience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audience_v1_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audience) ProtoMessage() {}

func (x *Audience) ProtoReflect() protoreflect.Message {
	mi := &file_audience_v1_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audience.ProtoReflect.Descriptor instead.
func (*Audience) Descriptor() ([]byte, []int) {
	return file_audience_v1_model_proto_rawDescGZIP(), []int{0}
}

func (x *Audience) GetAudienceId() string {
	if x != nil {
		return x.AudienceId
	}
	return ""
}

func (x *Audience) GetFeatureName() string {
	if x != nil {
		return x.FeatureName
	}
	return ""
}

func (x *Audience) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Audience) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Audience) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Audience) GetEnabledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EnabledAt
	}
	return nil
}

type BulkCreateAudience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudienceId  string `protobuf:"bytes,1,opt,name=audience_id,json=audienceId,proto3" json:"audience_id,omitempty"`
	FeatureName string `protobuf:"bytes,2,opt,name=feature_name,json=featureName,proto3" json:"feature_name,omitempty"`
	Enabled     bool   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *BulkCreateAudience) Reset() {
	*x = BulkCreateAudience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audience_v1_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkCreateAudience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkCreateAudience) ProtoMessage() {}

func (x *BulkCreateAudience) ProtoReflect() protoreflect.Message {
	mi := &file_audience_v1_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkCreateAudience.ProtoReflect.Descriptor instead.
func (*BulkCreateAudience) Descriptor() ([]byte, []int) {
	return file_audience_v1_model_proto_rawDescGZIP(), []int{1}
}

func (x *BulkCreateAudience) GetAudienceId() string {
	if x != nil {
		return x.AudienceId
	}
	return ""
}

func (x *BulkCreateAudience) GetFeatureName() string {
	if x != nil {
		return x.FeatureName
	}
	return ""
}

func (x *BulkCreateAudience) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_audience_v1_model_proto protoreflect.FileDescriptor

var file_audience_v1_model_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x72, 0x0a, 0x12, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x64,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x2a, 0x6c, 0x0a, 0x0e, 0x41, 0x75, 0x64, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x55, 0x44,
	0x49, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x55,
	0x44, 0x49, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x4e,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x55, 0x44, 0x49, 0x45,
	0x4e, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0xac, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6b, 0x72, 0x69, 0x72, 0x6e, 0x75, 0x72, 0x68, 0x69, 0x64,
	0x61, 0x79, 0x61, 0x74, 0x2f, 0x66, 0x66, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa,
	0x02, 0x0b, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b,
	0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_audience_v1_model_proto_rawDescOnce sync.Once
	file_audience_v1_model_proto_rawDescData = file_audience_v1_model_proto_rawDesc
)

func file_audience_v1_model_proto_rawDescGZIP() []byte {
	file_audience_v1_model_proto_rawDescOnce.Do(func() {
		file_audience_v1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_audience_v1_model_proto_rawDescData)
	})
	return file_audience_v1_model_proto_rawDescData
}

var file_audience_v1_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_audience_v1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_audience_v1_model_proto_goTypes = []interface{}{
	(AudienceStatus)(0),           // 0: audience.v1.AudienceStatus
	(*Audience)(nil),              // 1: audience.v1.Audience
	(*BulkCreateAudience)(nil),    // 2: audience.v1.BulkCreateAudience
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_audience_v1_model_proto_depIdxs = []int32{
	3, // 0: audience.v1.Audience.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: audience.v1.Audience.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: audience.v1.Audience.enabled_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_audience_v1_model_proto_init() }
func file_audience_v1_model_proto_init() {
	if File_audience_v1_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_audience_v1_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audience); i {
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
		file_audience_v1_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkCreateAudience); i {
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
			RawDescriptor: file_audience_v1_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_audience_v1_model_proto_goTypes,
		DependencyIndexes: file_audience_v1_model_proto_depIdxs,
		EnumInfos:         file_audience_v1_model_proto_enumTypes,
		MessageInfos:      file_audience_v1_model_proto_msgTypes,
	}.Build()
	File_audience_v1_model_proto = out.File
	file_audience_v1_model_proto_rawDesc = nil
	file_audience_v1_model_proto_goTypes = nil
	file_audience_v1_model_proto_depIdxs = nil
}
