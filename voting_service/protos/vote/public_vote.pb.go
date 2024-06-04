// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: public_vote.proto

package vote

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

type PublicVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ElectionId string `protobuf:"bytes,2,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,3,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PublicVote) Reset() {
	*x = PublicVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_vote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicVote) ProtoMessage() {}

func (x *PublicVote) ProtoReflect() protoreflect.Message {
	mi := &file_public_vote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicVote.ProtoReflect.Descriptor instead.
func (*PublicVote) Descriptor() ([]byte, []int) {
	return file_public_vote_proto_rawDescGZIP(), []int{0}
}

func (x *PublicVote) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublicVote) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *PublicVote) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *PublicVote) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_public_vote_proto protoreflect.FileDescriptor

var file_public_vote_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x79, 0x0a, 0x0a, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x3b,
	0x76, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_vote_proto_rawDescOnce sync.Once
	file_public_vote_proto_rawDescData = file_public_vote_proto_rawDesc
)

func file_public_vote_proto_rawDescGZIP() []byte {
	file_public_vote_proto_rawDescOnce.Do(func() {
		file_public_vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_vote_proto_rawDescData)
	})
	return file_public_vote_proto_rawDescData
}

var file_public_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_public_vote_proto_goTypes = []interface{}{
	(*PublicVote)(nil), // 0: protos.PublicVote
}
var file_public_vote_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_public_vote_proto_init() }
func file_public_vote_proto_init() {
	if File_public_vote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_vote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicVote); i {
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
			RawDescriptor: file_public_vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_public_vote_proto_goTypes,
		DependencyIndexes: file_public_vote_proto_depIdxs,
		MessageInfos:      file_public_vote_proto_msgTypes,
	}.Build()
	File_public_vote_proto = out.File
	file_public_vote_proto_rawDesc = nil
	file_public_vote_proto_goTypes = nil
	file_public_vote_proto_depIdxs = nil
}
