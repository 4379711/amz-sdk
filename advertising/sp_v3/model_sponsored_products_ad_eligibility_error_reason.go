package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsAdEligibilityErrorReason the model 'SponsoredProductsAdEligibilityErrorReason'
type SponsoredProductsAdEligibilityErrorReason string

// List of SponsoredProductsAdEligibilityErrorReason
const (
	SPONSOREDPRODUCTSADELIGIBILITYERRORREASON_AD_INELIGIBLE SponsoredProductsAdEligibilityErrorReason = "AD_INELIGIBLE"
)

// All allowed values of SponsoredProductsAdEligibilityErrorReason enum
var AllowedSponsoredProductsAdEligibilityErrorReasonEnumValues = []SponsoredProductsAdEligibilityErrorReason{
	"AD_INELIGIBLE",
}

func (v *SponsoredProductsAdEligibilityErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsAdEligibilityErrorReason(value)
	for _, existing := range AllowedSponsoredProductsAdEligibilityErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsAdEligibilityErrorReason", value)
}

// NewSponsoredProductsAdEligibilityErrorReasonFromValue returns a pointer to a valid SponsoredProductsAdEligibilityErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsAdEligibilityErrorReasonFromValue(v string) (*SponsoredProductsAdEligibilityErrorReason, error) {
	ev := SponsoredProductsAdEligibilityErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsAdEligibilityErrorReason: valid values are %v", v, AllowedSponsoredProductsAdEligibilityErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsAdEligibilityErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsAdEligibilityErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsAdEligibilityErrorReason value
func (v SponsoredProductsAdEligibilityErrorReason) Ptr() *SponsoredProductsAdEligibilityErrorReason {
	return &v
}

type NullableSponsoredProductsAdEligibilityErrorReason struct {
	value *SponsoredProductsAdEligibilityErrorReason
	isSet bool
}

func (v NullableSponsoredProductsAdEligibilityErrorReason) Get() *SponsoredProductsAdEligibilityErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsAdEligibilityErrorReason) Set(val *SponsoredProductsAdEligibilityErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdEligibilityErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdEligibilityErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdEligibilityErrorReason(val *SponsoredProductsAdEligibilityErrorReason) *NullableSponsoredProductsAdEligibilityErrorReason {
	return &NullableSponsoredProductsAdEligibilityErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdEligibilityErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdEligibilityErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
