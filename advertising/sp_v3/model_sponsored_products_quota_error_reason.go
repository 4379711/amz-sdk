package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsQuotaErrorReason the model 'SponsoredProductsQuotaErrorReason'
type SponsoredProductsQuotaErrorReason string

// List of SponsoredProductsQuotaErrorReason
const (
	SPONSOREDPRODUCTSQUOTAERRORREASON_QUOTA_EXCEEDED              SponsoredProductsQuotaErrorReason = "QUOTA_EXCEEDED"
	SPONSOREDPRODUCTSQUOTAERRORREASON_NON_ARCHIVED_QUOTA_EXCEEDED SponsoredProductsQuotaErrorReason = "NON_ARCHIVED_QUOTA_EXCEEDED"
)

// All allowed values of SponsoredProductsQuotaErrorReason enum
var AllowedSponsoredProductsQuotaErrorReasonEnumValues = []SponsoredProductsQuotaErrorReason{
	"QUOTA_EXCEEDED",
	"NON_ARCHIVED_QUOTA_EXCEEDED",
}

func (v *SponsoredProductsQuotaErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsQuotaErrorReason(value)
	for _, existing := range AllowedSponsoredProductsQuotaErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsQuotaErrorReason", value)
}

// NewSponsoredProductsQuotaErrorReasonFromValue returns a pointer to a valid SponsoredProductsQuotaErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsQuotaErrorReasonFromValue(v string) (*SponsoredProductsQuotaErrorReason, error) {
	ev := SponsoredProductsQuotaErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsQuotaErrorReason: valid values are %v", v, AllowedSponsoredProductsQuotaErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsQuotaErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsQuotaErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsQuotaErrorReason value
func (v SponsoredProductsQuotaErrorReason) Ptr() *SponsoredProductsQuotaErrorReason {
	return &v
}

type NullableSponsoredProductsQuotaErrorReason struct {
	value *SponsoredProductsQuotaErrorReason
	isSet bool
}

func (v NullableSponsoredProductsQuotaErrorReason) Get() *SponsoredProductsQuotaErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsQuotaErrorReason) Set(val *SponsoredProductsQuotaErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsQuotaErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsQuotaErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsQuotaErrorReason(val *SponsoredProductsQuotaErrorReason) *NullableSponsoredProductsQuotaErrorReason {
	return &NullableSponsoredProductsQuotaErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsQuotaErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsQuotaErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
