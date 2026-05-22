package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
	"gopkg.in/validator.v2"
)

// TargetingExpressionInner - struct for TargetingExpressionInner
type TargetingExpressionInner struct {
	ContentTargetingPredicate *ContentTargetingPredicate
	TargetingPredicate        *TargetingPredicate
	TargetingPredicateLegacy  *TargetingPredicateLegacy
	TargetingPredicateNested  *TargetingPredicateNested
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

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TargetingExpressionInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContentTargetingPredicate
	err = newDecoder(data).Decode(&dst.ContentTargetingPredicate)
	if err == nil {
		jsonContentTargetingPredicate, _ := sonic.Marshal(dst.ContentTargetingPredicate)
		if string(jsonContentTargetingPredicate) == "{}" { // empty struct
			dst.ContentTargetingPredicate = nil
		} else {
			if err = validator.Validate(dst.ContentTargetingPredicate); err != nil {
				dst.ContentTargetingPredicate = nil
			} else {
				match++
			}
		}
	} else {
		dst.ContentTargetingPredicate = nil
	}

	// try to unmarshal data into TargetingPredicate
	err = newDecoder(data).Decode(&dst.TargetingPredicate)
	if err == nil {
		jsonTargetingPredicate, _ := sonic.Marshal(dst.TargetingPredicate)
		if string(jsonTargetingPredicate) == "{}" { // empty struct
			dst.TargetingPredicate = nil
		} else {
			if err = validator.Validate(dst.TargetingPredicate); err != nil {
				dst.TargetingPredicate = nil
			} else {
				match++
			}
		}
	} else {
		dst.TargetingPredicate = nil
	}

	// try to unmarshal data into TargetingPredicateLegacy
	err = newDecoder(data).Decode(&dst.TargetingPredicateLegacy)
	if err == nil {
		jsonTargetingPredicateLegacy, _ := sonic.Marshal(dst.TargetingPredicateLegacy)
		if string(jsonTargetingPredicateLegacy) == "{}" { // empty struct
			dst.TargetingPredicateLegacy = nil
		} else {
			if err = validator.Validate(dst.TargetingPredicateLegacy); err != nil {
				dst.TargetingPredicateLegacy = nil
			} else {
				match++
			}
		}
	} else {
		dst.TargetingPredicateLegacy = nil
	}

	// try to unmarshal data into TargetingPredicateNested
	err = newDecoder(data).Decode(&dst.TargetingPredicateNested)
	if err == nil {
		jsonTargetingPredicateNested, _ := sonic.Marshal(dst.TargetingPredicateNested)
		if string(jsonTargetingPredicateNested) == "{}" { // empty struct
			dst.TargetingPredicateNested = nil
		} else {
			if err = validator.Validate(dst.TargetingPredicateNested); err != nil {
				dst.TargetingPredicateNested = nil
			} else {
				match++
			}
		}
	} else {
		dst.TargetingPredicateNested = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContentTargetingPredicate = nil
		dst.TargetingPredicate = nil
		dst.TargetingPredicateLegacy = nil
		dst.TargetingPredicateNested = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TargetingExpressionInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TargetingExpressionInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TargetingExpressionInner) MarshalJSON() ([]byte, error) {
	if src.ContentTargetingPredicate != nil {
		return sonic.Marshal(&src.ContentTargetingPredicate)
	}

	if src.TargetingPredicate != nil {
		return sonic.Marshal(&src.TargetingPredicate)
	}

	if src.TargetingPredicateLegacy != nil {
		return sonic.Marshal(&src.TargetingPredicateLegacy)
	}

	if src.TargetingPredicateNested != nil {
		return sonic.Marshal(&src.TargetingPredicateNested)
	}

	return nil, nil // no data in oneOf schemas
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
