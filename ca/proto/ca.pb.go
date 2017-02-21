// Code generated by protoc-gen-go.
// source: ca/proto/ca.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	ca/proto/ca.proto

It has these top-level messages:
	IssueCertificateRequest
	GenerateOCSPRequest
	OCSPResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/letsencrypt/boulder/core/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type IssueCertificateRequest struct {
	Csr              []byte `protobuf:"bytes,1,opt,name=csr" json:"csr,omitempty"`
	RegistrationID   *int64 `protobuf:"varint,2,opt,name=registrationID" json:"registrationID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IssueCertificateRequest) Reset()                    { *m = IssueCertificateRequest{} }
func (m *IssueCertificateRequest) String() string            { return proto1.CompactTextString(m) }
func (*IssueCertificateRequest) ProtoMessage()               {}
func (*IssueCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IssueCertificateRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *IssueCertificateRequest) GetRegistrationID() int64 {
	if m != nil && m.RegistrationID != nil {
		return *m.RegistrationID
	}
	return 0
}

type GenerateOCSPRequest struct {
	CertDER          []byte  `protobuf:"bytes,1,opt,name=certDER" json:"certDER,omitempty"`
	Status           *string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Reason           *int32  `protobuf:"varint,3,opt,name=reason" json:"reason,omitempty"`
	RevokedAt        *int64  `protobuf:"varint,4,opt,name=revokedAt" json:"revokedAt,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GenerateOCSPRequest) Reset()                    { *m = GenerateOCSPRequest{} }
func (m *GenerateOCSPRequest) String() string            { return proto1.CompactTextString(m) }
func (*GenerateOCSPRequest) ProtoMessage()               {}
func (*GenerateOCSPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GenerateOCSPRequest) GetCertDER() []byte {
	if m != nil {
		return m.CertDER
	}
	return nil
}

func (m *GenerateOCSPRequest) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *GenerateOCSPRequest) GetReason() int32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

func (m *GenerateOCSPRequest) GetRevokedAt() int64 {
	if m != nil && m.RevokedAt != nil {
		return *m.RevokedAt
	}
	return 0
}

type OCSPResponse struct {
	Response         []byte `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *OCSPResponse) Reset()                    { *m = OCSPResponse{} }
func (m *OCSPResponse) String() string            { return proto1.CompactTextString(m) }
func (*OCSPResponse) ProtoMessage()               {}
func (*OCSPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OCSPResponse) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto1.RegisterType((*IssueCertificateRequest)(nil), "ca.IssueCertificateRequest")
	proto1.RegisterType((*GenerateOCSPRequest)(nil), "ca.GenerateOCSPRequest")
	proto1.RegisterType((*OCSPResponse)(nil), "ca.OCSPResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for CertificateAuthority service

type CertificateAuthorityClient interface {
	IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*core.Certificate, error)
}

type certificateAuthorityClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAuthorityClient(cc *grpc.ClientConn) CertificateAuthorityClient {
	return &certificateAuthorityClient{cc}
}

func (c *certificateAuthorityClient) IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*core.Certificate, error) {
	out := new(core.Certificate)
	err := grpc.Invoke(ctx, "/ca.CertificateAuthority/IssueCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateAuthority service

type CertificateAuthorityServer interface {
	IssueCertificate(context.Context, *IssueCertificateRequest) (*core.Certificate, error)
}

func RegisterCertificateAuthorityServer(s *grpc.Server, srv CertificateAuthorityServer) {
	s.RegisterService(&_CertificateAuthority_serviceDesc, srv)
}

func _CertificateAuthority_IssueCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).IssueCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.CertificateAuthority/IssueCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).IssueCertificate(ctx, req.(*IssueCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAuthority_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ca.CertificateAuthority",
	HandlerType: (*CertificateAuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueCertificate",
			Handler:    _CertificateAuthority_IssueCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

// Client API for OCSPGenerator service

type OCSPGeneratorClient interface {
	GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error)
}

type oCSPGeneratorClient struct {
	cc *grpc.ClientConn
}

func NewOCSPGeneratorClient(cc *grpc.ClientConn) OCSPGeneratorClient {
	return &oCSPGeneratorClient{cc}
}

func (c *oCSPGeneratorClient) GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error) {
	out := new(OCSPResponse)
	err := grpc.Invoke(ctx, "/ca.OCSPGenerator/GenerateOCSP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OCSPGenerator service

type OCSPGeneratorServer interface {
	GenerateOCSP(context.Context, *GenerateOCSPRequest) (*OCSPResponse, error)
}

func RegisterOCSPGeneratorServer(s *grpc.Server, srv OCSPGeneratorServer) {
	s.RegisterService(&_OCSPGenerator_serviceDesc, srv)
}

func _OCSPGenerator_GenerateOCSP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateOCSPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCSPGeneratorServer).GenerateOCSP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.OCSPGenerator/GenerateOCSP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCSPGeneratorServer).GenerateOCSP(ctx, req.(*GenerateOCSPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OCSPGenerator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ca.OCSPGenerator",
	HandlerType: (*OCSPGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateOCSP",
			Handler:    _OCSPGenerator_GenerateOCSP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto1.RegisterFile("ca/proto/ca.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x9b, 0xe6, 0xed, 0x5b, 0x3b, 0xc6, 0x9a, 0xac, 0x1f, 0x0d, 0xf1, 0x12, 0x72, 0xca,
	0x29, 0x85, 0x5e, 0x05, 0xa1, 0x36, 0x22, 0x05, 0x41, 0xa9, 0x27, 0xc5, 0xcb, 0xb2, 0x8e, 0x1a,
	0x84, 0x6c, 0x9d, 0x99, 0x08, 0xfe, 0xf7, 0x92, 0x8f, 0x42, 0x10, 0xbd, 0xcd, 0xf2, 0xf0, 0xfc,
	0x76, 0x7e, 0x03, 0x81, 0xd1, 0xf3, 0x2d, 0x59, 0xb1, 0x73, 0xa3, 0xb3, 0x66, 0x50, 0x43, 0xa3,
	0xa3, 0x13, 0x63, 0x09, 0x77, 0x81, 0x25, 0x6c, 0xa3, 0xe4, 0x02, 0x66, 0x6b, 0xe6, 0x0a, 0x57,
	0x48, 0x52, 0xbc, 0x14, 0x46, 0x0b, 0x6e, 0xf0, 0xa3, 0x42, 0x16, 0xb5, 0x0f, 0xae, 0x61, 0x0a,
	0x9d, 0xd8, 0x49, 0x3d, 0x75, 0x0a, 0x53, 0xc2, 0xd7, 0x82, 0x85, 0xb4, 0x14, 0xb6, 0x5c, 0xe7,
	0xe1, 0x30, 0x76, 0x52, 0x37, 0x79, 0x80, 0xa3, 0x6b, 0x2c, 0x91, 0xb4, 0xe0, 0xed, 0xea, 0xfe,
	0x6e, 0xd7, 0x3d, 0x84, 0xb1, 0x41, 0x92, 0xfc, 0x6a, 0xd3, 0xf5, 0xa7, 0xf0, 0x9f, 0x45, 0x4b,
	0xc5, 0x4d, 0x6f, 0x52, 0xbf, 0x09, 0x35, 0xdb, 0x32, 0x74, 0x63, 0x27, 0x1d, 0xa9, 0x00, 0x26,
	0x84, 0x9f, 0xf6, 0x1d, 0x9f, 0x97, 0x12, 0xfe, 0x6b, 0xd0, 0x31, 0x78, 0x2d, 0x92, 0xb7, 0xb6,
	0x64, 0x54, 0x3e, 0xec, 0x51, 0x37, 0xb7, 0xd0, 0xc5, 0x13, 0x1c, 0xf7, 0xf6, 0x5e, 0x56, 0xf2,
	0x66, 0xa9, 0x90, 0x2f, 0x95, 0x83, 0xff, 0x53, 0x4a, 0x9d, 0x65, 0x46, 0x67, 0x7f, 0xa8, 0x46,
	0x41, 0xd6, 0x9c, 0xa4, 0x97, 0x24, 0x83, 0xc5, 0x0d, 0x1c, 0xd4, 0xff, 0x77, 0x7a, 0x96, 0xd4,
	0x39, 0x78, 0x7d, 0x57, 0x35, 0xab, 0x91, 0xbf, 0xd8, 0x47, 0x7e, 0x1d, 0xf4, 0x77, 0x4f, 0x06,
	0x97, 0xe3, 0xc7, 0x51, 0x73, 0xf1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x60, 0xcf, 0xa9,
	0xa0, 0x01, 0x00, 0x00,
}
