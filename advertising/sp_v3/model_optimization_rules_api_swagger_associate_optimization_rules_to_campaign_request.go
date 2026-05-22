package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest{}

// OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest Request body for create campaign to optimization rules association.
type OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest struct {
	// An array of rule identifiers.
	OptimizationRuleIds []string `json:"optimizationRuleIds"`
}

type _OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest

// NewOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest instantiates a new OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest(optimizationRuleIds []string) *OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest {
	this := OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest{}
	this.OptimizationRuleIds = optimizationRuleIds
	return &this
}

// NewOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequestWithDefaults instantiates a new OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequestWithDefaults() *OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest {
	this := OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest{}
	return &this
}

// GetOptimizationRuleIds returns the OptimizationRuleIds field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) GetOptimizationRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OptimizationRuleIds
}

// GetOptimizationRuleIdsOk returns a tuple with the OptimizationRuleIds field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) GetOptimizationRuleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimizationRuleIds, true
}

// SetOptimizationRuleIds sets field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) SetOptimizationRuleIds(v []string) {
	o.OptimizationRuleIds = v
}

func (o OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRuleIds"] = o.OptimizationRuleIds
	return toSerialize, nil
}

type NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest struct {
	value *OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest
	isSet bool
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) Get() *OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest {
	return v.value
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) Set(val *OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest(val *OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) *NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest {
	return &NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
