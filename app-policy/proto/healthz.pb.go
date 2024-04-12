// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: healthz.proto

package proto

import (
	fmt "fmt"

	proto1 "github.com/gogo/protobuf/proto"

	math "math"

	context "context"

	grpc "google.golang.org/grpc"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HealthCheckRequest struct {
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto1.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptorHealthz, []int{0} }

type HealthCheckResponse struct {
	Healthy bool `protobuf:"varint,1,opt,name=healthy,proto3" json:"healthy,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto1.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptorHealthz, []int{1} }

func (m *HealthCheckResponse) GetHealthy() bool {
	if m != nil {
		return m.Healthy
	}
	return false
}

func init() {
	proto1.RegisterType((*HealthCheckRequest)(nil), "dikastes.HealthCheckRequest")
	proto1.RegisterType((*HealthCheckResponse)(nil), "dikastes.HealthCheckResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Healthz service

type HealthzClient interface {
	CheckReadiness(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	CheckLiveness(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthzClient struct {
	cc *grpc.ClientConn
}

func NewHealthzClient(cc *grpc.ClientConn) HealthzClient {
	return &healthzClient{cc}
}

func (c *healthzClient) CheckReadiness(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/dikastes.Healthz/CheckReadiness", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthzClient) CheckLiveness(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/dikastes.Healthz/CheckLiveness", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Healthz service

type HealthzServer interface {
	CheckReadiness(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	CheckLiveness(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthzServer(s *grpc.Server, srv HealthzServer) {
	s.RegisterService(&_Healthz_serviceDesc, srv)
}

func _Healthz_CheckReadiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthzServer).CheckReadiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dikastes.Healthz/CheckReadiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthzServer).CheckReadiness(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Healthz_CheckLiveness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthzServer).CheckLiveness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dikastes.Healthz/CheckLiveness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthzServer).CheckLiveness(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Healthz_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dikastes.Healthz",
	HandlerType: (*HealthzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckReadiness",
			Handler:    _Healthz_CheckReadiness_Handler,
		},
		{
			MethodName: "CheckLiveness",
			Handler:    _Healthz_CheckLiveness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "healthz.proto",
}

func (m *HealthCheckRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheckRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *HealthCheckResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheckResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Healthy {
		dAtA[i] = 0x8
		i++
		if m.Healthy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintHealthz(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HealthCheckRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *HealthCheckResponse) Size() (n int) {
	var l int
	_ = l
	if m.Healthy {
		n += 2
	}
	return n
}

func sovHealthz(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHealthz(x uint64) (n int) {
	return sovHealthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HealthCheckRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthz
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
			return fmt.Errorf("proto: HealthCheckRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheckRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHealthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthz
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
func (m *HealthCheckResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthz
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
			return fmt.Errorf("proto: HealthCheckResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheckResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Healthy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Healthy = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipHealthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthz
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
func skipHealthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHealthz
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
					return 0, ErrIntOverflowHealthz
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
					return 0, ErrIntOverflowHealthz
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
				return 0, ErrInvalidLengthHealthz
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHealthz
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
				next, err := skipHealthz(dAtA[start:])
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
	ErrInvalidLengthHealthz = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHealthz   = fmt.Errorf("proto: integer overflow")
)

func init() { proto1.RegisterFile("healthz.proto", fileDescriptorHealthz) }

var fileDescriptorHealthz = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xa8, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xc9, 0xcc, 0x4e, 0x2c,
	0x2e, 0x49, 0x2d, 0x56, 0x12, 0xe1, 0x12, 0xf2, 0x00, 0x4b, 0x39, 0x67, 0xa4, 0x26, 0x67, 0x07,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0xe9, 0x73, 0x09, 0xa3, 0x88, 0x16, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0x43, 0xcc, 0xa9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08,
	0x82, 0x71, 0x8d, 0x96, 0x31, 0x72, 0xb1, 0x43, 0x74, 0x54, 0x09, 0xf9, 0x72, 0xf1, 0x41, 0xb5,
	0x25, 0xa6, 0x64, 0xe6, 0xa5, 0x16, 0x17, 0x0b, 0xc9, 0xe8, 0xc1, 0xec, 0xd3, 0xc3, 0xb4, 0x4c,
	0x4a, 0x16, 0x87, 0x2c, 0xd4, 0x52, 0x1f, 0x2e, 0x5e, 0xb0, 0x80, 0x4f, 0x66, 0x59, 0x2a, 0xc5,
	0xa6, 0x39, 0x09, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x51, 0xac, 0xe0, 0xe0, 0x48, 0x62, 0x03, 0x53, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90,
	0xe9, 0xfa, 0x7b, 0x26, 0x01, 0x00, 0x00,
}
