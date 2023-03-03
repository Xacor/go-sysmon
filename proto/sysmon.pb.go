// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: proto/sysmon.proto

package sysmonpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshRate     *durationpb.Duration `protobuf:"bytes,1,opt,name=refreshRate,proto3" json:"refreshRate,omitempty"`
	RefreshInterval *durationpb.Duration `protobuf:"bytes,2,opt,name=refreshInterval,proto3" json:"refreshInterval,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sysmon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sysmon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_sysmon_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetRefreshRate() *durationpb.Duration {
	if x != nil {
		return x.RefreshRate
	}
	return nil
}

func (x *Request) GetRefreshInterval() *durationpb.Duration {
	if x != nil {
		return x.RefreshInterval
	}
	return nil
}

type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoadAverage *LoadAverage           `protobuf:"bytes,1,opt,name=loadAverage,proto3" json:"loadAverage,omitempty"`
	TimeCreated *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timeCreated,proto3" json:"timeCreated,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sysmon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sysmon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_proto_sysmon_proto_rawDescGZIP(), []int{1}
}

func (x *Snapshot) GetLoadAverage() *LoadAverage {
	if x != nil {
		return x.LoadAverage
	}
	return nil
}

func (x *Snapshot) GetTimeCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeCreated
	}
	return nil
}

type LoadAverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Load1  float32 `protobuf:"fixed32,1,opt,name=load1,proto3" json:"load1,omitempty"`
	Load5  float32 `protobuf:"fixed32,2,opt,name=load5,proto3" json:"load5,omitempty"`
	Load15 float32 `protobuf:"fixed32,3,opt,name=load15,proto3" json:"load15,omitempty"`
}

func (x *LoadAverage) Reset() {
	*x = LoadAverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sysmon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadAverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadAverage) ProtoMessage() {}

func (x *LoadAverage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sysmon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadAverage.ProtoReflect.Descriptor instead.
func (*LoadAverage) Descriptor() ([]byte, []int) {
	return file_proto_sysmon_proto_rawDescGZIP(), []int{2}
}

func (x *LoadAverage) GetLoad1() float32 {
	if x != nil {
		return x.Load1
	}
	return 0
}

func (x *LoadAverage) GetLoad5() float32 {
	if x != nil {
		return x.Load5
	}
	return 0
}

func (x *LoadAverage) GetLoad15() float32 {
	if x != nil {
		return x.Load15
	}
	return 0
}

var File_proto_sysmon_proto protoreflect.FileDescriptor

var file_proto_sysmon_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x79, 0x73, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x79, 0x73, 0x6d, 0x6f, 0x6e, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x7f, 0x0a, 0x08, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x35, 0x0a, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x79, 0x73, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x3c,
	0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x51, 0x0a, 0x0b,
	0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x6f, 0x61, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6c, 0x6f, 0x61, 0x64,
	0x31, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x61, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x6c, 0x6f, 0x61, 0x64, 0x35, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x61, 0x64, 0x31,
	0x35, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x35, 0x32,
	0x3e, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x4d, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x79, 0x73, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x6d,
	0x6f, 0x6e, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x61,
	0x63, 0x6f, 0x72, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x79, 0x73, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x79,
	0x73, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sysmon_proto_rawDescOnce sync.Once
	file_proto_sysmon_proto_rawDescData = file_proto_sysmon_proto_rawDesc
)

func file_proto_sysmon_proto_rawDescGZIP() []byte {
	file_proto_sysmon_proto_rawDescOnce.Do(func() {
		file_proto_sysmon_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sysmon_proto_rawDescData)
	})
	return file_proto_sysmon_proto_rawDescData
}

var file_proto_sysmon_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_sysmon_proto_goTypes = []interface{}{
	(*Request)(nil),               // 0: sysmon.Request
	(*Snapshot)(nil),              // 1: sysmon.Snapshot
	(*LoadAverage)(nil),           // 2: sysmon.LoadAverage
	(*durationpb.Duration)(nil),   // 3: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_proto_sysmon_proto_depIdxs = []int32{
	3, // 0: sysmon.Request.refreshRate:type_name -> google.protobuf.Duration
	3, // 1: sysmon.Request.refreshInterval:type_name -> google.protobuf.Duration
	2, // 2: sysmon.Snapshot.loadAverage:type_name -> sysmon.LoadAverage
	4, // 3: sysmon.Snapshot.timeCreated:type_name -> google.protobuf.Timestamp
	0, // 4: sysmon.SysMon.GetSnapshot:input_type -> sysmon.Request
	1, // 5: sysmon.SysMon.GetSnapshot:output_type -> sysmon.Snapshot
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_sysmon_proto_init() }
func file_proto_sysmon_proto_init() {
	if File_proto_sysmon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sysmon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_sysmon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
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
		file_proto_sysmon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadAverage); i {
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
			RawDescriptor: file_proto_sysmon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sysmon_proto_goTypes,
		DependencyIndexes: file_proto_sysmon_proto_depIdxs,
		MessageInfos:      file_proto_sysmon_proto_msgTypes,
	}.Build()
	File_proto_sysmon_proto = out.File
	file_proto_sysmon_proto_rawDesc = nil
	file_proto_sysmon_proto_goTypes = nil
	file_proto_sysmon_proto_depIdxs = nil
}
