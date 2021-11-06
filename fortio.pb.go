// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/client/transform/fortio.proto

package main

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This proto represents that output format that Fortio expects when converted to JSON.
// Nighthawk can fill in this proto, and then unmarshal to the Fortio-compatible JSON.
// Therefore, this proto may not follow conventions. aip.dev/not-precedent
type FortioResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// From https://github.com/fortio/fortio/wiki/FAQ:
	// If using the command line, don't forget to use -a to auto save the json results in the current
	// data directory. And to use labels -labels "x y z" so the file name and the json include some
	// description of your run. You can later run fortio report to browse
	// graphically your results.
	// The Nighthawk equivalent of this field is structured as a repeated string, which will be
	// concatenated into this field, using ' ' as a separator.
	Labels string `protobuf:"bytes,1,opt,name=Labels,proto3" json:"Labels,omitempty"`
	// Start time of the load test execution.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	// Configured qps
	RequestedQPS uint32 `protobuf:"varint,3,opt,name=RequestedQPS,proto3" json:"RequestedQPS,omitempty"`
	// Configured duration
	RequestedDuration *durationpb.Duration `protobuf:"bytes,4,opt,name=RequestedDuration,proto3" json:"RequestedDuration,omitempty"`
	// Effective qps
	ActualQPS float64 `protobuf:"fixed64,5,opt,name=ActualQPS,proto3" json:"ActualQPS,omitempty"`
	// Effective duration
	ActualDuration float64 `protobuf:"fixed64,6,opt,name=ActualDuration,proto3" json:"ActualDuration,omitempty"`
	// The configured number of used connections
	NumThreads uint32 `protobuf:"varint,7,opt,name=NumThreads,proto3" json:"NumThreads,omitempty"`
	// Latency histogram
	DurationHistogram *DurationHistogram `protobuf:"bytes,8,opt,name=DurationHistogram,proto3" json:"DurationHistogram,omitempty"`
	// Fortio tracks per-return-code. We group by 2xx,3xx, etc.
	RetCodes map[string]uint64 `protobuf:"bytes,11,rep,name=RetCodes,proto3" json:"RetCodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Effective URI.
	URL string `protobuf:"bytes,12,opt,name=URL,proto3" json:"URL,omitempty"`
	// (Unstructured) version. We serialize a representation of the Nighthawk version into this field.
	Version string `protobuf:"bytes,13,opt,name=Version,proto3" json:"Version,omitempty"`
	// Was jitter used Y/N.
	Jitter bool `protobuf:"varint,14,opt,name=Jitter,proto3" json:"Jitter,omitempty"`
	// Run type. Can be "HTTP", "GRPC", and an empty string has also been observed.
	// We only set HTTP for now.
	RunType string `protobuf:"bytes,15,opt,name=RunType,proto3" json:"RunType,omitempty"`
	// Sizes and HeaderSizes seems to mean different things in Fortio depending on if the go stdclient
	// or DIY fastclient was used. We populate this field with response body sizes, and HeaderSizes
	// with header sizes. Both are tracked in bytes. Note: we don't use full histograms here, but
	// rather rely on StreamingStatistic. This means we don't populate percentiles, and only populate
	// min/mean/max/etc with respect to DurationHistogram.
	Sizes *DurationHistogram `protobuf:"bytes,16,opt,name=Sizes,proto3" json:"Sizes,omitempty"`
	// See the 'Sizes' field for detals.
	HeaderSizes *DurationHistogram `protobuf:"bytes,17,opt,name=HeaderSizes,proto3" json:"HeaderSizes,omitempty"`
	// Bytes sent should reflect the bytes sent after protocol serialization, compression, and
	// encryption are applied (i.e., the size of the buffers passed to the send() system call or its
	// equivalent).
	BytesSent uint64 `protobuf:"varint,18,opt,name=BytesSent,proto3" json:"BytesSent,omitempty"`
	// Bytes received should reflect the bytes read before decryption, de-compression, and protocol
	// de-serialization are applied (i.e., the size of bytes read from the recv() system call or its
	// equivalent).
	BytesReceived uint64 `protobuf:"varint,19,opt,name=BytesReceived,proto3" json:"BytesReceived,omitempty"`
}

func (x *FortioResult) Reset() {
	*x = FortioResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_client_transform_fortio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FortioResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FortioResult) ProtoMessage() {}

func (x *FortioResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_client_transform_fortio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FortioResult.ProtoReflect.Descriptor instead.
func (*FortioResult) Descriptor() ([]byte, []int) {
	return file_api_client_transform_fortio_proto_rawDescGZIP(), []int{0}
}

func (x *FortioResult) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

func (x *FortioResult) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *FortioResult) GetRequestedQPS() uint32 {
	if x != nil {
		return x.RequestedQPS
	}
	return 0
}

func (x *FortioResult) GetRequestedDuration() *durationpb.Duration {
	if x != nil {
		return x.RequestedDuration
	}
	return nil
}

func (x *FortioResult) GetActualQPS() float64 {
	if x != nil {
		return x.ActualQPS
	}
	return 0
}

func (x *FortioResult) GetActualDuration() float64 {
	if x != nil {
		return x.ActualDuration
	}
	return 0
}

func (x *FortioResult) GetNumThreads() uint32 {
	if x != nil {
		return x.NumThreads
	}
	return 0
}

func (x *FortioResult) GetDurationHistogram() *DurationHistogram {
	if x != nil {
		return x.DurationHistogram
	}
	return nil
}

func (x *FortioResult) GetRetCodes() map[string]uint64 {
	if x != nil {
		return x.RetCodes
	}
	return nil
}

func (x *FortioResult) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *FortioResult) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *FortioResult) GetJitter() bool {
	if x != nil {
		return x.Jitter
	}
	return false
}

func (x *FortioResult) GetRunType() string {
	if x != nil {
		return x.RunType
	}
	return ""
}

func (x *FortioResult) GetSizes() *DurationHistogram {
	if x != nil {
		return x.Sizes
	}
	return nil
}

func (x *FortioResult) GetHeaderSizes() *DurationHistogram {
	if x != nil {
		return x.HeaderSizes
	}
	return nil
}

func (x *FortioResult) GetBytesSent() uint64 {
	if x != nil {
		return x.BytesSent
	}
	return 0
}

func (x *FortioResult) GetBytesReceived() uint64 {
	if x != nil {
		return x.BytesReceived
	}
	return 0
}

type DurationHistogram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       uint64              `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Data        []*DataEntry        `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
	Min         float64             `protobuf:"fixed64,3,opt,name=Min,proto3" json:"Min,omitempty"`
	Max         float64             `protobuf:"fixed64,4,opt,name=Max,proto3" json:"Max,omitempty"`
	Sum         float64             `protobuf:"fixed64,5,opt,name=Sum,proto3" json:"Sum,omitempty"`
	Avg         float64             `protobuf:"fixed64,6,opt,name=Avg,proto3" json:"Avg,omitempty"`
	StdDev      float64             `protobuf:"fixed64,7,opt,name=StdDev,proto3" json:"StdDev,omitempty"`
	Percentiles []*FortioPercentile `protobuf:"bytes,8,rep,name=Percentiles,proto3" json:"Percentiles,omitempty"`
}

func (x *DurationHistogram) Reset() {
	*x = DurationHistogram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_client_transform_fortio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationHistogram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationHistogram) ProtoMessage() {}

func (x *DurationHistogram) ProtoReflect() protoreflect.Message {
	mi := &file_api_client_transform_fortio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationHistogram.ProtoReflect.Descriptor instead.
func (*DurationHistogram) Descriptor() ([]byte, []int) {
	return file_api_client_transform_fortio_proto_rawDescGZIP(), []int{1}
}

func (x *DurationHistogram) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *DurationHistogram) GetData() []*DataEntry {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DurationHistogram) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *DurationHistogram) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *DurationHistogram) GetSum() float64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *DurationHistogram) GetAvg() float64 {
	if x != nil {
		return x.Avg
	}
	return 0
}

func (x *DurationHistogram) GetStdDev() float64 {
	if x != nil {
		return x.StdDev
	}
	return 0
}

func (x *DurationHistogram) GetPercentiles() []*FortioPercentile {
	if x != nil {
		return x.Percentiles
	}
	return nil
}

type DataEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start   float64 `protobuf:"fixed64,1,opt,name=Start,proto3" json:"Start,omitempty"`
	End     float64 `protobuf:"fixed64,2,opt,name=End,proto3" json:"End,omitempty"`
	Percent float64 `protobuf:"fixed64,3,opt,name=Percent,proto3" json:"Percent,omitempty"`
	Count   uint64  `protobuf:"varint,4,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *DataEntry) Reset() {
	*x = DataEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_client_transform_fortio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataEntry) ProtoMessage() {}

func (x *DataEntry) ProtoReflect() protoreflect.Message {
	mi := &file_api_client_transform_fortio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataEntry.ProtoReflect.Descriptor instead.
func (*DataEntry) Descriptor() ([]byte, []int) {
	return file_api_client_transform_fortio_proto_rawDescGZIP(), []int{2}
}

func (x *DataEntry) GetStart() float64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *DataEntry) GetEnd() float64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *DataEntry) GetPercent() float64 {
	if x != nil {
		return x.Percent
	}
	return 0
}

func (x *DataEntry) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type FortioPercentile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Percentile float64 `protobuf:"fixed64,1,opt,name=Percentile,proto3" json:"Percentile,omitempty"`
	Value      float64 `protobuf:"fixed64,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *FortioPercentile) Reset() {
	*x = FortioPercentile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_client_transform_fortio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FortioPercentile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FortioPercentile) ProtoMessage() {}

func (x *FortioPercentile) ProtoReflect() protoreflect.Message {
	mi := &file_api_client_transform_fortio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FortioPercentile.ProtoReflect.Descriptor instead.
func (*FortioPercentile) Descriptor() ([]byte, []int) {
	return file_api_client_transform_fortio_proto_rawDescGZIP(), []int{3}
}

func (x *FortioPercentile) GetPercentile() float64 {
	if x != nil {
		return x.Percentile
	}
	return 0
}

func (x *FortioPercentile) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_api_client_transform_fortio_proto protoreflect.FileDescriptor

var file_api_client_transform_fortio_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x66, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x68, 0x61, 0x77, 0x6b, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb6, 0x06, 0x0a, 0x0c, 0x46, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x51,
	0x50, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x51, 0x50, 0x53, 0x12, 0x47, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x51, 0x50, 0x53, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x51, 0x50, 0x53, 0x12, 0x26, 0x0a,
	0x0e, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x75, 0x6d, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x4e, 0x75, 0x6d, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x51, 0x0a, 0x11, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x68, 0x61, 0x77, 0x6b, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x11, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x48, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x69, 0x67,
	0x68, 0x74, 0x68, 0x61, 0x77, 0x6b, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6f,
	0x72, 0x74, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x4a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x4a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x39, 0x0a, 0x05, 0x53, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x68, 0x61, 0x77, 0x6b, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x53, 0x69, 0x7a, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x68, 0x61, 0x77, 0x6b, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x69, 0x7a,
	0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x42, 0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x1a, 0x3b, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x3a, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x85, 0x02, 0x0a, 0x11, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x14,
	0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x68, 0x61, 0x77, 0x6b, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x4d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x61, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x4d, 0x61, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x75, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x41,
	0x76, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x41, 0x76, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x64, 0x44, 0x65, 0x76, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x53,
	0x74, 0x64, 0x44, 0x65, 0x76, 0x12, 0x44, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x69, 0x67,
	0x68, 0x74, 0x68, 0x61, 0x77, 0x6b, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6f,
	0x72, 0x74, 0x69, 0x6f, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x52, 0x0b,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x3a, 0x03, 0xf8, 0x42, 0x01,
	0x22, 0x68, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x45, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x4d, 0x0a, 0x10, 0x46, 0x6f,
	0x72, 0x74, 0x69, 0x6f, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_client_transform_fortio_proto_rawDescOnce sync.Once
	file_api_client_transform_fortio_proto_rawDescData = file_api_client_transform_fortio_proto_rawDesc
)

func file_api_client_transform_fortio_proto_rawDescGZIP() []byte {
	file_api_client_transform_fortio_proto_rawDescOnce.Do(func() {
		file_api_client_transform_fortio_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_client_transform_fortio_proto_rawDescData)
	})
	return file_api_client_transform_fortio_proto_rawDescData
}

var file_api_client_transform_fortio_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_client_transform_fortio_proto_goTypes = []interface{}{
	(*FortioResult)(nil),          // 0: nighthawk.client.FortioResult
	(*DurationHistogram)(nil),     // 1: nighthawk.client.DurationHistogram
	(*DataEntry)(nil),             // 2: nighthawk.client.DataEntry
	(*FortioPercentile)(nil),      // 3: nighthawk.client.FortioPercentile
	nil,                           // 4: nighthawk.client.FortioResult.RetCodesEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 6: google.protobuf.Duration
}
var file_api_client_transform_fortio_proto_depIdxs = []int32{
	5, // 0: nighthawk.client.FortioResult.StartTime:type_name -> google.protobuf.Timestamp
	6, // 1: nighthawk.client.FortioResult.RequestedDuration:type_name -> google.protobuf.Duration
	1, // 2: nighthawk.client.FortioResult.DurationHistogram:type_name -> nighthawk.client.DurationHistogram
	4, // 3: nighthawk.client.FortioResult.RetCodes:type_name -> nighthawk.client.FortioResult.RetCodesEntry
	1, // 4: nighthawk.client.FortioResult.Sizes:type_name -> nighthawk.client.DurationHistogram
	1, // 5: nighthawk.client.FortioResult.HeaderSizes:type_name -> nighthawk.client.DurationHistogram
	2, // 6: nighthawk.client.DurationHistogram.Data:type_name -> nighthawk.client.DataEntry
	3, // 7: nighthawk.client.DurationHistogram.Percentiles:type_name -> nighthawk.client.FortioPercentile
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_api_client_transform_fortio_proto_init() }
func file_api_client_transform_fortio_proto_init() {
	if File_api_client_transform_fortio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_client_transform_fortio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FortioResult); i {
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
		file_api_client_transform_fortio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationHistogram); i {
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
		file_api_client_transform_fortio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataEntry); i {
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
		file_api_client_transform_fortio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FortioPercentile); i {
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
			RawDescriptor: file_api_client_transform_fortio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_client_transform_fortio_proto_goTypes,
		DependencyIndexes: file_api_client_transform_fortio_proto_depIdxs,
		MessageInfos:      file_api_client_transform_fortio_proto_msgTypes,
	}.Build()
	File_api_client_transform_fortio_proto = out.File
	file_api_client_transform_fortio_proto_rawDesc = nil
	file_api_client_transform_fortio_proto_goTypes = nil
	file_api_client_transform_fortio_proto_depIdxs = nil
}