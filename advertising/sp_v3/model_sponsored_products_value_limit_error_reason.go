package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsValueLimitErrorReason the model 'SponsoredProductsValueLimitErrorReason'
type SponsoredProductsValueLimitErrorReason string

// List of SponsoredProductsValueLimitErrorReason
const (
	SPONSOREDPRODUCTSVALUELIMITERRORREASON_TOO_LOW            SponsoredProductsValueLimitErrorReason = "TOO_LOW"
	SPONSOREDPRODUCTSVALUELIMITERRORREASON_TOO_HIGH           SponsoredProductsValueLimitErrorReason = "TOO_HIGH"
	SPONSOREDPRODUCTSVALUELIMITERRORREASON_INVALID_ENUM_VALUE SponsoredProductsValueLimitErrorReason = "INVALID_ENUM_VALUE"
	SPONSOREDPRODUCTSVALUELIMITERRORREASON_NOT_IN_LIST        SponsoredProductsValueLimitErrorReason = "NOT_IN_LIST"
)

// All allowed values of SponsoredProductsValueLimitErrorReason enum
var AllowedSponsoredProductsValueLimitErrorReasonEnumValues = []SponsoredProductsValueLimitErrorReason{
	"TOO_LOW",
	"TOO_HIGH",
	"INVALID_ENUM_VALUE",
	"NOT_IN_LIST",
}

func (v *SponsoredProductsValueLimitErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsValueLimitErrorReason(value)
	for _, existing := range AllowedSponsoredProductsValueLimitErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsValueLimitErrorReason", value)
}

// NewSponsoredProductsValueLimitErrorReasonFromValue returns a pointer to a valid SponsoredProductsValueLimitErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsValueLimitErrorReasonFromValue(v string) (*SponsoredProductsValueLimitErrorReason, error) {
	ev := SponsoredProductsValueLimitErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsValueLimitErrorReason: valid values are %v", v, AllowedSponsoredProductsValueLimitErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsValueLimitErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsValueLimitErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsValueLimitErrorReason value
func (v SponsoredProductsValueLimitErrorReason) Ptr() *SponsoredProductsValueLimitErrorReason {
	return &v
}

type NullableSponsoredProductsValueLimitErrorReason struct {
	value *SponsoredProductsValueLimitErrorReason
	isSet bool
}

func (v NullableSponsoredProductsValueLimitErrorReason) Get() *SponsoredProductsValueLimitErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsValueLimitErrorReason) Set(val *SponsoredProductsValueLimitErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsValueLimitErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsValueLimitErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsValueLimitErrorReason(val *SponsoredProductsValueLimitErrorReason) *NullableSponsoredProductsValueLimitErrorReason {
	return &NullableSponsoredProductsValueLimitErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsValueLimitErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsValueLimitErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
