// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blockheadandtxs.proto

//包名，通过protoc生成时go文件时

package go_protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BlockHeadAndTxs struct {
	StaretBlockHash      []byte     `protobuf:"bytes,1,opt,name=StaretBlockHash,proto3" json:"StaretBlockHash,omitempty"`
	Bh                   *BlockHead `protobuf:"bytes,2,opt,name=Bh,proto3" json:"Bh,omitempty"`
	TxBs                 [][]byte   `protobuf:"bytes,3,rep,name=TxBs,proto3" json:"TxBs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BlockHeadAndTxs) Reset()         { *m = BlockHeadAndTxs{} }
func (m *BlockHeadAndTxs) String() string { return proto.CompactTextString(m) }
func (*BlockHeadAndTxs) ProtoMessage()    {}
func (*BlockHeadAndTxs) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ffea91f5dd76ba4, []int{0}
}
func (m *BlockHeadAndTxs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeadAndTxs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeadAndTxs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeadAndTxs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeadAndTxs.Merge(m, src)
}
func (m *BlockHeadAndTxs) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeadAndTxs) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeadAndTxs.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeadAndTxs proto.InternalMessageInfo

func (m *BlockHeadAndTxs) GetStaretBlockHash() []byte {
	if m != nil {
		return m.StaretBlockHash
	}
	return nil
}

func (m *BlockHeadAndTxs) GetBh() *BlockHead {
	if m != nil {
		return m.Bh
	}
	return nil
}

func (m *BlockHeadAndTxs) GetTxBs() [][]byte {
	if m != nil {
		return m.TxBs
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeadAndTxs)(nil), "go_protos.BlockHeadAndTxs")
}

func init() { proto.RegisterFile("blockheadandtxs.proto", fileDescriptor_6ffea91f5dd76ba4) }

var fileDescriptor_6ffea91f5dd76ba4 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xca, 0xc9, 0x4f,
	0xce, 0xce, 0x48, 0x4d, 0x4c, 0x49, 0xcc, 0x4b, 0x29, 0xa9, 0x28, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x4c, 0xcf, 0x8f, 0x07, 0xb3, 0x8a, 0xa5, 0xf8, 0xe1, 0x2a, 0x20, 0x72, 0x4a,
	0xa5, 0x5c, 0xfc, 0x4e, 0x20, 0x21, 0x8f, 0xd4, 0xc4, 0x14, 0xc7, 0xbc, 0x94, 0x90, 0x8a, 0x62,
	0x21, 0x0d, 0x2e, 0xfe, 0xe0, 0x92, 0xc4, 0xa2, 0xd4, 0x12, 0x88, 0x44, 0x62, 0x71, 0x86, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0xba, 0xb0, 0x90, 0x0a, 0x17, 0x93, 0x53, 0x86, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0xb7, 0x91, 0x88, 0x1e, 0xdc, 0x16, 0x3d, 0xb8, 0x89, 0x41, 0x4c, 0x4e, 0x19,
	0x42, 0x42, 0x5c, 0x2c, 0x21, 0x15, 0x4e, 0xc5, 0x12, 0xcc, 0x0a, 0xcc, 0x1a, 0x3c, 0x41, 0x60,
	0xb6, 0x93, 0xec, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe3, 0xb1, 0x1c, 0x43, 0x14, 0xb7, 0x9e, 0x3e, 0xdc, 0x80, 0x24, 0x36, 0x30, 0x6d, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x2b, 0x4a, 0x66, 0x07, 0xd1, 0x00, 0x00, 0x00,
}

func (m *BlockHeadAndTxs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeadAndTxs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeadAndTxs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.TxBs) > 0 {
		for iNdEx := len(m.TxBs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TxBs[iNdEx])
			copy(dAtA[i:], m.TxBs[iNdEx])
			i = encodeVarintBlockheadandtxs(dAtA, i, uint64(len(m.TxBs[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Bh != nil {
		{
			size, err := m.Bh.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlockheadandtxs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.StaretBlockHash) > 0 {
		i -= len(m.StaretBlockHash)
		copy(dAtA[i:], m.StaretBlockHash)
		i = encodeVarintBlockheadandtxs(dAtA, i, uint64(len(m.StaretBlockHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlockheadandtxs(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlockheadandtxs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockHeadAndTxs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StaretBlockHash)
	if l > 0 {
		n += 1 + l + sovBlockheadandtxs(uint64(l))
	}
	if m.Bh != nil {
		l = m.Bh.Size()
		n += 1 + l + sovBlockheadandtxs(uint64(l))
	}
	if len(m.TxBs) > 0 {
		for _, b := range m.TxBs {
			l = len(b)
			n += 1 + l + sovBlockheadandtxs(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBlockheadandtxs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlockheadandtxs(x uint64) (n int) {
	return sovBlockheadandtxs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockHeadAndTxs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockheadandtxs
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
			return fmt.Errorf("proto: BlockHeadAndTxs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeadAndTxs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StaretBlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockheadandtxs
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
				return ErrInvalidLengthBlockheadandtxs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockheadandtxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StaretBlockHash = append(m.StaretBlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.StaretBlockHash == nil {
				m.StaretBlockHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bh", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockheadandtxs
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
				return ErrInvalidLengthBlockheadandtxs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlockheadandtxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bh == nil {
				m.Bh = &BlockHead{}
			}
			if err := m.Bh.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxBs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockheadandtxs
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
				return ErrInvalidLengthBlockheadandtxs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockheadandtxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxBs = append(m.TxBs, make([]byte, postIndex-iNdEx))
			copy(m.TxBs[len(m.TxBs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockheadandtxs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockheadandtxs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBlockheadandtxs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockheadandtxs
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
					return 0, ErrIntOverflowBlockheadandtxs
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
					return 0, ErrIntOverflowBlockheadandtxs
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
				return 0, ErrInvalidLengthBlockheadandtxs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlockheadandtxs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlockheadandtxs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlockheadandtxs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockheadandtxs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlockheadandtxs = fmt.Errorf("proto: unexpected end of group")
)
