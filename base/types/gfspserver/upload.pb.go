// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: base/types/gfspserver/upload.proto

package gfspserver

import (
	context "context"
	fmt "fmt"
	gfsperrors "github.com/bnb-chain/greenfield-storage-provider/base/types/gfsperrors"
	gfsptask "github.com/bnb-chain/greenfield-storage-provider/base/types/gfsptask"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GfSpUploadObjectRequest struct {
	UploadObjectTask *gfsptask.GfSpUploadObjectTask `protobuf:"bytes,1,opt,name=upload_object_task,json=uploadObjectTask,proto3" json:"upload_object_task,omitempty"`
	Payload          []byte                         `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *GfSpUploadObjectRequest) Reset()         { *m = GfSpUploadObjectRequest{} }
func (m *GfSpUploadObjectRequest) String() string { return proto.CompactTextString(m) }
func (*GfSpUploadObjectRequest) ProtoMessage()    {}
func (*GfSpUploadObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_552ccd21b70a14be, []int{0}
}
func (m *GfSpUploadObjectRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GfSpUploadObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GfSpUploadObjectRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GfSpUploadObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GfSpUploadObjectRequest.Merge(m, src)
}
func (m *GfSpUploadObjectRequest) XXX_Size() int {
	return m.Size()
}
func (m *GfSpUploadObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GfSpUploadObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GfSpUploadObjectRequest proto.InternalMessageInfo

func (m *GfSpUploadObjectRequest) GetUploadObjectTask() *gfsptask.GfSpUploadObjectTask {
	if m != nil {
		return m.UploadObjectTask
	}
	return nil
}

func (m *GfSpUploadObjectRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type GfSpUploadObjectResponse struct {
	Err *gfsperrors.GfSpError `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (m *GfSpUploadObjectResponse) Reset()         { *m = GfSpUploadObjectResponse{} }
func (m *GfSpUploadObjectResponse) String() string { return proto.CompactTextString(m) }
func (*GfSpUploadObjectResponse) ProtoMessage()    {}
func (*GfSpUploadObjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_552ccd21b70a14be, []int{1}
}
func (m *GfSpUploadObjectResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GfSpUploadObjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GfSpUploadObjectResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GfSpUploadObjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GfSpUploadObjectResponse.Merge(m, src)
}
func (m *GfSpUploadObjectResponse) XXX_Size() int {
	return m.Size()
}
func (m *GfSpUploadObjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GfSpUploadObjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GfSpUploadObjectResponse proto.InternalMessageInfo

func (m *GfSpUploadObjectResponse) GetErr() *gfsperrors.GfSpError {
	if m != nil {
		return m.Err
	}
	return nil
}

type GfSpResumableUploadObjectRequest struct {
	ResumableUploadObjectTask *gfsptask.GfSpResumableUploadObjectTask `protobuf:"bytes,1,opt,name=resumable_upload_object_task,json=resumableUploadObjectTask,proto3" json:"resumable_upload_object_task,omitempty"`
	Payload                   []byte                                  `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *GfSpResumableUploadObjectRequest) Reset()         { *m = GfSpResumableUploadObjectRequest{} }
func (m *GfSpResumableUploadObjectRequest) String() string { return proto.CompactTextString(m) }
func (*GfSpResumableUploadObjectRequest) ProtoMessage()    {}
func (*GfSpResumableUploadObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_552ccd21b70a14be, []int{2}
}
func (m *GfSpResumableUploadObjectRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GfSpResumableUploadObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GfSpResumableUploadObjectRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GfSpResumableUploadObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GfSpResumableUploadObjectRequest.Merge(m, src)
}
func (m *GfSpResumableUploadObjectRequest) XXX_Size() int {
	return m.Size()
}
func (m *GfSpResumableUploadObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GfSpResumableUploadObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GfSpResumableUploadObjectRequest proto.InternalMessageInfo

func (m *GfSpResumableUploadObjectRequest) GetResumableUploadObjectTask() *gfsptask.GfSpResumableUploadObjectTask {
	if m != nil {
		return m.ResumableUploadObjectTask
	}
	return nil
}

func (m *GfSpResumableUploadObjectRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type GfSpResumableUploadObjectResponse struct {
	Err *gfsperrors.GfSpError `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (m *GfSpResumableUploadObjectResponse) Reset()         { *m = GfSpResumableUploadObjectResponse{} }
func (m *GfSpResumableUploadObjectResponse) String() string { return proto.CompactTextString(m) }
func (*GfSpResumableUploadObjectResponse) ProtoMessage()    {}
func (*GfSpResumableUploadObjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_552ccd21b70a14be, []int{3}
}
func (m *GfSpResumableUploadObjectResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GfSpResumableUploadObjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GfSpResumableUploadObjectResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GfSpResumableUploadObjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GfSpResumableUploadObjectResponse.Merge(m, src)
}
func (m *GfSpResumableUploadObjectResponse) XXX_Size() int {
	return m.Size()
}
func (m *GfSpResumableUploadObjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GfSpResumableUploadObjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GfSpResumableUploadObjectResponse proto.InternalMessageInfo

func (m *GfSpResumableUploadObjectResponse) GetErr() *gfsperrors.GfSpError {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*GfSpUploadObjectRequest)(nil), "base.types.gfspserver.GfSpUploadObjectRequest")
	proto.RegisterType((*GfSpUploadObjectResponse)(nil), "base.types.gfspserver.GfSpUploadObjectResponse")
	proto.RegisterType((*GfSpResumableUploadObjectRequest)(nil), "base.types.gfspserver.GfSpResumableUploadObjectRequest")
	proto.RegisterType((*GfSpResumableUploadObjectResponse)(nil), "base.types.gfspserver.GfSpResumableUploadObjectResponse")
}

func init() {
	proto.RegisterFile("base/types/gfspserver/upload.proto", fileDescriptor_552ccd21b70a14be)
}

var fileDescriptor_552ccd21b70a14be = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x8e, 0x8b, 0x04, 0x92, 0xe1, 0x50, 0x2c, 0x21, 0xd2, 0x08, 0x45, 0x69, 0x4e, 0xe5, 0x50,
	0x5b, 0x0a, 0x07, 0x38, 0x23, 0x01, 0x37, 0x90, 0x52, 0x50, 0x25, 0x2e, 0xc5, 0x49, 0xa7, 0x69,
	0x68, 0x1b, 0x1b, 0x3b, 0x29, 0xea, 0x3b, 0x70, 0x40, 0xbc, 0x01, 0x8f, 0xc0, 0x5b, 0x70, 0xec,
	0x71, 0x8f, 0xab, 0xf6, 0x45, 0x56, 0x71, 0xb6, 0xaa, 0x94, 0x4d, 0xa2, 0xdd, 0xbd, 0xe4, 0x67,
	0xe6, 0x9b, 0x99, 0xef, 0xfb, 0xec, 0xc1, 0x7e, 0xc4, 0x35, 0xb0, 0x7c, 0x27, 0x41, 0xb3, 0x64,
	0xa1, 0xa5, 0x06, 0xb5, 0x05, 0xc5, 0x0a, 0xb9, 0x16, 0x7c, 0x4e, 0xa5, 0x12, 0xb9, 0x20, 0xcf,
	0x4a, 0x0c, 0x35, 0x18, 0x7a, 0xc6, 0x38, 0xc3, 0x5a, 0x29, 0x28, 0x25, 0x94, 0x66, 0xe6, 0x55,
	0x55, 0x3a, 0x6e, 0x0d, 0x92, 0x73, 0xbd, 0x62, 0xe5, 0xa3, 0xca, 0xfb, 0xbf, 0x10, 0x7e, 0xfe,
	0x61, 0x31, 0x91, 0x5f, 0xcc, 0xb8, 0x4f, 0xd1, 0x77, 0x88, 0xf3, 0x10, 0x7e, 0x14, 0xa0, 0x73,
	0x32, 0xc5, 0xa4, 0x62, 0x31, 0x13, 0x26, 0x3e, 0x2b, 0xeb, 0x6c, 0xe4, 0xa1, 0xd1, 0xe3, 0xe0,
	0x25, 0xad, 0x51, 0x32, 0x3d, 0xeb, 0x9d, 0x3e, 0x73, 0xbd, 0x0a, 0xfb, 0x45, 0x2d, 0x42, 0x6c,
	0xfc, 0x48, 0xf2, 0x5d, 0x19, 0xb4, 0x7b, 0x1e, 0x1a, 0x3d, 0x09, 0x4f, 0xbf, 0xfe, 0x47, 0x6c,
	0xdf, 0x64, 0xa3, 0xa5, 0xc8, 0x34, 0x90, 0x00, 0x3f, 0x00, 0xa5, 0xae, 0xe7, 0x7b, 0xf5, 0xf9,
	0x95, 0x76, 0xc3, 0xe0, 0x5d, 0xf9, 0x19, 0x96, 0x60, 0xff, 0x1f, 0xc2, 0x5e, 0x19, 0x0a, 0x41,
	0x17, 0x1b, 0x1e, 0xad, 0xa1, 0x49, 0xa7, 0xc6, 0x2f, 0xd4, 0x29, 0x3f, 0x6b, 0x55, 0x1c, 0xb4,
	0x2a, 0x6e, 0x6c, 0x6e, 0xa4, 0x0f, 0x54, 0x5b, 0xaa, 0xc3, 0x83, 0x29, 0x1e, 0x76, 0x50, 0xbe,
	0xbf, 0x19, 0xc1, 0xdf, 0x1e, 0x7e, 0x7a, 0x76, 0x77, 0x02, 0x6a, 0x9b, 0xc6, 0x40, 0x7e, 0xe2,
	0x7e, 0xdd, 0x72, 0x42, 0x69, 0xe3, 0x85, 0xa3, 0x2d, 0x37, 0xc5, 0x61, 0xb7, 0xc6, 0x57, 0xf4,
	0x7d, 0x6b, 0x84, 0xc8, 0x1f, 0x84, 0x07, 0xad, 0x42, 0xc9, 0xeb, 0x8e, 0x96, 0x5d, 0xa7, 0xe9,
	0xbc, 0xb9, 0x7b, 0xe1, 0x99, 0xd4, 0xdb, 0x6f, 0xff, 0x0f, 0x2e, 0xda, 0x1f, 0x5c, 0x74, 0x79,
	0x70, 0xd1, 0xef, 0xa3, 0x6b, 0xed, 0x8f, 0xae, 0x75, 0x71, 0x74, 0xad, 0xaf, 0xef, 0x93, 0x34,
	0x5f, 0x16, 0x11, 0x8d, 0xc5, 0x86, 0x45, 0x59, 0x34, 0x8e, 0x97, 0x3c, 0xcd, 0x58, 0xa2, 0x00,
	0xb2, 0x45, 0x0a, 0xeb, 0xf9, 0x58, 0xe7, 0x42, 0xf1, 0x04, 0xc6, 0x52, 0x89, 0x6d, 0x3a, 0x07,
	0xc5, 0x1a, 0x17, 0x3b, 0x7a, 0x68, 0x16, 0xef, 0xd5, 0x55, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0x8e, 0xb5, 0xb0, 0xf8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GfSpUploadServiceClient is the client API for GfSpUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GfSpUploadServiceClient interface {
	GfSpUploadObject(ctx context.Context, opts ...grpc.CallOption) (GfSpUploadService_GfSpUploadObjectClient, error)
	// TODO(chris): It is recommended to use the segment buffer as a single request instead of the GRPC stream, as the performance
	//  of the GRPC stream is not satisfactory.
	GfSpResumableUploadObject(ctx context.Context, opts ...grpc.CallOption) (GfSpUploadService_GfSpResumableUploadObjectClient, error)
}

type gfSpUploadServiceClient struct {
	cc grpc1.ClientConn
}

func NewGfSpUploadServiceClient(cc grpc1.ClientConn) GfSpUploadServiceClient {
	return &gfSpUploadServiceClient{cc}
}

func (c *gfSpUploadServiceClient) GfSpUploadObject(ctx context.Context, opts ...grpc.CallOption) (GfSpUploadService_GfSpUploadObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GfSpUploadService_serviceDesc.Streams[0], "/base.types.gfspserver.GfSpUploadService/GfSpUploadObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &gfSpUploadServiceGfSpUploadObjectClient{stream}
	return x, nil
}

type GfSpUploadService_GfSpUploadObjectClient interface {
	Send(*GfSpUploadObjectRequest) error
	CloseAndRecv() (*GfSpUploadObjectResponse, error)
	grpc.ClientStream
}

type gfSpUploadServiceGfSpUploadObjectClient struct {
	grpc.ClientStream
}

func (x *gfSpUploadServiceGfSpUploadObjectClient) Send(m *GfSpUploadObjectRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gfSpUploadServiceGfSpUploadObjectClient) CloseAndRecv() (*GfSpUploadObjectResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GfSpUploadObjectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gfSpUploadServiceClient) GfSpResumableUploadObject(ctx context.Context, opts ...grpc.CallOption) (GfSpUploadService_GfSpResumableUploadObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GfSpUploadService_serviceDesc.Streams[1], "/base.types.gfspserver.GfSpUploadService/GfSpResumableUploadObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &gfSpUploadServiceGfSpResumableUploadObjectClient{stream}
	return x, nil
}

type GfSpUploadService_GfSpResumableUploadObjectClient interface {
	Send(*GfSpResumableUploadObjectRequest) error
	CloseAndRecv() (*GfSpResumableUploadObjectResponse, error)
	grpc.ClientStream
}

type gfSpUploadServiceGfSpResumableUploadObjectClient struct {
	grpc.ClientStream
}

func (x *gfSpUploadServiceGfSpResumableUploadObjectClient) Send(m *GfSpResumableUploadObjectRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gfSpUploadServiceGfSpResumableUploadObjectClient) CloseAndRecv() (*GfSpResumableUploadObjectResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GfSpResumableUploadObjectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GfSpUploadServiceServer is the server API for GfSpUploadService service.
type GfSpUploadServiceServer interface {
	GfSpUploadObject(GfSpUploadService_GfSpUploadObjectServer) error
	// TODO(chris): It is recommended to use the segment buffer as a single request instead of the GRPC stream, as the performance
	//  of the GRPC stream is not satisfactory.
	GfSpResumableUploadObject(GfSpUploadService_GfSpResumableUploadObjectServer) error
}

// UnimplementedGfSpUploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGfSpUploadServiceServer struct {
}

func (*UnimplementedGfSpUploadServiceServer) GfSpUploadObject(srv GfSpUploadService_GfSpUploadObjectServer) error {
	return status.Errorf(codes.Unimplemented, "method GfSpUploadObject not implemented")
}
func (*UnimplementedGfSpUploadServiceServer) GfSpResumableUploadObject(srv GfSpUploadService_GfSpResumableUploadObjectServer) error {
	return status.Errorf(codes.Unimplemented, "method GfSpResumableUploadObject not implemented")
}

func RegisterGfSpUploadServiceServer(s grpc1.Server, srv GfSpUploadServiceServer) {
	s.RegisterService(&_GfSpUploadService_serviceDesc, srv)
}

func _GfSpUploadService_GfSpUploadObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GfSpUploadServiceServer).GfSpUploadObject(&gfSpUploadServiceGfSpUploadObjectServer{stream})
}

type GfSpUploadService_GfSpUploadObjectServer interface {
	SendAndClose(*GfSpUploadObjectResponse) error
	Recv() (*GfSpUploadObjectRequest, error)
	grpc.ServerStream
}

type gfSpUploadServiceGfSpUploadObjectServer struct {
	grpc.ServerStream
}

func (x *gfSpUploadServiceGfSpUploadObjectServer) SendAndClose(m *GfSpUploadObjectResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gfSpUploadServiceGfSpUploadObjectServer) Recv() (*GfSpUploadObjectRequest, error) {
	m := new(GfSpUploadObjectRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GfSpUploadService_GfSpResumableUploadObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GfSpUploadServiceServer).GfSpResumableUploadObject(&gfSpUploadServiceGfSpResumableUploadObjectServer{stream})
}

type GfSpUploadService_GfSpResumableUploadObjectServer interface {
	SendAndClose(*GfSpResumableUploadObjectResponse) error
	Recv() (*GfSpResumableUploadObjectRequest, error)
	grpc.ServerStream
}

type gfSpUploadServiceGfSpResumableUploadObjectServer struct {
	grpc.ServerStream
}

func (x *gfSpUploadServiceGfSpResumableUploadObjectServer) SendAndClose(m *GfSpResumableUploadObjectResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gfSpUploadServiceGfSpResumableUploadObjectServer) Recv() (*GfSpResumableUploadObjectRequest, error) {
	m := new(GfSpResumableUploadObjectRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GfSpUploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.types.gfspserver.GfSpUploadService",
	HandlerType: (*GfSpUploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GfSpUploadObject",
			Handler:       _GfSpUploadService_GfSpUploadObject_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GfSpResumableUploadObject",
			Handler:       _GfSpUploadService_GfSpResumableUploadObject_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "base/types/gfspserver/upload.proto",
}

func (m *GfSpUploadObjectRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GfSpUploadObjectRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GfSpUploadObjectRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintUpload(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if m.UploadObjectTask != nil {
		{
			size, err := m.UploadObjectTask.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUpload(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GfSpUploadObjectResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GfSpUploadObjectResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GfSpUploadObjectResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Err != nil {
		{
			size, err := m.Err.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUpload(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GfSpResumableUploadObjectRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GfSpResumableUploadObjectRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GfSpResumableUploadObjectRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintUpload(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if m.ResumableUploadObjectTask != nil {
		{
			size, err := m.ResumableUploadObjectTask.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUpload(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GfSpResumableUploadObjectResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GfSpResumableUploadObjectResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GfSpResumableUploadObjectResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Err != nil {
		{
			size, err := m.Err.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUpload(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUpload(dAtA []byte, offset int, v uint64) int {
	offset -= sovUpload(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GfSpUploadObjectRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UploadObjectTask != nil {
		l = m.UploadObjectTask.Size()
		n += 1 + l + sovUpload(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovUpload(uint64(l))
	}
	return n
}

func (m *GfSpUploadObjectResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Err != nil {
		l = m.Err.Size()
		n += 1 + l + sovUpload(uint64(l))
	}
	return n
}

func (m *GfSpResumableUploadObjectRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResumableUploadObjectTask != nil {
		l = m.ResumableUploadObjectTask.Size()
		n += 1 + l + sovUpload(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovUpload(uint64(l))
	}
	return n
}

func (m *GfSpResumableUploadObjectResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Err != nil {
		l = m.Err.Size()
		n += 1 + l + sovUpload(uint64(l))
	}
	return n
}

func sovUpload(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUpload(x uint64) (n int) {
	return sovUpload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GfSpUploadObjectRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GfSpUploadObjectRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GfSpUploadObjectRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UploadObjectTask", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUpload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UploadObjectTask == nil {
				m.UploadObjectTask = &gfsptask.GfSpUploadObjectTask{}
			}
			if err := m.UploadObjectTask.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUpload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUpload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GfSpUploadObjectResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GfSpUploadObjectResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GfSpUploadObjectResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUpload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Err == nil {
				m.Err = &gfsperrors.GfSpError{}
			}
			if err := m.Err.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GfSpResumableUploadObjectRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GfSpResumableUploadObjectRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GfSpResumableUploadObjectRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResumableUploadObjectTask", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUpload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResumableUploadObjectTask == nil {
				m.ResumableUploadObjectTask = &gfsptask.GfSpResumableUploadObjectTask{}
			}
			if err := m.ResumableUploadObjectTask.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUpload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUpload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GfSpResumableUploadObjectResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GfSpResumableUploadObjectResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GfSpResumableUploadObjectResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUpload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Err == nil {
				m.Err = &gfsperrors.GfSpError{}
			}
			if err := m.Err.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUpload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpload
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUpload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUpload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthUpload
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUpload
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUpload
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUpload        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpload          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUpload = fmt.Errorf("proto: unexpected end of group")
)