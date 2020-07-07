package billingpb

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/copier"
	"github.com/paysuper/paysuper-proto/go/recurringpb"
	tools "github.com/paysuper/paysuper-tools/number"
	"time"
)

func (m *Merchant) GetPayoutCurrency() string {
	if m.Banking == nil {
		return ""
	}

	return m.Banking.Currency
}

func (m *Merchant) IsDeleted() bool {
	return m.Status == MerchantStatusDeleted
}

func (m *Merchant) HasAuthorizedEmail() bool {
	return m.Contacts != nil && m.Contacts.Authorized != nil && m.Contacts.Authorized.Email != ""
}

func (m *Merchant) GetAuthorizedEmail() string {
	return m.Contacts.Authorized.Email
}

func (m *Merchant) GetCompanyName() string {
	if m.Company == nil || m.Company.Name == "" {
		return ""
	}
	return m.Company.Name
}

func (m *Merchant) IsAgreementSigningStarted() bool {
	return m.AgreementSignatureData != nil && (!m.HasPspSignature || !m.HasMerchantSignature)
}

func (m *Merchant) GetPrintableStatus() string {
	status := "draft"

	if m.Status == MerchantStatusAgreementSigned {
		status = "live"
	}

	return status
}

func (m *Merchant) GetCompleteStepsCount() int32 {
	count := int32(0)

	if m.Steps.Company {
		count++
	}

	if m.Steps.Contacts {
		count++
	}

	if m.Steps.Banking {
		count++
	}

	if m.Steps.Tariff {
		count++
	}

	return count
}

func (m *Merchant) IsDataComplete() bool {
	return m.IsCompanyComplete() && m.IsContactsComplete() && m.IsBankingComplete() && m.HasTariff()
}

func (m *Merchant) HasTariff() bool {
	return m.Tariff != nil && len(m.Tariff.Payment) > 0 && m.Tariff.Payout != nil &&
		m.Tariff.Payout.MethodFixedFee > 0 && m.Tariff.Payout.MethodFixedFeeCurrency != "" &&
		m.Tariff.HomeRegion != ""

}

func (m *Merchant) HasPrimaryOnboardingUserName() bool {
	return m.User != nil && m.User.FirstName != "" && m.User.LastName != ""
}

func (m *Merchant) GetAddress() string {
	address := ""
	addressParts := []string{
		m.Company.Address,
		m.Company.AddressAdditional,
		m.Company.State,
		m.Company.City,
		m.Company.Country,
		m.Company.Zip,
	}

	for _, part := range addressParts {
		if len(part) > 0 {
			if len(address) > 0 {
				address += ", "
			}
			address += part
		}
	}

	return address
}

func (m *Merchant) IsCompanyComplete() bool {
	return m.Company != nil && m.Company.Name != "" && m.Company.AlternativeName != "" && m.Company.Website != "" &&
		m.Company.Country != "" && m.Company.Zip != "" && m.Company.City != "" &&
		m.Company.Address != "" && m.Company.RegistrationNumber != ""
}

func (m *Merchant) IsContactsComplete() bool {
	return m.Contacts != nil && m.Contacts.Authorized != nil && m.Contacts.Technical != nil &&
		m.Contacts.Authorized.Name != "" && m.Contacts.Authorized.Email != "" && m.Contacts.Authorized.Phone != "" &&
		m.Contacts.Authorized.Position != "" && m.Contacts.Technical.Name != "" && m.Contacts.Technical.Email != "" &&
		m.Contacts.Technical.Phone != ""
}

func (m *Merchant) IsBankingComplete() bool {
	return m.Banking != nil && m.Banking.Currency != "" && m.Banking.Name != "" && m.Banking.Address != "" &&
		m.Banking.AccountNumber != "" && m.Banking.Swift != ""
}

func (m *Merchant) IsHighRisk() bool {
	return m.MccCode == MccCodeHighRisk
}

func (m *Merchant) CanChangeStatusTo(status int32) bool {
	if status == MerchantStatusDraft && (m.Status == MerchantStatusPending || m.Status == MerchantStatusRejected) {
		return true
	}

	if status == MerchantStatusPending && m.Status == MerchantStatusDraft {
		return true
	}

	if status == MerchantStatusAccepted && m.Status == MerchantStatusPending {
		return true
	}

	if status == MerchantStatusAgreementSigning && m.Status == MerchantStatusAccepted {
		return true
	}

	if status == MerchantStatusAgreementSigned && m.Status == MerchantStatusAgreementSigning {
		return true
	}

	if status == MerchantStatusRejected && (m.Status == MerchantStatusPending || m.Status == MerchantStatusAgreementSigning) {
		return true
	}

	if status == MerchantStatusDeleted && (m.Status == MerchantStatusRejected || m.Status == MerchantStatusDraft) {
		return true
	}

	return false
}

func (m *Merchant) GetAuthorizedName() string {
	if m.Contacts == nil || m.Contacts.Authorized == nil || m.Contacts.Authorized.Name == "" {
		return ""
	}
	return m.Contacts.Authorized.Name
}

func (m *Merchant) GetOwnerName() string {
	if m.User == nil {
		return ""
	}

	return m.User.FirstName + " " + m.User.LastName
}

func (m *Merchant) GetOwnerEmail() string {
	if m.User == nil {
		return ""
	}

	return m.User.Email
}

func (m *Merchant) GetProcessingDefaultCurrency() string {
	if m.Banking == nil {
		return ""
	}

	return m.Banking.ProcessingDefaultCurrency
}

func (m *Project) IsProduction() bool {
	return m.Status == ProjectStatusInProduction
}

func (m *Project) IsDeleted() bool {
	return m.Status == ProjectStatusDeleted
}

func (m *Order) CanBeRecreated() bool {
	if m.PrivateStatus != recurringpb.OrderStatusPaymentSystemRejectOnCreate &&
		m.PrivateStatus != recurringpb.OrderStatusPaymentSystemReject &&
		m.PrivateStatus != recurringpb.OrderStatusProjectReject &&
		m.PrivateStatus != recurringpb.OrderStatusPaymentSystemDeclined &&
		m.PrivateStatus != recurringpb.OrderStatusNew &&
		m.PrivateStatus != recurringpb.OrderStatusPaymentSystemCreate &&
		m.PrivateStatus != recurringpb.OrderStatusPaymentSystemCanceled {

		return false
	}

	return true
}

func (m *Order) GetMerchantId() string {
	return m.Project.MerchantId
}

func (m *Order) GetPaymentMethodId() string {
	return m.PaymentMethod.Id
}

func (m *Order) IsDeclined() bool {
	return m.PrivateStatus == recurringpb.OrderStatusPaymentSystemDeclined || m.PrivateStatus == recurringpb.OrderStatusPaymentSystemCanceled
}

func (m *Order) GetDeclineReason() string {
	reason, ok := m.PaymentMethodTxnParams[TxnParamsFieldDeclineReason]

	if !ok {
		return ""
	}

	return reason
}

func (m *Order) GetPrivateDeclineCode() string {
	code, ok := m.PaymentMethodTxnParams[TxnParamsFieldDeclineCode]

	if !ok {
		return ""
	}

	return code
}

func (m *Order) GetPublicDeclineCode() string {
	code := m.GetPrivateDeclineCode()

	if code == "" {
		return ""
	}

	code, ok := DeclineCodeMap[code]

	if !ok {
		return ""
	}

	return code
}

func (m *Order) GetMerchantRoyaltyCurrency() string {
	if m.Project == nil {
		return ""
	}
	return m.Project.MerchantRoyaltyCurrency
}

func (m *Order) GetPaymentMethodName() string {
	return m.PaymentMethod.Name
}

func (m *Order) GetBankCardBrand() (string, error) {
	val, ok := m.PaymentRequisites[PaymentCreateBankCardFieldBrand]

	if !ok {
		return "", errors.New(orderBankCardBrandNotFound)
	}

	return val, nil
}

func (m *Order) GetCountry() string {
	if m.BillingAddress != nil && m.BillingAddress.Country != "" {
		return m.BillingAddress.Country
	}
	if m.User != nil && m.User.Address != nil && m.User.Address.Country != "" {
		return m.User.Address.Country
	}
	return ""
}

func (m *Order) GetPostalCode() string {
	if m.BillingAddress != nil && m.BillingAddress.PostalCode != "" {
		return m.BillingAddress.PostalCode
	}
	if m.User != nil && m.User.Address != nil && m.User.Address.PostalCode != "" {
		return m.User.Address.PostalCode
	}
	return ""
}

func (m *Order) HasEndedStatus() bool {
	return m.PrivateStatus == recurringpb.OrderStatusPaymentSystemReject || m.PrivateStatus == recurringpb.OrderStatusProjectComplete ||
		m.PrivateStatus == recurringpb.OrderStatusProjectReject || m.PrivateStatus == recurringpb.OrderStatusRefund ||
		m.PrivateStatus == recurringpb.OrderStatusChargeback
}

func (m *Order) RefundAllowed() bool {
	v, ok := orderRefundAllowedStatuses[m.PrivateStatus]

	return ok && v == true
}

func (m *Order) FormInputTimeIsEnded() bool {
	t, err := ptypes.Timestamp(m.ExpireDateToFormInput)

	return err != nil || t.Before(time.Now())
}

func (m *Order) GetProjectId() string {
	return m.Project.Id
}

func (m *Order) GetPublicStatus() string {
	st, ok := orderStatusPublicMapping[m.PrivateStatus]
	if !ok {
		return recurringpb.OrderPublicStatusPending
	}
	return st
}

func (m *Order) GetReceiptUserEmail() string {
	if m.User != nil {
		return m.User.Email
	}
	return ""
}

func (m *Order) GetReceiptUserPhone() string {
	if m.User != nil {
		return m.User.Phone
	}
	return ""
}

func (m *Order) GetState() string {
	if m.BillingAddress != nil && m.BillingAddress.State != "" {
		return m.BillingAddress.State
	}
	if m.User != nil && m.User.Address != nil && m.User.Address.State != "" {
		return m.User.Address.State
	}
	return ""
}

func (m *Order) SetNotificationStatus(key string, val bool) {
	if m.IsNotificationsSent == nil {
		m.IsNotificationsSent = make(map[string]bool)
	}
	m.IsNotificationsSent[key] = val
}

func (m *Order) GetNotificationStatus(key string) bool {
	if m.IsNotificationsSent == nil {
		return false
	}
	val, ok := m.IsNotificationsSent[key]
	if !ok {
		return false
	}
	return val
}

func (m *Order) GetCostPaymentMethodName() (string, error) {
	if m.PaymentMethod == nil {
		return "", errors.New(orderPaymentMethodNotSet)
	}
	if m.PaymentMethod.IsBankCard() {
		return m.GetBankCardBrand()
	}
	return m.GetPaymentMethodName(), nil
}

func (m *Order) GetPaymentSystemApiUrl() string {
	return m.PaymentMethod.Params.ApiUrl
}

func (m *Order) GetPaymentFormDataChangeResult() *PaymentFormDataChangeResponseItem {
	item := &PaymentFormDataChangeResponseItem{
		UserAddressDataRequired: m.UserAddressDataRequired,
		UserIpData: &UserIpData{
			Country: m.User.Address.Country,
			City:    m.User.Address.City,
			Zip:     m.User.Address.PostalCode,
		},
		Brand:                  "",
		HasVat:                 m.Tax.Rate > 0,
		VatRate:                tools.ToPrecise(m.Tax.Rate),
		Vat:                    m.Tax.Amount,
		VatInChargeCurrency:    tools.FormatAmount(m.GetTaxAmountInChargeCurrency()),
		ChargeAmount:           m.ChargeAmount,
		ChargeCurrency:         m.ChargeCurrency,
		Currency:               m.Currency,
		Amount:                 m.OrderAmount,
		TotalAmount:            m.TotalPaymentAmount,
		Items:                  m.Items,
		CountryPaymentsAllowed: true,
		CountryChangeAllowed:   m.CountryChangeAllowed(),
	}

	brand, err := m.GetBankCardBrand()
	if err == nil {
		item.Brand = brand
	}

	if m.CountryRestriction != nil {
		item.CountryPaymentsAllowed = m.CountryRestriction.PaymentsAllowed
	}

	return item
}

func (m *Order) GetTaxAmountInChargeCurrency() float64 {
	if m.Tax.Amount == 0 {
		return 0
	}
	if m.Currency == m.ChargeCurrency {
		return m.Tax.Amount
	}
	return tools.GetPercentPartFromAmount(m.ChargeAmount, m.Tax.Rate)
}

func (m *Order) IsDeclinedByCountry() bool {
	return m.PrivateStatus == recurringpb.OrderStatusPaymentSystemDeclined &&
		m.CountryRestriction != nil &&
		m.CountryRestriction.PaymentsAllowed == false &&
		m.CountryRestriction.ChangeAllowed == false
}

func (m *Order) CountryChangeAllowed() bool {
	return m.CountryRestriction == nil || m.CountryRestriction.ChangeAllowed == true
}

func (m *Order) NeedCallbackNotification() bool {
	status := m.GetPublicStatus()
	return status == recurringpb.OrderPublicStatusRefunded || status == recurringpb.OrderPublicStatusProcessed ||
		status == recurringpb.OrderPublicStatusChargeback || m.IsDeclined()
}

func (ou *OrderUser) HasAddress() bool {
	return ou.Address != nil && ou.Address.Country != ""
}

func (ou *OrderUser) GetCountry() string {
	if ou.Address == nil {
		return ""
	}
	return ou.Address.Country
}

func (m *PaymentMethod) IsValid() bool {
	return m.ExternalId != "" &&
		m.Type != "" &&
		m.Group != "" &&
		m.Name != "" &&
		m.TestSettings != nil &&
		m.ProductionSettings != nil
}

func (m *PaymentMethod) IsBankCard() bool {
	return m.Group == PaymentSystemGroupAliasBankCard
}

func (m *PaymentFormPaymentMethod) IsBankCard() bool {
	return m.Group == PaymentSystemGroupAliasBankCard
}

func (m *PaymentMethodOrder) IsBankCard() bool {
	return m.Group == PaymentSystemGroupAliasBankCard
}

func (m *PaymentMethodParams) IsSettingComplete() bool {
	return m.TerminalId != "" && m.Secret != "" && m.SecretCallback != ""
}

func (c *Country) GetVatCurrencyCode() string {
	if c.VatEnabled && c.VatCurrency != "" {
		return c.VatCurrency
	}
	return c.Currency
}

func (c *Country) GetPaymentRestrictions(isForHighRisk bool) (bool, bool) {
	if isForHighRisk {
		return c.HighRiskPaymentsAllowed, c.HighRiskChangeAllowed
	}

	return c.PaymentsAllowed, c.ChangeAllowed
}

func (m *MerchantPaymentMethodRequest) GetPerTransactionCurrency() string {
	return m.Commission.PerTransaction.Currency
}

func (m *MerchantPaymentMethodRequest) HasPerTransactionCurrency() bool {
	return m.Commission.PerTransaction.Currency != ""
}

func (m *MerchantPaymentMethodRequest) HasIntegration() bool {
	return m.Integration.TerminalId != "" && m.Integration.TerminalPassword != "" &&
		m.Integration.TerminalCallbackPassword != ""
}

func (p *Product) IsPricesContainDefaultCurrency() bool {
	_, err := p.GetPriceInCurrency(&PriceGroup{Currency: p.DefaultCurrency})
	if err != nil {
		_, err = p.GetPriceInCurrency(&PriceGroup{Currency: VirtualCurrencyPriceGroup})
	}
	return err == nil
}

func (p *Product) GetPriceInCurrency(group *PriceGroup) (float64, error) {
	for _, price := range p.Prices {
		if group.Currency == VirtualCurrencyPriceGroup && price.IsVirtualCurrency {
			return price.Amount, nil
		}

		if group.Region != "" && price.Region == group.Region {
			return price.Amount, nil
		}

		if group.Region == "" && price.Region == group.Currency {
			return price.Amount, nil
		}
	}

	return 0, ProductNoPriceInCurrencyError
}

func (p *Product) GetLocalizedName(lang string) (string, error) {
	v, ok := p.Name[lang]
	if !ok || len(v) == 0 {
		return "", errors.New(fmt.Sprintf(productNoNameInLanguage, lang))
	}
	return v, nil
}

func (p *Product) GetLocalizedDescription(lang string) (string, error) {
	v, ok := p.Description[lang]
	if !ok || len(v) == 0 {
		return "", errors.New(fmt.Sprintf(productNoDescriptionInLanguage, lang))
	}
	return v, nil
}

func (p *Product) GetLocalizedLongDescription(lang string) (string, error) {
	v, ok := p.LongDescription[lang]
	if !ok || len(v) == 0 {
		return "", errors.New(fmt.Sprintf(productNoLongDescriptionInLanguage, lang))
	}
	return v, nil
}

func (p *KeyProduct) GetLocalizedName(lang string) (string, error) {
	v, ok := p.Name[lang]
	if !ok || len(v) == 0 {
		return "", errors.New(fmt.Sprintf(productNoNameInLanguage, lang))
	}
	return v, nil
}

func (p *KeyProduct) GetLocalizedDescription(lang string) (string, error) {
	v, ok := p.Description[lang]
	if !ok || len(v) == 0 {
		return "", errors.New(fmt.Sprintf(productNoDescriptionInLanguage, lang))
	}
	return v, nil
}

func (p *KeyProduct) GetLocalizedLongDescription(lang string) (string, error) {
	v, ok := p.LongDescription[lang]
	if !ok || len(v) == 0 {
		return "", errors.New(fmt.Sprintf(productNoLongDescriptionInLanguage, lang))
	}
	return v, nil
}

func (p *KeyProduct) GetPriceInCurrencyAndPlatform(group *PriceGroup, platformId string) (float64, error) {
	// todo: add support for virtual currencies here?
	// note: virtual currency has IsVirtualCurrency=true && Currency=""
	for _, platform := range p.Platforms {
		if platform.Id == platformId {
			for _, price := range platform.Prices {
				if group.Region != "" && price.Region == group.Region {
					return price.Amount, nil
				}

				if group.Region == "" && price.Region == group.Currency {
					return price.Amount, nil
				}
			}
		}
	}

	return 0, errors.New(fmt.Sprintf(productNoPriceInCurrency, group.Region))
}

func (m *UserProfile) IsEmailVerified() bool {
	return m.Email.Confirmed
}

func (m *UserProfile) IsPersonalComplete() bool {
	return m.Personal != nil && m.Personal.FirstName != "" && m.Personal.LastName != "" && m.Personal.Position != ""
}

func (m *UserProfile) IsHelpCompleted() bool {
	return m.Help != nil && (m.Help.ProductPromotionAndDevelopment || m.Help.ReleasedGamePromotion || m.Help.InternationalSales || m.Help.Other)
}

func (m *UserProfile) IsCompanyCompleted() bool {
	return m.Company != nil && m.Company.CompanyName != "" && m.Company.Website != "" &&
		m.Company.AnnualIncome != nil && m.Company.NumberOfEmployees != nil && m.Company.KindOfActivity != "" &&
		m.Company.Monetization != nil && m.Company.Monetization.IsComplete() && m.Company.Platforms != nil &&
		m.Company.Platforms.IsComplete()
}

func (m *UserProfile) NeedConfirmEmail() bool {
	return m.IsPersonalComplete() && m.IsHelpCompleted() && m.IsCompanyCompleted() && !m.Email.IsConfirmationEmailSent
}

func (m *UserProfile) HasPersonChanges(profile *UserProfile) bool {
	return m.Personal != nil && profile.Personal != m.Personal
}

func (m *UserProfile) HasHelpChanges(profile *UserProfile) bool {
	return m.Help != nil && profile.Help != m.Help
}

func (m *UserProfile) HasCompanyChanges(profile *UserProfile) bool {
	return m.Company != nil && profile.Company != m.Company
}

func (m *UserProfile) HasCompanyAnnualIncomeChanges(profile *UserProfile) bool {
	return m.Company.AnnualIncome != nil && profile.Company.AnnualIncome != m.Company.AnnualIncome
}

func (m *UserProfile) HasCompanyNumberOfEmployeesChanges(profile *UserProfile) bool {
	return m.Company.NumberOfEmployees != nil && profile.Company.NumberOfEmployees != m.Company.NumberOfEmployees
}

func (m *UserProfile) HasCompanyMonetizationChanges(profile *UserProfile) bool {
	return m.Company.Monetization != nil && profile.Company.Monetization != m.Company.Monetization
}

func (m *UserProfile) HasCompanyPlatformsChanges(profile *UserProfile) bool {
	return m.Company.Platforms != nil && profile.Company.Platforms != m.Company.Platforms
}

func (m *UserProfileCompanyMonetization) IsComplete() bool {
	return m.PaidSubscription || m.InGameAdvertising || m.InGamePurchases || m.PremiumAccess || m.Other
}

func (m *UserProfileCompanyPlatforms) IsComplete() bool {
	return m.PcMac || m.GameConsole || m.MobileDevice || m.WebBrowser || m.Other
}

func (m *OnboardingRequest) HasIdentificationFields() bool {
	return m.Id != "" || (m.User != nil && m.User.Id != "")
}

func (r *ResponseErrorMessage) Error() string {
	return r.Message
}

func (r *ResponseError) Error() string {
	return r.Message.Message
}

func (r *ResponseErrorMessage) GetResponseErrorWithDetails(details string) *ResponseErrorMessage {
	newResponseErrorMessage := new(ResponseErrorMessage)
	err := copier.Copy(&newResponseErrorMessage, &r)

	if err != nil {
		r.Details = details
		return r
	}

	newResponseErrorMessage.Details = details
	return newResponseErrorMessage
}

func (m *RoyaltyReport) ChangesAvailable(newStatus string) bool {
	if m.Status == RoyaltyReportStatusAccepted {
		return false
	}

	if m.Status == RoyaltyReportStatusPending && newStatus != RoyaltyReportStatusAccepted &&
		newStatus != RoyaltyReportStatusDispute {
		return false
	}

	if m.Status == RoyaltyReportStatusCanceled && newStatus != RoyaltyReportStatusPending {
		return false
	}

	if m.Status == RoyaltyReportStatusDispute && newStatus != RoyaltyReportStatusPending {
		return false
	}

	return true
}

func (m *OrderViewPrivate) GetOrderType() string {
	if m.Refund != nil {
		return OrderTypeRefund
	}

	return OrderTypePayment
}

func (m *OrderViewPrivate) GetCardNumber() string {
	if m.PaymentMethod != nil && m.PaymentMethod.Card != nil {
		return m.PaymentMethod.Card.GetMasked()
	}

	return ""
}

func (m *MerchantBanking) GetBankingRequisites() string {
	if m == nil {
		return ""
	}

	return m.Name + "; " + m.Address + "; " + m.AccountNumber + "; " + m.Swift
}
