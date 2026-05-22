package orders_v0

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// OtherDeliveryAttributes Miscellaneous delivery attributes associated with the shipping address.
type OtherDeliveryAttributes string

// List of OtherDeliveryAttributes
const (
	OTHERDELIVERYATTRIBUTES_HAS_ACCESS_POINT OtherDeliveryAttributes = "HAS_ACCESS_POINT"
	OTHERDELIVERYATTRIBUTES_PALLET_ENABLED   OtherDeliveryAttributes = "PALLET_ENABLED"
	OTHERDELIVERYATTRIBUTES_PALLET_DISABLED  OtherDeliveryAttributes = "PALLET_DISABLED"
)

// All allowed values of OtherDeliveryAttributes enum
var AllowedOtherDeliveryAttributesEnumValues = []OtherDeliveryAttributes{
	OTHERDELIVERYATTRIBUTES_HAS_ACCESS_POINT,
	OTHERDELIVERYATTRIBUTES_PALLET_ENABLED,
	OTHERDELIVERYATTRIBUTES_PALLET_DISABLED,
}

func (v *OtherDeliveryAttributes) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OtherDeliveryAttributes(value)
	for _, existing := range AllowedOtherDeliveryAttributesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OtherDeliveryAttributes", value)
}

// NewOtherDeliveryAttributesFromValue returns a pointer to a valid OtherDeliveryAttributes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOtherDeliveryAttributesFromValue(v string) (*OtherDeliveryAttributes, error) {
	ev := OtherDeliveryAttributes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OtherDeliveryAttributes: valid values are %v", v, AllowedOtherDeliveryAttributesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OtherDeliveryAttributes) IsValid() bool {
	for _, existing := range AllowedOtherDeliveryAttributesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OtherDeliveryAttributes value
func (v OtherDeliveryAttributes) Ptr() *OtherDeliveryAttributes {
	return &v
}

type NullableOtherDeliveryAttributes struct {
	value *OtherDeliveryAttributes
	isSet bool
}

func (v NullableOtherDeliveryAttributes) Get() *OtherDeliveryAttributes {
	return v.value
}

func (v *NullableOtherDeliveryAttributes) Set(val *OtherDeliveryAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableOtherDeliveryAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableOtherDeliveryAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtherDeliveryAttributes(val *OtherDeliveryAttributes) *NullableOtherDeliveryAttributes {
	return &NullableOtherDeliveryAttributes{value: val, isSet: true}
}

func (v NullableOtherDeliveryAttributes) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOtherDeliveryAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
