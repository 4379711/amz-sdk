package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsCreateTargetingExpressionPredicateType The type of targeting expression. You can specify values for the following predicates: | Predicate | Description | | --- | --- | | `ASIN_CATEGORY_SAME_AS` | Target the category that is the same as the category expressed. | | `ASIN_BRAND_SAME_AS` | Target the brand that is the same as the brand expressed. | | `ASIN_PRICE_LESS_THAN` | Target a price that is less than the price expressed. | | `ASIN_PRICE_BETWEEN` | Target a price that is between the prices expressed. | | `ASIN_PRICE_GREATER_THAN` | Target a price that is greater than the price expressed. | | `ASIN_REVIEW_RATING_LESS_THAN` | Target a review rating less than the review rating that is expressed. | | `ASIN_REVIEW_RATING_BETWEEN` | Target a review rating that is between the review ratings expressed. | | `ASIN_REVIEW_RATING_GREATER_THAN` | Target a review rating that is greater than the review rating expressed. | | `ASIN_SAME_AS` | Target an ASIN that is the same as the ASIN expressed. | | `ASIN_IS_PRIME_SHIPPING_ELIGIBLE` | Target products that are Prime Shipping Eligible. This refinement can be applied at a category or brand level only. | | `ASIN_AGE_RANGE_SAME_AS` | Target an age range that is in the expressed range. This refinement can be applied for toys and games categories only. | | `ASIN_GENRE_SAME_AS` | Target products related to the expressed genre. This refinement can be applied for Books and eBooks categories only.   | | `ASIN_EXPANDED_FROM` | Target products similar in performance to the ASIN expressed.   | | `KEYWORD_GROUP_SAME_AS` | Target the keyword group that is the same as the keyword group expressed (Beta coming soon). |
type SponsoredProductsCreateTargetingExpressionPredicateType string

// List of SponsoredProductsCreateTargetingExpressionPredicateType
const (
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_CATEGORY_SAME_AS           SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_CATEGORY_SAME_AS"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_BRAND_SAME_AS              SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_BRAND_SAME_AS"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_PRICE_LESS_THAN            SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_PRICE_LESS_THAN"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_PRICE_BETWEEN              SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_PRICE_BETWEEN"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_PRICE_GREATER_THAN         SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_PRICE_GREATER_THAN"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_REVIEW_RATING_LESS_THAN    SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_REVIEW_RATING_LESS_THAN"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_REVIEW_RATING_BETWEEN      SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_REVIEW_RATING_BETWEEN"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_REVIEW_RATING_GREATER_THAN SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_REVIEW_RATING_GREATER_THAN"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_SAME_AS                    SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_SAME_AS"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_AGE_RANGE_SAME_AS          SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_AGE_RANGE_SAME_AS"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_GENRE_SAME_AS              SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_GENRE_SAME_AS"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_IS_PRIME_SHIPPING_ELIGIBLE SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_IS_PRIME_SHIPPING_ELIGIBLE"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_ASIN_EXPANDED_FROM              SponsoredProductsCreateTargetingExpressionPredicateType = "ASIN_EXPANDED_FROM"
	SPONSOREDPRODUCTSCREATETARGETINGEXPRESSIONPREDICATETYPE_KEYWORD_GROUP_SAME_AS           SponsoredProductsCreateTargetingExpressionPredicateType = "KEYWORD_GROUP_SAME_AS"
)

// All allowed values of SponsoredProductsCreateTargetingExpressionPredicateType enum
var AllowedSponsoredProductsCreateTargetingExpressionPredicateTypeEnumValues = []SponsoredProductsCreateTargetingExpressionPredicateType{
	"ASIN_CATEGORY_SAME_AS",
	"ASIN_BRAND_SAME_AS",
	"ASIN_PRICE_LESS_THAN",
	"ASIN_PRICE_BETWEEN",
	"ASIN_PRICE_GREATER_THAN",
	"ASIN_REVIEW_RATING_LESS_THAN",
	"ASIN_REVIEW_RATING_BETWEEN",
	"ASIN_REVIEW_RATING_GREATER_THAN",
	"ASIN_SAME_AS",
	"ASIN_AGE_RANGE_SAME_AS",
	"ASIN_GENRE_SAME_AS",
	"ASIN_IS_PRIME_SHIPPING_ELIGIBLE",
	"ASIN_EXPANDED_FROM",
	"KEYWORD_GROUP_SAME_AS",
}

func (v *SponsoredProductsCreateTargetingExpressionPredicateType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCreateTargetingExpressionPredicateType(value)
	for _, existing := range AllowedSponsoredProductsCreateTargetingExpressionPredicateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCreateTargetingExpressionPredicateType", value)
}

// NewSponsoredProductsCreateTargetingExpressionPredicateTypeFromValue returns a pointer to a valid SponsoredProductsCreateTargetingExpressionPredicateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCreateTargetingExpressionPredicateTypeFromValue(v string) (*SponsoredProductsCreateTargetingExpressionPredicateType, error) {
	ev := SponsoredProductsCreateTargetingExpressionPredicateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCreateTargetingExpressionPredicateType: valid values are %v", v, AllowedSponsoredProductsCreateTargetingExpressionPredicateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCreateTargetingExpressionPredicateType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCreateTargetingExpressionPredicateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCreateTargetingExpressionPredicateType value
func (v SponsoredProductsCreateTargetingExpressionPredicateType) Ptr() *SponsoredProductsCreateTargetingExpressionPredicateType {
	return &v
}

type NullableSponsoredProductsCreateTargetingExpressionPredicateType struct {
	value *SponsoredProductsCreateTargetingExpressionPredicateType
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetingExpressionPredicateType) Get() *SponsoredProductsCreateTargetingExpressionPredicateType {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetingExpressionPredicateType) Set(val *SponsoredProductsCreateTargetingExpressionPredicateType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetingExpressionPredicateType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetingExpressionPredicateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetingExpressionPredicateType(val *SponsoredProductsCreateTargetingExpressionPredicateType) *NullableSponsoredProductsCreateTargetingExpressionPredicateType {
	return &NullableSponsoredProductsCreateTargetingExpressionPredicateType{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetingExpressionPredicateType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetingExpressionPredicateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
