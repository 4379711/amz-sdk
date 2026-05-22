package awd_20240509

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// WeightUnitOfMeasurement Unit of measurement for the package weight.
type WeightUnitOfMeasurement string

// List of WeightUnitOfMeasurement
const (
	WEIGHTUNITOFMEASUREMENT_POUNDS    WeightUnitOfMeasurement = "POUNDS"
	WEIGHTUNITOFMEASUREMENT_KILOGRAMS WeightUnitOfMeasurement = "KILOGRAMS"
)

// All allowed values of WeightUnitOfMeasurement enum
var AllowedWeightUnitOfMeasurementEnumValues = []WeightUnitOfMeasurement{
	WEIGHTUNITOFMEASUREMENT_POUNDS,
	WEIGHTUNITOFMEASUREMENT_KILOGRAMS,
}

func (v *WeightUnitOfMeasurement) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WeightUnitOfMeasurement(value)
	for _, existing := range AllowedWeightUnitOfMeasurementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WeightUnitOfMeasurement", value)
}

// NewWeightUnitOfMeasurementFromValue returns a pointer to a valid WeightUnitOfMeasurement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWeightUnitOfMeasurementFromValue(v string) (*WeightUnitOfMeasurement, error) {
	ev := WeightUnitOfMeasurement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WeightUnitOfMeasurement: valid values are %v", v, AllowedWeightUnitOfMeasurementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WeightUnitOfMeasurement) IsValid() bool {
	for _, existing := range AllowedWeightUnitOfMeasurementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WeightUnitOfMeasurement value
func (v WeightUnitOfMeasurement) Ptr() *WeightUnitOfMeasurement {
	return &v
}

type NullableWeightUnitOfMeasurement struct {
	value *WeightUnitOfMeasurement
	isSet bool
}

func (v NullableWeightUnitOfMeasurement) Get() *WeightUnitOfMeasurement {
	return v.value
}

func (v *NullableWeightUnitOfMeasurement) Set(val *WeightUnitOfMeasurement) {
	v.value = val
	v.isSet = true
}

func (v NullableWeightUnitOfMeasurement) IsSet() bool {
	return v.isSet
}

func (v *NullableWeightUnitOfMeasurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeightUnitOfMeasurement(val *WeightUnitOfMeasurement) *NullableWeightUnitOfMeasurement {
	return &NullableWeightUnitOfMeasurement{value: val, isSet: true}
}

func (v NullableWeightUnitOfMeasurement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableWeightUnitOfMeasurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
