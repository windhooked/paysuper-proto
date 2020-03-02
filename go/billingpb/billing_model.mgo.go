package billingpb

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/paysuper/paysuper-proto/go/recurringpb"
	tools "github.com/paysuper/paysuper-tools/number"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MgoId struct {
	Id primitive.ObjectID `bson:"_id"`
}

type MgoMultiLang struct {
	Lang  string `bson:"lang"`
	Value string `bson:"value"`
}

type MgoOrderProject struct {
	Id                      primitive.ObjectID       `bson:"_id"` 
	MerchantId              primitive.ObjectID       `bson:"merchant_id"`
	Name                    []*MgoMultiLang          `bson:"name"`
	UrlSuccess              string                   `bson:"url_success"`
	UrlFail                 string                   `bson:"url_fail"`
	NotifyEmails            []string                 `bson:"notify_emails"`
	SecretKey               string                   `bson:"secret_key"`
	SendNotifyEmail         bool                     `bson:"send_notify_email"`
	UrlCheckAccount         string                   `bson:"url_check_account"`
	UrlProcessPayment       string                   `bson:"url_process_payment"`
	CallbackProtocol        string                   `bson:"callback_protocol"`
	UrlChargebackPayment    string                   `bson:"url_chargeback_payment"`
	UrlCancelPayment        string                   `bson:"url_cancel_payment"`
	UrlFraudPayment         string                   `bson:"url_fraud_payment"`
	UrlRefundPayment        string                   `bson:"url_refund_payment"`
	Status                  int32                    `bson:"status"`
	MerchantRoyaltyCurrency string                   `bson:"merchant_royalty_currency"`
	RedirectSettings        *ProjectRedirectSettings `bson:"redirect_settings"`
}

type MgoOrderPaymentMethod struct {
	Id              primitive.ObjectID   `bson:"_id"`
	Name            string               `bson:"name"`
	Handler         string               `bson:"handler"`
	ExternalId      string               `bson:"external_id"`
	Params          *PaymentMethodParams `bson:"params"`
	PaymentSystemId primitive.ObjectID   `bson:"payment_system_id"`
	Group           string               `bson:"group_alias"`
	Saved           bool                 `bson:"saved"`
	Card            *PaymentMethodCard   `bson:"card,omitempty"`
	Wallet          *PaymentMethodWallet `bson:"wallet,omitempty"`
	CryptoCurrency  *PaymentMethodCrypto `bson:"crypto_currency,omitempty"`
	RefundAllowed   bool                 `bson:"refund_allowed"`
}

type MgoOrderNotificationRefund struct {
	Amount        float64 `bson:"amount"`
	Currency      string  `bson:"currency"`
	Reason        string  `bson:"reason"`
	Code          string  `bson:"code"`
	ReceiptNumber string  `bson:"receipt_number"`
	ReceiptUrl    string  `bson:"receipt_url"`
}

type MgoOrder struct {
	Id                          primitive.ObjectID             `bson:"_id"`
	Uuid                        string                         `bson:"uuid"`
	Transaction                 string                         `bson:"pm_order_id"`
	Object                      string                         `bson:"object"`
	Status                      string                         `bson:"status"`
	PrivateStatus               int32                          `bson:"private_status"`
	Description                 string                         `bson:"description"`
	CreatedAt                   time.Time                      `bson:"created_at"`
	UpdatedAt                   time.Time                      `bson:"updated_at"`
	CanceledAt                  time.Time                      `bson:"canceled_at"`
	Canceled                    bool                           `bson:"canceled"`
	Cancellation                *OrderNotificationCancellation `bson:"cancellation"`
	Refunded                    bool                           `bson:"refunded"`
	RefundedAt                  time.Time                      `bson:"refunded_at"`
	ReceiptEmail                string                         `bson:"receipt_email"`
	ReceiptPhone                string                         `bson:"receipt_phone"`
	ReceiptNumber               string                         `bson:"receipt_number"`
	ReceiptUrl                  string                         `bson:"receipt_url"`
	AgreementVersion            string                         `bson:"agreement_version"`
	AgreementAccepted           bool                           `bson:"agreement_accepted"`
	NotifySale                  bool                           `bson:"notify_sale"`
	NotifySaleEmail             string                         `bson:"notify_sale_email"`
	Issuer                      *OrderIssuer                   `bson:"issuer"`
	TotalPaymentAmount          float64                        `bson:"total_payment_amount"`
	Currency                    string                         `bson:"currency"`
	User                        *OrderUser                     `bson:"user"`
	BillingAddress              *OrderBillingAddress           `bson:"billing_address"`
	Tax                         *OrderTax                      `bson:"tax"`
	PaymentMethod               *MgoOrderPaymentMethod         `bson:"payment_method"`
	Items                       []*MgoOrderItem                `bson:"items"`
	Refund                      *MgoOrderNotificationRefund    `bson:"refund"`
	Metadata                    map[string]string              `bson:"metadata"`
	PrivateMetadata             map[string]string              `bson:"private_metadata"`
	Project                     *MgoOrderProject               `bson:"project"`
	ProjectOrderId              string                         `bson:"project_order_id"`
	ProjectAccount              string                         `bson:"project_account"`
	ProjectLastRequestedAt      time.Time                      `bson:"project_last_requested_at"`
	ProjectParams               map[string]string              `bson:"project_params"`
	PaymentMethodOrderClosedAt  time.Time                      `bson:"pm_order_close_date"`
	IsJsonRequest               bool                           `bson:"created_by_json"`
	OrderAmount                 float64                        `bson:"private_amount"`
	PaymentMethodPayerAccount   string                         `bson:"pm_account"`
	PaymentMethodTxnParams      map[string]string              `bson:"pm_txn_params"`
	PaymentRequisites           map[string]string              `bson:"payment_requisites"`
	ExpireDateToFormInput       time.Time                      `bson:"expire_date_to_form_input"`
	UserAddressDataRequired     bool                           `bson:"user_address_data_required"`
	Products                    []string                       `bson:"products"`
	IsNotificationsSent         map[string]bool                `bson:"is_notifications_sent"`
	CountryRestriction          *CountryRestriction            `bson:"country_restriction"`
	ParentOrder                 *ParentOrder                   `bson:"parent_order"`
	ParentPaymentAt             time.Time                      `bson:"parent_payment_at"`
	Type                        string                         `bson:"type"`
	IsVatDeduction              bool                           `bson:"is_vat_deduction"`
	CountryCode                 string                         `bson:"country_code"`
	PlatformId                  string                         `bson:"platform_id"`
	ProductType                 string                         `bson:"product_type"`
	Keys                        []string                       `bson:"keys"`
	IsKeyProductNotified        bool                           `bson:"is_key_product_notified"`
	ReceiptId                   string                         `bson:"receipt_id"`
	IsBuyForVirtualCurrency     bool                           `bson:"is_buy_for_virtual_currency"`
	MccCode                     string                         `bson:"mcc_code"`
	OperatingCompanyId          string                         `bson:"operating_company_id"`
	IsHighRisk                  bool                           `bson:"is_high_risk"`
	ChargeCurrency              string                         `bson:"charge_currency"`
	ChargeAmount                float64                        `bson:"charge_amount"`
	PaymentIpCountry            string                         `bson:"payment_ip_country"`
	IsIpCountryMismatchBin      bool                           `bson:"is_ip_country_mismatch_bin"`
	BillingCountryChangedByUser bool                           `bson:"billing_country_changed_by_user"`
	IsRefundAllowed             bool                           `bson:"is_refund_allowed"`
	VatPayer                    string                         `bson:"vat_payer"`
	IsProduction                bool                           `bson:"is_production"`
	FormMode                    string                         `bson:"form_mode"`
}

type MgoOrderItem struct {
	Id          primitive.ObjectID `bson:"_id"`
	Object      string             `bson:"object"`
	Sku         string             `bson:"sku"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Amount      float64            `bson:"amount"`
	Currency    string             `bson:"currency"`
	Images      []string           `bson:"images"`
	Url         string             `bson:"url"`
	Metadata    map[string]string  `bson:"metadata"`
	Code        string             `bson:"code"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	PlatformId  string             `bson:"platform_id"`
}

func (m *Order) MarshalBSON() ([]byte, error) {
	projectOid, err := primitive.ObjectIDFromHex(m.Project.Id)

	if err != nil {
		return nil, err
	}

	merchantOid, err := primitive.ObjectIDFromHex(m.Project.MerchantId)

	if err != nil {
		return nil, err
	}

	st := &MgoOrder{
		Uuid:               m.Uuid,
		Transaction:        m.Transaction,
		Object:             "order",
		Status:             m.GetPublicStatus(),
		PrivateStatus:      m.PrivateStatus,
		Description:        m.Description,
		Canceled:           m.PrivateStatus == recurringpb.OrderStatusPaymentSystemCanceled,
		Cancellation:       m.Cancellation,
		Refunded:           m.PrivateStatus == recurringpb.OrderStatusRefund,
		ReceiptEmail:       m.GetReceiptUserEmail(),
		ReceiptPhone:       m.GetReceiptUserPhone(),
		ReceiptNumber:      m.ReceiptNumber,
		ReceiptUrl:         m.ReceiptUrl,
		AgreementVersion:   m.AgreementVersion,
		AgreementAccepted:  m.AgreementAccepted,
		NotifySale:         m.NotifySale,
		NotifySaleEmail:    m.NotifySaleEmail,
		Issuer:             m.Issuer,
		TotalPaymentAmount: m.TotalPaymentAmount,
		Currency:           m.Currency,
		User:               m.User,
		BillingAddress:     m.BillingAddress,
		Tax:                m.Tax,
		Items:              []*MgoOrderItem{},
		Metadata:           m.Metadata,
		PrivateMetadata:    m.PrivateMetadata,
		Project: &MgoOrderProject{
			Id:                      projectOid,
			MerchantId:              merchantOid,
			UrlSuccess:              m.Project.UrlSuccess,
			UrlFail:                 m.Project.UrlFail,
			NotifyEmails:            m.Project.NotifyEmails,
			SendNotifyEmail:         m.Project.SendNotifyEmail,
			SecretKey:               m.Project.SecretKey,
			UrlCheckAccount:         m.Project.UrlCheckAccount,
			UrlProcessPayment:       m.Project.UrlProcessPayment,
			CallbackProtocol:        m.Project.CallbackProtocol,
			UrlChargebackPayment:    m.Project.UrlChargebackPayment,
			UrlCancelPayment:        m.Project.UrlCancelPayment,
			UrlRefundPayment:        m.Project.UrlRefundPayment,
			UrlFraudPayment:         m.Project.UrlFraudPayment,
			Status:                  m.Project.Status,
			MerchantRoyaltyCurrency: m.Project.MerchantRoyaltyCurrency,
			RedirectSettings:        m.Project.RedirectSettings,
		},
		ProjectOrderId:              m.ProjectOrderId,
		ProjectAccount:              m.ProjectAccount,
		ProjectParams:               m.ProjectParams,
		IsJsonRequest:               m.IsJsonRequest,
		OrderAmount:                 m.OrderAmount,
		PaymentMethodPayerAccount:   m.PaymentMethodPayerAccount,
		PaymentMethodTxnParams:      m.PaymentMethodTxnParams,
		PaymentRequisites:           m.PaymentRequisites,
		UserAddressDataRequired:     m.UserAddressDataRequired,
		Products:                    m.Products,
		IsNotificationsSent:         m.IsNotificationsSent,
		CountryRestriction:          m.CountryRestriction,
		ParentOrder:                 m.ParentOrder,
		Type:                        m.Type,
		IsVatDeduction:              m.IsVatDeduction,
		CountryCode:                 m.GetCountry(),
		ProductType:                 m.ProductType,
		PlatformId:                  m.PlatformId,
		Keys:                        m.Keys,
		IsKeyProductNotified:        m.IsKeyProductNotified,
		ReceiptId:                   m.ReceiptId,
		IsBuyForVirtualCurrency:     m.IsBuyForVirtualCurrency,
		MccCode:                     m.MccCode,
		OperatingCompanyId:          m.OperatingCompanyId,
		IsHighRisk:                  m.IsHighRisk,
		ChargeCurrency:              m.ChargeCurrency,
		ChargeAmount:                m.ChargeAmount,
		PaymentIpCountry:            m.PaymentIpCountry,
		IsIpCountryMismatchBin:      m.IsIpCountryMismatchBin,
		BillingCountryChangedByUser: m.BillingCountryChangedByUser,
		IsRefundAllowed:             m.IsRefundAllowed,
		VatPayer:                    m.VatPayer,
		IsProduction:                m.IsProduction,
		FormMode:                    m.FormMode,
	}

	if m.Refund != nil {
		st.Refund = &MgoOrderNotificationRefund{
			Amount:        m.Refund.Amount,
			Currency:      m.Refund.Currency,
			Reason:        m.Refund.Reason,
			Code:          m.Refund.Code,
			ReceiptNumber: m.Refund.ReceiptNumber,
			ReceiptUrl:    m.Refund.ReceiptUrl,
		}
	}

	for _, v := range m.Items {
		item := &MgoOrderItem{
			Object:      v.Object,
			Sku:         v.Sku,
			Name:        v.Name,
			Description: v.Description,
			Amount:      v.Amount,
			Currency:    v.Currency,
			Images:      v.Images,
			Url:         v.Url,
			Metadata:    v.Metadata,
			Code:        v.Code,
			PlatformId:  v.PlatformId,
		}

		if len(v.Id) <= 0 {
			item.Id = primitive.NewObjectID()
		} else {
			itemOid, err := primitive.ObjectIDFromHex(v.Id)

			if err != nil {
				return nil, errors.New(ErrorInvalidObjectId)
			}
			item.Id = itemOid
		}

		item.CreatedAt, _ = ptypes.Timestamp(v.CreatedAt)
		item.CreatedAt, _ = ptypes.Timestamp(v.UpdatedAt)
		st.Items = append(st.Items, item)
	}

	if m.PaymentMethod != nil {
		paymentMethodOid, err := primitive.ObjectIDFromHex(m.PaymentMethod.Id)

		if err != nil {
			return nil, err
		}

		paymentSystemOid, err := primitive.ObjectIDFromHex(m.PaymentMethod.PaymentSystemId)

		if err != nil {
			return nil, err
		}

		st.PaymentMethod = &MgoOrderPaymentMethod{
			Id:              paymentMethodOid,
			Name:            m.PaymentMethod.Name,
			ExternalId:      m.PaymentMethod.ExternalId,
			Params:          m.PaymentMethod.Params,
			PaymentSystemId: paymentSystemOid,
			Group:           m.PaymentMethod.Group,
			Saved:           m.PaymentMethod.Saved,
			Handler:         m.PaymentMethod.Handler,
			RefundAllowed:   m.PaymentMethod.RefundAllowed,
		}

		if m.PaymentMethod.Card != nil {
			st.PaymentMethod.Card = m.PaymentMethod.Card
		}
		if m.PaymentMethod.Wallet != nil {
			st.PaymentMethod.Wallet = m.PaymentMethod.Wallet
		}
		if m.PaymentMethod.CryptoCurrency != nil {
			st.PaymentMethod.CryptoCurrency = m.PaymentMethod.CryptoCurrency
		}
	}

	if len(m.Id) <= 0 {
		st.Id = primitive.NewObjectID()
	} else {
		oid, err := primitive.ObjectIDFromHex(m.Id)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}

		st.Id = oid
	}

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	if m.UpdatedAt != nil {
		t, err := ptypes.Timestamp(m.UpdatedAt)

		if err != nil {
			return nil, err
		}

		st.UpdatedAt = t
	} else {
		st.UpdatedAt = time.Now()
	}

	if m.ProjectLastRequestedAt != nil {
		t, err := ptypes.Timestamp(m.ProjectLastRequestedAt)

		if err != nil {
			return nil, err
		}

		st.ProjectLastRequestedAt = t
	}

	if m.PaymentMethodOrderClosedAt != nil {
		t, err := ptypes.Timestamp(m.PaymentMethodOrderClosedAt)

		if err != nil {
			return nil, err
		}

		st.PaymentMethodOrderClosedAt = t
	}

	if m.ExpireDateToFormInput != nil {
		t, err := ptypes.Timestamp(m.ExpireDateToFormInput)

		if err != nil {
			return nil, err
		}

		st.ExpireDateToFormInput = t
	} else {
		st.ExpireDateToFormInput = time.Now()
	}

	if m.Project != nil && len(m.Project.Name) > 0 {
		for k, v := range m.Project.Name {
			st.Project.Name = append(st.Project.Name, &MgoMultiLang{Lang: k, Value: v})
		}
	}

	if m.CanceledAt != nil {
		t, err := ptypes.Timestamp(m.CanceledAt)

		if err != nil {
			return nil, err
		}

		st.CanceledAt = t
	}

	if m.RefundedAt != nil {
		t, err := ptypes.Timestamp(m.RefundedAt)

		if err != nil {
			return nil, err
		}

		st.RefundedAt = t
	}

	if m.ParentPaymentAt != nil {
		t, err := ptypes.Timestamp(m.ParentPaymentAt)

		if err != nil {
			return nil, err
		}

		st.ParentPaymentAt = t
	}

	return bson.Marshal(st)
}

func (m *Order) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoOrder)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Uuid = decoded.Uuid
	m.Transaction = decoded.Transaction
	m.Object = decoded.Object
	m.Status = decoded.Status
	m.PrivateStatus = decoded.PrivateStatus
	m.Description = decoded.Description
	m.Canceled = decoded.Canceled
	m.Cancellation = decoded.Cancellation
	m.Refunded = decoded.Refunded
	m.ReceiptEmail = decoded.ReceiptEmail
	m.ReceiptPhone = decoded.ReceiptPhone
	m.ReceiptNumber = decoded.ReceiptNumber
	m.ReceiptUrl = decoded.ReceiptUrl
	m.AgreementVersion = decoded.AgreementVersion
	m.AgreementAccepted = decoded.AgreementAccepted
	m.NotifySale = decoded.NotifySale
	m.NotifySaleEmail = decoded.NotifySaleEmail
	m.Issuer = decoded.Issuer
	m.TotalPaymentAmount = decoded.TotalPaymentAmount
	m.Currency = decoded.Currency
	m.User = decoded.User
	m.BillingAddress = decoded.BillingAddress
	m.Tax = decoded.Tax
	m.PaymentMethod = getPaymentMethodOrder(decoded.PaymentMethod)
	m.Items = []*OrderItem{}
	m.PlatformId = decoded.PlatformId
	m.ProductType = decoded.ProductType
	m.Keys = decoded.Keys
	m.IsKeyProductNotified = decoded.IsKeyProductNotified
	m.ReceiptId = decoded.ReceiptId
	m.IsBuyForVirtualCurrency = decoded.IsBuyForVirtualCurrency
	m.MccCode = decoded.MccCode
	m.OperatingCompanyId = decoded.OperatingCompanyId
	m.IsHighRisk = decoded.IsHighRisk
	m.ChargeCurrency = decoded.ChargeCurrency
	m.ChargeAmount = decoded.ChargeAmount
	m.PaymentIpCountry = decoded.PaymentIpCountry
	m.IsIpCountryMismatchBin = decoded.IsIpCountryMismatchBin
	m.BillingCountryChangedByUser = decoded.BillingCountryChangedByUser
	m.IsRefundAllowed = decoded.IsRefundAllowed
	m.FormMode = decoded.FormMode

	if decoded.Refund != nil {
		m.Refund = &OrderNotificationRefund{
			Amount:        decoded.Refund.Amount,
			Currency:      decoded.Refund.Currency,
			Reason:        decoded.Refund.Reason,
			Code:          decoded.Refund.Code,
			ReceiptNumber: decoded.Refund.ReceiptNumber,
			ReceiptUrl:    decoded.Refund.ReceiptUrl,
		}
	}
	m.Metadata = decoded.Metadata
	m.PrivateMetadata = decoded.PrivateMetadata
	m.Project = getOrderProject(decoded.Project)

	for _, v := range decoded.Items {
		item := &OrderItem{
			Id:          v.Id.Hex(),
			Object:      v.Object,
			Sku:         v.Sku,
			Name:        v.Name,
			Description: v.Description,
			Amount:      v.Amount,
			Currency:    v.Currency,
			Images:      v.Images,
			Url:         v.Url,
			Metadata:    v.Metadata,
			Code:        v.Code,
			PlatformId:  v.PlatformId,
		}
		item.CreatedAt, _ = ptypes.TimestampProto(v.CreatedAt)
		item.UpdatedAt, _ = ptypes.TimestampProto(v.UpdatedAt)
		m.Items = append(m.Items, item)
	}

	if decoded.Project != nil {
		nameLen := len(decoded.Project.Name)
		if nameLen > 0 {
			m.Project.Name = make(map[string]string, nameLen)

			for _, v := range decoded.Project.Name {
				m.Project.Name[v.Lang] = v.Value
			}
		}
	}

	m.ProjectOrderId = decoded.ProjectOrderId
	m.ProjectAccount = decoded.ProjectAccount
	m.ProjectParams = decoded.ProjectParams
	m.IsJsonRequest = decoded.IsJsonRequest
	m.OrderAmount = decoded.OrderAmount
	m.PaymentMethodPayerAccount = decoded.PaymentMethodPayerAccount
	m.PaymentMethodTxnParams = decoded.PaymentMethodTxnParams
	m.PaymentRequisites = decoded.PaymentRequisites
	m.UserAddressDataRequired = decoded.UserAddressDataRequired
	m.Products = decoded.Products
	m.IsNotificationsSent = decoded.IsNotificationsSent
	m.CountryRestriction = decoded.CountryRestriction
	m.ParentOrder = decoded.ParentOrder
	m.Type = decoded.Type
	m.IsVatDeduction = decoded.IsVatDeduction
	m.CountryCode = decoded.CountryCode
	m.VatPayer = decoded.VatPayer
	m.IsProduction = decoded.IsProduction

	m.PaymentMethodOrderClosedAt, err = ptypes.TimestampProto(decoded.PaymentMethodOrderClosedAt)
	if err != nil {
		return err
	}

	m.ParentPaymentAt, err = ptypes.TimestampProto(decoded.ParentPaymentAt)

	if err != nil {
		return err
	}

	m.ProjectLastRequestedAt, err = ptypes.TimestampProto(decoded.ProjectLastRequestedAt)
	if err != nil {
		return err
	}

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)
	if err != nil {
		return err
	}
	m.CanceledAt, err = ptypes.TimestampProto(decoded.CanceledAt)
	if err != nil {
		return err
	}

	m.RefundedAt, err = ptypes.TimestampProto(decoded.RefundedAt)
	if err != nil {
		return err
	}

	m.ExpireDateToFormInput, err = ptypes.TimestampProto(decoded.ExpireDateToFormInput)
	if err != nil {
		return err
	}

	return nil
}

func (m *PaymentFormPaymentMethod) IsBankCard() bool {
	return m.Group == recurringpb.PaymentSystemGroupAliasBankCard
}

func (m *PaymentMethod) IsBankCard() bool {
	return m.Group == recurringpb.PaymentSystemGroupAliasBankCard
}

func (m *PaymentMethodOrder) IsBankCard() bool {
	return m.Group == recurringpb.PaymentSystemGroupAliasBankCard
}

func (m *PaymentMethodOrder) IsCryptoCurrency() bool {
	return m.Group == recurringpb.PaymentSystemGroupAliasBitcoin
}

func getPaymentMethodOrder(in *MgoOrderPaymentMethod) *PaymentMethodOrder {
	if in == nil {
		return nil
	}

	result := &PaymentMethodOrder{
		Id:              in.Id.Hex(),
		Name:            in.Name,
		ExternalId:      in.ExternalId,
		Params:          in.Params,
		PaymentSystemId: in.PaymentSystemId.Hex(),
		Group:           in.Group,
		Saved:           in.Saved,
		RefundAllowed:   in.RefundAllowed,
		Handler:         in.Handler,
	}

	if in.Card != nil {
		result.Card = in.Card
	}
	if in.Wallet != nil {
		result.Wallet = in.Wallet
	}
	if in.CryptoCurrency != nil {
		result.CryptoCurrency = in.CryptoCurrency
	}

	return result
}

func getOrderProject(in *MgoOrderProject) *ProjectOrder {
	project := &ProjectOrder{
		Id:                      in.Id.Hex(),
		MerchantId:              in.MerchantId.Hex(),
		UrlSuccess:              in.UrlSuccess,
		UrlFail:                 in.UrlFail,
		NotifyEmails:            in.NotifyEmails,
		SendNotifyEmail:         in.SendNotifyEmail,
		SecretKey:               in.SecretKey,
		UrlCheckAccount:         in.UrlCheckAccount,
		UrlProcessPayment:       in.UrlProcessPayment,
		UrlChargebackPayment:    in.UrlChargebackPayment,
		UrlCancelPayment:        in.UrlCancelPayment,
		UrlRefundPayment:        in.UrlRefundPayment,
		UrlFraudPayment:         in.UrlFraudPayment,
		CallbackProtocol:        in.CallbackProtocol,
		Status:                  in.Status,
		MerchantRoyaltyCurrency: in.MerchantRoyaltyCurrency,
		RedirectSettings:        in.RedirectSettings,
	}

	if len(in.Name) > 0 {
		project.Name = make(map[string]string)

		for _, v := range in.Name {
			project.Name[v.Lang] = v.Value
		}
	}

	return project
}

func getOrderViewMoney(in *OrderViewMoney) *OrderViewMoney {
	if in == nil {
		return &OrderViewMoney{}
	}

	return &OrderViewMoney{
		Amount:   tools.ToPrecise(in.Amount),
		Currency: in.Currency,
	}
}

func getOrderViewItems(in []*MgoOrderItem) []*OrderItem {
	var items []*OrderItem

	if len(in) <= 0 {
		return items
	}

	for _, v := range in {
		item := &OrderItem{
			Id:          v.Id.Hex(),
			Object:      v.Object,
			Sku:         v.Sku,
			Name:        v.Name,
			Description: v.Description,
			Amount:      v.Amount,
			Currency:    v.Currency,
			Images:      v.Images,
			Url:         v.Url,
			Metadata:    v.Metadata,
			Code:        v.Code,
			PlatformId:  v.PlatformId,
		}

		item.CreatedAt, _ = ptypes.TimestampProto(v.CreatedAt)
		item.CreatedAt, _ = ptypes.TimestampProto(v.UpdatedAt)

		items = append(items, item)
	}

	return items
}

func (m *Id) MarshalBSON() ([]byte, error) {
	st := &MgoId{}
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, errors.New(ErrorInvalidObjectId)
	}
	st.Id = oid
	return bson.Marshal(st)
}

func (m *Id) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoId)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	return nil
}

type CountryAndRegionItem struct {
	Country string `bson:"iso_code_a2"`
	Region  string `bson:"region"`
}

type CountryAndRegionItems struct {
	Items []*CountryAndRegionItem `json:"items"`
}
