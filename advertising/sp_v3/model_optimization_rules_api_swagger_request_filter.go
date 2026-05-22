package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIRequestFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIRequestFilter{}

// OptimizationRulesAPIRequestFilter Filter used in search requests.
type OptimizationRulesAPIRequestFilter struct {
	OptimizationRuleFilter *OptimizationRulesAPIOptimizationRuleFilter `json:"optimizationRuleFilter,omitempty"`
	CampaignFilter         *OptimizationRulesAPICampaignFilter         `json:"campaignFilter,omitempty"`
}

// NewOptimizationRulesAPIRequestFilter instantiates a new OptimizationRulesAPIRequestFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIRequestFilter() *OptimizationRulesAPIRequestFilter {
	this := OptimizationRulesAPIRequestFilter{}
	return &this
}

// NewOptimizationRulesAPIRequestFilterWithDefaults instantiates a new OptimizationRulesAPIRequestFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIRequestFilterWithDefaults() *OptimizationRulesAPIRequestFilter {
	this := OptimizationRulesAPIRequestFilter{}
	return &this
}

// GetOptimizationRuleFilter returns the OptimizationRuleFilter field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRequestFilter) GetOptimizationRuleFilter() OptimizationRulesAPIOptimizationRuleFilter {
	if o == nil || IsNil(o.OptimizationRuleFilter) {
		var ret OptimizationRulesAPIOptimizationRuleFilter
		return ret
	}
	return *o.OptimizationRuleFilter
}

// GetOptimizationRuleFilterOk returns a tuple with the OptimizationRuleFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRequestFilter) GetOptimizationRuleFilterOk() (*OptimizationRulesAPIOptimizationRuleFilter, bool) {
	if o == nil || IsNil(o.OptimizationRuleFilter) {
		return nil, false
	}
	return o.OptimizationRuleFilter, true
}

// HasOptimizationRuleFilter returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRequestFilter) HasOptimizationRuleFilter() bool {
	if o != nil && !IsNil(o.OptimizationRuleFilter) {
		return true
	}

	return false
}

// SetOptimizationRuleFilter gets a reference to the given OptimizationRulesAPIOptimizationRuleFilter and assigns it to the OptimizationRuleFilter field.
func (o *OptimizationRulesAPIRequestFilter) SetOptimizationRuleFilter(v OptimizationRulesAPIOptimizationRuleFilter) {
	o.OptimizationRuleFilter = &v
}

// GetCampaignFilter returns the CampaignFilter field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRequestFilter) GetCampaignFilter() OptimizationRulesAPICampaignFilter {
	if o == nil || IsNil(o.CampaignFilter) {
		var ret OptimizationRulesAPICampaignFilter
		return ret
	}
	return *o.CampaignFilter
}

// GetCampaignFilterOk returns a tuple with the CampaignFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRequestFilter) GetCampaignFilterOk() (*OptimizationRulesAPICampaignFilter, bool) {
	if o == nil || IsNil(o.CampaignFilter) {
		return nil, false
	}
	return o.CampaignFilter, true
}

// HasCampaignFilter returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRequestFilter) HasCampaignFilter() bool {
	if o != nil && !IsNil(o.CampaignFilter) {
		return true
	}

	return false
}

// SetCampaignFilter gets a reference to the given OptimizationRulesAPICampaignFilter and assigns it to the CampaignFilter field.
func (o *OptimizationRulesAPIRequestFilter) SetCampaignFilter(v OptimizationRulesAPICampaignFilter) {
	o.CampaignFilter = &v
}

func (o OptimizationRulesAPIRequestFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptimizationRuleFilter) {
		toSerialize["optimizationRuleFilter"] = o.OptimizationRuleFilter
	}
	if !IsNil(o.CampaignFilter) {
		toSerialize["campaignFilter"] = o.CampaignFilter
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIRequestFilter struct {
	value *OptimizationRulesAPIRequestFilter
	isSet bool
}

func (v NullableOptimizationRulesAPIRequestFilter) Get() *OptimizationRulesAPIRequestFilter {
	return v.value
}

func (v *NullableOptimizationRulesAPIRequestFilter) Set(val *OptimizationRulesAPIRequestFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRequestFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRequestFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRequestFilter(val *OptimizationRulesAPIRequestFilter) *NullableOptimizationRulesAPIRequestFilter {
	return &NullableOptimizationRulesAPIRequestFilter{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRequestFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRequestFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
