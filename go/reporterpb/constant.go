package reporterpb

const (
	ServiceName    = "p1payreporter"
	ServiceVersion = "latest"

	ReportTypeTransactions        = "transactions"
	ReportTypeVat                 = "vat"
	ReportTypeVatTransactions     = "vat_transactions"
	ReportTypeRoyalty             = "royalty"
	ReportTypeRoyaltyTransactions = "royalty_transactions"
	ReportTypePayout              = "payout"
	ReportTypeAgreement           = "agreement"

	OutputExtensionXlsx = "xlsx"
	OutputExtensionCsv  = "csv"
	OutputExtensionPdf  = "pdf"

	RequestParameterAgreementNumber                             = "number"
	RequestParameterAgreementLegalName                          = "legal_name"
	RequestParameterAgreementAddress                            = "address"
	RequestParameterAgreementRegistrationNumber                 = "registration_number"
	RequestParameterAgreementPayoutCost                         = "payout_cost"
	RequestParameterAgreementMinimalPayoutLimit                 = "minimal_payout_limit"
	RequestParameterAgreementPayoutCurrency                     = "payout_currency"
	RequestParameterAgreementPSRate                             = "ps_rate"
	RequestParameterAgreementHomeRegion                         = "home_region"
	RequestParameterAgreementMerchantAuthorizedName             = "merchant_authorized_name"
	RequestParameterAgreementMerchantAuthorizedPosition         = "merchant_authorized_position"
	RequestParameterAgreementOperatingCompanyLegalName          = "oc_name"
	RequestParameterAgreementOperatingCompanyAddress            = "oc_address"
	RequestParameterAgreementOperatingCompanyRegistrationNumber = "oc_registration_number"
	RequestParameterAgreementOperatingCompanyAuthorizedName     = "oc_authorized_name"
	RequestParameterAgreementOperatingCompanyAuthorizedPosition = "oc_authorized_position"
)
