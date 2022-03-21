// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/passage/passage.proto

package passage

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PassageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book    string `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	Chapter int32  `protobuf:"varint,2,opt,name=chapter,proto3" json:"chapter,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *PassageRequest) Reset() {
	*x = PassageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_passage_passage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PassageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassageRequest) ProtoMessage() {}

func (x *PassageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_passage_passage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassageRequest.ProtoReflect.Descriptor instead.
func (*PassageRequest) Descriptor() ([]byte, []int) {
	return file_proto_passage_passage_proto_rawDescGZIP(), []int{0}
}

func (x *PassageRequest) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

func (x *PassageRequest) GetChapter() int32 {
	if x != nil {
		return x.Chapter
	}
	return 0
}

func (x *PassageRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Verse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verse   int32  `protobuf:"varint,1,opt,name=verse,proto3" json:"verse,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Verse) Reset() {
	*x = Verse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_passage_passage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Verse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Verse) ProtoMessage() {}

func (x *Verse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_passage_passage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Verse.ProtoReflect.Descriptor instead.
func (*Verse) Descriptor() ([]byte, []int) {
	return file_proto_passage_passage_proto_rawDescGZIP(), []int{1}
}

func (x *Verse) GetVerse() int32 {
	if x != nil {
		return x.Verse
	}
	return 0
}

func (x *Verse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Verse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PassageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verse []*Verse `protobuf:"bytes,1,rep,name=verse,proto3" json:"verse,omitempty"`
}

func (x *PassageResponse) Reset() {
	*x = PassageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_passage_passage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PassageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassageResponse) ProtoMessage() {}

func (x *PassageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_passage_passage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassageResponse.ProtoReflect.Descriptor instead.
func (*PassageResponse) Descriptor() ([]byte, []int) {
	return file_proto_passage_passage_proto_rawDescGZIP(), []int{2}
}

func (x *PassageResponse) GetVerse() []*Verse {
	if x != nil {
		return x.Verse
	}
	return nil
}

var File_proto_passage_passage_proto protoreflect.FileDescriptor

var file_proto_passage_passage_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x70, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0e, 0x50,
	0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x6f,
	0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x05, 0x56, 0x65, 0x72, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x2f, 0x0a, 0x0f, 0x50, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x65, 0x52, 0x05, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x32, 0x6e, 0x0a, 0x0e, 0x50, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x07, 0x50, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0f, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x7b, 0x62, 0x6f, 0x6f, 0x6b, 0x7d, 0x2f,
	0x7b, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x7d, 0x2f, 0x7b, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x7d, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_passage_passage_proto_rawDescOnce sync.Once
	file_proto_passage_passage_proto_rawDescData = file_proto_passage_passage_proto_rawDesc
)

func file_proto_passage_passage_proto_rawDescGZIP() []byte {
	file_proto_passage_passage_proto_rawDescOnce.Do(func() {
		file_proto_passage_passage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_passage_passage_proto_rawDescData)
	})
	return file_proto_passage_passage_proto_rawDescData
}

var file_proto_passage_passage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_passage_passage_proto_goTypes = []interface{}{
	(*PassageRequest)(nil),  // 0: PassageRequest
	(*Verse)(nil),           // 1: Verse
	(*PassageResponse)(nil), // 2: PassageResponse
}
var file_proto_passage_passage_proto_depIdxs = []int32{
	1, // 0: PassageResponse.verse:type_name -> Verse
	0, // 1: PassageService.Passage:input_type -> PassageRequest
	2, // 2: PassageService.Passage:output_type -> PassageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_passage_passage_proto_init() }
func file_proto_passage_passage_proto_init() {
	if File_proto_passage_passage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_passage_passage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PassageRequest); i {
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
		file_proto_passage_passage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Verse); i {
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
		file_proto_passage_passage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PassageResponse); i {
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
			RawDescriptor: file_proto_passage_passage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_passage_passage_proto_goTypes,
		DependencyIndexes: file_proto_passage_passage_proto_depIdxs,
		MessageInfos:      file_proto_passage_passage_proto_msgTypes,
	}.Build()
	File_proto_passage_passage_proto = out.File
	file_proto_passage_passage_proto_rawDesc = nil
	file_proto_passage_passage_proto_goTypes = nil
	file_proto_passage_passage_proto_depIdxs = nil
}