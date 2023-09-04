// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.3
// source: fixtures/fixtures.proto

package fixtures

import (
	plumber "github.com/tinkershack/muxfarm/plumber"
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

type IngestDoc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Muxfarmid *plumber.MuxfarmID     `protobuf:"bytes,1,opt,name=muxfarmid,proto3" json:"muxfarmid,omitempty"`
	Mediain   *plumber.MediaIn       `protobuf:"bytes,2,opt,name=mediain,proto3" json:"mediain,omitempty"`
	Status    *plumber.Status        `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *IngestDoc) Reset() {
	*x = IngestDoc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_fixtures_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestDoc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestDoc) ProtoMessage() {}

func (x *IngestDoc) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_fixtures_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestDoc.ProtoReflect.Descriptor instead.
func (*IngestDoc) Descriptor() ([]byte, []int) {
	return file_fixtures_fixtures_proto_rawDescGZIP(), []int{0}
}

func (x *IngestDoc) GetMuxfarmid() *plumber.MuxfarmID {
	if x != nil {
		return x.Muxfarmid
	}
	return nil
}

func (x *IngestDoc) GetMediain() *plumber.MediaIn {
	if x != nil {
		return x.Mediain
	}
	return nil
}

func (x *IngestDoc) GetStatus() *plumber.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *IngestDoc) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type IngestBatchDoc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Muxfarmid *plumber.MuxfarmID     `protobuf:"bytes,1,opt,name=muxfarmid,proto3" json:"muxfarmid,omitempty"`
	Ingestid  *plumber.MuxfarmID     `protobuf:"bytes,2,opt,name=ingestid,proto3" json:"ingestid,omitempty"`
	Media     *plumber.Media         `protobuf:"bytes,3,opt,name=media,proto3" json:"media,omitempty"`
	Status    *plumber.Status        `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *IngestBatchDoc) Reset() {
	*x = IngestBatchDoc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_fixtures_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestBatchDoc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestBatchDoc) ProtoMessage() {}

func (x *IngestBatchDoc) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_fixtures_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestBatchDoc.ProtoReflect.Descriptor instead.
func (*IngestBatchDoc) Descriptor() ([]byte, []int) {
	return file_fixtures_fixtures_proto_rawDescGZIP(), []int{1}
}

func (x *IngestBatchDoc) GetMuxfarmid() *plumber.MuxfarmID {
	if x != nil {
		return x.Muxfarmid
	}
	return nil
}

func (x *IngestBatchDoc) GetIngestid() *plumber.MuxfarmID {
	if x != nil {
		return x.Ingestid
	}
	return nil
}

func (x *IngestBatchDoc) GetMedia() *plumber.Media {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *IngestBatchDoc) GetStatus() *plumber.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *IngestBatchDoc) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_fixtures_fixtures_proto protoreflect.FileDescriptor

var file_fixtures_fixtures_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x78, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x66, 0x69, 0x78, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x1a, 0x15, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x6c, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe2, 0x01, 0x0a, 0x09, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x12, 0x38,
	0x0a, 0x09, 0x6d, 0x75, 0x78, 0x66, 0x61, 0x72, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x4d, 0x75, 0x78, 0x66, 0x61, 0x72, 0x6d, 0x49, 0x44, 0x52, 0x09, 0x6d,
	0x75, 0x78, 0x66, 0x61, 0x72, 0x6d, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6c, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x49, 0x6e, 0x52, 0x07, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x69, 0x6e, 0x12, 0x2d, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x99, 0x02, 0x0a, 0x0e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x6f, 0x63, 0x12, 0x38, 0x0a, 0x09, 0x6d, 0x75, 0x78, 0x66,
	0x61, 0x72, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x75,
	0x78, 0x66, 0x61, 0x72, 0x6d, 0x49, 0x44, 0x52, 0x09, 0x6d, 0x75, 0x78, 0x66, 0x61, 0x72, 0x6d,
	0x69, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70,
	0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x75, 0x78, 0x66, 0x61, 0x72, 0x6d, 0x49, 0x44,
	0x52, 0x08, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6c, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x73, 0x68, 0x61, 0x63, 0x6b, 0x2f, 0x6d, 0x75, 0x78, 0x66,
	0x61, 0x72, 0x6d, 0x2f, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fixtures_fixtures_proto_rawDescOnce sync.Once
	file_fixtures_fixtures_proto_rawDescData = file_fixtures_fixtures_proto_rawDesc
)

func file_fixtures_fixtures_proto_rawDescGZIP() []byte {
	file_fixtures_fixtures_proto_rawDescOnce.Do(func() {
		file_fixtures_fixtures_proto_rawDescData = protoimpl.X.CompressGZIP(file_fixtures_fixtures_proto_rawDescData)
	})
	return file_fixtures_fixtures_proto_rawDescData
}

var file_fixtures_fixtures_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fixtures_fixtures_proto_goTypes = []interface{}{
	(*IngestDoc)(nil),             // 0: fixtures.ingestDoc
	(*IngestBatchDoc)(nil),        // 1: fixtures.ingestBatchDoc
	(*plumber.MuxfarmID)(nil),     // 2: plumber.plumber.MuxfarmID
	(*plumber.MediaIn)(nil),       // 3: plumber.plumber.MediaIn
	(*plumber.Status)(nil),        // 4: plumber.state.Status
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*plumber.Media)(nil),         // 6: plumber.plumber.Media
}
var file_fixtures_fixtures_proto_depIdxs = []int32{
	2, // 0: fixtures.ingestDoc.muxfarmid:type_name -> plumber.plumber.MuxfarmID
	3, // 1: fixtures.ingestDoc.mediain:type_name -> plumber.plumber.MediaIn
	4, // 2: fixtures.ingestDoc.status:type_name -> plumber.state.Status
	5, // 3: fixtures.ingestDoc.timestamp:type_name -> google.protobuf.Timestamp
	2, // 4: fixtures.ingestBatchDoc.muxfarmid:type_name -> plumber.plumber.MuxfarmID
	2, // 5: fixtures.ingestBatchDoc.ingestid:type_name -> plumber.plumber.MuxfarmID
	6, // 6: fixtures.ingestBatchDoc.media:type_name -> plumber.plumber.Media
	4, // 7: fixtures.ingestBatchDoc.status:type_name -> plumber.state.Status
	5, // 8: fixtures.ingestBatchDoc.timestamp:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_fixtures_fixtures_proto_init() }
func file_fixtures_fixtures_proto_init() {
	if File_fixtures_fixtures_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fixtures_fixtures_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestDoc); i {
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
		file_fixtures_fixtures_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestBatchDoc); i {
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
			RawDescriptor: file_fixtures_fixtures_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fixtures_fixtures_proto_goTypes,
		DependencyIndexes: file_fixtures_fixtures_proto_depIdxs,
		MessageInfos:      file_fixtures_fixtures_proto_msgTypes,
	}.Build()
	File_fixtures_fixtures_proto = out.File
	file_fixtures_fixtures_proto_rawDesc = nil
	file_fixtures_fixtures_proto_goTypes = nil
	file_fixtures_fixtures_proto_depIdxs = nil
}
