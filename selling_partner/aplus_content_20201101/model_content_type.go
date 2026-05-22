package aplus_content_20201101

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ContentType The A+ Content document type.
type ContentType string

// List of ContentType
const (
	CONTENTTYPE_EBC ContentType = "EBC"
	CONTENTTYPE_EMC ContentType = "EMC"
)

// All allowed values of ContentType enum
var AllowedContentTypeEnumValues = []ContentType{
	CONTENTTYPE_EBC,
	CONTENTTYPE_EMC,
}

func (v *ContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContentType(value)
	for _, existing := range AllowedContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContentType", value)
}

// NewContentTypeFromValue returns a pointer to a valid ContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContentTypeFromValue(v string) (*ContentType, error) {
	ev := ContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContentType: valid values are %v", v, AllowedContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContentType) IsValid() bool {
	for _, existing := range AllowedContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContentType value
func (v ContentType) Ptr() *ContentType {
	return &v
}

type NullableContentType struct {
	value *ContentType
	isSet bool
}

func (v NullableContentType) Get() *ContentType {
	return v.value
}

func (v *NullableContentType) Set(val *ContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentType(val *ContentType) *NullableContentType {
	return &NullableContentType{value: val, isSet: true}
}

func (v NullableContentType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
