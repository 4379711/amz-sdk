package invoices_api_model_20240619

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ExportStatus The current status of the request.
type ExportStatus string

// List of ExportStatus
const (
	EXPORTSTATUS_REQUESTED  ExportStatus = "REQUESTED"
	EXPORTSTATUS_PROCESSING ExportStatus = "PROCESSING"
	EXPORTSTATUS_DONE       ExportStatus = "DONE"
	EXPORTSTATUS_ERROR      ExportStatus = "ERROR"
)

// All allowed values of ExportStatus enum
var AllowedExportStatusEnumValues = []ExportStatus{
	EXPORTSTATUS_REQUESTED,
	EXPORTSTATUS_PROCESSING,
	EXPORTSTATUS_DONE,
	EXPORTSTATUS_ERROR,
}

func (v *ExportStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExportStatus(value)
	for _, existing := range AllowedExportStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExportStatus", value)
}

// NewExportStatusFromValue returns a pointer to a valid ExportStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExportStatusFromValue(v string) (*ExportStatus, error) {
	ev := ExportStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExportStatus: valid values are %v", v, AllowedExportStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExportStatus) IsValid() bool {
	for _, existing := range AllowedExportStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExportStatus value
func (v ExportStatus) Ptr() *ExportStatus {
	return &v
}

type NullableExportStatus struct {
	value *ExportStatus
	isSet bool
}

func (v NullableExportStatus) Get() *ExportStatus {
	return v.value
}

func (v *NullableExportStatus) Set(val *ExportStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableExportStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableExportStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportStatus(val *ExportStatus) *NullableExportStatus {
	return &NullableExportStatus{value: val, isSet: true}
}

func (v NullableExportStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableExportStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
