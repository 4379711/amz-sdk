package report_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// AsyncReportAdProduct The advertising product.
type AsyncReportAdProduct string

// List of AsyncReportAdProduct
const (
	ASYNCREPORTADPRODUCT_SPONSORED_PRODUCTS   AsyncReportAdProduct = "SPONSORED_PRODUCTS"
	ASYNCREPORTADPRODUCT_SPONSORED_BRANDS     AsyncReportAdProduct = "SPONSORED_BRANDS"
	ASYNCREPORTADPRODUCT_SPONSORED_DISPLAY    AsyncReportAdProduct = "SPONSORED_DISPLAY"
	ASYNCREPORTADPRODUCT_SPONSORED_TELEVISION AsyncReportAdProduct = "SPONSORED_TELEVISION"
	ASYNCREPORTADPRODUCT_DEMAND_SIDE_PLATFORM AsyncReportAdProduct = "DEMAND_SIDE_PLATFORM"
	ASYNCREPORTADPRODUCT_ALL                  AsyncReportAdProduct = "ALL"
)

// All allowed values of AsyncReportAdProduct enum
var AllowedAsyncReportAdProductEnumValues = []AsyncReportAdProduct{
	"SPONSORED_PRODUCTS",
	"SPONSORED_BRANDS",
	"SPONSORED_DISPLAY",
	"SPONSORED_TELEVISION",
	"DEMAND_SIDE_PLATFORM",
	"ALL",
}

func (v *AsyncReportAdProduct) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = AsyncReportAdProduct(value)
	return nil
}

// NewAsyncReportAdProductFromValue returns a pointer to a valid AsyncReportAdProduct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAsyncReportAdProductFromValue(v string) (*AsyncReportAdProduct, error) {
	ev := AsyncReportAdProduct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AsyncReportAdProduct: valid values are %v", v, AllowedAsyncReportAdProductEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AsyncReportAdProduct) IsValid() bool {
	for _, existing := range AllowedAsyncReportAdProductEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AsyncReportAdProduct value
func (v AsyncReportAdProduct) Ptr() *AsyncReportAdProduct {
	return &v
}

type NullableAsyncReportAdProduct struct {
	value *AsyncReportAdProduct
	isSet bool
}

func (v NullableAsyncReportAdProduct) Get() *AsyncReportAdProduct {
	return v.value
}

func (v *NullableAsyncReportAdProduct) Set(val *AsyncReportAdProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncReportAdProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncReportAdProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncReportAdProduct(val *AsyncReportAdProduct) *NullableAsyncReportAdProduct {
	return &NullableAsyncReportAdProduct{value: val, isSet: true}
}

func (v NullableAsyncReportAdProduct) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsyncReportAdProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
