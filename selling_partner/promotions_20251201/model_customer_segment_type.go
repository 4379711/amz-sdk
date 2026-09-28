package promotions_20251201

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CustomerSegmentType The segment type for targeting specific customer cohorts.
type CustomerSegmentType string

// List of CustomerSegmentType
const (
	CUSTOMERSEGMENTTYPE_BRAND   CustomerSegmentType = "BRAND"
	CUSTOMERSEGMENTTYPE_PROGRAM CustomerSegmentType = "PROGRAM"
)

// All allowed values of CustomerSegmentType enum
var AllowedCustomerSegmentTypeEnumValues = []CustomerSegmentType{
	"BRAND",
	"PROGRAM",
}

func (v *CustomerSegmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerSegmentType(value)
	for _, existing := range AllowedCustomerSegmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerSegmentType", value)
}

// NewCustomerSegmentTypeFromValue returns a pointer to a valid CustomerSegmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerSegmentTypeFromValue(v string) (*CustomerSegmentType, error) {
	ev := CustomerSegmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerSegmentType: valid values are %v", v, AllowedCustomerSegmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerSegmentType) IsValid() bool {
	for _, existing := range AllowedCustomerSegmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomerSegmentType value
func (v CustomerSegmentType) Ptr() *CustomerSegmentType {
	return &v
}

type NullableCustomerSegmentType struct {
	value *CustomerSegmentType
	isSet bool
}

func (v NullableCustomerSegmentType) Get() *CustomerSegmentType {
	return v.value
}

func (v *NullableCustomerSegmentType) Set(val *CustomerSegmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSegmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSegmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSegmentType(val *CustomerSegmentType) *NullableCustomerSegmentType {
	return &NullableCustomerSegmentType{value: val, isSet: true}
}

func (v NullableCustomerSegmentType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCustomerSegmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
