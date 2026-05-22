package shipping_v2

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// IneligibilityReasonCode Reasons that make a shipment service offering ineligible.
type IneligibilityReasonCode string

// List of IneligibilityReasonCode
const (
	INELIGIBILITYREASONCODE_NO_COVERAGE                       IneligibilityReasonCode = "NO_COVERAGE"
	INELIGIBILITYREASONCODE_PICKUP_SLOT_RESTRICTION           IneligibilityReasonCode = "PICKUP_SLOT_RESTRICTION"
	INELIGIBILITYREASONCODE_UNSUPPORTED_VAS                   IneligibilityReasonCode = "UNSUPPORTED_VAS"
	INELIGIBILITYREASONCODE_VAS_COMBINATION_RESTRICTION       IneligibilityReasonCode = "VAS_COMBINATION_RESTRICTION"
	INELIGIBILITYREASONCODE_SIZE_RESTRICTIONS                 IneligibilityReasonCode = "SIZE_RESTRICTIONS"
	INELIGIBILITYREASONCODE_WEIGHT_RESTRICTIONS               IneligibilityReasonCode = "WEIGHT_RESTRICTIONS"
	INELIGIBILITYREASONCODE_LATE_DELIVERY                     IneligibilityReasonCode = "LATE_DELIVERY"
	INELIGIBILITYREASONCODE_PROGRAM_CONSTRAINTS               IneligibilityReasonCode = "PROGRAM_CONSTRAINTS"
	INELIGIBILITYREASONCODE_TERMS_AND_CONDITIONS_NOT_ACCEPTED IneligibilityReasonCode = "TERMS_AND_CONDITIONS_NOT_ACCEPTED"
	INELIGIBILITYREASONCODE_UNKNOWN                           IneligibilityReasonCode = "UNKNOWN"
)

// All allowed values of IneligibilityReasonCode enum
var AllowedIneligibilityReasonCodeEnumValues = []IneligibilityReasonCode{
	INELIGIBILITYREASONCODE_NO_COVERAGE,
	INELIGIBILITYREASONCODE_PICKUP_SLOT_RESTRICTION,
	INELIGIBILITYREASONCODE_UNSUPPORTED_VAS,
	INELIGIBILITYREASONCODE_VAS_COMBINATION_RESTRICTION,
	INELIGIBILITYREASONCODE_SIZE_RESTRICTIONS,
	INELIGIBILITYREASONCODE_WEIGHT_RESTRICTIONS,
	INELIGIBILITYREASONCODE_LATE_DELIVERY,
	INELIGIBILITYREASONCODE_PROGRAM_CONSTRAINTS,
	INELIGIBILITYREASONCODE_TERMS_AND_CONDITIONS_NOT_ACCEPTED,
	INELIGIBILITYREASONCODE_UNKNOWN,
}

func (v *IneligibilityReasonCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IneligibilityReasonCode(value)
	for _, existing := range AllowedIneligibilityReasonCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IneligibilityReasonCode", value)
}

// NewIneligibilityReasonCodeFromValue returns a pointer to a valid IneligibilityReasonCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIneligibilityReasonCodeFromValue(v string) (*IneligibilityReasonCode, error) {
	ev := IneligibilityReasonCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IneligibilityReasonCode: valid values are %v", v, AllowedIneligibilityReasonCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IneligibilityReasonCode) IsValid() bool {
	for _, existing := range AllowedIneligibilityReasonCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IneligibilityReasonCode value
func (v IneligibilityReasonCode) Ptr() *IneligibilityReasonCode {
	return &v
}

type NullableIneligibilityReasonCode struct {
	value *IneligibilityReasonCode
	isSet bool
}

func (v NullableIneligibilityReasonCode) Get() *IneligibilityReasonCode {
	return v.value
}

func (v *NullableIneligibilityReasonCode) Set(val *IneligibilityReasonCode) {
	v.value = val
	v.isSet = true
}

func (v NullableIneligibilityReasonCode) IsSet() bool {
	return v.isSet
}

func (v *NullableIneligibilityReasonCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIneligibilityReasonCode(val *IneligibilityReasonCode) *NullableIneligibilityReasonCode {
	return &NullableIneligibilityReasonCode{value: val, isSet: true}
}

func (v NullableIneligibilityReasonCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIneligibilityReasonCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
