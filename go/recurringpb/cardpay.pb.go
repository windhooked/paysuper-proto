// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: cardpay.proto

package recurringpb

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

type CardPayAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Phone   string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	State   string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Street  string `protobuf:"bytes,5,opt,name=street,proto3" json:"street,omitempty"`
	Zip     string `protobuf:"bytes,6,opt,name=zip,proto3" json:"zip,omitempty"`
}

func (x *CardPayAddress) Reset() {
	*x = CardPayAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardpay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardPayAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardPayAddress) ProtoMessage() {}

func (x *CardPayAddress) ProtoReflect() protoreflect.Message {
	mi := &file_cardpay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardPayAddress.ProtoReflect.Descriptor instead.
func (*CardPayAddress) Descriptor() ([]byte, []int) {
	return file_cardpay_proto_rawDescGZIP(), []int{0}
}

func (x *CardPayAddress) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CardPayAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CardPayAddress) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CardPayAddress) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CardPayAddress) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *CardPayAddress) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

type CardPayItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Count       int32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CardPayItem) Reset() {
	*x = CardPayItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardpay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardPayItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardPayItem) ProtoMessage() {}

func (x *CardPayItem) ProtoReflect() protoreflect.Message {
	mi := &file_cardpay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardPayItem.ProtoReflect.Descriptor instead.
func (*CardPayItem) Descriptor() ([]byte, []int) {
	return file_cardpay_proto_rawDescGZIP(), []int{1}
}

func (x *CardPayItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CardPayItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CardPayItem) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CardPayItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CardPayMerchantOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description     string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Items           []*CardPayItem  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	ShippingAddress *CardPayAddress `protobuf:"bytes,4,opt,name=shipping_address,json=shippingAddress,proto3" json:"shipping_address,omitempty"`
}

func (x *CardPayMerchantOrder) Reset() {
	*x = CardPayMerchantOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardpay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardPayMerchantOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardPayMerchantOrder) ProtoMessage() {}

func (x *CardPayMerchantOrder) ProtoReflect() protoreflect.Message {
	mi := &file_cardpay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardPayMerchantOrder.ProtoReflect.Descriptor instead.
func (*CardPayMerchantOrder) Descriptor() ([]byte, []int) {
	return file_cardpay_proto_rawDescGZIP(), []int{2}
}

func (x *CardPayMerchantOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CardPayMerchantOrder) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CardPayMerchantOrder) GetItems() []*CardPayItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CardPayMerchantOrder) GetShippingAddress() *CardPayAddress {
	if x != nil {
		return x.ShippingAddress
	}
	return nil
}

type CallbackCardPayBankCardAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Holder             string `protobuf:"bytes,1,opt,name=holder,proto3" json:"holder,omitempty"`
	IssuingCountryCode string `protobuf:"bytes,2,opt,name=issuing_country_code,json=issuingCountryCode,proto3" json:"issuing_country_code,omitempty"`
	MaskedPan          string `protobuf:"bytes,3,opt,name=masked_pan,json=maskedPan,proto3" json:"masked_pan,omitempty"`
	Token              string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CallbackCardPayBankCardAccount) Reset() {
	*x = CallbackCardPayBankCardAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardpay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackCardPayBankCardAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackCardPayBankCardAccount) ProtoMessage() {}

func (x *CallbackCardPayBankCardAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cardpay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackCardPayBankCardAccount.ProtoReflect.Descriptor instead.
func (*CallbackCardPayBankCardAccount) Descriptor() ([]byte, []int) {
	return file_cardpay_proto_rawDescGZIP(), []int{3}
}

func (x *CallbackCardPayBankCardAccount) GetHolder() string {
	if x != nil {
		return x.Holder
	}
	return ""
}

func (x *CallbackCardPayBankCardAccount) GetIssuingCountryCode() string {
	if x != nil {
		return x.IssuingCountryCode
	}
	return ""
}

func (x *CallbackCardPayBankCardAccount) GetMaskedPan() string {
	if x != nil {
		return x.MaskedPan
	}
	return ""
}

func (x *CallbackCardPayBankCardAccount) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CallbackCardPayCryptoCurrencyAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CryptoAddress       string `protobuf:"bytes,1,opt,name=crypto_address,json=cryptoAddress,proto3" json:"crypto_address,omitempty"`
	CryptoTransactionId string `protobuf:"bytes,2,opt,name=crypto_transaction_id,json=cryptoTransactionId,proto3" json:"crypto_transaction_id,omitempty"`
	PrcAmount           string `protobuf:"bytes,3,opt,name=prc_amount,json=prcAmount,proto3" json:"prc_amount,omitempty"`
	PrcCurrency         string `protobuf:"bytes,4,opt,name=prc_currency,json=prcCurrency,proto3" json:"prc_currency,omitempty"`
}

func (x *CallbackCardPayCryptoCurrencyAccount) Reset() {
	*x = CallbackCardPayCryptoCurrencyAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardpay_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackCardPayCryptoCurrencyAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackCardPayCryptoCurrencyAccount) ProtoMessage() {}

func (x *CallbackCardPayCryptoCurrencyAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cardpay_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackCardPayCryptoCurrencyAccount.ProtoReflect.Descriptor instead.
func (*CallbackCardPayCryptoCurrencyAccount) Descriptor() ([]byte, []int) {
	return file_cardpay_proto_rawDescGZIP(), []int{4}
}

func (x *CallbackCardPayCryptoCurrencyAccount) GetCryptoAddress() string {
	if x != nil {
		return x.CryptoAddress
	}
	return ""
}

func (x *CallbackCardPayCryptoCurrencyAccount) GetCryptoTransactionId() string {
	if x != nil {
		return x.CryptoTransactionId
	}
	return ""
}

func (x *CallbackCardPayCryptoCurrencyAccount) GetPrcAmount() string {
	if x != nil {
		return x.PrcAmount
	}
	return ""
}

func (x *CallbackCardPayCryptoCurrencyAccount) GetPrcCurrency() string {
	if x != nil {
		return x.PrcCurrency
	}
	return ""
}

type CardPayCustomer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Ip     string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Id     string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Locale string `protobuf:"bytes,4,opt,name=locale,proto3" json:"locale,omitempty"`
}

func (x *CardPayCustomer) Reset() {
	*x = CardPayCustomer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardpay_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardPayCustomer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardPayCustomer) ProtoMessage() {}

func (x *CardPayCustomer) ProtoReflect() protoreflect.Message {
	mi := &file_cardpay_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardPayCustomer.ProtoReflect.Descriptor instead.
func (*CardPayCustomer) Descriptor() ([]byte, []int) {
	return file_cardpay_proto_rawDescGZIP(), []int{5}
}

func (x *CardPayCustomer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CardPayCustomer) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CardPayCustomer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CardPayCustomer) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

type CardPayEWalletAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CardPayEWalletAccount) Reset() {
	*x = CardPayEWalletAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardpay_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardPayEWalletAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardPayEWalletAccount) ProtoMessage() {}

func (x *CardPayEWalletAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cardpay_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardPayEWalletAccount.ProtoReflect.Descriptor instead.
func (*CardPayEWalletAccount) Descriptor() ([]byte, []int) {
	return file_cardpay_proto_rawDescGZIP(), []int{6}
}

func (x *CardPayEWalletAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CallbackCardPayPaymentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount        float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	AuthCode      string  `protobuf:"bytes,3,opt,name=auth_code,json=authCode,proto3" json:"auth_code,omitempty"`
	Created       string  `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Currency      string  `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	DeclineCode   string  `protobuf:"bytes,6,opt,name=decline_code,json=declineCode,proto3" json:"decline_code,omitempty"`
	DeclineReason string  `protobuf:"bytes,7,opt,name=decline_reason,json=declineReason,proto3" json:"decline_reason,omitempty"`
	Description   string  `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Is_3D         bool    `protobuf:"varint,9,opt,name=is_3d,json=is3d,proto3" json:"is_3d,omitempty"`
	Note          string  `protobuf:"bytes,10,opt,name=note,proto3" json:"note,omitempty"`
	Rrn           string  `protobuf:"bytes,11,opt,name=rrn,proto3" json:"rrn,omitempty"`
	Status        string  `protobuf:"bytes,12,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CallbackCardPayPaymentData) Reset() {
	*x = CallbackCardPayPaymentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardpay_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackCardPayPaymentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackCardPayPaymentData) ProtoMessage() {}

func (x *CallbackCardPayPaymentData) ProtoReflect() protoreflect.Message {
	mi := &file_cardpay_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackCardPayPaymentData.ProtoReflect.Descriptor instead.
func (*CallbackCardPayPaymentData) Descriptor() ([]byte, []int) {
	return file_cardpay_proto_rawDescGZIP(), []int{7}
}

func (x *CallbackCardPayPaymentData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CallbackCardPayPaymentData) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CallbackCardPayPaymentData) GetAuthCode() string {
	if x != nil {
		return x.AuthCode
	}
	return ""
}

func (x *CallbackCardPayPaymentData) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *CallbackCardPayPaymentData) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CallbackCardPayPaymentData) GetDeclineCode() string {
	if x != nil {
		return x.DeclineCode
	}
	return ""
}

func (x *CallbackCardPayPaymentData) GetDeclineReason() string {
	if x != nil {
		return x.DeclineReason
	}
	return ""
}

func (x *CallbackCardPayPaymentData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CallbackCardPayPaymentData) GetIs_3D() bool {
	if x != nil {
		return x.Is_3D
	}
	return false
}

func (x *CallbackCardPayPaymentData) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *CallbackCardPayPaymentData) GetRrn() string {
	if x != nil {
		return x.Rrn
	}
	return ""
}

func (x *CallbackCardPayPaymentData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CardPayPaymentCallback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantOrder         *CardPayMerchantOrder                 `protobuf:"bytes,1,opt,name=merchant_order,json=merchantOrder,proto3" json:"merchant_order,omitempty"`
	PaymentMethod         string                                `protobuf:"bytes,2,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	CallbackTime          string                                `protobuf:"bytes,3,opt,name=callback_time,json=callbackTime,proto3" json:"callback_time,omitempty"`
	CardAccount           *CallbackCardPayBankCardAccount       `protobuf:"bytes,4,opt,name=card_account,json=cardAccount,proto3" json:"card_account,omitempty"`
	CryptocurrencyAccount *CallbackCardPayCryptoCurrencyAccount `protobuf:"bytes,5,opt,name=cryptocurrency_account,json=cryptocurrencyAccount,proto3" json:"cryptocurrency_account,omitempty"`
	Customer              *CardPayCustomer                      `protobuf:"bytes,6,opt,name=customer,proto3" json:"customer,omitempty"`
	EwalletAccount        *CardPayEWalletAccount                `protobuf:"bytes,7,opt,name=ewallet_account,json=ewalletAccount,proto3" json:"ewallet_account,omitempty"`
	PaymentData           *CallbackCardPayPaymentData           `protobuf:"bytes,8,opt,name=payment_data,json=paymentData,proto3" json:"payment_data,omitempty"`
	// @inject_tag: json:"-"
	Signature string `protobuf:"bytes,9,opt,name=signature,proto3" json:"-"`
}

func (x *CardPayPaymentCallback) Reset() {
	*x = CardPayPaymentCallback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardpay_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardPayPaymentCallback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardPayPaymentCallback) ProtoMessage() {}

func (x *CardPayPaymentCallback) ProtoReflect() protoreflect.Message {
	mi := &file_cardpay_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardPayPaymentCallback.ProtoReflect.Descriptor instead.
func (*CardPayPaymentCallback) Descriptor() ([]byte, []int) {
	return file_cardpay_proto_rawDescGZIP(), []int{8}
}

func (x *CardPayPaymentCallback) GetMerchantOrder() *CardPayMerchantOrder {
	if x != nil {
		return x.MerchantOrder
	}
	return nil
}

func (x *CardPayPaymentCallback) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *CardPayPaymentCallback) GetCallbackTime() string {
	if x != nil {
		return x.CallbackTime
	}
	return ""
}

func (x *CardPayPaymentCallback) GetCardAccount() *CallbackCardPayBankCardAccount {
	if x != nil {
		return x.CardAccount
	}
	return nil
}

func (x *CardPayPaymentCallback) GetCryptocurrencyAccount() *CallbackCardPayCryptoCurrencyAccount {
	if x != nil {
		return x.CryptocurrencyAccount
	}
	return nil
}

func (x *CardPayPaymentCallback) GetCustomer() *CardPayCustomer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *CardPayPaymentCallback) GetEwalletAccount() *CardPayEWalletAccount {
	if x != nil {
		return x.EwalletAccount
	}
	return nil
}

func (x *CardPayPaymentCallback) GetPaymentData() *CallbackCardPayPaymentData {
	if x != nil {
		return x.PaymentData
	}
	return nil
}

func (x *CardPayPaymentCallback) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_cardpay_proto protoreflect.FileDescriptor

var file_cardpay_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x72, 0x64, 0x70, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x64,
	0x50, 0x61, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x7a, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x7a, 0x69, 0x70, 0x22, 0x6f,
	0x0a, 0x0b, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22,
	0xb6, 0x01, 0x0a, 0x14, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x41, 0x0a, 0x10, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x1e, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x42, 0x61, 0x6e, 0x6b,
	0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x5f,
	0x70, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x73, 0x6b, 0x65,
	0x64, 0x50, 0x61, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc3, 0x01, 0x0a, 0x24, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x63, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x63, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x72, 0x63, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x63, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x22, 0x5f, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x65, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x45, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd6, 0x02, 0x0a, 0x1a, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x63, 0x6c,
	0x69, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x63, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x33, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x69, 0x73, 0x33, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x72, 0x6e,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x72, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0xbb, 0x04, 0x0a, 0x16, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x43,
	0x0a, 0x0e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x0d, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x49, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x42, 0x61,
	0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b, 0x63,
	0x61, 0x72, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x63, 0x0a, 0x16, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64,
	0x50, 0x61, 0x79, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x15, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x33, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x50,
	0x61, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0f, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x45, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x0c,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cardpay_proto_rawDescOnce sync.Once
	file_cardpay_proto_rawDescData = file_cardpay_proto_rawDesc
)

func file_cardpay_proto_rawDescGZIP() []byte {
	file_cardpay_proto_rawDescOnce.Do(func() {
		file_cardpay_proto_rawDescData = protoimpl.X.CompressGZIP(file_cardpay_proto_rawDescData)
	})
	return file_cardpay_proto_rawDescData
}

var file_cardpay_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cardpay_proto_goTypes = []interface{}{
	(*CardPayAddress)(nil),                       // 0: entity.CardPayAddress
	(*CardPayItem)(nil),                          // 1: entity.CardPayItem
	(*CardPayMerchantOrder)(nil),                 // 2: entity.CardPayMerchantOrder
	(*CallbackCardPayBankCardAccount)(nil),       // 3: entity.CallbackCardPayBankCardAccount
	(*CallbackCardPayCryptoCurrencyAccount)(nil), // 4: entity.CallbackCardPayCryptoCurrencyAccount
	(*CardPayCustomer)(nil),                      // 5: entity.CardPayCustomer
	(*CardPayEWalletAccount)(nil),                // 6: entity.CardPayEWalletAccount
	(*CallbackCardPayPaymentData)(nil),           // 7: entity.CallbackCardPayPaymentData
	(*CardPayPaymentCallback)(nil),               // 8: entity.CardPayPaymentCallback
}
var file_cardpay_proto_depIdxs = []int32{
	1, // 0: entity.CardPayMerchantOrder.items:type_name -> entity.CardPayItem
	0, // 1: entity.CardPayMerchantOrder.shipping_address:type_name -> entity.CardPayAddress
	2, // 2: entity.CardPayPaymentCallback.merchant_order:type_name -> entity.CardPayMerchantOrder
	3, // 3: entity.CardPayPaymentCallback.card_account:type_name -> entity.CallbackCardPayBankCardAccount
	4, // 4: entity.CardPayPaymentCallback.cryptocurrency_account:type_name -> entity.CallbackCardPayCryptoCurrencyAccount
	5, // 5: entity.CardPayPaymentCallback.customer:type_name -> entity.CardPayCustomer
	6, // 6: entity.CardPayPaymentCallback.ewallet_account:type_name -> entity.CardPayEWalletAccount
	7, // 7: entity.CardPayPaymentCallback.payment_data:type_name -> entity.CallbackCardPayPaymentData
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cardpay_proto_init() }
func file_cardpay_proto_init() {
	if File_cardpay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cardpay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardPayAddress); i {
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
		file_cardpay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardPayItem); i {
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
		file_cardpay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardPayMerchantOrder); i {
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
		file_cardpay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackCardPayBankCardAccount); i {
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
		file_cardpay_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackCardPayCryptoCurrencyAccount); i {
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
		file_cardpay_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardPayCustomer); i {
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
		file_cardpay_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardPayEWalletAccount); i {
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
		file_cardpay_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackCardPayPaymentData); i {
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
		file_cardpay_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardPayPaymentCallback); i {
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
			RawDescriptor: file_cardpay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cardpay_proto_goTypes,
		DependencyIndexes: file_cardpay_proto_depIdxs,
		MessageInfos:      file_cardpay_proto_msgTypes,
	}.Build()
	File_cardpay_proto = out.File
	file_cardpay_proto_rawDesc = nil
	file_cardpay_proto_goTypes = nil
	file_cardpay_proto_depIdxs = nil
}
