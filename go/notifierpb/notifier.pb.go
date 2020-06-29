// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: notifier.proto

package notifierpb

import (
	proto "github.com/golang/protobuf/proto"
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

type CheckUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url           string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	User          *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	SecretKey     string `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	IsLiveProject bool   `protobuf:"varint,4,opt,name=is_live_project,json=isLiveProject,proto3" json:"is_live_project,omitempty"`
	TestingCase   string `protobuf:"bytes,5,opt,name=testing_case,json=testingCase,proto3" json:"testing_case,omitempty"`
	OrderUuid     string `protobuf:"bytes,6,opt,name=order_uuid,json=orderUuid,proto3" json:"order_uuid,omitempty"`
	MerchantId    string `protobuf:"bytes,7,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
}

func (x *CheckUserRequest) Reset() {
	*x = CheckUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserRequest) ProtoMessage() {}

func (x *CheckUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserRequest.ProtoReflect.Descriptor instead.
func (*CheckUserRequest) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{0}
}

func (x *CheckUserRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CheckUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CheckUserRequest) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *CheckUserRequest) GetIsLiveProject() bool {
	if x != nil {
		return x.IsLiveProject
	}
	return false
}

func (x *CheckUserRequest) GetTestingCase() string {
	if x != nil {
		return x.TestingCase
	}
	return ""
}

func (x *CheckUserRequest) GetOrderUuid() string {
	if x != nil {
		return x.OrderUuid
	}
	return ""
}

func (x *CheckUserRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

type CheckUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CheckUserResponse) Reset() {
	*x = CheckUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserResponse) ProtoMessage() {}

func (x *CheckUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserResponse.ProtoReflect.Descriptor instead.
func (*CheckUserResponse) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{1}
}

func (x *CheckUserResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CheckUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	//
	// The unique identifier for the customer in the paysuper.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// @inject_tag: json:"object"
	//
	// The system constant that contains the returned object's type.
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object"`
	// @inject_tag: json:"external_id"
	//
	// The unique identifier for the customer in the merchant project.
	ExternalId string `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id"`
	// @inject_tag: json:"name"
	//
	// The customer's name.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	// @inject_tag: json:"email"
	//
	// The customer's email address.
	Email string `protobuf:"bytes,5,opt,name=email,proto3" json:"email"`
	// @inject_tag: json:"email_verified"
	//
	// Whether the customer's email address has been verified on the merchant side.
	EmailVerified bool `protobuf:"varint,6,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified"`
	// @inject_tag: json:"phone"
	//
	// The customer's phone number.
	Phone string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone"`
	// @inject_tag: json:"phone_verified"
	//
	// Whether the customer's phone number has been verified on the merchant side.
	PhoneVerified bool `protobuf:"varint,8,opt,name=phone_verified,json=phoneVerified,proto3" json:"phone_verified"`
	// @inject_tag: json:"ip"
	//
	// The customer's IP address.
	Ip string `protobuf:"bytes,9,opt,name=ip,proto3" json:"ip"`
	// @inject_tag: json:"locale"
	//
	// The customer's locale name. The language code in ISO 639-1 (for instance en-US).
	Locale string `protobuf:"bytes,10,opt,name=locale,proto3" json:"locale"`
	// @inject_tag: json:"address"
	//
	// The customer's address details.
	Address *BillingAddress `protobuf:"bytes,11,opt,name=address,proto3" json:"address"`
	// @inject_tag: json:"metadata"
	//
	// A string-value description that you can attach to the customer's object. It can be useful for storing additional information about your customer's payment.
	Metadata map[string]string `protobuf:"bytes,12,rep,name=metadata,proto3" json:"metadata" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// @inject_tag: json:"notify_new_region"
	//
	// Has the value true if the customer confirmed to receive the PaySuper newsletter about the beginning of payment acceptance in new regions.
	NotifyNewRegion bool `protobuf:"varint,13,opt,name=notify_new_region,json=notifyNewRegion,proto3" json:"notify_new_region"`
	// @inject_tag: json:"notify_new_region_email"
	//
	// The customer’s email for a newsletter about the beginning of payment acceptance in new regions.
	NotifyNewRegionEmail string `protobuf:"bytes,14,opt,name=notify_new_region_email,json=notifyNewRegionEmail,proto3" json:"notify_new_region_email"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *User) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetPhoneVerified() bool {
	if x != nil {
		return x.PhoneVerified
	}
	return false
}

func (x *User) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *User) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *User) GetAddress() *BillingAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *User) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *User) GetNotifyNewRegion() bool {
	if x != nil {
		return x.NotifyNewRegion
	}
	return false
}

func (x *User) GetNotifyNewRegionEmail() string {
	if x != nil {
		return x.NotifyNewRegionEmail
	}
	return ""
}

type BillingAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"country"
	//
	// The customer's country. Two-letter country code in ISO 3166-1, in uppercase (for instance US).
	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country"`
	// @inject_tag: json:"city"
	//
	// The customer’s city.
	City string `protobuf:"bytes,2,opt,name=city,proto3" json:"city"`
	// @inject_tag: json:"postal_code"
	//
	// The customer's postal code.
	PostalCode string `protobuf:"bytes,3,opt,name=postal_code,json=postalCode,proto3" json:"postal_code"`
	// @inject_tag: json:"state"
	//
	// The customer's state code in ISO 3166-2.
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state"`
}

func (x *BillingAddress) Reset() {
	*x = BillingAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingAddress) ProtoMessage() {}

func (x *BillingAddress) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingAddress.ProtoReflect.Descriptor instead.
func (*BillingAddress) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{3}
}

func (x *BillingAddress) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *BillingAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *BillingAddress) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *BillingAddress) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

var File_notifier_proto protoreflect.FileDescriptor

var file_notifier_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xf2, 0x01, 0x0a, 0x10, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69,
	0x73, 0x4c, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x45, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x93, 0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x32,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x11,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4e,
	0x65, 0x77, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x17, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x1a,
	0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x75, 0x0a, 0x0e,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x32, 0x59, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifier_proto_rawDescOnce sync.Once
	file_notifier_proto_rawDescData = file_notifier_proto_rawDesc
)

func file_notifier_proto_rawDescGZIP() []byte {
	file_notifier_proto_rawDescOnce.Do(func() {
		file_notifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifier_proto_rawDescData)
	})
	return file_notifier_proto_rawDescData
}

var file_notifier_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_notifier_proto_goTypes = []interface{}{
	(*CheckUserRequest)(nil),  // 0: notifier.CheckUserRequest
	(*CheckUserResponse)(nil), // 1: notifier.CheckUserResponse
	(*User)(nil),              // 2: notifier.User
	(*BillingAddress)(nil),    // 3: notifier.BillingAddress
	nil,                       // 4: notifier.User.MetadataEntry
}
var file_notifier_proto_depIdxs = []int32{
	2, // 0: notifier.CheckUserRequest.user:type_name -> notifier.User
	3, // 1: notifier.User.address:type_name -> notifier.BillingAddress
	4, // 2: notifier.User.metadata:type_name -> notifier.User.MetadataEntry
	0, // 3: notifier.NotifierService.CheckUser:input_type -> notifier.CheckUserRequest
	1, // 4: notifier.NotifierService.CheckUser:output_type -> notifier.CheckUserResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_notifier_proto_init() }
func file_notifier_proto_init() {
	if File_notifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserRequest); i {
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
		file_notifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserResponse); i {
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
		file_notifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_notifier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingAddress); i {
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
			RawDescriptor: file_notifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notifier_proto_goTypes,
		DependencyIndexes: file_notifier_proto_depIdxs,
		MessageInfos:      file_notifier_proto_msgTypes,
	}.Build()
	File_notifier_proto = out.File
	file_notifier_proto_rawDesc = nil
	file_notifier_proto_goTypes = nil
	file_notifier_proto_depIdxs = nil
}
