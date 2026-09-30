package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsCreateOrUpdateState Entity State. | State |  description | |-----------|------------| | `ENABLED` | The entity is enabled. | | `PAUSED` | The entity is disabled. |
type SponsoredProductsCreateOrUpdateState string

// List of SponsoredProductsCreateOrUpdateState
const (
	SPONSOREDPRODUCTSCREATEORUPDATESTATE_ENABLED  SponsoredProductsCreateOrUpdateState = "ENABLED"
	SPONSOREDPRODUCTSCREATEORUPDATESTATE_PAUSED   SponsoredProductsCreateOrUpdateState = "PAUSED"
	SPONSOREDPRODUCTSCREATEORUPDATESTATE_PROPOSED SponsoredProductsCreateOrUpdateState = "PROPOSED"
)

// All allowed values of SponsoredProductsCreateOrUpdateState enum
var AllowedSponsoredProductsCreateOrUpdateStateEnumValues = []SponsoredProductsCreateOrUpdateState{
	"ENABLED",
	"PAUSED",
	"PROPOSED",
}

func (v *SponsoredProductsCreateOrUpdateState) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsCreateOrUpdateState(value)
	return nil
}

// NewSponsoredProductsCreateOrUpdateStateFromValue returns a pointer to a valid SponsoredProductsCreateOrUpdateState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCreateOrUpdateStateFromValue(v string) (*SponsoredProductsCreateOrUpdateState, error) {
	ev := SponsoredProductsCreateOrUpdateState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCreateOrUpdateState: valid values are %v", v, AllowedSponsoredProductsCreateOrUpdateStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCreateOrUpdateState) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCreateOrUpdateState value
func (v SponsoredProductsCreateOrUpdateState) Ptr() *SponsoredProductsCreateOrUpdateState {
	return &v
}

type NullableSponsoredProductsCreateOrUpdateState struct {
	value *SponsoredProductsCreateOrUpdateState
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateState) Get() *SponsoredProductsCreateOrUpdateState {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateState) Set(val *SponsoredProductsCreateOrUpdateState) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateState) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateState(val *SponsoredProductsCreateOrUpdateState) *NullableSponsoredProductsCreateOrUpdateState {
	return &NullableSponsoredProductsCreateOrUpdateState{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
