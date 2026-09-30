package fulfillment_inbound_20240320

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// OperationStatus The status of an operation.
type OperationStatus string

// List of OperationStatus
const (
	OPERATIONSTATUS_SUCCESS     OperationStatus = "SUCCESS"
	OPERATIONSTATUS_FAILED      OperationStatus = "FAILED"
	OPERATIONSTATUS_IN_PROGRESS OperationStatus = "IN_PROGRESS"
)

// All allowed values of OperationStatus enum
var AllowedOperationStatusEnumValues = []OperationStatus{
	OPERATIONSTATUS_SUCCESS,
	OPERATIONSTATUS_FAILED,
	OPERATIONSTATUS_IN_PROGRESS,
}

func (v *OperationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = OperationStatus(value)
	return nil
}

// NewOperationStatusFromValue returns a pointer to a valid OperationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperationStatusFromValue(v string) (*OperationStatus, error) {
	ev := OperationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperationStatus: valid values are %v", v, AllowedOperationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperationStatus) IsValid() bool {
	for _, existing := range AllowedOperationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperationStatus value
func (v OperationStatus) Ptr() *OperationStatus {
	return &v
}

type NullableOperationStatus struct {
	value *OperationStatus
	isSet bool
}

func (v NullableOperationStatus) Get() *OperationStatus {
	return v.value
}

func (v *NullableOperationStatus) Set(val *OperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationStatus(val *OperationStatus) *NullableOperationStatus {
	return &NullableOperationStatus{value: val, isSet: true}
}

func (v NullableOperationStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
