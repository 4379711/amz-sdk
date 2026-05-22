package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// Tactic The advertising tactic associated with the campaign. The following table lists available tactic names: |Tactic Name|Type|Description| |-----------|-----|-----------| |T00020     |Contextual targeting | Choose individual products to show your ads in placements related to those products.<br> Choose individual categories to show your ads in placements related to those categories on and off Amazon.| |T00030     |Audiences or Contextual Targeting | Select individual products, categories, refined categories, or audiences to show your ads.|
type Tactic string

// List of Tactic
const (
	TACTIC_T00020 Tactic = "T00020"
	TACTIC_T00030 Tactic = "T00030"
)

// All allowed values of Tactic enum
var AllowedTacticEnumValues = []Tactic{
	"T00020",
	"T00030",
}

func (v *Tactic) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Tactic(value)
	for _, existing := range AllowedTacticEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Tactic", value)
}

// NewTacticFromValue returns a pointer to a valid Tactic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTacticFromValue(v string) (*Tactic, error) {
	ev := Tactic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Tactic: valid values are %v", v, AllowedTacticEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Tactic) IsValid() bool {
	for _, existing := range AllowedTacticEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Tactic value
func (v Tactic) Ptr() *Tactic {
	return &v
}

type NullableTactic struct {
	value *Tactic
	isSet bool
}

func (v NullableTactic) Get() *Tactic {
	return v.value
}

func (v *NullableTactic) Set(val *Tactic) {
	v.value = val
	v.isSet = true
}

func (v NullableTactic) IsSet() bool {
	return v.isSet
}

func (v *NullableTactic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTactic(val *Tactic) *NullableTactic {
	return &NullableTactic{value: val, isSet: true}
}

func (v NullableTactic) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTactic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
