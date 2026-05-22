package product_pricing_v0

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CustomerType Indicates whether to request Consumer or Business offers. Default is Consumer.
type CustomerType string

// List of CustomerType
const (
	CUSTOMERTYPE_CONSUMER CustomerType = "Consumer"
	CUSTOMERTYPE_BUSINESS CustomerType = "Business"
)

// All allowed values of CustomerType enum
var AllowedCustomerTypeEnumValues = []CustomerType{
	CUSTOMERTYPE_CONSUMER,
	CUSTOMERTYPE_BUSINESS,
}

func (v *CustomerType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerType(value)
	for _, existing := range AllowedCustomerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerType", value)
}

// NewCustomerTypeFromValue returns a pointer to a valid CustomerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerTypeFromValue(v string) (*CustomerType, error) {
	ev := CustomerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerType: valid values are %v", v, AllowedCustomerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerType) IsValid() bool {
	for _, existing := range AllowedCustomerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomerType value
func (v CustomerType) Ptr() *CustomerType {
	return &v
}

type NullableCustomerType struct {
	value *CustomerType
	isSet bool
}

func (v NullableCustomerType) Get() *CustomerType {
	return v.value
}

func (v *NullableCustomerType) Set(val *CustomerType) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerType) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerType(val *CustomerType) *NullableCustomerType {
	return &NullableCustomerType{value: val, isSet: true}
}

func (v NullableCustomerType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCustomerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
