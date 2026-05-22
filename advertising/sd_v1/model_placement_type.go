package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PlacementType Placement type where the rule should be applied, defaults to `ALL`(including home page, detail page, twitch and offsite).  **Future** More available placemenTypes will be supported.
type PlacementType string

// List of PlacementType
const (
	PLACEMENTTYPE_ALL PlacementType = "ALL"
)

// All allowed values of PlacementType enum
var AllowedPlacementTypeEnumValues = []PlacementType{
	"ALL",
}

func (v *PlacementType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlacementType(value)
	for _, existing := range AllowedPlacementTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlacementType", value)
}

// NewPlacementTypeFromValue returns a pointer to a valid PlacementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlacementTypeFromValue(v string) (*PlacementType, error) {
	ev := PlacementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlacementType: valid values are %v", v, AllowedPlacementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlacementType) IsValid() bool {
	for _, existing := range AllowedPlacementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlacementType value
func (v PlacementType) Ptr() *PlacementType {
	return &v
}

type NullablePlacementType struct {
	value *PlacementType
	isSet bool
}

func (v NullablePlacementType) Get() *PlacementType {
	return v.value
}

func (v *NullablePlacementType) Set(val *PlacementType) {
	v.value = val
	v.isSet = true
}

func (v NullablePlacementType) IsSet() bool {
	return v.isSet
}

func (v *NullablePlacementType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlacementType(val *PlacementType) *NullablePlacementType {
	return &NullablePlacementType{value: val, isSet: true}
}

func (v NullablePlacementType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePlacementType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
