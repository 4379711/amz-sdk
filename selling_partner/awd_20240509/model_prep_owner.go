package awd_20240509

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PrepOwner The owner of the preparations, if special preparations are required.
type PrepOwner string

// List of PrepOwner
const (
	PREPOWNER_AMAZON PrepOwner = "AMAZON"
	PREPOWNER_SELF   PrepOwner = "SELF"
)

// All allowed values of PrepOwner enum
var AllowedPrepOwnerEnumValues = []PrepOwner{
	PREPOWNER_AMAZON,
	PREPOWNER_SELF,
}

func (v *PrepOwner) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = PrepOwner(value)
	return nil
}

// NewPrepOwnerFromValue returns a pointer to a valid PrepOwner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrepOwnerFromValue(v string) (*PrepOwner, error) {
	ev := PrepOwner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrepOwner: valid values are %v", v, AllowedPrepOwnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrepOwner) IsValid() bool {
	for _, existing := range AllowedPrepOwnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrepOwner value
func (v PrepOwner) Ptr() *PrepOwner {
	return &v
}

type NullablePrepOwner struct {
	value *PrepOwner
	isSet bool
}

func (v NullablePrepOwner) Get() *PrepOwner {
	return v.value
}

func (v *NullablePrepOwner) Set(val *PrepOwner) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepOwner) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepOwner(val *PrepOwner) *NullablePrepOwner {
	return &NullablePrepOwner{value: val, isSet: true}
}

func (v NullablePrepOwner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePrepOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
