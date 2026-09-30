package fulfillment_inbound_20240320

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// LabelPrintType Indicates the type of print type for a given label.
type LabelPrintType string

// List of LabelPrintType
const (
	LABELPRINTTYPE_STANDARD_FORMAT  LabelPrintType = "STANDARD_FORMAT"
	LABELPRINTTYPE_THERMAL_PRINTING LabelPrintType = "THERMAL_PRINTING"
)

// All allowed values of LabelPrintType enum
var AllowedLabelPrintTypeEnumValues = []LabelPrintType{
	LABELPRINTTYPE_STANDARD_FORMAT,
	LABELPRINTTYPE_THERMAL_PRINTING,
}

func (v *LabelPrintType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = LabelPrintType(value)
	return nil
}

// NewLabelPrintTypeFromValue returns a pointer to a valid LabelPrintType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLabelPrintTypeFromValue(v string) (*LabelPrintType, error) {
	ev := LabelPrintType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LabelPrintType: valid values are %v", v, AllowedLabelPrintTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LabelPrintType) IsValid() bool {
	for _, existing := range AllowedLabelPrintTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LabelPrintType value
func (v LabelPrintType) Ptr() *LabelPrintType {
	return &v
}

type NullableLabelPrintType struct {
	value *LabelPrintType
	isSet bool
}

func (v NullableLabelPrintType) Get() *LabelPrintType {
	return v.value
}

func (v *NullableLabelPrintType) Set(val *LabelPrintType) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelPrintType) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelPrintType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelPrintType(val *LabelPrintType) *NullableLabelPrintType {
	return &NullableLabelPrintType{value: val, isSet: true}
}

func (v NullableLabelPrintType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLabelPrintType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
