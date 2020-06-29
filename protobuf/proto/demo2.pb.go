// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: protobuf/proto/demo2.proto

package example

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

type Person_Gender int32

const (
	Person_MALE    Person_Gender = 0
	Person_FEMALE  Person_Gender = 1
	Person_UNKNOWN Person_Gender = 2
)

// Enum value maps for Person_Gender.
var (
	Person_Gender_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
		2: "UNKNOWN",
	}
	Person_Gender_value = map[string]int32{
		"MALE":    0,
		"FEMALE":  1,
		"UNKNOWN": 2,
	}
)

func (x Person_Gender) Enum() *Person_Gender {
	p := new(Person_Gender)
	*p = x
	return p
}

func (x Person_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Person_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_proto_demo2_proto_enumTypes[0].Descriptor()
}

func (Person_Gender) Type() protoreflect.EnumType {
	return &file_protobuf_proto_demo2_proto_enumTypes[0]
}

func (x Person_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Person_Gender.Descriptor instead.
func (Person_Gender) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_proto_demo2_proto_rawDescGZIP(), []int{0, 0}
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32         `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Info *Person_Other `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_proto_demo2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_proto_demo2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_protobuf_proto_demo2_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetInfo() *Person_Other {
	if x != nil {
		return x.Info
	}
	return nil
}

type Person_Other struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr  string        `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Hobby string        `protobuf:"bytes,2,opt,name=hobby,proto3" json:"hobby,omitempty"`
	G     Person_Gender `protobuf:"varint,3,opt,name=g,proto3,enum=example.Person_Gender" json:"g,omitempty"`
}

func (x *Person_Other) Reset() {
	*x = Person_Other{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_proto_demo2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_Other) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Other) ProtoMessage() {}

func (x *Person_Other) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_proto_demo2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_Other.ProtoReflect.Descriptor instead.
func (*Person_Other) Descriptor() ([]byte, []int) {
	return file_protobuf_proto_demo2_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Person_Other) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Person_Other) GetHobby() string {
	if x != nil {
		return x.Hobby
	}
	return ""
}

func (x *Person_Other) GetG() Person_Gender {
	if x != nil {
		return x.G
	}
	return Person_MALE
}

var File_protobuf_proto_demo2_proto protoreflect.FileDescriptor

var file_protobuf_proto_demo2_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0xdf, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x1a, 0x57, 0x0a, 0x05, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68,
	0x6f, 0x62, 0x62, 0x79, 0x12, 0x24, 0x0a, 0x01, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x01, 0x67, 0x22, 0x2b, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_proto_demo2_proto_rawDescOnce sync.Once
	file_protobuf_proto_demo2_proto_rawDescData = file_protobuf_proto_demo2_proto_rawDesc
)

func file_protobuf_proto_demo2_proto_rawDescGZIP() []byte {
	file_protobuf_proto_demo2_proto_rawDescOnce.Do(func() {
		file_protobuf_proto_demo2_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_proto_demo2_proto_rawDescData)
	})
	return file_protobuf_proto_demo2_proto_rawDescData
}

var file_protobuf_proto_demo2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_proto_demo2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_proto_demo2_proto_goTypes = []interface{}{
	(Person_Gender)(0),   // 0: example.Person.Gender
	(*Person)(nil),       // 1: example.Person
	(*Person_Other)(nil), // 2: example.Person.Other
}
var file_protobuf_proto_demo2_proto_depIdxs = []int32{
	2, // 0: example.Person.info:type_name -> example.Person.Other
	0, // 1: example.Person.Other.g:type_name -> example.Person.Gender
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protobuf_proto_demo2_proto_init() }
func file_protobuf_proto_demo2_proto_init() {
	if File_protobuf_proto_demo2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_proto_demo2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_protobuf_proto_demo2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_Other); i {
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
			RawDescriptor: file_protobuf_proto_demo2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_proto_demo2_proto_goTypes,
		DependencyIndexes: file_protobuf_proto_demo2_proto_depIdxs,
		EnumInfos:         file_protobuf_proto_demo2_proto_enumTypes,
		MessageInfos:      file_protobuf_proto_demo2_proto_msgTypes,
	}.Build()
	File_protobuf_proto_demo2_proto = out.File
	file_protobuf_proto_demo2_proto_rawDesc = nil
	file_protobuf_proto_demo2_proto_goTypes = nil
	file_protobuf_proto_demo2_proto_depIdxs = nil
}
