package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPICreateOptimizationRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPICreateOptimizationRulesResponse{}

// OptimizationRulesAPICreateOptimizationRulesResponse Response object for CreateOptimizationRules API.
type OptimizationRulesAPICreateOptimizationRulesResponse struct {
	// An enumerated error code for machine use.
	Code      *string                                              `json:"code,omitempty"`
	Responses []OptimizationRulesAPISingleOptimizationRuleResponse `json:"responses,omitempty"`
}

// NewOptimizationRulesAPICreateOptimizationRulesResponse instantiates a new OptimizationRulesAPICreateOptimizationRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPICreateOptimizationRulesResponse() *OptimizationRulesAPICreateOptimizationRulesResponse {
	this := OptimizationRulesAPICreateOptimizationRulesResponse{}
	return &this
}

// NewOptimizationRulesAPICreateOptimizationRulesResponseWithDefaults instantiates a new OptimizationRulesAPICreateOptimizationRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPICreateOptimizationRulesResponseWithDefaults() *OptimizationRulesAPICreateOptimizationRulesResponse {
	this := OptimizationRulesAPICreateOptimizationRulesResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRulesAPICreateOptimizationRulesResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPICreateOptimizationRulesResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRulesAPICreateOptimizationRulesResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRulesAPICreateOptimizationRulesResponse) SetCode(v string) {
	o.Code = &v
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *OptimizationRulesAPICreateOptimizationRulesResponse) GetResponses() []OptimizationRulesAPISingleOptimizationRuleResponse {
	if o == nil || IsNil(o.Responses) {
		var ret []OptimizationRulesAPISingleOptimizationRuleResponse
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPICreateOptimizationRulesResponse) GetResponsesOk() ([]OptimizationRulesAPISingleOptimizationRuleResponse, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *OptimizationRulesAPICreateOptimizationRulesResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []OptimizationRulesAPISingleOptimizationRuleResponse and assigns it to the Responses field.
func (o *OptimizationRulesAPICreateOptimizationRulesResponse) SetResponses(v []OptimizationRulesAPISingleOptimizationRuleResponse) {
	o.Responses = v
}

func (o OptimizationRulesAPICreateOptimizationRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPICreateOptimizationRulesResponse struct {
	value *OptimizationRulesAPICreateOptimizationRulesResponse
	isSet bool
}

func (v NullableOptimizationRulesAPICreateOptimizationRulesResponse) Get() *OptimizationRulesAPICreateOptimizationRulesResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPICreateOptimizationRulesResponse) Set(val *OptimizationRulesAPICreateOptimizationRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPICreateOptimizationRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPICreateOptimizationRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPICreateOptimizationRulesResponse(val *OptimizationRulesAPICreateOptimizationRulesResponse) *NullableOptimizationRulesAPICreateOptimizationRulesResponse {
	return &NullableOptimizationRulesAPICreateOptimizationRulesResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPICreateOptimizationRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPICreateOptimizationRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
