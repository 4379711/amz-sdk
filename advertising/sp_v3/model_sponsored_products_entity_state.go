package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsEntityState The current resource state. | State | Description | | --- | --- | | `ENABLED` | Enabled State | | `PAUSED` | Paused State | | `ARCHIVED` | ARCHIVED State | | `ENABLING` | State for Draft Entity Only | | `USER_DELETED` | State for Draft Entity Only | | `OTHER` | Read Only |
type SponsoredProductsEntityState string

// List of SponsoredProductsEntityState
const (
	SPONSOREDPRODUCTSENTITYSTATE_ENABLED      SponsoredProductsEntityState = "ENABLED"
	SPONSOREDPRODUCTSENTITYSTATE_PAUSED       SponsoredProductsEntityState = "PAUSED"
	SPONSOREDPRODUCTSENTITYSTATE_ARCHIVED     SponsoredProductsEntityState = "ARCHIVED"
	SPONSOREDPRODUCTSENTITYSTATE_ENABLING     SponsoredProductsEntityState = "ENABLING"
	SPONSOREDPRODUCTSENTITYSTATE_USER_DELETED SponsoredProductsEntityState = "USER_DELETED"
	SPONSOREDPRODUCTSENTITYSTATE_OTHER        SponsoredProductsEntityState = "OTHER"
	SPONSOREDPRODUCTSENTITYSTATE_PROPOSED     SponsoredProductsEntityState = "PROPOSED"
)

// All allowed values of SponsoredProductsEntityState enum
var AllowedSponsoredProductsEntityStateEnumValues = []SponsoredProductsEntityState{
	"ENABLED",
	"PAUSED",
	"ARCHIVED",
	"ENABLING",
	"USER_DELETED",
	"OTHER",
	"PROPOSED",
}

func (v *SponsoredProductsEntityState) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsEntityState(value)
	return nil
}

// NewSponsoredProductsEntityStateFromValue returns a pointer to a valid SponsoredProductsEntityState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsEntityStateFromValue(v string) (*SponsoredProductsEntityState, error) {
	ev := SponsoredProductsEntityState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsEntityState: valid values are %v", v, AllowedSponsoredProductsEntityStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsEntityState) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsEntityStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsEntityState value
func (v SponsoredProductsEntityState) Ptr() *SponsoredProductsEntityState {
	return &v
}

type NullableSponsoredProductsEntityState struct {
	value *SponsoredProductsEntityState
	isSet bool
}

func (v NullableSponsoredProductsEntityState) Get() *SponsoredProductsEntityState {
	return v.value
}

func (v *NullableSponsoredProductsEntityState) Set(val *SponsoredProductsEntityState) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsEntityState) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsEntityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsEntityState(val *SponsoredProductsEntityState) *NullableSponsoredProductsEntityState {
	return &NullableSponsoredProductsEntityState{value: val, isSet: true}
}

func (v NullableSponsoredProductsEntityState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsEntityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
