package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsState Entity State. | State |  description | |-----------|------------| | `ENABLED` | The entity is enabled. | | `PAUSED` | The entity is disabled. | | `ARCHIVED` | The entity is archived. |
type SponsoredProductsState string

// List of SponsoredProductsState
const (
	SPONSOREDPRODUCTSSTATE_ENABLED  SponsoredProductsState = "ENABLED"
	SPONSOREDPRODUCTSSTATE_PAUSED   SponsoredProductsState = "PAUSED"
	SPONSOREDPRODUCTSSTATE_ARCHIVED SponsoredProductsState = "ARCHIVED"
	SPONSOREDPRODUCTSSTATE_PROPOSED SponsoredProductsState = "PROPOSED"
)

// All allowed values of SponsoredProductsState enum
var AllowedSponsoredProductsStateEnumValues = []SponsoredProductsState{
	"ENABLED",
	"PAUSED",
	"ARCHIVED",
	"PROPOSED",
}

func (v *SponsoredProductsState) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsState(value)
	return nil
}

// NewSponsoredProductsStateFromValue returns a pointer to a valid SponsoredProductsState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsStateFromValue(v string) (*SponsoredProductsState, error) {
	ev := SponsoredProductsState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsState: valid values are %v", v, AllowedSponsoredProductsStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsState) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsState value
func (v SponsoredProductsState) Ptr() *SponsoredProductsState {
	return &v
}

type NullableSponsoredProductsState struct {
	value *SponsoredProductsState
	isSet bool
}

func (v NullableSponsoredProductsState) Get() *SponsoredProductsState {
	return v.value
}

func (v *NullableSponsoredProductsState) Set(val *SponsoredProductsState) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsState) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsState(val *SponsoredProductsState) *NullableSponsoredProductsState {
	return &NullableSponsoredProductsState{value: val, isSet: true}
}

func (v NullableSponsoredProductsState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
