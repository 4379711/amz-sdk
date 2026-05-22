package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPICreateOptimizationRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPICreateOptimizationRulesRequest{}

// OptimizationRulesAPICreateOptimizationRulesRequest Request object for creating one or multiple optimization rules.
type OptimizationRulesAPICreateOptimizationRulesRequest struct {
	OptimizationRules []OptimizationRulesAPIOptimizationRuleWithoutRuleId `json:"optimizationRules"`
}

type _OptimizationRulesAPICreateOptimizationRulesRequest OptimizationRulesAPICreateOptimizationRulesRequest

// NewOptimizationRulesAPICreateOptimizationRulesRequest instantiates a new OptimizationRulesAPICreateOptimizationRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPICreateOptimizationRulesRequest(optimizationRules []OptimizationRulesAPIOptimizationRuleWithoutRuleId) *OptimizationRulesAPICreateOptimizationRulesRequest {
	this := OptimizationRulesAPICreateOptimizationRulesRequest{}
	this.OptimizationRules = optimizationRules
	return &this
}

// NewOptimizationRulesAPICreateOptimizationRulesRequestWithDefaults instantiates a new OptimizationRulesAPICreateOptimizationRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPICreateOptimizationRulesRequestWithDefaults() *OptimizationRulesAPICreateOptimizationRulesRequest {
	this := OptimizationRulesAPICreateOptimizationRulesRequest{}
	return &this
}

// GetOptimizationRules returns the OptimizationRules field value
func (o *OptimizationRulesAPICreateOptimizationRulesRequest) GetOptimizationRules() []OptimizationRulesAPIOptimizationRuleWithoutRuleId {
	if o == nil {
		var ret []OptimizationRulesAPIOptimizationRuleWithoutRuleId
		return ret
	}

	return o.OptimizationRules
}

// GetOptimizationRulesOk returns a tuple with the OptimizationRules field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPICreateOptimizationRulesRequest) GetOptimizationRulesOk() ([]OptimizationRulesAPIOptimizationRuleWithoutRuleId, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimizationRules, true
}

// SetOptimizationRules sets field value
func (o *OptimizationRulesAPICreateOptimizationRulesRequest) SetOptimizationRules(v []OptimizationRulesAPIOptimizationRuleWithoutRuleId) {
	o.OptimizationRules = v
}

func (o OptimizationRulesAPICreateOptimizationRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRules"] = o.OptimizationRules
	return toSerialize, nil
}

type NullableOptimizationRulesAPICreateOptimizationRulesRequest struct {
	value *OptimizationRulesAPICreateOptimizationRulesRequest
	isSet bool
}

func (v NullableOptimizationRulesAPICreateOptimizationRulesRequest) Get() *OptimizationRulesAPICreateOptimizationRulesRequest {
	return v.value
}

func (v *NullableOptimizationRulesAPICreateOptimizationRulesRequest) Set(val *OptimizationRulesAPICreateOptimizationRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPICreateOptimizationRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPICreateOptimizationRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPICreateOptimizationRulesRequest(val *OptimizationRulesAPICreateOptimizationRulesRequest) *NullableOptimizationRulesAPICreateOptimizationRulesRequest {
	return &NullableOptimizationRulesAPICreateOptimizationRulesRequest{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPICreateOptimizationRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPICreateOptimizationRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
