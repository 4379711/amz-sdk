package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
	"gopkg.in/validator.v2"
)

// SDTargetingBidRecommendationsResponseV31BidRecommendationsInner - struct for SDTargetingBidRecommendationsResponseV31BidRecommendationsInner
type SDTargetingBidRecommendationsResponseV31BidRecommendationsInner struct {
	SDTargetingBidRecommendationsResponseItemFailureV31 *SDTargetingBidRecommendationsResponseItemFailureV31
	SDTargetingBidRecommendationsResponseItemSuccessV31 *SDTargetingBidRecommendationsResponseItemSuccessV31
}

// SDTargetingBidRecommendationsResponseItemFailureV31AsSDTargetingBidRecommendationsResponseV31BidRecommendationsInner is a convenience function that returns SDTargetingBidRecommendationsResponseItemFailureV31 wrapped in SDTargetingBidRecommendationsResponseV31BidRecommendationsInner
func SDTargetingBidRecommendationsResponseItemFailureV31AsSDTargetingBidRecommendationsResponseV31BidRecommendationsInner(v *SDTargetingBidRecommendationsResponseItemFailureV31) SDTargetingBidRecommendationsResponseV31BidRecommendationsInner {
	return SDTargetingBidRecommendationsResponseV31BidRecommendationsInner{
		SDTargetingBidRecommendationsResponseItemFailureV31: v,
	}
}

// SDTargetingBidRecommendationsResponseItemSuccessV31AsSDTargetingBidRecommendationsResponseV31BidRecommendationsInner is a convenience function that returns SDTargetingBidRecommendationsResponseItemSuccessV31 wrapped in SDTargetingBidRecommendationsResponseV31BidRecommendationsInner
func SDTargetingBidRecommendationsResponseItemSuccessV31AsSDTargetingBidRecommendationsResponseV31BidRecommendationsInner(v *SDTargetingBidRecommendationsResponseItemSuccessV31) SDTargetingBidRecommendationsResponseV31BidRecommendationsInner {
	return SDTargetingBidRecommendationsResponseV31BidRecommendationsInner{
		SDTargetingBidRecommendationsResponseItemSuccessV31: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SDTargetingBidRecommendationsResponseV31BidRecommendationsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SDTargetingBidRecommendationsResponseItemFailureV31
	err = newDecoder(data).Decode(&dst.SDTargetingBidRecommendationsResponseItemFailureV31)
	if err == nil {
		jsonSDTargetingBidRecommendationsResponseItemFailureV31, _ := sonic.Marshal(dst.SDTargetingBidRecommendationsResponseItemFailureV31)
		if string(jsonSDTargetingBidRecommendationsResponseItemFailureV31) == "{}" { // empty struct
			dst.SDTargetingBidRecommendationsResponseItemFailureV31 = nil
		} else {
			if err = validator.Validate(dst.SDTargetingBidRecommendationsResponseItemFailureV31); err != nil {
				dst.SDTargetingBidRecommendationsResponseItemFailureV31 = nil
			} else {
				match++
			}
		}
	} else {
		dst.SDTargetingBidRecommendationsResponseItemFailureV31 = nil
	}

	// try to unmarshal data into SDTargetingBidRecommendationsResponseItemSuccessV31
	err = newDecoder(data).Decode(&dst.SDTargetingBidRecommendationsResponseItemSuccessV31)
	if err == nil {
		jsonSDTargetingBidRecommendationsResponseItemSuccessV31, _ := sonic.Marshal(dst.SDTargetingBidRecommendationsResponseItemSuccessV31)
		if string(jsonSDTargetingBidRecommendationsResponseItemSuccessV31) == "{}" { // empty struct
			dst.SDTargetingBidRecommendationsResponseItemSuccessV31 = nil
		} else {
			if err = validator.Validate(dst.SDTargetingBidRecommendationsResponseItemSuccessV31); err != nil {
				dst.SDTargetingBidRecommendationsResponseItemSuccessV31 = nil
			} else {
				match++
			}
		}
	} else {
		dst.SDTargetingBidRecommendationsResponseItemSuccessV31 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SDTargetingBidRecommendationsResponseItemFailureV31 = nil
		dst.SDTargetingBidRecommendationsResponseItemSuccessV31 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SDTargetingBidRecommendationsResponseV31BidRecommendationsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SDTargetingBidRecommendationsResponseV31BidRecommendationsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SDTargetingBidRecommendationsResponseV31BidRecommendationsInner) MarshalJSON() ([]byte, error) {
	if src.SDTargetingBidRecommendationsResponseItemFailureV31 != nil {
		return sonic.Marshal(&src.SDTargetingBidRecommendationsResponseItemFailureV31)
	}

	if src.SDTargetingBidRecommendationsResponseItemSuccessV31 != nil {
		return sonic.Marshal(&src.SDTargetingBidRecommendationsResponseItemSuccessV31)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SDTargetingBidRecommendationsResponseV31BidRecommendationsInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SDTargetingBidRecommendationsResponseItemFailureV31 != nil {
		return obj.SDTargetingBidRecommendationsResponseItemFailureV31
	}

	if obj.SDTargetingBidRecommendationsResponseItemSuccessV31 != nil {
		return obj.SDTargetingBidRecommendationsResponseItemSuccessV31
	}

	// all schemas are nil
	return nil
}

type NullableSDTargetingBidRecommendationsResponseV31BidRecommendationsInner struct {
	value *SDTargetingBidRecommendationsResponseV31BidRecommendationsInner
	isSet bool
}

func (v NullableSDTargetingBidRecommendationsResponseV31BidRecommendationsInner) Get() *SDTargetingBidRecommendationsResponseV31BidRecommendationsInner {
	return v.value
}

func (v *NullableSDTargetingBidRecommendationsResponseV31BidRecommendationsInner) Set(val *SDTargetingBidRecommendationsResponseV31BidRecommendationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingBidRecommendationsResponseV31BidRecommendationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingBidRecommendationsResponseV31BidRecommendationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingBidRecommendationsResponseV31BidRecommendationsInner(val *SDTargetingBidRecommendationsResponseV31BidRecommendationsInner) *NullableSDTargetingBidRecommendationsResponseV31BidRecommendationsInner {
	return &NullableSDTargetingBidRecommendationsResponseV31BidRecommendationsInner{value: val, isSet: true}
}

func (v NullableSDTargetingBidRecommendationsResponseV31BidRecommendationsInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingBidRecommendationsResponseV31BidRecommendationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
