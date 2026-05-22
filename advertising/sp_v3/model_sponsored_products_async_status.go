package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsAsyncStatus Entity state for async status
type SponsoredProductsAsyncStatus string

// List of SponsoredProductsAsyncStatus
const (
	SPONSOREDPRODUCTSASYNCSTATUS_WAITING SponsoredProductsAsyncStatus = "WAITING"
	SPONSOREDPRODUCTSASYNCSTATUS_PAUSED  SponsoredProductsAsyncStatus = "PAUSED"
	SPONSOREDPRODUCTSASYNCSTATUS_FAILED  SponsoredProductsAsyncStatus = "FAILED"
	SPONSOREDPRODUCTSASYNCSTATUS_SUCCEED SponsoredProductsAsyncStatus = "SUCCEED"
)

// All allowed values of SponsoredProductsAsyncStatus enum
var AllowedSponsoredProductsAsyncStatusEnumValues = []SponsoredProductsAsyncStatus{
	"WAITING",
	"PAUSED",
	"FAILED",
	"SUCCEED",
}

func (v *SponsoredProductsAsyncStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsAsyncStatus(value)
	for _, existing := range AllowedSponsoredProductsAsyncStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsAsyncStatus", value)
}

// NewSponsoredProductsAsyncStatusFromValue returns a pointer to a valid SponsoredProductsAsyncStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsAsyncStatusFromValue(v string) (*SponsoredProductsAsyncStatus, error) {
	ev := SponsoredProductsAsyncStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsAsyncStatus: valid values are %v", v, AllowedSponsoredProductsAsyncStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsAsyncStatus) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsAsyncStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsAsyncStatus value
func (v SponsoredProductsAsyncStatus) Ptr() *SponsoredProductsAsyncStatus {
	return &v
}

type NullableSponsoredProductsAsyncStatus struct {
	value *SponsoredProductsAsyncStatus
	isSet bool
}

func (v NullableSponsoredProductsAsyncStatus) Get() *SponsoredProductsAsyncStatus {
	return v.value
}

func (v *NullableSponsoredProductsAsyncStatus) Set(val *SponsoredProductsAsyncStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAsyncStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAsyncStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAsyncStatus(val *SponsoredProductsAsyncStatus) *NullableSponsoredProductsAsyncStatus {
	return &NullableSponsoredProductsAsyncStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsAsyncStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAsyncStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
