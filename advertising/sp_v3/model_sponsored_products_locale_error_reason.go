package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsLocaleErrorReason the model 'SponsoredProductsLocaleErrorReason'
type SponsoredProductsLocaleErrorReason string

// List of SponsoredProductsLocaleErrorReason
const (
	SPONSOREDPRODUCTSLOCALEERRORREASON_INVALID_LOCALE SponsoredProductsLocaleErrorReason = "INVALID_LOCALE"
)

// All allowed values of SponsoredProductsLocaleErrorReason enum
var AllowedSponsoredProductsLocaleErrorReasonEnumValues = []SponsoredProductsLocaleErrorReason{
	"INVALID_LOCALE",
}

func (v *SponsoredProductsLocaleErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsLocaleErrorReason(value)
	for _, existing := range AllowedSponsoredProductsLocaleErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsLocaleErrorReason", value)
}

// NewSponsoredProductsLocaleErrorReasonFromValue returns a pointer to a valid SponsoredProductsLocaleErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsLocaleErrorReasonFromValue(v string) (*SponsoredProductsLocaleErrorReason, error) {
	ev := SponsoredProductsLocaleErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsLocaleErrorReason: valid values are %v", v, AllowedSponsoredProductsLocaleErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsLocaleErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsLocaleErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsLocaleErrorReason value
func (v SponsoredProductsLocaleErrorReason) Ptr() *SponsoredProductsLocaleErrorReason {
	return &v
}

type NullableSponsoredProductsLocaleErrorReason struct {
	value *SponsoredProductsLocaleErrorReason
	isSet bool
}

func (v NullableSponsoredProductsLocaleErrorReason) Get() *SponsoredProductsLocaleErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsLocaleErrorReason) Set(val *SponsoredProductsLocaleErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsLocaleErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsLocaleErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsLocaleErrorReason(val *SponsoredProductsLocaleErrorReason) *NullableSponsoredProductsLocaleErrorReason {
	return &NullableSponsoredProductsLocaleErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsLocaleErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsLocaleErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
