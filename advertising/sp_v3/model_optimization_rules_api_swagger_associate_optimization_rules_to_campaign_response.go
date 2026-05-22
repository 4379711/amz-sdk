package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse{}

// OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse Response object for create campaign to optimization rules association.
type OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse struct {
	// An enumerated error code for machine use.
	Code      *string                                                         `json:"code,omitempty"`
	Responses []OptimizationRulesAPISingleOptimizationRuleAssociationResponse `json:"responses,omitempty"`
}

// NewOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse instantiates a new OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse() *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse {
	this := OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse{}
	return &this
}

// NewOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponseWithDefaults instantiates a new OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponseWithDefaults() *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse {
	this := OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) SetCode(v string) {
	o.Code = &v
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) GetResponses() []OptimizationRulesAPISingleOptimizationRuleAssociationResponse {
	if o == nil || IsNil(o.Responses) {
		var ret []OptimizationRulesAPISingleOptimizationRuleAssociationResponse
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) GetResponsesOk() ([]OptimizationRulesAPISingleOptimizationRuleAssociationResponse, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []OptimizationRulesAPISingleOptimizationRuleAssociationResponse and assigns it to the Responses field.
func (o *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) SetResponses(v []OptimizationRulesAPISingleOptimizationRuleAssociationResponse) {
	o.Responses = v
}

func (o OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse struct {
	value *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse
	isSet bool
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) Get() *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) Set(val *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse(val *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) *NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse {
	return &NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
