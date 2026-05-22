package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsUnsupportedOperationErrorReason the model 'SponsoredProductsUnsupportedOperationErrorReason'
type SponsoredProductsUnsupportedOperationErrorReason string

// List of SponsoredProductsUnsupportedOperationErrorReason
const (
	SPONSOREDPRODUCTSUNSUPPORTEDOPERATIONERRORREASON_UNSUPPORTED_OPERATION SponsoredProductsUnsupportedOperationErrorReason = "UNSUPPORTED_OPERATION"
)

// All allowed values of SponsoredProductsUnsupportedOperationErrorReason enum
var AllowedSponsoredProductsUnsupportedOperationErrorReasonEnumValues = []SponsoredProductsUnsupportedOperationErrorReason{
	"UNSUPPORTED_OPERATION",
}

func (v *SponsoredProductsUnsupportedOperationErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsUnsupportedOperationErrorReason(value)
	for _, existing := range AllowedSponsoredProductsUnsupportedOperationErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsUnsupportedOperationErrorReason", value)
}

// NewSponsoredProductsUnsupportedOperationErrorReasonFromValue returns a pointer to a valid SponsoredProductsUnsupportedOperationErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsUnsupportedOperationErrorReasonFromValue(v string) (*SponsoredProductsUnsupportedOperationErrorReason, error) {
	ev := SponsoredProductsUnsupportedOperationErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsUnsupportedOperationErrorReason: valid values are %v", v, AllowedSponsoredProductsUnsupportedOperationErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsUnsupportedOperationErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsUnsupportedOperationErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsUnsupportedOperationErrorReason value
func (v SponsoredProductsUnsupportedOperationErrorReason) Ptr() *SponsoredProductsUnsupportedOperationErrorReason {
	return &v
}

type NullableSponsoredProductsUnsupportedOperationErrorReason struct {
	value *SponsoredProductsUnsupportedOperationErrorReason
	isSet bool
}

func (v NullableSponsoredProductsUnsupportedOperationErrorReason) Get() *SponsoredProductsUnsupportedOperationErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsUnsupportedOperationErrorReason) Set(val *SponsoredProductsUnsupportedOperationErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUnsupportedOperationErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUnsupportedOperationErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUnsupportedOperationErrorReason(val *SponsoredProductsUnsupportedOperationErrorReason) *NullableSponsoredProductsUnsupportedOperationErrorReason {
	return &NullableSponsoredProductsUnsupportedOperationErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsUnsupportedOperationErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUnsupportedOperationErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
