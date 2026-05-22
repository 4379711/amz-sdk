package replenishment20221107

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// AutoEnrollmentPreference The auto-enrollment preference indicates whether the offer is opted-in to or opted-out of Amazon's auto-enrollment feature.
type AutoEnrollmentPreference string

// List of AutoEnrollmentPreference
const (
	AUTOENROLLMENTPREFERENCE_OPTED_IN  AutoEnrollmentPreference = "OPTED_IN"
	AUTOENROLLMENTPREFERENCE_OPTED_OUT AutoEnrollmentPreference = "OPTED_OUT"
)

// All allowed values of AutoEnrollmentPreference enum
var AllowedAutoEnrollmentPreferenceEnumValues = []AutoEnrollmentPreference{
	AUTOENROLLMENTPREFERENCE_OPTED_IN,
	AUTOENROLLMENTPREFERENCE_OPTED_OUT,
}

func (v *AutoEnrollmentPreference) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AutoEnrollmentPreference(value)
	for _, existing := range AllowedAutoEnrollmentPreferenceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AutoEnrollmentPreference", value)
}

// NewAutoEnrollmentPreferenceFromValue returns a pointer to a valid AutoEnrollmentPreference
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAutoEnrollmentPreferenceFromValue(v string) (*AutoEnrollmentPreference, error) {
	ev := AutoEnrollmentPreference(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AutoEnrollmentPreference: valid values are %v", v, AllowedAutoEnrollmentPreferenceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AutoEnrollmentPreference) IsValid() bool {
	for _, existing := range AllowedAutoEnrollmentPreferenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutoEnrollmentPreference value
func (v AutoEnrollmentPreference) Ptr() *AutoEnrollmentPreference {
	return &v
}

type NullableAutoEnrollmentPreference struct {
	value *AutoEnrollmentPreference
	isSet bool
}

func (v NullableAutoEnrollmentPreference) Get() *AutoEnrollmentPreference {
	return v.value
}

func (v *NullableAutoEnrollmentPreference) Set(val *AutoEnrollmentPreference) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoEnrollmentPreference) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoEnrollmentPreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoEnrollmentPreference(val *AutoEnrollmentPreference) *NullableAutoEnrollmentPreference {
	return &NullableAutoEnrollmentPreference{value: val, isSet: true}
}

func (v NullableAutoEnrollmentPreference) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAutoEnrollmentPreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
