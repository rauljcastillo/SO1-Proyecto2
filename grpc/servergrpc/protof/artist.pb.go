// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: artist.proto

package protosf

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

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Album string `protobuf:"bytes,2,opt,name=album,proto3" json:"album,omitempty"`
	Year  string `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	Rank  string `protobuf:"bytes,4,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_artist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_artist_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Req) GetAlbum() string {
	if x != nil {
		return x.Album
	}
	return ""
}

func (x *Req) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Req) GetRank() string {
	if x != nil {
		return x.Rank
	}
	return ""
}

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_artist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_artist_proto_rawDescGZIP(), []int{1}
}

var File_artist_proto protoreflect.FileDescriptor

var file_artist_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x63, 0x6f, 0x6e, 0x66, 0x70, 0x22, 0x57, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61,
	0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x05,
	0x0a, 0x03, 0x52, 0x65, 0x73, 0x32, 0x2b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x70,
	0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x70, 0x2e, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artist_proto_rawDescOnce sync.Once
	file_artist_proto_rawDescData = file_artist_proto_rawDesc
)

func file_artist_proto_rawDescGZIP() []byte {
	file_artist_proto_rawDescOnce.Do(func() {
		file_artist_proto_rawDescData = protoimpl.X.CompressGZIP(file_artist_proto_rawDescData)
	})
	return file_artist_proto_rawDescData
}

var file_artist_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_artist_proto_goTypes = []interface{}{
	(*Req)(nil), // 0: confp.Req
	(*Res)(nil), // 1: confp.Res
}
var file_artist_proto_depIdxs = []int32{
	0, // 0: confp.Info.GetInfo:input_type -> confp.Req
	1, // 1: confp.Info.GetInfo:output_type -> confp.Res
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_artist_proto_init() }
func file_artist_proto_init() {
	if File_artist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_artist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
			RawDescriptor: file_artist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_artist_proto_goTypes,
		DependencyIndexes: file_artist_proto_depIdxs,
		MessageInfos:      file_artist_proto_msgTypes,
	}.Build()
	File_artist_proto = out.File
	file_artist_proto_rawDesc = nil
	file_artist_proto_goTypes = nil
	file_artist_proto_depIdxs = nil
}
