// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: actor.proto

/*
	Package protos is a generated protocol buffer package.

	It is generated from these files:
		actor.proto
		group.proto
		permission.proto
		permission_service.proto
		role.proto
		role_service.proto

	It has these top-level messages:
		Actor
		Group
		Permission
		HasPermissionRequest
		HasPermissionResponse
		ListResourcePatternsRequest
		ListResourcePatternsResponse
		Role
		CreateRoleRequest
		CreateRoleResponse
		DeleteRoleRequest
		DeleteRoleResponse
		AssignRoleRequest
		AssignRoleResponse
		AssignRoleToGroupRequest
		AssignRoleToGroupResponse
		UnassignRoleFromGroupRequest
		UnassignRoleFromGroupResponse
		UnassignRoleRequest
		UnassignRoleResponse
		HasRoleRequest
		HasRoleResponse
		HasRoleForGroupRequest
		HasRoleForGroupResponse
		ListActorRolesRequest
		ListActorRolesResponse
		ListRolePermissionsRequest
		ListRolePermissionsResponse
*/
package protos

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

type Actor struct {
	ID        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *Actor) Reset()                    { *m = Actor{} }
func (m *Actor) String() string            { return proto.CompactTextString(m) }
func (*Actor) ProtoMessage()               {}
func (*Actor) Descriptor() ([]byte, []int) { return fileDescriptorActor, []int{0} }

func (m *Actor) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Actor) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*Actor)(nil), "cloud_foundry.perm.protos.Actor")
}
func (m *Actor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Actor) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintActor(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Namespace) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintActor(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	return i, nil
}

func encodeVarintActor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Actor) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovActor(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovActor(uint64(l))
	}
	return n
}

func sovActor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozActor(x uint64) (n int) {
	return sovActor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Actor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActor
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
			return fmt.Errorf("proto: Actor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Actor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActor
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
				return ErrInvalidLengthActor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActor
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
				return ErrInvalidLengthActor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthActor
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
func skipActor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActor
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
					return 0, ErrIntOverflowActor
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
					return 0, ErrIntOverflowActor
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
				return 0, ErrInvalidLengthActor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowActor
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
				next, err := skipActor(dAtA[start:])
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
	ErrInvalidLengthActor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActor   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("actor.proto", fileDescriptorActor) }

var fileDescriptorActor = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4c, 0x2e, 0xc9,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4c, 0xce, 0xc9, 0x2f, 0x4d, 0x89, 0x4f,
	0xcb, 0x2f, 0xcd, 0x4b, 0x29, 0xaa, 0xd4, 0x2b, 0x48, 0x2d, 0xca, 0x85, 0xc8, 0x14, 0x4b, 0xe9,
	0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb,
	0x83, 0xc5, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0xa8, 0x57, 0x72, 0xe6, 0x62,
	0x75, 0x04, 0x19, 0x2c, 0x24, 0xc6, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe9,
	0xc4, 0xf6, 0xe8, 0x9e, 0x3c, 0x93, 0xa7, 0x4b, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x0c, 0x17, 0x67,
	0x5e, 0x62, 0x6e, 0x6a, 0x71, 0x41, 0x62, 0x72, 0xaa, 0x04, 0x33, 0x48, 0x3a, 0x08, 0x21, 0xe0,
	0xc5, 0xc2, 0xc1, 0x28, 0xc0, 0xe4, 0x24, 0x71, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x10, 0xc5, 0x06, 0x71, 0x4d, 0x12, 0x84, 0x36,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xd5, 0x31, 0x41, 0xbe, 0x00, 0x00, 0x00,
}
