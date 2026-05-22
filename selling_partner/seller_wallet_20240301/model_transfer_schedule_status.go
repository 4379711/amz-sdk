package seller_wallet_20240301

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// TransferScheduleStatus The schedule status of the transfer.
type TransferScheduleStatus string

// List of TransferScheduleStatus
const (
	TRANSFERSCHEDULESTATUS_ENABLED  TransferScheduleStatus = "ENABLED"
	TRANSFERSCHEDULESTATUS_DISABLED TransferScheduleStatus = "DISABLED"
	TRANSFERSCHEDULESTATUS_EXPIRED  TransferScheduleStatus = "EXPIRED"
	TRANSFERSCHEDULESTATUS_DELETED  TransferScheduleStatus = "DELETED"
)

// All allowed values of TransferScheduleStatus enum
var AllowedTransferScheduleStatusEnumValues = []TransferScheduleStatus{
	TRANSFERSCHEDULESTATUS_ENABLED,
	TRANSFERSCHEDULESTATUS_DISABLED,
	TRANSFERSCHEDULESTATUS_EXPIRED,
	TRANSFERSCHEDULESTATUS_DELETED,
}

func (v *TransferScheduleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferScheduleStatus(value)
	for _, existing := range AllowedTransferScheduleStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferScheduleStatus", value)
}

// NewTransferScheduleStatusFromValue returns a pointer to a valid TransferScheduleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferScheduleStatusFromValue(v string) (*TransferScheduleStatus, error) {
	ev := TransferScheduleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferScheduleStatus: valid values are %v", v, AllowedTransferScheduleStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferScheduleStatus) IsValid() bool {
	for _, existing := range AllowedTransferScheduleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferScheduleStatus value
func (v TransferScheduleStatus) Ptr() *TransferScheduleStatus {
	return &v
}

type NullableTransferScheduleStatus struct {
	value *TransferScheduleStatus
	isSet bool
}

func (v NullableTransferScheduleStatus) Get() *TransferScheduleStatus {
	return v.value
}

func (v *NullableTransferScheduleStatus) Set(val *TransferScheduleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferScheduleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferScheduleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferScheduleStatus(val *TransferScheduleStatus) *NullableTransferScheduleStatus {
	return &NullableTransferScheduleStatus{value: val, isSet: true}
}

func (v NullableTransferScheduleStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransferScheduleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
