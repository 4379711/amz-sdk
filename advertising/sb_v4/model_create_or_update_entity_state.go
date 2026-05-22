package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CreateOrUpdateEntityState Entity state for create or update operation.
type CreateOrUpdateEntityState string

// List of CreateOrUpdateEntityState
const (
	CREATEORUPDATEENTITYSTATE_ENABLED CreateOrUpdateEntityState = "ENABLED"
	CREATEORUPDATEENTITYSTATE_PAUSED  CreateOrUpdateEntityState = "PAUSED"
)

// All allowed values of CreateOrUpdateEntityState enum
var AllowedCreateOrUpdateEntityStateEnumValues = []CreateOrUpdateEntityState{
	"ENABLED",
	"PAUSED",
}

func (v *CreateOrUpdateEntityState) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreateOrUpdateEntityState(value)
	for _, existing := range AllowedCreateOrUpdateEntityStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreateOrUpdateEntityState", value)
}

// NewCreateOrUpdateEntityStateFromValue returns a pointer to a valid CreateOrUpdateEntityState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateOrUpdateEntityStateFromValue(v string) (*CreateOrUpdateEntityState, error) {
	ev := CreateOrUpdateEntityState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateOrUpdateEntityState: valid values are %v", v, AllowedCreateOrUpdateEntityStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateOrUpdateEntityState) IsValid() bool {
	for _, existing := range AllowedCreateOrUpdateEntityStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateOrUpdateEntityState value
func (v CreateOrUpdateEntityState) Ptr() *CreateOrUpdateEntityState {
	return &v
}

type NullableCreateOrUpdateEntityState struct {
	value *CreateOrUpdateEntityState
	isSet bool
}

func (v NullableCreateOrUpdateEntityState) Get() *CreateOrUpdateEntityState {
	return v.value
}

func (v *NullableCreateOrUpdateEntityState) Set(val *CreateOrUpdateEntityState) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateEntityState) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateEntityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateEntityState(val *CreateOrUpdateEntityState) *NullableCreateOrUpdateEntityState {
	return &NullableCreateOrUpdateEntityState{value: val, isSet: true}
}

func (v NullableCreateOrUpdateEntityState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateOrUpdateEntityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
