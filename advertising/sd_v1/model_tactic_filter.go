package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// TacticFilter Optional. Restricts results to entities with the advertising tactic associated with the campaign. Must be one of the following table lists available tactic names: |Tactic Name|Type|Description| |-----------|-----|-----------| |T00020     |Contextual targeting | Choose individual products to show your ads in placements related to those products.<br> Choose individual categories to show your ads in placements related to those categories on and off Amazon.| |T00030     |Audiences or Contextual Targeting | Select individual products, categories, refined categories, or audiences to show your ads.|
type TacticFilter string

// List of TacticFilter
const (
	TACTICFILTER_T00020       TacticFilter = "T00020"
	TACTICFILTER_T00030       TacticFilter = "T00030"
	TACTICFILTER_T00020T00030 TacticFilter = "T00020,T00030"
)

// All allowed values of TacticFilter enum
var AllowedTacticFilterEnumValues = []TacticFilter{
	"T00020",
	"T00030",
	"T00020,T00030",
}

func (v *TacticFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TacticFilter(value)
	for _, existing := range AllowedTacticFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TacticFilter", value)
}

// NewTacticFilterFromValue returns a pointer to a valid TacticFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTacticFilterFromValue(v string) (*TacticFilter, error) {
	ev := TacticFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TacticFilter: valid values are %v", v, AllowedTacticFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TacticFilter) IsValid() bool {
	for _, existing := range AllowedTacticFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TacticFilter value
func (v TacticFilter) Ptr() *TacticFilter {
	return &v
}

type NullableTacticFilter struct {
	value *TacticFilter
	isSet bool
}

func (v NullableTacticFilter) Get() *TacticFilter {
	return v.value
}

func (v *NullableTacticFilter) Set(val *TacticFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTacticFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTacticFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTacticFilter(val *TacticFilter) *NullableTacticFilter {
	return &NullableTacticFilter{value: val, isSet: true}
}

func (v NullableTacticFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTacticFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
