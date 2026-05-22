package product_fees_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// IdType The type of product identifier used in a `FeesEstimateByIdRequest`.
type IdType string

// List of IdType
const (
	IDTYPE_ASIN       IdType = "ASIN"
	IDTYPE_SELLER_SKU IdType = "SellerSKU"
)

// All allowed values of IdType enum
var AllowedIdTypeEnumValues = []IdType{
	IDTYPE_ASIN,
	IDTYPE_SELLER_SKU,
}

func (v *IdType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdType(value)
	for _, existing := range AllowedIdTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdType", value)
}

// NewIdTypeFromValue returns a pointer to a valid IdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdTypeFromValue(v string) (*IdType, error) {
	ev := IdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IdType: valid values are %v", v, AllowedIdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdType) IsValid() bool {
	for _, existing := range AllowedIdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IdType value
func (v IdType) Ptr() *IdType {
	return &v
}

type NullableIdType struct {
	value *IdType
	isSet bool
}

func (v NullableIdType) Get() *IdType {
	return v.value
}

func (v *NullableIdType) Set(val *IdType) {
	v.value = val
	v.isSet = true
}

func (v NullableIdType) IsSet() bool {
	return v.isSet
}

func (v *NullableIdType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdType(val *IdType) *NullableIdType {
	return &NullableIdType{value: val, isSet: true}
}

func (v NullableIdType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIdType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
