package merchant_fulfillment_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// InputTargetType Indicates whether the additional seller input is at the item or shipment level.
type InputTargetType string

// List of InputTargetType
const (
	INPUTTARGETTYPE_SHIPMENT_LEVEL InputTargetType = "SHIPMENT_LEVEL"
	INPUTTARGETTYPE_ITEM_LEVEL     InputTargetType = "ITEM_LEVEL"
)

// All allowed values of InputTargetType enum
var AllowedInputTargetTypeEnumValues = []InputTargetType{
	INPUTTARGETTYPE_SHIPMENT_LEVEL,
	INPUTTARGETTYPE_ITEM_LEVEL,
}

func (v *InputTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InputTargetType(value)
	for _, existing := range AllowedInputTargetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InputTargetType", value)
}

// NewInputTargetTypeFromValue returns a pointer to a valid InputTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInputTargetTypeFromValue(v string) (*InputTargetType, error) {
	ev := InputTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InputTargetType: valid values are %v", v, AllowedInputTargetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InputTargetType) IsValid() bool {
	for _, existing := range AllowedInputTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InputTargetType value
func (v InputTargetType) Ptr() *InputTargetType {
	return &v
}

type NullableInputTargetType struct {
	value *InputTargetType
	isSet bool
}

func (v NullableInputTargetType) Get() *InputTargetType {
	return v.value
}

func (v *NullableInputTargetType) Set(val *InputTargetType) {
	v.value = val
	v.isSet = true
}

func (v NullableInputTargetType) IsSet() bool {
	return v.isSet
}

func (v *NullableInputTargetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputTargetType(val *InputTargetType) *NullableInputTargetType {
	return &NullableInputTargetType{value: val, isSet: true}
}

func (v NullableInputTargetType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInputTargetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
