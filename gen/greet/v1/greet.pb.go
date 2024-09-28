// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: greet/v1/greet.proto

package greetv1

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

type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_v1_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_v1_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_greet_v1_greet_proto_rawDescGZIP(), []int{0}
}

func (x *GreetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_v1_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_v1_greet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_greet_v1_greet_proto_rawDescGZIP(), []int{1}
}

func (x *GreetResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

type MulRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputNum int32 `protobuf:"varint,1,opt,name=input_num,json=inputNum,proto3" json:"input_num,omitempty"`
}

func (x *MulRequest) Reset() {
	*x = MulRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_v1_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MulRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MulRequest) ProtoMessage() {}

func (x *MulRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_v1_greet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MulRequest.ProtoReflect.Descriptor instead.
func (*MulRequest) Descriptor() ([]byte, []int) {
	return file_greet_v1_greet_proto_rawDescGZIP(), []int{2}
}

func (x *MulRequest) GetInputNum() int32 {
	if x != nil {
		return x.InputNum
	}
	return 0
}

type MulResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MulResponse) Reset() {
	*x = MulResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_v1_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MulResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MulResponse) ProtoMessage() {}

func (x *MulResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_v1_greet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MulResponse.ProtoReflect.Descriptor instead.
func (*MulResponse) Descriptor() ([]byte, []int) {
	return file_greet_v1_greet_proto_rawDescGZIP(), []int{3}
}

func (x *MulResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_greet_v1_greet_proto protoreflect.FileDescriptor

var file_greet_v1_greet_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x22, 0x22, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x29, 0x0a, 0x0a, 0x4d, 0x75, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4e, 0x75, 0x6d, 0x22, 0x25, 0x0a, 0x0b,
	0x4d, 0x75, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0x80, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x03, 0x4d, 0x75, 0x6c, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x74, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greet_v1_greet_proto_rawDescOnce sync.Once
	file_greet_v1_greet_proto_rawDescData = file_greet_v1_greet_proto_rawDesc
)

func file_greet_v1_greet_proto_rawDescGZIP() []byte {
	file_greet_v1_greet_proto_rawDescOnce.Do(func() {
		file_greet_v1_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_greet_v1_greet_proto_rawDescData)
	})
	return file_greet_v1_greet_proto_rawDescData
}

var file_greet_v1_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_greet_v1_greet_proto_goTypes = []any{
	(*GreetRequest)(nil),  // 0: greet.v1.GreetRequest
	(*GreetResponse)(nil), // 1: greet.v1.GreetResponse
	(*MulRequest)(nil),    // 2: greet.v1.MulRequest
	(*MulResponse)(nil),   // 3: greet.v1.MulResponse
}
var file_greet_v1_greet_proto_depIdxs = []int32{
	0, // 0: greet.v1.GreetService.Greet:input_type -> greet.v1.GreetRequest
	2, // 1: greet.v1.GreetService.Mul:input_type -> greet.v1.MulRequest
	1, // 2: greet.v1.GreetService.Greet:output_type -> greet.v1.GreetResponse
	3, // 3: greet.v1.GreetService.Mul:output_type -> greet.v1.MulResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greet_v1_greet_proto_init() }
func file_greet_v1_greet_proto_init() {
	if File_greet_v1_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greet_v1_greet_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GreetRequest); i {
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
		file_greet_v1_greet_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GreetResponse); i {
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
		file_greet_v1_greet_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MulRequest); i {
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
		file_greet_v1_greet_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MulResponse); i {
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
			RawDescriptor: file_greet_v1_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greet_v1_greet_proto_goTypes,
		DependencyIndexes: file_greet_v1_greet_proto_depIdxs,
		MessageInfos:      file_greet_v1_greet_proto_msgTypes,
	}.Build()
	File_greet_v1_greet_proto = out.File
	file_greet_v1_greet_proto_rawDesc = nil
	file_greet_v1_greet_proto_goTypes = nil
	file_greet_v1_greet_proto_depIdxs = nil
}
