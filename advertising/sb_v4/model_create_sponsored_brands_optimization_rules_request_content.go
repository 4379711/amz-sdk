package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsOptimizationRulesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsOptimizationRulesRequestContent{}

// CreateSponsoredBrandsOptimizationRulesRequestContent struct for CreateSponsoredBrandsOptimizationRulesRequestContent
type CreateSponsoredBrandsOptimizationRulesRequestContent struct {
	OptimizationRules []CreateOptimizationRule `json:"optimizationRules"`
}

type _CreateSponsoredBrandsOptimizationRulesRequestContent CreateSponsoredBrandsOptimizationRulesRequestContent

// NewCreateSponsoredBrandsOptimizationRulesRequestContent instantiates a new CreateSponsoredBrandsOptimizationRulesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsOptimizationRulesRequestContent(optimizationRules []CreateOptimizationRule) *CreateSponsoredBrandsOptimizationRulesRequestContent {
	this := CreateSponsoredBrandsOptimizationRulesRequestContent{}
	this.OptimizationRules = optimizationRules
	return &this
}

// NewCreateSponsoredBrandsOptimizationRulesRequestContentWithDefaults instantiates a new CreateSponsoredBrandsOptimizationRulesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsOptimizationRulesRequestContentWithDefaults() *CreateSponsoredBrandsOptimizationRulesRequestContent {
	this := CreateSponsoredBrandsOptimizationRulesRequestContent{}
	return &this
}

// GetOptimizationRules returns the OptimizationRules field value
func (o *CreateSponsoredBrandsOptimizationRulesRequestContent) GetOptimizationRules() []CreateOptimizationRule {
	if o == nil {
		var ret []CreateOptimizationRule
		return ret
	}

	return o.OptimizationRules
}

// GetOptimizationRulesOk returns a tuple with the OptimizationRules field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsOptimizationRulesRequestContent) GetOptimizationRulesOk() ([]CreateOptimizationRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimizationRules, true
}

// SetOptimizationRules sets field value
func (o *CreateSponsoredBrandsOptimizationRulesRequestContent) SetOptimizationRules(v []CreateOptimizationRule) {
	o.OptimizationRules = v
}

func (o CreateSponsoredBrandsOptimizationRulesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRules"] = o.OptimizationRules
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsOptimizationRulesRequestContent struct {
	value *CreateSponsoredBrandsOptimizationRulesRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsOptimizationRulesRequestContent) Get() *CreateSponsoredBrandsOptimizationRulesRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsOptimizationRulesRequestContent) Set(val *CreateSponsoredBrandsOptimizationRulesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsOptimizationRulesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsOptimizationRulesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsOptimizationRulesRequestContent(val *CreateSponsoredBrandsOptimizationRulesRequestContent) *NullableCreateSponsoredBrandsOptimizationRulesRequestContent {
	return &NullableCreateSponsoredBrandsOptimizationRulesRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsOptimizationRulesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsOptimizationRulesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
