package portfolios_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PortfolioValueLimitErrorReason the model 'PortfolioValueLimitErrorReason'
type PortfolioValueLimitErrorReason string

// List of PortfolioValueLimitErrorReason
const (
	PORTFOLIOVALUELIMITERRORREASON_TOO_LOW            PortfolioValueLimitErrorReason = "TOO_LOW"
	PORTFOLIOVALUELIMITERRORREASON_TOO_HIGH           PortfolioValueLimitErrorReason = "TOO_HIGH"
	PORTFOLIOVALUELIMITERRORREASON_INVALID_ENUM_VALUE PortfolioValueLimitErrorReason = "INVALID_ENUM_VALUE"
)

// All allowed values of PortfolioValueLimitErrorReason enum
var AllowedPortfolioValueLimitErrorReasonEnumValues = []PortfolioValueLimitErrorReason{
	"TOO_LOW",
	"TOO_HIGH",
	"INVALID_ENUM_VALUE",
}

func (v *PortfolioValueLimitErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = PortfolioValueLimitErrorReason(value)
	return nil
}

// NewPortfolioValueLimitErrorReasonFromValue returns a pointer to a valid PortfolioValueLimitErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioValueLimitErrorReasonFromValue(v string) (*PortfolioValueLimitErrorReason, error) {
	ev := PortfolioValueLimitErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioValueLimitErrorReason: valid values are %v", v, AllowedPortfolioValueLimitErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioValueLimitErrorReason) IsValid() bool {
	for _, existing := range AllowedPortfolioValueLimitErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioValueLimitErrorReason value
func (v PortfolioValueLimitErrorReason) Ptr() *PortfolioValueLimitErrorReason {
	return &v
}

type NullablePortfolioValueLimitErrorReason struct {
	value *PortfolioValueLimitErrorReason
	isSet bool
}

func (v NullablePortfolioValueLimitErrorReason) Get() *PortfolioValueLimitErrorReason {
	return v.value
}

func (v *NullablePortfolioValueLimitErrorReason) Set(val *PortfolioValueLimitErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioValueLimitErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioValueLimitErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioValueLimitErrorReason(val *PortfolioValueLimitErrorReason) *NullablePortfolioValueLimitErrorReason {
	return &NullablePortfolioValueLimitErrorReason{value: val, isSet: true}
}

func (v NullablePortfolioValueLimitErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioValueLimitErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
