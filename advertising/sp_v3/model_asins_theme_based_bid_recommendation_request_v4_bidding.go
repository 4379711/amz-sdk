package sp_v3

import "github.com/bytedance/sonic"

// checks if the AsinsThemeBasedBidRecommendationRequestV4Bidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsinsThemeBasedBidRecommendationRequestV4Bidding{}

// AsinsThemeBasedBidRecommendationRequestV4Bidding Bidding control configuration for the campaign.
type AsinsThemeBasedBidRecommendationRequestV4Bidding struct {
	// Placement adjustment configuration for the campaign.
	Adjustments []PlacementAdjustment `json:"adjustments,omitempty"`
	Strategy    BiddingStrategy       `json:"strategy"`
}

type _AsinsThemeBasedBidRecommendationRequestV4Bidding AsinsThemeBasedBidRecommendationRequestV4Bidding

// NewAsinsThemeBasedBidRecommendationRequestV4Bidding instantiates a new AsinsThemeBasedBidRecommendationRequestV4Bidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsinsThemeBasedBidRecommendationRequestV4Bidding(strategy BiddingStrategy) *AsinsThemeBasedBidRecommendationRequestV4Bidding {
	this := AsinsThemeBasedBidRecommendationRequestV4Bidding{}
	this.Strategy = strategy
	return &this
}

// NewAsinsThemeBasedBidRecommendationRequestV4BiddingWithDefaults instantiates a new AsinsThemeBasedBidRecommendationRequestV4Bidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsinsThemeBasedBidRecommendationRequestV4BiddingWithDefaults() *AsinsThemeBasedBidRecommendationRequestV4Bidding {
	this := AsinsThemeBasedBidRecommendationRequestV4Bidding{}
	return &this
}

// GetAdjustments returns the Adjustments field value if set, zero value otherwise.
func (o *AsinsThemeBasedBidRecommendationRequestV4Bidding) GetAdjustments() []PlacementAdjustment {
	if o == nil || IsNil(o.Adjustments) {
		var ret []PlacementAdjustment
		return ret
	}
	return o.Adjustments
}

// GetAdjustmentsOk returns a tuple with the Adjustments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsThemeBasedBidRecommendationRequestV4Bidding) GetAdjustmentsOk() ([]PlacementAdjustment, bool) {
	if o == nil || IsNil(o.Adjustments) {
		return nil, false
	}
	return o.Adjustments, true
}

// HasAdjustments returns a boolean if a field has been set.
func (o *AsinsThemeBasedBidRecommendationRequestV4Bidding) HasAdjustments() bool {
	if o != nil && !IsNil(o.Adjustments) {
		return true
	}

	return false
}

// SetAdjustments gets a reference to the given []PlacementAdjustment and assigns it to the Adjustments field.
func (o *AsinsThemeBasedBidRecommendationRequestV4Bidding) SetAdjustments(v []PlacementAdjustment) {
	o.Adjustments = v
}

// GetStrategy returns the Strategy field value
func (o *AsinsThemeBasedBidRecommendationRequestV4Bidding) GetStrategy() BiddingStrategy {
	if o == nil {
		var ret BiddingStrategy
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *AsinsThemeBasedBidRecommendationRequestV4Bidding) GetStrategyOk() (*BiddingStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *AsinsThemeBasedBidRecommendationRequestV4Bidding) SetStrategy(v BiddingStrategy) {
	o.Strategy = v
}

func (o AsinsThemeBasedBidRecommendationRequestV4Bidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Adjustments) {
		toSerialize["adjustments"] = o.Adjustments
	}
	toSerialize["strategy"] = o.Strategy
	return toSerialize, nil
}

type NullableAsinsThemeBasedBidRecommendationRequestV4Bidding struct {
	value *AsinsThemeBasedBidRecommendationRequestV4Bidding
	isSet bool
}

func (v NullableAsinsThemeBasedBidRecommendationRequestV4Bidding) Get() *AsinsThemeBasedBidRecommendationRequestV4Bidding {
	return v.value
}

func (v *NullableAsinsThemeBasedBidRecommendationRequestV4Bidding) Set(val *AsinsThemeBasedBidRecommendationRequestV4Bidding) {
	v.value = val
	v.isSet = true
}

func (v NullableAsinsThemeBasedBidRecommendationRequestV4Bidding) IsSet() bool {
	return v.isSet
}

func (v *NullableAsinsThemeBasedBidRecommendationRequestV4Bidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsinsThemeBasedBidRecommendationRequestV4Bidding(val *AsinsThemeBasedBidRecommendationRequestV4Bidding) *NullableAsinsThemeBasedBidRecommendationRequestV4Bidding {
	return &NullableAsinsThemeBasedBidRecommendationRequestV4Bidding{value: val, isSet: true}
}

func (v NullableAsinsThemeBasedBidRecommendationRequestV4Bidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsinsThemeBasedBidRecommendationRequestV4Bidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
