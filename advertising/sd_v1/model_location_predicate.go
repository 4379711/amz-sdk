package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// LocationPredicate The location category.
type LocationPredicate string

// List of LocationPredicate
const (
	LOCATIONPREDICATE_LOCATION LocationPredicate = "location"
)

// All allowed values of LocationPredicate enum
var AllowedLocationPredicateEnumValues = []LocationPredicate{
	"location",
}

func (v *LocationPredicate) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationPredicate(value)
	for _, existing := range AllowedLocationPredicateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocationPredicate", value)
}

// NewLocationPredicateFromValue returns a pointer to a valid LocationPredicate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationPredicateFromValue(v string) (*LocationPredicate, error) {
	ev := LocationPredicate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocationPredicate: valid values are %v", v, AllowedLocationPredicateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocationPredicate) IsValid() bool {
	for _, existing := range AllowedLocationPredicateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocationPredicate value
func (v LocationPredicate) Ptr() *LocationPredicate {
	return &v
}

type NullableLocationPredicate struct {
	value *LocationPredicate
	isSet bool
}

func (v NullableLocationPredicate) Get() *LocationPredicate {
	return v.value
}

func (v *NullableLocationPredicate) Set(val *LocationPredicate) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationPredicate) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationPredicate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationPredicate(val *LocationPredicate) *NullableLocationPredicate {
	return &NullableLocationPredicate{value: val, isSet: true}
}

func (v NullableLocationPredicate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLocationPredicate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
