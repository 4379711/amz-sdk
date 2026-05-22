package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRuleResponse{}

// OptimizationRuleResponse struct for OptimizationRuleResponse
type OptimizationRuleResponse struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Description *string `json:"description,omitempty"`
	// The identifier of the optimization rule.
	RuleId *string `json:"ruleId,omitempty"`
}

// NewOptimizationRuleResponse instantiates a new OptimizationRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRuleResponse() *OptimizationRuleResponse {
	this := OptimizationRuleResponse{}
	return &this
}

// NewOptimizationRuleResponseWithDefaults instantiates a new OptimizationRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRuleResponseWithDefaults() *OptimizationRuleResponse {
	this := OptimizationRuleResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRuleResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRuleResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRuleResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRuleResponse) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OptimizationRuleResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRuleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OptimizationRuleResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OptimizationRuleResponse) SetDescription(v string) {
	o.Description = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *OptimizationRuleResponse) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRuleResponse) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *OptimizationRuleResponse) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *OptimizationRuleResponse) SetRuleId(v string) {
	o.RuleId = &v
}

func (o OptimizationRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.RuleId) {
		toSerialize["ruleId"] = o.RuleId
	}
	return toSerialize, nil
}

type NullableOptimizationRuleResponse struct {
	value *OptimizationRuleResponse
	isSet bool
}

func (v NullableOptimizationRuleResponse) Get() *OptimizationRuleResponse {
	return v.value
}

func (v *NullableOptimizationRuleResponse) Set(val *OptimizationRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRuleResponse(val *OptimizationRuleResponse) *NullableOptimizationRuleResponse {
	return &NullableOptimizationRuleResponse{value: val, isSet: true}
}

func (v NullableOptimizationRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
