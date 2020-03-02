package billingpb

func (m *PaymentFormPaymentMethod) IsBankCard() bool {
	return m.Group == PaymentSystemGroupAliasBankCard
}

func (m *PaymentMethod) IsBankCard() bool {
	return m.Group == PaymentSystemGroupAliasBankCard
}

func (m *PaymentMethodOrder) IsBankCard() bool {
	return m.Group == PaymentSystemGroupAliasBankCard
}

func (m *PaymentMethodOrder) IsCryptoCurrency() bool {
	return m.Group == PaymentSystemGroupAliasBitcoin
}
