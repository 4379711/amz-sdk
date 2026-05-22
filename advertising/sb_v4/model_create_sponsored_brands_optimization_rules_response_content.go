package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsOptimizationRulesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsOptimizationRulesResponseContent{}

// CreateSponsoredBrandsOptimizationRulesResponseContent struct for CreateSponsoredBrandsOptimizationRulesResponseContent
type CreateSponsoredBrandsOptimizationRulesResponseContent struct {
	OptimizationRules BulkCreateOptimizationRuleOperationResponse `json:"optimizationRules"`
}

type _CreateSponsoredBrandsOptimizationRulesResponseContent CreateSponsoredBrandsOptimizationRulesResponseContent

// NewCreateSponsoredBrandsOptimizationRulesResponseContent instantiates a new CreateSponsoredBrandsOptimizationRulesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsOptimizationRulesResponseContent(optimizationRules BulkCreateOptimizationRuleOperationResponse) *CreateSponsoredBrandsOptimizationRulesResponseContent {
	this := CreateSponsoredBrandsOptimizationRulesResponseContent{}
	this.OptimizationRules = optimizationRules
	return &this
}

// NewCreateSponsoredBrandsOptimizationRulesResponseContentWithDefaults instantiates a new CreateSponsoredBrandsOptimizationRulesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsOptimizationRulesResponseContentWithDefaults() *CreateSponsoredBrandsOptimizationRulesResponseContent {
	this := CreateSponsoredBrandsOptimizationRulesResponseContent{}
	return &this
}

// GetOptimizationRules returns the OptimizationRules field value
func (o *CreateSponsoredBrandsOptimizationRulesResponseContent) GetOptimizationRules() BulkCreateOptimizationRuleOperationResponse {
	if o == nil {
		var ret BulkCreateOptimizationRuleOperationResponse
		return ret
	}

	return o.OptimizationRules
}

// GetOptimizationRulesOk returns a tuple with the OptimizationRules field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsOptimizationRulesResponseContent) GetOptimizationRulesOk() (*BulkCreateOptimizationRuleOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRules, true
}

// SetOptimizationRules sets field value
func (o *CreateSponsoredBrandsOptimizationRulesResponseContent) SetOptimizationRules(v BulkCreateOptimizationRuleOperationResponse) {
	o.OptimizationRules = v
}

func (o CreateSponsoredBrandsOptimizationRulesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRules"] = o.OptimizationRules
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsOptimizationRulesResponseContent struct {
	value *CreateSponsoredBrandsOptimizationRulesResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsOptimizationRulesResponseContent) Get() *CreateSponsoredBrandsOptimizationRulesResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsOptimizationRulesResponseContent) Set(val *CreateSponsoredBrandsOptimizationRulesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsOptimizationRulesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsOptimizationRulesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsOptimizationRulesResponseContent(val *CreateSponsoredBrandsOptimizationRulesResponseContent) *NullableCreateSponsoredBrandsOptimizationRulesResponseContent {
	return &NullableCreateSponsoredBrandsOptimizationRulesResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsOptimizationRulesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsOptimizationRulesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
