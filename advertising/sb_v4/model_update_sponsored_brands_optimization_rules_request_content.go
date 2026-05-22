package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsOptimizationRulesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsOptimizationRulesRequestContent{}

// UpdateSponsoredBrandsOptimizationRulesRequestContent struct for UpdateSponsoredBrandsOptimizationRulesRequestContent
type UpdateSponsoredBrandsOptimizationRulesRequestContent struct {
	OptimizationRules []UpdateOptimizationRule `json:"optimizationRules"`
}

type _UpdateSponsoredBrandsOptimizationRulesRequestContent UpdateSponsoredBrandsOptimizationRulesRequestContent

// NewUpdateSponsoredBrandsOptimizationRulesRequestContent instantiates a new UpdateSponsoredBrandsOptimizationRulesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsOptimizationRulesRequestContent(optimizationRules []UpdateOptimizationRule) *UpdateSponsoredBrandsOptimizationRulesRequestContent {
	this := UpdateSponsoredBrandsOptimizationRulesRequestContent{}
	this.OptimizationRules = optimizationRules
	return &this
}

// NewUpdateSponsoredBrandsOptimizationRulesRequestContentWithDefaults instantiates a new UpdateSponsoredBrandsOptimizationRulesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsOptimizationRulesRequestContentWithDefaults() *UpdateSponsoredBrandsOptimizationRulesRequestContent {
	this := UpdateSponsoredBrandsOptimizationRulesRequestContent{}
	return &this
}

// GetOptimizationRules returns the OptimizationRules field value
func (o *UpdateSponsoredBrandsOptimizationRulesRequestContent) GetOptimizationRules() []UpdateOptimizationRule {
	if o == nil {
		var ret []UpdateOptimizationRule
		return ret
	}

	return o.OptimizationRules
}

// GetOptimizationRulesOk returns a tuple with the OptimizationRules field value
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsOptimizationRulesRequestContent) GetOptimizationRulesOk() ([]UpdateOptimizationRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimizationRules, true
}

// SetOptimizationRules sets field value
func (o *UpdateSponsoredBrandsOptimizationRulesRequestContent) SetOptimizationRules(v []UpdateOptimizationRule) {
	o.OptimizationRules = v
}

func (o UpdateSponsoredBrandsOptimizationRulesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRules"] = o.OptimizationRules
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsOptimizationRulesRequestContent struct {
	value *UpdateSponsoredBrandsOptimizationRulesRequestContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsOptimizationRulesRequestContent) Get() *UpdateSponsoredBrandsOptimizationRulesRequestContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsOptimizationRulesRequestContent) Set(val *UpdateSponsoredBrandsOptimizationRulesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsOptimizationRulesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsOptimizationRulesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsOptimizationRulesRequestContent(val *UpdateSponsoredBrandsOptimizationRulesRequestContent) *NullableUpdateSponsoredBrandsOptimizationRulesRequestContent {
	return &NullableUpdateSponsoredBrandsOptimizationRulesRequestContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsOptimizationRulesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsOptimizationRulesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
