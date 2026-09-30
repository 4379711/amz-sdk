package vendor_direct_fulfillment_shipping_20211228

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ContainerLabelFormat The format of the container label.
type ContainerLabelFormat string

// List of ContainerLabelFormat
const (
	CONTAINERLABELFORMAT_PNG ContainerLabelFormat = "PNG"
	CONTAINERLABELFORMAT_ZPL ContainerLabelFormat = "ZPL"
)

// All allowed values of ContainerLabelFormat enum
var AllowedContainerLabelFormatEnumValues = []ContainerLabelFormat{
	CONTAINERLABELFORMAT_PNG,
	CONTAINERLABELFORMAT_ZPL,
}

func (v *ContainerLabelFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = ContainerLabelFormat(value)
	return nil
}

// NewContainerLabelFormatFromValue returns a pointer to a valid ContainerLabelFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContainerLabelFormatFromValue(v string) (*ContainerLabelFormat, error) {
	ev := ContainerLabelFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContainerLabelFormat: valid values are %v", v, AllowedContainerLabelFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContainerLabelFormat) IsValid() bool {
	for _, existing := range AllowedContainerLabelFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerLabelFormat value
func (v ContainerLabelFormat) Ptr() *ContainerLabelFormat {
	return &v
}

type NullableContainerLabelFormat struct {
	value *ContainerLabelFormat
	isSet bool
}

func (v NullableContainerLabelFormat) Get() *ContainerLabelFormat {
	return v.value
}

func (v *NullableContainerLabelFormat) Set(val *ContainerLabelFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerLabelFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerLabelFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerLabelFormat(val *ContainerLabelFormat) *NullableContainerLabelFormat {
	return &NullableContainerLabelFormat{value: val, isSet: true}
}

func (v NullableContainerLabelFormat) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContainerLabelFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
