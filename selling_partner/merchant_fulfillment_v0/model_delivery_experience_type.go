package merchant_fulfillment_v0

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// DeliveryExperienceType The delivery confirmation level.
type DeliveryExperienceType string

// List of DeliveryExperienceType
const (
	DELIVERYEXPERIENCETYPE_DELIVERY_CONFIRMATION_WITH_ADULT_SIGNATURE DeliveryExperienceType = "DeliveryConfirmationWithAdultSignature"
	DELIVERYEXPERIENCETYPE_DELIVERY_CONFIRMATION_WITH_SIGNATURE       DeliveryExperienceType = "DeliveryConfirmationWithSignature"
	DELIVERYEXPERIENCETYPE_DELIVERY_CONFIRMATION_WITHOUT_SIGNATURE    DeliveryExperienceType = "DeliveryConfirmationWithoutSignature"
	DELIVERYEXPERIENCETYPE_NO_TRACKING                                DeliveryExperienceType = "NoTracking"
)

// All allowed values of DeliveryExperienceType enum
var AllowedDeliveryExperienceTypeEnumValues = []DeliveryExperienceType{
	DELIVERYEXPERIENCETYPE_DELIVERY_CONFIRMATION_WITH_ADULT_SIGNATURE,
	DELIVERYEXPERIENCETYPE_DELIVERY_CONFIRMATION_WITH_SIGNATURE,
	DELIVERYEXPERIENCETYPE_DELIVERY_CONFIRMATION_WITHOUT_SIGNATURE,
	DELIVERYEXPERIENCETYPE_NO_TRACKING,
}

func (v *DeliveryExperienceType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeliveryExperienceType(value)
	for _, existing := range AllowedDeliveryExperienceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeliveryExperienceType", value)
}

// NewDeliveryExperienceTypeFromValue returns a pointer to a valid DeliveryExperienceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeliveryExperienceTypeFromValue(v string) (*DeliveryExperienceType, error) {
	ev := DeliveryExperienceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeliveryExperienceType: valid values are %v", v, AllowedDeliveryExperienceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeliveryExperienceType) IsValid() bool {
	for _, existing := range AllowedDeliveryExperienceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeliveryExperienceType value
func (v DeliveryExperienceType) Ptr() *DeliveryExperienceType {
	return &v
}

type NullableDeliveryExperienceType struct {
	value *DeliveryExperienceType
	isSet bool
}

func (v NullableDeliveryExperienceType) Get() *DeliveryExperienceType {
	return v.value
}

func (v *NullableDeliveryExperienceType) Set(val *DeliveryExperienceType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryExperienceType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryExperienceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryExperienceType(val *DeliveryExperienceType) *NullableDeliveryExperienceType {
	return &NullableDeliveryExperienceType{value: val, isSet: true}
}

func (v NullableDeliveryExperienceType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeliveryExperienceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
