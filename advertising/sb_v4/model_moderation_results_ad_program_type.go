package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ModerationResultsAdProgramType The program type of the ad.
type ModerationResultsAdProgramType string

// List of ModerationResultsAdProgramType
const (
	MODERATIONRESULTSADPROGRAMTYPE_SB_PRODUCT_COLLECTION ModerationResultsAdProgramType = "SB_PRODUCT_COLLECTION"
	MODERATIONRESULTSADPROGRAMTYPE_SB_STORE_SPOTLIGHT    ModerationResultsAdProgramType = "SB_STORE_SPOTLIGHT"
	MODERATIONRESULTSADPROGRAMTYPE_SB_VIDEO              ModerationResultsAdProgramType = "SB_VIDEO"
	MODERATIONRESULTSADPROGRAMTYPE_SPONSORED_PRODUCTS    ModerationResultsAdProgramType = "SPONSORED_PRODUCTS"
)

// All allowed values of ModerationResultsAdProgramType enum
var AllowedModerationResultsAdProgramTypeEnumValues = []ModerationResultsAdProgramType{
	"SB_PRODUCT_COLLECTION",
	"SB_STORE_SPOTLIGHT",
	"SB_VIDEO",
	"SPONSORED_PRODUCTS",
}

func (v *ModerationResultsAdProgramType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModerationResultsAdProgramType(value)
	for _, existing := range AllowedModerationResultsAdProgramTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModerationResultsAdProgramType", value)
}

// NewModerationResultsAdProgramTypeFromValue returns a pointer to a valid ModerationResultsAdProgramType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModerationResultsAdProgramTypeFromValue(v string) (*ModerationResultsAdProgramType, error) {
	ev := ModerationResultsAdProgramType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModerationResultsAdProgramType: valid values are %v", v, AllowedModerationResultsAdProgramTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModerationResultsAdProgramType) IsValid() bool {
	for _, existing := range AllowedModerationResultsAdProgramTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModerationResultsAdProgramType value
func (v ModerationResultsAdProgramType) Ptr() *ModerationResultsAdProgramType {
	return &v
}

type NullableModerationResultsAdProgramType struct {
	value *ModerationResultsAdProgramType
	isSet bool
}

func (v NullableModerationResultsAdProgramType) Get() *ModerationResultsAdProgramType {
	return v.value
}

func (v *NullableModerationResultsAdProgramType) Set(val *ModerationResultsAdProgramType) {
	v.value = val
	v.isSet = true
}

func (v NullableModerationResultsAdProgramType) IsSet() bool {
	return v.isSet
}

func (v *NullableModerationResultsAdProgramType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModerationResultsAdProgramType(val *ModerationResultsAdProgramType) *NullableModerationResultsAdProgramType {
	return &NullableModerationResultsAdProgramType{value: val, isSet: true}
}

func (v NullableModerationResultsAdProgramType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableModerationResultsAdProgramType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
