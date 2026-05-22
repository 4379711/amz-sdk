package shipping_v2

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// GenerationStatus Generation Status.
type GenerationStatus string

// List of GenerationStatus
const (
	GENERATIONSTATUS_COMPLETED   GenerationStatus = "Completed"
	GENERATIONSTATUS_IN_PROGRESS GenerationStatus = "InProgress"
)

// All allowed values of GenerationStatus enum
var AllowedGenerationStatusEnumValues = []GenerationStatus{
	GENERATIONSTATUS_COMPLETED,
	GENERATIONSTATUS_IN_PROGRESS,
}

func (v *GenerationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GenerationStatus(value)
	for _, existing := range AllowedGenerationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GenerationStatus", value)
}

// NewGenerationStatusFromValue returns a pointer to a valid GenerationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGenerationStatusFromValue(v string) (*GenerationStatus, error) {
	ev := GenerationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GenerationStatus: valid values are %v", v, AllowedGenerationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GenerationStatus) IsValid() bool {
	for _, existing := range AllowedGenerationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GenerationStatus value
func (v GenerationStatus) Ptr() *GenerationStatus {
	return &v
}

type NullableGenerationStatus struct {
	value *GenerationStatus
	isSet bool
}

func (v NullableGenerationStatus) Get() *GenerationStatus {
	return v.value
}

func (v *NullableGenerationStatus) Set(val *GenerationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerationStatus(val *GenerationStatus) *NullableGenerationStatus {
	return &NullableGenerationStatus{value: val, isSet: true}
}

func (v NullableGenerationStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenerationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
