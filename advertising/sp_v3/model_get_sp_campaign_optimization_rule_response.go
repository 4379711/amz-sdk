package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSPCampaignOptimizationRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSPCampaignOptimizationRuleResponse{}

// GetSPCampaignOptimizationRuleResponse struct for GetSPCampaignOptimizationRuleResponse
type GetSPCampaignOptimizationRuleResponse struct {
	CampaignOptimizationRule *CampaignOptimizationRule `json:"CampaignOptimizationRule,omitempty"`
}

// NewGetSPCampaignOptimizationRuleResponse instantiates a new GetSPCampaignOptimizationRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSPCampaignOptimizationRuleResponse() *GetSPCampaignOptimizationRuleResponse {
	this := GetSPCampaignOptimizationRuleResponse{}
	return &this
}

// NewGetSPCampaignOptimizationRuleResponseWithDefaults instantiates a new GetSPCampaignOptimizationRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSPCampaignOptimizationRuleResponseWithDefaults() *GetSPCampaignOptimizationRuleResponse {
	this := GetSPCampaignOptimizationRuleResponse{}
	return &this
}

// GetCampaignOptimizationRule returns the CampaignOptimizationRule field value if set, zero value otherwise.
func (o *GetSPCampaignOptimizationRuleResponse) GetCampaignOptimizationRule() CampaignOptimizationRule {
	if o == nil || IsNil(o.CampaignOptimizationRule) {
		var ret CampaignOptimizationRule
		return ret
	}
	return *o.CampaignOptimizationRule
}

// GetCampaignOptimizationRuleOk returns a tuple with the CampaignOptimizationRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSPCampaignOptimizationRuleResponse) GetCampaignOptimizationRuleOk() (*CampaignOptimizationRule, bool) {
	if o == nil || IsNil(o.CampaignOptimizationRule) {
		return nil, false
	}
	return o.CampaignOptimizationRule, true
}

// HasCampaignOptimizationRule returns a boolean if a field has been set.
func (o *GetSPCampaignOptimizationRuleResponse) HasCampaignOptimizationRule() bool {
	if o != nil && !IsNil(o.CampaignOptimizationRule) {
		return true
	}

	return false
}

// SetCampaignOptimizationRule gets a reference to the given CampaignOptimizationRule and assigns it to the CampaignOptimizationRule field.
func (o *GetSPCampaignOptimizationRuleResponse) SetCampaignOptimizationRule(v CampaignOptimizationRule) {
	o.CampaignOptimizationRule = &v
}

func (o GetSPCampaignOptimizationRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignOptimizationRule) {
		toSerialize["CampaignOptimizationRule"] = o.CampaignOptimizationRule
	}
	return toSerialize, nil
}

type NullableGetSPCampaignOptimizationRuleResponse struct {
	value *GetSPCampaignOptimizationRuleResponse
	isSet bool
}

func (v NullableGetSPCampaignOptimizationRuleResponse) Get() *GetSPCampaignOptimizationRuleResponse {
	return v.value
}

func (v *NullableGetSPCampaignOptimizationRuleResponse) Set(val *GetSPCampaignOptimizationRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSPCampaignOptimizationRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSPCampaignOptimizationRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSPCampaignOptimizationRuleResponse(val *GetSPCampaignOptimizationRuleResponse) *NullableGetSPCampaignOptimizationRuleResponse {
	return &NullableGetSPCampaignOptimizationRuleResponse{value: val, isSet: true}
}

func (v NullableGetSPCampaignOptimizationRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSPCampaignOptimizationRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
