package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsNegativeTargetingExpressionPredicateType The type of nagative targeting expression. You can only specify values for the following predicates: | Predicate | Description | | --- | --- | | `ASIN_BRAND_SAME_AS` | Target the brand that is the same as the brand expressed. | | `ASIN_SAME_AS` | Target an ASIN that is the same as the ASIN expressed. | | `OTHER` | Other Type.   |
type SponsoredProductsNegativeTargetingExpressionPredicateType string

// List of SponsoredProductsNegativeTargetingExpressionPredicateType
const (
	SPONSOREDPRODUCTSNEGATIVETARGETINGEXPRESSIONPREDICATETYPE_ASIN_BRAND_SAME_AS SponsoredProductsNegativeTargetingExpressionPredicateType = "ASIN_BRAND_SAME_AS"
	SPONSOREDPRODUCTSNEGATIVETARGETINGEXPRESSIONPREDICATETYPE_ASIN_SAME_AS       SponsoredProductsNegativeTargetingExpressionPredicateType = "ASIN_SAME_AS"
	SPONSOREDPRODUCTSNEGATIVETARGETINGEXPRESSIONPREDICATETYPE_OTHER              SponsoredProductsNegativeTargetingExpressionPredicateType = "OTHER"
)

// All allowed values of SponsoredProductsNegativeTargetingExpressionPredicateType enum
var AllowedSponsoredProductsNegativeTargetingExpressionPredicateTypeEnumValues = []SponsoredProductsNegativeTargetingExpressionPredicateType{
	"ASIN_BRAND_SAME_AS",
	"ASIN_SAME_AS",
	"OTHER",
}

func (v *SponsoredProductsNegativeTargetingExpressionPredicateType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsNegativeTargetingExpressionPredicateType(value)
	for _, existing := range AllowedSponsoredProductsNegativeTargetingExpressionPredicateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsNegativeTargetingExpressionPredicateType", value)
}

// NewSponsoredProductsNegativeTargetingExpressionPredicateTypeFromValue returns a pointer to a valid SponsoredProductsNegativeTargetingExpressionPredicateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsNegativeTargetingExpressionPredicateTypeFromValue(v string) (*SponsoredProductsNegativeTargetingExpressionPredicateType, error) {
	ev := SponsoredProductsNegativeTargetingExpressionPredicateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsNegativeTargetingExpressionPredicateType: valid values are %v", v, AllowedSponsoredProductsNegativeTargetingExpressionPredicateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsNegativeTargetingExpressionPredicateType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsNegativeTargetingExpressionPredicateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsNegativeTargetingExpressionPredicateType value
func (v SponsoredProductsNegativeTargetingExpressionPredicateType) Ptr() *SponsoredProductsNegativeTargetingExpressionPredicateType {
	return &v
}

type NullableSponsoredProductsNegativeTargetingExpressionPredicateType struct {
	value *SponsoredProductsNegativeTargetingExpressionPredicateType
	isSet bool
}

func (v NullableSponsoredProductsNegativeTargetingExpressionPredicateType) Get() *SponsoredProductsNegativeTargetingExpressionPredicateType {
	return v.value
}

func (v *NullableSponsoredProductsNegativeTargetingExpressionPredicateType) Set(val *SponsoredProductsNegativeTargetingExpressionPredicateType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeTargetingExpressionPredicateType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeTargetingExpressionPredicateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeTargetingExpressionPredicateType(val *SponsoredProductsNegativeTargetingExpressionPredicateType) *NullableSponsoredProductsNegativeTargetingExpressionPredicateType {
	return &NullableSponsoredProductsNegativeTargetingExpressionPredicateType{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeTargetingExpressionPredicateType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeTargetingExpressionPredicateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
