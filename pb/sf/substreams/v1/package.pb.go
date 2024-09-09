// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: sf/substreams/v1/package.proto

package pbsubstreams

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Needs to be one so this file can be used _directly_ as a
	// buf `Image` andor a ProtoSet for grpcurl and other tools
	ProtoFiles  []*descriptorpb.FileDescriptorProto `protobuf:"bytes,1,rep,name=proto_files,json=protoFiles,proto3" json:"proto_files,omitempty"`
	Version     uint64                              `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	Modules     *Modules                            `protobuf:"bytes,6,opt,name=modules,proto3" json:"modules,omitempty"`
	ModuleMeta  []*ModuleMetadata                   `protobuf:"bytes,7,rep,name=module_meta,json=moduleMeta,proto3" json:"module_meta,omitempty"`
	PackageMeta []*PackageMetadata                  `protobuf:"bytes,8,rep,name=package_meta,json=packageMeta,proto3" json:"package_meta,omitempty"`
	// Source network for Substreams to fetch its data from.
	Network    string     `protobuf:"bytes,9,opt,name=network,proto3" json:"network,omitempty"`
	SinkConfig *anypb.Any `protobuf:"bytes,10,opt,name=sink_config,json=sinkConfig,proto3" json:"sink_config,omitempty"`
	SinkModule string     `protobuf:"bytes,11,opt,name=sink_module,json=sinkModule,proto3" json:"sink_module,omitempty"`
	// image is the bytes to a JPEG, WebP or PNG file. Max size is 2 MiB
	Image        []byte                    `protobuf:"bytes,12,opt,name=image,proto3" json:"image,omitempty"`
	Networks     map[string]*NetworkParams `protobuf:"bytes,13,rep,name=networks,proto3" json:"networks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BlockFilters map[string]string         `protobuf:"bytes,14,rep,name=block_filters,json=blockFilters,proto3" json:"block_filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_v1_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_v1_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_sf_substreams_v1_package_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetProtoFiles() []*descriptorpb.FileDescriptorProto {
	if x != nil {
		return x.ProtoFiles
	}
	return nil
}

func (x *Package) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Package) GetModules() *Modules {
	if x != nil {
		return x.Modules
	}
	return nil
}

func (x *Package) GetModuleMeta() []*ModuleMetadata {
	if x != nil {
		return x.ModuleMeta
	}
	return nil
}

func (x *Package) GetPackageMeta() []*PackageMetadata {
	if x != nil {
		return x.PackageMeta
	}
	return nil
}

func (x *Package) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Package) GetSinkConfig() *anypb.Any {
	if x != nil {
		return x.SinkConfig
	}
	return nil
}

func (x *Package) GetSinkModule() string {
	if x != nil {
		return x.SinkModule
	}
	return ""
}

func (x *Package) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Package) GetNetworks() map[string]*NetworkParams {
	if x != nil {
		return x.Networks
	}
	return nil
}

func (x *Package) GetBlockFilters() map[string]string {
	if x != nil {
		return x.BlockFilters
	}
	return nil
}

type NetworkParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitialBlocks map[string]uint64 `protobuf:"bytes,1,rep,name=initialBlocks,proto3" json:"initialBlocks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Params        map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NetworkParams) Reset() {
	*x = NetworkParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_v1_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkParams) ProtoMessage() {}

func (x *NetworkParams) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_v1_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkParams.ProtoReflect.Descriptor instead.
func (*NetworkParams) Descriptor() ([]byte, []int) {
	return file_sf_substreams_v1_package_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkParams) GetInitialBlocks() map[string]uint64 {
	if x != nil {
		return x.InitialBlocks
	}
	return nil
}

func (x *NetworkParams) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type PackageMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Doc     string `protobuf:"bytes,4,opt,name=doc,proto3" json:"doc,omitempty"`
}

func (x *PackageMetadata) Reset() {
	*x = PackageMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_v1_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageMetadata) ProtoMessage() {}

func (x *PackageMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_v1_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageMetadata.ProtoReflect.Descriptor instead.
func (*PackageMetadata) Descriptor() ([]byte, []int) {
	return file_sf_substreams_v1_package_proto_rawDescGZIP(), []int{2}
}

func (x *PackageMetadata) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PackageMetadata) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PackageMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PackageMetadata) GetDoc() string {
	if x != nil {
		return x.Doc
	}
	return ""
}

type ModuleMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Corresponds to the index in `Package.metadata.package_meta`
	PackageIndex uint64 `protobuf:"varint,1,opt,name=package_index,json=packageIndex,proto3" json:"package_index,omitempty"`
	Doc          string `protobuf:"bytes,2,opt,name=doc,proto3" json:"doc,omitempty"`
}

func (x *ModuleMetadata) Reset() {
	*x = ModuleMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_v1_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleMetadata) ProtoMessage() {}

func (x *ModuleMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_v1_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleMetadata.ProtoReflect.Descriptor instead.
func (*ModuleMetadata) Descriptor() ([]byte, []int) {
	return file_sf_substreams_v1_package_proto_rawDescGZIP(), []int{3}
}

func (x *ModuleMetadata) GetPackageIndex() uint64 {
	if x != nil {
		return x.PackageIndex
	}
	return 0
}

func (x *ModuleMetadata) GetDoc() string {
	if x != nil {
		return x.Doc
	}
	return ""
}

var File_sf_substreams_v1_package_proto protoreflect.FileDescriptor

var file_sf_substreams_v1_package_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xec, 0x05, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x07,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x41, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x44, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x66, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x0a, 0x73, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x69, 0x6e, 0x6b, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x69, 0x6e, 0x6b, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x0d,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x50, 0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x5c, 0x0a, 0x0d, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x66,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x05, 0x22, 0xab,
	0x02, 0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x58, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x43, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x66, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x40, 0x0a, 0x12, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x63, 0x0a, 0x0f,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x6f, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f,
	0x63, 0x22, 0x47, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x63, 0x42, 0xcc, 0x01, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x73,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa,
	0x02, 0x10, 0x53, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x10, 0x53, 0x66, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x53, 0x66, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x53, 0x66, 0x3a, 0x3a, 0x53, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sf_substreams_v1_package_proto_rawDescOnce sync.Once
	file_sf_substreams_v1_package_proto_rawDescData = file_sf_substreams_v1_package_proto_rawDesc
)

func file_sf_substreams_v1_package_proto_rawDescGZIP() []byte {
	file_sf_substreams_v1_package_proto_rawDescOnce.Do(func() {
		file_sf_substreams_v1_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_substreams_v1_package_proto_rawDescData)
	})
	return file_sf_substreams_v1_package_proto_rawDescData
}

var file_sf_substreams_v1_package_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sf_substreams_v1_package_proto_goTypes = []interface{}{
	(*Package)(nil),                          // 0: sf.substreams.v1.Package
	(*NetworkParams)(nil),                    // 1: sf.substreams.v1.NetworkParams
	(*PackageMetadata)(nil),                  // 2: sf.substreams.v1.PackageMetadata
	(*ModuleMetadata)(nil),                   // 3: sf.substreams.v1.ModuleMetadata
	nil,                                      // 4: sf.substreams.v1.Package.NetworksEntry
	nil,                                      // 5: sf.substreams.v1.Package.BlockFiltersEntry
	nil,                                      // 6: sf.substreams.v1.NetworkParams.InitialBlocksEntry
	nil,                                      // 7: sf.substreams.v1.NetworkParams.ParamsEntry
	(*descriptorpb.FileDescriptorProto)(nil), // 8: google.protobuf.FileDescriptorProto
	(*Modules)(nil),                          // 9: sf.substreams.v1.Modules
	(*anypb.Any)(nil),                        // 10: google.protobuf.Any
}
var file_sf_substreams_v1_package_proto_depIdxs = []int32{
	8,  // 0: sf.substreams.v1.Package.proto_files:type_name -> google.protobuf.FileDescriptorProto
	9,  // 1: sf.substreams.v1.Package.modules:type_name -> sf.substreams.v1.Modules
	3,  // 2: sf.substreams.v1.Package.module_meta:type_name -> sf.substreams.v1.ModuleMetadata
	2,  // 3: sf.substreams.v1.Package.package_meta:type_name -> sf.substreams.v1.PackageMetadata
	10, // 4: sf.substreams.v1.Package.sink_config:type_name -> google.protobuf.Any
	4,  // 5: sf.substreams.v1.Package.networks:type_name -> sf.substreams.v1.Package.NetworksEntry
	5,  // 6: sf.substreams.v1.Package.block_filters:type_name -> sf.substreams.v1.Package.BlockFiltersEntry
	6,  // 7: sf.substreams.v1.NetworkParams.initialBlocks:type_name -> sf.substreams.v1.NetworkParams.InitialBlocksEntry
	7,  // 8: sf.substreams.v1.NetworkParams.params:type_name -> sf.substreams.v1.NetworkParams.ParamsEntry
	1,  // 9: sf.substreams.v1.Package.NetworksEntry.value:type_name -> sf.substreams.v1.NetworkParams
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_sf_substreams_v1_package_proto_init() }
func file_sf_substreams_v1_package_proto_init() {
	if File_sf_substreams_v1_package_proto != nil {
		return
	}
	file_sf_substreams_v1_modules_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sf_substreams_v1_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_sf_substreams_v1_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkParams); i {
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
		file_sf_substreams_v1_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageMetadata); i {
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
		file_sf_substreams_v1_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleMetadata); i {
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
			RawDescriptor: file_sf_substreams_v1_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_substreams_v1_package_proto_goTypes,
		DependencyIndexes: file_sf_substreams_v1_package_proto_depIdxs,
		MessageInfos:      file_sf_substreams_v1_package_proto_msgTypes,
	}.Build()
	File_sf_substreams_v1_package_proto = out.File
	file_sf_substreams_v1_package_proto_rawDesc = nil
	file_sf_substreams_v1_package_proto_goTypes = nil
	file_sf_substreams_v1_package_proto_depIdxs = nil
}
