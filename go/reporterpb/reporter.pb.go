// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: reporter.proto

package reporterpb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"status"
	//
	// The response status code.
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status"`
	// @inject_tag: json:"message,omitempty"
	//
	// The response error message (if any).
	Message *ResponseErrorMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// @inject_tag: json:"file_id"
	//
	// The unique identifier for the file.
	FileId string `protobuf:"bytes,3,opt,name=file_id,json=fileId,proto3" json:"file_id"`
}

func (x *CreateFileResponse) Reset() {
	*x = CreateFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileResponse) ProtoMessage() {}

func (x *CreateFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileResponse.ProtoReflect.Descriptor instead.
func (*CreateFileResponse) Descriptor() ([]byte, []int) {
	return file_reporter_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFileResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateFileResponse) GetMessage() *ResponseErrorMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *CreateFileResponse) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type ResponseErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag: json:"code"
	//
	// The response code.
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code"`
	//@inject_tag: json:"message"
	//
	// The response message.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	//@inject_tag: json:"details,omitempty"
	//
	// The response details.
	Details string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *ResponseErrorMessage) Reset() {
	*x = ResponseErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reporter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseErrorMessage) ProtoMessage() {}

func (x *ResponseErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_reporter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseErrorMessage.ProtoReflect.Descriptor instead.
func (*ResponseErrorMessage) Descriptor() ([]byte, []int) {
	return file_reporter_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseErrorMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ResponseErrorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseErrorMessage) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

type ReportFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag: json:"-" bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"-" bson:"_id"`
	// @inject_tag: json:"-" validate:"required,hexadecimal,len=24"
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"-" validate:"required,hexadecimal,len=24"`
	// @inject_tag: json:"-" validate:"required,hexadecimal,len=24"
	MerchantId string `protobuf:"bytes,3,opt,name=merchant_id,json=merchantId,proto3" json:"-" validate:"required,hexadecimal,len=24"`
	// @inject_tag: json:"report_type" validate:"required,oneof=transactions vat vat_transactions royalty royalty_transactions payout agreement"
	//
	// The report type. Available values: transactions, vat, vat_transactions, royalty, royalty_transactions, payout, agreement.
	ReportType string `protobuf:"bytes,4,opt,name=report_type,json=reportType,proto3" json:"report_type" validate:"required,oneof=transactions vat vat_transactions royalty royalty_transactions payout agreement"`
	// @inject_tag: json:"file_type" validate:"required,alpha"
	//
	// The report file type. Available values: pdf, xlsx, csv.
	FileType string `protobuf:"bytes,5,opt,name=file_type,json=fileType,proto3" json:"file_type" validate:"required,alpha"`
	// @inject_tag: json:"-"
	Params []byte `protobuf:"bytes,6,opt,name=params,proto3" json:"-"`
	// @inject_tag: json:"template" validate:"omitempty,hexadecimal"
	//
	// The unique identifier for the report template.
	Template string `protobuf:"bytes,7,opt,name=template,proto3" json:"template" validate:"omitempty,hexadecimal"`
	// @inject_tag: json:"retention_time"
	//
	// Retention time for the report.
	RetentionTime int32 `protobuf:"varint,8,opt,name=retention_time,json=retentionTime,proto3" json:"retention_time"`
	// @inject_tag: json:"send_notification"
	//
	// Has a true value if it's required to send a notification about the report file to the user.
	SendNotification bool `protobuf:"varint,9,opt,name=send_notification,json=sendNotification,proto3" json:"send_notification"`
	// @inject_tag: json:"created_at"
	//
	// The date of the report file creation.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	// @inject_tag: json:"notification_channel_id"
	//
	// Merchant's or user's identifier to send notification in correct channel
	NotificationChannelId string `protobuf:"bytes,11,opt,name=notification_channel_id,json=notificationChannelId,proto3" json:"notification_channel_id"`
	// @inject_tag: json:"skip_post_process"
	//
	// Has a true value if need to skip post process method call after the report file exported.
	SkipPostProcess bool `protobuf:"varint,12,opt,name=skip_post_process,json=skipPostProcess,proto3" json:"skip_post_process"`
}

func (x *ReportFile) Reset() {
	*x = ReportFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reporter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportFile) ProtoMessage() {}

func (x *ReportFile) ProtoReflect() protoreflect.Message {
	mi := &file_reporter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportFile.ProtoReflect.Descriptor instead.
func (*ReportFile) Descriptor() ([]byte, []int) {
	return file_reporter_proto_rawDescGZIP(), []int{2}
}

func (x *ReportFile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReportFile) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReportFile) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *ReportFile) GetReportType() string {
	if x != nil {
		return x.ReportType
	}
	return ""
}

func (x *ReportFile) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *ReportFile) GetParams() []byte {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *ReportFile) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

func (x *ReportFile) GetRetentionTime() int32 {
	if x != nil {
		return x.RetentionTime
	}
	return 0
}

func (x *ReportFile) GetSendNotification() bool {
	if x != nil {
		return x.SendNotification
	}
	return false
}

func (x *ReportFile) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ReportFile) GetNotificationChannelId() string {
	if x != nil {
		return x.NotificationChannelId
	}
	return ""
}

func (x *ReportFile) GetSkipPostProcess() bool {
	if x != nil {
		return x.SkipPostProcess
	}
	return false
}

type PostProcessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportFile    *ReportFile       `protobuf:"bytes,1,opt,name=report_file,json=reportFile,proto3" json:"report_file,omitempty"`
	FileName      string            `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	RetentionTime int64             `protobuf:"varint,3,opt,name=retention_time,json=retentionTime,proto3" json:"retention_time,omitempty"`
	File          []byte            `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	Params        map[string]string `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PostProcessRequest) Reset() {
	*x = PostProcessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reporter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostProcessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostProcessRequest) ProtoMessage() {}

func (x *PostProcessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reporter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostProcessRequest.ProtoReflect.Descriptor instead.
func (*PostProcessRequest) Descriptor() ([]byte, []int) {
	return file_reporter_proto_rawDescGZIP(), []int{3}
}

func (x *PostProcessRequest) GetReportFile() *ReportFile {
	if x != nil {
		return x.ReportFile
	}
	return nil
}

func (x *PostProcessRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *PostProcessRequest) GetRetentionTime() int64 {
	if x != nil {
		return x.RetentionTime
	}
	return 0
}

func (x *PostProcessRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *PostProcessRequest) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_reporter_proto protoreflect.FileDescriptor

var file_reporter_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x14,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xbb, 0x03, 0x0a,
	0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x65, 0x6e, 0x64, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x11, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x6b, 0x69, 0x70, 0x50,
	0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0xa0, 0x02, 0x0a, 0x12, 0x50,
	0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x40, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x55, 0x0a,
	0x0f, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14,
	0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reporter_proto_rawDescOnce sync.Once
	file_reporter_proto_rawDescData = file_reporter_proto_rawDesc
)

func file_reporter_proto_rawDescGZIP() []byte {
	file_reporter_proto_rawDescOnce.Do(func() {
		file_reporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_reporter_proto_rawDescData)
	})
	return file_reporter_proto_rawDescData
}

var file_reporter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_reporter_proto_goTypes = []interface{}{
	(*CreateFileResponse)(nil),   // 0: reporter.CreateFileResponse
	(*ResponseErrorMessage)(nil), // 1: reporter.ResponseErrorMessage
	(*ReportFile)(nil),           // 2: reporter.ReportFile
	(*PostProcessRequest)(nil),   // 3: reporter.PostProcessRequest
	nil,                          // 4: reporter.PostProcessRequest.ParamsEntry
	(*timestamp.Timestamp)(nil),  // 5: google.protobuf.Timestamp
}
var file_reporter_proto_depIdxs = []int32{
	1, // 0: reporter.CreateFileResponse.message:type_name -> reporter.ResponseErrorMessage
	5, // 1: reporter.ReportFile.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: reporter.PostProcessRequest.report_file:type_name -> reporter.ReportFile
	4, // 3: reporter.PostProcessRequest.params:type_name -> reporter.PostProcessRequest.ParamsEntry
	2, // 4: reporter.ReporterService.CreateFile:input_type -> reporter.ReportFile
	0, // 5: reporter.ReporterService.CreateFile:output_type -> reporter.CreateFileResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_reporter_proto_init() }
func file_reporter_proto_init() {
	if File_reporter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reporter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reporter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseErrorMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reporter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportFile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reporter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostProcessRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reporter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reporter_proto_goTypes,
		DependencyIndexes: file_reporter_proto_depIdxs,
		MessageInfos:      file_reporter_proto_msgTypes,
	}.Build()
	File_reporter_proto = out.File
	file_reporter_proto_rawDesc = nil
	file_reporter_proto_goTypes = nil
	file_reporter_proto_depIdxs = nil
}
