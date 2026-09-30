package sd_v1

import (
	"github.com/bytedance/sonic"
)

// SDTargetExpressionV32 - struct for SDTargetExpressionV32
type SDTargetExpressionV32 struct {
	SDContentTargetingPredicateV31 *SDContentTargetingPredicateV31
	SDTargetingPredicateNestedV31  *SDTargetingPredicateNestedV31
	SDTargetingPredicateV31        *SDTargetingPredicateV31
	jsonState                      *targetingJSONState
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
