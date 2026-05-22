package aplus_content_20201101

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ContentModuleType The type of A+ Content module.
type ContentModuleType string

// List of ContentModuleType
const (
	CONTENTMODULETYPE_STANDARD_COMPANY_LOGO              ContentModuleType = "STANDARD_COMPANY_LOGO"
	CONTENTMODULETYPE_STANDARD_COMPARISON_TABLE          ContentModuleType = "STANDARD_COMPARISON_TABLE"
	CONTENTMODULETYPE_STANDARD_FOUR_IMAGE_TEXT           ContentModuleType = "STANDARD_FOUR_IMAGE_TEXT"
	CONTENTMODULETYPE_STANDARD_FOUR_IMAGE_TEXT_QUADRANT  ContentModuleType = "STANDARD_FOUR_IMAGE_TEXT_QUADRANT"
	CONTENTMODULETYPE_STANDARD_HEADER_IMAGE_TEXT         ContentModuleType = "STANDARD_HEADER_IMAGE_TEXT"
	CONTENTMODULETYPE_STANDARD_IMAGE_SIDEBAR             ContentModuleType = "STANDARD_IMAGE_SIDEBAR"
	CONTENTMODULETYPE_STANDARD_IMAGE_TEXT_OVERLAY        ContentModuleType = "STANDARD_IMAGE_TEXT_OVERLAY"
	CONTENTMODULETYPE_STANDARD_MULTIPLE_IMAGE_TEXT       ContentModuleType = "STANDARD_MULTIPLE_IMAGE_TEXT"
	CONTENTMODULETYPE_STANDARD_PRODUCT_DESCRIPTION       ContentModuleType = "STANDARD_PRODUCT_DESCRIPTION"
	CONTENTMODULETYPE_STANDARD_SINGLE_IMAGE_HIGHLIGHTS   ContentModuleType = "STANDARD_SINGLE_IMAGE_HIGHLIGHTS"
	CONTENTMODULETYPE_STANDARD_SINGLE_IMAGE_SPECS_DETAIL ContentModuleType = "STANDARD_SINGLE_IMAGE_SPECS_DETAIL"
	CONTENTMODULETYPE_STANDARD_SINGLE_SIDE_IMAGE         ContentModuleType = "STANDARD_SINGLE_SIDE_IMAGE"
	CONTENTMODULETYPE_STANDARD_TECH_SPECS                ContentModuleType = "STANDARD_TECH_SPECS"
	CONTENTMODULETYPE_STANDARD_TEXT                      ContentModuleType = "STANDARD_TEXT"
	CONTENTMODULETYPE_STANDARD_THREE_IMAGE_TEXT          ContentModuleType = "STANDARD_THREE_IMAGE_TEXT"
)

// All allowed values of ContentModuleType enum
var AllowedContentModuleTypeEnumValues = []ContentModuleType{
	CONTENTMODULETYPE_STANDARD_COMPANY_LOGO,
	CONTENTMODULETYPE_STANDARD_COMPARISON_TABLE,
	CONTENTMODULETYPE_STANDARD_FOUR_IMAGE_TEXT,
	CONTENTMODULETYPE_STANDARD_FOUR_IMAGE_TEXT_QUADRANT,
	CONTENTMODULETYPE_STANDARD_HEADER_IMAGE_TEXT,
	CONTENTMODULETYPE_STANDARD_IMAGE_SIDEBAR,
	CONTENTMODULETYPE_STANDARD_IMAGE_TEXT_OVERLAY,
	CONTENTMODULETYPE_STANDARD_MULTIPLE_IMAGE_TEXT,
	CONTENTMODULETYPE_STANDARD_PRODUCT_DESCRIPTION,
	CONTENTMODULETYPE_STANDARD_SINGLE_IMAGE_HIGHLIGHTS,
	CONTENTMODULETYPE_STANDARD_SINGLE_IMAGE_SPECS_DETAIL,
	CONTENTMODULETYPE_STANDARD_SINGLE_SIDE_IMAGE,
	CONTENTMODULETYPE_STANDARD_TECH_SPECS,
	CONTENTMODULETYPE_STANDARD_TEXT,
	CONTENTMODULETYPE_STANDARD_THREE_IMAGE_TEXT,
}

func (v *ContentModuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContentModuleType(value)
	for _, existing := range AllowedContentModuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContentModuleType", value)
}

// NewContentModuleTypeFromValue returns a pointer to a valid ContentModuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContentModuleTypeFromValue(v string) (*ContentModuleType, error) {
	ev := ContentModuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContentModuleType: valid values are %v", v, AllowedContentModuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContentModuleType) IsValid() bool {
	for _, existing := range AllowedContentModuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContentModuleType value
func (v ContentModuleType) Ptr() *ContentModuleType {
	return &v
}

type NullableContentModuleType struct {
	value *ContentModuleType
	isSet bool
}

func (v NullableContentModuleType) Get() *ContentModuleType {
	return v.value
}

func (v *NullableContentModuleType) Set(val *ContentModuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableContentModuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableContentModuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentModuleType(val *ContentModuleType) *NullableContentModuleType {
	return &NullableContentModuleType{value: val, isSet: true}
}

func (v NullableContentModuleType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContentModuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
