package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsDuplicateValueErrorReason the model 'SponsoredProductsDuplicateValueErrorReason'
type SponsoredProductsDuplicateValueErrorReason string

// List of SponsoredProductsDuplicateValueErrorReason
const (
	SPONSOREDPRODUCTSDUPLICATEVALUEERRORREASON_DUPLICATE_VALUE                 SponsoredProductsDuplicateValueErrorReason = "DUPLICATE_VALUE"
	SPONSOREDPRODUCTSDUPLICATEVALUEERRORREASON_NAME_NOT_UNIQUE                 SponsoredProductsDuplicateValueErrorReason = "NAME_NOT_UNIQUE"
	SPONSOREDPRODUCTSDUPLICATEVALUEERRORREASON_MARKETPLACE_ATTRIBUTES_REPEATED SponsoredProductsDuplicateValueErrorReason = "MARKETPLACE_ATTRIBUTES_REPEATED"
)

// All allowed values of SponsoredProductsDuplicateValueErrorReason enum
var AllowedSponsoredProductsDuplicateValueErrorReasonEnumValues = []SponsoredProductsDuplicateValueErrorReason{
	"DUPLICATE_VALUE",
	"NAME_NOT_UNIQUE",
	"MARKETPLACE_ATTRIBUTES_REPEATED",
}

func (v *SponsoredProductsDuplicateValueErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsDuplicateValueErrorReason(value)
	for _, existing := range AllowedSponsoredProductsDuplicateValueErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsDuplicateValueErrorReason", value)
}

// NewSponsoredProductsDuplicateValueErrorReasonFromValue returns a pointer to a valid SponsoredProductsDuplicateValueErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsDuplicateValueErrorReasonFromValue(v string) (*SponsoredProductsDuplicateValueErrorReason, error) {
	ev := SponsoredProductsDuplicateValueErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsDuplicateValueErrorReason: valid values are %v", v, AllowedSponsoredProductsDuplicateValueErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsDuplicateValueErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsDuplicateValueErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsDuplicateValueErrorReason value
func (v SponsoredProductsDuplicateValueErrorReason) Ptr() *SponsoredProductsDuplicateValueErrorReason {
	return &v
}

type NullableSponsoredProductsDuplicateValueErrorReason struct {
	value *SponsoredProductsDuplicateValueErrorReason
	isSet bool
}

func (v NullableSponsoredProductsDuplicateValueErrorReason) Get() *SponsoredProductsDuplicateValueErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsDuplicateValueErrorReason) Set(val *SponsoredProductsDuplicateValueErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDuplicateValueErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDuplicateValueErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDuplicateValueErrorReason(val *SponsoredProductsDuplicateValueErrorReason) *NullableSponsoredProductsDuplicateValueErrorReason {
	return &NullableSponsoredProductsDuplicateValueErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsDuplicateValueErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDuplicateValueErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
