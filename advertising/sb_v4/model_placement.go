package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// Placement List of bid adjustments for placements. - HOME - The home page of the Amazon store. - DETAIL_PAGE - Product detail pages within the Amazon store. - OTHER - Other placement groups. Such as search pages in the Amazon Store.
type Placement string

// List of Placement
const (
	PLACEMENT_HOME        Placement = "HOME"
	PLACEMENT_DETAIL_PAGE Placement = "DETAIL_PAGE"
	PLACEMENT_OTHER       Placement = "OTHER"
)

// All allowed values of Placement enum
var AllowedPlacementEnumValues = []Placement{
	"HOME",
	"DETAIL_PAGE",
	"OTHER",
}

func (v *Placement) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Placement(value)
	for _, existing := range AllowedPlacementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Placement", value)
}

// NewPlacementFromValue returns a pointer to a valid Placement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlacementFromValue(v string) (*Placement, error) {
	ev := Placement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Placement: valid values are %v", v, AllowedPlacementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Placement) IsValid() bool {
	for _, existing := range AllowedPlacementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Placement value
func (v Placement) Ptr() *Placement {
	return &v
}

type NullablePlacement struct {
	value *Placement
	isSet bool
}

func (v NullablePlacement) Get() *Placement {
	return v.value
}

func (v *NullablePlacement) Set(val *Placement) {
	v.value = val
	v.isSet = true
}

func (v NullablePlacement) IsSet() bool {
	return v.isSet
}

func (v *NullablePlacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlacement(val *Placement) *NullablePlacement {
	return &NullablePlacement{value: val, isSet: true}
}

func (v NullablePlacement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePlacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
