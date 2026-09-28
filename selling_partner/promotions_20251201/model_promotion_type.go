package promotions_20251201

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PromotionType The type of promotion, categorized by the discount rules and eligibility criteria.
type PromotionType string

// List of PromotionType
const (
	PROMOTIONTYPE_BASKET_BUILDING PromotionType = "BASKET_BUILDING"
	PROMOTIONTYPE_DEAL            PromotionType = "DEAL"
	PROMOTIONTYPE_PRICE_DISCOUNT  PromotionType = "PRICE_DISCOUNT"
	PROMOTIONTYPE_COUPON          PromotionType = "COUPON"
)

// All allowed values of PromotionType enum
var AllowedPromotionTypeEnumValues = []PromotionType{
	"BASKET_BUILDING",
	"DEAL",
	"PRICE_DISCOUNT",
	"COUPON",
}

func (v *PromotionType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PromotionType(value)
	for _, existing := range AllowedPromotionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PromotionType", value)
}

// NewPromotionTypeFromValue returns a pointer to a valid PromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionTypeFromValue(v string) (*PromotionType, error) {
	ev := PromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionType: valid values are %v", v, AllowedPromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionType) IsValid() bool {
	for _, existing := range AllowedPromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PromotionType value
func (v PromotionType) Ptr() *PromotionType {
	return &v
}

type NullablePromotionType struct {
	value *PromotionType
	isSet bool
}

func (v NullablePromotionType) Get() *PromotionType {
	return v.value
}

func (v *NullablePromotionType) Set(val *PromotionType) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionType) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionType(val *PromotionType) *NullablePromotionType {
	return &NullablePromotionType{value: val, isSet: true}
}

func (v NullablePromotionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePromotionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
