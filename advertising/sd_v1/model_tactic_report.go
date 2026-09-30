package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// TacticReport The advertising tactic associated with the campaign. The following table lists available tactic names: |Tactic Name|Type|Description| |-----------|-----|-----------| |T00020     |Contextual targeting | Choose individual products to show your ads in placements related to those products.<br> Choose individual categories to show your ads in placements related to those categories on and off Amazon.| |T00030     |Audiences or Contextual Targeting | Select individual products, categories, refined categories, or audiences to show your ads.|
type TacticReport string

// List of TacticReport
const (
	TACTICREPORT_T00020 TacticReport = "T00020"
	TACTICREPORT_T00030 TacticReport = "T00030"
)

// All allowed values of TacticReport enum
var AllowedTacticReportEnumValues = []TacticReport{
	"T00020",
	"T00030",
}

func (v *TacticReport) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = TacticReport(value)
	return nil
}

// NewTacticReportFromValue returns a pointer to a valid TacticReport
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTacticReportFromValue(v string) (*TacticReport, error) {
	ev := TacticReport(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TacticReport: valid values are %v", v, AllowedTacticReportEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TacticReport) IsValid() bool {
	for _, existing := range AllowedTacticReportEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TacticReport value
func (v TacticReport) Ptr() *TacticReport {
	return &v
}

type NullableTacticReport struct {
	value *TacticReport
	isSet bool
}

func (v NullableTacticReport) Get() *TacticReport {
	return v.value
}

func (v *NullableTacticReport) Set(val *TacticReport) {
	v.value = val
	v.isSet = true
}

func (v NullableTacticReport) IsSet() bool {
	return v.isSet
}

func (v *NullableTacticReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTacticReport(val *TacticReport) *NullableTacticReport {
	return &NullableTacticReport{value: val, isSet: true}
}

func (v NullableTacticReport) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTacticReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
