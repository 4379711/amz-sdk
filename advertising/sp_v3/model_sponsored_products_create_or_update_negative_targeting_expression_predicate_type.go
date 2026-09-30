package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType The type of nagative targeting expression. You can only specify values for the following predicates: | Predicate | Description | | --- | --- | | `ASIN_BRAND_SAME_AS` | Target the brand that is the same as the brand expressed. | | `ASIN_SAME_AS` | Target an ASIN that is the same as the ASIN expressed. |
type SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType string

// List of SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType
const (
	SPONSOREDPRODUCTSCREATEORUPDATENEGATIVETARGETINGEXPRESSIONPREDICATETYPE_BRAND_SAME_AS SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType = "ASIN_BRAND_SAME_AS"
	SPONSOREDPRODUCTSCREATEORUPDATENEGATIVETARGETINGEXPRESSIONPREDICATETYPE_SAME_AS       SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType = "ASIN_SAME_AS"
)

// All allowed values of SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType enum
var AllowedSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateTypeEnumValues = []SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType{
	"ASIN_BRAND_SAME_AS",
	"ASIN_SAME_AS",
}

func (v *SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType(value)
	return nil
}

// NewSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateTypeFromValue returns a pointer to a valid SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateTypeFromValue(v string) (*SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType, error) {
	ev := SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType: valid values are %v", v, AllowedSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType value
func (v SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType) Ptr() *SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType {
	return &v
}

type NullableSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType struct {
	value *SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType) Get() *SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType) Set(val *SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType(val *SponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType) *NullableSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType {
	return &NullableSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateNegativeTargetingExpressionPredicateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
