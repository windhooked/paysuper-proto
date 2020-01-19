package billingpb

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *Merchant) ChangesAllowed() bool {
	return m.Status == MerchantStatusDraft
}

func (m *Merchant) GetPayoutCurrency() string {
	if m.Banking == nil {
		return ""
	}

	return m.Banking.Currency
}

func (m *Merchant) CanGenerateAgreement() bool {
	return (m.Status == MerchantStatusAgreementSigning ||
		m.Status == MerchantStatusAgreementSigned) && m.Banking != nil && m.Company.Country != "" &&
		m.Contacts != nil && m.Contacts.Authorized != nil
}

func (m *Merchant) IsFullySigned() bool {
	return m.HasMerchantSignature && m.HasPspSignature
}

func (m *Merchant) IsDeleted() bool {
	return m.Status == MerchantStatusDeleted
}

func (m *Project) IsProduction() bool {
	return m.Status == ProjectStatusInProduction
}

func (m *Project) IsDeleted() bool {
	return m.Status == ProjectStatusDeleted
}

func (m *Project) NeedChangeStatusToDraft(req *Project) bool {
	if m.Status != ProjectStatusTestCompleted &&
		m.Status != ProjectStatusInProduction {
		return false
	}

	if m.CallbackProtocol == ProjectCallbackProtocolEmpty &&
		req.CallbackProtocol == ProjectCallbackProtocolDefault {
		return true
	}

	if req.UrlCheckAccount != "" &&
		req.UrlCheckAccount != m.UrlCheckAccount {
		return true
	}

	if req.UrlProcessPayment != "" &&
		req.UrlProcessPayment != m.UrlProcessPayment {
		return true
	}

	return false
}

func (m *OrderUser) IsIdentified() bool {
	_, err := primitive.ObjectIDFromHex(m.Id)
	return err == nil
}

func (m *PaymentMethod) IsValid() bool {
	return m.ExternalId != "" &&
		m.Type != "" &&
		m.Group != "" &&
		m.Name != "" &&
		m.TestSettings != nil &&
		m.ProductionSettings != nil
}

func (m *Merchant) HasAuthorizedEmail() bool {
	return m.Contacts != nil && m.Contacts.Authorized != nil && m.Contacts.Authorized.Email != ""
}

func (m *Merchant) GetAuthorizedEmail() string {
	return m.Contacts.Authorized.Email
}

func (m *Merchant) GetAuthorizedName() string {
	if m.Contacts == nil || m.Contacts.Authorized == nil || m.Contacts.Authorized.Name == "" {
		return ""
	}
	return m.Contacts.Authorized.Name
}

func (m *Merchant) GetCompanyName() string {
	if m.Company == nil || m.Company.Name == "" {
		return ""
	}
	return m.Company.Name
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

func (m *Merchant) IsAgreementSigningStarted() bool {
	return m.AgreementSignatureData != nil && (!m.HasPspSignature || !m.HasMerchantSignature)
}

func (m *Merchant) CanSignAgreement(singerType int32) bool {
	return m.AgreementSignatureData != nil &&
		(singerType == SignerTypeMerchant && m.Status == MerchantStatusAccepted) ||
		(singerType == SignerTypePs && m.Status == MerchantStatusAgreementSigning)
}

func (m *Merchant) IsAgreementSigned() bool {
	return m.HasMerchantSignature && m.HasPspSignature
}

func (m *Merchant) GetPrintableStatus() string {
	status := "draft"

	if m.Status == MerchantStatusAgreementSigned {
		status = "life"
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

func (m *Merchant) GetMerchantSignatureId() string {
	return m.AgreementSignatureData.MerchantSignatureId
}

func (m *Merchant) GetPaysuperSignatureId() string {
	return m.AgreementSignatureData.PsSignatureId
}

func (m *Merchant) GetMerchantSignUrl() *MerchantAgreementSignatureDataSignUrl {
	return m.AgreementSignatureData.MerchantSignUrl
}

func (m *Merchant) GetPaysuperSignUrl() *MerchantAgreementSignatureDataSignUrl {
	return m.AgreementSignatureData.PsSignUrl
}

func (m *Merchant) IsPaysuperSignatureId(signatureId string) bool {
	return m.AgreementSignatureData.PsSignatureId == signatureId
}

func (m *Merchant) IsMerchantSignature(signatureId string) bool {
	return m.AgreementSignatureData.MerchantSignatureId == signatureId
}

func (m *Merchant) HasTariff() bool {
	return m.Tariff != nil && len(m.Tariff.Payment) > 0 && m.Tariff.Payout != nil &&
		m.Tariff.Payout.MethodFixedFee > 0 && m.Tariff.Payout.MethodFixedFeeCurrency != "" && m.Tariff.HomeRegion != ""

}

func (m *Merchant) HasPrimaryOnboardingUserName() bool {
	return m.User != nil && m.User.FirstName != "" && m.User.LastName != ""
}

func (m *Merchant) GetAddress() string {
	address := m.Company.Address + ", " + m.Company.AddressAdditional

	if m.Company.State != "" {
		address += ", " + m.Company.State
	}

	address += ", " + m.Company.City + ", " + m.Company.Country + ", " + m.Company.Zip

	return address
}

func (m *PaymentMethodParams) IsSettingComplete() bool {
	return m.TerminalId != "" && m.Secret != "" && m.SecretCallback != ""
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

func (m *Project) GetVirtualCurrencyRate(group *PriceGroup) (float64, error) {
	for _, price := range m.VirtualCurrency.Prices {
		if group.Region != "" && price.Region == group.Region {
			return price.Amount, nil
		}

		if group.Region == "" && price.Region == group.Currency {
			return price.Amount, nil
		}
	}

	return 0, errors.New(fmt.Sprintf("no price in currency %s", group.Region))
}

func (m *Merchant) IsHighRisk() bool {
	return m.MccCode == MccCodeHighRisk
}

func (m *UserRole) IsOwner() bool {
	return m.Role == RoleMerchantOwner
}

func (m *UserRole) IsAdmin() bool {
	return m.Role == RoleSystemAdmin
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

func (ou *OrderUser) HasAddress() bool {
	return ou.Address != nil && ou.Address.Country != ""
}

func (ou *OrderUser) GetCountry() string {
	if ou.Address == nil {
		return ""
	}
	return ou.Address.Country
}
