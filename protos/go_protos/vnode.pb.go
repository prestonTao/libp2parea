// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vnode.proto

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

type Vnodeinfo struct {
	Nid                  []byte   `protobuf:"bytes,1,opt,name=Nid,proto3" json:"Nid,omitempty"`
	Index                uint64   `protobuf:"varint,2,opt,name=Index,proto3" json:"Index,omitempty"`
	Vid                  []byte   `protobuf:"bytes,3,opt,name=Vid,proto3" json:"Vid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vnodeinfo) Reset()         { *m = Vnodeinfo{} }
func (m *Vnodeinfo) String() string { return proto.CompactTextString(m) }
func (*Vnodeinfo) ProtoMessage()    {}
func (*Vnodeinfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd943148590b397, []int{0}
}
func (m *Vnodeinfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vnodeinfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vnodeinfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vnodeinfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vnodeinfo.Merge(m, src)
}
func (m *Vnodeinfo) XXX_Size() int {
	return m.Size()
}
func (m *Vnodeinfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Vnodeinfo.DiscardUnknown(m)
}

var xxx_messageInfo_Vnodeinfo proto.InternalMessageInfo

func (m *Vnodeinfo) GetNid() []byte {
	if m != nil {
		return m.Nid
	}
	return nil
}

func (m *Vnodeinfo) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Vnodeinfo) GetVid() []byte {
	if m != nil {
		return m.Vid
	}
	return nil
}

type FindVnodeVO struct {
	Self                 *Vnodeinfo `protobuf:"bytes,1,opt,name=Self,proto3" json:"Self,omitempty"`
	Target               *Vnodeinfo `protobuf:"bytes,2,opt,name=Target,proto3" json:"Target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindVnodeVO) Reset()         { *m = FindVnodeVO{} }
func (m *FindVnodeVO) String() string { return proto.CompactTextString(m) }
func (*FindVnodeVO) ProtoMessage()    {}
func (*FindVnodeVO) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd943148590b397, []int{1}
}
func (m *FindVnodeVO) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FindVnodeVO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FindVnodeVO.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FindVnodeVO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindVnodeVO.Merge(m, src)
}
func (m *FindVnodeVO) XXX_Size() int {
	return m.Size()
}
func (m *FindVnodeVO) XXX_DiscardUnknown() {
	xxx_messageInfo_FindVnodeVO.DiscardUnknown(m)
}

var xxx_messageInfo_FindVnodeVO proto.InternalMessageInfo

func (m *FindVnodeVO) GetSelf() *Vnodeinfo {
	if m != nil {
		return m.Self
	}
	return nil
}

func (m *FindVnodeVO) GetTarget() *Vnodeinfo {
	if m != nil {
		return m.Target
	}
	return nil
}

type VnodeinfoRepeated struct {
	Vnodes               []*Vnodeinfo `protobuf:"bytes,1,rep,name=vnodes,proto3" json:"vnodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *VnodeinfoRepeated) Reset()         { *m = VnodeinfoRepeated{} }
func (m *VnodeinfoRepeated) String() string { return proto.CompactTextString(m) }
func (*VnodeinfoRepeated) ProtoMessage()    {}
func (*VnodeinfoRepeated) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd943148590b397, []int{2}
}
func (m *VnodeinfoRepeated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VnodeinfoRepeated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VnodeinfoRepeated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VnodeinfoRepeated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VnodeinfoRepeated.Merge(m, src)
}
func (m *VnodeinfoRepeated) XXX_Size() int {
	return m.Size()
}
func (m *VnodeinfoRepeated) XXX_DiscardUnknown() {
	xxx_messageInfo_VnodeinfoRepeated.DiscardUnknown(m)
}

var xxx_messageInfo_VnodeinfoRepeated proto.InternalMessageInfo

func (m *VnodeinfoRepeated) GetVnodes() []*Vnodeinfo {
	if m != nil {
		return m.Vnodes
	}
	return nil
}

func init() {
	proto.RegisterType((*Vnodeinfo)(nil), "go_protos.Vnodeinfo")
	proto.RegisterType((*FindVnodeVO)(nil), "go_protos.FindVnodeVO")
	proto.RegisterType((*VnodeinfoRepeated)(nil), "go_protos.VnodeinfoRepeated")
}

func init() { proto.RegisterFile("vnode.proto", fileDescriptor_4fd943148590b397) }

var fileDescriptor_4fd943148590b397 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xcb, 0xcb, 0x4f,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xcf, 0x8f, 0x07, 0xb3, 0x8a, 0x95,
	0x5c, 0xb9, 0x38, 0xc3, 0x40, 0x32, 0x99, 0x79, 0x69, 0xf9, 0x42, 0x02, 0x5c, 0xcc, 0x7e, 0x99,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x20, 0xa6, 0x90, 0x08, 0x17, 0xab, 0x67, 0x5e,
	0x4a, 0x6a, 0x85, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x84, 0x03, 0x52, 0x17, 0x96, 0x99,
	0x22, 0xc1, 0x0c, 0x51, 0x17, 0x96, 0x99, 0xa2, 0x94, 0xca, 0xc5, 0xed, 0x96, 0x99, 0x97, 0x02,
	0x36, 0x2a, 0xcc, 0x5f, 0x48, 0x83, 0x8b, 0x25, 0x38, 0x35, 0x27, 0x0d, 0x6c, 0x12, 0xb7, 0x91,
	0x88, 0x1e, 0xdc, 0x3e, 0x3d, 0xb8, 0x65, 0x41, 0x60, 0x15, 0x42, 0x3a, 0x5c, 0x6c, 0x21, 0x89,
	0x45, 0xe9, 0xa9, 0x25, 0x60, 0x1b, 0x70, 0xa9, 0x85, 0xaa, 0x51, 0x72, 0xe4, 0x12, 0x44, 0x08,
	0xa6, 0x16, 0xa4, 0x26, 0x96, 0xa4, 0xa6, 0x80, 0x8c, 0x00, 0x7b, 0xae, 0x58, 0x82, 0x51, 0x81,
	0x19, 0xb7, 0x11, 0x10, 0x35, 0x4e, 0xb2, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x51, 0xdc, 0x7a, 0xfa, 0x70, 0x0d, 0x49, 0x6c,
	0x60, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x61, 0x22, 0xcc, 0x30, 0x01, 0x00, 0x00,
}

func (m *Vnodeinfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vnodeinfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vnodeinfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Vid) > 0 {
		i -= len(m.Vid)
		copy(dAtA[i:], m.Vid)
		i = encodeVarintVnode(dAtA, i, uint64(len(m.Vid)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Index != 0 {
		i = encodeVarintVnode(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Nid) > 0 {
		i -= len(m.Nid)
		copy(dAtA[i:], m.Nid)
		i = encodeVarintVnode(dAtA, i, uint64(len(m.Nid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FindVnodeVO) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FindVnodeVO) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FindVnodeVO) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Target != nil {
		{
			size, err := m.Target.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVnode(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Self != nil {
		{
			size, err := m.Self.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVnode(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VnodeinfoRepeated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VnodeinfoRepeated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VnodeinfoRepeated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Vnodes) > 0 {
		for iNdEx := len(m.Vnodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Vnodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVnode(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintVnode(dAtA []byte, offset int, v uint64) int {
	offset -= sovVnode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Vnodeinfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Nid)
	if l > 0 {
		n += 1 + l + sovVnode(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovVnode(uint64(m.Index))
	}
	l = len(m.Vid)
	if l > 0 {
		n += 1 + l + sovVnode(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FindVnodeVO) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Self != nil {
		l = m.Self.Size()
		n += 1 + l + sovVnode(uint64(l))
	}
	if m.Target != nil {
		l = m.Target.Size()
		n += 1 + l + sovVnode(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *VnodeinfoRepeated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Vnodes) > 0 {
		for _, e := range m.Vnodes {
			l = e.Size()
			n += 1 + l + sovVnode(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVnode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVnode(x uint64) (n int) {
	return sovVnode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vnodeinfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVnode
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
			return fmt.Errorf("proto: Vnodeinfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vnodeinfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVnode
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
				return ErrInvalidLengthVnode
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVnode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nid = append(m.Nid[:0], dAtA[iNdEx:postIndex]...)
			if m.Nid == nil {
				m.Nid = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVnode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVnode
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
				return ErrInvalidLengthVnode
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVnode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vid = append(m.Vid[:0], dAtA[iNdEx:postIndex]...)
			if m.Vid == nil {
				m.Vid = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVnode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVnode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVnode
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
func (m *FindVnodeVO) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVnode
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
			return fmt.Errorf("proto: FindVnodeVO: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FindVnodeVO: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Self", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVnode
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
				return ErrInvalidLengthVnode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVnode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Self == nil {
				m.Self = &Vnodeinfo{}
			}
			if err := m.Self.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVnode
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
				return ErrInvalidLengthVnode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVnode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Target == nil {
				m.Target = &Vnodeinfo{}
			}
			if err := m.Target.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVnode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVnode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVnode
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
func (m *VnodeinfoRepeated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVnode
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
			return fmt.Errorf("proto: VnodeinfoRepeated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VnodeinfoRepeated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vnodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVnode
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
				return ErrInvalidLengthVnode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVnode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vnodes = append(m.Vnodes, &Vnodeinfo{})
			if err := m.Vnodes[len(m.Vnodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVnode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVnode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVnode
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
func skipVnode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVnode
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
					return 0, ErrIntOverflowVnode
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
					return 0, ErrIntOverflowVnode
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
				return 0, ErrInvalidLengthVnode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVnode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVnode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVnode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVnode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVnode = fmt.Errorf("proto: unexpected end of group")
)
