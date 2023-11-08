// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crude/crude/resource.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Resource struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RName     string `protobuf:"bytes,2,opt,name=rName,proto3" json:"rName,omitempty"`
	RCategory string `protobuf:"bytes,3,opt,name=rCategory,proto3" json:"rCategory,omitempty"`
	RColour   string `protobuf:"bytes,4,opt,name=rColour,proto3" json:"rColour,omitempty"`
	RSize     string `protobuf:"bytes,5,opt,name=rSize,proto3" json:"rSize,omitempty"`
	RQuantity uint64 `protobuf:"varint,6,opt,name=rQuantity,proto3" json:"rQuantity,omitempty"`
	Creator   string `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4e983fdb4b8595a, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Resource) GetRName() string {
	if m != nil {
		return m.RName
	}
	return ""
}

func (m *Resource) GetRCategory() string {
	if m != nil {
		return m.RCategory
	}
	return ""
}

func (m *Resource) GetRColour() string {
	if m != nil {
		return m.RColour
	}
	return ""
}

func (m *Resource) GetRSize() string {
	if m != nil {
		return m.RSize
	}
	return ""
}

func (m *Resource) GetRQuantity() uint64 {
	if m != nil {
		return m.RQuantity
	}
	return 0
}

func (m *Resource) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Resource)(nil), "crude.crude.Resource")
}

func init() { proto.RegisterFile("crude/crude/resource.proto", fileDescriptor_a4e983fdb4b8595a) }

var fileDescriptor_a4e983fdb4b8595a = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2e, 0x2a, 0x4d,
	0x49, 0xd5, 0x87, 0x90, 0x45, 0xa9, 0xc5, 0xf9, 0xa5, 0x45, 0xc9, 0xa9, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0xdc, 0x60, 0x51, 0x3d, 0x30, 0xa9, 0xb4, 0x8d, 0x91, 0x8b, 0x23, 0x08, 0x2a,
	0x2f, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x94, 0x99,
	0x22, 0x24, 0xc2, 0xc5, 0x5a, 0xe4, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0xe1, 0x08, 0xc9, 0x70, 0x71, 0x16, 0x39, 0x27, 0x96, 0xa4, 0xa6, 0xe7, 0x17, 0x55, 0x4a,
	0x30, 0x83, 0x65, 0x10, 0x02, 0x42, 0x12, 0x5c, 0xec, 0x45, 0xce, 0xf9, 0x39, 0xf9, 0xa5, 0x45,
	0x12, 0x2c, 0x60, 0x39, 0x18, 0x17, 0x6c, 0x5a, 0x70, 0x66, 0x55, 0xaa, 0x04, 0x2b, 0xd4, 0x34,
	0x10, 0x07, 0x6c, 0x5a, 0x60, 0x69, 0x62, 0x5e, 0x49, 0x66, 0x49, 0xa5, 0x04, 0x1b, 0xd8, 0x6a,
	0x84, 0x00, 0xc8, 0xb4, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0x09, 0x76, 0x88, 0x69, 0x50,
	0xae, 0x93, 0xee, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x09, 0x43, 0x7c,
	0x5d, 0x01, 0xf5, 0x7d, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0xef, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb6, 0x2f, 0x94, 0x42, 0x19, 0x01, 0x00, 0x00,
}

func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	if m.RQuantity != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.RQuantity))
		i--
		dAtA[i] = 0x30
	}
	if len(m.RSize) > 0 {
		i -= len(m.RSize)
		copy(dAtA[i:], m.RSize)
		i = encodeVarintResource(dAtA, i, uint64(len(m.RSize)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RColour) > 0 {
		i -= len(m.RColour)
		copy(dAtA[i:], m.RColour)
		i = encodeVarintResource(dAtA, i, uint64(len(m.RColour)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RCategory) > 0 {
		i -= len(m.RCategory)
		copy(dAtA[i:], m.RCategory)
		i = encodeVarintResource(dAtA, i, uint64(len(m.RCategory)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RName) > 0 {
		i -= len(m.RName)
		copy(dAtA[i:], m.RName)
		i = encodeVarintResource(dAtA, i, uint64(len(m.RName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintResource(dAtA []byte, offset int, v uint64) int {
	offset -= sovResource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Resource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovResource(uint64(m.Id))
	}
	l = len(m.RName)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.RCategory)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.RColour)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.RSize)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	if m.RQuantity != 0 {
		n += 1 + sovResource(uint64(m.RQuantity))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func sovResource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResource(x uint64) (n int) {
	return sovResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
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
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
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
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
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
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RSize = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RQuantity", wireType)
			}
			m.RQuantity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RQuantity |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
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
func skipResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
				return 0, ErrInvalidLengthResource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResource = fmt.Errorf("proto: unexpected end of group")
)