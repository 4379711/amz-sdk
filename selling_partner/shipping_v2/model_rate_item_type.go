package shipping_v2

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// RateItemType Type of the rateItem.
type RateItemType string

// List of RateItemType
const (
	RATEITEMTYPE_MANDATORY RateItemType = "MANDATORY"
	RATEITEMTYPE_OPTIONAL  RateItemType = "OPTIONAL"
	RATEITEMTYPE_INCLUDED  RateItemType = "INCLUDED"
)

// All allowed values of RateItemType enum
var AllowedRateItemTypeEnumValues = []RateItemType{
	RATEITEMTYPE_MANDATORY,
	RATEITEMTYPE_OPTIONAL,
	RATEITEMTYPE_INCLUDED,
}

func (v *RateItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RateItemType(value)
	for _, existing := range AllowedRateItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RateItemType", value)
}

// NewRateItemTypeFromValue returns a pointer to a valid RateItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRateItemTypeFromValue(v string) (*RateItemType, error) {
	ev := RateItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RateItemType: valid values are %v", v, AllowedRateItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RateItemType) IsValid() bool {
	for _, existing := range AllowedRateItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RateItemType value
func (v RateItemType) Ptr() *RateItemType {
	return &v
}

type NullableRateItemType struct {
	value *RateItemType
	isSet bool
}

func (v NullableRateItemType) Get() *RateItemType {
	return v.value
}

func (v *NullableRateItemType) Set(val *RateItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableRateItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableRateItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateItemType(val *RateItemType) *NullableRateItemType {
	return &NullableRateItemType{value: val, isSet: true}
}

func (v NullableRateItemType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRateItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
