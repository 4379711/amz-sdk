package seller_wallet_20240301

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// BalanceType The type of bank account balance.
type BalanceType string

// List of BalanceType
const (
	BALANCETYPE_AVAILABLE  BalanceType = "AVAILABLE"
	BALANCETYPE_LOCKED_IN  BalanceType = "LOCKED_IN"
	BALANCETYPE_LOCKED_OUT BalanceType = "LOCKED_OUT"
	BALANCETYPE_TOTAL      BalanceType = "TOTAL"
)

// All allowed values of BalanceType enum
var AllowedBalanceTypeEnumValues = []BalanceType{
	BALANCETYPE_AVAILABLE,
	BALANCETYPE_LOCKED_IN,
	BALANCETYPE_LOCKED_OUT,
	BALANCETYPE_TOTAL,
}

func (v *BalanceType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = BalanceType(value)
	return nil
}

// NewBalanceTypeFromValue returns a pointer to a valid BalanceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBalanceTypeFromValue(v string) (*BalanceType, error) {
	ev := BalanceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BalanceType: valid values are %v", v, AllowedBalanceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BalanceType) IsValid() bool {
	for _, existing := range AllowedBalanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BalanceType value
func (v BalanceType) Ptr() *BalanceType {
	return &v
}

type NullableBalanceType struct {
	value *BalanceType
	isSet bool
}

func (v NullableBalanceType) Get() *BalanceType {
	return v.value
}

func (v *NullableBalanceType) Set(val *BalanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceType(val *BalanceType) *NullableBalanceType {
	return &NullableBalanceType{value: val, isSet: true}
}

func (v NullableBalanceType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBalanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
