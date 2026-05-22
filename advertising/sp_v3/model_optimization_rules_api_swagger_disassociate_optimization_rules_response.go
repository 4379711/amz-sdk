package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIDisassociateOptimizationRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDisassociateOptimizationRulesResponse{}

// OptimizationRulesAPIDisassociateOptimizationRulesResponse Response object for deleting optimization rules associations.
type OptimizationRulesAPIDisassociateOptimizationRulesResponse struct {
	Success []OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem `json:"success"`
	Error   []OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem   `json:"error"`
}

type _OptimizationRulesAPIDisassociateOptimizationRulesResponse OptimizationRulesAPIDisassociateOptimizationRulesResponse

// NewOptimizationRulesAPIDisassociateOptimizationRulesResponse instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDisassociateOptimizationRulesResponse(success []OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem, error_ []OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) *OptimizationRulesAPIDisassociateOptimizationRulesResponse {
	this := OptimizationRulesAPIDisassociateOptimizationRulesResponse{}
	this.Success = success
	this.Error = error_
	return &this
}

// NewOptimizationRulesAPIDisassociateOptimizationRulesResponseWithDefaults instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDisassociateOptimizationRulesResponseWithDefaults() *OptimizationRulesAPIDisassociateOptimizationRulesResponse {
	this := OptimizationRulesAPIDisassociateOptimizationRulesResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponse) GetSuccess() []OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem {
	if o == nil {
		var ret []OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponse) GetSuccessOk() ([]OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Success, true
}

// SetSuccess sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponse) SetSuccess(v []OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) {
	o.Success = v
}

// GetError returns the Error field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponse) GetError() []OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem {
	if o == nil {
		var ret []OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponse) GetErrorOk() ([]OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error, true
}

// SetError sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponse) SetError(v []OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) {
	o.Error = v
}

func (o OptimizationRulesAPIDisassociateOptimizationRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDisassociateOptimizationRulesResponse struct {
	value *OptimizationRulesAPIDisassociateOptimizationRulesResponse
	isSet bool
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesResponse) Get() *OptimizationRulesAPIDisassociateOptimizationRulesResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponse) Set(val *OptimizationRulesAPIDisassociateOptimizationRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDisassociateOptimizationRulesResponse(val *OptimizationRulesAPIDisassociateOptimizationRulesResponse) *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponse {
	return &NullableOptimizationRulesAPIDisassociateOptimizationRulesResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
