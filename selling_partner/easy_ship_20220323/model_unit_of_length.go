package easy_ship_20220323

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// UnitOfLength The unit of measurement used to measure the length.
type UnitOfLength string

// List of UnitOfLength
const (
	UNITOFLENGTH_CM UnitOfLength = "Cm"
)

// All allowed values of UnitOfLength enum
var AllowedUnitOfLengthEnumValues = []UnitOfLength{
	UNITOFLENGTH_CM,
}

func (v *UnitOfLength) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UnitOfLength(value)
	for _, existing := range AllowedUnitOfLengthEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UnitOfLength", value)
}

// NewUnitOfLengthFromValue returns a pointer to a valid UnitOfLength
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUnitOfLengthFromValue(v string) (*UnitOfLength, error) {
	ev := UnitOfLength(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UnitOfLength: valid values are %v", v, AllowedUnitOfLengthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UnitOfLength) IsValid() bool {
	for _, existing := range AllowedUnitOfLengthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnitOfLength value
func (v UnitOfLength) Ptr() *UnitOfLength {
	return &v
}

type NullableUnitOfLength struct {
	value *UnitOfLength
	isSet bool
}

func (v NullableUnitOfLength) Get() *UnitOfLength {
	return v.value
}

func (v *NullableUnitOfLength) Set(val *UnitOfLength) {
	v.value = val
	v.isSet = true
}

func (v NullableUnitOfLength) IsSet() bool {
	return v.isSet
}

func (v *NullableUnitOfLength) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnitOfLength(val *UnitOfLength) *NullableUnitOfLength {
	return &NullableUnitOfLength{value: val, isSet: true}
}

func (v NullableUnitOfLength) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnitOfLength) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
