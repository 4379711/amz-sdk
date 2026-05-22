package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIOptimizationRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIOptimizationRulesResponse{}

// OptimizationRulesAPIOptimizationRulesResponse Response object for CreateOptimizationRules and UpdateOptimizationRules API.
type OptimizationRulesAPIOptimizationRulesResponse struct {
	// An enumerated error code for machine use.
	Code      *string                                              `json:"code,omitempty"`
	Responses []OptimizationRulesAPISingleOptimizationRuleResponse `json:"responses,omitempty"`
}

// NewOptimizationRulesAPIOptimizationRulesResponse instantiates a new OptimizationRulesAPIOptimizationRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIOptimizationRulesResponse() *OptimizationRulesAPIOptimizationRulesResponse {
	this := OptimizationRulesAPIOptimizationRulesResponse{}
	return &this
}

// NewOptimizationRulesAPIOptimizationRulesResponseWithDefaults instantiates a new OptimizationRulesAPIOptimizationRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIOptimizationRulesResponseWithDefaults() *OptimizationRulesAPIOptimizationRulesResponse {
	this := OptimizationRulesAPIOptimizationRulesResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRulesResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRulesResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRulesResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRulesAPIOptimizationRulesResponse) SetCode(v string) {
	o.Code = &v
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRulesResponse) GetResponses() []OptimizationRulesAPISingleOptimizationRuleResponse {
	if o == nil || IsNil(o.Responses) {
		var ret []OptimizationRulesAPISingleOptimizationRuleResponse
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRulesResponse) GetResponsesOk() ([]OptimizationRulesAPISingleOptimizationRuleResponse, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRulesResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []OptimizationRulesAPISingleOptimizationRuleResponse and assigns it to the Responses field.
func (o *OptimizationRulesAPIOptimizationRulesResponse) SetResponses(v []OptimizationRulesAPISingleOptimizationRuleResponse) {
	o.Responses = v
}

func (o OptimizationRulesAPIOptimizationRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIOptimizationRulesResponse struct {
	value *OptimizationRulesAPIOptimizationRulesResponse
	isSet bool
}

func (v NullableOptimizationRulesAPIOptimizationRulesResponse) Get() *OptimizationRulesAPIOptimizationRulesResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPIOptimizationRulesResponse) Set(val *OptimizationRulesAPIOptimizationRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIOptimizationRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIOptimizationRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIOptimizationRulesResponse(val *OptimizationRulesAPIOptimizationRulesResponse) *NullableOptimizationRulesAPIOptimizationRulesResponse {
	return &NullableOptimizationRulesAPIOptimizationRulesResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIOptimizationRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIOptimizationRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
