package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ThemeBasedBidRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThemeBasedBidRecommendation{}

// ThemeBasedBidRecommendation struct for ThemeBasedBidRecommendation
type ThemeBasedBidRecommendation struct {
	Theme Theme `json:"theme"`
	// The bid recommendations for targeting expressions listed in the request.
	BidRecommendationsForTargetingExpressions []BidRecommendationPerTargetingExpression `json:"bidRecommendationsForTargetingExpressions"`
	ImpactMetrics                             NullableImpactMetrics                     `json:"impactMetrics,omitempty"`
}

type _ThemeBasedBidRecommendation ThemeBasedBidRecommendation

// NewThemeBasedBidRecommendation instantiates a new ThemeBasedBidRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThemeBasedBidRecommendation(theme Theme, bidRecommendationsForTargetingExpressions []BidRecommendationPerTargetingExpression) *ThemeBasedBidRecommendation {
	this := ThemeBasedBidRecommendation{}
	this.Theme = theme
	this.BidRecommendationsForTargetingExpressions = bidRecommendationsForTargetingExpressions
	return &this
}

// NewThemeBasedBidRecommendationWithDefaults instantiates a new ThemeBasedBidRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThemeBasedBidRecommendationWithDefaults() *ThemeBasedBidRecommendation {
	this := ThemeBasedBidRecommendation{}
	return &this
}

// GetTheme returns the Theme field value
func (o *ThemeBasedBidRecommendation) GetTheme() Theme {
	if o == nil {
		var ret Theme
		return ret
	}

	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *ThemeBasedBidRecommendation) GetThemeOk() (*Theme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Theme, true
}

// SetTheme sets field value
func (o *ThemeBasedBidRecommendation) SetTheme(v Theme) {
	o.Theme = v
}

// GetBidRecommendationsForTargetingExpressions returns the BidRecommendationsForTargetingExpressions field value
func (o *ThemeBasedBidRecommendation) GetBidRecommendationsForTargetingExpressions() []BidRecommendationPerTargetingExpression {
	if o == nil {
		var ret []BidRecommendationPerTargetingExpression
		return ret
	}

	return o.BidRecommendationsForTargetingExpressions
}

// GetBidRecommendationsForTargetingExpressionsOk returns a tuple with the BidRecommendationsForTargetingExpressions field value
// and a boolean to check if the value has been set.
func (o *ThemeBasedBidRecommendation) GetBidRecommendationsForTargetingExpressionsOk() ([]BidRecommendationPerTargetingExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.BidRecommendationsForTargetingExpressions, true
}

// SetBidRecommendationsForTargetingExpressions sets field value
func (o *ThemeBasedBidRecommendation) SetBidRecommendationsForTargetingExpressions(v []BidRecommendationPerTargetingExpression) {
	o.BidRecommendationsForTargetingExpressions = v
}

// GetImpactMetrics returns the ImpactMetrics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThemeBasedBidRecommendation) GetImpactMetrics() ImpactMetrics {
	if o == nil || IsNil(o.ImpactMetrics.Get()) {
		var ret ImpactMetrics
		return ret
	}
	return *o.ImpactMetrics.Get()
}

// GetImpactMetricsOk returns a tuple with the ImpactMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThemeBasedBidRecommendation) GetImpactMetricsOk() (*ImpactMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImpactMetrics.Get(), o.ImpactMetrics.IsSet()
}

// HasImpactMetrics returns a boolean if a field has been set.
func (o *ThemeBasedBidRecommendation) HasImpactMetrics() bool {
	if o != nil && o.ImpactMetrics.IsSet() {
		return true
	}

	return false
}

// SetImpactMetrics gets a reference to the given NullableImpactMetrics and assigns it to the ImpactMetrics field.
func (o *ThemeBasedBidRecommendation) SetImpactMetrics(v ImpactMetrics) {
	o.ImpactMetrics.Set(&v)
}

// SetImpactMetricsNil sets the value for ImpactMetrics to be an explicit nil
func (o *ThemeBasedBidRecommendation) SetImpactMetricsNil() {
	o.ImpactMetrics.Set(nil)
}

// UnsetImpactMetrics ensures that no value is present for ImpactMetrics, not even an explicit nil
func (o *ThemeBasedBidRecommendation) UnsetImpactMetrics() {
	o.ImpactMetrics.Unset()
}

func (o ThemeBasedBidRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["theme"] = o.Theme
	toSerialize["bidRecommendationsForTargetingExpressions"] = o.BidRecommendationsForTargetingExpressions
	if o.ImpactMetrics.IsSet() {
		toSerialize["impactMetrics"] = o.ImpactMetrics.Get()
	}
	return toSerialize, nil
}

type NullableThemeBasedBidRecommendation struct {
	value *ThemeBasedBidRecommendation
	isSet bool
}

func (v NullableThemeBasedBidRecommendation) Get() *ThemeBasedBidRecommendation {
	return v.value
}

func (v *NullableThemeBasedBidRecommendation) Set(val *ThemeBasedBidRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableThemeBasedBidRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableThemeBasedBidRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThemeBasedBidRecommendation(val *ThemeBasedBidRecommendation) *NullableThemeBasedBidRecommendation {
	return &NullableThemeBasedBidRecommendation{value: val, isSet: true}
}

func (v NullableThemeBasedBidRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThemeBasedBidRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
