// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nodes.proto

package nodes_pb

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

type LinkedListNode struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Next                 []byte   `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty"`
	Prev                 []byte   `protobuf:"bytes,3,opt,name=prev,proto3" json:"prev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkedListNode) Reset()         { *m = LinkedListNode{} }
func (m *LinkedListNode) String() string { return proto.CompactTextString(m) }
func (*LinkedListNode) ProtoMessage()    {}
func (*LinkedListNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_39b18f0c01aa3995, []int{0}
}

func (m *LinkedListNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkedListNode.Unmarshal(m, b)
}
func (m *LinkedListNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkedListNode.Marshal(b, m, deterministic)
}
func (m *LinkedListNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkedListNode.Merge(m, src)
}
func (m *LinkedListNode) XXX_Size() int {
	return xxx_messageInfo_LinkedListNode.Size(m)
}
func (m *LinkedListNode) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkedListNode.DiscardUnknown(m)
}

var xxx_messageInfo_LinkedListNode proto.InternalMessageInfo

func (m *LinkedListNode) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *LinkedListNode) GetNext() []byte {
	if m != nil {
		return m.Next
	}
	return nil
}

func (m *LinkedListNode) GetPrev() []byte {
	if m != nil {
		return m.Prev
	}
	return nil
}

func init() {
	proto.RegisterType((*LinkedListNode)(nil), "nodes.LinkedListNode")
}

func init() { proto.RegisterFile("nodes.proto", fileDescriptor_39b18f0c01aa3995) }

var fileDescriptor_39b18f0c01aa3995 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcb, 0x4f, 0x49,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x7c, 0xb8, 0xf8, 0x7c,
	0x32, 0xf3, 0xb2, 0x53, 0x53, 0x7c, 0x32, 0x8b, 0x4b, 0xfc, 0xf2, 0x53, 0x52, 0x85, 0x84, 0xb8,
	0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c, 0x90, 0x58,
	0x5e, 0x6a, 0x45, 0x89, 0x04, 0x13, 0x44, 0x0c, 0xc4, 0x06, 0x89, 0x15, 0x14, 0xa5, 0x96, 0x49,
	0x30, 0x43, 0xc4, 0x40, 0x6c, 0x27, 0xae, 0x28, 0x0e, 0xb0, 0xb1, 0xf1, 0x05, 0x49, 0x49, 0x6c,
	0x60, 0x7b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x48, 0xb0, 0xc1, 0x76, 0x00, 0x00,
	0x00,
}
