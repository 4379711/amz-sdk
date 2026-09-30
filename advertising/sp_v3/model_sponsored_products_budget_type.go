package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsBudgetType the model 'SponsoredProductsBudgetType'
type SponsoredProductsBudgetType string

// List of SponsoredProductsBudgetType
const (
	SPONSOREDPRODUCTSBUDGETTYPE_DAILY SponsoredProductsBudgetType = "DAILY"
	SPONSOREDPRODUCTSBUDGETTYPE_OTHER SponsoredProductsBudgetType = "OTHER"
)

// All allowed values of SponsoredProductsBudgetType enum
var AllowedSponsoredProductsBudgetTypeEnumValues = []SponsoredProductsBudgetType{
	"DAILY",
	"OTHER",
}

func (v *SponsoredProductsBudgetType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsBudgetType(value)
	return nil
}

// NewSponsoredProductsBudgetTypeFromValue returns a pointer to a valid SponsoredProductsBudgetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsBudgetTypeFromValue(v string) (*SponsoredProductsBudgetType, error) {
	ev := SponsoredProductsBudgetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsBudgetType: valid values are %v", v, AllowedSponsoredProductsBudgetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsBudgetType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsBudgetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsBudgetType value
func (v SponsoredProductsBudgetType) Ptr() *SponsoredProductsBudgetType {
	return &v
}

type NullableSponsoredProductsBudgetType struct {
	value *SponsoredProductsBudgetType
	isSet bool
}

func (v NullableSponsoredProductsBudgetType) Get() *SponsoredProductsBudgetType {
	return v.value
}

func (v *NullableSponsoredProductsBudgetType) Set(val *SponsoredProductsBudgetType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBudgetType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBudgetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBudgetType(val *SponsoredProductsBudgetType) *NullableSponsoredProductsBudgetType {
	return &NullableSponsoredProductsBudgetType{value: val, isSet: true}
}

func (v NullableSponsoredProductsBudgetType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBudgetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
