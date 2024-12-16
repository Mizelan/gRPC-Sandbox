// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.2
// source: proto/v1/entity.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Age   int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_proto_v1_entity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_entity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_v1_entity_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

var file_proto_v1_entity_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50001,
		Name:          "fieldoption.proto.v1.db_column",
		Tag:           "bytes,50001,opt,name=db_column",
		Filename:      "proto/v1/entity.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50002,
		Name:          "fieldoption.proto.v1.primary_key",
		Tag:           "varint,50002,opt,name=primary_key",
		Filename:      "proto/v1/entity.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50003,
		Name:          "fieldoption.proto.v1.index",
		Tag:           "varint,50003,opt,name=index",
		Filename:      "proto/v1/entity.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50004,
		Name:          "fieldoption.proto.v1.db_type",
		Tag:           "bytes,50004,opt,name=db_type",
		Filename:      "proto/v1/entity.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string db_column = 50001;
	E_DbColumn = &file_proto_v1_entity_proto_extTypes[0] // 데이터베이스 컬럼 이름
	// optional bool primary_key = 50002;
	E_PrimaryKey = &file_proto_v1_entity_proto_extTypes[1] // 기본키 여부
	// optional bool index = 50003;
	E_Index = &file_proto_v1_entity_proto_extTypes[2] // 인덱스 필요 여부
	// optional string db_type = 50004;
	E_DbType = &file_proto_v1_entity_proto_extTypes[3] // 데이터베이스 타입
)

var File_proto_v1_entity_proto protoreflect.FileDescriptor

var file_proto_v1_entity_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0x8a, 0xb5, 0x18, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x90, 0xb5, 0x18, 0x01, 0xa2, 0xb5, 0x18, 0x0b, 0x56, 0x41, 0x52, 0x43, 0x48, 0x41, 0x52,
	0x28, 0x33, 0x36, 0x29, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0x8a, 0xb5, 0x18, 0x0d, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x98, 0xb5, 0x18, 0x01, 0xa2, 0xb5,
	0x18, 0x0c, 0x56, 0x41, 0x52, 0x43, 0x48, 0x41, 0x52, 0x28, 0x32, 0x35, 0x35, 0x29, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x17, 0x8a, 0xb5, 0x18, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65,
	0xa2, 0xb5, 0x18, 0x07, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x3a, 0x3f, 0x0a, 0x09, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x62, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x88, 0x01,
	0x01, 0x3a, 0x43, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd2, 0x86, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x3a, 0x38, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3,
	0x86, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01,
	0x3a, 0x3b, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd4, 0x86, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x16, 0x5a,
	0x14, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_entity_proto_rawDescOnce sync.Once
	file_proto_v1_entity_proto_rawDescData = file_proto_v1_entity_proto_rawDesc
)

func file_proto_v1_entity_proto_rawDescGZIP() []byte {
	file_proto_v1_entity_proto_rawDescOnce.Do(func() {
		file_proto_v1_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_entity_proto_rawDescData)
	})
	return file_proto_v1_entity_proto_rawDescData
}

var file_proto_v1_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_v1_entity_proto_goTypes = []any{
	(*User)(nil),                      // 0: fieldoption.proto.v1.User
	(*descriptorpb.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_proto_v1_entity_proto_depIdxs = []int32{
	1, // 0: fieldoption.proto.v1.db_column:extendee -> google.protobuf.FieldOptions
	1, // 1: fieldoption.proto.v1.primary_key:extendee -> google.protobuf.FieldOptions
	1, // 2: fieldoption.proto.v1.index:extendee -> google.protobuf.FieldOptions
	1, // 3: fieldoption.proto.v1.db_type:extendee -> google.protobuf.FieldOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1_entity_proto_init() }
func file_proto_v1_entity_proto_init() {
	if File_proto_v1_entity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_proto_v1_entity_proto_goTypes,
		DependencyIndexes: file_proto_v1_entity_proto_depIdxs,
		MessageInfos:      file_proto_v1_entity_proto_msgTypes,
		ExtensionInfos:    file_proto_v1_entity_proto_extTypes,
	}.Build()
	File_proto_v1_entity_proto = out.File
	file_proto_v1_entity_proto_rawDesc = nil
	file_proto_v1_entity_proto_goTypes = nil
	file_proto_v1_entity_proto_depIdxs = nil
}
