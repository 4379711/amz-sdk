package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// BidOptimizationStrategy DEPRECATED [PLANNED SHUTOFF DATE 3/31/2024] **Note: This feature has been deprecated and planned to shutoff on 03/31/2024. After the shut off date, we will ignore this field in the request and treat it as null. You will not get an error if you supply this field in the request. Based on customer feedback, we are rethinking this feature in context to Goal based campaigns to help advertiser reach NTB customers at scale with transparent reporting. Meanwhile, if you have any feedback or suggestion related to this feature then please reach out to our customer support teams.  The bid optimization strategy. - MAXIMIZE_IMMEDIATE_SALES - The default bidding strategy. The campaign is optimized to maximize sale. - MAXIMIZE_NEW_TO_BRAND_CUSTOMERS - The campaign is optimized to acquire more new-to-brand customers.
type BidOptimizationStrategy string

// List of BidOptimizationStrategy
const (
	BIDOPTIMIZATIONSTRATEGY_IMMEDIATE_SALES        BidOptimizationStrategy = "MAXIMIZE_IMMEDIATE_SALES"
	BIDOPTIMIZATIONSTRATEGY_NEW_TO_BRAND_CUSTOMERS BidOptimizationStrategy = "MAXIMIZE_NEW_TO_BRAND_CUSTOMERS"
)

// All allowed values of BidOptimizationStrategy enum
var AllowedBidOptimizationStrategyEnumValues = []BidOptimizationStrategy{
	"MAXIMIZE_IMMEDIATE_SALES",
	"MAXIMIZE_NEW_TO_BRAND_CUSTOMERS",
}

func (v *BidOptimizationStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BidOptimizationStrategy(value)
	for _, existing := range AllowedBidOptimizationStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BidOptimizationStrategy", value)
}

// NewBidOptimizationStrategyFromValue returns a pointer to a valid BidOptimizationStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBidOptimizationStrategyFromValue(v string) (*BidOptimizationStrategy, error) {
	ev := BidOptimizationStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BidOptimizationStrategy: valid values are %v", v, AllowedBidOptimizationStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BidOptimizationStrategy) IsValid() bool {
	for _, existing := range AllowedBidOptimizationStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BidOptimizationStrategy value
func (v BidOptimizationStrategy) Ptr() *BidOptimizationStrategy {
	return &v
}

type NullableBidOptimizationStrategy struct {
	value *BidOptimizationStrategy
	isSet bool
}

func (v NullableBidOptimizationStrategy) Get() *BidOptimizationStrategy {
	return v.value
}

func (v *NullableBidOptimizationStrategy) Set(val *BidOptimizationStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableBidOptimizationStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableBidOptimizationStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidOptimizationStrategy(val *BidOptimizationStrategy) *NullableBidOptimizationStrategy {
	return &NullableBidOptimizationStrategy{value: val, isSet: true}
}

func (v NullableBidOptimizationStrategy) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidOptimizationStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
