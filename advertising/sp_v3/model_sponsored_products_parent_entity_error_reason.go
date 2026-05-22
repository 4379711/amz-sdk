package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsParentEntityErrorReason the model 'SponsoredProductsParentEntityErrorReason'
type SponsoredProductsParentEntityErrorReason string

// List of SponsoredProductsParentEntityErrorReason
const (
	SPONSOREDPRODUCTSPARENTENTITYERRORREASON_DOES_NOT_TARGET_THESE_MARKETPLACES SponsoredProductsParentEntityErrorReason = "PARENT_ENTITY_DOES_NOT_TARGET_THESE_MARKETPLACES"
	SPONSOREDPRODUCTSPARENTENTITYERRORREASON_ARCHIVED                           SponsoredProductsParentEntityErrorReason = "PARENT_ENTITY_ARCHIVED"
	SPONSOREDPRODUCTSPARENTENTITYERRORREASON_NOT_FOUND                          SponsoredProductsParentEntityErrorReason = "PARENT_ENTITY_NOT_FOUND"
)

// All allowed values of SponsoredProductsParentEntityErrorReason enum
var AllowedSponsoredProductsParentEntityErrorReasonEnumValues = []SponsoredProductsParentEntityErrorReason{
	"PARENT_ENTITY_DOES_NOT_TARGET_THESE_MARKETPLACES",
	"PARENT_ENTITY_ARCHIVED",
	"PARENT_ENTITY_NOT_FOUND",
}

func (v *SponsoredProductsParentEntityErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsParentEntityErrorReason(value)
	for _, existing := range AllowedSponsoredProductsParentEntityErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsParentEntityErrorReason", value)
}

// NewSponsoredProductsParentEntityErrorReasonFromValue returns a pointer to a valid SponsoredProductsParentEntityErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsParentEntityErrorReasonFromValue(v string) (*SponsoredProductsParentEntityErrorReason, error) {
	ev := SponsoredProductsParentEntityErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsParentEntityErrorReason: valid values are %v", v, AllowedSponsoredProductsParentEntityErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsParentEntityErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsParentEntityErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsParentEntityErrorReason value
func (v SponsoredProductsParentEntityErrorReason) Ptr() *SponsoredProductsParentEntityErrorReason {
	return &v
}

type NullableSponsoredProductsParentEntityErrorReason struct {
	value *SponsoredProductsParentEntityErrorReason
	isSet bool
}

func (v NullableSponsoredProductsParentEntityErrorReason) Get() *SponsoredProductsParentEntityErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsParentEntityErrorReason) Set(val *SponsoredProductsParentEntityErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsParentEntityErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsParentEntityErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsParentEntityErrorReason(val *SponsoredProductsParentEntityErrorReason) *NullableSponsoredProductsParentEntityErrorReason {
	return &NullableSponsoredProductsParentEntityErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsParentEntityErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsParentEntityErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
