package vendor_invoices

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// NetCostUnitOfMeasure This field represents weight unit of measure of items that are ordered by cases and supporting priced by weight.
type NetCostUnitOfMeasure string

// List of NetCostUnitOfMeasure
const (
	NETCOSTUNITOFMEASURE_POUNDS    NetCostUnitOfMeasure = "POUNDS"
	NETCOSTUNITOFMEASURE_OUNCES    NetCostUnitOfMeasure = "OUNCES"
	NETCOSTUNITOFMEASURE_GRAMS     NetCostUnitOfMeasure = "GRAMS"
	NETCOSTUNITOFMEASURE_KILOGRAMS NetCostUnitOfMeasure = "KILOGRAMS"
)

// All allowed values of NetCostUnitOfMeasure enum
var AllowedNetCostUnitOfMeasureEnumValues = []NetCostUnitOfMeasure{
	NETCOSTUNITOFMEASURE_POUNDS,
	NETCOSTUNITOFMEASURE_OUNCES,
	NETCOSTUNITOFMEASURE_GRAMS,
	NETCOSTUNITOFMEASURE_KILOGRAMS,
}

func (v *NetCostUnitOfMeasure) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetCostUnitOfMeasure(value)
	for _, existing := range AllowedNetCostUnitOfMeasureEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetCostUnitOfMeasure", value)
}

// NewNetCostUnitOfMeasureFromValue returns a pointer to a valid NetCostUnitOfMeasure
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetCostUnitOfMeasureFromValue(v string) (*NetCostUnitOfMeasure, error) {
	ev := NetCostUnitOfMeasure(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetCostUnitOfMeasure: valid values are %v", v, AllowedNetCostUnitOfMeasureEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetCostUnitOfMeasure) IsValid() bool {
	for _, existing := range AllowedNetCostUnitOfMeasureEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetCostUnitOfMeasure value
func (v NetCostUnitOfMeasure) Ptr() *NetCostUnitOfMeasure {
	return &v
}

type NullableNetCostUnitOfMeasure struct {
	value *NetCostUnitOfMeasure
	isSet bool
}

func (v NullableNetCostUnitOfMeasure) Get() *NetCostUnitOfMeasure {
	return v.value
}

func (v *NullableNetCostUnitOfMeasure) Set(val *NetCostUnitOfMeasure) {
	v.value = val
	v.isSet = true
}

func (v NullableNetCostUnitOfMeasure) IsSet() bool {
	return v.isSet
}

func (v *NullableNetCostUnitOfMeasure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetCostUnitOfMeasure(val *NetCostUnitOfMeasure) *NullableNetCostUnitOfMeasure {
	return &NullableNetCostUnitOfMeasure{value: val, isSet: true}
}

func (v NullableNetCostUnitOfMeasure) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableNetCostUnitOfMeasure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
