package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsApplicableMarketplacesErrorReason the model 'SponsoredProductsApplicableMarketplacesErrorReason'
type SponsoredProductsApplicableMarketplacesErrorReason string

// List of SponsoredProductsApplicableMarketplacesErrorReason
const (
	SPONSOREDPRODUCTSAPPLICABLEMARKETPLACESERRORREASON_APPLICABLE_MARKETPLACES_MISMATCH_ERROR SponsoredProductsApplicableMarketplacesErrorReason = "APPLICABLE_MARKETPLACES_MISMATCH_ERROR"
)

// All allowed values of SponsoredProductsApplicableMarketplacesErrorReason enum
var AllowedSponsoredProductsApplicableMarketplacesErrorReasonEnumValues = []SponsoredProductsApplicableMarketplacesErrorReason{
	"APPLICABLE_MARKETPLACES_MISMATCH_ERROR",
}

func (v *SponsoredProductsApplicableMarketplacesErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsApplicableMarketplacesErrorReason(value)
	for _, existing := range AllowedSponsoredProductsApplicableMarketplacesErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsApplicableMarketplacesErrorReason", value)
}

// NewSponsoredProductsApplicableMarketplacesErrorReasonFromValue returns a pointer to a valid SponsoredProductsApplicableMarketplacesErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsApplicableMarketplacesErrorReasonFromValue(v string) (*SponsoredProductsApplicableMarketplacesErrorReason, error) {
	ev := SponsoredProductsApplicableMarketplacesErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsApplicableMarketplacesErrorReason: valid values are %v", v, AllowedSponsoredProductsApplicableMarketplacesErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsApplicableMarketplacesErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsApplicableMarketplacesErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsApplicableMarketplacesErrorReason value
func (v SponsoredProductsApplicableMarketplacesErrorReason) Ptr() *SponsoredProductsApplicableMarketplacesErrorReason {
	return &v
}

type NullableSponsoredProductsApplicableMarketplacesErrorReason struct {
	value *SponsoredProductsApplicableMarketplacesErrorReason
	isSet bool
}

func (v NullableSponsoredProductsApplicableMarketplacesErrorReason) Get() *SponsoredProductsApplicableMarketplacesErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsApplicableMarketplacesErrorReason) Set(val *SponsoredProductsApplicableMarketplacesErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsApplicableMarketplacesErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsApplicableMarketplacesErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsApplicableMarketplacesErrorReason(val *SponsoredProductsApplicableMarketplacesErrorReason) *NullableSponsoredProductsApplicableMarketplacesErrorReason {
	return &NullableSponsoredProductsApplicableMarketplacesErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsApplicableMarketplacesErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsApplicableMarketplacesErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
