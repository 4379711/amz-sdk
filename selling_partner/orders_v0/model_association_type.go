package orders_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// AssociationType The type of association an item has with an order item.
type AssociationType string

// List of AssociationType
const (
	ASSOCIATIONTYPE_VALUE_ADD_SERVICE AssociationType = "VALUE_ADD_SERVICE"
)

// All allowed values of AssociationType enum
var AllowedAssociationTypeEnumValues = []AssociationType{
	ASSOCIATIONTYPE_VALUE_ADD_SERVICE,
}

func (v *AssociationType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssociationType(value)
	for _, existing := range AllowedAssociationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssociationType", value)
}

// NewAssociationTypeFromValue returns a pointer to a valid AssociationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssociationTypeFromValue(v string) (*AssociationType, error) {
	ev := AssociationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssociationType: valid values are %v", v, AllowedAssociationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssociationType) IsValid() bool {
	for _, existing := range AllowedAssociationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssociationType value
func (v AssociationType) Ptr() *AssociationType {
	return &v
}

type NullableAssociationType struct {
	value *AssociationType
	isSet bool
}

func (v NullableAssociationType) Get() *AssociationType {
	return v.value
}

func (v *NullableAssociationType) Set(val *AssociationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationType(val *AssociationType) *NullableAssociationType {
	return &NullableAssociationType{value: val, isSet: true}
}

func (v NullableAssociationType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssociationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
