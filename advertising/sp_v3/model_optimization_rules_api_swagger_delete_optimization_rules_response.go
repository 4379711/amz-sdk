package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIDeleteOptimizationRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDeleteOptimizationRulesResponse{}

// OptimizationRulesAPIDeleteOptimizationRulesResponse Response object from deleting optimization rules.
type OptimizationRulesAPIDeleteOptimizationRulesResponse struct {
	Success []OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem `json:"success,omitempty"`
	Error   []OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem   `json:"error,omitempty"`
}

// NewOptimizationRulesAPIDeleteOptimizationRulesResponse instantiates a new OptimizationRulesAPIDeleteOptimizationRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDeleteOptimizationRulesResponse() *OptimizationRulesAPIDeleteOptimizationRulesResponse {
	this := OptimizationRulesAPIDeleteOptimizationRulesResponse{}
	return &this
}

// NewOptimizationRulesAPIDeleteOptimizationRulesResponseWithDefaults instantiates a new OptimizationRulesAPIDeleteOptimizationRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDeleteOptimizationRulesResponseWithDefaults() *OptimizationRulesAPIDeleteOptimizationRulesResponse {
	this := OptimizationRulesAPIDeleteOptimizationRulesResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponse) GetSuccess() []OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem {
	if o == nil || IsNil(o.Success) {
		var ret []OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponse) GetSuccessOk() ([]OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem and assigns it to the Success field.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponse) SetSuccess(v []OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponse) GetError() []OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem {
	if o == nil || IsNil(o.Error) {
		var ret []OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponse) GetErrorOk() ([]OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem and assigns it to the Error field.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponse) SetError(v []OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) {
	o.Error = v
}

func (o OptimizationRulesAPIDeleteOptimizationRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDeleteOptimizationRulesResponse struct {
	value *OptimizationRulesAPIDeleteOptimizationRulesResponse
	isSet bool
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesResponse) Get() *OptimizationRulesAPIDeleteOptimizationRulesResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesResponse) Set(val *OptimizationRulesAPIDeleteOptimizationRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDeleteOptimizationRulesResponse(val *OptimizationRulesAPIDeleteOptimizationRulesResponse) *NullableOptimizationRulesAPIDeleteOptimizationRulesResponse {
	return &NullableOptimizationRulesAPIDeleteOptimizationRulesResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
