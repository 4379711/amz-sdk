package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BiddingStrategyRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BiddingStrategyRecommendation{}

// BiddingStrategyRecommendation Contains suggested recommendation for the campaign bidding strategy.
type BiddingStrategyRecommendation struct {
	// The suggested bidding strategy value for the campaign. | Value | Strategy name | Description | |----------------|---------------|-------------| | `LEGACY_FOR_SALES` | Dynamic bids - down only | Lowers your bids in real time when your ad may be less likely to convert to a sale. Campaigns created before the release of the bidding controls feature used this setting by default. | | `AUTO_FOR_SALES` | Dynamic bids - up and down | Increases or decreases your bids in real time by a maximum of 100%. With this setting bids increase when your ad is more likely to convert to a sale, and bids decrease when less likely to convert to a sale. | | `MANUAL` | Fixed bid | Uses your exact bid and any placement adjustments you set, and is not subject to dynamic bidding. |
	SuggestedBiddingStrategy *string `json:"suggestedBiddingStrategy,omitempty"`
	// Type of suggested action.
	Action *string `json:"action,omitempty"`
}

// NewBiddingStrategyRecommendation instantiates a new BiddingStrategyRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiddingStrategyRecommendation() *BiddingStrategyRecommendation {
	this := BiddingStrategyRecommendation{}
	return &this
}

// NewBiddingStrategyRecommendationWithDefaults instantiates a new BiddingStrategyRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiddingStrategyRecommendationWithDefaults() *BiddingStrategyRecommendation {
	this := BiddingStrategyRecommendation{}
	return &this
}

// GetSuggestedBiddingStrategy returns the SuggestedBiddingStrategy field value if set, zero value otherwise.
func (o *BiddingStrategyRecommendation) GetSuggestedBiddingStrategy() string {
	if o == nil || IsNil(o.SuggestedBiddingStrategy) {
		var ret string
		return ret
	}
	return *o.SuggestedBiddingStrategy
}

// GetSuggestedBiddingStrategyOk returns a tuple with the SuggestedBiddingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiddingStrategyRecommendation) GetSuggestedBiddingStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.SuggestedBiddingStrategy) {
		return nil, false
	}
	return o.SuggestedBiddingStrategy, true
}

// HasSuggestedBiddingStrategy returns a boolean if a field has been set.
func (o *BiddingStrategyRecommendation) HasSuggestedBiddingStrategy() bool {
	if o != nil && !IsNil(o.SuggestedBiddingStrategy) {
		return true
	}

	return false
}

// SetSuggestedBiddingStrategy gets a reference to the given string and assigns it to the SuggestedBiddingStrategy field.
func (o *BiddingStrategyRecommendation) SetSuggestedBiddingStrategy(v string) {
	o.SuggestedBiddingStrategy = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BiddingStrategyRecommendation) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiddingStrategyRecommendation) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BiddingStrategyRecommendation) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BiddingStrategyRecommendation) SetAction(v string) {
	o.Action = &v
}

func (o BiddingStrategyRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuggestedBiddingStrategy) {
		toSerialize["suggestedBiddingStrategy"] = o.SuggestedBiddingStrategy
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableBiddingStrategyRecommendation struct {
	value *BiddingStrategyRecommendation
	isSet bool
}

func (v NullableBiddingStrategyRecommendation) Get() *BiddingStrategyRecommendation {
	return v.value
}

func (v *NullableBiddingStrategyRecommendation) Set(val *BiddingStrategyRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableBiddingStrategyRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableBiddingStrategyRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiddingStrategyRecommendation(val *BiddingStrategyRecommendation) *NullableBiddingStrategyRecommendation {
	return &NullableBiddingStrategyRecommendation{value: val, isSet: true}
}

func (v NullableBiddingStrategyRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBiddingStrategyRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
