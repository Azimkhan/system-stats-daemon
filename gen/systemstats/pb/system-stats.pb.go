// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.1
// source: system-stats.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type SystemStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuLoadAverage []*CPULoadAverage      `protobuf:"bytes,1,rep,name=cpuLoadAverage,proto3" json:"cpuLoadAverage,omitempty"`
	DiskLoad       []*DiskLoad            `protobuf:"bytes,2,rep,name=diskLoad,proto3" json:"diskLoad,omitempty"`
	Timestamp      *timestamppb.Timestamp `protobuf:"bytes,99,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SystemStatsResponse) Reset() {
	*x = SystemStatsResponse{}
	mi := &file_system_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SystemStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemStatsResponse) ProtoMessage() {}

func (x *SystemStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemStatsResponse.ProtoReflect.Descriptor instead.
func (*SystemStatsResponse) Descriptor() ([]byte, []int) {
	return file_system_stats_proto_rawDescGZIP(), []int{0}
}

func (x *SystemStatsResponse) GetCpuLoadAverage() []*CPULoadAverage {
	if x != nil {
		return x.CpuLoadAverage
	}
	return nil
}

func (x *SystemStatsResponse) GetDiskLoad() []*DiskLoad {
	if x != nil {
		return x.DiskLoad
	}
	return nil
}

func (x *SystemStatsResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type CPULoadAverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinutesAgo  uint32  `protobuf:"varint,1,opt,name=minutesAgo,proto3" json:"minutesAgo,omitempty"`
	AverageLoad float32 `protobuf:"fixed32,2,opt,name=averageLoad,proto3" json:"averageLoad,omitempty"`
}

func (x *CPULoadAverage) Reset() {
	*x = CPULoadAverage{}
	mi := &file_system_stats_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CPULoadAverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPULoadAverage) ProtoMessage() {}

func (x *CPULoadAverage) ProtoReflect() protoreflect.Message {
	mi := &file_system_stats_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPULoadAverage.ProtoReflect.Descriptor instead.
func (*CPULoadAverage) Descriptor() ([]byte, []int) {
	return file_system_stats_proto_rawDescGZIP(), []int{1}
}

func (x *CPULoadAverage) GetMinutesAgo() uint32 {
	if x != nil {
		return x.MinutesAgo
	}
	return 0
}

func (x *CPULoadAverage) GetAverageLoad() float32 {
	if x != nil {
		return x.AverageLoad
	}
	return 0
}

type DiskLoad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device                string  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	TransactionsPerSecond float32 `protobuf:"fixed32,2,opt,name=transactionsPerSecond,proto3" json:"transactionsPerSecond,omitempty"`
	Throughput            float32 `protobuf:"fixed32,3,opt,name=throughput,proto3" json:"throughput,omitempty"`
}

func (x *DiskLoad) Reset() {
	*x = DiskLoad{}
	mi := &file_system_stats_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiskLoad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskLoad) ProtoMessage() {}

func (x *DiskLoad) ProtoReflect() protoreflect.Message {
	mi := &file_system_stats_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskLoad.ProtoReflect.Descriptor instead.
func (*DiskLoad) Descriptor() ([]byte, []int) {
	return file_system_stats_proto_rawDescGZIP(), []int{2}
}

func (x *DiskLoad) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *DiskLoad) GetTransactionsPerSecond() float32 {
	if x != nil {
		return x.TransactionsPerSecond
	}
	return 0
}

func (x *DiskLoad) GetThroughput() float32 {
	if x != nil {
		return x.Throughput
	}
	return 0
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	mi := &file_system_stats_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_stats_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_system_stats_proto_rawDescGZIP(), []int{3}
}

var File_system_stats_proto protoreflect.FileDescriptor

var file_system_stats_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x13, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x0e, 0x63, 0x70, 0x75, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x50, 0x55, 0x4c, 0x6f, 0x61, 0x64, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x63, 0x70, 0x75, 0x4c, 0x6f, 0x61, 0x64, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x4c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x4c,
	0x6f, 0x61, 0x64, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x52, 0x0a, 0x0e, 0x43, 0x50, 0x55, 0x4c, 0x6f,
	0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x73, 0x41, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x41, 0x67, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x22, 0x78, 0x0a, 0x08, 0x44,
	0x69, 0x73, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x34, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x15,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x65, 0x72, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68,
	0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x68, 0x72, 0x6f, 0x75,
	0x67, 0x68, 0x70, 0x75, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x4f, 0x0a, 0x12, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0d, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_stats_proto_rawDescOnce sync.Once
	file_system_stats_proto_rawDescData = file_system_stats_proto_rawDesc
)

func file_system_stats_proto_rawDescGZIP() []byte {
	file_system_stats_proto_rawDescOnce.Do(func() {
		file_system_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_stats_proto_rawDescData)
	})
	return file_system_stats_proto_rawDescData
}

var file_system_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_system_stats_proto_goTypes = []any{
	(*SystemStatsResponse)(nil),   // 0: SystemStatsResponse
	(*CPULoadAverage)(nil),        // 1: CPULoadAverage
	(*DiskLoad)(nil),              // 2: DiskLoad
	(*EmptyRequest)(nil),          // 3: EmptyRequest
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_system_stats_proto_depIdxs = []int32{
	1, // 0: SystemStatsResponse.cpuLoadAverage:type_name -> CPULoadAverage
	2, // 1: SystemStatsResponse.diskLoad:type_name -> DiskLoad
	4, // 2: SystemStatsResponse.timestamp:type_name -> google.protobuf.Timestamp
	3, // 3: SystemStatsService.GetSystemStats:input_type -> EmptyRequest
	0, // 4: SystemStatsService.GetSystemStats:output_type -> SystemStatsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_system_stats_proto_init() }
func file_system_stats_proto_init() {
	if File_system_stats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_stats_proto_goTypes,
		DependencyIndexes: file_system_stats_proto_depIdxs,
		MessageInfos:      file_system_stats_proto_msgTypes,
	}.Build()
	File_system_stats_proto = out.File
	file_system_stats_proto_rawDesc = nil
	file_system_stats_proto_goTypes = nil
	file_system_stats_proto_depIdxs = nil
}