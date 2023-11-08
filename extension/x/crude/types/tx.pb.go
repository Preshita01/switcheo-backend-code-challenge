// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crude/crude/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

type MsgCreateResource struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	RName     string `protobuf:"bytes,2,opt,name=rName,proto3" json:"rName,omitempty"`
	RCategory string `protobuf:"bytes,3,opt,name=rCategory,proto3" json:"rCategory,omitempty"`
	RColour   string `protobuf:"bytes,4,opt,name=rColour,proto3" json:"rColour,omitempty"`
	RSize     string `protobuf:"bytes,5,opt,name=rSize,proto3" json:"rSize,omitempty"`
	RQuantity string `protobuf:"bytes,6,opt,name=rQuantity,proto3" json:"rQuantity,omitempty"`
}

func (m *MsgCreateResource) Reset()         { *m = MsgCreateResource{} }
func (m *MsgCreateResource) String() string { return proto.CompactTextString(m) }
func (*MsgCreateResource) ProtoMessage()    {}
func (*MsgCreateResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b55c834faf70bc, []int{0}
}
func (m *MsgCreateResource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateResource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateResource.Merge(m, src)
}
func (m *MsgCreateResource) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateResource) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateResource.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateResource proto.InternalMessageInfo

func (m *MsgCreateResource) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateResource) GetRName() string {
	if m != nil {
		return m.RName
	}
	return ""
}

func (m *MsgCreateResource) GetRCategory() string {
	if m != nil {
		return m.RCategory
	}
	return ""
}

func (m *MsgCreateResource) GetRColour() string {
	if m != nil {
		return m.RColour
	}
	return ""
}

func (m *MsgCreateResource) GetRSize() string {
	if m != nil {
		return m.RSize
	}
	return ""
}

func (m *MsgCreateResource) GetRQuantity() string {
	if m != nil {
		return m.RQuantity
	}
	return ""
}

type MsgCreateResourceResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreateResourceResponse) Reset()         { *m = MsgCreateResourceResponse{} }
func (m *MsgCreateResourceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateResourceResponse) ProtoMessage()    {}
func (*MsgCreateResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b55c834faf70bc, []int{1}
}
func (m *MsgCreateResourceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateResourceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateResourceResponse.Merge(m, src)
}
func (m *MsgCreateResourceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateResourceResponse proto.InternalMessageInfo

func (m *MsgCreateResourceResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreateResource)(nil), "crude.crude.MsgCreateResource")
	proto.RegisterType((*MsgCreateResourceResponse)(nil), "crude.crude.MsgCreateResourceResponse")
}

func init() { proto.RegisterFile("crude/crude/tx.proto", fileDescriptor_60b55c834faf70bc) }

var fileDescriptor_60b55c834faf70bc = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2e, 0x2a, 0x4d,
	0x49, 0xd5, 0x87, 0x90, 0x25, 0x15, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xdc, 0x60, 0xbe,
	0x1e, 0x98, 0x54, 0x5a, 0xcf, 0xc8, 0x25, 0xe8, 0x5b, 0x9c, 0xee, 0x5c, 0x94, 0x9a, 0x58, 0x92,
	0x1a, 0x94, 0x5a, 0x9c, 0x5f, 0x5a, 0x94, 0x9c, 0x2a, 0x24, 0xc1, 0xc5, 0x9e, 0x0c, 0x12, 0xc9,
	0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0x44, 0xb8, 0x58, 0x8b, 0xfc,
	0x12, 0x73, 0x53, 0x25, 0x98, 0xc0, 0xe2, 0x10, 0x8e, 0x90, 0x0c, 0x17, 0x67, 0x91, 0x73, 0x62,
	0x49, 0x6a, 0x7a, 0x7e, 0x51, 0xa5, 0x04, 0x33, 0x58, 0x06, 0x21, 0x00, 0x32, 0xad, 0xc8, 0x39,
	0x3f, 0x27, 0xbf, 0xb4, 0x48, 0x82, 0x05, 0x62, 0x1a, 0x94, 0x0b, 0x36, 0x2d, 0x38, 0xb3, 0x2a,
	0x55, 0x82, 0x15, 0x6a, 0x1a, 0x88, 0x03, 0x36, 0x2d, 0xb0, 0x34, 0x31, 0xaf, 0x24, 0xb3, 0xa4,
	0x52, 0x82, 0x0d, 0x6a, 0x1a, 0x4c, 0x40, 0x49, 0x9b, 0x4b, 0x12, 0xc3, 0xc1, 0x41, 0xa9, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x60, 0x37, 0xb3, 0x04, 0x31,
	0x65, 0xa6, 0x18, 0xc5, 0x73, 0x31, 0xfb, 0x16, 0xa7, 0x0b, 0x45, 0x70, 0xf1, 0xa1, 0xf9, 0x50,
	0x4e, 0x0f, 0x29, 0x14, 0xf4, 0x30, 0x0c, 0x94, 0x52, 0xc3, 0x2f, 0x0f, 0xb3, 0xd0, 0x49, 0xf7,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x84, 0x21, 0xc1, 0x5e, 0x01, 0x0b,
	0xfe, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x14, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x9c, 0xc1, 0x9b, 0x1e, 0x9a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateResource(ctx context.Context, in *MsgCreateResource, opts ...grpc.CallOption) (*MsgCreateResourceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateResource(ctx context.Context, in *MsgCreateResource, opts ...grpc.CallOption) (*MsgCreateResourceResponse, error) {
	out := new(MsgCreateResourceResponse)
	err := c.cc.Invoke(ctx, "/crude.crude.Msg/CreateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateResource(context.Context, *MsgCreateResource) (*MsgCreateResourceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateResource(ctx context.Context, req *MsgCreateResource) (*MsgCreateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateResource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crude.crude.Msg/CreateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateResource(ctx, req.(*MsgCreateResource))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crude.crude.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResource",
			Handler:    _Msg_CreateResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crude/crude/tx.proto",
}

func (m *MsgCreateResource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateResource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateResource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RQuantity) > 0 {
		i -= len(m.RQuantity)
		copy(dAtA[i:], m.RQuantity)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RQuantity)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.RSize) > 0 {
		i -= len(m.RSize)
		copy(dAtA[i:], m.RSize)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RSize)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RColour) > 0 {
		i -= len(m.RColour)
		copy(dAtA[i:], m.RColour)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RColour)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RCategory) > 0 {
		i -= len(m.RCategory)
		copy(dAtA[i:], m.RCategory)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RCategory)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RName) > 0 {
		i -= len(m.RName)
		copy(dAtA[i:], m.RName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateResourceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateResourceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateResourceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateResource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RCategory)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RColour)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RSize)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RQuantity)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateResourceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateResource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateResource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateResource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RCategory", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RCategory = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RColour", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RColour = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RSize", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RSize = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RQuantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RQuantity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateResourceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateResourceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateResourceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
