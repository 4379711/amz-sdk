package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RuleRecommendationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleRecommendationError{}

// RuleRecommendationError struct for RuleRecommendationError
type RuleRecommendationError struct {
	// campaignId
	CampaignId *string                        `json:"campaignId,omitempty"`
	Error      *CampaignOptimizationRuleError `json:"Error,omitempty"`
}

// NewRuleRecommendationError instantiates a new RuleRecommendationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleRecommendationError() *RuleRecommendationError {
	this := RuleRecommendationError{}
	return &this
}

// NewRuleRecommendationErrorWithDefaults instantiates a new RuleRecommendationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleRecommendationErrorWithDefaults() *RuleRecommendationError {
	this := RuleRecommendationError{}
	return &this
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *RuleRecommendationError) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRecommendationError) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *RuleRecommendationError) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *RuleRecommendationError) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RuleRecommendationError) GetError() CampaignOptimizationRuleError {
	if o == nil || IsNil(o.Error) {
		var ret CampaignOptimizationRuleError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRecommendationError) GetErrorOk() (*CampaignOptimizationRuleError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RuleRecommendationError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given CampaignOptimizationRuleError and assigns it to the Error field.
func (o *RuleRecommendationError) SetError(v CampaignOptimizationRuleError) {
	o.Error = &v
}

func (o RuleRecommendationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.Error) {
		toSerialize["Error"] = o.Error
	}
	return toSerialize, nil
}

type NullableRuleRecommendationError struct {
	value *RuleRecommendationError
	isSet bool
}

func (v NullableRuleRecommendationError) Get() *RuleRecommendationError {
	return v.value
}

func (v *NullableRuleRecommendationError) Set(val *RuleRecommendationError) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleRecommendationError) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleRecommendationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleRecommendationError(val *RuleRecommendationError) *NullableRuleRecommendationError {
	return &NullableRuleRecommendationError{value: val, isSet: true}
}

func (v NullableRuleRecommendationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleRecommendationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
