package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SBTargetingSupplySource [UPDATE: As of 05/28/2024, `STREAMING_VIDEO` is deprecated].   The supply source where the target will be used. Use `AMAZON` for placements on Amazon website. Use `STREAMING_VIDEO` for off-site video placements such as IMDb TV.
type SBTargetingSupplySource string

// List of SBTargetingSupplySource
const (
	SBTARGETINGSUPPLYSOURCE_AMAZON          SBTargetingSupplySource = "AMAZON"
	SBTARGETINGSUPPLYSOURCE_STREAMING_VIDEO SBTargetingSupplySource = "STREAMING_VIDEO"
)

// All allowed values of SBTargetingSupplySource enum
var AllowedSBTargetingSupplySourceEnumValues = []SBTargetingSupplySource{
	"AMAZON",
	"STREAMING_VIDEO",
}

func (v *SBTargetingSupplySource) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SBTargetingSupplySource(value)
	for _, existing := range AllowedSBTargetingSupplySourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SBTargetingSupplySource", value)
}

// NewSBTargetingSupplySourceFromValue returns a pointer to a valid SBTargetingSupplySource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSBTargetingSupplySourceFromValue(v string) (*SBTargetingSupplySource, error) {
	ev := SBTargetingSupplySource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SBTargetingSupplySource: valid values are %v", v, AllowedSBTargetingSupplySourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SBTargetingSupplySource) IsValid() bool {
	for _, existing := range AllowedSBTargetingSupplySourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SBTargetingSupplySource value
func (v SBTargetingSupplySource) Ptr() *SBTargetingSupplySource {
	return &v
}

type NullableSBTargetingSupplySource struct {
	value *SBTargetingSupplySource
	isSet bool
}

func (v NullableSBTargetingSupplySource) Get() *SBTargetingSupplySource {
	return v.value
}

func (v *NullableSBTargetingSupplySource) Set(val *SBTargetingSupplySource) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingSupplySource) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingSupplySource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingSupplySource(val *SBTargetingSupplySource) *NullableSBTargetingSupplySource {
	return &NullableSBTargetingSupplySource{value: val, isSet: true}
}

func (v NullableSBTargetingSupplySource) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingSupplySource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
