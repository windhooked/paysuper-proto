package document_signerpb

const (
	ServiceName    = "psdocumentsigner"
	ServiceVersion = "latest"

	SignerRoleNameMerchant = "Merchant"
	SignerRoleNamePaysuper = "Paysuper"

	MetadataFieldMerchantId       = "merchant_id"
	MetadataFieldPayoutDocumentId = "payout_document_id"

	StorageTypeAgreement = "agreement"
	StorageTypeReport    = "report"

	RequestTypeCreateEmbeddedWithTemplate = "create_embedded_with_template"
	RequestTypeCreateEmbedded             = "create_embedded"
	RequestTypeCreateWebsite              = "create_website"
)
