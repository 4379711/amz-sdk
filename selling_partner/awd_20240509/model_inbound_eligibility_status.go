package awd_20240509

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// InboundEligibilityStatus Enum denoting the package inbound eligibility.
type InboundEligibilityStatus string

// List of InboundEligibilityStatus
const (
	INBOUNDELIGIBILITYSTATUS_ELIGIBLE   InboundEligibilityStatus = "ELIGIBLE"
	INBOUNDELIGIBILITYSTATUS_INELIGIBLE InboundEligibilityStatus = "INELIGIBLE"
)

// All allowed values of InboundEligibilityStatus enum
var AllowedInboundEligibilityStatusEnumValues = []InboundEligibilityStatus{
	INBOUNDELIGIBILITYSTATUS_ELIGIBLE,
	INBOUNDELIGIBILITYSTATUS_INELIGIBLE,
}

func (v *InboundEligibilityStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = InboundEligibilityStatus(value)
	return nil
}

// NewInboundEligibilityStatusFromValue returns a pointer to a valid InboundEligibilityStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInboundEligibilityStatusFromValue(v string) (*InboundEligibilityStatus, error) {
	ev := InboundEligibilityStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InboundEligibilityStatus: valid values are %v", v, AllowedInboundEligibilityStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InboundEligibilityStatus) IsValid() bool {
	for _, existing := range AllowedInboundEligibilityStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InboundEligibilityStatus value
func (v InboundEligibilityStatus) Ptr() *InboundEligibilityStatus {
	return &v
}

type NullableInboundEligibilityStatus struct {
	value *InboundEligibilityStatus
	isSet bool
}

func (v NullableInboundEligibilityStatus) Get() *InboundEligibilityStatus {
	return v.value
}

func (v *NullableInboundEligibilityStatus) Set(val *InboundEligibilityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundEligibilityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundEligibilityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundEligibilityStatus(val *InboundEligibilityStatus) *NullableInboundEligibilityStatus {
	return &NullableInboundEligibilityStatus{value: val, isSet: true}
}

func (v NullableInboundEligibilityStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInboundEligibilityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
