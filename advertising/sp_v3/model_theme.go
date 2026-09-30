package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// Theme The bid recommendation theme. This API currently supports `CONVERSION_OPPORTUNITIES`, `PRIME_DAY`, `FALL_PRIME_DEAL_EVENT`, and `BFCM_HOLIDAY` themes.
type Theme string

// List of Theme
const (
	THEME_CONVERSION_OPPORTUNITIES Theme = "CONVERSION_OPPORTUNITIES"
	THEME_PRIME_DAY                Theme = "PRIME_DAY"
	THEME_FALL_PRIME_DEAL_EVENT    Theme = "FALL_PRIME_DEAL_EVENT"
	THEME_BFCM_HOLIDAY             Theme = "BFCM_HOLIDAY"
)

// All allowed values of Theme enum
var AllowedThemeEnumValues = []Theme{
	"CONVERSION_OPPORTUNITIES",
	"PRIME_DAY",
	"FALL_PRIME_DEAL_EVENT",
	"BFCM_HOLIDAY",
}

func (v *Theme) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = Theme(value)
	return nil
}

// NewThemeFromValue returns a pointer to a valid Theme
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewThemeFromValue(v string) (*Theme, error) {
	ev := Theme(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Theme: valid values are %v", v, AllowedThemeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Theme) IsValid() bool {
	for _, existing := range AllowedThemeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Theme value
func (v Theme) Ptr() *Theme {
	return &v
}

type NullableTheme struct {
	value *Theme
	isSet bool
}

func (v NullableTheme) Get() *Theme {
	return v.value
}

func (v *NullableTheme) Set(val *Theme) {
	v.value = val
	v.isSet = true
}

func (v NullableTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTheme(val *Theme) *NullableTheme {
	return &NullableTheme{value: val, isSet: true}
}

func (v NullableTheme) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
