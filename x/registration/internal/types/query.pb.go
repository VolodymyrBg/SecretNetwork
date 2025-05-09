// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: secret/registration/v1beta1/query.proto

package types

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type QueryEncryptedSeedRequest struct {
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *QueryEncryptedSeedRequest) Reset()         { *m = QueryEncryptedSeedRequest{} }
func (m *QueryEncryptedSeedRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEncryptedSeedRequest) ProtoMessage()    {}
func (*QueryEncryptedSeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ee71413f073b37c, []int{0}
}
func (m *QueryEncryptedSeedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEncryptedSeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEncryptedSeedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEncryptedSeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEncryptedSeedRequest.Merge(m, src)
}
func (m *QueryEncryptedSeedRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEncryptedSeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEncryptedSeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEncryptedSeedRequest proto.InternalMessageInfo

type QueryEncryptedSeedResponse struct {
	EncryptedSeed []byte `protobuf:"bytes,1,opt,name=encrypted_seed,json=encryptedSeed,proto3" json:"encrypted_seed,omitempty"`
}

func (m *QueryEncryptedSeedResponse) Reset()         { *m = QueryEncryptedSeedResponse{} }
func (m *QueryEncryptedSeedResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEncryptedSeedResponse) ProtoMessage()    {}
func (*QueryEncryptedSeedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ee71413f073b37c, []int{1}
}
func (m *QueryEncryptedSeedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEncryptedSeedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEncryptedSeedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEncryptedSeedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEncryptedSeedResponse.Merge(m, src)
}
func (m *QueryEncryptedSeedResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEncryptedSeedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEncryptedSeedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEncryptedSeedResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryEncryptedSeedRequest)(nil), "secret.registration.v1beta1.QueryEncryptedSeedRequest")
	proto.RegisterType((*QueryEncryptedSeedResponse)(nil), "secret.registration.v1beta1.QueryEncryptedSeedResponse")
}

func init() {
	proto.RegisterFile("secret/registration/v1beta1/query.proto", fileDescriptor_7ee71413f073b37c)
}

var fileDescriptor_7ee71413f073b37c = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0xcd, 0x8a, 0xd4, 0x40,
	0x18, 0x4c, 0x94, 0x5d, 0xa1, 0x71, 0x15, 0x1a, 0xf1, 0x27, 0xbb, 0x34, 0xcb, 0xe0, 0xea, 0x7a,
	0xd8, 0x6e, 0x57, 0x65, 0x3d, 0x0a, 0xca, 0x9e, 0x16, 0x04, 0x67, 0x3d, 0x79, 0x59, 0x92, 0xcc,
	0x67, 0x1b, 0x66, 0xa6, 0xbb, 0xa7, 0xbb, 0xa3, 0x13, 0x06, 0x2f, 0x3e, 0x81, 0xe0, 0x4b, 0xcc,
	0x23, 0xf8, 0x08, 0x73, 0x1c, 0xf0, 0xe2, 0x51, 0x33, 0x3e, 0x88, 0xa4, 0xd3, 0x23, 0x19, 0x88,
	0x01, 0xf1, 0x96, 0xf4, 0x57, 0x5f, 0x55, 0x7d, 0x55, 0xe8, 0xbe, 0x81, 0x54, 0x83, 0x65, 0x1a,
	0x78, 0x66, 0xac, 0x8e, 0x6d, 0x26, 0x05, 0x7b, 0x7f, 0x9c, 0x80, 0x8d, 0x8f, 0xd9, 0x24, 0x07,
	0x5d, 0x50, 0xa5, 0xa5, 0x95, 0x78, 0xb7, 0x06, 0xd2, 0x26, 0x90, 0x7a, 0x60, 0x74, 0x83, 0x4b,
	0x2e, 0x1d, 0x8e, 0x55, 0x5f, 0xf5, 0x4a, 0xb4, 0xcb, 0xa5, 0xe4, 0x23, 0x60, 0xee, 0x2f, 0xc9,
	0xdf, 0x32, 0x18, 0x2b, 0xeb, 0xf9, 0xa2, 0x3d, 0x3f, 0x8c, 0x55, 0xc6, 0x62, 0x21, 0xa4, 0x75,
	0x8c, 0xc6, 0x4f, 0x0f, 0xba, 0x6c, 0x8d, 0x0d, 0xf7, 0xb0, 0x07, 0x5d, 0x30, 0x0e, 0x02, 0x4c,
	0xe6, 0x19, 0x7b, 0x4f, 0xd0, 0x9d, 0x57, 0xd5, 0x39, 0xa7, 0x22, 0xd5, 0x85, 0xb2, 0x30, 0x38,
	0x07, 0x18, 0xf4, 0x61, 0x92, 0x83, 0xb1, 0xf8, 0x16, 0xba, 0xa2, 0xf2, 0xe4, 0x62, 0x08, 0xc5,
	0xed, 0x70, 0x3f, 0x3c, 0xbc, 0xda, 0xdf, 0x56, 0x79, 0x72, 0x06, 0x45, 0xef, 0x05, 0x8a, 0xda,
	0xb6, 0x8c, 0x92, 0xc2, 0x00, 0x3e, 0x40, 0xd7, 0x60, 0x3d, 0xb8, 0x30, 0x00, 0x03, 0xbf, 0xbd,
	0x03, 0x4d, 0xf8, 0xa3, 0xf9, 0x65, 0xb4, 0xe5, 0x58, 0x30, 0x47, 0x5b, 0xaf, 0xa7, 0x67, 0x50,
	0xe0, 0x9b, 0xb4, 0x3e, 0x9f, 0xae, 0xb3, 0xa1, 0xa7, 0x55, 0x36, 0xd1, 0x3e, 0xed, 0x88, 0x99,
	0x56, 0x8e, 0xee, 0x7e, 0xfa, 0xf6, 0xeb, 0xcb, 0x25, 0x82, 0xf7, 0xda, 0x8f, 0xb6, 0xd3, 0xa3,
	0x21, 0x14, 0x78, 0x86, 0xae, 0xf7, 0x1b, 0xe3, 0xff, 0x93, 0xa4, 0x4e, 0xf2, 0x10, 0xdf, 0x6b,
	0x97, 0x6c, 0x3e, 0x3a, 0xf1, 0xaf, 0x21, 0xda, 0xd9, 0x08, 0x0c, 0x9f, 0x74, 0x6a, 0xfc, 0xb5,
	0x97, 0xe8, 0xe9, 0x3f, 0xef, 0xd5, 0xcd, 0xf4, 0x4e, 0x9c, 0xe5, 0x87, 0x98, 0xb6, 0x5b, 0xfe,
	0xd3, 0xcf, 0x51, 0xd5, 0x1a, 0x9b, 0xf9, 0xf2, 0x3f, 0x3e, 0x8f, 0x17, 0x3f, 0x49, 0x30, 0x2f,
	0x49, 0xb8, 0x28, 0x49, 0xb8, 0x2c, 0x49, 0xf8, 0xa3, 0x24, 0xe1, 0xe7, 0x15, 0x09, 0x96, 0x2b,
	0x12, 0x7c, 0x5f, 0x91, 0xe0, 0xcd, 0x33, 0x9e, 0xd9, 0x77, 0x79, 0x42, 0x53, 0x39, 0x66, 0x26,
	0xd5, 0x76, 0x14, 0x27, 0x86, 0x9d, 0x3b, 0x97, 0x2f, 0xc1, 0x7e, 0x90, 0x7a, 0xc8, 0xa6, 0x9b,
	0xa2, 0x99, 0xb0, 0xa0, 0x45, 0x3c, 0x62, 0xb6, 0x50, 0x60, 0x92, 0x6d, 0x97, 0xff, 0xe3, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x62, 0xa4, 0x1c, 0x7a, 0x03, 0x00, 0x00,
}

func (this *QueryEncryptedSeedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryEncryptedSeedRequest)
	if !ok {
		that2, ok := that.(QueryEncryptedSeedRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.PubKey, that1.PubKey) {
		return false
	}
	return true
}
func (this *QueryEncryptedSeedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryEncryptedSeedResponse)
	if !ok {
		that2, ok := that.(QueryEncryptedSeedResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.EncryptedSeed, that1.EncryptedSeed) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Returns the key used for transactions
	TxKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Key, error)
	// Returns the key used for registration
	RegistrationKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Key, error)
	// Returns the encrypted seed for a registered node by public key
	EncryptedSeed(ctx context.Context, in *QueryEncryptedSeedRequest, opts ...grpc.CallOption) (*QueryEncryptedSeedResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) TxKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Key, error) {
	out := new(Key)
	err := c.cc.Invoke(ctx, "/secret.registration.v1beta1.Query/TxKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RegistrationKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Key, error) {
	out := new(Key)
	err := c.cc.Invoke(ctx, "/secret.registration.v1beta1.Query/RegistrationKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EncryptedSeed(ctx context.Context, in *QueryEncryptedSeedRequest, opts ...grpc.CallOption) (*QueryEncryptedSeedResponse, error) {
	out := new(QueryEncryptedSeedResponse)
	err := c.cc.Invoke(ctx, "/secret.registration.v1beta1.Query/EncryptedSeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Returns the key used for transactions
	TxKey(context.Context, *emptypb.Empty) (*Key, error)
	// Returns the key used for registration
	RegistrationKey(context.Context, *emptypb.Empty) (*Key, error)
	// Returns the encrypted seed for a registered node by public key
	EncryptedSeed(context.Context, *QueryEncryptedSeedRequest) (*QueryEncryptedSeedResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) TxKey(ctx context.Context, req *emptypb.Empty) (*Key, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxKey not implemented")
}
func (*UnimplementedQueryServer) RegistrationKey(ctx context.Context, req *emptypb.Empty) (*Key, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrationKey not implemented")
}
func (*UnimplementedQueryServer) EncryptedSeed(ctx context.Context, req *QueryEncryptedSeedRequest) (*QueryEncryptedSeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncryptedSeed not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_TxKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TxKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secret.registration.v1beta1.Query/TxKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TxKey(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RegistrationKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RegistrationKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secret.registration.v1beta1.Query/RegistrationKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RegistrationKey(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EncryptedSeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEncryptedSeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EncryptedSeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secret.registration.v1beta1.Query/EncryptedSeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EncryptedSeed(ctx, req.(*QueryEncryptedSeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "secret.registration.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TxKey",
			Handler:    _Query_TxKey_Handler,
		},
		{
			MethodName: "RegistrationKey",
			Handler:    _Query_RegistrationKey_Handler,
		},
		{
			MethodName: "EncryptedSeed",
			Handler:    _Query_EncryptedSeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secret/registration/v1beta1/query.proto",
}

func (m *QueryEncryptedSeedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEncryptedSeedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEncryptedSeedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEncryptedSeedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEncryptedSeedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEncryptedSeedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptedSeed) > 0 {
		i -= len(m.EncryptedSeed)
		copy(dAtA[i:], m.EncryptedSeed)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EncryptedSeed)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryEncryptedSeedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEncryptedSeedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EncryptedSeed)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryEncryptedSeedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEncryptedSeedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEncryptedSeedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEncryptedSeedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEncryptedSeedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEncryptedSeedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedSeed", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedSeed = append(m.EncryptedSeed[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptedSeed == nil {
				m.EncryptedSeed = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
