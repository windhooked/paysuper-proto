// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: paylink.proto

package billingpb

import (
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

type CreatePaylinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag: json:"id" validate:"omitempty,hexadecimal,len=24"
	//
	// The unique identifier for the payment link.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" validate:"omitempty,hexadecimal,len=24"`
	// @inject_tag: json:"expires_at" validate:"omitempty,numeric,gte=0"
	//
	// The date of the payment link expiration.
	ExpiresAt int64 `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at" validate:"omitempty,numeric,gte=0"`
	//@inject_tag: json:"products" validate:"required,gte=1,dive,hexadecimal,len=24" required:"true"
	//
	// The list of the payment link's products.
	Products []string `protobuf:"bytes,3,rep,name=products,proto3" json:"products" validate:"required,gte=1,dive,hexadecimal,len=24" required:"true"`
	//@inject_tag: json:"merchant_id" validate:"required,hexadecimal,len=24" required:"true"
	//
	// The unique identifier for the merchant.
	MerchantId string `protobuf:"bytes,4,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id" validate:"required,hexadecimal,len=24" required:"true"`
	//@inject_tag: json:"project_id" validate:"required,hexadecimal,len=24" required:"true"
	//
	// The unique identifier for the project.
	ProjectId string `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3" json:"project_id" validate:"required,hexadecimal,len=24" required:"true"`
	// @inject_tag: json:"name" validate:"required" required:"true"
	//
	// The payment link's name.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name" validate:"required" required:"true"`
	// @inject_tag: json:"no_expiry_date"
	//
	// Has a true value if the payment link has no expiry date.
	NoExpiryDate bool `protobuf:"varint,7,opt,name=no_expiry_date,json=noExpiryDate,proto3" json:"no_expiry_date"`
	// @inject_tag: json:"products_type" validate="required,oneof=product key" required:"true"
	//
	// The type of products. Available values: product, key.
	ProductsType string `protobuf:"bytes,8,opt,name=products_type,json=productsType,proto3" json:"products_type" required:"true"`
}

func (x *CreatePaylinkRequest) Reset() {
	*x = CreatePaylinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paylink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaylinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaylinkRequest) ProtoMessage() {}

func (x *CreatePaylinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paylink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaylinkRequest.ProtoReflect.Descriptor instead.
func (*CreatePaylinkRequest) Descriptor() ([]byte, []int) {
	return file_paylink_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePaylinkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePaylinkRequest) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *CreatePaylinkRequest) GetProducts() []string {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *CreatePaylinkRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *CreatePaylinkRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreatePaylinkRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePaylinkRequest) GetNoExpiryDate() bool {
	if x != nil {
		return x.NoExpiryDate
	}
	return false
}

func (x *CreatePaylinkRequest) GetProductsType() string {
	if x != nil {
		return x.ProductsType
	}
	return ""
}

type Paylink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	//
	// The unique identifier for the payment link.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// @inject_tag: json:"object"
	//
	// The system constant that contains the returned object's type. Const value: paylink.
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object"`
	// @inject_tag: json:"products"
	//
	// The list of the payment link's products.
	Products []string `protobuf:"bytes,3,rep,name=products,proto3" json:"products"`
	// @inject_tag: json:"expires_at"
	//
	// The date of the payment link expiration.
	ExpiresAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at"`
	// @inject_tag: json:"created_at"
	//
	// The date of the payment link creation.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	// @inject_tag: json:"updated_at"
	//
	// The date of the payment link last update.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	// @inject_tag: json:"merchant_id"
	//
	// The unique identifier for the merchant.
	MerchantId string `protobuf:"bytes,7,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id"`
	// @inject_tag: json:"project_id"
	//
	// The unique identifier for the project.
	ProjectId string `protobuf:"bytes,8,opt,name=project_id,json=projectId,proto3" json:"project_id"`
	// @inject_tag: json:"name"
	//
	// The payment link's name.
	Name string `protobuf:"bytes,9,opt,name=name,proto3" json:"name"`
	// @inject_tag: json:"is_expired"
	//
	// Has a true value if the payment link has expired.
	IsExpired bool `protobuf:"varint,10,opt,name=is_expired,json=isExpired,proto3" json:"is_expired"`
	// @inject_tag: json:"visits"
	//
	// The total number of visits.
	Visits int32 `protobuf:"varint,11,opt,name=visits,proto3" json:"visits"`
	// @inject_tag: json:"no_expiry_date"
	//
	// Has a true value if the payment link has no expiry date.
	NoExpiryDate bool `protobuf:"varint,12,opt,name=no_expiry_date,json=noExpiryDate,proto3" json:"no_expiry_date"`
	// @inject_tag: json:"products_type"
	//
	// The type of products. Available values: product, key.
	ProductsType string `protobuf:"bytes,13,opt,name=products_type,json=productsType,proto3" json:"products_type"`
	// @inject_tag: json:"-"
	Deleted bool `protobuf:"varint,14,opt,name=deleted,proto3" json:"-"`
	// @inject_tag: json:"total_transactions"
	//
	// The total number of transactions.
	TotalTransactions int32 `protobuf:"varint,15,opt,name=total_transactions,json=totalTransactions,proto3" json:"total_transactions"`
	// @inject_tag: json:"sales_count"
	//
	// The total number of sales.
	SalesCount int32 `protobuf:"varint,16,opt,name=sales_count,json=salesCount,proto3" json:"sales_count"`
	// @inject_tag: json:"returns_count"
	//
	// The total number of returns.
	ReturnsCount int32 `protobuf:"varint,17,opt,name=returns_count,json=returnsCount,proto3" json:"returns_count"`
	// @inject_tag: json:"conversion"
	//
	// The conversion - sales per visits relation.
	Conversion float64 `protobuf:"fixed64,18,opt,name=conversion,proto3" json:"conversion"`
	// @inject_tag: json:"gross_sales_amount"
	//
	// The gross amount of sales.
	GrossSalesAmount float64 `protobuf:"fixed64,19,opt,name=gross_sales_amount,json=grossSalesAmount,proto3" json:"gross_sales_amount"`
	// @inject_tag: json:"gross_returns_amount"
	//
	// The gross amount of returns.
	GrossReturnsAmount float64 `protobuf:"fixed64,20,opt,name=gross_returns_amount,json=grossReturnsAmount,proto3" json:"gross_returns_amount"`
	// @inject_tag: json:"gross_total_amount"
	//
	// The gross revenue.
	GrossTotalAmount float64 `protobuf:"fixed64,21,opt,name=gross_total_amount,json=grossTotalAmount,proto3" json:"gross_total_amount"`
	// @inject_tag: json:"transactions_currency"
	//
	// The transactions currency. Three-letter currency code in ISO 4217, in uppercase.
	TransactionsCurrency string `protobuf:"bytes,22,opt,name=transactions_currency,json=transactionsCurrency,proto3" json:"transactions_currency"`
	// @inject_tag: json:"short_link"
	//
	// The short identity of paylink for use the link shortening service.
	ShortLink string `protobuf:"bytes,23,opt,name=short_link,json=shortLink,proto3" json:"short_link"`
}

func (x *Paylink) Reset() {
	*x = Paylink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paylink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paylink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paylink) ProtoMessage() {}

func (x *Paylink) ProtoReflect() protoreflect.Message {
	mi := &file_paylink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paylink.ProtoReflect.Descriptor instead.
func (*Paylink) Descriptor() ([]byte, []int) {
	return file_paylink_proto_rawDescGZIP(), []int{1}
}

func (x *Paylink) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Paylink) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *Paylink) GetProducts() []string {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Paylink) GetExpiresAt() *timestamp.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *Paylink) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Paylink) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Paylink) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *Paylink) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Paylink) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Paylink) GetIsExpired() bool {
	if x != nil {
		return x.IsExpired
	}
	return false
}

func (x *Paylink) GetVisits() int32 {
	if x != nil {
		return x.Visits
	}
	return 0
}

func (x *Paylink) GetNoExpiryDate() bool {
	if x != nil {
		return x.NoExpiryDate
	}
	return false
}

func (x *Paylink) GetProductsType() string {
	if x != nil {
		return x.ProductsType
	}
	return ""
}

func (x *Paylink) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *Paylink) GetTotalTransactions() int32 {
	if x != nil {
		return x.TotalTransactions
	}
	return 0
}

func (x *Paylink) GetSalesCount() int32 {
	if x != nil {
		return x.SalesCount
	}
	return 0
}

func (x *Paylink) GetReturnsCount() int32 {
	if x != nil {
		return x.ReturnsCount
	}
	return 0
}

func (x *Paylink) GetConversion() float64 {
	if x != nil {
		return x.Conversion
	}
	return 0
}

func (x *Paylink) GetGrossSalesAmount() float64 {
	if x != nil {
		return x.GrossSalesAmount
	}
	return 0
}

func (x *Paylink) GetGrossReturnsAmount() float64 {
	if x != nil {
		return x.GrossReturnsAmount
	}
	return 0
}

func (x *Paylink) GetGrossTotalAmount() float64 {
	if x != nil {
		return x.GrossTotalAmount
	}
	return 0
}

func (x *Paylink) GetTransactionsCurrency() string {
	if x != nil {
		return x.TransactionsCurrency
	}
	return ""
}

func (x *Paylink) GetShortLink() string {
	if x != nil {
		return x.ShortLink
	}
	return ""
}

type StatCommon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"-" bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"-" bson:"_id"`
	// @inject_tag: json:"paylink_id" bson:"-"
	//
	// The unique identifier for the payment link.
	PaylinkId string `protobuf:"bytes,2,opt,name=paylink_id,json=paylinkId,proto3" json:"paylink_id" bson:"-"`
	// @inject_tag: json:"visits" bson:"-"
	//
	// The total number of visits.
	Visits int32 `protobuf:"varint,3,opt,name=visits,proto3" json:"visits" bson:"-"`
	// @inject_tag: json:"total_transactions" bson:"total_transactions"
	//
	// The total number of transactions.
	TotalTransactions int32 `protobuf:"varint,4,opt,name=total_transactions,json=totalTransactions,proto3" json:"total_transactions" bson:"total_transactions"`
	// @inject_tag: json:"sales_count" bson:"sales_count"
	//
	// The total number of sales.
	SalesCount int32 `protobuf:"varint,5,opt,name=sales_count,json=salesCount,proto3" json:"sales_count" bson:"sales_count"`
	// @inject_tag: json:"returns_count" bson:"returns_count"
	//
	// The total number of returns.
	ReturnsCount int32 `protobuf:"varint,6,opt,name=returns_count,json=returnsCount,proto3" json:"returns_count" bson:"returns_count"`
	// @inject_tag: json:"gross_sales_amount" bson:"gross_sales_amount"
	//
	// The gross amount of sales.
	GrossSalesAmount float64 `protobuf:"fixed64,7,opt,name=gross_sales_amount,json=grossSalesAmount,proto3" json:"gross_sales_amount" bson:"gross_sales_amount"`
	// @inject_tag: json:"gross_returns_amount" bson:"gross_returns_amount"
	//
	// The gross amount of returns.
	GrossReturnsAmount float64 `protobuf:"fixed64,8,opt,name=gross_returns_amount,json=grossReturnsAmount,proto3" json:"gross_returns_amount" bson:"gross_returns_amount"`
	// @inject_tag: json:"gross_total_amount" bson:"gross_total_amount"
	//
	// The gross revenue.
	GrossTotalAmount float64 `protobuf:"fixed64,9,opt,name=gross_total_amount,json=grossTotalAmount,proto3" json:"gross_total_amount" bson:"gross_total_amount"`
	// @inject_tag: json:"transactions_currency" bson:"transactions_currency"
	//
	// The transactions currency. Three-letter currency code in ISO 4217, in uppercase.
	TransactionsCurrency string `protobuf:"bytes,10,opt,name=transactions_currency,json=transactionsCurrency,proto3" json:"transactions_currency" bson:"transactions_currency"`
	// @inject_tag: json:"conversion" bson:"-"
	//
	// The conversion - sales per visits relation.
	Conversion float64 `protobuf:"fixed64,11,opt,name=conversion,proto3" json:"conversion" bson:"-"`
	// @inject_tag: json:"country_code" bson:"-"
	//
	// Two-letter country code in ISO 3166-1, in uppercase (for instance US).
	CountryCode string `protobuf:"bytes,12,opt,name=country_code,json=countryCode,proto3" json:"country_code" bson:"-"`
	// @inject_tag: json:"date" bson:"-"
	//
	// The date of the summary.
	Date string `protobuf:"bytes,13,opt,name=date,proto3" json:"date" bson:"-"`
	// @inject_tag: json:"referrer_host" bson:"-"
	//
	// The address of the webpage where the request originated.
	ReferrerHost string `protobuf:"bytes,14,opt,name=referrer_host,json=referrerHost,proto3" json:"referrer_host" bson:"-"`
	// @inject_tag: json:"utm" bson:"-"
	//
	// The UTM-tags list.
	Utm *Utm `protobuf:"bytes,15,opt,name=utm,proto3" json:"utm" bson:"-"`
}

func (x *StatCommon) Reset() {
	*x = StatCommon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paylink_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatCommon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatCommon) ProtoMessage() {}

func (x *StatCommon) ProtoReflect() protoreflect.Message {
	mi := &file_paylink_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatCommon.ProtoReflect.Descriptor instead.
func (*StatCommon) Descriptor() ([]byte, []int) {
	return file_paylink_proto_rawDescGZIP(), []int{2}
}

func (x *StatCommon) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StatCommon) GetPaylinkId() string {
	if x != nil {
		return x.PaylinkId
	}
	return ""
}

func (x *StatCommon) GetVisits() int32 {
	if x != nil {
		return x.Visits
	}
	return 0
}

func (x *StatCommon) GetTotalTransactions() int32 {
	if x != nil {
		return x.TotalTransactions
	}
	return 0
}

func (x *StatCommon) GetSalesCount() int32 {
	if x != nil {
		return x.SalesCount
	}
	return 0
}

func (x *StatCommon) GetReturnsCount() int32 {
	if x != nil {
		return x.ReturnsCount
	}
	return 0
}

func (x *StatCommon) GetGrossSalesAmount() float64 {
	if x != nil {
		return x.GrossSalesAmount
	}
	return 0
}

func (x *StatCommon) GetGrossReturnsAmount() float64 {
	if x != nil {
		return x.GrossReturnsAmount
	}
	return 0
}

func (x *StatCommon) GetGrossTotalAmount() float64 {
	if x != nil {
		return x.GrossTotalAmount
	}
	return 0
}

func (x *StatCommon) GetTransactionsCurrency() string {
	if x != nil {
		return x.TransactionsCurrency
	}
	return ""
}

func (x *StatCommon) GetConversion() float64 {
	if x != nil {
		return x.Conversion
	}
	return 0
}

func (x *StatCommon) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *StatCommon) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *StatCommon) GetReferrerHost() string {
	if x != nil {
		return x.ReferrerHost
	}
	return ""
}

func (x *StatCommon) GetUtm() *Utm {
	if x != nil {
		return x.Utm
	}
	return nil
}

type Utm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"utm_source"
	//
	// The UTM-tag of the advertising system, for example: Bing Ads, Google Adwords.
	UtmSource string `protobuf:"bytes,1,opt,name=utm_source,json=utmSource,proto3" json:"utm_source"`
	// @inject_tag: json:"utm_medium"
	//
	// The UTM-tag of the traffic type, e.g.: cpc, cpm, email newsletter.
	UtmMedium string `protobuf:"bytes,2,opt,name=utm_medium,json=utmMedium,proto3" json:"utm_medium"`
	// @inject_tag: json:"utm_campaign"
	//
	// The UTM-tag of the advertising campaign, for example: Online games, Simulation game.
	UtmCampaign string `protobuf:"bytes,3,opt,name=utm_campaign,json=utmCampaign,proto3" json:"utm_campaign"`
}

func (x *Utm) Reset() {
	*x = Utm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paylink_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Utm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Utm) ProtoMessage() {}

func (x *Utm) ProtoReflect() protoreflect.Message {
	mi := &file_paylink_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Utm.ProtoReflect.Descriptor instead.
func (*Utm) Descriptor() ([]byte, []int) {
	return file_paylink_proto_rawDescGZIP(), []int{3}
}

func (x *Utm) GetUtmSource() string {
	if x != nil {
		return x.UtmSource
	}
	return ""
}

func (x *Utm) GetUtmMedium() string {
	if x != nil {
		return x.UtmMedium
	}
	return ""
}

func (x *Utm) GetUtmCampaign() string {
	if x != nil {
		return x.UtmCampaign
	}
	return ""
}

type GroupStatCommon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"top" bson:"top"
	//
	// The list of statistical results for the period.
	Top []*StatCommon `protobuf:"bytes,1,rep,name=top,proto3" json:"top" bson:"top"`
	// @inject_tag: json:"total" bson:"total"
	//
	// The statistical results for the country.
	Total *StatCommon `protobuf:"bytes,2,opt,name=total,proto3" json:"total" bson:"total"`
}

func (x *GroupStatCommon) Reset() {
	*x = GroupStatCommon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paylink_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupStatCommon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupStatCommon) ProtoMessage() {}

func (x *GroupStatCommon) ProtoReflect() protoreflect.Message {
	mi := &file_paylink_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupStatCommon.ProtoReflect.Descriptor instead.
func (*GroupStatCommon) Descriptor() ([]byte, []int) {
	return file_paylink_proto_rawDescGZIP(), []int{4}
}

func (x *GroupStatCommon) GetTop() []*StatCommon {
	if x != nil {
		return x.Top
	}
	return nil
}

func (x *GroupStatCommon) GetTotal() *StatCommon {
	if x != nil {
		return x.Total
	}
	return nil
}

var File_paylink_proto protoreflect.FileDescriptor

var file_paylink_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x69, 0x6e, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x6f, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e, 0x6f, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x54, 0x79, 0x70, 0x65, 0x22, 0xe5, 0x06, 0x0a,
	0x07, 0x50, 0x61, 0x79, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x69, 0x73, 0x69, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x76, 0x69, 0x73, 0x69, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x6f, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x6e, 0x6f, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x2d, 0x0a,
	0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10,
	0x67, 0x72, 0x6f, 0x73, 0x73, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x30, 0x0a, 0x14, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12,
	0x67, 0x72, 0x6f, 0x73, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10,
	0x67, 0x72, 0x6f, 0x73, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x33, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x22, 0xa7, 0x04, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6c, 0x69, 0x6e, 0x6b,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x69, 0x73, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x76, 0x69, 0x73, 0x69, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6c,
	0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x12, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x67, 0x72, 0x6f,
	0x73, 0x73, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a,
	0x14, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x67, 0x72, 0x6f,
	0x73, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x12, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x67, 0x72, 0x6f,
	0x73, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a,
	0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x03, 0x75, 0x74, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x55, 0x74, 0x6d, 0x52, 0x03, 0x75, 0x74, 0x6d, 0x22, 0x66,
	0x0a, 0x03, 0x55, 0x74, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x74, 0x6d, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x74, 0x6d, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x74, 0x6d, 0x5f, 0x6d, 0x65, 0x64, 0x69,
	0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x74, 0x6d, 0x4d, 0x65, 0x64,
	0x69, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x74, 0x6d, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x74, 0x6d, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x22, 0x63, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x03, 0x74, 0x6f, 0x70,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x6f, 0x70,
	0x12, 0x29, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x0e, 0x5a, 0x0c, 0x2e,
	0x2f, 0x3b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_paylink_proto_rawDescOnce sync.Once
	file_paylink_proto_rawDescData = file_paylink_proto_rawDesc
)

func file_paylink_proto_rawDescGZIP() []byte {
	file_paylink_proto_rawDescOnce.Do(func() {
		file_paylink_proto_rawDescData = protoimpl.X.CompressGZIP(file_paylink_proto_rawDescData)
	})
	return file_paylink_proto_rawDescData
}

var file_paylink_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_paylink_proto_goTypes = []interface{}{
	(*CreatePaylinkRequest)(nil), // 0: paylink.CreatePaylinkRequest
	(*Paylink)(nil),              // 1: paylink.Paylink
	(*StatCommon)(nil),           // 2: paylink.StatCommon
	(*Utm)(nil),                  // 3: paylink.Utm
	(*GroupStatCommon)(nil),      // 4: paylink.GroupStatCommon
	(*timestamp.Timestamp)(nil),  // 5: google.protobuf.Timestamp
}
var file_paylink_proto_depIdxs = []int32{
	5, // 0: paylink.Paylink.expires_at:type_name -> google.protobuf.Timestamp
	5, // 1: paylink.Paylink.created_at:type_name -> google.protobuf.Timestamp
	5, // 2: paylink.Paylink.updated_at:type_name -> google.protobuf.Timestamp
	3, // 3: paylink.StatCommon.utm:type_name -> paylink.Utm
	2, // 4: paylink.GroupStatCommon.top:type_name -> paylink.StatCommon
	2, // 5: paylink.GroupStatCommon.total:type_name -> paylink.StatCommon
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_paylink_proto_init() }
func file_paylink_proto_init() {
	if File_paylink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paylink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaylinkRequest); i {
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
		file_paylink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paylink); i {
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
		file_paylink_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatCommon); i {
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
		file_paylink_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Utm); i {
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
		file_paylink_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupStatCommon); i {
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
			RawDescriptor: file_paylink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_paylink_proto_goTypes,
		DependencyIndexes: file_paylink_proto_depIdxs,
		MessageInfos:      file_paylink_proto_msgTypes,
	}.Build()
	File_paylink_proto = out.File
	file_paylink_proto_rawDesc = nil
	file_paylink_proto_goTypes = nil
	file_paylink_proto_depIdxs = nil
}
