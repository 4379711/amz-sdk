package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AssociateSponsoredBrandsOptimizationRulesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociateSponsoredBrandsOptimizationRulesRequestContent{}

// AssociateSponsoredBrandsOptimizationRulesRequestContent struct for AssociateSponsoredBrandsOptimizationRulesRequestContent
type AssociateSponsoredBrandsOptimizationRulesRequestContent struct {
	OptimizationRuleAssociations []OptimizationRuleToEntityMapping `json:"optimizationRuleAssociations"`
}

type _AssociateSponsoredBrandsOptimizationRulesRequestContent AssociateSponsoredBrandsOptimizationRulesRequestContent

// NewAssociateSponsoredBrandsOptimizationRulesRequestContent instantiates a new AssociateSponsoredBrandsOptimizationRulesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociateSponsoredBrandsOptimizationRulesRequestContent(optimizationRuleAssociations []OptimizationRuleToEntityMapping) *AssociateSponsoredBrandsOptimizationRulesRequestContent {
	this := AssociateSponsoredBrandsOptimizationRulesRequestContent{}
	this.OptimizationRuleAssociations = optimizationRuleAssociations
	return &this
}

// NewAssociateSponsoredBrandsOptimizationRulesRequestContentWithDefaults instantiates a new AssociateSponsoredBrandsOptimizationRulesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociateSponsoredBrandsOptimizationRulesRequestContentWithDefaults() *AssociateSponsoredBrandsOptimizationRulesRequestContent {
	this := AssociateSponsoredBrandsOptimizationRulesRequestContent{}
	return &this
}

// GetOptimizationRuleAssociations returns the OptimizationRuleAssociations field value
func (o *AssociateSponsoredBrandsOptimizationRulesRequestContent) GetOptimizationRuleAssociations() []OptimizationRuleToEntityMapping {
	if o == nil {
		var ret []OptimizationRuleToEntityMapping
		return ret
	}

	return o.OptimizationRuleAssociations
}

// GetOptimizationRuleAssociationsOk returns a tuple with the OptimizationRuleAssociations field value
// and a boolean to check if the value has been set.
func (o *AssociateSponsoredBrandsOptimizationRulesRequestContent) GetOptimizationRuleAssociationsOk() ([]OptimizationRuleToEntityMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimizationRuleAssociations, true
}

// SetOptimizationRuleAssociations sets field value
func (o *AssociateSponsoredBrandsOptimizationRulesRequestContent) SetOptimizationRuleAssociations(v []OptimizationRuleToEntityMapping) {
	o.OptimizationRuleAssociations = v
}

func (o AssociateSponsoredBrandsOptimizationRulesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRuleAssociations"] = o.OptimizationRuleAssociations
	return toSerialize, nil
}

type NullableAssociateSponsoredBrandsOptimizationRulesRequestContent struct {
	value *AssociateSponsoredBrandsOptimizationRulesRequestContent
	isSet bool
}

func (v NullableAssociateSponsoredBrandsOptimizationRulesRequestContent) Get() *AssociateSponsoredBrandsOptimizationRulesRequestContent {
	return v.value
}

func (v *NullableAssociateSponsoredBrandsOptimizationRulesRequestContent) Set(val *AssociateSponsoredBrandsOptimizationRulesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociateSponsoredBrandsOptimizationRulesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociateSponsoredBrandsOptimizationRulesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociateSponsoredBrandsOptimizationRulesRequestContent(val *AssociateSponsoredBrandsOptimizationRulesRequestContent) *NullableAssociateSponsoredBrandsOptimizationRulesRequestContent {
	return &NullableAssociateSponsoredBrandsOptimizationRulesRequestContent{value: val, isSet: true}
}

func (v NullableAssociateSponsoredBrandsOptimizationRulesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssociateSponsoredBrandsOptimizationRulesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
