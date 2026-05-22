package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIAssociateOptimizationRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIAssociateOptimizationRulesResponse{}

// OptimizationRulesAPIAssociateOptimizationRulesResponse Response object for creating rules associations.
type OptimizationRulesAPIAssociateOptimizationRulesResponse struct {
	Success []OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem `json:"success"`
	Error   []OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem   `json:"error"`
}

type _OptimizationRulesAPIAssociateOptimizationRulesResponse OptimizationRulesAPIAssociateOptimizationRulesResponse

// NewOptimizationRulesAPIAssociateOptimizationRulesResponse instantiates a new OptimizationRulesAPIAssociateOptimizationRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIAssociateOptimizationRulesResponse(success []OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem, error_ []OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) *OptimizationRulesAPIAssociateOptimizationRulesResponse {
	this := OptimizationRulesAPIAssociateOptimizationRulesResponse{}
	this.Success = success
	this.Error = error_
	return &this
}

// NewOptimizationRulesAPIAssociateOptimizationRulesResponseWithDefaults instantiates a new OptimizationRulesAPIAssociateOptimizationRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIAssociateOptimizationRulesResponseWithDefaults() *OptimizationRulesAPIAssociateOptimizationRulesResponse {
	this := OptimizationRulesAPIAssociateOptimizationRulesResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponse) GetSuccess() []OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem {
	if o == nil {
		var ret []OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponse) GetSuccessOk() ([]OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Success, true
}

// SetSuccess sets field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponse) SetSuccess(v []OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) {
	o.Success = v
}

// GetError returns the Error field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponse) GetError() []OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem {
	if o == nil {
		var ret []OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponse) GetErrorOk() ([]OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error, true
}

// SetError sets field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponse) SetError(v []OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) {
	o.Error = v
}

func (o OptimizationRulesAPIAssociateOptimizationRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableOptimizationRulesAPIAssociateOptimizationRulesResponse struct {
	value *OptimizationRulesAPIAssociateOptimizationRulesResponse
	isSet bool
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesResponse) Get() *OptimizationRulesAPIAssociateOptimizationRulesResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesResponse) Set(val *OptimizationRulesAPIAssociateOptimizationRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIAssociateOptimizationRulesResponse(val *OptimizationRulesAPIAssociateOptimizationRulesResponse) *NullableOptimizationRulesAPIAssociateOptimizationRulesResponse {
	return &NullableOptimizationRulesAPIAssociateOptimizationRulesResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
