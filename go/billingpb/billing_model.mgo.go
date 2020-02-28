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

type MgoCommission struct {
	Id struct {
		PaymentMethodId primitive.ObjectID `bson:"pm_id"`
		ProjectId       primitive.ObjectID `bson:"project_id"`
	} `bson:"_id"`
	PaymentMethodCommission float64   `bson:"pm_commission"`
	PspCommission           float64   `bson:"psp_commission"`
	ToUserCommission        float64   `bson:"total_commission_to_user"`
	StartDate               time.Time `bson:"start_date"`
}

type MgoCommissionBilling struct {
	Id                      primitive.ObjectID `bson:"_id"`
	PaymentMethodId         primitive.ObjectID `bson:"pm_id"`
	ProjectId               primitive.ObjectID `bson:"project_id"`
	PaymentMethodCommission float64            `bson:"pm_commission"`
	PspCommission           float64            `bson:"psp_commission"`
	TotalCommissionToUser   float64            `bson:"total_commission_to_user"`
	StartDate               time.Time          `bson:"start_date"`
	CreatedAt               time.Time          `bson:"created_at"`
	UpdatedAt               time.Time          `bson:"updated_at"`
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

type MgoRefundOrder struct {
	Id   primitive.ObjectID `bson:"id"`
	Uuid string             `bson:"uuid"`
}

type MgoRefund struct {
	Id             primitive.ObjectID `bson:"_id"`
	OriginalOrder  *MgoRefundOrder    `bson:"original_order"`
	ExternalId     string             `bson:"external_id"`
	Amount         float64            `bson:"amount"`
	CreatorId      primitive.ObjectID `bson:"creator_id"`
	Currency       string             `bson:"currency"`
	Status         int32              `bson:"status"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
	PayerData      *RefundPayerData   `bson:"payer_data"`
	SalesTax       float32            `bson:"sales_tax"`
	IsChargeback   bool               `bson:"is_chargeback"`
	CreatedOrderId primitive.ObjectID `bson:"created_order_id,omitempty"`
	Reason         string             `bson:"reason"`
}

type MgoCountry struct {
	Id                      primitive.ObjectID   `bson:"_id"`
	IsoCodeA2               string               `bson:"iso_code_a2"`
	Region                  string               `bson:"region"`
	Currency                string               `bson:"currency"`
	PaymentsAllowed         bool                 `bson:"payments_allowed"`
	ChangeAllowed           bool                 `bson:"change_allowed"`
	VatEnabled              bool                 `bson:"vat_enabled"`
	VatCurrency             string               `bson:"vat_currency"`
	PriceGroupId            string               `bson:"price_group_id"`
	VatThreshold            *CountryVatThreshold `bson:"vat_threshold"`
	VatPeriodMonth          int32                `bson:"vat_period_month"`
	VatDeadlineDays         int32                `bson:"vat_deadline_days"`
	VatStoreYears           int32                `bson:"vat_store_years"`
	VatCurrencyRatesPolicy  string               `bson:"vat_currency_rates_policy"`
	VatCurrencyRatesSource  string               `bson:"vat_currency_rates_source"`
	PayerTariffRegion       string               `bson:"payer_tariff_region"`
	CreatedAt               time.Time            `bson:"created_at"`
	UpdatedAt               time.Time            `bson:"updated_at"`
	HighRiskPaymentsAllowed bool                 `bson:"high_risk_payments_allowed"`
	HighRiskChangeAllowed   bool                 `bson:"high_risk_change_allowed"`
}

type MgoPayoutCostSystem struct {
	Id                    primitive.ObjectID `bson:"_id"`
	IntrabankCostAmount   float64            `bson:"intrabank_cost_amount"`
	IntrabankCostCurrency string             `bson:"intrabank_cost_currency"`
	InterbankCostAmount   float64            `bson:"interbank_cost_amount"`
	InterbankCostCurrency string             `bson:"interbank_cost_currency"`
	IsActive              bool               `bson:"is_active"`
	CreatedAt             time.Time          `bson:"created_at"`
}

type MgoRoyaltyReport struct {
	Id                 primitive.ObjectID    `bson:"_id"`
	MerchantId         primitive.ObjectID    `bson:"merchant_id"`
	CreatedAt          time.Time             `bson:"created_at"`
	UpdatedAt          time.Time             `bson:"updated_at"`
	PayoutDate         time.Time             `bson:"payout_date"`
	Status             string                `bson:"status"`
	PeriodFrom         time.Time             `bson:"period_from"`
	PeriodTo           time.Time             `bson:"period_to"`
	AcceptExpireAt     time.Time             `bson:"accept_expire_at"`
	AcceptedAt         time.Time             `bson:"accepted_at"`
	Totals             *RoyaltyReportTotals  `bson:"totals"`
	Currency           string                `bson:"currency"`
	Summary            *RoyaltyReportSummary `bson:"summary"`
	DisputeReason      string                `bson:"dispute_reason"`
	DisputeStartedAt   time.Time             `bson:"dispute_started_at"`
	DisputeClosedAt    time.Time             `bson:"dispute_closed_at"`
	IsAutoAccepted     bool                  `bson:"is_auto_accepted"`
	PayoutDocumentId   string                `bson:"payout_document_id"`
	OperatingCompanyId string                `bson:"operating_company_id"`
}

type MgoRoyaltyReportCorrectionItem struct {
	AccountingEntryId primitive.ObjectID `bson:"accounting_entry_id"`
	Amount            float64            `bson:"amount"`
	Currency          string             `bson:"currency"`
	Reason            string             `bson:"reason"`
	EntryDate         time.Time          `bson:"entry_date"`
}

type MgoRoyaltyReportChanges struct {
	Id              primitive.ObjectID `bson:"_id"`
	RoyaltyReportId primitive.ObjectID `bson:"royalty_report_id"`
	Source          string             `bson:"source"`
	Ip              string             `bson:"ip"`
	Hash            string             `bson:"hash"`
	CreatedAt       time.Time          `bson:"created_at"`
}

type MgoOrderViewPrivate struct {
	Id                                         primitive.ObjectID             `bson:"_id" json:"-"`
	Uuid                                       string                         `bson:"uuid" json:"uuid"`
	TotalPaymentAmount                         float64                        `bson:"total_payment_amount" json:"total_payment_amount"`
	Currency                                   string                         `bson:"currency" json:"currency"`
	Project                                    *MgoOrderProject               `bson:"project" json:"project"`
	CreatedAt                                  time.Time                      `bson:"created_at" json:"created_at"`
	Transaction                                string                         `bson:"pm_order_id" json:"transaction"`
	PaymentMethod                              *MgoOrderPaymentMethod         `bson:"payment_method" json:"payment_method"`
	CountryCode                                string                         `bson:"country_code" json:"country_code"`
	MerchantId                                 primitive.ObjectID             `bson:"merchant_id" json:"merchant_id"`
	Locale                                     string                         `bson:"locale" json:"locale"`
	Status                                     string                         `bson:"status" json:"status"`
	TransactionDate                            time.Time                      `bson:"pm_order_close_date" json:"transaction_date"`
	User                                       *OrderUser                     `bson:"user" json:"user"`
	BillingAddress                             *OrderBillingAddress           `bson:"billing_address" json:"billing_address"`
	Type                                       string                         `bson:"type" json:"type"`
	IsVatDeduction                             bool                           `bson:"is_vat_deduction" json:"is_vat_deduction"`
	PaymentGrossRevenueLocal                   *OrderViewMoney                `bson:"payment_gross_revenue_local" json:"payment_gross_revenue_local"`
	PaymentGrossRevenueOrigin                  *OrderViewMoney                `bson:"payment_gross_revenue_origin" json:"payment_gross_revenue_origin"`
	PaymentGrossRevenue                        *OrderViewMoney                `bson:"payment_gross_revenue" json:"payment_gross_revenue"`
	PaymentTaxFee                              *OrderViewMoney                `bson:"payment_tax_fee" json:"payment_tax_fee"`
	PaymentTaxFeeLocal                         *OrderViewMoney                `bson:"payment_tax_fee_local" json:"payment_tax_fee_local"`
	PaymentTaxFeeOrigin                        *OrderViewMoney                `bson:"payment_tax_fee_origin" json:"payment_tax_fee_origin"`
	PaymentTaxFeeCurrencyExchangeFee           *OrderViewMoney                `bson:"payment_tax_fee_currency_exchange_fee" json:"payment_tax_fee_currency_exchange_fee"`
	PaymentTaxFeeTotal                         *OrderViewMoney                `bson:"payment_tax_fee_total" json:"payment_tax_fee_total"`
	PaymentGrossRevenueFx                      *OrderViewMoney                `bson:"payment_gross_revenue_fx" json:"payment_gross_revenue_fx"`
	PaymentGrossRevenueFxTaxFee                *OrderViewMoney                `bson:"payment_gross_revenue_fx_tax_fee" json:"payment_gross_revenue_fx_tax_fee"`
	PaymentGrossRevenueFxProfit                *OrderViewMoney                `bson:"payment_gross_revenue_fx_profit" json:"payment_gross_revenue_fx_profit"`
	GrossRevenue                               *OrderViewMoney                `bson:"gross_revenue" json:"gross_revenue"`
	TaxFee                                     *OrderViewMoney                `bson:"tax_fee" json:"tax_fee"`
	TaxFeeCurrencyExchangeFee                  *OrderViewMoney                `bson:"tax_fee_currency_exchange_fee" json:"tax_fee_currency_exchange_fee"`
	TaxFeeTotal                                *OrderViewMoney                `bson:"tax_fee_total" json:"tax_fee_total"`
	MethodFeeTotal                             *OrderViewMoney                `bson:"method_fee_total" json:"method_fee_total"`
	MethodFeeTariff                            *OrderViewMoney                `bson:"method_fee_tariff" json:"method_fee_tariff"`
	PaysuperMethodFeeTariffSelfCost            *OrderViewMoney                `bson:"paysuper_method_fee_tariff_self_cost" json:"paysuper_method_fee_tariff_self_cost"`
	PaysuperMethodFeeProfit                    *OrderViewMoney                `bson:"paysuper_method_fee_profit" json:"paysuper_method_fee_profit"`
	MethodFixedFeeTariff                       *OrderViewMoney                `bson:"method_fixed_fee_tariff" json:"method_fixed_fee_tariff"`
	PaysuperMethodFixedFeeTariffFxProfit       *OrderViewMoney                `bson:"paysuper_method_fixed_fee_tariff_fx_profit" json:"paysuper_method_fixed_fee_tariff_fx_profit"`
	PaysuperMethodFixedFeeTariffSelfCost       *OrderViewMoney                `bson:"paysuper_method_fixed_fee_tariff_self_cost" json:"paysuper_method_fixed_fee_tariff_self_cost"`
	PaysuperMethodFixedFeeTariffTotalProfit    *OrderViewMoney                `bson:"paysuper_method_fixed_fee_tariff_total_profit" json:"paysuper_method_fixed_fee_tariff_total_profit"`
	PaysuperFixedFee                           *OrderViewMoney                `bson:"paysuper_fixed_fee" json:"paysuper_fixed_fee"`
	PaysuperFixedFeeFxProfit                   *OrderViewMoney                `bson:"paysuper_fixed_fee_fx_profit" json:"paysuper_fixed_fee_fx_profit"`
	FeesTotal                                  *OrderViewMoney                `bson:"fees_total" json:"fees_total"`
	FeesTotalLocal                             *OrderViewMoney                `bson:"fees_total_local" json:"fees_total_local"`
	NetRevenue                                 *OrderViewMoney                `bson:"net_revenue" json:"net_revenue"`
	PaysuperMethodTotalProfit                  *OrderViewMoney                `bson:"paysuper_method_total_profit" json:"paysuper_method_total_profit"`
	PaysuperTotalProfit                        *OrderViewMoney                `bson:"paysuper_total_profit" json:"paysuper_total_profit"`
	PaymentRefundGrossRevenueLocal             *OrderViewMoney                `bson:"payment_refund_gross_revenue_local" json:"payment_refund_gross_revenue_local"`
	PaymentRefundGrossRevenueOrigin            *OrderViewMoney                `bson:"payment_refund_gross_revenue_origin" json:"payment_refund_gross_revenue_origin"`
	PaymentRefundGrossRevenue                  *OrderViewMoney                `bson:"payment_refund_gross_revenue" json:"payment_refund_gross_revenue"`
	PaymentRefundTaxFee                        *OrderViewMoney                `bson:"payment_refund_tax_fee" json:"payment_refund_tax_fee"`
	PaymentRefundTaxFeeLocal                   *OrderViewMoney                `bson:"payment_refund_tax_fee_local" json:"payment_refund_tax_fee_local"`
	PaymentRefundTaxFeeOrigin                  *OrderViewMoney                `bson:"payment_refund_tax_fee_origin" json:"payment_refund_tax_fee_origin"`
	PaymentRefundFeeTariff                     *OrderViewMoney                `bson:"payment_refund_fee_tariff" json:"payment_refund_fee_tariff"`
	MethodRefundFixedFeeTariff                 *OrderViewMoney                `bson:"method_refund_fixed_fee_tariff" json:"method_refund_fixed_fee_tariff"`
	RefundGrossRevenue                         *OrderViewMoney                `bson:"refund_gross_revenue" json:"refund_gross_revenue"`
	RefundGrossRevenueFx                       *OrderViewMoney                `bson:"refund_gross_revenue_fx" json:"refund_gross_revenue_fx"`
	MethodRefundFeeTariff                      *OrderViewMoney                `bson:"method_refund_fee_tariff" json:"method_refund_fee_tariff"`
	PaysuperMethodRefundFeeTariffProfit        *OrderViewMoney                `bson:"paysuper_method_refund_fee_tariff_profit" json:"paysuper_method_refund_fee_tariff_profit"`
	PaysuperMethodRefundFixedFeeTariffSelfCost *OrderViewMoney                `bson:"paysuper_method_refund_fixed_fee_tariff_self_cost" json:"paysuper_method_refund_fixed_fee_tariff_self_cost"`
	MerchantRefundFixedFeeTariff               *OrderViewMoney                `bson:"merchant_refund_fixed_fee_tariff" json:"merchant_refund_fixed_fee_tariff"`
	PaysuperMethodRefundFixedFeeTariffProfit   *OrderViewMoney                `bson:"paysuper_method_refund_fixed_fee_tariff_profit" json:"paysuper_method_refund_fixed_fee_tariff_profit"`
	RefundTaxFee                               *OrderViewMoney                `bson:"refund_tax_fee" json:"refund_tax_fee"`
	RefundTaxFeeCurrencyExchangeFee            *OrderViewMoney                `bson:"refund_tax_fee_currency_exchange_fee" json:"refund_tax_fee_currency_exchange_fee"`
	PaysuperRefundTaxFeeCurrencyExchangeFee    *OrderViewMoney                `bson:"paysuper_refund_tax_fee_currency_exchange_fee" json:"paysuper_refund_tax_fee_currency_exchange_fee"`
	RefundTaxFeeTotal                          *OrderViewMoney                `bson:"refund_tax_fee_total" json:"refund_tax_fee_total"`
	RefundReverseRevenue                       *OrderViewMoney                `bson:"refund_reverse_revenue" json:"refund_reverse_revenue"`
	RefundFeesTotal                            *OrderViewMoney                `bson:"refund_fees_total" json:"refund_fees_total"`
	RefundFeesTotalLocal                       *OrderViewMoney                `bson:"refund_fees_total_local" json:"refund_fees_total_local"`
	PaysuperRefundTotalProfit                  *OrderViewMoney                `bson:"paysuper_refund_total_profit" json:"paysuper_refund_total_profit"`
	Issuer                                     *OrderIssuer                   `bson:"issuer"`
	Items                                      []*MgoOrderItem                `bson:"items"`
	MerchantPayoutCurrency                     string                         `bson:"merchant_payout_currency"`
	ParentOrder                                *ParentOrder                   `bson:"parent_order"`
	Refund                                     *MgoOrderNotificationRefund    `bson:"refund"`
	Cancellation                               *OrderNotificationCancellation `bson:"cancellation"`
	MccCode                                    string                         `bson:"mcc_code"`
	OperatingCompanyId                         string                         `bson:"operating_company_id"`
	IsHighRisk                                 bool                           `bson:"is_high_risk"`
	RefundAllowed                              bool                           `bson:"refund_allowed"`
	VatPayer                                   string                         `bson:"vat_payer"`
	IsProduction                               bool                           `bson:"is_production"`
	TaxRate                                    float64                        `bson:"tax_rate"`
	MerchantInfo                               *OrderViewMerchantInfo         `bson:"merchant_info"`
	OrderChargeBeforeVat                       *OrderViewMoney                `bson:"order_charge_before_vat"`
}

type MgoOrderViewPublic struct {
	Id                                      primitive.ObjectID             `bson:"_id"`
	Uuid                                    string                         `bson:"uuid"`
	TotalPaymentAmount                      float64                        `bson:"total_payment_amount"`
	Currency                                string                         `bson:"currency"`
	Project                                 *MgoOrderProject               `bson:"project"`
	CreatedAt                               time.Time                      `bson:"created_at"`
	Transaction                             string                         `bson:"pm_order_id"`
	PaymentMethod                           *MgoOrderPaymentMethod         `bson:"payment_method"`
	CountryCode                             string                         `bson:"country_code"`
	MerchantId                              primitive.ObjectID             `bson:"merchant_id"`
	Locale                                  string                         `bson:"locale"`
	Status                                  string                         `bson:"status"`
	TransactionDate                         time.Time                      `bson:"pm_order_close_date"`
	User                                    *OrderUser                     `bson:"user"`
	BillingAddress                          *OrderBillingAddress           `bson:"billing_address"`
	Type                                    string                         `bson:"type"`
	IsVatDeduction                          bool                           `bson:"is_vat_deduction"`
	GrossRevenue                            *OrderViewMoney                `bson:"gross_revenue"`
	TaxFee                                  *OrderViewMoney                `bson:"tax_fee"`
	TaxFeeCurrencyExchangeFee               *OrderViewMoney                `bson:"tax_fee_currency_exchange_fee"`
	TaxFeeTotal                             *OrderViewMoney                `bson:"tax_fee_total"`
	MethodFeeTotal                          *OrderViewMoney                `bson:"method_fee_total"`
	MethodFeeTariff                         *OrderViewMoney                `bson:"method_fee_tariff"`
	MethodFixedFeeTariff                    *OrderViewMoney                `bson:"method_fixed_fee_tariff"`
	PaysuperFixedFee                        *OrderViewMoney                `bson:"paysuper_fixed_fee"`
	FeesTotal                               *OrderViewMoney                `bson:"fees_total"`
	FeesTotalLocal                          *OrderViewMoney                `bson:"fees_total_local"`
	NetRevenue                              *OrderViewMoney                `bson:"net_revenue"`
	RefundGrossRevenue                      *OrderViewMoney                `bson:"refund_gross_revenue"`
	MethodRefundFeeTariff                   *OrderViewMoney                `bson:"method_refund_fee_tariff"`
	MerchantRefundFixedFeeTariff            *OrderViewMoney                `bson:"merchant_refund_fixed_fee_tariff"`
	RefundTaxFee                            *OrderViewMoney                `bson:"refund_tax_fee"`
	RefundTaxFeeCurrencyExchangeFee         *OrderViewMoney                `bson:"refund_tax_fee_currency_exchange_fee"`
	PaysuperRefundTaxFeeCurrencyExchangeFee *OrderViewMoney                `bson:"paysuper_refund_tax_fee_currency_exchange_fee"`
	RefundReverseRevenue                    *OrderViewMoney                `bson:"refund_reverse_revenue"`
	RefundFeesTotal                         *OrderViewMoney                `bson:"refund_fees_total"`
	RefundFeesTotalLocal                    *OrderViewMoney                `bson:"refund_fees_total_local"`
	Issuer                                  *OrderIssuer                   `bson:"issuer"`
	Items                                   []*MgoOrderItem                `bson:"items"`
	MerchantPayoutCurrency                  string                         `bson:"merchant_payout_currency"`
	ParentOrder                             *ParentOrder                   `bson:"parent_order"`
	Refund                                  *MgoOrderNotificationRefund    `bson:"refund"`
	Cancellation                            *OrderNotificationCancellation `bson:"cancellation"`
	OperatingCompanyId                      string                         `bson:"operating_company_id"`
	RefundAllowed                           bool                           `bson:"refund_allowed"`
	OrderCharge                             *OrderViewMoney                `bson:"order_charge"`
	OrderChargeBeforeVat                    *OrderViewMoney                `bson:"order_charge_before_vat"`
	PaymentIpCountry                        string                         `bson:"payment_ip_country"`
	IsIpCountryMismatchBin                  bool                           `bson:"is_ip_country_mismatch_bin"`
	BillingCountryChangedByUser             bool                           `bson:"billing_country_changed_by_user"`
	VatPayer                                string                         `bson:"vat_payer"`
	IsProduction                            bool                           `bson:"is_production"`
}

type MgoMerchantBalance struct {
	Id             primitive.ObjectID `bson:"_id"`
	MerchantId     primitive.ObjectID `bson:"merchant_id"`
	Currency       string             `bson:"currency"`
	Debit          float64            `bson:"debit"`
	Credit         float64            `bson:"credit"`
	RollingReserve float64            `bson:"rolling_reserve"`
	Total          float64            `bson:"total"`
	CreatedAt      time.Time          `bson:"created_at"`
}

func (m *MerchantBalance) MarshalBSON() ([]byte, error) {
	st := &MgoMerchantBalance{
		Currency:       m.Currency,
		Debit:          m.Debit,
		Credit:         m.Credit,
		RollingReserve: m.RollingReserve,
		Total:          m.Total,
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

	merchantOid, err := primitive.ObjectIDFromHex(m.MerchantId)

	if err != nil {
		return nil, errors.New(ErrorInvalidObjectId)
	}

	st.MerchantId = merchantOid

	if m.CreatedAt != nil {
		t, err := ptypes.Timestamp(m.CreatedAt)

		if err != nil {
			return nil, err
		}

		st.CreatedAt = t
	} else {
		st.CreatedAt = time.Now()
	}

	return bson.Marshal(st)
}

func (m *MerchantBalance) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoMerchantBalance)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.MerchantId = decoded.MerchantId.Hex()
	m.Currency = decoded.Currency
	m.Debit = decoded.Debit
	m.Credit = decoded.Credit
	m.RollingReserve = decoded.RollingReserve
	m.Total = decoded.Total

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *Country) MarshalBSON() ([]byte, error) {
	st := &MgoCountry{
		IsoCodeA2:               m.IsoCodeA2,
		Region:                  m.Region,
		Currency:                m.Currency,
		PaymentsAllowed:         m.PaymentsAllowed,
		ChangeAllowed:           m.ChangeAllowed,
		VatEnabled:              m.VatEnabled,
		PriceGroupId:            m.PriceGroupId,
		VatCurrency:             m.VatCurrency,
		VatThreshold:            m.VatThreshold,
		VatPeriodMonth:          m.VatPeriodMonth,
		VatStoreYears:           m.VatStoreYears,
		VatCurrencyRatesPolicy:  m.VatCurrencyRatesPolicy,
		VatCurrencyRatesSource:  m.VatCurrencyRatesSource,
		PayerTariffRegion:       m.PayerTariffRegion,
		HighRiskPaymentsAllowed: m.HighRiskPaymentsAllowed,
		HighRiskChangeAllowed:   m.HighRiskChangeAllowed,
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

	return bson.Marshal(st)
}

func (m *Country) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoCountry)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.IsoCodeA2 = decoded.IsoCodeA2
	m.Region = decoded.Region
	m.Currency = decoded.Currency
	m.PaymentsAllowed = decoded.PaymentsAllowed
	m.ChangeAllowed = decoded.ChangeAllowed
	m.VatEnabled = decoded.VatEnabled
	m.PriceGroupId = decoded.PriceGroupId
	m.VatCurrency = decoded.VatCurrency
	m.VatThreshold = decoded.VatThreshold
	m.VatPeriodMonth = decoded.VatPeriodMonth
	m.VatDeadlineDays = decoded.VatDeadlineDays
	m.VatStoreYears = decoded.VatStoreYears
	m.VatCurrencyRatesPolicy = decoded.VatCurrencyRatesPolicy
	m.VatCurrencyRatesSource = decoded.VatCurrencyRatesSource
	m.PayerTariffRegion = decoded.PayerTariffRegion
	m.HighRiskPaymentsAllowed = decoded.HighRiskPaymentsAllowed
	m.HighRiskChangeAllowed = decoded.HighRiskChangeAllowed

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (m *Commission) MarshalBSON() ([]byte, error) {
	paymentMethodOid, err := primitive.ObjectIDFromHex(m.PaymentMethodId)

	if err != nil {
		return nil, err
	}

	projectOid, err := primitive.ObjectIDFromHex(m.ProjectId)

	if err != nil {
		return nil, err
	}

	st := &MgoCommissionBilling{
		PaymentMethodId:         paymentMethodOid,
		ProjectId:               projectOid,
		PaymentMethodCommission: m.PaymentMethodCommission,
		PspCommission:           m.PspCommission,
		TotalCommissionToUser:   m.TotalCommissionToUser,
	}

	t, err := ptypes.Timestamp(m.StartDate)

	if err != nil {
		return nil, err
	}

	st.StartDate = t

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

	return bson.Marshal(st)
}

func (m *Commission) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoCommissionBilling)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.PaymentMethodId = decoded.PaymentMethodId.Hex()
	m.ProjectId = decoded.ProjectId.Hex()
	m.PaymentMethodCommission = decoded.PaymentMethodCommission
	m.PspCommission = decoded.PspCommission
	m.TotalCommissionToUser = decoded.TotalCommissionToUser

	m.StartDate, err = ptypes.TimestampProto(decoded.StartDate)

	if err != nil {
		return err
	}

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

	return err
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

func (m *Refund) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.OriginalOrder.Id)

	if err != nil {
		return nil, err
	}

	creatorOid, err := primitive.ObjectIDFromHex(m.CreatorId)

	if err != nil {
		return nil, err
	}

	st := &MgoRefund{
		OriginalOrder: &MgoRefundOrder{
			Id:   oid,
			Uuid: m.OriginalOrder.Uuid,
		},
		ExternalId:   m.ExternalId,
		Amount:       m.Amount,
		CreatorId:    creatorOid,
		Currency:     m.Currency,
		Status:       m.Status,
		PayerData:    m.PayerData,
		SalesTax:     m.SalesTax,
		IsChargeback: m.IsChargeback,
		Reason:       m.Reason,
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

	if m.CreatedOrderId != "" {
		oid, err := primitive.ObjectIDFromHex(m.CreatedOrderId)

		if err != nil {
			return nil, errors.New(ErrorInvalidObjectId)
		}

		st.CreatedOrderId = oid
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

	return bson.Marshal(st)
}

func (m *Refund) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoRefund)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.OriginalOrder = &RefundOrder{
		Id:   decoded.OriginalOrder.Id.Hex(),
		Uuid: decoded.OriginalOrder.Uuid,
	}
	m.ExternalId = decoded.ExternalId
	m.Amount = decoded.Amount
	m.CreatorId = decoded.CreatorId.Hex()
	m.Currency = decoded.Currency
	m.Status = decoded.Status
	m.PayerData = decoded.PayerData
	m.SalesTax = decoded.SalesTax
	m.IsChargeback = decoded.IsChargeback
	m.Reason = decoded.Reason

	if !decoded.CreatedOrderId.IsZero() {
		m.CreatedOrderId = decoded.CreatedOrderId.Hex()
	}

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	m.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

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

func (m *PayoutCostSystem) MarshalBSON() ([]byte, error) {
	st := &MgoPayoutCostSystem{
		IntrabankCostAmount:   m.IntrabankCostAmount,
		IntrabankCostCurrency: m.IntrabankCostCurrency,
		InterbankCostAmount:   m.InterbankCostAmount,
		InterbankCostCurrency: m.InterbankCostCurrency,
		IsActive:              m.IsActive,
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

	t, err := ptypes.Timestamp(m.CreatedAt)

	if err != nil {
		return nil, err
	}

	st.CreatedAt = t

	return bson.Marshal(st)
}

func (m *PayoutCostSystem) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoPayoutCostSystem)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}
	m.Id = decoded.Id.Hex()
	m.IntrabankCostAmount = decoded.IntrabankCostAmount
	m.IntrabankCostCurrency = decoded.IntrabankCostCurrency
	m.InterbankCostAmount = decoded.InterbankCostAmount
	m.InterbankCostCurrency = decoded.InterbankCostCurrency
	m.IsActive = decoded.IsActive

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *RoyaltyReport) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, err
	}

	merchantOid, err := primitive.ObjectIDFromHex(m.MerchantId)

	if err != nil {
		return nil, err
	}

	st := &MgoRoyaltyReport{
		Id:                 oid,
		MerchantId:         merchantOid,
		Status:             m.Status,
		Totals:             m.Totals,
		Currency:           m.Currency,
		Summary:            m.Summary,
		DisputeReason:      m.DisputeReason,
		IsAutoAccepted:     m.IsAutoAccepted,
		PayoutDocumentId:   m.PayoutDocumentId,
		OperatingCompanyId: m.OperatingCompanyId,
	}

	if m.PayoutDate != nil {
		t, err := ptypes.Timestamp(m.PayoutDate)

		if err != nil {
			return nil, err
		}

		st.PayoutDate = t
	}

	t, err := ptypes.Timestamp(m.PeriodFrom)

	if err != nil {
		return nil, err
	}

	st.PeriodFrom = t
	t, err = ptypes.Timestamp(m.PeriodTo)

	if err != nil {
		return nil, err
	}

	st.PeriodTo = t
	t, err = ptypes.Timestamp(m.AcceptExpireAt)

	if err != nil {
		return nil, err
	}

	st.AcceptExpireAt = t

	if m.AcceptedAt != nil {
		t, err = ptypes.Timestamp(m.AcceptedAt)

		if err != nil {
			return nil, err
		}

		st.AcceptedAt = t
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

	if m.DisputeStartedAt != nil {
		t, err := ptypes.Timestamp(m.DisputeStartedAt)

		if err != nil {
			return nil, err
		}

		st.DisputeStartedAt = t
	}

	if m.DisputeClosedAt != nil {
		t, err := ptypes.Timestamp(m.DisputeClosedAt)

		if err != nil {
			return nil, err
		}

		st.DisputeClosedAt = t
	}

	return bson.Marshal(st)
}

func (m *RoyaltyReport) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoRoyaltyReport)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.MerchantId = decoded.MerchantId.Hex()
	m.Status = decoded.Status
	m.Totals = decoded.Totals
	m.Currency = decoded.Currency
	m.Summary = decoded.Summary
	m.DisputeReason = decoded.DisputeReason
	m.IsAutoAccepted = decoded.IsAutoAccepted
	m.PayoutDocumentId = decoded.PayoutDocumentId
	m.OperatingCompanyId = decoded.OperatingCompanyId

	m.PayoutDate, err = ptypes.TimestampProto(decoded.PayoutDate)
	if err != nil {
		return err
	}

	m.PeriodFrom, err = ptypes.TimestampProto(decoded.PeriodFrom)
	if err != nil {
		return err
	}

	m.PeriodTo, err = ptypes.TimestampProto(decoded.PeriodTo)
	if err != nil {
		return err
	}

	m.AcceptExpireAt, err = ptypes.TimestampProto(decoded.AcceptExpireAt)
	if err != nil {
		return err
	}

	m.AcceptedAt, err = ptypes.TimestampProto(decoded.AcceptedAt)
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

	m.DisputeStartedAt, err = ptypes.TimestampProto(decoded.DisputeStartedAt)
	if err != nil {
		return err
	}

	m.DisputeClosedAt, err = ptypes.TimestampProto(decoded.DisputeClosedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m *RoyaltyReportChanges) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, err
	}

	royaltyReportOid, err := primitive.ObjectIDFromHex(m.RoyaltyReportId)

	if err != nil {
		return nil, err
	}

	st := &MgoRoyaltyReportChanges{
		Id:              oid,
		RoyaltyReportId: royaltyReportOid,
		Source:          m.Source,
		Ip:              m.Ip,
		Hash:            m.Hash,
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

	return bson.Marshal(st)
}

func (m *RoyaltyReportChanges) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoRoyaltyReportChanges)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.RoyaltyReportId = decoded.RoyaltyReportId.Hex()
	m.Source = decoded.Source
	m.Ip = decoded.Ip
	m.Hash = decoded.Hash

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (m *RoyaltyReportCorrectionItem) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.AccountingEntryId)

	if err != nil {
		return nil, err
	}

	st := &MgoRoyaltyReportCorrectionItem{
		AccountingEntryId: oid,
		Amount:            m.Amount,
		Currency:          m.Currency,
		Reason:            m.Reason,
	}

	t, err := ptypes.Timestamp(m.EntryDate)

	if err != nil {
		return nil, err
	}

	st.EntryDate = t

	return bson.Marshal(st)
}

func (m *RoyaltyReportCorrectionItem) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoRoyaltyReportCorrectionItem)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.AccountingEntryId = decoded.AccountingEntryId.Hex()
	m.Amount = decoded.Amount
	m.Currency = decoded.Currency
	m.Reason = decoded.Reason

	m.EntryDate, err = ptypes.TimestampProto(decoded.EntryDate)
	if err != nil {
		return err
	}

	return nil
}

func (m *OrderViewPrivate) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoOrderViewPrivate)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Uuid = decoded.Uuid
	m.TotalPaymentAmount = decoded.TotalPaymentAmount
	m.Currency = decoded.Currency
	m.Project = getOrderProject(decoded.Project)
	m.Transaction = decoded.Transaction
	m.PaymentMethod = getPaymentMethodOrder(decoded.PaymentMethod)
	m.CountryCode = decoded.CountryCode
	m.MerchantId = decoded.MerchantId.Hex()
	m.Locale = decoded.Locale
	m.Status = decoded.Status
	m.User = decoded.User
	m.BillingAddress = decoded.BillingAddress
	m.Type = decoded.Type
	m.Issuer = decoded.Issuer
	m.MerchantPayoutCurrency = decoded.MerchantPayoutCurrency
	m.IsVatDeduction = decoded.IsVatDeduction
	m.ParentOrder = decoded.ParentOrder
	m.Cancellation = decoded.Cancellation

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

	m.PaymentGrossRevenueLocal = getOrderViewMoney(decoded.PaymentGrossRevenueLocal)
	m.PaymentGrossRevenueOrigin = getOrderViewMoney(decoded.PaymentGrossRevenueOrigin)
	m.PaymentGrossRevenue = getOrderViewMoney(decoded.PaymentGrossRevenue)
	m.PaymentTaxFee = getOrderViewMoney(decoded.PaymentTaxFee)
	m.PaymentTaxFeeLocal = getOrderViewMoney(decoded.PaymentTaxFeeLocal)
	m.PaymentTaxFeeOrigin = getOrderViewMoney(decoded.PaymentTaxFeeOrigin)
	m.PaymentTaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.PaymentTaxFeeCurrencyExchangeFee)
	m.PaymentTaxFeeTotal = getOrderViewMoney(decoded.PaymentTaxFeeTotal)
	m.PaymentGrossRevenueFx = getOrderViewMoney(decoded.PaymentGrossRevenueFx)
	m.PaymentGrossRevenueFxTaxFee = getOrderViewMoney(decoded.PaymentGrossRevenueFxTaxFee)
	m.PaymentGrossRevenueFxProfit = getOrderViewMoney(decoded.PaymentGrossRevenueFxProfit)
	m.GrossRevenue = getOrderViewMoney(decoded.GrossRevenue)
	m.TaxFee = getOrderViewMoney(decoded.TaxFee)
	m.TaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.TaxFeeCurrencyExchangeFee)
	m.TaxFeeTotal = getOrderViewMoney(decoded.TaxFeeTotal)
	m.MethodFeeTotal = getOrderViewMoney(decoded.MethodFeeTotal)
	m.MethodFeeTariff = getOrderViewMoney(decoded.MethodFeeTariff)
	m.PaysuperMethodFeeTariffSelfCost = getOrderViewMoney(decoded.PaysuperMethodFeeTariffSelfCost)
	m.PaysuperMethodFeeProfit = getOrderViewMoney(decoded.PaysuperMethodFeeProfit)
	m.MethodFixedFeeTariff = getOrderViewMoney(decoded.MethodFixedFeeTariff)
	m.PaysuperMethodFixedFeeTariffFxProfit = getOrderViewMoney(decoded.PaysuperMethodFixedFeeTariffFxProfit)
	m.PaysuperMethodFixedFeeTariffSelfCost = getOrderViewMoney(decoded.PaysuperMethodFixedFeeTariffSelfCost)
	m.PaysuperMethodFixedFeeTariffTotalProfit = getOrderViewMoney(decoded.PaysuperMethodFixedFeeTariffTotalProfit)
	m.PaysuperFixedFee = getOrderViewMoney(decoded.PaysuperFixedFee)
	m.PaysuperFixedFeeFxProfit = getOrderViewMoney(decoded.PaysuperFixedFeeFxProfit)
	m.FeesTotal = getOrderViewMoney(decoded.FeesTotal)
	m.FeesTotalLocal = getOrderViewMoney(decoded.FeesTotalLocal)
	m.NetRevenue = getOrderViewMoney(decoded.NetRevenue)
	m.PaysuperMethodTotalProfit = getOrderViewMoney(decoded.PaysuperMethodTotalProfit)
	m.PaysuperTotalProfit = getOrderViewMoney(decoded.PaysuperTotalProfit)
	m.PaymentRefundGrossRevenueLocal = getOrderViewMoney(decoded.PaymentRefundGrossRevenueLocal)
	m.PaymentRefundGrossRevenueOrigin = getOrderViewMoney(decoded.PaymentRefundGrossRevenueOrigin)
	m.PaymentRefundGrossRevenue = getOrderViewMoney(decoded.PaymentRefundGrossRevenue)
	m.PaymentRefundTaxFee = getOrderViewMoney(decoded.PaymentRefundTaxFee)
	m.PaymentRefundTaxFeeLocal = getOrderViewMoney(decoded.PaymentRefundTaxFeeLocal)
	m.PaymentRefundTaxFeeOrigin = getOrderViewMoney(decoded.PaymentRefundTaxFeeOrigin)
	m.PaymentRefundFeeTariff = getOrderViewMoney(decoded.PaymentRefundFeeTariff)
	m.MethodRefundFixedFeeTariff = getOrderViewMoney(decoded.MethodRefundFixedFeeTariff)
	m.RefundGrossRevenue = getOrderViewMoney(decoded.RefundGrossRevenue)
	m.RefundGrossRevenueFx = getOrderViewMoney(decoded.RefundGrossRevenueFx)
	m.MethodRefundFeeTariff = getOrderViewMoney(decoded.MethodRefundFeeTariff)
	m.PaysuperMethodRefundFeeTariffProfit = getOrderViewMoney(decoded.PaysuperMethodRefundFeeTariffProfit)
	m.PaysuperMethodRefundFixedFeeTariffSelfCost = getOrderViewMoney(decoded.PaysuperMethodRefundFixedFeeTariffSelfCost)
	m.MerchantRefundFixedFeeTariff = getOrderViewMoney(decoded.MerchantRefundFixedFeeTariff)
	m.PaysuperMethodRefundFixedFeeTariffProfit = getOrderViewMoney(decoded.PaysuperMethodRefundFixedFeeTariffProfit)
	m.RefundTaxFee = getOrderViewMoney(decoded.RefundTaxFee)
	m.RefundTaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.RefundTaxFeeCurrencyExchangeFee)
	m.PaysuperRefundTaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.PaysuperRefundTaxFeeCurrencyExchangeFee)
	m.RefundTaxFeeTotal = getOrderViewMoney(decoded.RefundTaxFeeTotal)
	m.RefundReverseRevenue = getOrderViewMoney(decoded.RefundReverseRevenue)
	m.RefundFeesTotal = getOrderViewMoney(decoded.RefundFeesTotal)
	m.RefundFeesTotalLocal = getOrderViewMoney(decoded.RefundFeesTotalLocal)
	m.PaysuperRefundTotalProfit = getOrderViewMoney(decoded.PaysuperRefundTotalProfit)
	m.Items = getOrderViewItems(decoded.Items)
	m.MccCode = decoded.MccCode
	m.OperatingCompanyId = decoded.OperatingCompanyId
	m.IsHighRisk = decoded.IsHighRisk
	m.RefundAllowed = decoded.RefundAllowed
	m.VatPayer = decoded.VatPayer
	m.IsProduction = decoded.IsProduction
	m.TaxRate = decoded.TaxRate
	m.MerchantInfo = decoded.MerchantInfo
	m.OrderChargeBeforeVat = decoded.OrderChargeBeforeVat

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.TransactionDate, err = ptypes.TimestampProto(decoded.TransactionDate)
	if err != nil {
		return err
	}

	return nil
}

func (m *OrderViewPublic) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoOrderViewPublic)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Id = decoded.Id.Hex()
	m.Uuid = decoded.Uuid
	m.TotalPaymentAmount = decoded.TotalPaymentAmount
	m.Currency = decoded.Currency
	m.Project = getOrderProject(decoded.Project)
	m.Transaction = decoded.Transaction
	m.PaymentMethod = getPaymentMethodOrder(decoded.PaymentMethod)
	m.CountryCode = decoded.CountryCode
	m.MerchantId = decoded.MerchantId.Hex()
	m.Locale = decoded.Locale
	m.Status = decoded.Status
	m.User = decoded.User
	m.BillingAddress = decoded.BillingAddress
	m.Type = decoded.Type
	m.Issuer = decoded.Issuer
	m.MerchantPayoutCurrency = decoded.MerchantPayoutCurrency
	m.IsVatDeduction = decoded.IsVatDeduction
	m.ParentOrder = decoded.ParentOrder
	m.Cancellation = decoded.Cancellation

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

	m.GrossRevenue = getOrderViewMoney(decoded.GrossRevenue)
	m.TaxFee = getOrderViewMoney(decoded.TaxFee)
	m.TaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.TaxFeeCurrencyExchangeFee)
	m.TaxFeeTotal = getOrderViewMoney(decoded.TaxFeeTotal)
	m.MethodFeeTotal = getOrderViewMoney(decoded.MethodFeeTotal)
	m.MethodFeeTariff = getOrderViewMoney(decoded.MethodFeeTariff)
	m.MethodFixedFeeTariff = getOrderViewMoney(decoded.MethodFixedFeeTariff)
	m.PaysuperFixedFee = getOrderViewMoney(decoded.PaysuperFixedFee)
	m.FeesTotal = getOrderViewMoney(decoded.FeesTotal)
	m.FeesTotalLocal = getOrderViewMoney(decoded.FeesTotalLocal)
	m.NetRevenue = getOrderViewMoney(decoded.NetRevenue)
	m.RefundGrossRevenue = getOrderViewMoney(decoded.RefundGrossRevenue)
	m.MethodRefundFeeTariff = getOrderViewMoney(decoded.MethodRefundFeeTariff)
	m.MerchantRefundFixedFeeTariff = getOrderViewMoney(decoded.MerchantRefundFixedFeeTariff)
	m.RefundTaxFee = getOrderViewMoney(decoded.RefundTaxFee)
	m.RefundTaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.RefundTaxFeeCurrencyExchangeFee)
	m.PaysuperRefundTaxFeeCurrencyExchangeFee = getOrderViewMoney(decoded.PaysuperRefundTaxFeeCurrencyExchangeFee)
	m.RefundReverseRevenue = getOrderViewMoney(decoded.RefundReverseRevenue)
	m.RefundFeesTotal = getOrderViewMoney(decoded.RefundFeesTotal)
	m.RefundFeesTotalLocal = getOrderViewMoney(decoded.RefundFeesTotalLocal)
	m.Items = getOrderViewItems(decoded.Items)
	m.OperatingCompanyId = decoded.OperatingCompanyId
	m.RefundAllowed = decoded.RefundAllowed
	m.OrderCharge = decoded.OrderCharge
	m.PaymentIpCountry = decoded.PaymentIpCountry
	m.IsIpCountryMismatchBin = decoded.IsIpCountryMismatchBin
	m.BillingCountryChangedByUser = decoded.BillingCountryChangedByUser
	m.VatPayer = decoded.VatPayer
	m.IsProduction = decoded.IsProduction

	m.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)
	if err != nil {
		return err
	}

	m.TransactionDate, err = ptypes.TimestampProto(decoded.TransactionDate)
	if err != nil {
		return err
	}

	return nil
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
