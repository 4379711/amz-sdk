package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// EntityState The current resource state.
type EntityState string

// List of EntityState
const (
	ENTITYSTATE_ENABLED EntityState = "ENABLED"
	ENTITYSTATE_PAUSED  EntityState = "PAUSED"
)

// All allowed values of EntityState enum
var AllowedEntityStateEnumValues = []EntityState{
	"ENABLED",
	"PAUSED",
}

func (v *EntityState) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntityState(value)
	for _, existing := range AllowedEntityStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntityState", value)
}

// NewEntityStateFromValue returns a pointer to a valid EntityState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntityStateFromValue(v string) (*EntityState, error) {
	ev := EntityState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntityState: valid values are %v", v, AllowedEntityStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntityState) IsValid() bool {
	for _, existing := range AllowedEntityStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityState value
func (v EntityState) Ptr() *EntityState {
	return &v
}

type NullableEntityState struct {
	value *EntityState
	isSet bool
}

func (v NullableEntityState) Get() *EntityState {
	return v.value
}

func (v *NullableEntityState) Set(val *EntityState) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityState) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityState(val *EntityState) *NullableEntityState {
	return &NullableEntityState{value: val, isSet: true}
}

func (v NullableEntityState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEntityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
