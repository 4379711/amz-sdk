package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the MultiCountryAsinsThemeBasedBidRecommendationRequestBidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiCountryAsinsThemeBasedBidRecommendationRequestBidding{}

// MultiCountryAsinsThemeBasedBidRecommendationRequestBidding Bidding control configuration for the campaign.
type MultiCountryAsinsThemeBasedBidRecommendationRequestBidding struct {
	// Placement adjustment configuration for the campaign.
	Adjustments []BidPlacementAdjustment `json:"adjustments,omitempty"`
	Strategy    BiddingStrategy          `json:"strategy"`
}

type _MultiCountryAsinsThemeBasedBidRecommendationRequestBidding MultiCountryAsinsThemeBasedBidRecommendationRequestBidding

// NewMultiCountryAsinsThemeBasedBidRecommendationRequestBidding instantiates a new MultiCountryAsinsThemeBasedBidRecommendationRequestBidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiCountryAsinsThemeBasedBidRecommendationRequestBidding(strategy BiddingStrategy) *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding {
	this := MultiCountryAsinsThemeBasedBidRecommendationRequestBidding{}
	this.Strategy = strategy
	return &this
}

// NewMultiCountryAsinsThemeBasedBidRecommendationRequestBiddingWithDefaults instantiates a new MultiCountryAsinsThemeBasedBidRecommendationRequestBidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiCountryAsinsThemeBasedBidRecommendationRequestBiddingWithDefaults() *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding {
	this := MultiCountryAsinsThemeBasedBidRecommendationRequestBidding{}
	return &this
}

// GetAdjustments returns the Adjustments field value if set, zero value otherwise.
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding) GetAdjustments() []BidPlacementAdjustment {
	if o == nil || IsNil(o.Adjustments) {
		var ret []BidPlacementAdjustment
		return ret
	}
	return o.Adjustments
}

// GetAdjustmentsOk returns a tuple with the Adjustments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding) GetAdjustmentsOk() ([]BidPlacementAdjustment, bool) {
	if o == nil || IsNil(o.Adjustments) {
		return nil, false
	}
	return o.Adjustments, true
}

// HasAdjustments returns a boolean if a field has been set.
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding) HasAdjustments() bool {
	if o != nil && !IsNil(o.Adjustments) {
		return true
	}

	return false
}

// SetAdjustments gets a reference to the given []BidPlacementAdjustment and assigns it to the Adjustments field.
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding) SetAdjustments(v []BidPlacementAdjustment) {
	o.Adjustments = v
}

// GetStrategy returns the Strategy field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding) GetStrategy() BiddingStrategy {
	if o == nil {
		var ret BiddingStrategy
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding) GetStrategyOk() (*BiddingStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding) SetStrategy(v BiddingStrategy) {
	o.Strategy = v
}

func (o MultiCountryAsinsThemeBasedBidRecommendationRequestBidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Adjustments) {
		toSerialize["adjustments"] = o.Adjustments
	}
	toSerialize["strategy"] = o.Strategy
	return toSerialize, nil
}

type NullableMultiCountryAsinsThemeBasedBidRecommendationRequestBidding struct {
	value *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding
	isSet bool
}

func (v NullableMultiCountryAsinsThemeBasedBidRecommendationRequestBidding) Get() *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding {
	return v.value
}

func (v *NullableMultiCountryAsinsThemeBasedBidRecommendationRequestBidding) Set(val *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiCountryAsinsThemeBasedBidRecommendationRequestBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiCountryAsinsThemeBasedBidRecommendationRequestBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiCountryAsinsThemeBasedBidRecommendationRequestBidding(val *MultiCountryAsinsThemeBasedBidRecommendationRequestBidding) *NullableMultiCountryAsinsThemeBasedBidRecommendationRequestBidding {
	return &NullableMultiCountryAsinsThemeBasedBidRecommendationRequestBidding{value: val, isSet: true}
}

func (v NullableMultiCountryAsinsThemeBasedBidRecommendationRequestBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMultiCountryAsinsThemeBasedBidRecommendationRequestBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
