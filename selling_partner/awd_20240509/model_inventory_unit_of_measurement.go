package awd_20240509

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// InventoryUnitOfMeasurement Unit of measurement for the inventory.
type InventoryUnitOfMeasurement string

// List of InventoryUnitOfMeasurement
const (
	INVENTORYUNITOFMEASUREMENT_PRODUCT_UNITS InventoryUnitOfMeasurement = "PRODUCT_UNITS"
	INVENTORYUNITOFMEASUREMENT_CASES         InventoryUnitOfMeasurement = "CASES"
	INVENTORYUNITOFMEASUREMENT_PALLETS       InventoryUnitOfMeasurement = "PALLETS"
)

// All allowed values of InventoryUnitOfMeasurement enum
var AllowedInventoryUnitOfMeasurementEnumValues = []InventoryUnitOfMeasurement{
	INVENTORYUNITOFMEASUREMENT_PRODUCT_UNITS,
	INVENTORYUNITOFMEASUREMENT_CASES,
	INVENTORYUNITOFMEASUREMENT_PALLETS,
}

func (v *InventoryUnitOfMeasurement) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InventoryUnitOfMeasurement(value)
	for _, existing := range AllowedInventoryUnitOfMeasurementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InventoryUnitOfMeasurement", value)
}

// NewInventoryUnitOfMeasurementFromValue returns a pointer to a valid InventoryUnitOfMeasurement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInventoryUnitOfMeasurementFromValue(v string) (*InventoryUnitOfMeasurement, error) {
	ev := InventoryUnitOfMeasurement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InventoryUnitOfMeasurement: valid values are %v", v, AllowedInventoryUnitOfMeasurementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InventoryUnitOfMeasurement) IsValid() bool {
	for _, existing := range AllowedInventoryUnitOfMeasurementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InventoryUnitOfMeasurement value
func (v InventoryUnitOfMeasurement) Ptr() *InventoryUnitOfMeasurement {
	return &v
}

type NullableInventoryUnitOfMeasurement struct {
	value *InventoryUnitOfMeasurement
	isSet bool
}

func (v NullableInventoryUnitOfMeasurement) Get() *InventoryUnitOfMeasurement {
	return v.value
}

func (v *NullableInventoryUnitOfMeasurement) Set(val *InventoryUnitOfMeasurement) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryUnitOfMeasurement) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryUnitOfMeasurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryUnitOfMeasurement(val *InventoryUnitOfMeasurement) *NullableInventoryUnitOfMeasurement {
	return &NullableInventoryUnitOfMeasurement{value: val, isSet: true}
}

func (v NullableInventoryUnitOfMeasurement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInventoryUnitOfMeasurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
