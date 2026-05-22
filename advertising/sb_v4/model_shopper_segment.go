package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ShopperSegment List of bid adjustments for shopper segments. - NEW_TO_BRAND_PURCHASE - The shopper segment where shopper has not purchased product from the brand.
type ShopperSegment string

// List of ShopperSegment
const (
	SHOPPERSEGMENT_NEW_TO_BRAND_PURCHASE ShopperSegment = "NEW_TO_BRAND_PURCHASE"
)

// All allowed values of ShopperSegment enum
var AllowedShopperSegmentEnumValues = []ShopperSegment{
	"NEW_TO_BRAND_PURCHASE",
}

func (v *ShopperSegment) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShopperSegment(value)
	for _, existing := range AllowedShopperSegmentEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShopperSegment", value)
}

// NewShopperSegmentFromValue returns a pointer to a valid ShopperSegment
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShopperSegmentFromValue(v string) (*ShopperSegment, error) {
	ev := ShopperSegment(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShopperSegment: valid values are %v", v, AllowedShopperSegmentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShopperSegment) IsValid() bool {
	for _, existing := range AllowedShopperSegmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShopperSegment value
func (v ShopperSegment) Ptr() *ShopperSegment {
	return &v
}

type NullableShopperSegment struct {
	value *ShopperSegment
	isSet bool
}

func (v NullableShopperSegment) Get() *ShopperSegment {
	return v.value
}

func (v *NullableShopperSegment) Set(val *ShopperSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableShopperSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableShopperSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopperSegment(val *ShopperSegment) *NullableShopperSegment {
	return &NullableShopperSegment{value: val, isSet: true}
}

func (v NullableShopperSegment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShopperSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
