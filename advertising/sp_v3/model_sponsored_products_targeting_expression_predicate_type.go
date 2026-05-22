package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsTargetingExpressionPredicateType The type of targeting expression. You can specify values for the following predicates: | Predicate | Description | | --- | --- | | `QUERY_BROAD_REL_MATCHES` | Auto Targeting - cannot be manually created - corresponds to the `Loose match` target type in the UI, this will show your ad to shoppers who use search terms loosely related to your products.| | `QUERY_HIGH_REL_MATCHES` | Auto Targeting - cannot be manually created - corresponds to the `Close match` target type in the UI, this will show your ad to shoppers who use search terms closely related to your products.| | `ASIN_ACCESSORY_RELATED` | Auto Targeting - cannot be manually created - corresponds to the `Complements` target type in the UI, this will show your ad to shoppers who view the detail pages of products that complement your product.| | `ASIN_SUBSTITUTE_RELATED` | Auto Targeting - cannot be manually created - corresponds to the `Substitutes` target type in the UI, this will show your ad to shoppers who use detail pages of products similar to yours.| | `ASIN_CATEGORY_SAME_AS` | Target the category that is the same as the category expressed | | `ASIN_BRAND_SAME_AS` | Target the brand that is the same as the brand expressed. | | `ASIN_PRICE_LESS_THAN` | Target a price that is less than the price expressed. | | `ASIN_PRICE_BETWEEN` | Target a price that is between the prices expressed. | | `ASIN_PRICE_GREATER_THAN` | Target a price that is greater than the price expressed. | | `ASIN_REVIEW_RATING_LESS_THAN` | Target a review rating less than the review rating that is expressed. | | `ASIN_REVIEW_RATING_BETWEEN` | Target a review rating that is between the review ratings expressed. | | `ASIN_REVIEW_RATING_GREATER_THAN` | Target a review rating that is greater than the review rating expressed. | | `ASIN_SAME_AS` | Target an ASIN that is the same as the ASIN expressed. | | `ASIN_IS_PRIME_SHIPPING_ELIGIBLE` | Target products that are Prime Shipping Eligible. This refinement can be applied at a category or brand level only. | | `ASIN_AGE_RANGE_SAME_AS` | Target an age range that is in the expressed range. This refinement can be applied for toys and games categories only. | | `ASIN_GENRE_SAME_AS` | Target products related to the expressed genre. This refinement can be applied for Books and eBooks categories only.   | | `ASIN_EXPANDED_FROM` | Target products similar in performance to the ASIN expressed.   | | `KEYWORD_GROUP_SAME_AS` | Target the keyword group that is the same as the keyword group expressed (Beta coming soon). | | `OTHER` | Other Type.   |
type SponsoredProductsTargetingExpressionPredicateType string

// List of SponsoredProductsTargetingExpressionPredicateType
const (
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_CATEGORY_SAME_AS           SponsoredProductsTargetingExpressionPredicateType = "ASIN_CATEGORY_SAME_AS"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_BRAND_SAME_AS              SponsoredProductsTargetingExpressionPredicateType = "ASIN_BRAND_SAME_AS"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_PRICE_LESS_THAN            SponsoredProductsTargetingExpressionPredicateType = "ASIN_PRICE_LESS_THAN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_PRICE_BETWEEN              SponsoredProductsTargetingExpressionPredicateType = "ASIN_PRICE_BETWEEN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_PRICE_GREATER_THAN         SponsoredProductsTargetingExpressionPredicateType = "ASIN_PRICE_GREATER_THAN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_REVIEW_RATING_LESS_THAN    SponsoredProductsTargetingExpressionPredicateType = "ASIN_REVIEW_RATING_LESS_THAN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_REVIEW_RATING_BETWEEN      SponsoredProductsTargetingExpressionPredicateType = "ASIN_REVIEW_RATING_BETWEEN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_REVIEW_RATING_GREATER_THAN SponsoredProductsTargetingExpressionPredicateType = "ASIN_REVIEW_RATING_GREATER_THAN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_SAME_AS                    SponsoredProductsTargetingExpressionPredicateType = "ASIN_SAME_AS"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_QUERY_BROAD_REL_MATCHES         SponsoredProductsTargetingExpressionPredicateType = "QUERY_BROAD_REL_MATCHES"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_QUERY_HIGH_REL_MATCHES          SponsoredProductsTargetingExpressionPredicateType = "QUERY_HIGH_REL_MATCHES"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_SUBSTITUTE_RELATED         SponsoredProductsTargetingExpressionPredicateType = "ASIN_SUBSTITUTE_RELATED"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_ACCESSORY_RELATED          SponsoredProductsTargetingExpressionPredicateType = "ASIN_ACCESSORY_RELATED"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_AGE_RANGE_SAME_AS          SponsoredProductsTargetingExpressionPredicateType = "ASIN_AGE_RANGE_SAME_AS"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_GENRE_SAME_AS              SponsoredProductsTargetingExpressionPredicateType = "ASIN_GENRE_SAME_AS"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_IS_PRIME_SHIPPING_ELIGIBLE SponsoredProductsTargetingExpressionPredicateType = "ASIN_IS_PRIME_SHIPPING_ELIGIBLE"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_ASIN_EXPANDED_FROM              SponsoredProductsTargetingExpressionPredicateType = "ASIN_EXPANDED_FROM"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_KEYWORD_GROUP_SAME_AS           SponsoredProductsTargetingExpressionPredicateType = "KEYWORD_GROUP_SAME_AS"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPE_OTHER                           SponsoredProductsTargetingExpressionPredicateType = "OTHER"
)

// All allowed values of SponsoredProductsTargetingExpressionPredicateType enum
var AllowedSponsoredProductsTargetingExpressionPredicateTypeEnumValues = []SponsoredProductsTargetingExpressionPredicateType{
	"ASIN_CATEGORY_SAME_AS",
	"ASIN_BRAND_SAME_AS",
	"ASIN_PRICE_LESS_THAN",
	"ASIN_PRICE_BETWEEN",
	"ASIN_PRICE_GREATER_THAN",
	"ASIN_REVIEW_RATING_LESS_THAN",
	"ASIN_REVIEW_RATING_BETWEEN",
	"ASIN_REVIEW_RATING_GREATER_THAN",
	"ASIN_SAME_AS",
	"QUERY_BROAD_REL_MATCHES",
	"QUERY_HIGH_REL_MATCHES",
	"ASIN_SUBSTITUTE_RELATED",
	"ASIN_ACCESSORY_RELATED",
	"ASIN_AGE_RANGE_SAME_AS",
	"ASIN_GENRE_SAME_AS",
	"ASIN_IS_PRIME_SHIPPING_ELIGIBLE",
	"ASIN_EXPANDED_FROM",
	"KEYWORD_GROUP_SAME_AS",
	"OTHER",
}

func (v *SponsoredProductsTargetingExpressionPredicateType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsTargetingExpressionPredicateType(value)
	for _, existing := range AllowedSponsoredProductsTargetingExpressionPredicateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsTargetingExpressionPredicateType", value)
}

// NewSponsoredProductsTargetingExpressionPredicateTypeFromValue returns a pointer to a valid SponsoredProductsTargetingExpressionPredicateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsTargetingExpressionPredicateTypeFromValue(v string) (*SponsoredProductsTargetingExpressionPredicateType, error) {
	ev := SponsoredProductsTargetingExpressionPredicateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsTargetingExpressionPredicateType: valid values are %v", v, AllowedSponsoredProductsTargetingExpressionPredicateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsTargetingExpressionPredicateType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsTargetingExpressionPredicateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsTargetingExpressionPredicateType value
func (v SponsoredProductsTargetingExpressionPredicateType) Ptr() *SponsoredProductsTargetingExpressionPredicateType {
	return &v
}

type NullableSponsoredProductsTargetingExpressionPredicateType struct {
	value *SponsoredProductsTargetingExpressionPredicateType
	isSet bool
}

func (v NullableSponsoredProductsTargetingExpressionPredicateType) Get() *SponsoredProductsTargetingExpressionPredicateType {
	return v.value
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateType) Set(val *SponsoredProductsTargetingExpressionPredicateType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetingExpressionPredicateType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetingExpressionPredicateType(val *SponsoredProductsTargetingExpressionPredicateType) *NullableSponsoredProductsTargetingExpressionPredicateType {
	return &NullableSponsoredProductsTargetingExpressionPredicateType{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetingExpressionPredicateType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
