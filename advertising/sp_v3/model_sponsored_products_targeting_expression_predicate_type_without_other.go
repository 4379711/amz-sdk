package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsTargetingExpressionPredicateTypeWithoutOther The type of targeting expression. You can specify values for the following predicates: | Predicate | Description | | --- | --- | | `QUERY_BROAD_REL_MATCHES` | Auto Targeting - cannot be manually created - corresponds to the `Loose match` target type in the UI, this will show your ad to shoppers who use search terms loosely related to your products.| | `QUERY_HIGH_REL_MATCHES` | Auto Targeting - cannot be manually created - corresponds to the `Close match` target type in the UI, this will show your ad to shoppers who use search terms closely related to your products.| | `ASIN_ACCESSORY_RELATED` | Auto Targeting - cannot be manually created - corresponds to the `Complements` target type in the UI, this will show your ad to shoppers who view the detail pages of products that complement your product.| | `ASIN_SUBSTITUTE_RELATED` | Auto Targeting - cannot be manually created - corresponds to the `Substitutes` target type in the UI, this will show your ad to shoppers who use detail pages of products similar to yours.| | `ASIN_CATEGORY_SAME_AS` | Target the category that is the same as the category expressed | | `ASIN_BRAND_SAME_AS` | Target the brand that is the same as the brand expressed. | | `ASIN_PRICE_LESS_THAN` | Target a price that is less than the price expressed. | | `ASIN_PRICE_BETWEEN` | Target a price that is between the prices expressed. | | `ASIN_PRICE_GREATER_THAN` | Target a price that is greater than the price expressed. | | `ASIN_REVIEW_RATING_LESS_THAN` | Target a review rating less than the review rating that is expressed. | | `ASIN_REVIEW_RATING_BETWEEN` | Target a review rating that is between the review ratings expressed. | | `ASIN_REVIEW_RATING_GREATER_THAN` | Target a review rating that is greater than the review rating expressed. | | `ASIN_SAME_AS` | Target an ASIN that is the same as the ASIN expressed. | | `ASIN_IS_PRIME_SHIPPING_ELIGIBLE` | Target products that are Prime Shipping Eligible. This refinement can be applied at a category or brand level only. | | `ASIN_AGE_RANGE_SAME_AS` | Target an age range that is in the expressed range. This refinement can be applied for toys and games categories only. | | `ASIN_GENRE_SAME_AS` | Target products related to the expressed genre. This refinement can be applied for Books and eBooks categories only.   | | `ASIN_EXPANDED_FROM` | Target products similar in performance to the ASIN expressed.   | | `KEYWORD_GROUP_SAME_AS` | Target the keyword group that is the same as the keyword group expressed (Beta coming soon). | | `OTHER` | Other Type.   |
type SponsoredProductsTargetingExpressionPredicateTypeWithoutOther string

// List of SponsoredProductsTargetingExpressionPredicateTypeWithoutOther
const (
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_CATEGORY_SAME_AS           SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_CATEGORY_SAME_AS"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_BRAND_SAME_AS              SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_BRAND_SAME_AS"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_PRICE_LESS_THAN            SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_PRICE_LESS_THAN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_PRICE_BETWEEN              SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_PRICE_BETWEEN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_PRICE_GREATER_THAN         SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_PRICE_GREATER_THAN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_REVIEW_RATING_LESS_THAN    SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_REVIEW_RATING_LESS_THAN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_REVIEW_RATING_BETWEEN      SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_REVIEW_RATING_BETWEEN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_REVIEW_RATING_GREATER_THAN SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_REVIEW_RATING_GREATER_THAN"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_SAME_AS                    SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_SAME_AS"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_QUERY_BROAD_REL_MATCHES         SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "QUERY_BROAD_REL_MATCHES"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_QUERY_HIGH_REL_MATCHES          SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "QUERY_HIGH_REL_MATCHES"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_SUBSTITUTE_RELATED         SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_SUBSTITUTE_RELATED"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_ACCESSORY_RELATED          SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_ACCESSORY_RELATED"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_AGE_RANGE_SAME_AS          SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_AGE_RANGE_SAME_AS"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_GENRE_SAME_AS              SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_GENRE_SAME_AS"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_IS_PRIME_SHIPPING_ELIGIBLE SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_IS_PRIME_SHIPPING_ELIGIBLE"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_ASIN_EXPANDED_FROM              SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "ASIN_EXPANDED_FROM"
	SPONSOREDPRODUCTSTARGETINGEXPRESSIONPREDICATETYPEWITHOUTOTHER_KEYWORD_GROUP_SAME_AS           SponsoredProductsTargetingExpressionPredicateTypeWithoutOther = "KEYWORD_GROUP_SAME_AS"
)

// All allowed values of SponsoredProductsTargetingExpressionPredicateTypeWithoutOther enum
var AllowedSponsoredProductsTargetingExpressionPredicateTypeWithoutOtherEnumValues = []SponsoredProductsTargetingExpressionPredicateTypeWithoutOther{
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
}

func (v *SponsoredProductsTargetingExpressionPredicateTypeWithoutOther) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsTargetingExpressionPredicateTypeWithoutOther(value)
	for _, existing := range AllowedSponsoredProductsTargetingExpressionPredicateTypeWithoutOtherEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsTargetingExpressionPredicateTypeWithoutOther", value)
}

// NewSponsoredProductsTargetingExpressionPredicateTypeWithoutOtherFromValue returns a pointer to a valid SponsoredProductsTargetingExpressionPredicateTypeWithoutOther
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsTargetingExpressionPredicateTypeWithoutOtherFromValue(v string) (*SponsoredProductsTargetingExpressionPredicateTypeWithoutOther, error) {
	ev := SponsoredProductsTargetingExpressionPredicateTypeWithoutOther(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsTargetingExpressionPredicateTypeWithoutOther: valid values are %v", v, AllowedSponsoredProductsTargetingExpressionPredicateTypeWithoutOtherEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsTargetingExpressionPredicateTypeWithoutOther) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsTargetingExpressionPredicateTypeWithoutOtherEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsTargetingExpressionPredicateTypeWithoutOther value
func (v SponsoredProductsTargetingExpressionPredicateTypeWithoutOther) Ptr() *SponsoredProductsTargetingExpressionPredicateTypeWithoutOther {
	return &v
}

type NullableSponsoredProductsTargetingExpressionPredicateTypeWithoutOther struct {
	value *SponsoredProductsTargetingExpressionPredicateTypeWithoutOther
	isSet bool
}

func (v NullableSponsoredProductsTargetingExpressionPredicateTypeWithoutOther) Get() *SponsoredProductsTargetingExpressionPredicateTypeWithoutOther {
	return v.value
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateTypeWithoutOther) Set(val *SponsoredProductsTargetingExpressionPredicateTypeWithoutOther) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetingExpressionPredicateTypeWithoutOther) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateTypeWithoutOther) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetingExpressionPredicateTypeWithoutOther(val *SponsoredProductsTargetingExpressionPredicateTypeWithoutOther) *NullableSponsoredProductsTargetingExpressionPredicateTypeWithoutOther {
	return &NullableSponsoredProductsTargetingExpressionPredicateTypeWithoutOther{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetingExpressionPredicateTypeWithoutOther) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetingExpressionPredicateTypeWithoutOther) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
