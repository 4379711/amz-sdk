package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsBiddingStrategy The bidding strategy. | Value | Strategy name | Description | |----------------|---------------|-------------| | `LEGACY_FOR_SALES` | Dynamic bids - down only | Lowers your bids in real time when your ad may be less likely to convert to a sale. Campaigns created before the release of the bidding controls feature used this setting by default. | | `AUTO_FOR_SALES` | Dynamic bids - up and down | Increases or decreases your bids in real time by a maximum of 100%. With this setting bids increase when your ad is more likely to convert to a sale, and bids decrease when less likely to convert to a sale. | | `MANUAL` | Fixed bid | Uses your exact bid and any placement adjustments you set, and is not subject to dynamic bidding. | | `RULE_BASED` | Rule based bidding | See Rule based bidding documentation https://advertising.amazon.com/API/docs/en-us/sponsored-products/rule-based-bidding/overview |
type SponsoredProductsBiddingStrategy string

// List of SponsoredProductsBiddingStrategy
const (
	SPONSOREDPRODUCTSBIDDINGSTRATEGY_LEGACY_FOR_SALES SponsoredProductsBiddingStrategy = "LEGACY_FOR_SALES"
	SPONSOREDPRODUCTSBIDDINGSTRATEGY_AUTO_FOR_SALES   SponsoredProductsBiddingStrategy = "AUTO_FOR_SALES"
	SPONSOREDPRODUCTSBIDDINGSTRATEGY_MANUAL           SponsoredProductsBiddingStrategy = "MANUAL"
	SPONSOREDPRODUCTSBIDDINGSTRATEGY_RULE_BASED       SponsoredProductsBiddingStrategy = "RULE_BASED"
	SPONSOREDPRODUCTSBIDDINGSTRATEGY_OTHER            SponsoredProductsBiddingStrategy = "OTHER"
)

// All allowed values of SponsoredProductsBiddingStrategy enum
var AllowedSponsoredProductsBiddingStrategyEnumValues = []SponsoredProductsBiddingStrategy{
	"LEGACY_FOR_SALES",
	"AUTO_FOR_SALES",
	"MANUAL",
	"RULE_BASED",
	"OTHER",
}

func (v *SponsoredProductsBiddingStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsBiddingStrategy(value)
	for _, existing := range AllowedSponsoredProductsBiddingStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsBiddingStrategy", value)
}

// NewSponsoredProductsBiddingStrategyFromValue returns a pointer to a valid SponsoredProductsBiddingStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsBiddingStrategyFromValue(v string) (*SponsoredProductsBiddingStrategy, error) {
	ev := SponsoredProductsBiddingStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsBiddingStrategy: valid values are %v", v, AllowedSponsoredProductsBiddingStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsBiddingStrategy) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsBiddingStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsBiddingStrategy value
func (v SponsoredProductsBiddingStrategy) Ptr() *SponsoredProductsBiddingStrategy {
	return &v
}

type NullableSponsoredProductsBiddingStrategy struct {
	value *SponsoredProductsBiddingStrategy
	isSet bool
}

func (v NullableSponsoredProductsBiddingStrategy) Get() *SponsoredProductsBiddingStrategy {
	return v.value
}

func (v *NullableSponsoredProductsBiddingStrategy) Set(val *SponsoredProductsBiddingStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBiddingStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBiddingStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBiddingStrategy(val *SponsoredProductsBiddingStrategy) *NullableSponsoredProductsBiddingStrategy {
	return &NullableSponsoredProductsBiddingStrategy{value: val, isSet: true}
}

func (v NullableSponsoredProductsBiddingStrategy) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBiddingStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
