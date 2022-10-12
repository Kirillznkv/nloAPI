// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.21.5
// source: nlo.proto

package pkg

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_nlo_proto_msgTypes[0]
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
	return file_nlo_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string                 `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Frequency float64                `protobuf:"fixed64,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_nlo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_nlo_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Response) GetFrequency() float64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *Response) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_nlo_proto protoreflect.FileDescriptor

var file_nlo_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x81, 0x01, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x66, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x32, 0x2c, 0x0a, 0x03, 0x4e, 0x6c, 0x6f, 0x12, 0x25, 0x0a, 0x02, 0x44, 0x6f, 0x12, 0x0c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nlo_proto_rawDescOnce sync.Once
	file_nlo_proto_rawDescData = file_nlo_proto_rawDesc
)

func file_nlo_proto_rawDescGZIP() []byte {
	file_nlo_proto_rawDescOnce.Do(func() {
		file_nlo_proto_rawDescData = protoimpl.X.CompressGZIP(file_nlo_proto_rawDescData)
	})
	return file_nlo_proto_rawDescData
}

var file_nlo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nlo_proto_goTypes = []interface{}{
	(*Request)(nil),               // 0: api.Request
	(*Response)(nil),              // 1: api.Response
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_nlo_proto_depIdxs = []int32{
	2, // 0: api.Response.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: api.Nlo.Do:input_type -> api.Request
	1, // 2: api.Nlo.Do:output_type -> api.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nlo_proto_init() }
func file_nlo_proto_init() {
	if File_nlo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nlo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_nlo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_nlo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nlo_proto_goTypes,
		DependencyIndexes: file_nlo_proto_depIdxs,
		MessageInfos:      file_nlo_proto_msgTypes,
	}.Build()
	File_nlo_proto = out.File
	file_nlo_proto_rawDesc = nil
	file_nlo_proto_goTypes = nil
	file_nlo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NloClient is the client API for Nlo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NloClient interface {
	Do(ctx context.Context, in *Request, opts ...grpc.CallOption) (Nlo_DoClient, error)
}

type nloClient struct {
	cc grpc.ClientConnInterface
}

func NewNloClient(cc grpc.ClientConnInterface) NloClient {
	return &nloClient{cc}
}

func (c *nloClient) Do(ctx context.Context, in *Request, opts ...grpc.CallOption) (Nlo_DoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Nlo_serviceDesc.Streams[0], "/api.Nlo/Do", opts...)
	if err != nil {
		return nil, err
	}
	x := &nloDoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Nlo_DoClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type nloDoClient struct {
	grpc.ClientStream
}

func (x *nloDoClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NloServer is the server API for Nlo service.
type NloServer interface {
	Do(*Request, Nlo_DoServer) error
}

// UnimplementedNloServer can be embedded to have forward compatible implementations.
type UnimplementedNloServer struct {
}

func (*UnimplementedNloServer) Do(*Request, Nlo_DoServer) error {
	return status.Errorf(codes.Unimplemented, "method Do not implemented")
}

func RegisterNloServer(s *grpc.Server, srv NloServer) {
	s.RegisterService(&_Nlo_serviceDesc, srv)
}

func _Nlo_Do_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NloServer).Do(m, &nloDoServer{stream})
}

type Nlo_DoServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type nloDoServer struct {
	grpc.ServerStream
}

func (x *nloDoServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Nlo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Nlo",
	HandlerType: (*NloServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Do",
			Handler:       _Nlo_Do_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "nlo.proto",
}