package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalEntityServingStatus the model 'SponsoredProductsGlobalEntityServingStatus'
type SponsoredProductsGlobalEntityServingStatus string

// List of SponsoredProductsGlobalEntityServingStatus
const (
	SPONSOREDPRODUCTSGLOBALENTITYSERVINGSTATUS_ACTIVE                 SponsoredProductsGlobalEntityServingStatus = "ACTIVE"
	SPONSOREDPRODUCTSGLOBALENTITYSERVINGSTATUS_INACTIVE               SponsoredProductsGlobalEntityServingStatus = "INACTIVE"
	SPONSOREDPRODUCTSGLOBALENTITYSERVINGSTATUS_ACTIVE_WITH_EXCEPTIONS SponsoredProductsGlobalEntityServingStatus = "ACTIVE_WITH_EXCEPTIONS"
	SPONSOREDPRODUCTSGLOBALENTITYSERVINGSTATUS_ARCHIVED               SponsoredProductsGlobalEntityServingStatus = "ARCHIVED"
)

// All allowed values of SponsoredProductsGlobalEntityServingStatus enum
var AllowedSponsoredProductsGlobalEntityServingStatusEnumValues = []SponsoredProductsGlobalEntityServingStatus{
	"ACTIVE",
	"INACTIVE",
	"ACTIVE_WITH_EXCEPTIONS",
	"ARCHIVED",
}

func (v *SponsoredProductsGlobalEntityServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalEntityServingStatus(value)
	for _, existing := range AllowedSponsoredProductsGlobalEntityServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalEntityServingStatus", value)
}

// NewSponsoredProductsGlobalEntityServingStatusFromValue returns a pointer to a valid SponsoredProductsGlobalEntityServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalEntityServingStatusFromValue(v string) (*SponsoredProductsGlobalEntityServingStatus, error) {
	ev := SponsoredProductsGlobalEntityServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalEntityServingStatus: valid values are %v", v, AllowedSponsoredProductsGlobalEntityServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalEntityServingStatus) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalEntityServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalEntityServingStatus value
func (v SponsoredProductsGlobalEntityServingStatus) Ptr() *SponsoredProductsGlobalEntityServingStatus {
	return &v
}

type NullableSponsoredProductsGlobalEntityServingStatus struct {
	value *SponsoredProductsGlobalEntityServingStatus
	isSet bool
}

func (v NullableSponsoredProductsGlobalEntityServingStatus) Get() *SponsoredProductsGlobalEntityServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsGlobalEntityServingStatus) Set(val *SponsoredProductsGlobalEntityServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalEntityServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalEntityServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalEntityServingStatus(val *SponsoredProductsGlobalEntityServingStatus) *NullableSponsoredProductsGlobalEntityServingStatus {
	return &NullableSponsoredProductsGlobalEntityServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalEntityServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalEntityServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
