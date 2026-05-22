package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPISearchOptimizationRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPISearchOptimizationRulesResponse{}

// OptimizationRulesAPISearchOptimizationRulesResponse struct for OptimizationRulesAPISearchOptimizationRulesResponse
type OptimizationRulesAPISearchOptimizationRulesResponse struct {
	// An enumerated error code for machine use.
	Code              *string                                `json:"code,omitempty"`
	NextToken         *string                                `json:"nextToken,omitempty"`
	OptimizationRules []OptimizationRulesAPIOptimizationRule `json:"optimizationRules,omitempty"`
}

// NewOptimizationRulesAPISearchOptimizationRulesResponse instantiates a new OptimizationRulesAPISearchOptimizationRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPISearchOptimizationRulesResponse() *OptimizationRulesAPISearchOptimizationRulesResponse {
	this := OptimizationRulesAPISearchOptimizationRulesResponse{}
	return &this
}

// NewOptimizationRulesAPISearchOptimizationRulesResponseWithDefaults instantiates a new OptimizationRulesAPISearchOptimizationRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPISearchOptimizationRulesResponseWithDefaults() *OptimizationRulesAPISearchOptimizationRulesResponse {
	this := OptimizationRulesAPISearchOptimizationRulesResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) SetCode(v string) {
	o.Code = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) SetNextToken(v string) {
	o.NextToken = &v
}

// GetOptimizationRules returns the OptimizationRules field value if set, zero value otherwise.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) GetOptimizationRules() []OptimizationRulesAPIOptimizationRule {
	if o == nil || IsNil(o.OptimizationRules) {
		var ret []OptimizationRulesAPIOptimizationRule
		return ret
	}
	return o.OptimizationRules
}

// GetOptimizationRulesOk returns a tuple with the OptimizationRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) GetOptimizationRulesOk() ([]OptimizationRulesAPIOptimizationRule, bool) {
	if o == nil || IsNil(o.OptimizationRules) {
		return nil, false
	}
	return o.OptimizationRules, true
}

// HasOptimizationRules returns a boolean if a field has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) HasOptimizationRules() bool {
	if o != nil && !IsNil(o.OptimizationRules) {
		return true
	}

	return false
}

// SetOptimizationRules gets a reference to the given []OptimizationRulesAPIOptimizationRule and assigns it to the OptimizationRules field.
func (o *OptimizationRulesAPISearchOptimizationRulesResponse) SetOptimizationRules(v []OptimizationRulesAPIOptimizationRule) {
	o.OptimizationRules = v
}

func (o OptimizationRulesAPISearchOptimizationRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.OptimizationRules) {
		toSerialize["optimizationRules"] = o.OptimizationRules
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPISearchOptimizationRulesResponse struct {
	value *OptimizationRulesAPISearchOptimizationRulesResponse
	isSet bool
}

func (v NullableOptimizationRulesAPISearchOptimizationRulesResponse) Get() *OptimizationRulesAPISearchOptimizationRulesResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPISearchOptimizationRulesResponse) Set(val *OptimizationRulesAPISearchOptimizationRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPISearchOptimizationRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPISearchOptimizationRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPISearchOptimizationRulesResponse(val *OptimizationRulesAPISearchOptimizationRulesResponse) *NullableOptimizationRulesAPISearchOptimizationRulesResponse {
	return &NullableOptimizationRulesAPISearchOptimizationRulesResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPISearchOptimizationRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPISearchOptimizationRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
