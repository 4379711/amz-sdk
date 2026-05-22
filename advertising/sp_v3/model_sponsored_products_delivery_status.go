package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsDeliveryStatus DeliveryStatus enum defines if this entity is being used to serve ads.
type SponsoredProductsDeliveryStatus string

// List of SponsoredProductsDeliveryStatus
const (
	SPONSOREDPRODUCTSDELIVERYSTATUS_DELIVERING     SponsoredProductsDeliveryStatus = "DELIVERING"
	SPONSOREDPRODUCTSDELIVERYSTATUS_NOT_DELIVERING SponsoredProductsDeliveryStatus = "NOT_DELIVERING"
	SPONSOREDPRODUCTSDELIVERYSTATUS_UNAVAILABLE    SponsoredProductsDeliveryStatus = "UNAVAILABLE"
)

// All allowed values of SponsoredProductsDeliveryStatus enum
var AllowedSponsoredProductsDeliveryStatusEnumValues = []SponsoredProductsDeliveryStatus{
	"DELIVERING",
	"NOT_DELIVERING",
	"UNAVAILABLE",
}

func (v *SponsoredProductsDeliveryStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsDeliveryStatus(value)
	for _, existing := range AllowedSponsoredProductsDeliveryStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsDeliveryStatus", value)
}

// NewSponsoredProductsDeliveryStatusFromValue returns a pointer to a valid SponsoredProductsDeliveryStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsDeliveryStatusFromValue(v string) (*SponsoredProductsDeliveryStatus, error) {
	ev := SponsoredProductsDeliveryStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsDeliveryStatus: valid values are %v", v, AllowedSponsoredProductsDeliveryStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsDeliveryStatus) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsDeliveryStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsDeliveryStatus value
func (v SponsoredProductsDeliveryStatus) Ptr() *SponsoredProductsDeliveryStatus {
	return &v
}

type NullableSponsoredProductsDeliveryStatus struct {
	value *SponsoredProductsDeliveryStatus
	isSet bool
}

func (v NullableSponsoredProductsDeliveryStatus) Get() *SponsoredProductsDeliveryStatus {
	return v.value
}

func (v *NullableSponsoredProductsDeliveryStatus) Set(val *SponsoredProductsDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeliveryStatus(val *SponsoredProductsDeliveryStatus) *NullableSponsoredProductsDeliveryStatus {
	return &NullableSponsoredProductsDeliveryStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeliveryStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
