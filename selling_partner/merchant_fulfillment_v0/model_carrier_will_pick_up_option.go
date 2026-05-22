package merchant_fulfillment_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// CarrierWillPickUpOption Carrier will pick up option.
type CarrierWillPickUpOption string

// List of CarrierWillPickUpOption
const (
	CARRIERWILLPICKUPOPTION_CARRIER_WILL_PICK_UP  CarrierWillPickUpOption = "CarrierWillPickUp"
	CARRIERWILLPICKUPOPTION_SHIPPER_WILL_DROP_OFF CarrierWillPickUpOption = "ShipperWillDropOff"
	CARRIERWILLPICKUPOPTION_NO_PREFERENCE         CarrierWillPickUpOption = "NoPreference"
)

// All allowed values of CarrierWillPickUpOption enum
var AllowedCarrierWillPickUpOptionEnumValues = []CarrierWillPickUpOption{
	CARRIERWILLPICKUPOPTION_CARRIER_WILL_PICK_UP,
	CARRIERWILLPICKUPOPTION_SHIPPER_WILL_DROP_OFF,
	CARRIERWILLPICKUPOPTION_NO_PREFERENCE,
}

func (v *CarrierWillPickUpOption) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CarrierWillPickUpOption(value)
	for _, existing := range AllowedCarrierWillPickUpOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CarrierWillPickUpOption", value)
}

// NewCarrierWillPickUpOptionFromValue returns a pointer to a valid CarrierWillPickUpOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCarrierWillPickUpOptionFromValue(v string) (*CarrierWillPickUpOption, error) {
	ev := CarrierWillPickUpOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CarrierWillPickUpOption: valid values are %v", v, AllowedCarrierWillPickUpOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CarrierWillPickUpOption) IsValid() bool {
	for _, existing := range AllowedCarrierWillPickUpOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CarrierWillPickUpOption value
func (v CarrierWillPickUpOption) Ptr() *CarrierWillPickUpOption {
	return &v
}

type NullableCarrierWillPickUpOption struct {
	value *CarrierWillPickUpOption
	isSet bool
}

func (v NullableCarrierWillPickUpOption) Get() *CarrierWillPickUpOption {
	return v.value
}

func (v *NullableCarrierWillPickUpOption) Set(val *CarrierWillPickUpOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierWillPickUpOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierWillPickUpOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierWillPickUpOption(val *CarrierWillPickUpOption) *NullableCarrierWillPickUpOption {
	return &NullableCarrierWillPickUpOption{value: val, isSet: true}
}

func (v NullableCarrierWillPickUpOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCarrierWillPickUpOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
