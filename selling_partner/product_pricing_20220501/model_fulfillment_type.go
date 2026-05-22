package product_pricing_20220501

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// FulfillmentType Indicates whether the item is fulfilled by Amazon or by the seller (merchant).
type FulfillmentType string

// List of FulfillmentType
const (
	FULFILLMENTTYPE_AFN FulfillmentType = "AFN"
	FULFILLMENTTYPE_MFN FulfillmentType = "MFN"
)

// All allowed values of FulfillmentType enum
var AllowedFulfillmentTypeEnumValues = []FulfillmentType{
	FULFILLMENTTYPE_AFN,
	FULFILLMENTTYPE_MFN,
}

func (v *FulfillmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FulfillmentType(value)
	for _, existing := range AllowedFulfillmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FulfillmentType", value)
}

// NewFulfillmentTypeFromValue returns a pointer to a valid FulfillmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFulfillmentTypeFromValue(v string) (*FulfillmentType, error) {
	ev := FulfillmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FulfillmentType: valid values are %v", v, AllowedFulfillmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FulfillmentType) IsValid() bool {
	for _, existing := range AllowedFulfillmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FulfillmentType value
func (v FulfillmentType) Ptr() *FulfillmentType {
	return &v
}

type NullableFulfillmentType struct {
	value *FulfillmentType
	isSet bool
}

func (v NullableFulfillmentType) Get() *FulfillmentType {
	return v.value
}

func (v *NullableFulfillmentType) Set(val *FulfillmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentType(val *FulfillmentType) *NullableFulfillmentType {
	return &NullableFulfillmentType{value: val, isSet: true}
}

func (v NullableFulfillmentType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
