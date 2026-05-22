package sp_v3

import "github.com/bytedance/sonic"

// checks if the OptimizationRulesAPIDeleteOptimizationRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDeleteOptimizationRulesRequest{}

// OptimizationRulesAPIDeleteOptimizationRulesRequest Request body for deleting optimization rules.
type OptimizationRulesAPIDeleteOptimizationRulesRequest struct {
	// An array of rule identifiers.
	OptimizationRuleIds []string `json:"optimizationRuleIds"`
}

type _OptimizationRulesAPIDeleteOptimizationRulesRequest OptimizationRulesAPIDeleteOptimizationRulesRequest

// NewOptimizationRulesAPIDeleteOptimizationRulesRequest instantiates a new OptimizationRulesAPIDeleteOptimizationRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDeleteOptimizationRulesRequest(optimizationRuleIds []string) *OptimizationRulesAPIDeleteOptimizationRulesRequest {
	this := OptimizationRulesAPIDeleteOptimizationRulesRequest{}
	this.OptimizationRuleIds = optimizationRuleIds
	return &this
}

// NewOptimizationRulesAPIDeleteOptimizationRulesRequestWithDefaults instantiates a new OptimizationRulesAPIDeleteOptimizationRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDeleteOptimizationRulesRequestWithDefaults() *OptimizationRulesAPIDeleteOptimizationRulesRequest {
	this := OptimizationRulesAPIDeleteOptimizationRulesRequest{}
	return &this
}

// GetOptimizationRuleIds returns the OptimizationRuleIds field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesRequest) GetOptimizationRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OptimizationRuleIds
}

// GetOptimizationRuleIdsOk returns a tuple with the OptimizationRuleIds field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDeleteOptimizationRulesRequest) GetOptimizationRuleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimizationRuleIds, true
}

// SetOptimizationRuleIds sets field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesRequest) SetOptimizationRuleIds(v []string) {
	o.OptimizationRuleIds = v
}

func (o OptimizationRulesAPIDeleteOptimizationRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRuleIds"] = o.OptimizationRuleIds
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDeleteOptimizationRulesRequest struct {
	value *OptimizationRulesAPIDeleteOptimizationRulesRequest
	isSet bool
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesRequest) Get() *OptimizationRulesAPIDeleteOptimizationRulesRequest {
	return v.value
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesRequest) Set(val *OptimizationRulesAPIDeleteOptimizationRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDeleteOptimizationRulesRequest(val *OptimizationRulesAPIDeleteOptimizationRulesRequest) *NullableOptimizationRulesAPIDeleteOptimizationRulesRequest {
	return &NullableOptimizationRulesAPIDeleteOptimizationRulesRequest{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
