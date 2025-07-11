// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: master.proto

package proto_gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterWorkerReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkerIp      string                 `protobuf:"bytes,1,opt,name=worker_ip,json=workerIp,proto3" json:"worker_ip,omitempty"`
	Uuid          string                 `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterWorkerReq) Reset() {
	*x = RegisterWorkerReq{}
	mi := &file_master_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterWorkerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterWorkerReq) ProtoMessage() {}

func (x *RegisterWorkerReq) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterWorkerReq.ProtoReflect.Descriptor instead.
func (*RegisterWorkerReq) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterWorkerReq) GetWorkerIp() string {
	if x != nil {
		return x.WorkerIp
	}
	return ""
}

func (x *RegisterWorkerReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type RegisterWorkerRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsSuccess     bool                   `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Id            int64                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterWorkerRes) Reset() {
	*x = RegisterWorkerRes{}
	mi := &file_master_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterWorkerRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterWorkerRes) ProtoMessage() {}

func (x *RegisterWorkerRes) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterWorkerRes.ProtoReflect.Descriptor instead.
func (*RegisterWorkerRes) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterWorkerRes) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *RegisterWorkerRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_master_proto protoreflect.FileDescriptor

const file_master_proto_rawDesc = "" +
	"\n" +
	"\fmaster.proto\"D\n" +
	"\x11RegisterWorkerReq\x12\x1b\n" +
	"\tworker_ip\x18\x01 \x01(\tR\bworkerIp\x12\x12\n" +
	"\x04uuid\x18\x02 \x01(\tR\x04uuid\"B\n" +
	"\x11RegisterWorkerRes\x12\x1d\n" +
	"\n" +
	"is_success\x18\x01 \x01(\bR\tisSuccess\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\x03R\x02id2B\n" +
	"\x06Master\x128\n" +
	"\x0eRegisterWorker\x12\x12.RegisterWorkerReq\x1a\x12.RegisterWorkerResB\rZ\v./proto_genb\x06proto3"

var (
	file_master_proto_rawDescOnce sync.Once
	file_master_proto_rawDescData []byte
)

func file_master_proto_rawDescGZIP() []byte {
	file_master_proto_rawDescOnce.Do(func() {
		file_master_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_master_proto_rawDesc), len(file_master_proto_rawDesc)))
	})
	return file_master_proto_rawDescData
}

var file_master_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_master_proto_goTypes = []any{
	(*RegisterWorkerReq)(nil), // 0: RegisterWorkerReq
	(*RegisterWorkerRes)(nil), // 1: RegisterWorkerRes
}
var file_master_proto_depIdxs = []int32{
	0, // 0: Master.RegisterWorker:input_type -> RegisterWorkerReq
	1, // 1: Master.RegisterWorker:output_type -> RegisterWorkerRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_master_proto_init() }
func file_master_proto_init() {
	if File_master_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_master_proto_rawDesc), len(file_master_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_master_proto_goTypes,
		DependencyIndexes: file_master_proto_depIdxs,
		MessageInfos:      file_master_proto_msgTypes,
	}.Build()
	File_master_proto = out.File
	file_master_proto_goTypes = nil
	file_master_proto_depIdxs = nil
}
