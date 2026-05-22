package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ResultStatus the model 'ResultStatus'
type ResultStatus string

// List of ResultStatus
const (
	RESULTSTATUS_SUCCESS         ResultStatus = "SUCCESS"
	RESULTSTATUS_FAILURE         ResultStatus = "FAILURE"
	RESULTSTATUS_PARTIAL_FAILURE ResultStatus = "PARTIAL_FAILURE"
)

// All allowed values of ResultStatus enum
var AllowedResultStatusEnumValues = []ResultStatus{
	"SUCCESS",
	"FAILURE",
	"PARTIAL_FAILURE",
}

func (v *ResultStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResultStatus(value)
	for _, existing := range AllowedResultStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResultStatus", value)
}

// NewResultStatusFromValue returns a pointer to a valid ResultStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResultStatusFromValue(v string) (*ResultStatus, error) {
	ev := ResultStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResultStatus: valid values are %v", v, AllowedResultStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResultStatus) IsValid() bool {
	for _, existing := range AllowedResultStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResultStatus value
func (v ResultStatus) Ptr() *ResultStatus {
	return &v
}

type NullableResultStatus struct {
	value *ResultStatus
	isSet bool
}

func (v NullableResultStatus) Get() *ResultStatus {
	return v.value
}

func (v *NullableResultStatus) Set(val *ResultStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableResultStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableResultStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultStatus(val *ResultStatus) *NullableResultStatus {
	return &NullableResultStatus{value: val, isSet: true}
}

func (v NullableResultStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableResultStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
