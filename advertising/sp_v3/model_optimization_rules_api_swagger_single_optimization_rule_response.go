package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPISingleOptimizationRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPISingleOptimizationRuleResponse{}

// OptimizationRulesAPISingleOptimizationRuleResponse Response object for operations involving a single optimization rule.
type OptimizationRulesAPISingleOptimizationRuleResponse struct {
	// An enumerated success or error code for machine use.
	Code             *string                               `json:"code,omitempty"`
	OptimizationRule *OptimizationRulesAPIOptimizationRule `json:"optimizationRule,omitempty"`
	// A human-readable description of the error, if unsuccessful.
	Details *string `json:"details,omitempty"`
}

// NewOptimizationRulesAPISingleOptimizationRuleResponse instantiates a new OptimizationRulesAPISingleOptimizationRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPISingleOptimizationRuleResponse() *OptimizationRulesAPISingleOptimizationRuleResponse {
	this := OptimizationRulesAPISingleOptimizationRuleResponse{}
	return &this
}

// NewOptimizationRulesAPISingleOptimizationRuleResponseWithDefaults instantiates a new OptimizationRulesAPISingleOptimizationRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPISingleOptimizationRuleResponseWithDefaults() *OptimizationRulesAPISingleOptimizationRuleResponse {
	this := OptimizationRulesAPISingleOptimizationRuleResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) SetCode(v string) {
	o.Code = &v
}

// GetOptimizationRule returns the OptimizationRule field value if set, zero value otherwise.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) GetOptimizationRule() OptimizationRulesAPIOptimizationRule {
	if o == nil || IsNil(o.OptimizationRule) {
		var ret OptimizationRulesAPIOptimizationRule
		return ret
	}
	return *o.OptimizationRule
}

// GetOptimizationRuleOk returns a tuple with the OptimizationRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) GetOptimizationRuleOk() (*OptimizationRulesAPIOptimizationRule, bool) {
	if o == nil || IsNil(o.OptimizationRule) {
		return nil, false
	}
	return o.OptimizationRule, true
}

// HasOptimizationRule returns a boolean if a field has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) HasOptimizationRule() bool {
	if o != nil && !IsNil(o.OptimizationRule) {
		return true
	}

	return false
}

// SetOptimizationRule gets a reference to the given OptimizationRulesAPIOptimizationRule and assigns it to the OptimizationRule field.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) SetOptimizationRule(v OptimizationRulesAPIOptimizationRule) {
	o.OptimizationRule = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *OptimizationRulesAPISingleOptimizationRuleResponse) SetDetails(v string) {
	o.Details = &v
}

func (o OptimizationRulesAPISingleOptimizationRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.OptimizationRule) {
		toSerialize["optimizationRule"] = o.OptimizationRule
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPISingleOptimizationRuleResponse struct {
	value *OptimizationRulesAPISingleOptimizationRuleResponse
	isSet bool
}

func (v NullableOptimizationRulesAPISingleOptimizationRuleResponse) Get() *OptimizationRulesAPISingleOptimizationRuleResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPISingleOptimizationRuleResponse) Set(val *OptimizationRulesAPISingleOptimizationRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPISingleOptimizationRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPISingleOptimizationRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPISingleOptimizationRuleResponse(val *OptimizationRulesAPISingleOptimizationRuleResponse) *NullableOptimizationRulesAPISingleOptimizationRuleResponse {
	return &NullableOptimizationRulesAPISingleOptimizationRuleResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPISingleOptimizationRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPISingleOptimizationRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
