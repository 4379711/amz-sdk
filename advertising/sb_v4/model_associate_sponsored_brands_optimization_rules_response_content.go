package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AssociateSponsoredBrandsOptimizationRulesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociateSponsoredBrandsOptimizationRulesResponseContent{}

// AssociateSponsoredBrandsOptimizationRulesResponseContent struct for AssociateSponsoredBrandsOptimizationRulesResponseContent
type AssociateSponsoredBrandsOptimizationRulesResponseContent struct {
	OptimizationRuleAssociations BulkAssociationsOptimizationRuleResponse `json:"optimizationRuleAssociations"`
}

type _AssociateSponsoredBrandsOptimizationRulesResponseContent AssociateSponsoredBrandsOptimizationRulesResponseContent

// NewAssociateSponsoredBrandsOptimizationRulesResponseContent instantiates a new AssociateSponsoredBrandsOptimizationRulesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociateSponsoredBrandsOptimizationRulesResponseContent(optimizationRuleAssociations BulkAssociationsOptimizationRuleResponse) *AssociateSponsoredBrandsOptimizationRulesResponseContent {
	this := AssociateSponsoredBrandsOptimizationRulesResponseContent{}
	this.OptimizationRuleAssociations = optimizationRuleAssociations
	return &this
}

// NewAssociateSponsoredBrandsOptimizationRulesResponseContentWithDefaults instantiates a new AssociateSponsoredBrandsOptimizationRulesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociateSponsoredBrandsOptimizationRulesResponseContentWithDefaults() *AssociateSponsoredBrandsOptimizationRulesResponseContent {
	this := AssociateSponsoredBrandsOptimizationRulesResponseContent{}
	return &this
}

// GetOptimizationRuleAssociations returns the OptimizationRuleAssociations field value
func (o *AssociateSponsoredBrandsOptimizationRulesResponseContent) GetOptimizationRuleAssociations() BulkAssociationsOptimizationRuleResponse {
	if o == nil {
		var ret BulkAssociationsOptimizationRuleResponse
		return ret
	}

	return o.OptimizationRuleAssociations
}

// GetOptimizationRuleAssociationsOk returns a tuple with the OptimizationRuleAssociations field value
// and a boolean to check if the value has been set.
func (o *AssociateSponsoredBrandsOptimizationRulesResponseContent) GetOptimizationRuleAssociationsOk() (*BulkAssociationsOptimizationRuleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleAssociations, true
}

// SetOptimizationRuleAssociations sets field value
func (o *AssociateSponsoredBrandsOptimizationRulesResponseContent) SetOptimizationRuleAssociations(v BulkAssociationsOptimizationRuleResponse) {
	o.OptimizationRuleAssociations = v
}

func (o AssociateSponsoredBrandsOptimizationRulesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRuleAssociations"] = o.OptimizationRuleAssociations
	return toSerialize, nil
}

type NullableAssociateSponsoredBrandsOptimizationRulesResponseContent struct {
	value *AssociateSponsoredBrandsOptimizationRulesResponseContent
	isSet bool
}

func (v NullableAssociateSponsoredBrandsOptimizationRulesResponseContent) Get() *AssociateSponsoredBrandsOptimizationRulesResponseContent {
	return v.value
}

func (v *NullableAssociateSponsoredBrandsOptimizationRulesResponseContent) Set(val *AssociateSponsoredBrandsOptimizationRulesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociateSponsoredBrandsOptimizationRulesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociateSponsoredBrandsOptimizationRulesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociateSponsoredBrandsOptimizationRulesResponseContent(val *AssociateSponsoredBrandsOptimizationRulesResponseContent) *NullableAssociateSponsoredBrandsOptimizationRulesResponseContent {
	return &NullableAssociateSponsoredBrandsOptimizationRulesResponseContent{value: val, isSet: true}
}

func (v NullableAssociateSponsoredBrandsOptimizationRulesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssociateSponsoredBrandsOptimizationRulesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
