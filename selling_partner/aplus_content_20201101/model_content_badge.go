package aplus_content_20201101

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ContentBadge A flag that provides additional information about an A+ Content document.
type ContentBadge string

// List of ContentBadge
const (
	CONTENTBADGE_BULK      ContentBadge = "BULK"
	CONTENTBADGE_GENERATED ContentBadge = "GENERATED"
	CONTENTBADGE_LAUNCHPAD ContentBadge = "LAUNCHPAD"
	CONTENTBADGE_PREMIUM   ContentBadge = "PREMIUM"
	CONTENTBADGE_STANDARD  ContentBadge = "STANDARD"
)

// All allowed values of ContentBadge enum
var AllowedContentBadgeEnumValues = []ContentBadge{
	CONTENTBADGE_BULK,
	CONTENTBADGE_GENERATED,
	CONTENTBADGE_LAUNCHPAD,
	CONTENTBADGE_PREMIUM,
	CONTENTBADGE_STANDARD,
}

func (v *ContentBadge) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContentBadge(value)
	for _, existing := range AllowedContentBadgeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContentBadge", value)
}

// NewContentBadgeFromValue returns a pointer to a valid ContentBadge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContentBadgeFromValue(v string) (*ContentBadge, error) {
	ev := ContentBadge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContentBadge: valid values are %v", v, AllowedContentBadgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContentBadge) IsValid() bool {
	for _, existing := range AllowedContentBadgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContentBadge value
func (v ContentBadge) Ptr() *ContentBadge {
	return &v
}

type NullableContentBadge struct {
	value *ContentBadge
	isSet bool
}

func (v NullableContentBadge) Get() *ContentBadge {
	return v.value
}

func (v *NullableContentBadge) Set(val *ContentBadge) {
	v.value = val
	v.isSet = true
}

func (v NullableContentBadge) IsSet() bool {
	return v.isSet
}

func (v *NullableContentBadge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentBadge(val *ContentBadge) *NullableContentBadge {
	return &NullableContentBadge{value: val, isSet: true}
}

func (v NullableContentBadge) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContentBadge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
