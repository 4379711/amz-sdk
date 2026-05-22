package sp_v3

import "github.com/bytedance/sonic"

// checks if the OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest{}

// OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest Request body object for deleting campaign to optimization rules association.
type OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest struct {
	// An array of rule identifiers, all of which to disassociate from the campaign.
	OptimizationRuleIds []string `json:"optimizationRuleIds"`
}

type _OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest

// NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest(optimizationRuleIds []string) *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest {
	this := OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest{}
	this.OptimizationRuleIds = optimizationRuleIds
	return &this
}

// NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequestWithDefaults instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequestWithDefaults() *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest {
	this := OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest{}
	return &this
}

// GetOptimizationRuleIds returns the OptimizationRuleIds field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) GetOptimizationRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OptimizationRuleIds
}

// GetOptimizationRuleIdsOk returns a tuple with the OptimizationRuleIds field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) GetOptimizationRuleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimizationRuleIds, true
}

// SetOptimizationRuleIds sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) SetOptimizationRuleIds(v []string) {
	o.OptimizationRuleIds = v
}

func (o OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRuleIds"] = o.OptimizationRuleIds
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest struct {
	value *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest
	isSet bool
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) Get() *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest {
	return v.value
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) Set(val *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest(val *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest {
	return &NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
