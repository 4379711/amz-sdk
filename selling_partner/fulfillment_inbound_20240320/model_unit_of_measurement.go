package fulfillment_inbound_20240320

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// UnitOfMeasurement Unit of linear measure.
type UnitOfMeasurement string

// List of UnitOfMeasurement
const (
	UNITOFMEASUREMENT_IN UnitOfMeasurement = "IN"
	UNITOFMEASUREMENT_CM UnitOfMeasurement = "CM"
)

// All allowed values of UnitOfMeasurement enum
var AllowedUnitOfMeasurementEnumValues = []UnitOfMeasurement{
	UNITOFMEASUREMENT_IN,
	UNITOFMEASUREMENT_CM,
}

func (v *UnitOfMeasurement) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UnitOfMeasurement(value)
	for _, existing := range AllowedUnitOfMeasurementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UnitOfMeasurement", value)
}

// NewUnitOfMeasurementFromValue returns a pointer to a valid UnitOfMeasurement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUnitOfMeasurementFromValue(v string) (*UnitOfMeasurement, error) {
	ev := UnitOfMeasurement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UnitOfMeasurement: valid values are %v", v, AllowedUnitOfMeasurementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UnitOfMeasurement) IsValid() bool {
	for _, existing := range AllowedUnitOfMeasurementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnitOfMeasurement value
func (v UnitOfMeasurement) Ptr() *UnitOfMeasurement {
	return &v
}

type NullableUnitOfMeasurement struct {
	value *UnitOfMeasurement
	isSet bool
}

func (v NullableUnitOfMeasurement) Get() *UnitOfMeasurement {
	return v.value
}

func (v *NullableUnitOfMeasurement) Set(val *UnitOfMeasurement) {
	v.value = val
	v.isSet = true
}

func (v NullableUnitOfMeasurement) IsSet() bool {
	return v.isSet
}

func (v *NullableUnitOfMeasurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnitOfMeasurement(val *UnitOfMeasurement) *NullableUnitOfMeasurement {
	return &NullableUnitOfMeasurement{value: val, isSet: true}
}

func (v NullableUnitOfMeasurement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnitOfMeasurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
