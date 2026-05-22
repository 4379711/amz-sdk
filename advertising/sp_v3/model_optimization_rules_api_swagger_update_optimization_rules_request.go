package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIUpdateOptimizationRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIUpdateOptimizationRulesRequest{}

// OptimizationRulesAPIUpdateOptimizationRulesRequest Request object for updating one or multiple optimization rules.
type OptimizationRulesAPIUpdateOptimizationRulesRequest struct {
	OptimizationRules []OptimizationRulesAPIOptimizationRule `json:"optimizationRules"`
}

type _OptimizationRulesAPIUpdateOptimizationRulesRequest OptimizationRulesAPIUpdateOptimizationRulesRequest

// NewOptimizationRulesAPIUpdateOptimizationRulesRequest instantiates a new OptimizationRulesAPIUpdateOptimizationRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIUpdateOptimizationRulesRequest(optimizationRules []OptimizationRulesAPIOptimizationRule) *OptimizationRulesAPIUpdateOptimizationRulesRequest {
	this := OptimizationRulesAPIUpdateOptimizationRulesRequest{}
	this.OptimizationRules = optimizationRules
	return &this
}

// NewOptimizationRulesAPIUpdateOptimizationRulesRequestWithDefaults instantiates a new OptimizationRulesAPIUpdateOptimizationRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIUpdateOptimizationRulesRequestWithDefaults() *OptimizationRulesAPIUpdateOptimizationRulesRequest {
	this := OptimizationRulesAPIUpdateOptimizationRulesRequest{}
	return &this
}

// GetOptimizationRules returns the OptimizationRules field value
func (o *OptimizationRulesAPIUpdateOptimizationRulesRequest) GetOptimizationRules() []OptimizationRulesAPIOptimizationRule {
	if o == nil {
		var ret []OptimizationRulesAPIOptimizationRule
		return ret
	}

	return o.OptimizationRules
}

// GetOptimizationRulesOk returns a tuple with the OptimizationRules field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateOptimizationRulesRequest) GetOptimizationRulesOk() ([]OptimizationRulesAPIOptimizationRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimizationRules, true
}

// SetOptimizationRules sets field value
func (o *OptimizationRulesAPIUpdateOptimizationRulesRequest) SetOptimizationRules(v []OptimizationRulesAPIOptimizationRule) {
	o.OptimizationRules = v
}

func (o OptimizationRulesAPIUpdateOptimizationRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRules"] = o.OptimizationRules
	return toSerialize, nil
}

type NullableOptimizationRulesAPIUpdateOptimizationRulesRequest struct {
	value *OptimizationRulesAPIUpdateOptimizationRulesRequest
	isSet bool
}

func (v NullableOptimizationRulesAPIUpdateOptimizationRulesRequest) Get() *OptimizationRulesAPIUpdateOptimizationRulesRequest {
	return v.value
}

func (v *NullableOptimizationRulesAPIUpdateOptimizationRulesRequest) Set(val *OptimizationRulesAPIUpdateOptimizationRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIUpdateOptimizationRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIUpdateOptimizationRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIUpdateOptimizationRulesRequest(val *OptimizationRulesAPIUpdateOptimizationRulesRequest) *NullableOptimizationRulesAPIUpdateOptimizationRulesRequest {
	return &NullableOptimizationRulesAPIUpdateOptimizationRulesRequest{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIUpdateOptimizationRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIUpdateOptimizationRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
