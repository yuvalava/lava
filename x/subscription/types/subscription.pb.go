// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: subscription/subscription.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Subscription struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Consumer    string `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Block       int64  `protobuf:"varint,3,opt,name=block,proto3" json:"block,omitempty"`
	PlanIndex   string `protobuf:"bytes,4,opt,name=plan_index,json=planIndex,proto3" json:"plan_index,omitempty"`
	PlanBlock   int64  `protobuf:"varint,5,opt,name=plan_block,json=planBlock,proto3" json:"plan_block,omitempty"`
	IsYearly    bool   `protobuf:"varint,6,opt,name=is_yearly,json=isYearly,proto3" json:"is_yearly,omitempty"`
	ExpiryTime  uint64 `protobuf:"varint,7,opt,name=expiry_time,json=expiryTime,proto3" json:"expiry_time,omitempty"`
	UsedCU      uint64 `protobuf:"varint,8,opt,name=usedCU,proto3" json:"usedCU,omitempty"`
	RemainingCU uint64 `protobuf:"varint,9,opt,name=remainingCU,proto3" json:"remainingCU,omitempty"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac47bc0f89224537, []int{0}
}
func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return m.Size()
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Subscription) GetConsumer() string {
	if m != nil {
		return m.Consumer
	}
	return ""
}

func (m *Subscription) GetBlock() int64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Subscription) GetPlanIndex() string {
	if m != nil {
		return m.PlanIndex
	}
	return ""
}

func (m *Subscription) GetPlanBlock() int64 {
	if m != nil {
		return m.PlanBlock
	}
	return 0
}

func (m *Subscription) GetIsYearly() bool {
	if m != nil {
		return m.IsYearly
	}
	return false
}

func (m *Subscription) GetExpiryTime() uint64 {
	if m != nil {
		return m.ExpiryTime
	}
	return 0
}

func (m *Subscription) GetUsedCU() uint64 {
	if m != nil {
		return m.UsedCU
	}
	return 0
}

func (m *Subscription) GetRemainingCU() uint64 {
	if m != nil {
		return m.RemainingCU
	}
	return 0
}

func init() {
	proto.RegisterType((*Subscription)(nil), "lavanet.lava.subscription.Subscription")
}

func init() { proto.RegisterFile("subscription/subscription.proto", fileDescriptor_ac47bc0f89224537) }

var fileDescriptor_ac47bc0f89224537 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xeb, 0xfe, 0x4d, 0x6e, 0xbf, 0xc9, 0xfa, 0x84, 0x0c, 0x08, 0x37, 0x62, 0xca, 0x80,
	0x92, 0x81, 0x37, 0x68, 0x25, 0x24, 0xd6, 0x40, 0x07, 0x58, 0x2a, 0x27, 0xb5, 0x8a, 0x45, 0x62,
	0x47, 0xb6, 0x83, 0x9a, 0xb7, 0xe8, 0x63, 0x31, 0x76, 0x64, 0x44, 0xed, 0x8b, 0xa0, 0x38, 0xa5,
	0x0a, 0xd3, 0xd5, 0xef, 0xfc, 0xee, 0x59, 0x0e, 0xcc, 0x4c, 0x95, 0x9a, 0x4c, 0x8b, 0xd2, 0x0a,
	0x25, 0xe3, 0x2e, 0x44, 0xa5, 0x56, 0x56, 0xe1, 0xcb, 0x9c, 0x7d, 0x30, 0xc9, 0x6d, 0xd4, 0xdc,
	0xa8, 0xfb, 0x70, 0xbb, 0xeb, 0xc3, 0xbf, 0xa7, 0x4e, 0x80, 0x09, 0x4c, 0x32, 0xcd, 0x99, 0x55,
	0x9a, 0xa0, 0x00, 0x85, 0x7e, 0xf2, 0x8b, 0xf8, 0x0a, 0xbc, 0x4c, 0x49, 0x53, 0x15, 0x5c, 0x93,
	0xbe, 0x53, 0x67, 0xc6, 0xff, 0x61, 0x94, 0xe6, 0x2a, 0x7b, 0x27, 0x83, 0x00, 0x85, 0x83, 0xa4,
	0x05, 0x7c, 0x03, 0x50, 0xe6, 0x4c, 0xae, 0x84, 0x5c, 0xf3, 0x2d, 0x19, 0xba, 0x8e, 0xdf, 0x24,
	0x8f, 0x4d, 0x70, 0xd6, 0x6d, 0x73, 0xe4, 0x9a, 0x4e, 0xcf, 0x5d, 0xfb, 0x1a, 0x7c, 0x61, 0x56,
	0x35, 0x67, 0x3a, 0xaf, 0xc9, 0x38, 0x40, 0xa1, 0x97, 0x78, 0xc2, 0xbc, 0x38, 0xc6, 0x33, 0x98,
	0xf2, 0x6d, 0x29, 0x74, 0xbd, 0xb2, 0xa2, 0xe0, 0x64, 0x12, 0xa0, 0x70, 0x98, 0x40, 0x1b, 0x3d,
	0x8b, 0x82, 0xe3, 0x0b, 0x18, 0x57, 0x86, 0xaf, 0x17, 0x4b, 0xe2, 0x39, 0x77, 0x22, 0x1c, 0xc0,
	0x54, 0xf3, 0x82, 0x09, 0x29, 0xe4, 0x66, 0xb1, 0x24, 0xbe, 0x93, 0xdd, 0x68, 0xfe, 0xf0, 0x79,
	0xa0, 0x68, 0x7f, 0xa0, 0xe8, 0xfb, 0x40, 0xd1, 0xee, 0x48, 0x7b, 0xfb, 0x23, 0xed, 0x7d, 0x1d,
	0x69, 0xef, 0xf5, 0x6e, 0x23, 0xec, 0x5b, 0x95, 0x46, 0x99, 0x2a, 0xe2, 0xd3, 0xa4, 0xee, 0xc6,
	0xdb, 0x3f, 0xab, 0xc7, 0xb6, 0x2e, 0xb9, 0x49, 0xc7, 0x6e, 0xfc, 0xfb, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x13, 0x58, 0xb8, 0xa3, 0x9f, 0x01, 0x00, 0x00,
}

func (m *Subscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Subscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Subscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RemainingCU != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.RemainingCU))
		i--
		dAtA[i] = 0x48
	}
	if m.UsedCU != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.UsedCU))
		i--
		dAtA[i] = 0x40
	}
	if m.ExpiryTime != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.ExpiryTime))
		i--
		dAtA[i] = 0x38
	}
	if m.IsYearly {
		i--
		if m.IsYearly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.PlanBlock != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.PlanBlock))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PlanIndex) > 0 {
		i -= len(m.PlanIndex)
		copy(dAtA[i:], m.PlanIndex)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.PlanIndex)))
		i--
		dAtA[i] = 0x22
	}
	if m.Block != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Consumer) > 0 {
		i -= len(m.Consumer)
		copy(dAtA[i:], m.Consumer)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.Consumer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSubscription(dAtA []byte, offset int, v uint64) int {
	offset -= sovSubscription(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Subscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	l = len(m.Consumer)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovSubscription(uint64(m.Block))
	}
	l = len(m.PlanIndex)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.PlanBlock != 0 {
		n += 1 + sovSubscription(uint64(m.PlanBlock))
	}
	if m.IsYearly {
		n += 2
	}
	if m.ExpiryTime != 0 {
		n += 1 + sovSubscription(uint64(m.ExpiryTime))
	}
	if m.UsedCU != 0 {
		n += 1 + sovSubscription(uint64(m.UsedCU))
	}
	if m.RemainingCU != 0 {
		n += 1 + sovSubscription(uint64(m.RemainingCU))
	}
	return n
}

func sovSubscription(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSubscription(x uint64) (n int) {
	return sovSubscription(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Subscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubscription
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
			return fmt.Errorf("proto: Subscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Subscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consumer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Consumer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlanIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanBlock", wireType)
			}
			m.PlanBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsYearly", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsYearly = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiryTime", wireType)
			}
			m.ExpiryTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiryTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedCU", wireType)
			}
			m.UsedCU = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UsedCU |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingCU", wireType)
			}
			m.RemainingCU = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemainingCU |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSubscription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubscription
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
func skipSubscription(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
				return 0, ErrInvalidLengthSubscription
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSubscription
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSubscription
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSubscription        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubscription          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSubscription = fmt.Errorf("proto: unexpected end of group")
)