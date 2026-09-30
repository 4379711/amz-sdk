package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// LandingPageThrottledErrorCode the model 'LandingPageThrottledErrorCode'
type LandingPageThrottledErrorCode string

// List of LandingPageThrottledErrorCode
const (
	LANDINGPAGETHROTTLEDERRORCODE_THROTTLED LandingPageThrottledErrorCode = "THROTTLED"
)

// All allowed values of LandingPageThrottledErrorCode enum
var AllowedLandingPageThrottledErrorCodeEnumValues = []LandingPageThrottledErrorCode{
	"THROTTLED",
}

func (v *LandingPageThrottledErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = LandingPageThrottledErrorCode(value)
	return nil
}

// NewLandingPageThrottledErrorCodeFromValue returns a pointer to a valid LandingPageThrottledErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLandingPageThrottledErrorCodeFromValue(v string) (*LandingPageThrottledErrorCode, error) {
	ev := LandingPageThrottledErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LandingPageThrottledErrorCode: valid values are %v", v, AllowedLandingPageThrottledErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LandingPageThrottledErrorCode) IsValid() bool {
	for _, existing := range AllowedLandingPageThrottledErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LandingPageThrottledErrorCode value
func (v LandingPageThrottledErrorCode) Ptr() *LandingPageThrottledErrorCode {
	return &v
}

type NullableLandingPageThrottledErrorCode struct {
	value *LandingPageThrottledErrorCode
	isSet bool
}

func (v NullableLandingPageThrottledErrorCode) Get() *LandingPageThrottledErrorCode {
	return v.value
}

func (v *NullableLandingPageThrottledErrorCode) Set(val *LandingPageThrottledErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableLandingPageThrottledErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableLandingPageThrottledErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandingPageThrottledErrorCode(val *LandingPageThrottledErrorCode) *NullableLandingPageThrottledErrorCode {
	return &NullableLandingPageThrottledErrorCode{value: val, isSet: true}
}

func (v NullableLandingPageThrottledErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLandingPageThrottledErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
