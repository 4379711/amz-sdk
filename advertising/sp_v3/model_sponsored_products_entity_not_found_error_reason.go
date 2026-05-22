package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsEntityNotFoundErrorReason the model 'SponsoredProductsEntityNotFoundErrorReason'
type SponsoredProductsEntityNotFoundErrorReason string

// List of SponsoredProductsEntityNotFoundErrorReason
const (
	SPONSOREDPRODUCTSENTITYNOTFOUNDERRORREASON_ENTITY_NOT_FOUND SponsoredProductsEntityNotFoundErrorReason = "ENTITY_NOT_FOUND"
)

// All allowed values of SponsoredProductsEntityNotFoundErrorReason enum
var AllowedSponsoredProductsEntityNotFoundErrorReasonEnumValues = []SponsoredProductsEntityNotFoundErrorReason{
	"ENTITY_NOT_FOUND",
}

func (v *SponsoredProductsEntityNotFoundErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsEntityNotFoundErrorReason(value)
	for _, existing := range AllowedSponsoredProductsEntityNotFoundErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsEntityNotFoundErrorReason", value)
}

// NewSponsoredProductsEntityNotFoundErrorReasonFromValue returns a pointer to a valid SponsoredProductsEntityNotFoundErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsEntityNotFoundErrorReasonFromValue(v string) (*SponsoredProductsEntityNotFoundErrorReason, error) {
	ev := SponsoredProductsEntityNotFoundErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsEntityNotFoundErrorReason: valid values are %v", v, AllowedSponsoredProductsEntityNotFoundErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsEntityNotFoundErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsEntityNotFoundErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsEntityNotFoundErrorReason value
func (v SponsoredProductsEntityNotFoundErrorReason) Ptr() *SponsoredProductsEntityNotFoundErrorReason {
	return &v
}

type NullableSponsoredProductsEntityNotFoundErrorReason struct {
	value *SponsoredProductsEntityNotFoundErrorReason
	isSet bool
}

func (v NullableSponsoredProductsEntityNotFoundErrorReason) Get() *SponsoredProductsEntityNotFoundErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsEntityNotFoundErrorReason) Set(val *SponsoredProductsEntityNotFoundErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsEntityNotFoundErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsEntityNotFoundErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsEntityNotFoundErrorReason(val *SponsoredProductsEntityNotFoundErrorReason) *NullableSponsoredProductsEntityNotFoundErrorReason {
	return &NullableSponsoredProductsEntityNotFoundErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsEntityNotFoundErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsEntityNotFoundErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
