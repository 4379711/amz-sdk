package sd_v1

import (
	"github.com/bytedance/sonic"
)

// TargetingExpressionInner - struct for TargetingExpressionInner
type TargetingExpressionInner struct {
	ContentTargetingPredicate *ContentTargetingPredicate
	TargetingPredicate        *TargetingPredicate
	TargetingPredicateLegacy  *TargetingPredicateLegacy
	TargetingPredicateNested  *TargetingPredicateNested
	jsonState                 *targetingJSONState
}

// ContentTargetingPredicateAsTargetingExpressionInner is a convenience function that returns ContentTargetingPredicate wrapped in TargetingExpressionInner
func ContentTargetingPredicateAsTargetingExpressionInner(v *ContentTargetingPredicate) TargetingExpressionInner {
	return TargetingExpressionInner{
		ContentTargetingPredicate: v,
	}
}

// TargetingPredicateAsTargetingExpressionInner is a convenience function that returns TargetingPredicate wrapped in TargetingExpressionInner
func TargetingPredicateAsTargetingExpressionInner(v *TargetingPredicate) TargetingExpressionInner {
	return TargetingExpressionInner{
		TargetingPredicate: v,
	}
}

// TargetingPredicateLegacyAsTargetingExpressionInner is a convenience function that returns TargetingPredicateLegacy wrapped in TargetingExpressionInner
func TargetingPredicateLegacyAsTargetingExpressionInner(v *TargetingPredicateLegacy) TargetingExpressionInner {
	return TargetingExpressionInner{
		TargetingPredicateLegacy: v,
	}
}

// TargetingPredicateNestedAsTargetingExpressionInner is a convenience function that returns TargetingPredicateNested wrapped in TargetingExpressionInner
func TargetingPredicateNestedAsTargetingExpressionInner(v *TargetingPredicateNested) TargetingExpressionInner {
	return TargetingExpressionInner{
		TargetingPredicateNested: v,
	}
}

// Get the actual instance
func (obj *TargetingExpressionInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ContentTargetingPredicate != nil {
		return obj.ContentTargetingPredicate
	}

	if obj.TargetingPredicate != nil {
		return obj.TargetingPredicate
	}

	if obj.TargetingPredicateLegacy != nil {
		return obj.TargetingPredicateLegacy
	}

	if obj.TargetingPredicateNested != nil {
		return obj.TargetingPredicateNested
	}

	// all schemas are nil
	return nil
}

type NullableTargetingExpressionInner struct {
	value *TargetingExpressionInner
	isSet bool
}

func (v NullableTargetingExpressionInner) Get() *TargetingExpressionInner {
	return v.value
}

func (v *NullableTargetingExpressionInner) Set(val *TargetingExpressionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingExpressionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingExpressionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingExpressionInner(val *TargetingExpressionInner) *NullableTargetingExpressionInner {
	return &NullableTargetingExpressionInner{value: val, isSet: true}
}

func (v NullableTargetingExpressionInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingExpressionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
