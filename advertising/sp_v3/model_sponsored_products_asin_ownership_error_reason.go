package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsAsinOwnershipErrorReason the model 'SponsoredProductsAsinOwnershipErrorReason'
type SponsoredProductsAsinOwnershipErrorReason string

// List of SponsoredProductsAsinOwnershipErrorReason
const (
	SPONSOREDPRODUCTSASINOWNERSHIPERRORREASON_ASIN_NOT_OWNED_BY_AUTHOR SponsoredProductsAsinOwnershipErrorReason = "ASIN_NOT_OWNED_BY_AUTHOR"
)

// All allowed values of SponsoredProductsAsinOwnershipErrorReason enum
var AllowedSponsoredProductsAsinOwnershipErrorReasonEnumValues = []SponsoredProductsAsinOwnershipErrorReason{
	"ASIN_NOT_OWNED_BY_AUTHOR",
}

func (v *SponsoredProductsAsinOwnershipErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsAsinOwnershipErrorReason(value)
	for _, existing := range AllowedSponsoredProductsAsinOwnershipErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsAsinOwnershipErrorReason", value)
}

// NewSponsoredProductsAsinOwnershipErrorReasonFromValue returns a pointer to a valid SponsoredProductsAsinOwnershipErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsAsinOwnershipErrorReasonFromValue(v string) (*SponsoredProductsAsinOwnershipErrorReason, error) {
	ev := SponsoredProductsAsinOwnershipErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsAsinOwnershipErrorReason: valid values are %v", v, AllowedSponsoredProductsAsinOwnershipErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsAsinOwnershipErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsAsinOwnershipErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsAsinOwnershipErrorReason value
func (v SponsoredProductsAsinOwnershipErrorReason) Ptr() *SponsoredProductsAsinOwnershipErrorReason {
	return &v
}

type NullableSponsoredProductsAsinOwnershipErrorReason struct {
	value *SponsoredProductsAsinOwnershipErrorReason
	isSet bool
}

func (v NullableSponsoredProductsAsinOwnershipErrorReason) Get() *SponsoredProductsAsinOwnershipErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsAsinOwnershipErrorReason) Set(val *SponsoredProductsAsinOwnershipErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAsinOwnershipErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAsinOwnershipErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAsinOwnershipErrorReason(val *SponsoredProductsAsinOwnershipErrorReason) *NullableSponsoredProductsAsinOwnershipErrorReason {
	return &NullableSponsoredProductsAsinOwnershipErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsAsinOwnershipErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAsinOwnershipErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
