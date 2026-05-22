package product_eligibility

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// Check - A union of all the checks that we would want to skip
type Check struct {
	SkipAllBillingChecks *SkipAllBillingChecks
}

// SkipAllBillingChecksAsCheck is a convenience function that returns SkipAllBillingChecks wrapped in Check
func SkipAllBillingChecksAsCheck(v *SkipAllBillingChecks) Check {
	return Check{
		SkipAllBillingChecks: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Check) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SkipAllBillingChecks
	err = sonic.Unmarshal(data, &dst.SkipAllBillingChecks)
	if err == nil {
		jsonSkipAllBillingChecks, _ := sonic.Marshal(dst.SkipAllBillingChecks)
		if string(jsonSkipAllBillingChecks) == "{}" { // empty struct
			dst.SkipAllBillingChecks = nil
		} else {
			match++
		}
	} else {
		dst.SkipAllBillingChecks = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SkipAllBillingChecks = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Check)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Check)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Check) MarshalJSON() ([]byte, error) {
	if src.SkipAllBillingChecks != nil {
		return sonic.Marshal(&src.SkipAllBillingChecks)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Check) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SkipAllBillingChecks != nil {
		return obj.SkipAllBillingChecks
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj Check) GetActualInstanceValue() interface{} {
	if obj.SkipAllBillingChecks != nil {
		return *obj.SkipAllBillingChecks
	}

	// all schemas are nil
	return nil
}

type NullableCheck struct {
	value *Check
	isSet bool
}

func (v NullableCheck) Get() *Check {
	return v.value
}

func (v *NullableCheck) Set(val *Check) {
	v.value = val
	v.isSet = true
}

func (v NullableCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheck(val *Check) *NullableCheck {
	return &NullableCheck{value: val, isSet: true}
}

func (v NullableCheck) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
