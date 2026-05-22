package product_pricing_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// FulfillmentChannelType Indicates whether the item is fulfilled by Amazon or by the seller (merchant).
type FulfillmentChannelType string

// List of FulfillmentChannelType
const (
	FULFILLMENTCHANNELTYPE_AMAZON   FulfillmentChannelType = "Amazon"
	FULFILLMENTCHANNELTYPE_MERCHANT FulfillmentChannelType = "Merchant"
)

// All allowed values of FulfillmentChannelType enum
var AllowedFulfillmentChannelTypeEnumValues = []FulfillmentChannelType{
	FULFILLMENTCHANNELTYPE_AMAZON,
	FULFILLMENTCHANNELTYPE_MERCHANT,
}

func (v *FulfillmentChannelType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FulfillmentChannelType(value)
	for _, existing := range AllowedFulfillmentChannelTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FulfillmentChannelType", value)
}

// NewFulfillmentChannelTypeFromValue returns a pointer to a valid FulfillmentChannelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFulfillmentChannelTypeFromValue(v string) (*FulfillmentChannelType, error) {
	ev := FulfillmentChannelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FulfillmentChannelType: valid values are %v", v, AllowedFulfillmentChannelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FulfillmentChannelType) IsValid() bool {
	for _, existing := range AllowedFulfillmentChannelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FulfillmentChannelType value
func (v FulfillmentChannelType) Ptr() *FulfillmentChannelType {
	return &v
}

type NullableFulfillmentChannelType struct {
	value *FulfillmentChannelType
	isSet bool
}

func (v NullableFulfillmentChannelType) Get() *FulfillmentChannelType {
	return v.value
}

func (v *NullableFulfillmentChannelType) Set(val *FulfillmentChannelType) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentChannelType) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentChannelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentChannelType(val *FulfillmentChannelType) *NullableFulfillmentChannelType {
	return &NullableFulfillmentChannelType{value: val, isSet: true}
}

func (v NullableFulfillmentChannelType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentChannelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
