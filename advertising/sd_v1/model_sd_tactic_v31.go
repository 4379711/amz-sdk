package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SDTacticV31 The advertising tactic associated with the campaign. The following table lists available tactic names: |Tactic Name|Type|Description|         |-----------|-----|-----------|         |T00020 &nbsp;    |Products&nbsp;| Products: Choose individual products to show your ads in placements related to those products.<br>Categories: Choose individual categories to show your ads in placements related to those categories.|         |T00030&nbsp;|Audiences or Contextual Targeting &nbsp;|Select individual products, categories, refined categories, or audiences to show your ads.|
type SDTacticV31 string

// List of SDTacticV31
const (
	SDTACTICV31_T00020 SDTacticV31 = "T00020"
	SDTACTICV31_T00030 SDTacticV31 = "T00030"
)

// All allowed values of SDTacticV31 enum
var AllowedSDTacticV31EnumValues = []SDTacticV31{
	"T00020",
	"T00030",
}

func (v *SDTacticV31) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDTacticV31(value)
	for _, existing := range AllowedSDTacticV31EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDTacticV31", value)
}

// NewSDTacticV31FromValue returns a pointer to a valid SDTacticV31
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDTacticV31FromValue(v string) (*SDTacticV31, error) {
	ev := SDTacticV31(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDTacticV31: valid values are %v", v, AllowedSDTacticV31EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDTacticV31) IsValid() bool {
	for _, existing := range AllowedSDTacticV31EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDTacticV31 value
func (v SDTacticV31) Ptr() *SDTacticV31 {
	return &v
}

type NullableSDTacticV31 struct {
	value *SDTacticV31
	isSet bool
}

func (v NullableSDTacticV31) Get() *SDTacticV31 {
	return v.value
}

func (v *NullableSDTacticV31) Set(val *SDTacticV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTacticV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTacticV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTacticV31(val *SDTacticV31) *NullableSDTacticV31 {
	return &NullableSDTacticV31{value: val, isSet: true}
}

func (v NullableSDTacticV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTacticV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
