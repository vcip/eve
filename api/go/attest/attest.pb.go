// Code generated by protoc-gen-go. DO NOT EDIT.
// source: attest.proto

package certs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ZAttestResponseCode int32

const (
	ZAttestResponseCode_ATTEST_RESPONSE_SUCCESS ZAttestResponseCode = 0
	ZAttestResponseCode_ATTEST_RESPONSE_FAILURE ZAttestResponseCode = 1
)

var ZAttestResponseCode_name = map[int32]string{
	0: "ATTEST_RESPONSE_SUCCESS",
	1: "ATTEST_RESPONSE_FAILURE",
}

var ZAttestResponseCode_value = map[string]int32{
	"ATTEST_RESPONSE_SUCCESS": 0,
	"ATTEST_RESPONSE_FAILURE": 1,
}

func (x ZAttestResponseCode) String() string {
	return proto.EnumName(ZAttestResponseCode_name, int32(x))
}

func (ZAttestResponseCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_208cf26448842a3f, []int{0}
}

type ZAttestNonceResp struct {
	Nonce                []byte   `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZAttestNonceResp) Reset()         { *m = ZAttestNonceResp{} }
func (m *ZAttestNonceResp) String() string { return proto.CompactTextString(m) }
func (*ZAttestNonceResp) ProtoMessage()    {}
func (*ZAttestNonceResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_208cf26448842a3f, []int{0}
}

func (m *ZAttestNonceResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZAttestNonceResp.Unmarshal(m, b)
}
func (m *ZAttestNonceResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZAttestNonceResp.Marshal(b, m, deterministic)
}
func (m *ZAttestNonceResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZAttestNonceResp.Merge(m, src)
}
func (m *ZAttestNonceResp) XXX_Size() int {
	return xxx_messageInfo_ZAttestNonceResp.Size(m)
}
func (m *ZAttestNonceResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ZAttestNonceResp.DiscardUnknown(m)
}

var xxx_messageInfo_ZAttestNonceResp proto.InternalMessageInfo

func (m *ZAttestNonceResp) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// This is the request payload for POST /api/v1/edgeDevice/attestQuote
// or /api/v2/edgeDevice/attestQuote
// The message is assumed to be protected by a TLS session bound to the
// device certificate for v1.
// The message is assumed to be protected by signing envelope for v2.
type ZAttestQuoteReq struct {
	AttestData []byte `protobuf:"bytes,1,opt,name=attestData,proto3" json:"attestData,omitempty"`
	//nonce is included in attestData
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZAttestQuoteReq) Reset()         { *m = ZAttestQuoteReq{} }
func (m *ZAttestQuoteReq) String() string { return proto.CompactTextString(m) }
func (*ZAttestQuoteReq) ProtoMessage()    {}
func (*ZAttestQuoteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_208cf26448842a3f, []int{1}
}

func (m *ZAttestQuoteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZAttestQuoteReq.Unmarshal(m, b)
}
func (m *ZAttestQuoteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZAttestQuoteReq.Marshal(b, m, deterministic)
}
func (m *ZAttestQuoteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZAttestQuoteReq.Merge(m, src)
}
func (m *ZAttestQuoteReq) XXX_Size() int {
	return xxx_messageInfo_ZAttestQuoteReq.Size(m)
}
func (m *ZAttestQuoteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ZAttestQuoteReq.DiscardUnknown(m)
}

var xxx_messageInfo_ZAttestQuoteReq proto.InternalMessageInfo

func (m *ZAttestQuoteReq) GetAttestData() []byte {
	if m != nil {
		return m.AttestData
	}
	return nil
}

func (m *ZAttestQuoteReq) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// This is the response payload for POST /api/v1/edgeDevice/attestQuote
// or /api/v2/edgeDevice/attestQuote
// The message is assumed to be protected by a TLS session bound to the
// device certificate for v1.
// The message is assumed to be protected by signing envelope for v2.
type ZAttestQuoteResp struct {
	Response             ZAttestResponseCode `protobuf:"varint,1,opt,name=response,proto3,enum=ZAttestResponseCode" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ZAttestQuoteResp) Reset()         { *m = ZAttestQuoteResp{} }
func (m *ZAttestQuoteResp) String() string { return proto.CompactTextString(m) }
func (*ZAttestQuoteResp) ProtoMessage()    {}
func (*ZAttestQuoteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_208cf26448842a3f, []int{2}
}

func (m *ZAttestQuoteResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZAttestQuoteResp.Unmarshal(m, b)
}
func (m *ZAttestQuoteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZAttestQuoteResp.Marshal(b, m, deterministic)
}
func (m *ZAttestQuoteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZAttestQuoteResp.Merge(m, src)
}
func (m *ZAttestQuoteResp) XXX_Size() int {
	return xxx_messageInfo_ZAttestQuoteResp.Size(m)
}
func (m *ZAttestQuoteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ZAttestQuoteResp.DiscardUnknown(m)
}

var xxx_messageInfo_ZAttestQuoteResp proto.InternalMessageInfo

func (m *ZAttestQuoteResp) GetResponse() ZAttestResponseCode {
	if m != nil {
		return m.Response
	}
	return ZAttestResponseCode_ATTEST_RESPONSE_SUCCESS
}

func init() {
	proto.RegisterEnum("ZAttestResponseCode", ZAttestResponseCode_name, ZAttestResponseCode_value)
	proto.RegisterType((*ZAttestNonceResp)(nil), "ZAttestNonceResp")
	proto.RegisterType((*ZAttestQuoteReq)(nil), "ZAttestQuoteReq")
	proto.RegisterType((*ZAttestQuoteResp)(nil), "ZAttestQuoteResp")
}

func init() { proto.RegisterFile("attest.proto", fileDescriptor_208cf26448842a3f) }

var fileDescriptor_208cf26448842a3f = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xad, 0xa0, 0xe8, 0x50, 0x34, 0xc4, 0x82, 0x05, 0x45, 0x25, 0x5e, 0x8a, 0xe0, 0x46,
	0xf4, 0x17, 0xc4, 0x34, 0x05, 0x41, 0x1a, 0xdd, 0x4d, 0x2f, 0xbd, 0x94, 0xed, 0x66, 0x8c, 0x81,
	0x36, 0x13, 0xb3, 0xbb, 0x3d, 0xf8, 0xeb, 0xc5, 0x64, 0xa9, 0x1e, 0x7a, 0x9c, 0xef, 0x7d, 0xbc,
	0x07, 0x03, 0x7d, 0x69, 0x0c, 0x6a, 0xc3, 0xea, 0x86, 0x0c, 0x05, 0x23, 0xf0, 0xe6, 0x51, 0x0b,
	0xa6, 0x54, 0x29, 0xe4, 0xa8, 0x6b, 0x7f, 0x00, 0x07, 0xd5, 0xef, 0x31, 0xec, 0xdd, 0xf4, 0x46,
	0x7d, 0xde, 0x1d, 0x41, 0x0a, 0xa7, 0xce, 0x7c, 0xb7, 0x64, 0x90, 0xe3, 0x97, 0x7f, 0x05, 0xd0,
	0x95, 0x8d, 0xa5, 0x91, 0xce, 0xfe, 0x47, 0xfc, 0x4b, 0x38, 0xd6, 0x65, 0x51, 0x49, 0x63, 0x1b,
	0x1c, 0xee, 0xb7, 0xf1, 0x1f, 0x08, 0xc6, 0xdb, 0x69, 0x57, 0xa8, 0x6b, 0xff, 0x01, 0x8e, 0x1a,
	0xd4, 0x35, 0x55, 0xba, 0x5b, 0x3f, 0x79, 0x1c, 0x30, 0x27, 0x71, 0xc7, 0x63, 0xca, 0x91, 0x6f,
	0xad, 0xbb, 0x14, 0xce, 0x76, 0x08, 0xfe, 0x05, 0x9c, 0x47, 0x59, 0x96, 0x88, 0x6c, 0xc1, 0x13,
	0xf1, 0x96, 0x4e, 0x45, 0xb2, 0x10, 0xb3, 0x38, 0x4e, 0x84, 0xf0, 0xf6, 0x76, 0x85, 0x93, 0xe8,
	0xe5, 0x75, 0xc6, 0x13, 0xaf, 0xf7, 0x3c, 0x81, 0x6b, 0x45, 0x6b, 0xf6, 0x8d, 0x39, 0xe6, 0x92,
	0xa9, 0x15, 0xd9, 0x9c, 0x59, 0x8d, 0xcd, 0xa6, 0x54, 0xd8, 0x3d, 0x6d, 0x7e, 0x5b, 0x94, 0xe6,
	0xd3, 0x2e, 0x99, 0xa2, 0x75, 0xb8, 0xfa, 0xb8, 0xc7, 0xbc, 0xc0, 0x10, 0x37, 0x18, 0xca, 0xba,
	0x0c, 0x0b, 0x0a, 0x15, 0x36, 0x46, 0x2f, 0x0f, 0x5b, 0xf7, 0xe9, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x07, 0xd9, 0xb9, 0x8a, 0x70, 0x01, 0x00, 0x00,
}
