package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsInvalidInputErrorReason the model 'SponsoredProductsInvalidInputErrorReason'
type SponsoredProductsInvalidInputErrorReason string

// List of SponsoredProductsInvalidInputErrorReason
const (
	SPONSOREDPRODUCTSINVALIDINPUTERRORREASON_INVALID_TOKEN SponsoredProductsInvalidInputErrorReason = "INVALID_TOKEN"
)

// All allowed values of SponsoredProductsInvalidInputErrorReason enum
var AllowedSponsoredProductsInvalidInputErrorReasonEnumValues = []SponsoredProductsInvalidInputErrorReason{
	"INVALID_TOKEN",
}

func (v *SponsoredProductsInvalidInputErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsInvalidInputErrorReason(value)
	for _, existing := range AllowedSponsoredProductsInvalidInputErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsInvalidInputErrorReason", value)
}

// NewSponsoredProductsInvalidInputErrorReasonFromValue returns a pointer to a valid SponsoredProductsInvalidInputErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsInvalidInputErrorReasonFromValue(v string) (*SponsoredProductsInvalidInputErrorReason, error) {
	ev := SponsoredProductsInvalidInputErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsInvalidInputErrorReason: valid values are %v", v, AllowedSponsoredProductsInvalidInputErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsInvalidInputErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsInvalidInputErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsInvalidInputErrorReason value
func (v SponsoredProductsInvalidInputErrorReason) Ptr() *SponsoredProductsInvalidInputErrorReason {
	return &v
}

type NullableSponsoredProductsInvalidInputErrorReason struct {
	value *SponsoredProductsInvalidInputErrorReason
	isSet bool
}

func (v NullableSponsoredProductsInvalidInputErrorReason) Get() *SponsoredProductsInvalidInputErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsInvalidInputErrorReason) Set(val *SponsoredProductsInvalidInputErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsInvalidInputErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsInvalidInputErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsInvalidInputErrorReason(val *SponsoredProductsInvalidInputErrorReason) *NullableSponsoredProductsInvalidInputErrorReason {
	return &NullableSponsoredProductsInvalidInputErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsInvalidInputErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsInvalidInputErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
