// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.1
// source: print.proto

package print

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

type PrintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PrintRequest) Reset() {
	*x = PrintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_print_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintRequest) ProtoMessage() {}

func (x *PrintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_print_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintRequest.ProtoReflect.Descriptor instead.
func (*PrintRequest) Descriptor() ([]byte, []int) {
	return file_print_proto_rawDescGZIP(), []int{0}
}

func (x *PrintRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PrintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PrintResponse) Reset() {
	*x = PrintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_print_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintResponse) ProtoMessage() {}

func (x *PrintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_print_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintResponse.ProtoReflect.Descriptor instead.
func (*PrintResponse) Descriptor() ([]byte, []int) {
	return file_print_proto_rawDescGZIP(), []int{1}
}

func (x *PrintResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_print_proto protoreflect.FileDescriptor

var file_print_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a,
	0x0c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x29, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x36, 0x0a, 0x05,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x0d, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x61, 0x72, 0x74, 0x68, 0x61, 0x59, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x5f,
	0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_print_proto_rawDescOnce sync.Once
	file_print_proto_rawDescData = file_print_proto_rawDesc
)

func file_print_proto_rawDescGZIP() []byte {
	file_print_proto_rawDescOnce.Do(func() {
		file_print_proto_rawDescData = protoimpl.X.CompressGZIP(file_print_proto_rawDescData)
	})
	return file_print_proto_rawDescData
}

var file_print_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_print_proto_goTypes = []interface{}{
	(*PrintRequest)(nil),  // 0: PrintRequest
	(*PrintResponse)(nil), // 1: PrintResponse
}
var file_print_proto_depIdxs = []int32{
	0, // 0: Print.PrintHello:input_type -> PrintRequest
	1, // 1: Print.PrintHello:output_type -> PrintResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_print_proto_init() }
func file_print_proto_init() {
	if File_print_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_print_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintRequest); i {
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
		file_print_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintResponse); i {
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
			RawDescriptor: file_print_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_print_proto_goTypes,
		DependencyIndexes: file_print_proto_depIdxs,
		MessageInfos:      file_print_proto_msgTypes,
	}.Build()
	File_print_proto = out.File
	file_print_proto_rawDesc = nil
	file_print_proto_goTypes = nil
	file_print_proto_depIdxs = nil
}
