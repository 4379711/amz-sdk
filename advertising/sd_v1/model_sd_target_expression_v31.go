package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
	"gopkg.in/validator.v2"
)

// SDTargetExpressionV31 - struct for SDTargetExpressionV31
type SDTargetExpressionV31 struct {
	SDTargetingPredicateNestedV31 *SDTargetingPredicateNestedV31
	SDTargetingPredicateV31       *SDTargetingPredicateV31
}

// SDTargetingPredicateNestedV31AsSDTargetExpressionV31 is a convenience function that returns SDTargetingPredicateNestedV31 wrapped in SDTargetExpressionV31
func SDTargetingPredicateNestedV31AsSDTargetExpressionV31(v *SDTargetingPredicateNestedV31) SDTargetExpressionV31 {
	return SDTargetExpressionV31{
		SDTargetingPredicateNestedV31: v,
	}
}

// SDTargetingPredicateV31AsSDTargetExpressionV31 is a convenience function that returns SDTargetingPredicateV31 wrapped in SDTargetExpressionV31
func SDTargetingPredicateV31AsSDTargetExpressionV31(v *SDTargetingPredicateV31) SDTargetExpressionV31 {
	return SDTargetExpressionV31{
		SDTargetingPredicateV31: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SDTargetExpressionV31) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SDTargetingPredicateNestedV31
	err = newDecoder(data).Decode(&dst.SDTargetingPredicateNestedV31)
	if err == nil {
		jsonSDTargetingPredicateNestedV31, _ := sonic.Marshal(dst.SDTargetingPredicateNestedV31)
		if string(jsonSDTargetingPredicateNestedV31) == "{}" { // empty struct
			dst.SDTargetingPredicateNestedV31 = nil
		} else {
			if err = validator.Validate(dst.SDTargetingPredicateNestedV31); err != nil {
				dst.SDTargetingPredicateNestedV31 = nil
			} else {
				match++
			}
		}
	} else {
		dst.SDTargetingPredicateNestedV31 = nil
	}

	// try to unmarshal data into SDTargetingPredicateV31
	err = newDecoder(data).Decode(&dst.SDTargetingPredicateV31)
	if err == nil {
		jsonSDTargetingPredicateV31, _ := sonic.Marshal(dst.SDTargetingPredicateV31)
		if string(jsonSDTargetingPredicateV31) == "{}" { // empty struct
			dst.SDTargetingPredicateV31 = nil
		} else {
			if err = validator.Validate(dst.SDTargetingPredicateV31); err != nil {
				dst.SDTargetingPredicateV31 = nil
			} else {
				match++
			}
		}
	} else {
		dst.SDTargetingPredicateV31 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SDTargetingPredicateNestedV31 = nil
		dst.SDTargetingPredicateV31 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SDTargetExpressionV31)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SDTargetExpressionV31)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SDTargetExpressionV31) MarshalJSON() ([]byte, error) {
	if src.SDTargetingPredicateNestedV31 != nil {
		return sonic.Marshal(&src.SDTargetingPredicateNestedV31)
	}

	if src.SDTargetingPredicateV31 != nil {
		return sonic.Marshal(&src.SDTargetingPredicateV31)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SDTargetExpressionV31) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SDTargetingPredicateNestedV31 != nil {
		return obj.SDTargetingPredicateNestedV31
	}

	if obj.SDTargetingPredicateV31 != nil {
		return obj.SDTargetingPredicateV31
	}

	// all schemas are nil
	return nil
}

type NullableSDTargetExpressionV31 struct {
	value *SDTargetExpressionV31
	isSet bool
}

func (v NullableSDTargetExpressionV31) Get() *SDTargetExpressionV31 {
	return v.value
}

func (v *NullableSDTargetExpressionV31) Set(val *SDTargetExpressionV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetExpressionV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetExpressionV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetExpressionV31(val *SDTargetExpressionV31) *NullableSDTargetExpressionV31 {
	return &NullableSDTargetExpressionV31{value: val, isSet: true}
}

func (v NullableSDTargetExpressionV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetExpressionV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
