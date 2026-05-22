package supply_sources_20200701

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SupplySourceStatusReadOnly The `SupplySource` status.
type SupplySourceStatusReadOnly string

// List of SupplySourceStatusReadOnly
const (
	SUPPLYSOURCESTATUSREADONLY_ACTIVE   SupplySourceStatusReadOnly = "Active"
	SUPPLYSOURCESTATUSREADONLY_INACTIVE SupplySourceStatusReadOnly = "Inactive"
	SUPPLYSOURCESTATUSREADONLY_ARCHIVED SupplySourceStatusReadOnly = "Archived"
)

// All allowed values of SupplySourceStatusReadOnly enum
var AllowedSupplySourceStatusReadOnlyEnumValues = []SupplySourceStatusReadOnly{
	SUPPLYSOURCESTATUSREADONLY_ACTIVE,
	SUPPLYSOURCESTATUSREADONLY_INACTIVE,
	SUPPLYSOURCESTATUSREADONLY_ARCHIVED,
}

func (v *SupplySourceStatusReadOnly) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupplySourceStatusReadOnly(value)
	for _, existing := range AllowedSupplySourceStatusReadOnlyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupplySourceStatusReadOnly", value)
}

// NewSupplySourceStatusReadOnlyFromValue returns a pointer to a valid SupplySourceStatusReadOnly
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupplySourceStatusReadOnlyFromValue(v string) (*SupplySourceStatusReadOnly, error) {
	ev := SupplySourceStatusReadOnly(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupplySourceStatusReadOnly: valid values are %v", v, AllowedSupplySourceStatusReadOnlyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupplySourceStatusReadOnly) IsValid() bool {
	for _, existing := range AllowedSupplySourceStatusReadOnlyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SupplySourceStatusReadOnly value
func (v SupplySourceStatusReadOnly) Ptr() *SupplySourceStatusReadOnly {
	return &v
}

type NullableSupplySourceStatusReadOnly struct {
	value *SupplySourceStatusReadOnly
	isSet bool
}

func (v NullableSupplySourceStatusReadOnly) Get() *SupplySourceStatusReadOnly {
	return v.value
}

func (v *NullableSupplySourceStatusReadOnly) Set(val *SupplySourceStatusReadOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplySourceStatusReadOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplySourceStatusReadOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplySourceStatusReadOnly(val *SupplySourceStatusReadOnly) *NullableSupplySourceStatusReadOnly {
	return &NullableSupplySourceStatusReadOnly{value: val, isSet: true}
}

func (v NullableSupplySourceStatusReadOnly) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSupplySourceStatusReadOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
