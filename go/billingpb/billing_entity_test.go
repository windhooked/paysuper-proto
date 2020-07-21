package billingpb

import (
	"testing"
)

func TestMerchant_GetAddress(t *testing.T) {
	tests := []struct {
		name     string
		merchant *Merchant
		want     string
	}{
		{
			name:     "Full merchant address",
			merchant: &Merchant{Company: &MerchantCompanyInfo{Address: "Address", AddressAdditional: "AddressAdditional", State: "State", Country: "Country", City: "City", Zip: "Zip"}},
			want:     "Address, AddressAdditional, State, City, Country, Zip",
		},
		{
			name:     "Partial merchant address",
			merchant: &Merchant{Company: &MerchantCompanyInfo{Address: "Address", AddressAdditional: "AddressAdditional", Zip: "Zip"}},
			want:     "Address, AddressAdditional, Zip",
		},
		{
			name:     "No address",
			merchant: &Merchant{Company: &MerchantCompanyInfo{AddressAdditional: "AddressAdditional", State: "State", Country: "Country", City: "City", Zip: "Zip"}},
			want:     "AddressAdditional, State, City, Country, Zip",
		},
		{
			name:     "No address at all",
			merchant: &Merchant{Company: &MerchantCompanyInfo{}},
			want:     "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.merchant.GetAddress(); got != tt.want {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderViewPrivate_GetCustomerAddress(t *testing.T) {
	billingAddress := &OrderBillingAddress{
		Country:    "US",
		City:       "ELMIRA",
		State:      "NY",
		PostalCode: "14901",
	}
	geoLocationAddress := &OrderBillingAddress{
		Country:    "US",
		City:       "CUBA",
		State:      "AL",
		PostalCode: "36907",
	}

	tests := []struct {
		name   string
		entity *OrderViewPrivate
		want   *OrderBillingAddress
	}{
		{
			name:   "Object is null",
			entity: nil,
			want:   nil,
		},
		{
			name:   "User.Address and BillingAddress is null",
			entity: &OrderViewPrivate{},
			want:   nil,
		},
		{
			name: "Return BillingAddress",
			entity: &OrderViewPrivate{
				BillingAddress: billingAddress,
			},
			want: billingAddress,
		},
		{
			name: "Return User.Address",
			entity: &OrderViewPrivate{
				User: &OrderUser{
					Address: geoLocationAddress,
				},
			},
			want: geoLocationAddress,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.entity.GetCustomerAddress(); got != tt.want {
				t.Errorf("GetCustomerAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
