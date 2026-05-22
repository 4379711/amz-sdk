package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsCreateOrUpdateBiddingStrategy The bidding strategy. `strategy` is required for create requests if dynamicBidding is provided, but is optional for update requests. | Value | Strategy name | Description | |----------------|---------------|-------------| | `LEGACY_FOR_SALES` | Dynamic bids - down only | Lowers your bids in real time when your ad may be less likely to convert to a sale. Campaigns created before the release of the bidding controls feature used this setting by default. | | `AUTO_FOR_SALES` | Dynamic bids - up and down | Increases or decreases your bids in real time by a maximum of 100%. With this setting bids increase when your ad is more likely to convert to a sale, and bids decrease when less likely to convert to a sale. | | `MANUAL` | Fixed bid | Uses your exact bid and any placement adjustments you set, and is not subject to dynamic bidding. | | `RULE_BASED` | Rule based bidding | See Rule based bidding documentation https://advertising.amazon.com/API/docs/en-us/sponsored-products/rule-based-bidding/overview |
type SponsoredProductsCreateOrUpdateBiddingStrategy string

// List of SponsoredProductsCreateOrUpdateBiddingStrategy
const (
	SPONSOREDPRODUCTSCREATEORUPDATEBIDDINGSTRATEGY_LEGACY_FOR_SALES SponsoredProductsCreateOrUpdateBiddingStrategy = "LEGACY_FOR_SALES"
	SPONSOREDPRODUCTSCREATEORUPDATEBIDDINGSTRATEGY_AUTO_FOR_SALES   SponsoredProductsCreateOrUpdateBiddingStrategy = "AUTO_FOR_SALES"
	SPONSOREDPRODUCTSCREATEORUPDATEBIDDINGSTRATEGY_MANUAL           SponsoredProductsCreateOrUpdateBiddingStrategy = "MANUAL"
	SPONSOREDPRODUCTSCREATEORUPDATEBIDDINGSTRATEGY_RULE_BASED       SponsoredProductsCreateOrUpdateBiddingStrategy = "RULE_BASED"
)

// All allowed values of SponsoredProductsCreateOrUpdateBiddingStrategy enum
var AllowedSponsoredProductsCreateOrUpdateBiddingStrategyEnumValues = []SponsoredProductsCreateOrUpdateBiddingStrategy{
	"LEGACY_FOR_SALES",
	"AUTO_FOR_SALES",
	"MANUAL",
	"RULE_BASED",
}

func (v *SponsoredProductsCreateOrUpdateBiddingStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCreateOrUpdateBiddingStrategy(value)
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateBiddingStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCreateOrUpdateBiddingStrategy", value)
}

// NewSponsoredProductsCreateOrUpdateBiddingStrategyFromValue returns a pointer to a valid SponsoredProductsCreateOrUpdateBiddingStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCreateOrUpdateBiddingStrategyFromValue(v string) (*SponsoredProductsCreateOrUpdateBiddingStrategy, error) {
	ev := SponsoredProductsCreateOrUpdateBiddingStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCreateOrUpdateBiddingStrategy: valid values are %v", v, AllowedSponsoredProductsCreateOrUpdateBiddingStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCreateOrUpdateBiddingStrategy) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateBiddingStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCreateOrUpdateBiddingStrategy value
func (v SponsoredProductsCreateOrUpdateBiddingStrategy) Ptr() *SponsoredProductsCreateOrUpdateBiddingStrategy {
	return &v
}

type NullableSponsoredProductsCreateOrUpdateBiddingStrategy struct {
	value *SponsoredProductsCreateOrUpdateBiddingStrategy
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateBiddingStrategy) Get() *SponsoredProductsCreateOrUpdateBiddingStrategy {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateBiddingStrategy) Set(val *SponsoredProductsCreateOrUpdateBiddingStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateBiddingStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateBiddingStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateBiddingStrategy(val *SponsoredProductsCreateOrUpdateBiddingStrategy) *NullableSponsoredProductsCreateOrUpdateBiddingStrategy {
	return &NullableSponsoredProductsCreateOrUpdateBiddingStrategy{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateBiddingStrategy) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateBiddingStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
