package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// MultiCountryTargetingExpressionType the model 'MultiCountryTargetingExpressionType'
type MultiCountryTargetingExpressionType string

// List of MultiCountryTargetingExpressionType
const (
	MULTICOUNTRYTARGETINGEXPRESSIONTYPE_CLOSE_MATCH             MultiCountryTargetingExpressionType = "CLOSE_MATCH"
	MULTICOUNTRYTARGETINGEXPRESSIONTYPE_LOOSE_MATCH             MultiCountryTargetingExpressionType = "LOOSE_MATCH"
	MULTICOUNTRYTARGETINGEXPRESSIONTYPE_SUBSTITUTES             MultiCountryTargetingExpressionType = "SUBSTITUTES"
	MULTICOUNTRYTARGETINGEXPRESSIONTYPE_COMPLEMENTS             MultiCountryTargetingExpressionType = "COMPLEMENTS"
	MULTICOUNTRYTARGETINGEXPRESSIONTYPE_KEYWORD_BROAD_MATCH     MultiCountryTargetingExpressionType = "KEYWORD_BROAD_MATCH"
	MULTICOUNTRYTARGETINGEXPRESSIONTYPE_KEYWORD_EXACT_MATCH     MultiCountryTargetingExpressionType = "KEYWORD_EXACT_MATCH"
	MULTICOUNTRYTARGETINGEXPRESSIONTYPE_KEYWORD_PHRASE_MATCH    MultiCountryTargetingExpressionType = "KEYWORD_PHRASE_MATCH"
	MULTICOUNTRYTARGETINGEXPRESSIONTYPE_PAT_ASIN                MultiCountryTargetingExpressionType = "PAT_ASIN"
	MULTICOUNTRYTARGETINGEXPRESSIONTYPE_PAT_CATEGORY            MultiCountryTargetingExpressionType = "PAT_CATEGORY"
	MULTICOUNTRYTARGETINGEXPRESSIONTYPE_PAT_CATEGORY_REFINEMENT MultiCountryTargetingExpressionType = "PAT_CATEGORY_REFINEMENT"
	MULTICOUNTRYTARGETINGEXPRESSIONTYPE_KEYWORD_GROUP           MultiCountryTargetingExpressionType = "KEYWORD_GROUP"
)

// All allowed values of MultiCountryTargetingExpressionType enum
var AllowedMultiCountryTargetingExpressionTypeEnumValues = []MultiCountryTargetingExpressionType{
	"CLOSE_MATCH",
	"LOOSE_MATCH",
	"SUBSTITUTES",
	"COMPLEMENTS",
	"KEYWORD_BROAD_MATCH",
	"KEYWORD_EXACT_MATCH",
	"KEYWORD_PHRASE_MATCH",
	"PAT_ASIN",
	"PAT_CATEGORY",
	"PAT_CATEGORY_REFINEMENT",
	"KEYWORD_GROUP",
}

func (v *MultiCountryTargetingExpressionType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MultiCountryTargetingExpressionType(value)
	for _, existing := range AllowedMultiCountryTargetingExpressionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MultiCountryTargetingExpressionType", value)
}

// NewMultiCountryTargetingExpressionTypeFromValue returns a pointer to a valid MultiCountryTargetingExpressionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMultiCountryTargetingExpressionTypeFromValue(v string) (*MultiCountryTargetingExpressionType, error) {
	ev := MultiCountryTargetingExpressionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MultiCountryTargetingExpressionType: valid values are %v", v, AllowedMultiCountryTargetingExpressionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MultiCountryTargetingExpressionType) IsValid() bool {
	for _, existing := range AllowedMultiCountryTargetingExpressionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MultiCountryTargetingExpressionType value
func (v MultiCountryTargetingExpressionType) Ptr() *MultiCountryTargetingExpressionType {
	return &v
}

type NullableMultiCountryTargetingExpressionType struct {
	value *MultiCountryTargetingExpressionType
	isSet bool
}

func (v NullableMultiCountryTargetingExpressionType) Get() *MultiCountryTargetingExpressionType {
	return v.value
}

func (v *NullableMultiCountryTargetingExpressionType) Set(val *MultiCountryTargetingExpressionType) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiCountryTargetingExpressionType) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiCountryTargetingExpressionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiCountryTargetingExpressionType(val *MultiCountryTargetingExpressionType) *NullableMultiCountryTargetingExpressionType {
	return &NullableMultiCountryTargetingExpressionType{value: val, isSet: true}
}

func (v NullableMultiCountryTargetingExpressionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMultiCountryTargetingExpressionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
