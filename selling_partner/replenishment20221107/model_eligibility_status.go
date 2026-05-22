package replenishment20221107

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// EligibilityStatus The current eligibility status of an offer.
type EligibilityStatus string

// List of EligibilityStatus
const (
	ELIGIBILITYSTATUS_ELIGIBLE                    EligibilityStatus = "ELIGIBLE"
	ELIGIBILITYSTATUS_INELIGIBLE                  EligibilityStatus = "INELIGIBLE"
	ELIGIBILITYSTATUS_SUSPENDED                   EligibilityStatus = "SUSPENDED"
	ELIGIBILITYSTATUS_REPLENISHMENT_ONLY_ORDERING EligibilityStatus = "REPLENISHMENT_ONLY_ORDERING"
)

// All allowed values of EligibilityStatus enum
var AllowedEligibilityStatusEnumValues = []EligibilityStatus{
	ELIGIBILITYSTATUS_ELIGIBLE,
	ELIGIBILITYSTATUS_INELIGIBLE,
	ELIGIBILITYSTATUS_SUSPENDED,
	ELIGIBILITYSTATUS_REPLENISHMENT_ONLY_ORDERING,
}

func (v *EligibilityStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EligibilityStatus(value)
	for _, existing := range AllowedEligibilityStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EligibilityStatus", value)
}

// NewEligibilityStatusFromValue returns a pointer to a valid EligibilityStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEligibilityStatusFromValue(v string) (*EligibilityStatus, error) {
	ev := EligibilityStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EligibilityStatus: valid values are %v", v, AllowedEligibilityStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EligibilityStatus) IsValid() bool {
	for _, existing := range AllowedEligibilityStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EligibilityStatus value
func (v EligibilityStatus) Ptr() *EligibilityStatus {
	return &v
}

type NullableEligibilityStatus struct {
	value *EligibilityStatus
	isSet bool
}

func (v NullableEligibilityStatus) Get() *EligibilityStatus {
	return v.value
}

func (v *NullableEligibilityStatus) Set(val *EligibilityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEligibilityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEligibilityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEligibilityStatus(val *EligibilityStatus) *NullableEligibilityStatus {
	return &NullableEligibilityStatus{value: val, isSet: true}
}

func (v NullableEligibilityStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEligibilityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
