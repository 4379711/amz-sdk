package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ThemeBasedBidRecommendationV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThemeBasedBidRecommendationV4{}

// ThemeBasedBidRecommendationV4 struct for ThemeBasedBidRecommendationV4
type ThemeBasedBidRecommendationV4 struct {
	Theme Theme `json:"theme"`
	// The bid recommendations for targeting expressions listed in the request.
	BidRecommendationsForTargetingExpressions []BidRecommendationPerTargetingExpressionV4 `json:"bidRecommendationsForTargetingExpressions"`
}

type _ThemeBasedBidRecommendationV4 ThemeBasedBidRecommendationV4

// NewThemeBasedBidRecommendationV4 instantiates a new ThemeBasedBidRecommendationV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThemeBasedBidRecommendationV4(theme Theme, bidRecommendationsForTargetingExpressions []BidRecommendationPerTargetingExpressionV4) *ThemeBasedBidRecommendationV4 {
	this := ThemeBasedBidRecommendationV4{}
	this.Theme = theme
	this.BidRecommendationsForTargetingExpressions = bidRecommendationsForTargetingExpressions
	return &this
}

// NewThemeBasedBidRecommendationV4WithDefaults instantiates a new ThemeBasedBidRecommendationV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThemeBasedBidRecommendationV4WithDefaults() *ThemeBasedBidRecommendationV4 {
	this := ThemeBasedBidRecommendationV4{}
	return &this
}

// GetTheme returns the Theme field value
func (o *ThemeBasedBidRecommendationV4) GetTheme() Theme {
	if o == nil {
		var ret Theme
		return ret
	}

	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *ThemeBasedBidRecommendationV4) GetThemeOk() (*Theme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Theme, true
}

// SetTheme sets field value
func (o *ThemeBasedBidRecommendationV4) SetTheme(v Theme) {
	o.Theme = v
}

// GetBidRecommendationsForTargetingExpressions returns the BidRecommendationsForTargetingExpressions field value
func (o *ThemeBasedBidRecommendationV4) GetBidRecommendationsForTargetingExpressions() []BidRecommendationPerTargetingExpressionV4 {
	if o == nil {
		var ret []BidRecommendationPerTargetingExpressionV4
		return ret
	}

	return o.BidRecommendationsForTargetingExpressions
}

// GetBidRecommendationsForTargetingExpressionsOk returns a tuple with the BidRecommendationsForTargetingExpressions field value
// and a boolean to check if the value has been set.
func (o *ThemeBasedBidRecommendationV4) GetBidRecommendationsForTargetingExpressionsOk() ([]BidRecommendationPerTargetingExpressionV4, bool) {
	if o == nil {
		return nil, false
	}
	return o.BidRecommendationsForTargetingExpressions, true
}

// SetBidRecommendationsForTargetingExpressions sets field value
func (o *ThemeBasedBidRecommendationV4) SetBidRecommendationsForTargetingExpressions(v []BidRecommendationPerTargetingExpressionV4) {
	o.BidRecommendationsForTargetingExpressions = v
}

func (o ThemeBasedBidRecommendationV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["theme"] = o.Theme
	toSerialize["bidRecommendationsForTargetingExpressions"] = o.BidRecommendationsForTargetingExpressions
	return toSerialize, nil
}

type NullableThemeBasedBidRecommendationV4 struct {
	value *ThemeBasedBidRecommendationV4
	isSet bool
}

func (v NullableThemeBasedBidRecommendationV4) Get() *ThemeBasedBidRecommendationV4 {
	return v.value
}

func (v *NullableThemeBasedBidRecommendationV4) Set(val *ThemeBasedBidRecommendationV4) {
	v.value = val
	v.isSet = true
}

func (v NullableThemeBasedBidRecommendationV4) IsSet() bool {
	return v.isSet
}

func (v *NullableThemeBasedBidRecommendationV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThemeBasedBidRecommendationV4(val *ThemeBasedBidRecommendationV4) *NullableThemeBasedBidRecommendationV4 {
	return &NullableThemeBasedBidRecommendationV4{value: val, isSet: true}
}

func (v NullableThemeBasedBidRecommendationV4) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThemeBasedBidRecommendationV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
