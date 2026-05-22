package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DisassociateSponsoredBrandsOptimizationRulesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisassociateSponsoredBrandsOptimizationRulesRequestContent{}

// DisassociateSponsoredBrandsOptimizationRulesRequestContent struct for DisassociateSponsoredBrandsOptimizationRulesRequestContent
type DisassociateSponsoredBrandsOptimizationRulesRequestContent struct {
	OptimizationRuleDisassociations []OptimizationRuleToEntityMapping `json:"optimizationRuleDisassociations"`
}

type _DisassociateSponsoredBrandsOptimizationRulesRequestContent DisassociateSponsoredBrandsOptimizationRulesRequestContent

// NewDisassociateSponsoredBrandsOptimizationRulesRequestContent instantiates a new DisassociateSponsoredBrandsOptimizationRulesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisassociateSponsoredBrandsOptimizationRulesRequestContent(optimizationRuleDisassociations []OptimizationRuleToEntityMapping) *DisassociateSponsoredBrandsOptimizationRulesRequestContent {
	this := DisassociateSponsoredBrandsOptimizationRulesRequestContent{}
	this.OptimizationRuleDisassociations = optimizationRuleDisassociations
	return &this
}

// NewDisassociateSponsoredBrandsOptimizationRulesRequestContentWithDefaults instantiates a new DisassociateSponsoredBrandsOptimizationRulesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisassociateSponsoredBrandsOptimizationRulesRequestContentWithDefaults() *DisassociateSponsoredBrandsOptimizationRulesRequestContent {
	this := DisassociateSponsoredBrandsOptimizationRulesRequestContent{}
	return &this
}

// GetOptimizationRuleDisassociations returns the OptimizationRuleDisassociations field value
func (o *DisassociateSponsoredBrandsOptimizationRulesRequestContent) GetOptimizationRuleDisassociations() []OptimizationRuleToEntityMapping {
	if o == nil {
		var ret []OptimizationRuleToEntityMapping
		return ret
	}

	return o.OptimizationRuleDisassociations
}

// GetOptimizationRuleDisassociationsOk returns a tuple with the OptimizationRuleDisassociations field value
// and a boolean to check if the value has been set.
func (o *DisassociateSponsoredBrandsOptimizationRulesRequestContent) GetOptimizationRuleDisassociationsOk() ([]OptimizationRuleToEntityMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimizationRuleDisassociations, true
}

// SetOptimizationRuleDisassociations sets field value
func (o *DisassociateSponsoredBrandsOptimizationRulesRequestContent) SetOptimizationRuleDisassociations(v []OptimizationRuleToEntityMapping) {
	o.OptimizationRuleDisassociations = v
}

func (o DisassociateSponsoredBrandsOptimizationRulesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRuleDisassociations"] = o.OptimizationRuleDisassociations
	return toSerialize, nil
}

type NullableDisassociateSponsoredBrandsOptimizationRulesRequestContent struct {
	value *DisassociateSponsoredBrandsOptimizationRulesRequestContent
	isSet bool
}

func (v NullableDisassociateSponsoredBrandsOptimizationRulesRequestContent) Get() *DisassociateSponsoredBrandsOptimizationRulesRequestContent {
	return v.value
}

func (v *NullableDisassociateSponsoredBrandsOptimizationRulesRequestContent) Set(val *DisassociateSponsoredBrandsOptimizationRulesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDisassociateSponsoredBrandsOptimizationRulesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDisassociateSponsoredBrandsOptimizationRulesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisassociateSponsoredBrandsOptimizationRulesRequestContent(val *DisassociateSponsoredBrandsOptimizationRulesRequestContent) *NullableDisassociateSponsoredBrandsOptimizationRulesRequestContent {
	return &NullableDisassociateSponsoredBrandsOptimizationRulesRequestContent{value: val, isSet: true}
}

func (v NullableDisassociateSponsoredBrandsOptimizationRulesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDisassociateSponsoredBrandsOptimizationRulesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
