package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SDCostTypeV31 Determines what performance metric the bid recommendations will be optimized for. |Name|Description| |----|----------|-----------| |cpc|The bid recommendations will be optimized for clicks triggered by the ad.| |vcpm|The bid recommendations will be optimized for viewed impressions triggered by the ad. $1 is the minimum bid for vCPM.|
type SDCostTypeV31 string

// List of SDCostTypeV31
const (
	SDCOSTTYPEV31_CPC  SDCostTypeV31 = "cpc"
	SDCOSTTYPEV31_VCPM SDCostTypeV31 = "vcpm"
)

// All allowed values of SDCostTypeV31 enum
var AllowedSDCostTypeV31EnumValues = []SDCostTypeV31{
	"cpc",
	"vcpm",
}

func (v *SDCostTypeV31) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDCostTypeV31(value)
	for _, existing := range AllowedSDCostTypeV31EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDCostTypeV31", value)
}

// NewSDCostTypeV31FromValue returns a pointer to a valid SDCostTypeV31
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDCostTypeV31FromValue(v string) (*SDCostTypeV31, error) {
	ev := SDCostTypeV31(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDCostTypeV31: valid values are %v", v, AllowedSDCostTypeV31EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDCostTypeV31) IsValid() bool {
	for _, existing := range AllowedSDCostTypeV31EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDCostTypeV31 value
func (v SDCostTypeV31) Ptr() *SDCostTypeV31 {
	return &v
}

type NullableSDCostTypeV31 struct {
	value *SDCostTypeV31
	isSet bool
}

func (v NullableSDCostTypeV31) Get() *SDCostTypeV31 {
	return v.value
}

func (v *NullableSDCostTypeV31) Set(val *SDCostTypeV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDCostTypeV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDCostTypeV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDCostTypeV31(val *SDCostTypeV31) *NullableSDCostTypeV31 {
	return &NullableSDCostTypeV31{value: val, isSet: true}
}

func (v NullableSDCostTypeV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDCostTypeV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
