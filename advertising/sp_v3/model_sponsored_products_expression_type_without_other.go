package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsExpressionTypeWithoutOther the model 'SponsoredProductsExpressionTypeWithoutOther'
type SponsoredProductsExpressionTypeWithoutOther string

// List of SponsoredProductsExpressionTypeWithoutOther
const (
	SPONSOREDPRODUCTSEXPRESSIONTYPEWITHOUTOTHER_AUTO   SponsoredProductsExpressionTypeWithoutOther = "AUTO"
	SPONSOREDPRODUCTSEXPRESSIONTYPEWITHOUTOTHER_MANUAL SponsoredProductsExpressionTypeWithoutOther = "MANUAL"
)

// All allowed values of SponsoredProductsExpressionTypeWithoutOther enum
var AllowedSponsoredProductsExpressionTypeWithoutOtherEnumValues = []SponsoredProductsExpressionTypeWithoutOther{
	"AUTO",
	"MANUAL",
}

func (v *SponsoredProductsExpressionTypeWithoutOther) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsExpressionTypeWithoutOther(value)
	for _, existing := range AllowedSponsoredProductsExpressionTypeWithoutOtherEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsExpressionTypeWithoutOther", value)
}

// NewSponsoredProductsExpressionTypeWithoutOtherFromValue returns a pointer to a valid SponsoredProductsExpressionTypeWithoutOther
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsExpressionTypeWithoutOtherFromValue(v string) (*SponsoredProductsExpressionTypeWithoutOther, error) {
	ev := SponsoredProductsExpressionTypeWithoutOther(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsExpressionTypeWithoutOther: valid values are %v", v, AllowedSponsoredProductsExpressionTypeWithoutOtherEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsExpressionTypeWithoutOther) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsExpressionTypeWithoutOtherEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsExpressionTypeWithoutOther value
func (v SponsoredProductsExpressionTypeWithoutOther) Ptr() *SponsoredProductsExpressionTypeWithoutOther {
	return &v
}

type NullableSponsoredProductsExpressionTypeWithoutOther struct {
	value *SponsoredProductsExpressionTypeWithoutOther
	isSet bool
}

func (v NullableSponsoredProductsExpressionTypeWithoutOther) Get() *SponsoredProductsExpressionTypeWithoutOther {
	return v.value
}

func (v *NullableSponsoredProductsExpressionTypeWithoutOther) Set(val *SponsoredProductsExpressionTypeWithoutOther) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsExpressionTypeWithoutOther) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsExpressionTypeWithoutOther) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsExpressionTypeWithoutOther(val *SponsoredProductsExpressionTypeWithoutOther) *NullableSponsoredProductsExpressionTypeWithoutOther {
	return &NullableSponsoredProductsExpressionTypeWithoutOther{value: val, isSet: true}
}

func (v NullableSponsoredProductsExpressionTypeWithoutOther) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsExpressionTypeWithoutOther) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
