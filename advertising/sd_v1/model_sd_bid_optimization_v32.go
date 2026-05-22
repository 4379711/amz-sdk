package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SDBidOptimizationV32 Determines what the recommended bids will be optimized for.   |Name|CostType|Description| |----|--------|-----------| |reach|vcpm|Optimize for viewable impressions. $1 is the minimum bid for vCPM.| |clicks|cpc|Optimize for page visits| |conversions|cpc|Optimize for conversion|
type SDBidOptimizationV32 string

// List of SDBidOptimizationV32
const (
	SDBIDOPTIMIZATIONV32_REACH       SDBidOptimizationV32 = "reach"
	SDBIDOPTIMIZATIONV32_CLICKS      SDBidOptimizationV32 = "clicks"
	SDBIDOPTIMIZATIONV32_CONVERSIONS SDBidOptimizationV32 = "conversions"
)

// All allowed values of SDBidOptimizationV32 enum
var AllowedSDBidOptimizationV32EnumValues = []SDBidOptimizationV32{
	"reach",
	"clicks",
	"conversions",
}

func (v *SDBidOptimizationV32) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDBidOptimizationV32(value)
	for _, existing := range AllowedSDBidOptimizationV32EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDBidOptimizationV32", value)
}

// NewSDBidOptimizationV32FromValue returns a pointer to a valid SDBidOptimizationV32
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDBidOptimizationV32FromValue(v string) (*SDBidOptimizationV32, error) {
	ev := SDBidOptimizationV32(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDBidOptimizationV32: valid values are %v", v, AllowedSDBidOptimizationV32EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDBidOptimizationV32) IsValid() bool {
	for _, existing := range AllowedSDBidOptimizationV32EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDBidOptimizationV32 value
func (v SDBidOptimizationV32) Ptr() *SDBidOptimizationV32 {
	return &v
}

type NullableSDBidOptimizationV32 struct {
	value *SDBidOptimizationV32
	isSet bool
}

func (v NullableSDBidOptimizationV32) Get() *SDBidOptimizationV32 {
	return v.value
}

func (v *NullableSDBidOptimizationV32) Set(val *SDBidOptimizationV32) {
	v.value = val
	v.isSet = true
}

func (v NullableSDBidOptimizationV32) IsSet() bool {
	return v.isSet
}

func (v *NullableSDBidOptimizationV32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDBidOptimizationV32(val *SDBidOptimizationV32) *NullableSDBidOptimizationV32 {
	return &NullableSDBidOptimizationV32{value: val, isSet: true}
}

func (v NullableSDBidOptimizationV32) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDBidOptimizationV32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
