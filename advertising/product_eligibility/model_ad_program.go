package product_eligibility

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// AdProgram This defines the AdPrograms supported
type AdProgram string

// List of AdProgram
const (
	ADPROGRAM_SB   AdProgram = "SB"
	ADPROGRAM_SD   AdProgram = "SD"
	ADPROGRAM_MAAS AdProgram = "MAAS"
	ADPROGRAM_DTC  AdProgram = "DTC"
	ADPROGRAM_SPOT AdProgram = "SPOT"
)

// All allowed values of AdProgram enum
var AllowedAdProgramEnumValues = []AdProgram{
	"SB",
	"SD",
	"MAAS",
	"DTC",
	"SPOT",
}

func (v *AdProgram) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdProgram(value)
	for _, existing := range AllowedAdProgramEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdProgram", value)
}

// NewAdProgramFromValue returns a pointer to a valid AdProgram
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdProgramFromValue(v string) (*AdProgram, error) {
	ev := AdProgram(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdProgram: valid values are %v", v, AllowedAdProgramEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdProgram) IsValid() bool {
	for _, existing := range AllowedAdProgramEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdProgram value
func (v AdProgram) Ptr() *AdProgram {
	return &v
}

type NullableAdProgram struct {
	value *AdProgram
	isSet bool
}

func (v NullableAdProgram) Get() *AdProgram {
	return v.value
}

func (v *NullableAdProgram) Set(val *AdProgram) {
	v.value = val
	v.isSet = true
}

func (v NullableAdProgram) IsSet() bool {
	return v.isSet
}

func (v *NullableAdProgram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdProgram(val *AdProgram) *NullableAdProgram {
	return &NullableAdProgram{value: val, isSet: true}
}

func (v NullableAdProgram) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdProgram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
