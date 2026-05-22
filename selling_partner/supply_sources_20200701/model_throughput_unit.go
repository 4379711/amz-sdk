package supply_sources_20200701

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ThroughputUnit The throughput unit
type ThroughputUnit string

// List of ThroughputUnit
const (
	THROUGHPUTUNIT_ORDER ThroughputUnit = "Order"
)

// All allowed values of ThroughputUnit enum
var AllowedThroughputUnitEnumValues = []ThroughputUnit{
	THROUGHPUTUNIT_ORDER,
}

func (v *ThroughputUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThroughputUnit(value)
	for _, existing := range AllowedThroughputUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThroughputUnit", value)
}

// NewThroughputUnitFromValue returns a pointer to a valid ThroughputUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewThroughputUnitFromValue(v string) (*ThroughputUnit, error) {
	ev := ThroughputUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ThroughputUnit: valid values are %v", v, AllowedThroughputUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ThroughputUnit) IsValid() bool {
	for _, existing := range AllowedThroughputUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ThroughputUnit value
func (v ThroughputUnit) Ptr() *ThroughputUnit {
	return &v
}

type NullableThroughputUnit struct {
	value *ThroughputUnit
	isSet bool
}

func (v NullableThroughputUnit) Get() *ThroughputUnit {
	return v.value
}

func (v *NullableThroughputUnit) Set(val *ThroughputUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableThroughputUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableThroughputUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThroughputUnit(val *ThroughputUnit) *NullableThroughputUnit {
	return &NullableThroughputUnit{value: val, isSet: true}
}

func (v NullableThroughputUnit) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThroughputUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
