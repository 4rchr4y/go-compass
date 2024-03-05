// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: rayplugin1.0.proto

package pluginproto

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

type Schema_NestedBlock_EmbeddedMode int32

const (
	Schema_NestedBlock_INVALID Schema_NestedBlock_EmbeddedMode = 0
)

// Enum value maps for Schema_NestedBlock_EmbeddedMode.
var (
	Schema_NestedBlock_EmbeddedMode_name = map[int32]string{
		0: "INVALID",
	}
	Schema_NestedBlock_EmbeddedMode_value = map[string]int32{
		"INVALID": 0,
	}
)

func (x Schema_NestedBlock_EmbeddedMode) Enum() *Schema_NestedBlock_EmbeddedMode {
	p := new(Schema_NestedBlock_EmbeddedMode)
	*p = x
	return p
}

func (x Schema_NestedBlock_EmbeddedMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schema_NestedBlock_EmbeddedMode) Descriptor() protoreflect.EnumDescriptor {
	return file_rayplugin1_0_proto_enumTypes[0].Descriptor()
}

func (Schema_NestedBlock_EmbeddedMode) Type() protoreflect.EnumType {
	return &file_rayplugin1_0_proto_enumTypes[0]
}

func (x Schema_NestedBlock_EmbeddedMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schema_NestedBlock_EmbeddedMode.Descriptor instead.
func (Schema_NestedBlock_EmbeddedMode) EnumDescriptor() ([]byte, []int) {
	return file_rayplugin1_0_proto_rawDescGZIP(), []int{0, 2, 0}
}

type Schema_Object_EmbeddedMode int32

const (
	Schema_Object_INVALID Schema_Object_EmbeddedMode = 0
)

// Enum value maps for Schema_Object_EmbeddedMode.
var (
	Schema_Object_EmbeddedMode_name = map[int32]string{
		0: "INVALID",
	}
	Schema_Object_EmbeddedMode_value = map[string]int32{
		"INVALID": 0,
	}
)

func (x Schema_Object_EmbeddedMode) Enum() *Schema_Object_EmbeddedMode {
	p := new(Schema_Object_EmbeddedMode)
	*p = x
	return p
}

func (x Schema_Object_EmbeddedMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schema_Object_EmbeddedMode) Descriptor() protoreflect.EnumDescriptor {
	return file_rayplugin1_0_proto_enumTypes[1].Descriptor()
}

func (Schema_Object_EmbeddedMode) Type() protoreflect.EnumType {
	return &file_rayplugin1_0_proto_enumTypes[1]
}

func (x Schema_Object_EmbeddedMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schema_Object_EmbeddedMode.Descriptor instead.
func (Schema_Object_EmbeddedMode) EnumDescriptor() ([]byte, []int) {
	return file_rayplugin1_0_proto_rawDescGZIP(), []int{0, 3, 0}
}

type Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *Schema_Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *Schema) Reset() {
	*x = Schema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rayplugin1_0_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_rayplugin1_0_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_rayplugin1_0_proto_rawDescGZIP(), []int{0}
}

func (x *Schema) GetBlock() *Schema_Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type DescribeSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeSchema) Reset() {
	*x = DescribeSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rayplugin1_0_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSchema) ProtoMessage() {}

func (x *DescribeSchema) ProtoReflect() protoreflect.Message {
	mi := &file_rayplugin1_0_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSchema.ProtoReflect.Descriptor instead.
func (*DescribeSchema) Descriptor() ([]byte, []int) {
	return file_rayplugin1_0_proto_rawDescGZIP(), []int{1}
}

type Schema_Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes  []*Schema_Attribute   `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	BlockTypes  []*Schema_NestedBlock `protobuf:"bytes,2,rep,name=block_types,json=blockTypes,proto3" json:"block_types,omitempty"`
	Description string                `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Deprecated  bool                  `protobuf:"varint,4,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
}

func (x *Schema_Block) Reset() {
	*x = Schema_Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rayplugin1_0_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema_Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema_Block) ProtoMessage() {}

func (x *Schema_Block) ProtoReflect() protoreflect.Message {
	mi := &file_rayplugin1_0_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema_Block.ProtoReflect.Descriptor instead.
func (*Schema_Block) Descriptor() ([]byte, []int) {
	return file_rayplugin1_0_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Schema_Block) GetAttributes() []*Schema_Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Schema_Block) GetBlockTypes() []*Schema_NestedBlock {
	if x != nil {
		return x.BlockTypes
	}
	return nil
}

func (x *Schema_Block) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Schema_Block) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

type Schema_Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         []byte         `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	EmbeddedType *Schema_Object `protobuf:"bytes,2,opt,name=embedded_type,json=embeddedType,proto3" json:"embedded_type,omitempty"`
	Description  string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Required     bool           `protobuf:"varint,4,opt,name=required,proto3" json:"required,omitempty"`
	Optional     bool           `protobuf:"varint,5,opt,name=optional,proto3" json:"optional,omitempty"`
	Deprecated   bool           `protobuf:"varint,6,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
}

func (x *Schema_Attribute) Reset() {
	*x = Schema_Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rayplugin1_0_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema_Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema_Attribute) ProtoMessage() {}

func (x *Schema_Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_rayplugin1_0_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema_Attribute.ProtoReflect.Descriptor instead.
func (*Schema_Attribute) Descriptor() ([]byte, []int) {
	return file_rayplugin1_0_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Schema_Attribute) GetType() []byte {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Schema_Attribute) GetEmbeddedType() *Schema_Object {
	if x != nil {
		return x.EmbeddedType
	}
	return nil
}

func (x *Schema_Attribute) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Schema_Attribute) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *Schema_Attribute) GetOptional() bool {
	if x != nil {
		return x.Optional
	}
	return false
}

func (x *Schema_Attribute) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

type Schema_NestedBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block   *Schema_Block                   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	Nesting Schema_NestedBlock_EmbeddedMode `protobuf:"varint,3,opt,name=nesting,proto3,enum=pluginproto.Schema_NestedBlock_EmbeddedMode" json:"nesting,omitempty"`
}

func (x *Schema_NestedBlock) Reset() {
	*x = Schema_NestedBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rayplugin1_0_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema_NestedBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema_NestedBlock) ProtoMessage() {}

func (x *Schema_NestedBlock) ProtoReflect() protoreflect.Message {
	mi := &file_rayplugin1_0_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema_NestedBlock.ProtoReflect.Descriptor instead.
func (*Schema_NestedBlock) Descriptor() ([]byte, []int) {
	return file_rayplugin1_0_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Schema_NestedBlock) GetBlock() *Schema_Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *Schema_NestedBlock) GetNesting() Schema_NestedBlock_EmbeddedMode {
	if x != nil {
		return x.Nesting
	}
	return Schema_NestedBlock_INVALID
}

type Schema_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes []*Schema_Attribute        `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Nesting    Schema_Object_EmbeddedMode `protobuf:"varint,3,opt,name=nesting,proto3,enum=pluginproto.Schema_Object_EmbeddedMode" json:"nesting,omitempty"`
}

func (x *Schema_Object) Reset() {
	*x = Schema_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rayplugin1_0_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema_Object) ProtoMessage() {}

func (x *Schema_Object) ProtoReflect() protoreflect.Message {
	mi := &file_rayplugin1_0_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema_Object.ProtoReflect.Descriptor instead.
func (*Schema_Object) Descriptor() ([]byte, []int) {
	return file_rayplugin1_0_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Schema_Object) GetAttributes() []*Schema_Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Schema_Object) GetNesting() Schema_Object_EmbeddedMode {
	if x != nil {
		return x.Nesting
	}
	return Schema_Object_INVALID
}

type DescribeSchema_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeSchema_Request) Reset() {
	*x = DescribeSchema_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rayplugin1_0_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSchema_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSchema_Request) ProtoMessage() {}

func (x *DescribeSchema_Request) ProtoReflect() protoreflect.Message {
	mi := &file_rayplugin1_0_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSchema_Request.ProtoReflect.Descriptor instead.
func (*DescribeSchema_Request) Descriptor() ([]byte, []int) {
	return file_rayplugin1_0_proto_rawDescGZIP(), []int{1, 0}
}

type DescribeSchema_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider *Schema `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *DescribeSchema_Response) Reset() {
	*x = DescribeSchema_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rayplugin1_0_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSchema_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSchema_Response) ProtoMessage() {}

func (x *DescribeSchema_Response) ProtoReflect() protoreflect.Message {
	mi := &file_rayplugin1_0_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSchema_Response.ProtoReflect.Descriptor instead.
func (*DescribeSchema_Response) Descriptor() ([]byte, []int) {
	return file_rayplugin1_0_proto_rawDescGZIP(), []int{1, 1}
}

func (x *DescribeSchema_Response) GetProvider() *Schema {
	if x != nil {
		return x.Provider
	}
	return nil
}

var File_rayplugin1_0_proto protoreflect.FileDescriptor

var file_rayplugin1_0_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x61, 0x79, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x31, 0x2e, 0x30, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb3, 0x06, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x2f, 0x0a, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0xca, 0x01,
	0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65,
	0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x1a, 0xda, 0x01, 0x0a, 0x09, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x0d,
	0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x0c, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x1a, 0xa3, 0x01, 0x0a, 0x0b, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2f, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x46, 0x0a, 0x07, 0x6e, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x6e, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x1b, 0x0a, 0x0c, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x1a, 0xa7, 0x01,
	0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x07, 0x6e, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x07, 0x6e, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x1b, 0x0a, 0x0c, 0x45, 0x6d,
	0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x22, 0x58, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x32, 0x67, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x5b, 0x0a,
	0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x23, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x34, 0x72, 0x63, 0x68, 0x72, 0x34, 0x79,
	0x2f, 0x67, 0x6f, 0x72, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rayplugin1_0_proto_rawDescOnce sync.Once
	file_rayplugin1_0_proto_rawDescData = file_rayplugin1_0_proto_rawDesc
)

func file_rayplugin1_0_proto_rawDescGZIP() []byte {
	file_rayplugin1_0_proto_rawDescOnce.Do(func() {
		file_rayplugin1_0_proto_rawDescData = protoimpl.X.CompressGZIP(file_rayplugin1_0_proto_rawDescData)
	})
	return file_rayplugin1_0_proto_rawDescData
}

var file_rayplugin1_0_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_rayplugin1_0_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rayplugin1_0_proto_goTypes = []interface{}{
	(Schema_NestedBlock_EmbeddedMode)(0), // 0: pluginproto.Schema.NestedBlock.EmbeddedMode
	(Schema_Object_EmbeddedMode)(0),      // 1: pluginproto.Schema.Object.EmbeddedMode
	(*Schema)(nil),                       // 2: pluginproto.Schema
	(*DescribeSchema)(nil),               // 3: pluginproto.DescribeSchema
	(*Schema_Block)(nil),                 // 4: pluginproto.Schema.Block
	(*Schema_Attribute)(nil),             // 5: pluginproto.Schema.Attribute
	(*Schema_NestedBlock)(nil),           // 6: pluginproto.Schema.NestedBlock
	(*Schema_Object)(nil),                // 7: pluginproto.Schema.Object
	(*DescribeSchema_Request)(nil),       // 8: pluginproto.DescribeSchema.Request
	(*DescribeSchema_Response)(nil),      // 9: pluginproto.DescribeSchema.Response
}
var file_rayplugin1_0_proto_depIdxs = []int32{
	4,  // 0: pluginproto.Schema.block:type_name -> pluginproto.Schema.Block
	5,  // 1: pluginproto.Schema.Block.attributes:type_name -> pluginproto.Schema.Attribute
	6,  // 2: pluginproto.Schema.Block.block_types:type_name -> pluginproto.Schema.NestedBlock
	7,  // 3: pluginproto.Schema.Attribute.embedded_type:type_name -> pluginproto.Schema.Object
	4,  // 4: pluginproto.Schema.NestedBlock.block:type_name -> pluginproto.Schema.Block
	0,  // 5: pluginproto.Schema.NestedBlock.nesting:type_name -> pluginproto.Schema.NestedBlock.EmbeddedMode
	5,  // 6: pluginproto.Schema.Object.attributes:type_name -> pluginproto.Schema.Attribute
	1,  // 7: pluginproto.Schema.Object.nesting:type_name -> pluginproto.Schema.Object.EmbeddedMode
	2,  // 8: pluginproto.DescribeSchema.Response.provider:type_name -> pluginproto.Schema
	8,  // 9: pluginproto.Provider.DescribeSchema:input_type -> pluginproto.DescribeSchema.Request
	9,  // 10: pluginproto.Provider.DescribeSchema:output_type -> pluginproto.DescribeSchema.Response
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_rayplugin1_0_proto_init() }
func file_rayplugin1_0_proto_init() {
	if File_rayplugin1_0_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rayplugin1_0_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema); i {
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
		file_rayplugin1_0_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSchema); i {
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
		file_rayplugin1_0_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema_Block); i {
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
		file_rayplugin1_0_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema_Attribute); i {
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
		file_rayplugin1_0_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema_NestedBlock); i {
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
		file_rayplugin1_0_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema_Object); i {
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
		file_rayplugin1_0_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSchema_Request); i {
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
		file_rayplugin1_0_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSchema_Response); i {
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
			RawDescriptor: file_rayplugin1_0_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rayplugin1_0_proto_goTypes,
		DependencyIndexes: file_rayplugin1_0_proto_depIdxs,
		EnumInfos:         file_rayplugin1_0_proto_enumTypes,
		MessageInfos:      file_rayplugin1_0_proto_msgTypes,
	}.Build()
	File_rayplugin1_0_proto = out.File
	file_rayplugin1_0_proto_rawDesc = nil
	file_rayplugin1_0_proto_goTypes = nil
	file_rayplugin1_0_proto_depIdxs = nil
}
