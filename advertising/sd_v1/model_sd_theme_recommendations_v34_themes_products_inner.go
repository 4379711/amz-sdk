package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
	"gopkg.in/validator.v2"
)

// SDThemeRecommendationsV34ThemesProductsInner - struct for SDThemeRecommendationsV34ThemesProductsInner
type SDThemeRecommendationsV34ThemesProductsInner struct {
	SDProductTargetingRecommendationsSuccessV34 *SDProductTargetingRecommendationsSuccessV34
}

// SDProductTargetingRecommendationsSuccessV34AsSDThemeRecommendationsV34ThemesProductsInner is a convenience function that returns SDProductTargetingRecommendationsSuccessV34 wrapped in SDThemeRecommendationsV34ThemesProductsInner
func SDProductTargetingRecommendationsSuccessV34AsSDThemeRecommendationsV34ThemesProductsInner(v *SDProductTargetingRecommendationsSuccessV34) SDThemeRecommendationsV34ThemesProductsInner {
	return SDThemeRecommendationsV34ThemesProductsInner{
		SDProductTargetingRecommendationsSuccessV34: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SDThemeRecommendationsV34ThemesProductsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SDProductTargetingRecommendationsSuccessV34
	err = newDecoder(data).Decode(&dst.SDProductTargetingRecommendationsSuccessV34)
	if err == nil {
		jsonSDProductTargetingRecommendationsSuccessV34, _ := sonic.Marshal(dst.SDProductTargetingRecommendationsSuccessV34)
		if string(jsonSDProductTargetingRecommendationsSuccessV34) == "{}" { // empty struct
			dst.SDProductTargetingRecommendationsSuccessV34 = nil
		} else {
			if err = validator.Validate(dst.SDProductTargetingRecommendationsSuccessV34); err != nil {
				dst.SDProductTargetingRecommendationsSuccessV34 = nil
			} else {
				match++
			}
		}
	} else {
		dst.SDProductTargetingRecommendationsSuccessV34 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SDProductTargetingRecommendationsSuccessV34 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SDThemeRecommendationsV34ThemesProductsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SDThemeRecommendationsV34ThemesProductsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SDThemeRecommendationsV34ThemesProductsInner) MarshalJSON() ([]byte, error) {
	if src.SDProductTargetingRecommendationsSuccessV34 != nil {
		return sonic.Marshal(&src.SDProductTargetingRecommendationsSuccessV34)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SDThemeRecommendationsV34ThemesProductsInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SDProductTargetingRecommendationsSuccessV34 != nil {
		return obj.SDProductTargetingRecommendationsSuccessV34
	}

	// all schemas are nil
	return nil
}

type NullableSDThemeRecommendationsV34ThemesProductsInner struct {
	value *SDThemeRecommendationsV34ThemesProductsInner
	isSet bool
}

func (v NullableSDThemeRecommendationsV34ThemesProductsInner) Get() *SDThemeRecommendationsV34ThemesProductsInner {
	return v.value
}

func (v *NullableSDThemeRecommendationsV34ThemesProductsInner) Set(val *SDThemeRecommendationsV34ThemesProductsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSDThemeRecommendationsV34ThemesProductsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSDThemeRecommendationsV34ThemesProductsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDThemeRecommendationsV34ThemesProductsInner(val *SDThemeRecommendationsV34ThemesProductsInner) *NullableSDThemeRecommendationsV34ThemesProductsInner {
	return &NullableSDThemeRecommendationsV34ThemesProductsInner{value: val, isSet: true}
}

func (v NullableSDThemeRecommendationsV34ThemesProductsInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDThemeRecommendationsV34ThemesProductsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
