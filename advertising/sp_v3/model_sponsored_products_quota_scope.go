package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsQuotaScope the model 'SponsoredProductsQuotaScope'
type SponsoredProductsQuotaScope string

// List of SponsoredProductsQuotaScope
const (
	SPONSOREDPRODUCTSQUOTASCOPE_ACCOUNT       SponsoredProductsQuotaScope = "ACCOUNT"
	SPONSOREDPRODUCTSQUOTASCOPE_PARENT_ENTITY SponsoredProductsQuotaScope = "PARENT_ENTITY"
)

// All allowed values of SponsoredProductsQuotaScope enum
var AllowedSponsoredProductsQuotaScopeEnumValues = []SponsoredProductsQuotaScope{
	"ACCOUNT",
	"PARENT_ENTITY",
}

func (v *SponsoredProductsQuotaScope) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsQuotaScope(value)
	for _, existing := range AllowedSponsoredProductsQuotaScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsQuotaScope", value)
}

// NewSponsoredProductsQuotaScopeFromValue returns a pointer to a valid SponsoredProductsQuotaScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsQuotaScopeFromValue(v string) (*SponsoredProductsQuotaScope, error) {
	ev := SponsoredProductsQuotaScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsQuotaScope: valid values are %v", v, AllowedSponsoredProductsQuotaScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsQuotaScope) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsQuotaScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsQuotaScope value
func (v SponsoredProductsQuotaScope) Ptr() *SponsoredProductsQuotaScope {
	return &v
}

type NullableSponsoredProductsQuotaScope struct {
	value *SponsoredProductsQuotaScope
	isSet bool
}

func (v NullableSponsoredProductsQuotaScope) Get() *SponsoredProductsQuotaScope {
	return v.value
}

func (v *NullableSponsoredProductsQuotaScope) Set(val *SponsoredProductsQuotaScope) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsQuotaScope) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsQuotaScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsQuotaScope(val *SponsoredProductsQuotaScope) *NullableSponsoredProductsQuotaScope {
	return &NullableSponsoredProductsQuotaScope{value: val, isSet: true}
}

func (v NullableSponsoredProductsQuotaScope) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsQuotaScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
