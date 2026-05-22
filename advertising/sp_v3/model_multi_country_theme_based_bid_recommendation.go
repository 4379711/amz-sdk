package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the MultiCountryThemeBasedBidRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiCountryThemeBasedBidRecommendation{}

// MultiCountryThemeBasedBidRecommendation struct for MultiCountryThemeBasedBidRecommendation
type MultiCountryThemeBasedBidRecommendation struct {
	Theme Theme `json:"theme"`
	// The bid recommendations for targeting expressions listed in the request.
	BidRecommendationsForTargetingExpressions []MultiCountryBidRecommendationPerTargetingExpression `json:"bidRecommendationsForTargetingExpressions"`
}

type _MultiCountryThemeBasedBidRecommendation MultiCountryThemeBasedBidRecommendation

// NewMultiCountryThemeBasedBidRecommendation instantiates a new MultiCountryThemeBasedBidRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiCountryThemeBasedBidRecommendation(theme Theme, bidRecommendationsForTargetingExpressions []MultiCountryBidRecommendationPerTargetingExpression) *MultiCountryThemeBasedBidRecommendation {
	this := MultiCountryThemeBasedBidRecommendation{}
	this.Theme = theme
	this.BidRecommendationsForTargetingExpressions = bidRecommendationsForTargetingExpressions
	return &this
}

// NewMultiCountryThemeBasedBidRecommendationWithDefaults instantiates a new MultiCountryThemeBasedBidRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiCountryThemeBasedBidRecommendationWithDefaults() *MultiCountryThemeBasedBidRecommendation {
	this := MultiCountryThemeBasedBidRecommendation{}
	return &this
}

// GetTheme returns the Theme field value
func (o *MultiCountryThemeBasedBidRecommendation) GetTheme() Theme {
	if o == nil {
		var ret Theme
		return ret
	}

	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *MultiCountryThemeBasedBidRecommendation) GetThemeOk() (*Theme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Theme, true
}

// SetTheme sets field value
func (o *MultiCountryThemeBasedBidRecommendation) SetTheme(v Theme) {
	o.Theme = v
}

// GetBidRecommendationsForTargetingExpressions returns the BidRecommendationsForTargetingExpressions field value
func (o *MultiCountryThemeBasedBidRecommendation) GetBidRecommendationsForTargetingExpressions() []MultiCountryBidRecommendationPerTargetingExpression {
	if o == nil {
		var ret []MultiCountryBidRecommendationPerTargetingExpression
		return ret
	}

	return o.BidRecommendationsForTargetingExpressions
}

// GetBidRecommendationsForTargetingExpressionsOk returns a tuple with the BidRecommendationsForTargetingExpressions field value
// and a boolean to check if the value has been set.
func (o *MultiCountryThemeBasedBidRecommendation) GetBidRecommendationsForTargetingExpressionsOk() ([]MultiCountryBidRecommendationPerTargetingExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.BidRecommendationsForTargetingExpressions, true
}

// SetBidRecommendationsForTargetingExpressions sets field value
func (o *MultiCountryThemeBasedBidRecommendation) SetBidRecommendationsForTargetingExpressions(v []MultiCountryBidRecommendationPerTargetingExpression) {
	o.BidRecommendationsForTargetingExpressions = v
}

func (o MultiCountryThemeBasedBidRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["theme"] = o.Theme
	toSerialize["bidRecommendationsForTargetingExpressions"] = o.BidRecommendationsForTargetingExpressions
	return toSerialize, nil
}

type NullableMultiCountryThemeBasedBidRecommendation struct {
	value *MultiCountryThemeBasedBidRecommendation
	isSet bool
}

func (v NullableMultiCountryThemeBasedBidRecommendation) Get() *MultiCountryThemeBasedBidRecommendation {
	return v.value
}

func (v *NullableMultiCountryThemeBasedBidRecommendation) Set(val *MultiCountryThemeBasedBidRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiCountryThemeBasedBidRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiCountryThemeBasedBidRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiCountryThemeBasedBidRecommendation(val *MultiCountryThemeBasedBidRecommendation) *NullableMultiCountryThemeBasedBidRecommendation {
	return &NullableMultiCountryThemeBasedBidRecommendation{value: val, isSet: true}
}

func (v NullableMultiCountryThemeBasedBidRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMultiCountryThemeBasedBidRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
