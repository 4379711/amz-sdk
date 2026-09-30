package shipping_v2

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ChannelType The shipment source channel type.
type ChannelType string

// List of ChannelType
const (
	CHANNELTYPE_AMAZON   ChannelType = "AMAZON"
	CHANNELTYPE_EXTERNAL ChannelType = "EXTERNAL"
)

// All allowed values of ChannelType enum
var AllowedChannelTypeEnumValues = []ChannelType{
	CHANNELTYPE_AMAZON,
	CHANNELTYPE_EXTERNAL,
}

func (v *ChannelType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = ChannelType(value)
	return nil
}

// NewChannelTypeFromValue returns a pointer to a valid ChannelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelTypeFromValue(v string) (*ChannelType, error) {
	ev := ChannelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelType: valid values are %v", v, AllowedChannelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelType) IsValid() bool {
	for _, existing := range AllowedChannelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChannelType value
func (v ChannelType) Ptr() *ChannelType {
	return &v
}

type NullableChannelType struct {
	value *ChannelType
	isSet bool
}

func (v NullableChannelType) Get() *ChannelType {
	return v.value
}

func (v *NullableChannelType) Set(val *ChannelType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelType(val *ChannelType) *NullableChannelType {
	return &NullableChannelType{value: val, isSet: true}
}

func (v NullableChannelType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableChannelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
