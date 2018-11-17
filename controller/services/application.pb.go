// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: controller/services/application.proto

package services // import "github.com/argoproj/argo-cd/controller/services"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ResourcesQuery struct {
	ApplicationName      *string  `protobuf:"bytes,1,req,name=applicationName" json:"applicationName,omitempty"`
	Group                *string  `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	Version              *string  `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Kind                 *string  `protobuf:"bytes,4,opt,name=kind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourcesQuery) Reset()         { *m = ResourcesQuery{} }
func (m *ResourcesQuery) String() string { return proto.CompactTextString(m) }
func (*ResourcesQuery) ProtoMessage()    {}
func (*ResourcesQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_b7c966df2a6bd563, []int{0}
}
func (m *ResourcesQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourcesQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourcesQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ResourcesQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcesQuery.Merge(dst, src)
}
func (m *ResourcesQuery) XXX_Size() int {
	return m.Size()
}
func (m *ResourcesQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcesQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcesQuery proto.InternalMessageInfo

func (m *ResourcesQuery) GetApplicationName() string {
	if m != nil && m.ApplicationName != nil {
		return *m.ApplicationName
	}
	return ""
}

func (m *ResourcesQuery) GetGroup() string {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return ""
}

func (m *ResourcesQuery) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *ResourcesQuery) GetKind() string {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return ""
}

type ResourcesResponse struct {
	Items                []*v1alpha1.ResourceState `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ResourcesResponse) Reset()         { *m = ResourcesResponse{} }
func (m *ResourcesResponse) String() string { return proto.CompactTextString(m) }
func (*ResourcesResponse) ProtoMessage()    {}
func (*ResourcesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_b7c966df2a6bd563, []int{1}
}
func (m *ResourcesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourcesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourcesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ResourcesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcesResponse.Merge(dst, src)
}
func (m *ResourcesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ResourcesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcesResponse proto.InternalMessageInfo

func (m *ResourcesResponse) GetItems() []*v1alpha1.ResourceState {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourcesQuery)(nil), "github.com.argoproj.argo_cd.controller.services.ResourcesQuery")
	proto.RegisterType((*ResourcesResponse)(nil), "github.com.argoproj.argo_cd.controller.services.ResourcesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ApplicationService service

type ApplicationServiceClient interface {
	// Resources returns information about expected and observed application resources
	Resources(ctx context.Context, in *ResourcesQuery, opts ...grpc.CallOption) (*ResourcesResponse, error)
}

type applicationServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServiceClient(cc *grpc.ClientConn) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) Resources(ctx context.Context, in *ResourcesQuery, opts ...grpc.CallOption) (*ResourcesResponse, error) {
	out := new(ResourcesResponse)
	err := c.cc.Invoke(ctx, "/github.com.argoproj.argo_cd.controller.services.ApplicationService/Resources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApplicationService service

type ApplicationServiceServer interface {
	// Resources returns information about expected and observed application resources
	Resources(context.Context, *ResourcesQuery) (*ResourcesResponse, error)
}

func RegisterApplicationServiceServer(s *grpc.Server, srv ApplicationServiceServer) {
	s.RegisterService(&_ApplicationService_serviceDesc, srv)
}

func _ApplicationService_Resources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourcesQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Resources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.argoproj.argo_cd.controller.services.ApplicationService/Resources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Resources(ctx, req.(*ResourcesQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.argoproj.argo_cd.controller.services.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resources",
			Handler:    _ApplicationService_Resources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/services/application.proto",
}

func (m *ResourcesQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourcesQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ApplicationName == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("applicationName")
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApplication(dAtA, i, uint64(len(*m.ApplicationName)))
		i += copy(dAtA[i:], *m.ApplicationName)
	}
	if m.Group != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApplication(dAtA, i, uint64(len(*m.Group)))
		i += copy(dAtA[i:], *m.Group)
	}
	if m.Version != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintApplication(dAtA, i, uint64(len(*m.Version)))
		i += copy(dAtA[i:], *m.Version)
	}
	if m.Kind != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintApplication(dAtA, i, uint64(len(*m.Kind)))
		i += copy(dAtA[i:], *m.Kind)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ResourcesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourcesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintApplication(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintApplication(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResourcesQuery) Size() (n int) {
	var l int
	_ = l
	if m.ApplicationName != nil {
		l = len(*m.ApplicationName)
		n += 1 + l + sovApplication(uint64(l))
	}
	if m.Group != nil {
		l = len(*m.Group)
		n += 1 + l + sovApplication(uint64(l))
	}
	if m.Version != nil {
		l = len(*m.Version)
		n += 1 + l + sovApplication(uint64(l))
	}
	if m.Kind != nil {
		l = len(*m.Kind)
		n += 1 + l + sovApplication(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResourcesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovApplication(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApplication(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApplication(x uint64) (n int) {
	return sovApplication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourcesQuery) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
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
			return fmt.Errorf("proto: ResourcesQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourcesQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.ApplicationName = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Group = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Version = &s
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Kind = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplication
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("applicationName")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResourcesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
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
			return fmt.Errorf("proto: ResourcesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourcesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &v1alpha1.ResourceState{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplication
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
func skipApplication(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApplication
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
					return 0, ErrIntOverflowApplication
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
					return 0, ErrIntOverflowApplication
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
				return 0, ErrInvalidLengthApplication
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApplication
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
				next, err := skipApplication(dAtA[start:])
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
	ErrInvalidLengthApplication = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApplication   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("controller/services/application.proto", fileDescriptor_application_b7c966df2a6bd563)
}

var fileDescriptor_application_b7c966df2a6bd563 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcb, 0x4a, 0x03, 0x31,
	0x14, 0x35, 0x7d, 0x20, 0x8d, 0xa0, 0x18, 0x5c, 0x0c, 0x5d, 0x94, 0x52, 0x10, 0x66, 0x63, 0x42,
	0xbb, 0x17, 0xb1, 0x2b, 0xdd, 0x08, 0x4e, 0x77, 0x2e, 0x94, 0x98, 0xb9, 0x4c, 0x63, 0xa7, 0x73,
	0x43, 0x92, 0x29, 0xb8, 0xf1, 0x23, 0xfc, 0x0b, 0xff, 0xc4, 0xa5, 0x9f, 0x20, 0xfd, 0x12, 0x71,
	0x86, 0xf4, 0x21, 0x52, 0xd0, 0xdd, 0xb9, 0x87, 0x39, 0xe7, 0xcc, 0x3d, 0x37, 0xf4, 0x54, 0x61,
	0xe1, 0x2d, 0xe6, 0x39, 0x58, 0xe1, 0xc0, 0x2e, 0xb4, 0x02, 0x27, 0xa4, 0x31, 0xb9, 0x56, 0xd2,
	0x6b, 0x2c, 0xb8, 0xb1, 0xe8, 0x91, 0x89, 0x4c, 0xfb, 0x69, 0xf9, 0xc8, 0x15, 0xce, 0xb9, 0xb4,
	0x19, 0x1a, 0x8b, 0x4f, 0x15, 0x78, 0x50, 0x29, 0x5f, 0x5b, 0xf0, 0x60, 0xd1, 0xbd, 0x5e, 0x0b,
	0x44, 0x10, 0x54, 0xe0, 0x4c, 0xa5, 0xc2, 0xcc, 0x32, 0x21, 0x8d, 0xde, 0x0a, 0x12, 0x8b, 0xa1,
	0xcc, 0xcd, 0x54, 0x0e, 0x45, 0x06, 0x05, 0x58, 0xe9, 0x21, 0xad, 0xb3, 0x07, 0x2f, 0xf4, 0x30,
	0x01, 0x87, 0xa5, 0x55, 0xe0, 0x6e, 0x4b, 0xb0, 0xcf, 0x2c, 0xa6, 0x47, 0x1b, 0xca, 0x1b, 0x39,
	0x87, 0x88, 0xf4, 0x1b, 0x71, 0x27, 0xf9, 0x49, 0xb3, 0x13, 0xda, 0xce, 0x2c, 0x96, 0x26, 0x6a,
	0xf4, 0x49, 0xdc, 0x49, 0xea, 0x81, 0x45, 0x74, 0x7f, 0x01, 0xd6, 0x69, 0x2c, 0xa2, 0x66, 0xc5,
	0x87, 0x91, 0x31, 0xda, 0x9a, 0xe9, 0x22, 0x8d, 0x5a, 0x15, 0x5d, 0xe1, 0x81, 0xa3, 0xc7, 0xab,
	0xfc, 0x04, 0x9c, 0xc1, 0xc2, 0x01, 0xbb, 0xa7, 0x6d, 0xed, 0x61, 0xee, 0x22, 0xd2, 0x6f, 0xc6,
	0x07, 0xa3, 0x2b, 0xbe, 0xab, 0x20, 0x33, 0xcb, 0xf8, 0xf7, 0xbe, 0x7c, 0xb3, 0xd8, 0xb0, 0x2f,
	0x0f, 0xe6, 0x13, 0x2f, 0x3d, 0x24, 0xb5, 0xed, 0xe8, 0x8d, 0x50, 0x76, 0xb9, 0xfe, 0x7a, 0x52,
	0xf7, 0xca, 0x5e, 0x09, 0xed, 0xac, 0x7e, 0x86, 0x5d, 0xf0, 0x3f, 0x9e, 0x85, 0x6f, 0x17, 0xd9,
	0x1d, 0xff, 0xdf, 0x20, 0x34, 0x31, 0xd8, 0x1b, 0x9f, 0xbf, 0x2f, 0x7b, 0xe4, 0x63, 0xd9, 0x23,
	0x9f, 0xcb, 0x1e, 0xb9, 0x13, 0xbb, 0x2e, 0xff, 0xcb, 0x6b, 0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x8d, 0xcd, 0xc9, 0xf5, 0x83, 0x02, 0x00, 0x00,
}