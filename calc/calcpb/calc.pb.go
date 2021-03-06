// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.12.3
// source: calc/calcpb/calc.proto

package calcpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SumMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int64 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int64 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *SumMessage) Reset() {
	*x = SumMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calcpb_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumMessage) ProtoMessage() {}

func (x *SumMessage) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calcpb_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumMessage.ProtoReflect.Descriptor instead.
func (*SumMessage) Descriptor() ([]byte, []int) {
	return file_calc_calcpb_calc_proto_rawDescGZIP(), []int{0}
}

func (x *SumMessage) GetNum1() int64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *SumMessage) GetNum2() int64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type PrimeNumMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeNum int64 `protobuf:"varint,1,opt,name=primeNum,proto3" json:"primeNum,omitempty"`
}

func (x *PrimeNumMessage) Reset() {
	*x = PrimeNumMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calcpb_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumMessage) ProtoMessage() {}

func (x *PrimeNumMessage) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calcpb_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumMessage.ProtoReflect.Descriptor instead.
func (*PrimeNumMessage) Descriptor() ([]byte, []int) {
	return file_calc_calcpb_calc_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeNumMessage) GetPrimeNum() int64 {
	if x != nil {
		return x.PrimeNum
	}
	return 0
}

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SumMessage *SumMessage `protobuf:"bytes,1,opt,name=sumMessage,proto3" json:"sumMessage,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calcpb_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calcpb_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calc_calcpb_calc_proto_rawDescGZIP(), []int{2}
}

func (x *SumRequest) GetSumMessage() *SumMessage {
	if x != nil {
		return x.SumMessage
	}
	return nil
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calcpb_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calcpb_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calc_calcpb_calc_proto_rawDescGZIP(), []int{3}
}

func (x *SumResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type PrimeNumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeNumMessage *PrimeNumMessage `protobuf:"bytes,1,opt,name=primeNumMessage,proto3" json:"primeNumMessage,omitempty"`
}

func (x *PrimeNumRequest) Reset() {
	*x = PrimeNumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calcpb_calc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumRequest) ProtoMessage() {}

func (x *PrimeNumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calcpb_calc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumRequest.ProtoReflect.Descriptor instead.
func (*PrimeNumRequest) Descriptor() ([]byte, []int) {
	return file_calc_calcpb_calc_proto_rawDescGZIP(), []int{4}
}

func (x *PrimeNumRequest) GetPrimeNumMessage() *PrimeNumMessage {
	if x != nil {
		return x.PrimeNumMessage
	}
	return nil
}

type PrimeNumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeNumResponse) Reset() {
	*x = PrimeNumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calcpb_calc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumResponse) ProtoMessage() {}

func (x *PrimeNumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calcpb_calc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumResponse.ProtoReflect.Descriptor instead.
func (*PrimeNumResponse) Descriptor() ([]byte, []int) {
	return file_calc_calcpb_calc_proto_rawDescGZIP(), []int{5}
}

func (x *PrimeNumResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calc_calcpb_calc_proto protoreflect.FileDescriptor

var file_calc_calcpb_calc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x61, 0x6c, 0x63, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x2f, 0x63, 0x61,
	0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x6c, 0x63, 0x22, 0x34,
	0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x6e, 0x75, 0x6d, 0x32, 0x22, 0x2d, 0x0a, 0x0f, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x4e, 0x75, 0x6d, 0x22, 0x3e, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x0a, 0x73, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x53, 0x75, 0x6d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x73, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x52, 0x0a, 0x0f, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a,
	0x0f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a,
	0x0a, 0x10, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x94, 0x01, 0x0a, 0x11, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x30, 0x0a, 0x07, 0x43, 0x61, 0x6c, 0x63, 0x53, 0x75, 0x6d, 0x12, 0x10, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x18, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x63, 0x61, 0x6c, 0x63, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calc_calcpb_calc_proto_rawDescOnce sync.Once
	file_calc_calcpb_calc_proto_rawDescData = file_calc_calcpb_calc_proto_rawDesc
)

func file_calc_calcpb_calc_proto_rawDescGZIP() []byte {
	file_calc_calcpb_calc_proto_rawDescOnce.Do(func() {
		file_calc_calcpb_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calc_calcpb_calc_proto_rawDescData)
	})
	return file_calc_calcpb_calc_proto_rawDescData
}

var file_calc_calcpb_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_calc_calcpb_calc_proto_goTypes = []interface{}{
	(*SumMessage)(nil),       // 0: calc.SumMessage
	(*PrimeNumMessage)(nil),  // 1: calc.PrimeNumMessage
	(*SumRequest)(nil),       // 2: calc.SumRequest
	(*SumResponse)(nil),      // 3: calc.SumResponse
	(*PrimeNumRequest)(nil),  // 4: calc.PrimeNumRequest
	(*PrimeNumResponse)(nil), // 5: calc.PrimeNumResponse
}
var file_calc_calcpb_calc_proto_depIdxs = []int32{
	0, // 0: calc.SumRequest.sumMessage:type_name -> calc.SumMessage
	1, // 1: calc.PrimeNumRequest.primeNumMessage:type_name -> calc.PrimeNumMessage
	2, // 2: calc.CalculatorService.CalcSum:input_type -> calc.SumRequest
	4, // 3: calc.CalculatorService.PrimeNumberDecomposition:input_type -> calc.PrimeNumRequest
	3, // 4: calc.CalculatorService.CalcSum:output_type -> calc.SumResponse
	5, // 5: calc.CalculatorService.PrimeNumberDecomposition:output_type -> calc.PrimeNumResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_calc_calcpb_calc_proto_init() }
func file_calc_calcpb_calc_proto_init() {
	if File_calc_calcpb_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calc_calcpb_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumMessage); i {
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
		file_calc_calcpb_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumMessage); i {
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
		file_calc_calcpb_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_calc_calcpb_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_calc_calcpb_calc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumRequest); i {
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
		file_calc_calcpb_calc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumResponse); i {
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
			RawDescriptor: file_calc_calcpb_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calc_calcpb_calc_proto_goTypes,
		DependencyIndexes: file_calc_calcpb_calc_proto_depIdxs,
		MessageInfos:      file_calc_calcpb_calc_proto_msgTypes,
	}.Build()
	File_calc_calcpb_calc_proto = out.File
	file_calc_calcpb_calc_proto_rawDesc = nil
	file_calc_calcpb_calc_proto_goTypes = nil
	file_calc_calcpb_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	CalcSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) CalcSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calc.CalculatorService/CalcSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calc.CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*PrimeNumResponse, error) {
	m := new(PrimeNumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	CalcSum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumberDecomposition(*PrimeNumRequest, CalculatorService_PrimeNumberDecompositionServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) CalcSum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalcSum not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeNumberDecomposition(*PrimeNumRequest, CalculatorService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_CalcSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).CalcSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalculatorService/CalcSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).CalcSum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *PrimeNumResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalcSum",
			Handler:    _CalculatorService_CalcSum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calc/calcpb/calc.proto",
}
