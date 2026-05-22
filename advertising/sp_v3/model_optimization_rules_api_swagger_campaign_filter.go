package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPICampaignFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPICampaignFilter{}

// OptimizationRulesAPICampaignFilter Filter on campaigns.
type OptimizationRulesAPICampaignFilter struct {
	CampaignId *OptimizationRulesAPIEntityFieldFilter `json:"campaignId,omitempty"`
}

// NewOptimizationRulesAPICampaignFilter instantiates a new OptimizationRulesAPICampaignFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPICampaignFilter() *OptimizationRulesAPICampaignFilter {
	this := OptimizationRulesAPICampaignFilter{}
	return &this
}

// NewOptimizationRulesAPICampaignFilterWithDefaults instantiates a new OptimizationRulesAPICampaignFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPICampaignFilterWithDefaults() *OptimizationRulesAPICampaignFilter {
	this := OptimizationRulesAPICampaignFilter{}
	return &this
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *OptimizationRulesAPICampaignFilter) GetCampaignId() OptimizationRulesAPIEntityFieldFilter {
	if o == nil || IsNil(o.CampaignId) {
		var ret OptimizationRulesAPIEntityFieldFilter
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPICampaignFilter) GetCampaignIdOk() (*OptimizationRulesAPIEntityFieldFilter, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *OptimizationRulesAPICampaignFilter) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given OptimizationRulesAPIEntityFieldFilter and assigns it to the CampaignId field.
func (o *OptimizationRulesAPICampaignFilter) SetCampaignId(v OptimizationRulesAPIEntityFieldFilter) {
	o.CampaignId = &v
}

func (o OptimizationRulesAPICampaignFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPICampaignFilter struct {
	value *OptimizationRulesAPICampaignFilter
	isSet bool
}

func (v NullableOptimizationRulesAPICampaignFilter) Get() *OptimizationRulesAPICampaignFilter {
	return v.value
}

func (v *NullableOptimizationRulesAPICampaignFilter) Set(val *OptimizationRulesAPICampaignFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPICampaignFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPICampaignFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPICampaignFilter(val *OptimizationRulesAPICampaignFilter) *NullableOptimizationRulesAPICampaignFilter {
	return &NullableOptimizationRulesAPICampaignFilter{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPICampaignFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPICampaignFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
