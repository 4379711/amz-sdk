package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
	"gopkg.in/validator.v2"
)

// SDThemeRecommendationsThemesProductsInner - struct for SDThemeRecommendationsThemesProductsInner
type SDThemeRecommendationsThemesProductsInner struct {
	SDProductTargetingRecommendationsSuccess *SDProductTargetingRecommendationsSuccess
}

// SDProductTargetingRecommendationsSuccessAsSDThemeRecommendationsThemesProductsInner is a convenience function that returns SDProductTargetingRecommendationsSuccess wrapped in SDThemeRecommendationsThemesProductsInner
func SDProductTargetingRecommendationsSuccessAsSDThemeRecommendationsThemesProductsInner(v *SDProductTargetingRecommendationsSuccess) SDThemeRecommendationsThemesProductsInner {
	return SDThemeRecommendationsThemesProductsInner{
		SDProductTargetingRecommendationsSuccess: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SDThemeRecommendationsThemesProductsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SDProductTargetingRecommendationsSuccess
	err = newDecoder(data).Decode(&dst.SDProductTargetingRecommendationsSuccess)
	if err == nil {
		jsonSDProductTargetingRecommendationsSuccess, _ := sonic.Marshal(dst.SDProductTargetingRecommendationsSuccess)
		if string(jsonSDProductTargetingRecommendationsSuccess) == "{}" { // empty struct
			dst.SDProductTargetingRecommendationsSuccess = nil
		} else {
			if err = validator.Validate(dst.SDProductTargetingRecommendationsSuccess); err != nil {
				dst.SDProductTargetingRecommendationsSuccess = nil
			} else {
				match++
			}
		}
	} else {
		dst.SDProductTargetingRecommendationsSuccess = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SDProductTargetingRecommendationsSuccess = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SDThemeRecommendationsThemesProductsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SDThemeRecommendationsThemesProductsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SDThemeRecommendationsThemesProductsInner) MarshalJSON() ([]byte, error) {
	if src.SDProductTargetingRecommendationsSuccess != nil {
		return sonic.Marshal(&src.SDProductTargetingRecommendationsSuccess)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SDThemeRecommendationsThemesProductsInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SDProductTargetingRecommendationsSuccess != nil {
		return obj.SDProductTargetingRecommendationsSuccess
	}

	// all schemas are nil
	return nil
}

type NullableSDThemeRecommendationsThemesProductsInner struct {
	value *SDThemeRecommendationsThemesProductsInner
	isSet bool
}

func (v NullableSDThemeRecommendationsThemesProductsInner) Get() *SDThemeRecommendationsThemesProductsInner {
	return v.value
}

func (v *NullableSDThemeRecommendationsThemesProductsInner) Set(val *SDThemeRecommendationsThemesProductsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSDThemeRecommendationsThemesProductsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSDThemeRecommendationsThemesProductsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDThemeRecommendationsThemesProductsInner(val *SDThemeRecommendationsThemesProductsInner) *NullableSDThemeRecommendationsThemesProductsInner {
	return &NullableSDThemeRecommendationsThemesProductsInner{value: val, isSet: true}
}

func (v NullableSDThemeRecommendationsThemesProductsInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDThemeRecommendationsThemesProductsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
