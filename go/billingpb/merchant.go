package billingpb

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
