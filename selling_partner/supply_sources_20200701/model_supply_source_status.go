package supply_sources_20200701

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SupplySourceStatus The `SupplySource` status
type SupplySourceStatus string

// List of SupplySourceStatus
const (
	SUPPLYSOURCESTATUS_ACTIVE   SupplySourceStatus = "Active"
	SUPPLYSOURCESTATUS_INACTIVE SupplySourceStatus = "Inactive"
)

// All allowed values of SupplySourceStatus enum
var AllowedSupplySourceStatusEnumValues = []SupplySourceStatus{
	SUPPLYSOURCESTATUS_ACTIVE,
	SUPPLYSOURCESTATUS_INACTIVE,
}

func (v *SupplySourceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupplySourceStatus(value)
	for _, existing := range AllowedSupplySourceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupplySourceStatus", value)
}

// NewSupplySourceStatusFromValue returns a pointer to a valid SupplySourceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupplySourceStatusFromValue(v string) (*SupplySourceStatus, error) {
	ev := SupplySourceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupplySourceStatus: valid values are %v", v, AllowedSupplySourceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupplySourceStatus) IsValid() bool {
	for _, existing := range AllowedSupplySourceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SupplySourceStatus value
func (v SupplySourceStatus) Ptr() *SupplySourceStatus {
	return &v
}

type NullableSupplySourceStatus struct {
	value *SupplySourceStatus
	isSet bool
}

func (v NullableSupplySourceStatus) Get() *SupplySourceStatus {
	return v.value
}

func (v *NullableSupplySourceStatus) Set(val *SupplySourceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplySourceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplySourceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplySourceStatus(val *SupplySourceStatus) *NullableSupplySourceStatus {
	return &NullableSupplySourceStatus{value: val, isSet: true}
}

func (v NullableSupplySourceStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSupplySourceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
