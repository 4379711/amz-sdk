package portfolios_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PortfolioQuotaScope the model 'PortfolioQuotaScope'
type PortfolioQuotaScope string

// List of PortfolioQuotaScope
const (
	PORTFOLIOQUOTASCOPE_ACCOUNT PortfolioQuotaScope = "ACCOUNT"
)

// All allowed values of PortfolioQuotaScope enum
var AllowedPortfolioQuotaScopeEnumValues = []PortfolioQuotaScope{
	"ACCOUNT",
}

func (v *PortfolioQuotaScope) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = PortfolioQuotaScope(value)
	return nil
}

// NewPortfolioQuotaScopeFromValue returns a pointer to a valid PortfolioQuotaScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioQuotaScopeFromValue(v string) (*PortfolioQuotaScope, error) {
	ev := PortfolioQuotaScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioQuotaScope: valid values are %v", v, AllowedPortfolioQuotaScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioQuotaScope) IsValid() bool {
	for _, existing := range AllowedPortfolioQuotaScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioQuotaScope value
func (v PortfolioQuotaScope) Ptr() *PortfolioQuotaScope {
	return &v
}

type NullablePortfolioQuotaScope struct {
	value *PortfolioQuotaScope
	isSet bool
}

func (v NullablePortfolioQuotaScope) Get() *PortfolioQuotaScope {
	return v.value
}

func (v *NullablePortfolioQuotaScope) Set(val *PortfolioQuotaScope) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioQuotaScope) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioQuotaScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioQuotaScope(val *PortfolioQuotaScope) *NullablePortfolioQuotaScope {
	return &NullablePortfolioQuotaScope{value: val, isSet: true}
}

func (v NullablePortfolioQuotaScope) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioQuotaScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
