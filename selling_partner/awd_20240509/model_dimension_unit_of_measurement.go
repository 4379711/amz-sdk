package awd_20240509

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// DimensionUnitOfMeasurement Unit of measurement for package dimensions.
type DimensionUnitOfMeasurement string

// List of DimensionUnitOfMeasurement
const (
	DIMENSIONUNITOFMEASUREMENT_INCHES      DimensionUnitOfMeasurement = "INCHES"
	DIMENSIONUNITOFMEASUREMENT_CENTIMETERS DimensionUnitOfMeasurement = "CENTIMETERS"
)

// All allowed values of DimensionUnitOfMeasurement enum
var AllowedDimensionUnitOfMeasurementEnumValues = []DimensionUnitOfMeasurement{
	DIMENSIONUNITOFMEASUREMENT_INCHES,
	DIMENSIONUNITOFMEASUREMENT_CENTIMETERS,
}

func (v *DimensionUnitOfMeasurement) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DimensionUnitOfMeasurement(value)
	for _, existing := range AllowedDimensionUnitOfMeasurementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DimensionUnitOfMeasurement", value)
}

// NewDimensionUnitOfMeasurementFromValue returns a pointer to a valid DimensionUnitOfMeasurement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDimensionUnitOfMeasurementFromValue(v string) (*DimensionUnitOfMeasurement, error) {
	ev := DimensionUnitOfMeasurement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DimensionUnitOfMeasurement: valid values are %v", v, AllowedDimensionUnitOfMeasurementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DimensionUnitOfMeasurement) IsValid() bool {
	for _, existing := range AllowedDimensionUnitOfMeasurementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DimensionUnitOfMeasurement value
func (v DimensionUnitOfMeasurement) Ptr() *DimensionUnitOfMeasurement {
	return &v
}

type NullableDimensionUnitOfMeasurement struct {
	value *DimensionUnitOfMeasurement
	isSet bool
}

func (v NullableDimensionUnitOfMeasurement) Get() *DimensionUnitOfMeasurement {
	return v.value
}

func (v *NullableDimensionUnitOfMeasurement) Set(val *DimensionUnitOfMeasurement) {
	v.value = val
	v.isSet = true
}

func (v NullableDimensionUnitOfMeasurement) IsSet() bool {
	return v.isSet
}

func (v *NullableDimensionUnitOfMeasurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimensionUnitOfMeasurement(val *DimensionUnitOfMeasurement) *NullableDimensionUnitOfMeasurement {
	return &NullableDimensionUnitOfMeasurement{value: val, isSet: true}
}

func (v NullableDimensionUnitOfMeasurement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDimensionUnitOfMeasurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
