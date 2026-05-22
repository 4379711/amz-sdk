package fulfillment_outbound_20200701

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// InvalidItemReasonCode A code for why the item is invalid for return.
type InvalidItemReasonCode string

// List of InvalidItemReasonCode
const (
	INVALIDITEMREASONCODE_INVALID_VALUES          InvalidItemReasonCode = "InvalidValues"
	INVALIDITEMREASONCODE_DUPLICATE_REQUEST       InvalidItemReasonCode = "DuplicateRequest"
	INVALIDITEMREASONCODE_NO_COMPLETED_SHIP_ITEMS InvalidItemReasonCode = "NoCompletedShipItems"
	INVALIDITEMREASONCODE_NO_RETURNABLE_QUANTITY  InvalidItemReasonCode = "NoReturnableQuantity"
)

// All allowed values of InvalidItemReasonCode enum
var AllowedInvalidItemReasonCodeEnumValues = []InvalidItemReasonCode{
	INVALIDITEMREASONCODE_INVALID_VALUES,
	INVALIDITEMREASONCODE_DUPLICATE_REQUEST,
	INVALIDITEMREASONCODE_NO_COMPLETED_SHIP_ITEMS,
	INVALIDITEMREASONCODE_NO_RETURNABLE_QUANTITY,
}

func (v *InvalidItemReasonCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvalidItemReasonCode(value)
	for _, existing := range AllowedInvalidItemReasonCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InvalidItemReasonCode", value)
}

// NewInvalidItemReasonCodeFromValue returns a pointer to a valid InvalidItemReasonCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvalidItemReasonCodeFromValue(v string) (*InvalidItemReasonCode, error) {
	ev := InvalidItemReasonCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvalidItemReasonCode: valid values are %v", v, AllowedInvalidItemReasonCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvalidItemReasonCode) IsValid() bool {
	for _, existing := range AllowedInvalidItemReasonCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InvalidItemReasonCode value
func (v InvalidItemReasonCode) Ptr() *InvalidItemReasonCode {
	return &v
}

type NullableInvalidItemReasonCode struct {
	value *InvalidItemReasonCode
	isSet bool
}

func (v NullableInvalidItemReasonCode) Get() *InvalidItemReasonCode {
	return v.value
}

func (v *NullableInvalidItemReasonCode) Set(val *InvalidItemReasonCode) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidItemReasonCode) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidItemReasonCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidItemReasonCode(val *InvalidItemReasonCode) *NullableInvalidItemReasonCode {
	return &NullableInvalidItemReasonCode{value: val, isSet: true}
}

func (v NullableInvalidItemReasonCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvalidItemReasonCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
