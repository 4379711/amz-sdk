package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RuleRecommendationMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleRecommendationMetrics{}

// RuleRecommendationMetrics Performance Metrics supported by the rule recommendation
type RuleRecommendationMetrics struct {
	// return on ad spend value
	Roas *float64 `json:"roas,omitempty"`
}

// NewRuleRecommendationMetrics instantiates a new RuleRecommendationMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleRecommendationMetrics() *RuleRecommendationMetrics {
	this := RuleRecommendationMetrics{}
	return &this
}

// NewRuleRecommendationMetricsWithDefaults instantiates a new RuleRecommendationMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleRecommendationMetricsWithDefaults() *RuleRecommendationMetrics {
	this := RuleRecommendationMetrics{}
	return &this
}

// GetRoas returns the Roas field value if set, zero value otherwise.
func (o *RuleRecommendationMetrics) GetRoas() float64 {
	if o == nil || IsNil(o.Roas) {
		var ret float64
		return ret
	}
	return *o.Roas
}

// GetRoasOk returns a tuple with the Roas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRecommendationMetrics) GetRoasOk() (*float64, bool) {
	if o == nil || IsNil(o.Roas) {
		return nil, false
	}
	return o.Roas, true
}

// HasRoas returns a boolean if a field has been set.
func (o *RuleRecommendationMetrics) HasRoas() bool {
	if o != nil && !IsNil(o.Roas) {
		return true
	}

	return false
}

// SetRoas gets a reference to the given float64 and assigns it to the Roas field.
func (o *RuleRecommendationMetrics) SetRoas(v float64) {
	o.Roas = &v
}

func (o RuleRecommendationMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roas) {
		toSerialize["roas"] = o.Roas
	}
	return toSerialize, nil
}

type NullableRuleRecommendationMetrics struct {
	value *RuleRecommendationMetrics
	isSet bool
}

func (v NullableRuleRecommendationMetrics) Get() *RuleRecommendationMetrics {
	return v.value
}

func (v *NullableRuleRecommendationMetrics) Set(val *RuleRecommendationMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleRecommendationMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleRecommendationMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleRecommendationMetrics(val *RuleRecommendationMetrics) *NullableRuleRecommendationMetrics {
	return &NullableRuleRecommendationMetrics{value: val, isSet: true}
}

func (v NullableRuleRecommendationMetrics) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleRecommendationMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
