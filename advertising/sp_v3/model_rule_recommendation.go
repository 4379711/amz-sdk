package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RuleRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleRecommendation{}

// RuleRecommendation struct for RuleRecommendation
type RuleRecommendation struct {
	// campaignId
	CampaignId         *string                    `json:"campaignId,omitempty"`
	PerformanceMetrics *RuleRecommendationMetrics `json:"performanceMetrics,omitempty"`
}

// NewRuleRecommendation instantiates a new RuleRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleRecommendation() *RuleRecommendation {
	this := RuleRecommendation{}
	return &this
}

// NewRuleRecommendationWithDefaults instantiates a new RuleRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleRecommendationWithDefaults() *RuleRecommendation {
	this := RuleRecommendation{}
	return &this
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *RuleRecommendation) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRecommendation) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *RuleRecommendation) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *RuleRecommendation) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetPerformanceMetrics returns the PerformanceMetrics field value if set, zero value otherwise.
func (o *RuleRecommendation) GetPerformanceMetrics() RuleRecommendationMetrics {
	if o == nil || IsNil(o.PerformanceMetrics) {
		var ret RuleRecommendationMetrics
		return ret
	}
	return *o.PerformanceMetrics
}

// GetPerformanceMetricsOk returns a tuple with the PerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRecommendation) GetPerformanceMetricsOk() (*RuleRecommendationMetrics, bool) {
	if o == nil || IsNil(o.PerformanceMetrics) {
		return nil, false
	}
	return o.PerformanceMetrics, true
}

// HasPerformanceMetrics returns a boolean if a field has been set.
func (o *RuleRecommendation) HasPerformanceMetrics() bool {
	if o != nil && !IsNil(o.PerformanceMetrics) {
		return true
	}

	return false
}

// SetPerformanceMetrics gets a reference to the given RuleRecommendationMetrics and assigns it to the PerformanceMetrics field.
func (o *RuleRecommendation) SetPerformanceMetrics(v RuleRecommendationMetrics) {
	o.PerformanceMetrics = &v
}

func (o RuleRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.PerformanceMetrics) {
		toSerialize["performanceMetrics"] = o.PerformanceMetrics
	}
	return toSerialize, nil
}

type NullableRuleRecommendation struct {
	value *RuleRecommendation
	isSet bool
}

func (v NullableRuleRecommendation) Get() *RuleRecommendation {
	return v.value
}

func (v *NullableRuleRecommendation) Set(val *RuleRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleRecommendation(val *RuleRecommendation) *NullableRuleRecommendation {
	return &NullableRuleRecommendation{value: val, isSet: true}
}

func (v NullableRuleRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
