package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsMissingValueErrorReason the model 'SponsoredProductsMissingValueErrorReason'
type SponsoredProductsMissingValueErrorReason string

// List of SponsoredProductsMissingValueErrorReason
const (
	SPONSOREDPRODUCTSMISSINGVALUEERRORREASON_MISSING_VALUE SponsoredProductsMissingValueErrorReason = "MISSING_VALUE"
)

// All allowed values of SponsoredProductsMissingValueErrorReason enum
var AllowedSponsoredProductsMissingValueErrorReasonEnumValues = []SponsoredProductsMissingValueErrorReason{
	"MISSING_VALUE",
}

func (v *SponsoredProductsMissingValueErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsMissingValueErrorReason(value)
	return nil
}

// NewSponsoredProductsMissingValueErrorReasonFromValue returns a pointer to a valid SponsoredProductsMissingValueErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsMissingValueErrorReasonFromValue(v string) (*SponsoredProductsMissingValueErrorReason, error) {
	ev := SponsoredProductsMissingValueErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsMissingValueErrorReason: valid values are %v", v, AllowedSponsoredProductsMissingValueErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsMissingValueErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsMissingValueErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsMissingValueErrorReason value
func (v SponsoredProductsMissingValueErrorReason) Ptr() *SponsoredProductsMissingValueErrorReason {
	return &v
}

type NullableSponsoredProductsMissingValueErrorReason struct {
	value *SponsoredProductsMissingValueErrorReason
	isSet bool
}

func (v NullableSponsoredProductsMissingValueErrorReason) Get() *SponsoredProductsMissingValueErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsMissingValueErrorReason) Set(val *SponsoredProductsMissingValueErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMissingValueErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMissingValueErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMissingValueErrorReason(val *SponsoredProductsMissingValueErrorReason) *NullableSponsoredProductsMissingValueErrorReason {
	return &NullableSponsoredProductsMissingValueErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsMissingValueErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMissingValueErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
