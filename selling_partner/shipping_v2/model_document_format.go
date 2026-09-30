package shipping_v2

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// DocumentFormat The file format of the document.
type DocumentFormat string

// List of DocumentFormat
const (
	DOCUMENTFORMAT_PDF DocumentFormat = "PDF"
	DOCUMENTFORMAT_PNG DocumentFormat = "PNG"
	DOCUMENTFORMAT_ZPL DocumentFormat = "ZPL"
)

// All allowed values of DocumentFormat enum
var AllowedDocumentFormatEnumValues = []DocumentFormat{
	DOCUMENTFORMAT_PDF,
	DOCUMENTFORMAT_PNG,
	DOCUMENTFORMAT_ZPL,
}

func (v *DocumentFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = DocumentFormat(value)
	return nil
}

// NewDocumentFormatFromValue returns a pointer to a valid DocumentFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentFormatFromValue(v string) (*DocumentFormat, error) {
	ev := DocumentFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DocumentFormat: valid values are %v", v, AllowedDocumentFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentFormat) IsValid() bool {
	for _, existing := range AllowedDocumentFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocumentFormat value
func (v DocumentFormat) Ptr() *DocumentFormat {
	return &v
}

type NullableDocumentFormat struct {
	value *DocumentFormat
	isSet bool
}

func (v NullableDocumentFormat) Get() *DocumentFormat {
	return v.value
}

func (v *NullableDocumentFormat) Set(val *DocumentFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentFormat(val *DocumentFormat) *NullableDocumentFormat {
	return &NullableDocumentFormat{value: val, isSet: true}
}

func (v NullableDocumentFormat) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDocumentFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
