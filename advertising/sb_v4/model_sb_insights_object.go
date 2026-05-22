package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
	"gopkg.in/validator.v2"
)

// SBInsightsObject - struct for SBInsightsObject
type SBInsightsObject struct {
	KeywordInsight *KeywordInsight
}

// KeywordInsightAsSBInsightsObject is a convenience function that returns KeywordInsight wrapped in SBInsightsObject
func KeywordInsightAsSBInsightsObject(v *KeywordInsight) SBInsightsObject {
	return SBInsightsObject{
		KeywordInsight: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SBInsightsObject) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into KeywordInsight
	err = newDecoder(data).Decode(&dst.KeywordInsight)
	if err == nil {
		jsonKeywordInsight, _ := sonic.Marshal(dst.KeywordInsight)
		if string(jsonKeywordInsight) == "{}" { // empty struct
			dst.KeywordInsight = nil
		} else {
			if err = validator.Validate(dst.KeywordInsight); err != nil {
				dst.KeywordInsight = nil
			} else {
				match++
			}
		}
	} else {
		dst.KeywordInsight = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.KeywordInsight = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SBInsightsObject)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SBInsightsObject)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SBInsightsObject) MarshalJSON() ([]byte, error) {
	if src.KeywordInsight != nil {
		return sonic.Marshal(&src.KeywordInsight)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SBInsightsObject) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.KeywordInsight != nil {
		return obj.KeywordInsight
	}

	// all schemas are nil
	return nil
}

type NullableSBInsightsObject struct {
	value *SBInsightsObject
	isSet bool
}

func (v NullableSBInsightsObject) Get() *SBInsightsObject {
	return v.value
}

func (v *NullableSBInsightsObject) Set(val *SBInsightsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsObject(val *SBInsightsObject) *NullableSBInsightsObject {
	return &NullableSBInsightsObject{value: val, isSet: true}
}

func (v NullableSBInsightsObject) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
