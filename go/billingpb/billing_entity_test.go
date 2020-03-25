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
