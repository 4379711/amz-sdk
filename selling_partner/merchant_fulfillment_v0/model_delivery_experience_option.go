package merchant_fulfillment_v0

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// DeliveryExperienceOption The delivery confirmation level.
type DeliveryExperienceOption string

// List of DeliveryExperienceOption
const (
	DELIVERYEXPERIENCEOPTION_DELIVERY_CONFIRMATION_WITH_ADULT_SIGNATURE DeliveryExperienceOption = "DeliveryConfirmationWithAdultSignature"
	DELIVERYEXPERIENCEOPTION_DELIVERY_CONFIRMATION_WITH_SIGNATURE       DeliveryExperienceOption = "DeliveryConfirmationWithSignature"
	DELIVERYEXPERIENCEOPTION_DELIVERY_CONFIRMATION_WITHOUT_SIGNATURE    DeliveryExperienceOption = "DeliveryConfirmationWithoutSignature"
	DELIVERYEXPERIENCEOPTION_NO_TRACKING                                DeliveryExperienceOption = "NoTracking"
	DELIVERYEXPERIENCEOPTION_NO_PREFERENCE                              DeliveryExperienceOption = "NoPreference"
)

// All allowed values of DeliveryExperienceOption enum
var AllowedDeliveryExperienceOptionEnumValues = []DeliveryExperienceOption{
	DELIVERYEXPERIENCEOPTION_DELIVERY_CONFIRMATION_WITH_ADULT_SIGNATURE,
	DELIVERYEXPERIENCEOPTION_DELIVERY_CONFIRMATION_WITH_SIGNATURE,
	DELIVERYEXPERIENCEOPTION_DELIVERY_CONFIRMATION_WITHOUT_SIGNATURE,
	DELIVERYEXPERIENCEOPTION_NO_TRACKING,
	DELIVERYEXPERIENCEOPTION_NO_PREFERENCE,
}

func (v *DeliveryExperienceOption) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeliveryExperienceOption(value)
	for _, existing := range AllowedDeliveryExperienceOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeliveryExperienceOption", value)
}

// NewDeliveryExperienceOptionFromValue returns a pointer to a valid DeliveryExperienceOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeliveryExperienceOptionFromValue(v string) (*DeliveryExperienceOption, error) {
	ev := DeliveryExperienceOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeliveryExperienceOption: valid values are %v", v, AllowedDeliveryExperienceOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeliveryExperienceOption) IsValid() bool {
	for _, existing := range AllowedDeliveryExperienceOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeliveryExperienceOption value
func (v DeliveryExperienceOption) Ptr() *DeliveryExperienceOption {
	return &v
}

type NullableDeliveryExperienceOption struct {
	value *DeliveryExperienceOption
	isSet bool
}

func (v NullableDeliveryExperienceOption) Get() *DeliveryExperienceOption {
	return v.value
}

func (v *NullableDeliveryExperienceOption) Set(val *DeliveryExperienceOption) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryExperienceOption) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryExperienceOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryExperienceOption(val *DeliveryExperienceOption) *NullableDeliveryExperienceOption {
	return &NullableDeliveryExperienceOption{value: val, isSet: true}
}

func (v NullableDeliveryExperienceOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeliveryExperienceOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
