package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SBInsightsMatchType The match type. For more information, see [match types](https://advertising.amazon.com/help#GHTRFDZRJPW6764R) in the Amazon Advertising support center.
type SBInsightsMatchType string

// List of SBInsightsMatchType
const (
	SBINSIGHTSMATCHTYPE_EXACT  SBInsightsMatchType = "EXACT"
	SBINSIGHTSMATCHTYPE_PHRASE SBInsightsMatchType = "PHRASE"
	SBINSIGHTSMATCHTYPE_BROAD  SBInsightsMatchType = "BROAD"
)

// All allowed values of SBInsightsMatchType enum
var AllowedSBInsightsMatchTypeEnumValues = []SBInsightsMatchType{
	"EXACT",
	"PHRASE",
	"BROAD",
}

func (v *SBInsightsMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SBInsightsMatchType(value)
	return nil
}

// NewSBInsightsMatchTypeFromValue returns a pointer to a valid SBInsightsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSBInsightsMatchTypeFromValue(v string) (*SBInsightsMatchType, error) {
	ev := SBInsightsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SBInsightsMatchType: valid values are %v", v, AllowedSBInsightsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SBInsightsMatchType) IsValid() bool {
	for _, existing := range AllowedSBInsightsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SBInsightsMatchType value
func (v SBInsightsMatchType) Ptr() *SBInsightsMatchType {
	return &v
}

type NullableSBInsightsMatchType struct {
	value *SBInsightsMatchType
	isSet bool
}

func (v NullableSBInsightsMatchType) Get() *SBInsightsMatchType {
	return v.value
}

func (v *NullableSBInsightsMatchType) Set(val *SBInsightsMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsMatchType(val *SBInsightsMatchType) *NullableSBInsightsMatchType {
	return &NullableSBInsightsMatchType{value: val, isSet: true}
}

func (v NullableSBInsightsMatchType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
