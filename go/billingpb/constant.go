package billingpb

const (
	ServiceName    = "p1paybilling"
	ServiceVersion = "latest"

	ResponseStatusOk          = int32(200)
	ResponseStatusNotModified = int32(304)
	ResponseStatusBadData     = int32(400)
	ResponseStatusNotFound    = int32(404)
	ResponseStatusForbidden   = int32(403)
	ResponseStatusGone        = int32(410)
	ResponseStatusSystemError = int32(500)
	ResponseStatusTemporary   = int32(410)

	OrderTypeOrder  = "order"
	OrderTypeRefund = "refund"

	VatReportStatusThreshold = "threshold"
	VatReportStatusExpired   = "expired"
	VatReportStatusPending   = "pending"
	VatReportStatusNeedToPay = "need_to_pay"
	VatReportStatusPaid      = "paid"
	VatReportStatusOverdue   = "overdue"
	VatReportStatusCanceled  = "canceled"

	ErrorGrpcServiceCallFailed = "gRPC call failed"
)
