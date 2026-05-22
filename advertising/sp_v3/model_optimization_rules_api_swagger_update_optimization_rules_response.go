package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIUpdateOptimizationRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIUpdateOptimizationRulesResponse{}

// OptimizationRulesAPIUpdateOptimizationRulesResponse Response object for UpdateOptimizationRules API.
type OptimizationRulesAPIUpdateOptimizationRulesResponse struct {
	// An enumerated error code for machine use.
	Code      *string                                              `json:"code,omitempty"`
	Responses []OptimizationRulesAPISingleOptimizationRuleResponse `json:"responses,omitempty"`
}

// NewOptimizationRulesAPIUpdateOptimizationRulesResponse instantiates a new OptimizationRulesAPIUpdateOptimizationRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIUpdateOptimizationRulesResponse() *OptimizationRulesAPIUpdateOptimizationRulesResponse {
	this := OptimizationRulesAPIUpdateOptimizationRulesResponse{}
	return &this
}

// NewOptimizationRulesAPIUpdateOptimizationRulesResponseWithDefaults instantiates a new OptimizationRulesAPIUpdateOptimizationRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIUpdateOptimizationRulesResponseWithDefaults() *OptimizationRulesAPIUpdateOptimizationRulesResponse {
	this := OptimizationRulesAPIUpdateOptimizationRulesResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRulesAPIUpdateOptimizationRulesResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateOptimizationRulesResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRulesAPIUpdateOptimizationRulesResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRulesAPIUpdateOptimizationRulesResponse) SetCode(v string) {
	o.Code = &v
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *OptimizationRulesAPIUpdateOptimizationRulesResponse) GetResponses() []OptimizationRulesAPISingleOptimizationRuleResponse {
	if o == nil || IsNil(o.Responses) {
		var ret []OptimizationRulesAPISingleOptimizationRuleResponse
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateOptimizationRulesResponse) GetResponsesOk() ([]OptimizationRulesAPISingleOptimizationRuleResponse, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *OptimizationRulesAPIUpdateOptimizationRulesResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []OptimizationRulesAPISingleOptimizationRuleResponse and assigns it to the Responses field.
func (o *OptimizationRulesAPIUpdateOptimizationRulesResponse) SetResponses(v []OptimizationRulesAPISingleOptimizationRuleResponse) {
	o.Responses = v
}

func (o OptimizationRulesAPIUpdateOptimizationRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIUpdateOptimizationRulesResponse struct {
	value *OptimizationRulesAPIUpdateOptimizationRulesResponse
	isSet bool
}

func (v NullableOptimizationRulesAPIUpdateOptimizationRulesResponse) Get() *OptimizationRulesAPIUpdateOptimizationRulesResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPIUpdateOptimizationRulesResponse) Set(val *OptimizationRulesAPIUpdateOptimizationRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIUpdateOptimizationRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIUpdateOptimizationRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIUpdateOptimizationRulesResponse(val *OptimizationRulesAPIUpdateOptimizationRulesResponse) *NullableOptimizationRulesAPIUpdateOptimizationRulesResponse {
	return &NullableOptimizationRulesAPIUpdateOptimizationRulesResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIUpdateOptimizationRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIUpdateOptimizationRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
