package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ProductLocation The product location of the campaign. - SOLD_ON_AMAZON - For products sold on Amazon websites. - NOT_SOLD_ON_AMAZON - For products not sold on Amazon websites. - SOLD_ON_DTC - Deprecated (For products sold on DTC websites).
type ProductLocation string

// List of ProductLocation
const (
	PRODUCTLOCATION_SOLD_ON_AMAZON     ProductLocation = "SOLD_ON_AMAZON"
	PRODUCTLOCATION_NOT_SOLD_ON_AMAZON ProductLocation = "NOT_SOLD_ON_AMAZON"
	PRODUCTLOCATION_SOLD_ON_DTC        ProductLocation = "SOLD_ON_DTC"
)

// All allowed values of ProductLocation enum
var AllowedProductLocationEnumValues = []ProductLocation{
	"SOLD_ON_AMAZON",
	"NOT_SOLD_ON_AMAZON",
	"SOLD_ON_DTC",
}

func (v *ProductLocation) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductLocation(value)
	for _, existing := range AllowedProductLocationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductLocation", value)
}

// NewProductLocationFromValue returns a pointer to a valid ProductLocation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductLocationFromValue(v string) (*ProductLocation, error) {
	ev := ProductLocation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductLocation: valid values are %v", v, AllowedProductLocationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductLocation) IsValid() bool {
	for _, existing := range AllowedProductLocationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductLocation value
func (v ProductLocation) Ptr() *ProductLocation {
	return &v
}

type NullableProductLocation struct {
	value *ProductLocation
	isSet bool
}

func (v NullableProductLocation) Get() *ProductLocation {
	return v.value
}

func (v *NullableProductLocation) Set(val *ProductLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableProductLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableProductLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductLocation(val *ProductLocation) *NullableProductLocation {
	return &NullableProductLocation{value: val, isSet: true}
}

func (v NullableProductLocation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
