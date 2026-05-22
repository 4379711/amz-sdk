package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
	"gopkg.in/validator.v2"
)

// SDTargetExpressionV32 - struct for SDTargetExpressionV32
type SDTargetExpressionV32 struct {
	SDContentTargetingPredicateV31 *SDContentTargetingPredicateV31
	SDTargetingPredicateNestedV31  *SDTargetingPredicateNestedV31
	SDTargetingPredicateV31        *SDTargetingPredicateV31
}

// SDContentTargetingPredicateV31AsSDTargetExpressionV32 is a convenience function that returns SDContentTargetingPredicateV31 wrapped in SDTargetExpressionV32
func SDContentTargetingPredicateV31AsSDTargetExpressionV32(v *SDContentTargetingPredicateV31) SDTargetExpressionV32 {
	return SDTargetExpressionV32{
		SDContentTargetingPredicateV31: v,
	}
}

// SDTargetingPredicateNestedV31AsSDTargetExpressionV32 is a convenience function that returns SDTargetingPredicateNestedV31 wrapped in SDTargetExpressionV32
func SDTargetingPredicateNestedV31AsSDTargetExpressionV32(v *SDTargetingPredicateNestedV31) SDTargetExpressionV32 {
	return SDTargetExpressionV32{
		SDTargetingPredicateNestedV31: v,
	}
}

// SDTargetingPredicateV31AsSDTargetExpressionV32 is a convenience function that returns SDTargetingPredicateV31 wrapped in SDTargetExpressionV32
func SDTargetingPredicateV31AsSDTargetExpressionV32(v *SDTargetingPredicateV31) SDTargetExpressionV32 {
	return SDTargetExpressionV32{
		SDTargetingPredicateV31: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SDTargetExpressionV32) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SDContentTargetingPredicateV31
	err = newDecoder(data).Decode(&dst.SDContentTargetingPredicateV31)
	if err == nil {
		jsonSDContentTargetingPredicateV31, _ := sonic.Marshal(dst.SDContentTargetingPredicateV31)
		if string(jsonSDContentTargetingPredicateV31) == "{}" { // empty struct
			dst.SDContentTargetingPredicateV31 = nil
		} else {
			if err = validator.Validate(dst.SDContentTargetingPredicateV31); err != nil {
				dst.SDContentTargetingPredicateV31 = nil
			} else {
				match++
			}
		}
	} else {
		dst.SDContentTargetingPredicateV31 = nil
	}

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
		dst.SDContentTargetingPredicateV31 = nil
		dst.SDTargetingPredicateNestedV31 = nil
		dst.SDTargetingPredicateV31 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SDTargetExpressionV32)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SDTargetExpressionV32)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SDTargetExpressionV32) MarshalJSON() ([]byte, error) {
	if src.SDContentTargetingPredicateV31 != nil {
		return sonic.Marshal(&src.SDContentTargetingPredicateV31)
	}

	if src.SDTargetingPredicateNestedV31 != nil {
		return sonic.Marshal(&src.SDTargetingPredicateNestedV31)
	}

	if src.SDTargetingPredicateV31 != nil {
		return sonic.Marshal(&src.SDTargetingPredicateV31)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SDTargetExpressionV32) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SDContentTargetingPredicateV31 != nil {
		return obj.SDContentTargetingPredicateV31
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

type NullableSDTargetExpressionV32 struct {
	value *SDTargetExpressionV32
	isSet bool
}

func (v NullableSDTargetExpressionV32) Get() *SDTargetExpressionV32 {
	return v.value
}

func (v *NullableSDTargetExpressionV32) Set(val *SDTargetExpressionV32) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetExpressionV32) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetExpressionV32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetExpressionV32(val *SDTargetExpressionV32) *NullableSDTargetExpressionV32 {
	return &NullableSDTargetExpressionV32{value: val, isSet: true}
}

func (v NullableSDTargetExpressionV32) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetExpressionV32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
