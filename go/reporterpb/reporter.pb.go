// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reporter.proto

package reporterpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateFileResponse struct {
	// @inject_tag: json:"status"
	//
	// The response status code.
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"message,omitempty"
	//
	// The response error message (if any).
	Message *ResponseErrorMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// @inject_tag: json:"file_id"
	//
	// The unique identifier for the file.
	FileId               string   `protobuf:"bytes,3,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFileResponse) Reset()         { *m = CreateFileResponse{} }
func (m *CreateFileResponse) String() string { return proto.CompactTextString(m) }
func (*CreateFileResponse) ProtoMessage()    {}
func (*CreateFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47040e06999d0d97, []int{0}
}

func (m *CreateFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFileResponse.Unmarshal(m, b)
}
func (m *CreateFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFileResponse.Marshal(b, m, deterministic)
}
func (m *CreateFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFileResponse.Merge(m, src)
}
func (m *CreateFileResponse) XXX_Size() int {
	return xxx_messageInfo_CreateFileResponse.Size(m)
}
func (m *CreateFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFileResponse proto.InternalMessageInfo

func (m *CreateFileResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreateFileResponse) GetMessage() *ResponseErrorMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *CreateFileResponse) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

type ResponseErrorMessage struct {
	//@inject_tag: json:"code"
	//
	// The response code.
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	//@inject_tag: json:"message"
	//
	// The response message.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	//@inject_tag: json:"details,omitempty"
	//
	// The response details.
	Details              string   `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseErrorMessage) Reset()         { *m = ResponseErrorMessage{} }
func (m *ResponseErrorMessage) String() string { return proto.CompactTextString(m) }
func (*ResponseErrorMessage) ProtoMessage()    {}
func (*ResponseErrorMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_47040e06999d0d97, []int{1}
}

func (m *ResponseErrorMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseErrorMessage.Unmarshal(m, b)
}
func (m *ResponseErrorMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseErrorMessage.Marshal(b, m, deterministic)
}
func (m *ResponseErrorMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseErrorMessage.Merge(m, src)
}
func (m *ResponseErrorMessage) XXX_Size() int {
	return xxx_messageInfo_ResponseErrorMessage.Size(m)
}
func (m *ResponseErrorMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseErrorMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseErrorMessage proto.InternalMessageInfo

func (m *ResponseErrorMessage) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ResponseErrorMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ResponseErrorMessage) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type ReportFile struct {
	//@inject_tag: json:"-" bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"-" validate:"required,hexadecimal,len=24"
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// @inject_tag: json:"-" validate:"required,hexadecimal,len=24"
	MerchantId string `protobuf:"bytes,3,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	// @inject_tag: json:"report_type" validate:"required,oneof=transactions vat vat_transactions royalty royalty_transactions payout agreement"
	//
	// The report type. Available values: transactions, vat, vat_transactions, royalty, royalty_transactions, payout, agreement.
	ReportType string `protobuf:"bytes,4,opt,name=report_type,json=reportType,proto3" json:"report_type,omitempty"`
	// @inject_tag: json:"file_type" validate:"required,alpha"
	//
	// The report file type. Available values: pdf, xlsx, csv.
	FileType string `protobuf:"bytes,5,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	// @inject_tag: json:"-"
	Params []byte `protobuf:"bytes,6,opt,name=params,proto3" json:"params,omitempty"`
	// @inject_tag: json:"template" validate:"omitempty,hexadecimal"
	//
	// The unique identifier for the report template.
	Template string `protobuf:"bytes,7,opt,name=template,proto3" json:"template,omitempty"`
	// @inject_tag: json:"retention_time"
	//
	// Retention time for the report.
	RetentionTime int32 `protobuf:"varint,8,opt,name=retention_time,json=retentionTime,proto3" json:"retention_time,omitempty"`
	// @inject_tag: json:"send_notification"
	//
	// Has a true value if it's required to send a notification about the report file to the user.
	SendNotification bool `protobuf:"varint,9,opt,name=send_notification,json=sendNotification,proto3" json:"send_notification,omitempty"`
	// @inject_tag: json:"created_at"
	//
	// The date of the report file creation.
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ReportFile) Reset()         { *m = ReportFile{} }
func (m *ReportFile) String() string { return proto.CompactTextString(m) }
func (*ReportFile) ProtoMessage()    {}
func (*ReportFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_47040e06999d0d97, []int{2}
}

func (m *ReportFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportFile.Unmarshal(m, b)
}
func (m *ReportFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportFile.Marshal(b, m, deterministic)
}
func (m *ReportFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportFile.Merge(m, src)
}
func (m *ReportFile) XXX_Size() int {
	return xxx_messageInfo_ReportFile.Size(m)
}
func (m *ReportFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportFile.DiscardUnknown(m)
}

var xxx_messageInfo_ReportFile proto.InternalMessageInfo

func (m *ReportFile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReportFile) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ReportFile) GetMerchantId() string {
	if m != nil {
		return m.MerchantId
	}
	return ""
}

func (m *ReportFile) GetReportType() string {
	if m != nil {
		return m.ReportType
	}
	return ""
}

func (m *ReportFile) GetFileType() string {
	if m != nil {
		return m.FileType
	}
	return ""
}

func (m *ReportFile) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *ReportFile) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *ReportFile) GetRetentionTime() int32 {
	if m != nil {
		return m.RetentionTime
	}
	return 0
}

func (m *ReportFile) GetSendNotification() bool {
	if m != nil {
		return m.SendNotification
	}
	return false
}

func (m *ReportFile) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type PostProcessRequest struct {
	ReportFile           *ReportFile `protobuf:"bytes,1,opt,name=report_file,json=reportFile,proto3" json:"report_file,omitempty"`
	FileName             string      `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	RetentionTime        int64       `protobuf:"varint,3,opt,name=retention_time,json=retentionTime,proto3" json:"retention_time,omitempty"`
	File                 []byte      `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PostProcessRequest) Reset()         { *m = PostProcessRequest{} }
func (m *PostProcessRequest) String() string { return proto.CompactTextString(m) }
func (*PostProcessRequest) ProtoMessage()    {}
func (*PostProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47040e06999d0d97, []int{3}
}

func (m *PostProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostProcessRequest.Unmarshal(m, b)
}
func (m *PostProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostProcessRequest.Marshal(b, m, deterministic)
}
func (m *PostProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostProcessRequest.Merge(m, src)
}
func (m *PostProcessRequest) XXX_Size() int {
	return xxx_messageInfo_PostProcessRequest.Size(m)
}
func (m *PostProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostProcessRequest proto.InternalMessageInfo

func (m *PostProcessRequest) GetReportFile() *ReportFile {
	if m != nil {
		return m.ReportFile
	}
	return nil
}

func (m *PostProcessRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *PostProcessRequest) GetRetentionTime() int64 {
	if m != nil {
		return m.RetentionTime
	}
	return 0
}

func (m *PostProcessRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateFileResponse)(nil), "reporter.CreateFileResponse")
	proto.RegisterType((*ResponseErrorMessage)(nil), "reporter.ResponseErrorMessage")
	proto.RegisterType((*ReportFile)(nil), "reporter.ReportFile")
	proto.RegisterType((*PostProcessRequest)(nil), "reporter.PostProcessRequest")
}

func init() { proto.RegisterFile("reporter.proto", fileDescriptor_47040e06999d0d97) }

var fileDescriptor_47040e06999d0d97 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x35, 0xdd, 0x6e, 0xda, 0xdc, 0x5d, 0x57, 0xbd, 0x2c, 0xeb, 0x50, 0xc5, 0x2d, 0x05, 0x21,
	0x20, 0x64, 0xa1, 0x22, 0xe8, 0xa3, 0x2b, 0x0a, 0xfb, 0xe0, 0xb2, 0x8c, 0xf5, 0xc5, 0x07, 0xc3,
	0x34, 0xb9, 0xad, 0x03, 0x49, 0x26, 0xce, 0x4c, 0x85, 0x7d, 0xf2, 0xc7, 0xf8, 0xa3, 0xfc, 0x3b,
	0x32, 0x93, 0x4c, 0xbb, 0x6a, 0xdf, 0x72, 0xcf, 0x3d, 0x73, 0x3f, 0xce, 0xb9, 0x81, 0x13, 0x4d,
	0xad, 0xd2, 0x96, 0x74, 0xd6, 0x6a, 0x65, 0x15, 0x8e, 0x43, 0x3c, 0x39, 0x5f, 0x2b, 0xb5, 0xae,
	0xe8, 0xc2, 0xe3, 0xcb, 0xcd, 0xea, 0xc2, 0xca, 0x9a, 0x8c, 0x15, 0x75, 0xdb, 0x51, 0x67, 0x3f,
	0x01, 0xdf, 0x69, 0x12, 0x96, 0x3e, 0xc8, 0x8a, 0x38, 0x99, 0x56, 0x35, 0x86, 0xf0, 0x0c, 0x62,
	0x63, 0x85, 0xdd, 0x18, 0x16, 0x4d, 0xa3, 0xf4, 0x90, 0xf7, 0x11, 0xbe, 0x86, 0x51, 0x4d, 0xc6,
	0x88, 0x35, 0xb1, 0xc1, 0x34, 0x4a, 0x8f, 0xe6, 0xcf, 0xb2, 0x6d, 0xeb, 0xf0, 0xf8, 0xbd, 0xd6,
	0x4a, 0x7f, 0xec, 0x58, 0x3c, 0xd0, 0xf1, 0x31, 0x8c, 0x56, 0xb2, 0xa2, 0x5c, 0x96, 0xec, 0x60,
	0x1a, 0xa5, 0x09, 0x8f, 0x5d, 0x78, 0x55, 0xce, 0xbe, 0xc2, 0xe9, 0xbe, 0x97, 0x88, 0x30, 0x2c,
	0x54, 0x49, 0x7e, 0x80, 0x84, 0xfb, 0x6f, 0x64, 0x7f, 0xb7, 0x4f, 0x76, 0xe5, 0x19, 0x8c, 0x4a,
	0xb2, 0x42, 0x56, 0xa6, 0x2f, 0x1f, 0xc2, 0xd9, 0xef, 0x01, 0x00, 0xf7, 0x33, 0xba, 0x0d, 0xf1,
	0x04, 0x06, 0xb2, 0xec, 0x8b, 0x0e, 0x64, 0xe9, 0xe6, 0xda, 0x18, 0xd2, 0x6e, 0xae, 0xae, 0x64,
	0xec, 0xc2, 0xab, 0x12, 0xcf, 0xe1, 0xa8, 0x26, 0x5d, 0x7c, 0x13, 0x8d, 0xdd, 0x0d, 0x0d, 0x01,
	0xea, 0x08, 0xdd, 0xee, 0xb9, 0xbd, 0x6d, 0x89, 0x0d, 0x3b, 0x42, 0x07, 0x2d, 0x6e, 0x5b, 0xc2,
	0x27, 0x90, 0xf8, 0x95, 0x7d, 0xfa, 0xd0, 0xa7, 0xc7, 0x0e, 0xf0, 0xc9, 0x33, 0x88, 0x5b, 0xa1,
	0x45, 0x6d, 0x58, 0x3c, 0x8d, 0xd2, 0x63, 0xde, 0x47, 0x38, 0x81, 0xb1, 0xa5, 0xba, 0xad, 0x84,
	0x25, 0x36, 0xea, 0xde, 0x84, 0x18, 0x9f, 0x3b, 0xa3, 0x2d, 0x35, 0x56, 0xaa, 0x26, 0x77, 0x46,
	0xb2, 0xb1, 0x77, 0xe7, 0xfe, 0x16, 0x5d, 0xc8, 0x9a, 0xf0, 0x05, 0x3c, 0x32, 0xd4, 0x94, 0x79,
	0xa3, 0xac, 0x5c, 0xc9, 0x42, 0xb8, 0x04, 0x4b, 0xa6, 0x51, 0x3a, 0xe6, 0x0f, 0x5d, 0xe2, 0xfa,
	0x0e, 0x8e, 0x6f, 0x00, 0x0a, 0xef, 0x7f, 0x99, 0x0b, 0xcb, 0xc0, 0x9b, 0x3a, 0xc9, 0xba, 0xab,
	0xc9, 0xc2, 0xd5, 0x64, 0x8b, 0x70, 0x35, 0x3c, 0xe9, 0xd9, 0x6f, 0xed, 0xec, 0x57, 0x04, 0x78,
	0xa3, 0x8c, 0xbd, 0xd1, 0xaa, 0x20, 0x63, 0x38, 0x7d, 0xdf, 0x90, 0xb1, 0xf8, 0x6a, 0xab, 0x8b,
	0x5b, 0xd6, 0x4b, 0x7d, 0x34, 0x3f, 0xbd, 0x7b, 0x27, 0xc1, 0x8c, 0xa0, 0x96, 0x37, 0x26, 0xa8,
	0xd5, 0x88, 0x3a, 0xb8, 0xeb, 0xd5, 0xba, 0x16, 0xf5, 0xbe, 0xcd, 0x9d, 0x1f, 0x07, 0xff, 0x6e,
	0x8e, 0x30, 0xf4, 0x3d, 0x87, 0x5e, 0x52, 0xff, 0x3d, 0xff, 0x0c, 0x0f, 0x78, 0xdf, 0xfa, 0x13,
	0xe9, 0x1f, 0xb2, 0x20, 0xbc, 0x04, 0xd8, 0xdd, 0x3c, 0xee, 0x1d, 0x6d, 0xf2, 0x74, 0x87, 0xfe,
	0xff, 0x7f, 0xcc, 0xee, 0x5d, 0x1e, 0x7f, 0x81, 0x40, 0x68, 0x97, 0xcb, 0xd8, 0x2b, 0xf5, 0xf2,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xc0, 0x4e, 0x7c, 0x89, 0x03, 0x00, 0x00,
}
