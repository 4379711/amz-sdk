package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// QueryTermMatchType Defines how would the string resource field (e.g. campaign name, ad group name) be matched with the query term in filter.
type QueryTermMatchType string

// List of QueryTermMatchType
const (
	QUERYTERMMATCHTYPE_BROAD_MATCH QueryTermMatchType = "BROAD_MATCH"
	QUERYTERMMATCHTYPE_EXACT_MATCH QueryTermMatchType = "EXACT_MATCH"
)

// All allowed values of QueryTermMatchType enum
var AllowedQueryTermMatchTypeEnumValues = []QueryTermMatchType{
	"BROAD_MATCH",
	"EXACT_MATCH",
}

func (v *QueryTermMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryTermMatchType(value)
	for _, existing := range AllowedQueryTermMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryTermMatchType", value)
}

// NewQueryTermMatchTypeFromValue returns a pointer to a valid QueryTermMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryTermMatchTypeFromValue(v string) (*QueryTermMatchType, error) {
	ev := QueryTermMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryTermMatchType: valid values are %v", v, AllowedQueryTermMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryTermMatchType) IsValid() bool {
	for _, existing := range AllowedQueryTermMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTermMatchType value
func (v QueryTermMatchType) Ptr() *QueryTermMatchType {
	return &v
}

type NullableQueryTermMatchType struct {
	value *QueryTermMatchType
	isSet bool
}

func (v NullableQueryTermMatchType) Get() *QueryTermMatchType {
	return v.value
}

func (v *NullableQueryTermMatchType) Set(val *QueryTermMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryTermMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryTermMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryTermMatchType(val *QueryTermMatchType) *NullableQueryTermMatchType {
	return &NullableQueryTermMatchType{value: val, isSet: true}
}

func (v NullableQueryTermMatchType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableQueryTermMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
