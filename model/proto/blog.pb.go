// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: blog.proto

package proto

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

type PostID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID int32 `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (x *PostID) Reset() {
	*x = PostID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostID) ProtoMessage() {}

func (x *PostID) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostID.ProtoReflect.Descriptor instead.
func (*PostID) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{0}
}

func (x *PostID) GetPostID() int32 {
	if x != nil {
		return x.PostID
	}
	return 0
}

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID          int32  `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
	Content         string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Author          string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	PublicationDate string `protobuf:"bytes,4,opt,name=publicationDate,proto3" json:"publicationDate,omitempty"`
	Tags            string `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{1}
}

func (x *Blog) GetPostID() int32 {
	if x != nil {
		return x.PostID
	}
	return 0
}

func (x *Blog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Blog) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Blog) GetPublicationDate() string {
	if x != nil {
		return x.PublicationDate
	}
	return ""
}

func (x *Blog) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

var File_blog_proto protoreflect.FileDescriptor

var file_blog_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x20, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x44, 0x22, 0x8e, 0x01, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x32, 0xa8, 0x01, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x1a, 0x0b,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x24, 0x0a,
	0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x6c, 0x6f,
	0x67, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x06, 0x55, 0x50, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x1a, 0x0b, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22,
	0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_blog_proto_rawDescOnce sync.Once
	file_blog_proto_rawDescData = file_blog_proto_rawDesc
)

func file_blog_proto_rawDescGZIP() []byte {
	file_blog_proto_rawDescOnce.Do(func() {
		file_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_proto_rawDescData)
	})
	return file_blog_proto_rawDescData
}

var file_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_blog_proto_goTypes = []interface{}{
	(*PostID)(nil), // 0: model.PostID
	(*Blog)(nil),   // 1: model.Blog
}
var file_blog_proto_depIdxs = []int32{
	1, // 0: model.HelloService.Create:input_type -> model.Blog
	0, // 1: model.HelloService.Read:input_type -> model.PostID
	1, // 2: model.HelloService.UPdate:input_type -> model.Blog
	0, // 3: model.HelloService.Delete:input_type -> model.PostID
	1, // 4: model.HelloService.Create:output_type -> model.Blog
	1, // 5: model.HelloService.Read:output_type -> model.Blog
	1, // 6: model.HelloService.UPdate:output_type -> model.Blog
	1, // 7: model.HelloService.Delete:output_type -> model.Blog
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blog_proto_init() }
func file_blog_proto_init() {
	if File_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostID); i {
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
		file_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
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
			RawDescriptor: file_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_proto_goTypes,
		DependencyIndexes: file_blog_proto_depIdxs,
		MessageInfos:      file_blog_proto_msgTypes,
	}.Build()
	File_blog_proto = out.File
	file_blog_proto_rawDesc = nil
	file_blog_proto_goTypes = nil
	file_blog_proto_depIdxs = nil
}
