package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ShopperCohortType The shopper cohort type. The shopperCohortType is required to specify the type of shopper cohort used to apply bid adjustments. Currently only \"AUDIENCE_SEGMENT\" is supported.
type ShopperCohortType string

// List of ShopperCohortType
const (
	SHOPPERCOHORTTYPE_AUDIENCE_SEGMENT ShopperCohortType = "AUDIENCE_SEGMENT"
)

// All allowed values of ShopperCohortType enum
var AllowedShopperCohortTypeEnumValues = []ShopperCohortType{
	"AUDIENCE_SEGMENT",
}

func (v *ShopperCohortType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShopperCohortType(value)
	for _, existing := range AllowedShopperCohortTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShopperCohortType", value)
}

// NewShopperCohortTypeFromValue returns a pointer to a valid ShopperCohortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShopperCohortTypeFromValue(v string) (*ShopperCohortType, error) {
	ev := ShopperCohortType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShopperCohortType: valid values are %v", v, AllowedShopperCohortTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShopperCohortType) IsValid() bool {
	for _, existing := range AllowedShopperCohortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShopperCohortType value
func (v ShopperCohortType) Ptr() *ShopperCohortType {
	return &v
}

type NullableShopperCohortType struct {
	value *ShopperCohortType
	isSet bool
}

func (v NullableShopperCohortType) Get() *ShopperCohortType {
	return v.value
}

func (v *NullableShopperCohortType) Set(val *ShopperCohortType) {
	v.value = val
	v.isSet = true
}

func (v NullableShopperCohortType) IsSet() bool {
	return v.isSet
}

func (v *NullableShopperCohortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopperCohortType(val *ShopperCohortType) *NullableShopperCohortType {
	return &NullableShopperCohortType{value: val, isSet: true}
}

func (v NullableShopperCohortType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShopperCohortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
