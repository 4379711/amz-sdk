package shipping

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ServiceType The type of shipping service that will be used for the service offering.
type ServiceType string

// List of ServiceType
const (
	SERVICETYPE_AMAZON_SHIPPING_GROUND   ServiceType = "Amazon Shipping Ground"
	SERVICETYPE_AMAZON_SHIPPING_STANDARD ServiceType = "Amazon Shipping Standard"
	SERVICETYPE_AMAZON_SHIPPING_PREMIUM  ServiceType = "Amazon Shipping Premium"
)

// All allowed values of ServiceType enum
var AllowedServiceTypeEnumValues = []ServiceType{
	SERVICETYPE_AMAZON_SHIPPING_GROUND,
	SERVICETYPE_AMAZON_SHIPPING_STANDARD,
	SERVICETYPE_AMAZON_SHIPPING_PREMIUM,
}

func (v *ServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceType(value)
	for _, existing := range AllowedServiceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceType", value)
}

// NewServiceTypeFromValue returns a pointer to a valid ServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceTypeFromValue(v string) (*ServiceType, error) {
	ev := ServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceType: valid values are %v", v, AllowedServiceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceType) IsValid() bool {
	for _, existing := range AllowedServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceType value
func (v ServiceType) Ptr() *ServiceType {
	return &v
}

type NullableServiceType struct {
	value *ServiceType
	isSet bool
}

func (v NullableServiceType) Get() *ServiceType {
	return v.value
}

func (v *NullableServiceType) Set(val *ServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceType(val *ServiceType) *NullableServiceType {
	return &NullableServiceType{value: val, isSet: true}
}

func (v NullableServiceType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
