package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DisassociateSponsoredBrandsOptimizationRulesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisassociateSponsoredBrandsOptimizationRulesResponseContent{}

// DisassociateSponsoredBrandsOptimizationRulesResponseContent struct for DisassociateSponsoredBrandsOptimizationRulesResponseContent
type DisassociateSponsoredBrandsOptimizationRulesResponseContent struct {
	OptimizationRuleDisassociations BulkDisassociationsOptimizationRuleResponse `json:"optimizationRuleDisassociations"`
}

type _DisassociateSponsoredBrandsOptimizationRulesResponseContent DisassociateSponsoredBrandsOptimizationRulesResponseContent

// NewDisassociateSponsoredBrandsOptimizationRulesResponseContent instantiates a new DisassociateSponsoredBrandsOptimizationRulesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisassociateSponsoredBrandsOptimizationRulesResponseContent(optimizationRuleDisassociations BulkDisassociationsOptimizationRuleResponse) *DisassociateSponsoredBrandsOptimizationRulesResponseContent {
	this := DisassociateSponsoredBrandsOptimizationRulesResponseContent{}
	this.OptimizationRuleDisassociations = optimizationRuleDisassociations
	return &this
}

// NewDisassociateSponsoredBrandsOptimizationRulesResponseContentWithDefaults instantiates a new DisassociateSponsoredBrandsOptimizationRulesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisassociateSponsoredBrandsOptimizationRulesResponseContentWithDefaults() *DisassociateSponsoredBrandsOptimizationRulesResponseContent {
	this := DisassociateSponsoredBrandsOptimizationRulesResponseContent{}
	return &this
}

// GetOptimizationRuleDisassociations returns the OptimizationRuleDisassociations field value
func (o *DisassociateSponsoredBrandsOptimizationRulesResponseContent) GetOptimizationRuleDisassociations() BulkDisassociationsOptimizationRuleResponse {
	if o == nil {
		var ret BulkDisassociationsOptimizationRuleResponse
		return ret
	}

	return o.OptimizationRuleDisassociations
}

// GetOptimizationRuleDisassociationsOk returns a tuple with the OptimizationRuleDisassociations field value
// and a boolean to check if the value has been set.
func (o *DisassociateSponsoredBrandsOptimizationRulesResponseContent) GetOptimizationRuleDisassociationsOk() (*BulkDisassociationsOptimizationRuleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleDisassociations, true
}

// SetOptimizationRuleDisassociations sets field value
func (o *DisassociateSponsoredBrandsOptimizationRulesResponseContent) SetOptimizationRuleDisassociations(v BulkDisassociationsOptimizationRuleResponse) {
	o.OptimizationRuleDisassociations = v
}

func (o DisassociateSponsoredBrandsOptimizationRulesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRuleDisassociations"] = o.OptimizationRuleDisassociations
	return toSerialize, nil
}

type NullableDisassociateSponsoredBrandsOptimizationRulesResponseContent struct {
	value *DisassociateSponsoredBrandsOptimizationRulesResponseContent
	isSet bool
}

func (v NullableDisassociateSponsoredBrandsOptimizationRulesResponseContent) Get() *DisassociateSponsoredBrandsOptimizationRulesResponseContent {
	return v.value
}

func (v *NullableDisassociateSponsoredBrandsOptimizationRulesResponseContent) Set(val *DisassociateSponsoredBrandsOptimizationRulesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDisassociateSponsoredBrandsOptimizationRulesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDisassociateSponsoredBrandsOptimizationRulesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisassociateSponsoredBrandsOptimizationRulesResponseContent(val *DisassociateSponsoredBrandsOptimizationRulesResponseContent) *NullableDisassociateSponsoredBrandsOptimizationRulesResponseContent {
	return &NullableDisassociateSponsoredBrandsOptimizationRulesResponseContent{value: val, isSet: true}
}

func (v NullableDisassociateSponsoredBrandsOptimizationRulesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDisassociateSponsoredBrandsOptimizationRulesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
