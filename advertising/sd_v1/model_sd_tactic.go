package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SDTactic The advertising tactic associated with the campaign. The following table lists available tactic names: |Tactic Name|Type|Description|         |-----------|-----|-----------|         |T00020 &nbsp;    |Products&nbsp;| Products: Choose individual products to show your ads in placements related to those products.<br>Categories: Choose individual categories to show your ads in placements related to those categories.
type SDTactic string

// List of SDTactic
const (
	SDTACTIC_T00020 SDTactic = "T00020"
)

// All allowed values of SDTactic enum
var AllowedSDTacticEnumValues = []SDTactic{
	"T00020",
}

func (v *SDTactic) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDTactic(value)
	for _, existing := range AllowedSDTacticEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDTactic", value)
}

// NewSDTacticFromValue returns a pointer to a valid SDTactic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDTacticFromValue(v string) (*SDTactic, error) {
	ev := SDTactic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDTactic: valid values are %v", v, AllowedSDTacticEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDTactic) IsValid() bool {
	for _, existing := range AllowedSDTacticEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDTactic value
func (v SDTactic) Ptr() *SDTactic {
	return &v
}

type NullableSDTactic struct {
	value *SDTactic
	isSet bool
}

func (v NullableSDTactic) Get() *SDTactic {
	return v.value
}

func (v *NullableSDTactic) Set(val *SDTactic) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTactic) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTactic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTactic(val *SDTactic) *NullableSDTactic {
	return &NullableSDTactic{value: val, isSet: true}
}

func (v NullableSDTactic) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTactic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
