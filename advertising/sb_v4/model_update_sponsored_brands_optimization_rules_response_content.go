package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsOptimizationRulesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsOptimizationRulesResponseContent{}

// UpdateSponsoredBrandsOptimizationRulesResponseContent struct for UpdateSponsoredBrandsOptimizationRulesResponseContent
type UpdateSponsoredBrandsOptimizationRulesResponseContent struct {
	OptimizationRules BulkUpdateOptimizationRuleOperationResponse `json:"optimizationRules"`
}

type _UpdateSponsoredBrandsOptimizationRulesResponseContent UpdateSponsoredBrandsOptimizationRulesResponseContent

// NewUpdateSponsoredBrandsOptimizationRulesResponseContent instantiates a new UpdateSponsoredBrandsOptimizationRulesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsOptimizationRulesResponseContent(optimizationRules BulkUpdateOptimizationRuleOperationResponse) *UpdateSponsoredBrandsOptimizationRulesResponseContent {
	this := UpdateSponsoredBrandsOptimizationRulesResponseContent{}
	this.OptimizationRules = optimizationRules
	return &this
}

// NewUpdateSponsoredBrandsOptimizationRulesResponseContentWithDefaults instantiates a new UpdateSponsoredBrandsOptimizationRulesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsOptimizationRulesResponseContentWithDefaults() *UpdateSponsoredBrandsOptimizationRulesResponseContent {
	this := UpdateSponsoredBrandsOptimizationRulesResponseContent{}
	return &this
}

// GetOptimizationRules returns the OptimizationRules field value
func (o *UpdateSponsoredBrandsOptimizationRulesResponseContent) GetOptimizationRules() BulkUpdateOptimizationRuleOperationResponse {
	if o == nil {
		var ret BulkUpdateOptimizationRuleOperationResponse
		return ret
	}

	return o.OptimizationRules
}

// GetOptimizationRulesOk returns a tuple with the OptimizationRules field value
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsOptimizationRulesResponseContent) GetOptimizationRulesOk() (*BulkUpdateOptimizationRuleOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRules, true
}

// SetOptimizationRules sets field value
func (o *UpdateSponsoredBrandsOptimizationRulesResponseContent) SetOptimizationRules(v BulkUpdateOptimizationRuleOperationResponse) {
	o.OptimizationRules = v
}

func (o UpdateSponsoredBrandsOptimizationRulesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRules"] = o.OptimizationRules
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsOptimizationRulesResponseContent struct {
	value *UpdateSponsoredBrandsOptimizationRulesResponseContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsOptimizationRulesResponseContent) Get() *UpdateSponsoredBrandsOptimizationRulesResponseContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsOptimizationRulesResponseContent) Set(val *UpdateSponsoredBrandsOptimizationRulesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsOptimizationRulesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsOptimizationRulesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsOptimizationRulesResponseContent(val *UpdateSponsoredBrandsOptimizationRulesResponseContent) *NullableUpdateSponsoredBrandsOptimizationRulesResponseContent {
	return &NullableUpdateSponsoredBrandsOptimizationRulesResponseContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsOptimizationRulesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsOptimizationRulesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
