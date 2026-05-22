package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the MultiCountryBidRecommendationPerTargetingExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiCountryBidRecommendationPerTargetingExpression{}

// MultiCountryBidRecommendationPerTargetingExpression struct for MultiCountryBidRecommendationPerTargetingExpression
type MultiCountryBidRecommendationPerTargetingExpression struct {
	CountrySuggestedBids map[string][]float64            `json:"countrySuggestedBids"`
	Expression           MultiCountryTargetingExpression `json:"expression"`
}

type _MultiCountryBidRecommendationPerTargetingExpression MultiCountryBidRecommendationPerTargetingExpression

// NewMultiCountryBidRecommendationPerTargetingExpression instantiates a new MultiCountryBidRecommendationPerTargetingExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiCountryBidRecommendationPerTargetingExpression(countrySuggestedBids map[string][]float64, expression MultiCountryTargetingExpression) *MultiCountryBidRecommendationPerTargetingExpression {
	this := MultiCountryBidRecommendationPerTargetingExpression{}
	this.CountrySuggestedBids = countrySuggestedBids
	this.Expression = expression
	return &this
}

// NewMultiCountryBidRecommendationPerTargetingExpressionWithDefaults instantiates a new MultiCountryBidRecommendationPerTargetingExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiCountryBidRecommendationPerTargetingExpressionWithDefaults() *MultiCountryBidRecommendationPerTargetingExpression {
	this := MultiCountryBidRecommendationPerTargetingExpression{}
	return &this
}

// GetCountrySuggestedBids returns the CountrySuggestedBids field value
func (o *MultiCountryBidRecommendationPerTargetingExpression) GetCountrySuggestedBids() map[string][]float64 {
	if o == nil {
		var ret map[string][]float64
		return ret
	}

	return o.CountrySuggestedBids
}

// GetCountrySuggestedBidsOk returns a tuple with the CountrySuggestedBids field value
// and a boolean to check if the value has been set.
func (o *MultiCountryBidRecommendationPerTargetingExpression) GetCountrySuggestedBidsOk() (*map[string][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountrySuggestedBids, true
}

// SetCountrySuggestedBids sets field value
func (o *MultiCountryBidRecommendationPerTargetingExpression) SetCountrySuggestedBids(v map[string][]float64) {
	o.CountrySuggestedBids = v
}

// GetExpression returns the Expression field value
func (o *MultiCountryBidRecommendationPerTargetingExpression) GetExpression() MultiCountryTargetingExpression {
	if o == nil {
		var ret MultiCountryTargetingExpression
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *MultiCountryBidRecommendationPerTargetingExpression) GetExpressionOk() (*MultiCountryTargetingExpression, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *MultiCountryBidRecommendationPerTargetingExpression) SetExpression(v MultiCountryTargetingExpression) {
	o.Expression = v
}

func (o MultiCountryBidRecommendationPerTargetingExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countrySuggestedBids"] = o.CountrySuggestedBids
	toSerialize["expression"] = o.Expression
	return toSerialize, nil
}

type NullableMultiCountryBidRecommendationPerTargetingExpression struct {
	value *MultiCountryBidRecommendationPerTargetingExpression
	isSet bool
}

func (v NullableMultiCountryBidRecommendationPerTargetingExpression) Get() *MultiCountryBidRecommendationPerTargetingExpression {
	return v.value
}

func (v *NullableMultiCountryBidRecommendationPerTargetingExpression) Set(val *MultiCountryBidRecommendationPerTargetingExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiCountryBidRecommendationPerTargetingExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiCountryBidRecommendationPerTargetingExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiCountryBidRecommendationPerTargetingExpression(val *MultiCountryBidRecommendationPerTargetingExpression) *NullableMultiCountryBidRecommendationPerTargetingExpression {
	return &NullableMultiCountryBidRecommendationPerTargetingExpression{value: val, isSet: true}
}

func (v NullableMultiCountryBidRecommendationPerTargetingExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMultiCountryBidRecommendationPerTargetingExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
