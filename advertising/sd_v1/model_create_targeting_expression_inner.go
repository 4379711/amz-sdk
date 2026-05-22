package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
	"gopkg.in/validator.v2"
)

// CreateTargetingExpressionInner - struct for CreateTargetingExpressionInner
type CreateTargetingExpressionInner struct {
	ContentTargetingPredicate *ContentTargetingPredicate
	TargetingPredicate        *TargetingPredicate
	TargetingPredicateNested  *TargetingPredicateNested
}

// ContentTargetingPredicateAsCreateTargetingExpressionInner is a convenience function that returns ContentTargetingPredicate wrapped in CreateTargetingExpressionInner
func ContentTargetingPredicateAsCreateTargetingExpressionInner(v *ContentTargetingPredicate) CreateTargetingExpressionInner {
	return CreateTargetingExpressionInner{
		ContentTargetingPredicate: v,
	}
}

// TargetingPredicateAsCreateTargetingExpressionInner is a convenience function that returns TargetingPredicate wrapped in CreateTargetingExpressionInner
func TargetingPredicateAsCreateTargetingExpressionInner(v *TargetingPredicate) CreateTargetingExpressionInner {
	return CreateTargetingExpressionInner{
		TargetingPredicate: v,
	}
}

// TargetingPredicateNestedAsCreateTargetingExpressionInner is a convenience function that returns TargetingPredicateNested wrapped in CreateTargetingExpressionInner
func TargetingPredicateNestedAsCreateTargetingExpressionInner(v *TargetingPredicateNested) CreateTargetingExpressionInner {
	return CreateTargetingExpressionInner{
		TargetingPredicateNested: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateTargetingExpressionInner) UnmarshalJSON(data []byte) error {
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
		dst.TargetingPredicateNested = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateTargetingExpressionInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateTargetingExpressionInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateTargetingExpressionInner) MarshalJSON() ([]byte, error) {
	if src.ContentTargetingPredicate != nil {
		return sonic.Marshal(&src.ContentTargetingPredicate)
	}

	if src.TargetingPredicate != nil {
		return sonic.Marshal(&src.TargetingPredicate)
	}

	if src.TargetingPredicateNested != nil {
		return sonic.Marshal(&src.TargetingPredicateNested)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateTargetingExpressionInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ContentTargetingPredicate != nil {
		return obj.ContentTargetingPredicate
	}

	if obj.TargetingPredicate != nil {
		return obj.TargetingPredicate
	}

	if obj.TargetingPredicateNested != nil {
		return obj.TargetingPredicateNested
	}

	// all schemas are nil
	return nil
}

type NullableCreateTargetingExpressionInner struct {
	value *CreateTargetingExpressionInner
	isSet bool
}

func (v NullableCreateTargetingExpressionInner) Get() *CreateTargetingExpressionInner {
	return v.value
}

func (v *NullableCreateTargetingExpressionInner) Set(val *CreateTargetingExpressionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTargetingExpressionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTargetingExpressionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTargetingExpressionInner(val *CreateTargetingExpressionInner) *NullableCreateTargetingExpressionInner {
	return &NullableCreateTargetingExpressionInner{value: val, isSet: true}
}

func (v NullableCreateTargetingExpressionInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateTargetingExpressionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
