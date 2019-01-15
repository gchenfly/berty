// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/p2p/push.proto

package p2p // import "berty.tech/core/api/p2p"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DevicePushType int32

const (
	DevicePushType_UnknownDevicePushType DevicePushType = 0
	DevicePushType_APNS                  DevicePushType = 1
	DevicePushType_FCM                   DevicePushType = 2
)

var DevicePushType_name = map[int32]string{
	0: "UnknownDevicePushType",
	1: "APNS",
	2: "FCM",
}
var DevicePushType_value = map[string]int32{
	"UnknownDevicePushType": 0,
	"APNS":                  1,
	"FCM":                   2,
}

func (x DevicePushType) String() string {
	return proto.EnumName(DevicePushType_name, int32(x))
}
func (DevicePushType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_push_557b63745c727bbd, []int{0}
}

type DevicePushToDecrypted struct {
	Nonce                []byte         `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	PushType             DevicePushType `protobuf:"varint,2,opt,name=push_type,json=pushType,proto3,enum=berty.p2p.DevicePushType" json:"push_type,omitempty"`
	PushId               []byte         `protobuf:"bytes,3,opt,name=push_id,json=pushId,proto3" json:"push_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DevicePushToDecrypted) Reset()         { *m = DevicePushToDecrypted{} }
func (m *DevicePushToDecrypted) String() string { return proto.CompactTextString(m) }
func (*DevicePushToDecrypted) ProtoMessage()    {}
func (*DevicePushToDecrypted) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_557b63745c727bbd, []int{0}
}
func (m *DevicePushToDecrypted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DevicePushToDecrypted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DevicePushToDecrypted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DevicePushToDecrypted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevicePushToDecrypted.Merge(dst, src)
}
func (m *DevicePushToDecrypted) XXX_Size() int {
	return m.Size()
}
func (m *DevicePushToDecrypted) XXX_DiscardUnknown() {
	xxx_messageInfo_DevicePushToDecrypted.DiscardUnknown(m)
}

var xxx_messageInfo_DevicePushToDecrypted proto.InternalMessageInfo

func (m *DevicePushToDecrypted) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *DevicePushToDecrypted) GetPushType() DevicePushType {
	if m != nil {
		return m.PushType
	}
	return DevicePushType_UnknownDevicePushType
}

func (m *DevicePushToDecrypted) GetPushId() []byte {
	if m != nil {
		return m.PushId
	}
	return nil
}

type DevicePushIdentifier struct {
	PackageID            string   `protobuf:"bytes,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	DeviceToken          string   `protobuf:"bytes,2,opt,name=device_token,json=deviceToken,proto3" json:"device_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DevicePushIdentifier) Reset()         { *m = DevicePushIdentifier{} }
func (m *DevicePushIdentifier) String() string { return proto.CompactTextString(m) }
func (*DevicePushIdentifier) ProtoMessage()    {}
func (*DevicePushIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_557b63745c727bbd, []int{1}
}
func (m *DevicePushIdentifier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DevicePushIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DevicePushIdentifier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DevicePushIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevicePushIdentifier.Merge(dst, src)
}
func (m *DevicePushIdentifier) XXX_Size() int {
	return m.Size()
}
func (m *DevicePushIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_DevicePushIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_DevicePushIdentifier proto.InternalMessageInfo

func (m *DevicePushIdentifier) GetPackageID() string {
	if m != nil {
		return m.PackageID
	}
	return ""
}

func (m *DevicePushIdentifier) GetDeviceToken() string {
	if m != nil {
		return m.DeviceToken
	}
	return ""
}

func init() {
	proto.RegisterType((*DevicePushToDecrypted)(nil), "berty.p2p.DevicePushToDecrypted")
	proto.RegisterType((*DevicePushIdentifier)(nil), "berty.p2p.DevicePushIdentifier")
	proto.RegisterEnum("berty.p2p.DevicePushType", DevicePushType_name, DevicePushType_value)
}
func (m *DevicePushToDecrypted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevicePushToDecrypted) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Nonce) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPush(dAtA, i, uint64(len(m.Nonce)))
		i += copy(dAtA[i:], m.Nonce)
	}
	if m.PushType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPush(dAtA, i, uint64(m.PushType))
	}
	if len(m.PushId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPush(dAtA, i, uint64(len(m.PushId)))
		i += copy(dAtA[i:], m.PushId)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DevicePushIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevicePushIdentifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PackageID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPush(dAtA, i, uint64(len(m.PackageID)))
		i += copy(dAtA[i:], m.PackageID)
	}
	if len(m.DeviceToken) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPush(dAtA, i, uint64(len(m.DeviceToken)))
		i += copy(dAtA[i:], m.DeviceToken)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPush(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DevicePushToDecrypted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovPush(uint64(l))
	}
	if m.PushType != 0 {
		n += 1 + sovPush(uint64(m.PushType))
	}
	l = len(m.PushId)
	if l > 0 {
		n += 1 + l + sovPush(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DevicePushIdentifier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PackageID)
	if l > 0 {
		n += 1 + l + sovPush(uint64(l))
	}
	l = len(m.DeviceToken)
	if l > 0 {
		n += 1 + l + sovPush(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPush(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPush(x uint64) (n int) {
	return sovPush(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DevicePushToDecrypted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPush
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DevicePushToDecrypted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevicePushToDecrypted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPush
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushType", wireType)
			}
			m.PushType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PushType |= (DevicePushType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPush
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PushId = append(m.PushId[:0], dAtA[iNdEx:postIndex]...)
			if m.PushId == nil {
				m.PushId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPush(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPush
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
func (m *DevicePushIdentifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPush
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DevicePushIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevicePushIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackageID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPush
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PackageID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPush
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPush(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPush
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
func skipPush(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPush
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
					return 0, ErrIntOverflowPush
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPush
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPush
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPush
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPush(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPush = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPush   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/p2p/push.proto", fileDescriptor_push_557b63745c727bbd) }

var fileDescriptor_push_557b63745c727bbd = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0x59, 0xf8, 0xff, 0x81, 0x8e, 0x48, 0xc8, 0x06, 0x02, 0x78, 0xa8, 0xc8, 0x09, 0x8d,
	0x69, 0x13, 0x4c, 0x3c, 0x9a, 0x88, 0xc4, 0xa4, 0x07, 0x0d, 0xa9, 0x78, 0xf1, 0x42, 0xa0, 0x1d,
	0x4b, 0x43, 0xb2, 0x3b, 0x29, 0x8b, 0xa6, 0x17, 0x9f, 0xc3, 0x47, 0xf2, 0xe8, 0x13, 0x18, 0x53,
	0x5f, 0xc4, 0xec, 0x96, 0xa8, 0xdc, 0x66, 0xbe, 0x9d, 0xef, 0xb7, 0xdf, 0x0c, 0xf0, 0x39, 0xc5,
	0x2e, 0x0d, 0xc9, 0xa5, 0xcd, 0x7a, 0xe9, 0x50, 0x22, 0x95, 0xe4, 0xd6, 0x02, 0x13, 0x95, 0x3a,
	0x34, 0xa4, 0x83, 0x66, 0x24, 0x23, 0x69, 0x54, 0x57, 0x57, 0xf9, 0x40, 0xff, 0x05, 0x5a, 0x63,
	0x7c, 0x8a, 0x03, 0x9c, 0x6c, 0xd6, 0xcb, 0xa9, 0x1c, 0x63, 0x90, 0xa4, 0xa4, 0x30, 0xe4, 0x4d,
	0xf8, 0x2f, 0xa4, 0x08, 0xb0, 0xc3, 0x7a, 0x6c, 0x50, 0xf3, 0xf3, 0x86, 0x9f, 0x83, 0xa5, 0xe9,
	0x33, 0x95, 0x12, 0x76, 0x8a, 0x3d, 0x36, 0xa8, 0x0f, 0xbb, 0xce, 0xcf, 0x1f, 0xce, 0x1f, 0x54,
	0x4a, 0xe8, 0x57, 0x69, 0x5b, 0xf1, 0x36, 0x54, 0x8c, 0x2f, 0x0e, 0x3b, 0x25, 0xc3, 0x2b, 0xeb,
	0xd6, 0x0b, 0xfb, 0x11, 0x34, 0x7f, 0x4d, 0x5e, 0x88, 0x42, 0xc5, 0x8f, 0x31, 0x26, 0xfc, 0x14,
	0x80, 0xe6, 0xc1, 0x6a, 0x1e, 0xa1, 0xf6, 0xe8, 0x0c, 0xd6, 0x68, 0x3f, 0xfb, 0x38, 0xb4, 0x26,
	0xb9, 0xea, 0x8d, 0x7d, 0x6b, 0x3b, 0xe0, 0x85, 0xfc, 0x08, 0x6a, 0xa1, 0xa1, 0xcc, 0x94, 0x5c,
	0xa1, 0x30, 0xc9, 0x2c, 0x7f, 0x2f, 0xd7, 0xa6, 0x5a, 0x3a, 0xb9, 0x80, 0xfa, 0x6e, 0x3a, 0xde,
	0x85, 0xd6, 0xbd, 0x58, 0x09, 0xf9, 0x2c, 0x76, 0x1f, 0x1a, 0x05, 0x5e, 0x85, 0x7f, 0x97, 0x93,
	0xdb, 0xbb, 0x06, 0xe3, 0x15, 0x28, 0x5d, 0x5f, 0xdd, 0x34, 0x8a, 0xa3, 0xe3, 0xb7, 0xcc, 0x66,
	0xef, 0x99, 0xcd, 0x3e, 0x33, 0x9b, 0xbd, 0x7e, 0xd9, 0x85, 0x87, 0x76, 0xbe, 0xb7, 0xc2, 0x60,
	0xe9, 0x06, 0x32, 0x41, 0x77, 0x7b, 0xff, 0x45, 0xd9, 0x9c, 0xf6, 0xec, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x35, 0x55, 0x60, 0xcd, 0x91, 0x01, 0x00, 0x00,
}
