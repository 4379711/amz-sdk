package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// BiddingStrategy The bidding strategy selected for the campaign. Use `LEGACY_FOR_SALES` to lower your bid in real time when your ad may be less likely to convert to a sale. Use `AUTO_FOR_SALES` to increase your bid in real time when your ad may be more likely to convert to a sale or lower your bid when less likely to convert to a sale. Use `MANUAL` to use your exact bid along with any manual adjustments.
type BiddingStrategy string

// List of BiddingStrategy
const (
	BIDDINGSTRATEGY_LEGACY_FOR_SALES BiddingStrategy = "LEGACY_FOR_SALES"
	BIDDINGSTRATEGY_AUTO_FOR_SALES   BiddingStrategy = "AUTO_FOR_SALES"
	BIDDINGSTRATEGY_MANUAL           BiddingStrategy = "MANUAL"
	BIDDINGSTRATEGY_RULE_BASED       BiddingStrategy = "RULE_BASED"
)

// All allowed values of BiddingStrategy enum
var AllowedBiddingStrategyEnumValues = []BiddingStrategy{
	"LEGACY_FOR_SALES",
	"AUTO_FOR_SALES",
	"MANUAL",
	"RULE_BASED",
}

func (v *BiddingStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BiddingStrategy(value)
	for _, existing := range AllowedBiddingStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BiddingStrategy", value)
}

// NewBiddingStrategyFromValue returns a pointer to a valid BiddingStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBiddingStrategyFromValue(v string) (*BiddingStrategy, error) {
	ev := BiddingStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BiddingStrategy: valid values are %v", v, AllowedBiddingStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BiddingStrategy) IsValid() bool {
	for _, existing := range AllowedBiddingStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BiddingStrategy value
func (v BiddingStrategy) Ptr() *BiddingStrategy {
	return &v
}

type NullableBiddingStrategy struct {
	value *BiddingStrategy
	isSet bool
}

func (v NullableBiddingStrategy) Get() *BiddingStrategy {
	return v.value
}

func (v *NullableBiddingStrategy) Set(val *BiddingStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableBiddingStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableBiddingStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiddingStrategy(val *BiddingStrategy) *NullableBiddingStrategy {
	return &NullableBiddingStrategy{value: val, isSet: true}
}

func (v NullableBiddingStrategy) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBiddingStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
