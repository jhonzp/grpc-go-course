// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculator_pb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Adding struct {
	Addone               int32    `protobuf:"varint,1,opt,name=addone,proto3" json:"addone,omitempty"`
	Addtwo               int32    `protobuf:"varint,2,opt,name=addtwo,proto3" json:"addtwo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Adding) Reset()         { *m = Adding{} }
func (m *Adding) String() string { return proto.CompactTextString(m) }
func (*Adding) ProtoMessage()    {}
func (*Adding) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{0}
}

func (m *Adding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Adding.Unmarshal(m, b)
}
func (m *Adding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Adding.Marshal(b, m, deterministic)
}
func (m *Adding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Adding.Merge(m, src)
}
func (m *Adding) XXX_Size() int {
	return xxx_messageInfo_Adding.Size(m)
}
func (m *Adding) XXX_DiscardUnknown() {
	xxx_messageInfo_Adding.DiscardUnknown(m)
}

var xxx_messageInfo_Adding proto.InternalMessageInfo

func (m *Adding) GetAddone() int32 {
	if m != nil {
		return m.Addone
	}
	return 0
}

func (m *Adding) GetAddtwo() int32 {
	if m != nil {
		return m.Addtwo
	}
	return 0
}

type AddRequest struct {
	Adding               *Adding  `protobuf:"bytes,1,opt,name=adding,proto3" json:"adding,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{1}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetAdding() *Adding {
	if m != nil {
		return m.Adding
	}
	return nil
}

type AddResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{2}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeNumber struct {
	PrimeNumber          int32    `protobuf:"varint,1,opt,name=primeNumber,proto3" json:"primeNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumber) Reset()         { *m = PrimeNumber{} }
func (m *PrimeNumber) String() string { return proto.CompactTextString(m) }
func (*PrimeNumber) ProtoMessage()    {}
func (*PrimeNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{3}
}

func (m *PrimeNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumber.Unmarshal(m, b)
}
func (m *PrimeNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumber.Marshal(b, m, deterministic)
}
func (m *PrimeNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumber.Merge(m, src)
}
func (m *PrimeNumber) XXX_Size() int {
	return xxx_messageInfo_PrimeNumber.Size(m)
}
func (m *PrimeNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumber.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumber proto.InternalMessageInfo

func (m *PrimeNumber) GetPrimeNumber() int32 {
	if m != nil {
		return m.PrimeNumber
	}
	return 0
}

type PrimeNumberDescompositionRequest struct {
	PrimeNumber          *PrimeNumber `protobuf:"bytes,1,opt,name=primeNumber,proto3" json:"primeNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PrimeNumberDescompositionRequest) Reset()         { *m = PrimeNumberDescompositionRequest{} }
func (m *PrimeNumberDescompositionRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDescompositionRequest) ProtoMessage()    {}
func (*PrimeNumberDescompositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{4}
}

func (m *PrimeNumberDescompositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDescompositionRequest.Unmarshal(m, b)
}
func (m *PrimeNumberDescompositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDescompositionRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDescompositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDescompositionRequest.Merge(m, src)
}
func (m *PrimeNumberDescompositionRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDescompositionRequest.Size(m)
}
func (m *PrimeNumberDescompositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDescompositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDescompositionRequest proto.InternalMessageInfo

func (m *PrimeNumberDescompositionRequest) GetPrimeNumber() *PrimeNumber {
	if m != nil {
		return m.PrimeNumber
	}
	return nil
}

type PrimeNumberDescompositionResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDescompositionResponse) Reset()         { *m = PrimeNumberDescompositionResponse{} }
func (m *PrimeNumberDescompositionResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDescompositionResponse) ProtoMessage()    {}
func (*PrimeNumberDescompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{5}
}

func (m *PrimeNumberDescompositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDescompositionResponse.Unmarshal(m, b)
}
func (m *PrimeNumberDescompositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDescompositionResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDescompositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDescompositionResponse.Merge(m, src)
}
func (m *PrimeNumberDescompositionResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDescompositionResponse.Size(m)
}
func (m *PrimeNumberDescompositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDescompositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDescompositionResponse proto.InternalMessageInfo

func (m *PrimeNumberDescompositionResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type AVGNumberRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AVGNumberRequest) Reset()         { *m = AVGNumberRequest{} }
func (m *AVGNumberRequest) String() string { return proto.CompactTextString(m) }
func (*AVGNumberRequest) ProtoMessage()    {}
func (*AVGNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{6}
}

func (m *AVGNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AVGNumberRequest.Unmarshal(m, b)
}
func (m *AVGNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AVGNumberRequest.Marshal(b, m, deterministic)
}
func (m *AVGNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AVGNumberRequest.Merge(m, src)
}
func (m *AVGNumberRequest) XXX_Size() int {
	return xxx_messageInfo_AVGNumberRequest.Size(m)
}
func (m *AVGNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AVGNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AVGNumberRequest proto.InternalMessageInfo

func (m *AVGNumberRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type AVGResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AVGResponse) Reset()         { *m = AVGResponse{} }
func (m *AVGResponse) String() string { return proto.CompactTextString(m) }
func (*AVGResponse) ProtoMessage()    {}
func (*AVGResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{7}
}

func (m *AVGResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AVGResponse.Unmarshal(m, b)
}
func (m *AVGResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AVGResponse.Marshal(b, m, deterministic)
}
func (m *AVGResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AVGResponse.Merge(m, src)
}
func (m *AVGResponse) XXX_Size() int {
	return xxx_messageInfo_AVGResponse.Size(m)
}
func (m *AVGResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AVGResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AVGResponse proto.InternalMessageInfo

func (m *AVGResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Adding)(nil), "calculator.Adding")
	proto.RegisterType((*AddRequest)(nil), "calculator.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "calculator.AddResponse")
	proto.RegisterType((*PrimeNumber)(nil), "calculator.PrimeNumber")
	proto.RegisterType((*PrimeNumberDescompositionRequest)(nil), "calculator.PrimeNumberDescompositionRequest")
	proto.RegisterType((*PrimeNumberDescompositionResponse)(nil), "calculator.PrimeNumberDescompositionResponse")
	proto.RegisterType((*AVGNumberRequest)(nil), "calculator.AVGNumberRequest")
	proto.RegisterType((*AVGResponse)(nil), "calculator.AVGResponse")
}

func init() {
	proto.RegisterFile("calculator/calculator_pb/calculator.proto", fileDescriptor_5b44e1cc294e667d)
}

var fileDescriptor_5b44e1cc294e667d = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4b, 0x3a, 0x41,
	0x18, 0x75, 0xfd, 0xf1, 0xdb, 0xc3, 0xb7, 0x21, 0x35, 0x07, 0x35, 0xe9, 0x60, 0x03, 0x81, 0x49,
	0x69, 0xd8, 0xc5, 0xe8, 0xb4, 0x19, 0x78, 0x08, 0x22, 0x0c, 0x3c, 0x04, 0x11, 0xbb, 0x3b, 0x83,
	0x2c, 0xb8, 0x33, 0xd3, 0xcc, 0xac, 0x5e, 0xfa, 0x43, 0xfa, 0x73, 0x83, 0x71, 0xd6, 0x1d, 0x35,
	0xab, 0xdb, 0xbe, 0xb7, 0xdf, 0xf7, 0xde, 0xfb, 0x1e, 0x03, 0xe7, 0x49, 0x34, 0x4f, 0xf2, 0x79,
	0xa4, 0xb9, 0xec, 0x97, 0x9f, 0x6f, 0x22, 0x76, 0x50, 0x4f, 0x48, 0xae, 0x39, 0x82, 0x92, 0xc1,
	0x43, 0xf0, 0x43, 0x42, 0x52, 0x36, 0x43, 0x75, 0xf0, 0x23, 0x42, 0x38, 0xa3, 0x4d, 0xaf, 0xed,
	0x75, 0xfe, 0x4f, 0x2c, 0xb2, 0xbc, 0x5e, 0xf2, 0x66, 0x75, 0xcd, 0xeb, 0x25, 0xc7, 0x43, 0x80,
	0x90, 0x90, 0x09, 0x7d, 0xcf, 0xa9, 0xd2, 0xa8, 0x6b, 0xa6, 0x52, 0x36, 0x33, 0xdb, 0xc1, 0x00,
	0xf5, 0x1c, 0xdb, 0x95, 0xc3, 0xc4, 0x4e, 0xe0, 0x33, 0x08, 0xcc, 0xa6, 0x12, 0x9c, 0x29, 0x63,
	0x20, 0xa9, 0xca, 0xe7, 0xba, 0x30, 0x5e, 0x21, 0xdc, 0x87, 0xe0, 0x49, 0xa6, 0x19, 0x7d, 0xcc,
	0xb3, 0x98, 0x4a, 0xd4, 0x86, 0x40, 0x94, 0xd0, 0xce, 0xba, 0x14, 0x7e, 0x85, 0xb6, 0xb3, 0x70,
	0x4f, 0x55, 0xc2, 0x33, 0xc1, 0x55, 0xaa, 0x53, 0xce, 0x8a, 0x9c, 0x37, 0xbb, 0x2a, 0xc1, 0xa0,
	0xe1, 0x86, 0x75, 0x24, 0x36, 0xe5, 0x6f, 0xe1, 0xf4, 0x07, 0xf9, 0x5f, 0x8e, 0xe9, 0xc2, 0x61,
	0x38, 0x1d, 0x5b, 0x59, 0x9b, 0xa5, 0x0e, 0x3e, 0x73, 0x8f, 0xb1, 0xc8, 0xf4, 0x33, 0x1d, 0xef,
	0x91, 0xf4, 0x0a, 0xc9, 0xc1, 0x67, 0x15, 0x8e, 0x46, 0xeb, 0xdc, 0xcf, 0x54, 0x2e, 0xd2, 0x84,
	0xa2, 0x21, 0xfc, 0x0b, 0x09, 0x41, 0xf5, 0xad, 0xfe, 0xad, 0x67, 0xab, 0xb1, 0xc3, 0xaf, 0x5c,
	0x70, 0x05, 0x7d, 0xc0, 0xf1, 0xde, 0xfb, 0xd0, 0xc5, 0x9e, 0x8a, 0xbe, 0x6d, 0xb9, 0x75, 0xf9,
	0xc7, 0xe9, 0xc2, 0xfb, 0xca, 0x43, 0x0f, 0x50, 0x1b, 0xf1, 0x4c, 0xe4, 0x9a, 0x86, 0x8b, 0x48,
	0x46, 0x33, 0x8a, 0x4e, 0x36, 0xa2, 0x6e, 0x95, 0xb7, 0x75, 0x48, 0x59, 0x17, 0xae, 0x74, 0xbc,
	0xbb, 0xda, 0xcb, 0x41, 0xf9, 0x57, 0xc4, 0xb1, 0x6f, 0x1e, 0xfe, 0xf5, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xac, 0x26, 0x6c, 0x7b, 0x25, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	//unary
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	//Server Streaming
	PrimeNumberDescomposition(ctx context.Context, in *PrimeNumberDescompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDescompositionClient, error)
	//Client Streaming
	ComputeAvarage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAvarageClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDescomposition(ctx context.Context, in *PrimeNumberDescompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDescompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeNumberDescomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDescompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDescompositionClient interface {
	Recv() (*PrimeNumberDescompositionResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDescompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDescompositionClient) Recv() (*PrimeNumberDescompositionResponse, error) {
	m := new(PrimeNumberDescompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAvarage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAvarageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/ComputeAvarage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAvarageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAvarageClient interface {
	Send(*AVGNumberRequest) error
	CloseAndRecv() (*AVGResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAvarageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAvarageClient) Send(m *AVGNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAvarageClient) CloseAndRecv() (*AVGResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AVGResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	//unary
	Add(context.Context, *AddRequest) (*AddResponse, error)
	//Server Streaming
	PrimeNumberDescomposition(*PrimeNumberDescompositionRequest, CalculatorService_PrimeNumberDescompositionServer) error
	//Client Streaming
	ComputeAvarage(CalculatorService_ComputeAvarageServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeNumberDescomposition(req *PrimeNumberDescompositionRequest, srv CalculatorService_PrimeNumberDescompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDescomposition not implemented")
}
func (*UnimplementedCalculatorServiceServer) ComputeAvarage(srv CalculatorService_ComputeAvarageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAvarage not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDescomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDescompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDescomposition(m, &calculatorServicePrimeNumberDescompositionServer{stream})
}

type CalculatorService_PrimeNumberDescompositionServer interface {
	Send(*PrimeNumberDescompositionResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDescompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDescompositionServer) Send(m *PrimeNumberDescompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAvarage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAvarage(&calculatorServiceComputeAvarageServer{stream})
}

type CalculatorService_ComputeAvarageServer interface {
	SendAndClose(*AVGResponse) error
	Recv() (*AVGNumberRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAvarageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAvarageServer) SendAndClose(m *AVGResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAvarageServer) Recv() (*AVGNumberRequest, error) {
	m := new(AVGNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalculatorService_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDescomposition",
			Handler:       _CalculatorService_PrimeNumberDescomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAvarage",
			Handler:       _CalculatorService_ComputeAvarage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculator_pb/calculator.proto",
}
