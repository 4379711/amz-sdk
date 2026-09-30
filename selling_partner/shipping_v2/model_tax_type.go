package shipping_v2

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// TaxType Indicates the type of tax.
type TaxType string

// List of TaxType
const (
	TAXTYPE_GST TaxType = "GST"
)

// All allowed values of TaxType enum
var AllowedTaxTypeEnumValues = []TaxType{
	TAXTYPE_GST,
}

func (v *TaxType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = TaxType(value)
	return nil
}

// NewTaxTypeFromValue returns a pointer to a valid TaxType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaxTypeFromValue(v string) (*TaxType, error) {
	ev := TaxType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaxType: valid values are %v", v, AllowedTaxTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaxType) IsValid() bool {
	for _, existing := range AllowedTaxTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaxType value
func (v TaxType) Ptr() *TaxType {
	return &v
}

type NullableTaxType struct {
	value *TaxType
	isSet bool
}

func (v NullableTaxType) Get() *TaxType {
	return v.value
}

func (v *NullableTaxType) Set(val *TaxType) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxType) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxType(val *TaxType) *NullableTaxType {
	return &NullableTaxType{value: val, isSet: true}
}

func (v NullableTaxType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTaxType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
