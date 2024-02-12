// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/reporter/oracle_reporter.proto

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

type OracleReporter struct {
	Reporter    string `protobuf:"bytes,1,opt,name=reporter,proto3" json:"reporter,omitempty"`
	TotalTokens uint64 `protobuf:"varint,2,opt,name=totalTokens,proto3" json:"totalTokens,omitempty"`
}

func (m *OracleReporter) Reset()         { *m = OracleReporter{} }
func (m *OracleReporter) String() string { return proto.CompactTextString(m) }
func (*OracleReporter) ProtoMessage()    {}
func (*OracleReporter) Descriptor() ([]byte, []int) {
	return fileDescriptor_28310cb3dcf79802, []int{0}
}
func (m *OracleReporter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleReporter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleReporter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleReporter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleReporter.Merge(m, src)
}
func (m *OracleReporter) XXX_Size() int {
	return m.Size()
}
func (m *OracleReporter) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleReporter.DiscardUnknown(m)
}

var xxx_messageInfo_OracleReporter proto.InternalMessageInfo

func (m *OracleReporter) GetReporter() string {
	if m != nil {
		return m.Reporter
	}
	return ""
}

func (m *OracleReporter) GetTotalTokens() uint64 {
	if m != nil {
		return m.TotalTokens
	}
	return 0
}

func init() {
	proto.RegisterType((*OracleReporter)(nil), "layer.reporter.OracleReporter")
}

func init() {
	proto.RegisterFile("layer/reporter/oracle_reporter.proto", fileDescriptor_28310cb3dcf79802)
}

var fileDescriptor_28310cb3dcf79802 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0x49, 0xac, 0x4c,
	0x2d, 0xd2, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0x49, 0x2d, 0xd2, 0xcf, 0x2f, 0x4a, 0x4c, 0xce,
	0x49, 0x8d, 0x87, 0xf1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0xc0, 0xaa, 0xf4, 0x60,
	0xa2, 0x4a, 0x7e, 0x5c, 0x7c, 0xfe, 0x60, 0x85, 0x41, 0x50, 0x11, 0x21, 0x29, 0x2e, 0x0e, 0x98,
	0xac, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0x2f, 0xa4, 0xc0, 0xc5, 0x5d, 0x92, 0x5f,
	0x92, 0x98, 0x13, 0x92, 0x9f, 0x9d, 0x9a, 0x57, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x12, 0x84,
	0x2c, 0xe4, 0xe4, 0x7a, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31,
	0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xda, 0xe9,
	0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x25, 0xa9, 0x39, 0x39, 0xf9, 0x45,
	0xba, 0x99, 0xf9, 0xfa, 0x10, 0x47, 0x57, 0x20, 0x9c, 0x5d, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4,
	0x06, 0x76, 0xad, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x30, 0xe5, 0x3a, 0x8d, 0xd5, 0x00, 0x00,
	0x00,
}

func (m *OracleReporter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleReporter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleReporter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalTokens != 0 {
		i = encodeVarintOracleReporter(dAtA, i, uint64(m.TotalTokens))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Reporter) > 0 {
		i -= len(m.Reporter)
		copy(dAtA[i:], m.Reporter)
		i = encodeVarintOracleReporter(dAtA, i, uint64(len(m.Reporter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracleReporter(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracleReporter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OracleReporter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Reporter)
	if l > 0 {
		n += 1 + l + sovOracleReporter(uint64(l))
	}
	if m.TotalTokens != 0 {
		n += 1 + sovOracleReporter(uint64(m.TotalTokens))
	}
	return n
}

func sovOracleReporter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracleReporter(x uint64) (n int) {
	return sovOracleReporter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OracleReporter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracleReporter
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
			return fmt.Errorf("proto: OracleReporter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleReporter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracleReporter
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
				return ErrInvalidLengthOracleReporter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracleReporter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalTokens", wireType)
			}
			m.TotalTokens = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracleReporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalTokens |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracleReporter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracleReporter
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
func skipOracleReporter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracleReporter
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
					return 0, ErrIntOverflowOracleReporter
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
					return 0, ErrIntOverflowOracleReporter
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
				return 0, ErrInvalidLengthOracleReporter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracleReporter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracleReporter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracleReporter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracleReporter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracleReporter = fmt.Errorf("proto: unexpected end of group")
)
