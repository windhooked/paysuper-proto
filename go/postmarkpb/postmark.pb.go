// Code generated by protoc-gen-go. DO NOT EDIT.
// source: postmark.proto

package postmarkpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type PayloadHeader struct {
	//@inject_tag: json:"Name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"Name"`
	//@inject_tag: json:"Value"
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-" validate:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-" validate:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-" validate:"-"`
}

func (m *PayloadHeader) Reset()         { *m = PayloadHeader{} }
func (m *PayloadHeader) String() string { return proto.CompactTextString(m) }
func (*PayloadHeader) ProtoMessage()    {}
func (*PayloadHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_e023c86970429eaa, []int{0}
}

func (m *PayloadHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayloadHeader.Unmarshal(m, b)
}
func (m *PayloadHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayloadHeader.Marshal(b, m, deterministic)
}
func (m *PayloadHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayloadHeader.Merge(m, src)
}
func (m *PayloadHeader) XXX_Size() int {
	return xxx_messageInfo_PayloadHeader.Size(m)
}
func (m *PayloadHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_PayloadHeader.DiscardUnknown(m)
}

var xxx_messageInfo_PayloadHeader proto.InternalMessageInfo

func (m *PayloadHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PayloadHeader) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type PayloadAttachment struct {
	//@inject_tag: json:"Name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"Name"`
	//@inject_tag: json:"Content"
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"Content"`
	//@inject_tag: json:"ContentType"
	ContentType          string   `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"ContentType"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-" validate:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-" validate:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-" validate:"-"`
}

func (m *PayloadAttachment) Reset()         { *m = PayloadAttachment{} }
func (m *PayloadAttachment) String() string { return proto.CompactTextString(m) }
func (*PayloadAttachment) ProtoMessage()    {}
func (*PayloadAttachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e023c86970429eaa, []int{1}
}

func (m *PayloadAttachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayloadAttachment.Unmarshal(m, b)
}
func (m *PayloadAttachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayloadAttachment.Marshal(b, m, deterministic)
}
func (m *PayloadAttachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayloadAttachment.Merge(m, src)
}
func (m *PayloadAttachment) XXX_Size() int {
	return xxx_messageInfo_PayloadAttachment.Size(m)
}
func (m *PayloadAttachment) XXX_DiscardUnknown() {
	xxx_messageInfo_PayloadAttachment.DiscardUnknown(m)
}

var xxx_messageInfo_PayloadAttachment proto.InternalMessageInfo

func (m *PayloadAttachment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PayloadAttachment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *PayloadAttachment) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

type Payload struct {
	//@inject_tag: protobuf:"varint,1,opt,name=template_id,json=TemplateId,proto3"
	TemplateId int32 `protobuf:"varint,1,opt,name=template_id,json=TemplateId,proto3" json:"template_id,omitempty"`
	//@inject_tag: `protobuf:"bytes,2,opt,name=template_alias,json=TemplateAlias,proto3"
	TemplateAlias string `protobuf:"bytes,2,opt,name=template_alias,json=TemplateAlias,proto3" json:"template_alias,omitempty"`
	//@inject_tag: protobuf:"bytes,3,rep,name=template_model,proto3" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"
	TemplateModel map[string]string `protobuf:"bytes,3,rep,name=template_model,proto3" json:"template_model,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//@inject_tag: protobuf:"varint,4,opt,name=inline_css,json=InlineCss,proto3"
	InlineCss bool `protobuf:"varint,4,opt,name=inline_css,json=InlineCss,proto3" json:"inline_css,omitempty"`
	//@inject_tag: protobuf:"bytes,5,opt,name=from,json=From,proto3"
	From string `protobuf:"bytes,5,opt,name=from,json=From,proto3" json:"from,omitempty"`
	//@inject_tag: protobuf:"bytes,6,opt,name=to,json=To,proto3"
	To string `protobuf:"bytes,6,opt,name=to,json=To,proto3" json:"to,omitempty"`
	//@inject_tag: protobuf:"bytes,7,opt,name=cc,json=Cc,proto3"
	Cc string `protobuf:"bytes,7,opt,name=cc,json=Cc,proto3" json:"cc,omitempty"`
	//@inject_tag: protobuf:"bytes,8,opt,name=bcc,json=Bcc,proto3"
	Bcc string `protobuf:"bytes,8,opt,name=bcc,json=Bcc,proto3" json:"bcc,omitempty"`
	//@inject_tag: protobuf:"bytes,9,opt,name=tag,json=Tag,proto3"
	Tag string `protobuf:"bytes,9,opt,name=tag,json=Tag,proto3" json:"tag,omitempty"`
	//@inject_tag: protobuf:"bytes,10,opt,name=reply_to,json=ReplyTo,proto3"
	ReplyTo string `protobuf:"bytes,10,opt,name=reply_to,json=ReplyTo,proto3" json:"reply_to,omitempty"`
	//@inject_tag: protobuf:"bytes,11,rep,name=headers,json=Headers,proto3"
	Headers []*PayloadHeader `protobuf:"bytes,11,rep,name=headers,json=Headers,proto3" json:"headers,omitempty"`
	//@inject_tag: protobuf:"varint,12,opt,name=track_opens,json=TrackOpens,proto3"
	TrackOpens bool `protobuf:"varint,12,opt,name=track_opens,json=TrackOpens,proto3" json:"track_opens,omitempty"`
	//@inject_tag: protobuf:"bytes,13,opt,name=track_links,json=TrackLinks,proto3"
	TrackLinks string `protobuf:"bytes,13,opt,name=track_links,json=TrackLinks,proto3" json:"track_links,omitempty"`
	//@inject_tag: protobuf:"bytes,14,rep,name=Attachments,json=Attachments,proto3"
	Attachments []*PayloadAttachment `protobuf:"bytes,14,rep,name=Attachments,json=Attachments,proto3" json:"attachments,omitempty"`
	//@inject_tag: protobuf:"bytes,15,rep,name=metadata,json=Metadata,proto3" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"
	Metadata map[string]string `protobuf:"bytes,15,rep,name=metadata,json=Metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//@inject_tag: protobuf:"bytes,16,opt,name=template_object_model,json=TemplateModel,proto3"
	TemplateObjectModel  *_struct.Struct `protobuf:"bytes,16,opt,name=template_object_model,json=TemplateModel,proto3" json:"template_object_model,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-" bson:"-" structure:"-" validate:"-"`
	XXX_unrecognized     []byte          `json:"-" bson:"-" structure:"-" validate:"-"`
	XXX_sizecache        int32           `json:"-" bson:"-" structure:"-" validate:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_e023c86970429eaa, []int{2}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetTemplateId() int32 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func (m *Payload) GetTemplateAlias() string {
	if m != nil {
		return m.TemplateAlias
	}
	return ""
}

func (m *Payload) GetTemplateModel() map[string]string {
	if m != nil {
		return m.TemplateModel
	}
	return nil
}

func (m *Payload) GetInlineCss() bool {
	if m != nil {
		return m.InlineCss
	}
	return false
}

func (m *Payload) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Payload) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Payload) GetCc() string {
	if m != nil {
		return m.Cc
	}
	return ""
}

func (m *Payload) GetBcc() string {
	if m != nil {
		return m.Bcc
	}
	return ""
}

func (m *Payload) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Payload) GetReplyTo() string {
	if m != nil {
		return m.ReplyTo
	}
	return ""
}

func (m *Payload) GetHeaders() []*PayloadHeader {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Payload) GetTrackOpens() bool {
	if m != nil {
		return m.TrackOpens
	}
	return false
}

func (m *Payload) GetTrackLinks() string {
	if m != nil {
		return m.TrackLinks
	}
	return ""
}

func (m *Payload) GetAttachments() []*PayloadAttachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Payload) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Payload) GetTemplateObjectModel() *_struct.Struct {
	if m != nil {
		return m.TemplateObjectModel
	}
	return nil
}

func init() {
	proto.RegisterType((*PayloadHeader)(nil), "postmark.PayloadHeader")
	proto.RegisterType((*PayloadAttachment)(nil), "postmark.PayloadAttachment")
	proto.RegisterType((*Payload)(nil), "postmark.Payload")
	proto.RegisterMapType((map[string]string)(nil), "postmark.Payload.MetadataEntry")
	proto.RegisterMapType((map[string]string)(nil), "postmark.Payload.TemplateModelEntry")
}

func init() {
	proto.RegisterFile("postmark.proto", fileDescriptor_e023c86970429eaa)
}

var fileDescriptor_e023c86970429eaa = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0xd5, 0x76, 0x5d, 0xdb, 0xeb, 0x0f, 0x86, 0x01, 0x61, 0x0a, 0x68, 0xa5, 0x02, 0xa9,
	0x4f, 0x9d, 0x80, 0x17, 0x60, 0x42, 0x62, 0x20, 0x24, 0xd0, 0x98, 0x86, 0x42, 0x9f, 0x78, 0x89,
	0x1c, 0xc7, 0xeb, 0x42, 0x93, 0x38, 0x8a, 0xaf, 0x48, 0xf9, 0xa3, 0xf9, 0x1f, 0x90, 0xcf, 0x4e,
	0x3a, 0xd4, 0xbd, 0xf0, 0x76, 0xf7, 0xc9, 0xf7, 0x72, 0xbe, 0xfb, 0xda, 0x30, 0x29, 0xb4, 0xc1,
	0x4c, 0x94, 0x9b, 0x65, 0x51, 0x6a, 0xd4, 0xac, 0x5f, 0xe7, 0xd3, 0x27, 0x6b, 0xad, 0xd7, 0xa9,
	0x3a, 0x21, 0x1e, 0x6d, 0xaf, 0x4e, 0x0c, 0x96, 0x5b, 0x89, 0x4e, 0x37, 0x7f, 0x0b, 0xe3, 0xef,
	0xa2, 0x4a, 0xb5, 0x88, 0xbf, 0x28, 0x11, 0xab, 0x92, 0x31, 0x38, 0xc8, 0x45, 0xa6, 0x78, 0x6b,
	0xd6, 0x5a, 0x0c, 0x02, 0x8a, 0xd9, 0x7d, 0xe8, 0xfe, 0x16, 0xe9, 0x56, 0xf1, 0x36, 0x41, 0x97,
	0xcc, 0x63, 0xb8, 0xeb, 0x4b, 0xcf, 0x10, 0x85, 0xbc, 0xce, 0x54, 0x8e, 0xb7, 0x96, 0x73, 0xe8,
	0x49, 0x9d, 0xa3, 0xca, 0xd1, 0xff, 0xa0, 0x4e, 0xd9, 0x33, 0x18, 0xf9, 0x30, 0xc4, 0xaa, 0x50,
	0xbc, 0x43, 0x9f, 0x87, 0x9e, 0xad, 0xaa, 0x42, 0xcd, 0xff, 0x74, 0xa1, 0xe7, 0xdb, 0xb0, 0x63,
	0x18, 0xa2, 0xca, 0x8a, 0x54, 0xa0, 0x0a, 0x93, 0x98, 0x7a, 0x74, 0x03, 0xa8, 0xd1, 0xd7, 0x98,
	0xbd, 0x80, 0x49, 0x23, 0x10, 0x69, 0x22, 0x8c, 0x6f, 0x38, 0xae, 0xe9, 0x99, 0x85, 0xec, 0xfc,
	0x86, 0x2c, 0xd3, 0xb1, 0x4a, 0x79, 0x67, 0xd6, 0x59, 0x0c, 0x5f, 0x3d, 0x5f, 0x36, 0x5b, 0xf4,
	0x2d, 0x97, 0x2b, 0xaf, 0xbb, 0xb0, 0xb2, 0xcf, 0x39, 0x96, 0xd5, 0xee, 0x67, 0xc4, 0xd8, 0x53,
	0x80, 0x24, 0x4f, 0x93, 0x5c, 0x85, 0xd2, 0x18, 0x7e, 0x30, 0x6b, 0x2d, 0xfa, 0xc1, 0xc0, 0x91,
	0x4f, 0xc6, 0xd8, 0x85, 0x5c, 0x95, 0x3a, 0xe3, 0x5d, 0xb7, 0x10, 0x1b, 0xb3, 0x09, 0xb4, 0x51,
	0xf3, 0x43, 0x22, 0x6d, 0xd4, 0x36, 0x97, 0x92, 0xf7, 0x5c, 0x2e, 0x25, 0x3b, 0x82, 0x4e, 0x24,
	0x25, 0xef, 0x13, 0xb0, 0xa1, 0x25, 0x28, 0xd6, 0x7c, 0xe0, 0x08, 0x8a, 0x35, 0x7b, 0x04, 0xfd,
	0x52, 0x15, 0x69, 0x15, 0xa2, 0xe6, 0xe0, 0xb6, 0x4a, 0xf9, 0x4a, 0xb3, 0x97, 0xd0, 0xbb, 0x26,
	0x33, 0x0d, 0x1f, 0xd2, 0x5c, 0x0f, 0xf7, 0xe6, 0x72, 0x66, 0x07, 0xb5, 0x8e, 0x36, 0x5b, 0x0a,
	0xb9, 0x09, 0x75, 0xa1, 0x72, 0xc3, 0x47, 0x34, 0x05, 0x10, 0xba, 0xb4, 0x64, 0x27, 0x48, 0x93,
	0x7c, 0x63, 0xf8, 0x98, 0x3a, 0x3a, 0xc1, 0x37, 0x4b, 0xd8, 0x7b, 0x18, 0x8a, 0xe6, 0x1a, 0x18,
	0x3e, 0xa1, 0xc6, 0x8f, 0xf7, 0x1a, 0xef, 0xae, 0x4a, 0x70, 0x53, 0xcf, 0x4e, 0xa1, 0x9f, 0x29,
	0x14, 0xb1, 0x40, 0xc1, 0xef, 0x50, 0xed, 0xf1, 0xbe, 0x19, 0x17, 0x5e, 0xe1, 0x7c, 0x68, 0x0a,
	0xd8, 0x39, 0x3c, 0x68, 0xfc, 0xd4, 0xd1, 0x2f, 0x25, 0xd1, 0xdb, 0x7a, 0x34, 0x6b, 0xd1, 0xf8,
	0xee, 0x09, 0x2c, 0xeb, 0x27, 0xb0, 0xfc, 0x41, 0x4f, 0x20, 0xb8, 0x57, 0x57, 0x5d, 0x52, 0x11,
	0xf9, 0x39, 0xfd, 0x00, 0x6c, 0xdf, 0x74, 0x6b, 0xc0, 0x46, 0x55, 0xfe, 0x5a, 0xdb, 0xf0, 0xf6,
	0x47, 0xf1, 0xae, 0xfd, 0xa6, 0x35, 0x3d, 0x85, 0xf1, 0x3f, 0x27, 0xfd, 0x9f, 0xe2, 0x8f, 0xa3,
	0x9f, 0x50, 0xcf, 0x5d, 0x44, 0xd1, 0x21, 0x1d, 0xf9, 0xf5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x93, 0xba, 0x0f, 0xd7, 0xdf, 0x03, 0x00, 0x00,
}
