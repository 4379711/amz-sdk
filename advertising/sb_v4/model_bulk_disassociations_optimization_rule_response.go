package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BulkDisassociationsOptimizationRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkDisassociationsOptimizationRuleResponse{}

// BulkDisassociationsOptimizationRuleResponse struct for BulkDisassociationsOptimizationRuleResponse
type BulkDisassociationsOptimizationRuleResponse struct {
	Success []OptimizationRuleToEntityMappingSuccessResponseItem `json:"success,omitempty"`
	Error   []OptimizationRuleFailureResponseItem                `json:"error,omitempty"`
}

// NewBulkDisassociationsOptimizationRuleResponse instantiates a new BulkDisassociationsOptimizationRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkDisassociationsOptimizationRuleResponse() *BulkDisassociationsOptimizationRuleResponse {
	this := BulkDisassociationsOptimizationRuleResponse{}
	return &this
}

// NewBulkDisassociationsOptimizationRuleResponseWithDefaults instantiates a new BulkDisassociationsOptimizationRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkDisassociationsOptimizationRuleResponseWithDefaults() *BulkDisassociationsOptimizationRuleResponse {
	this := BulkDisassociationsOptimizationRuleResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BulkDisassociationsOptimizationRuleResponse) GetSuccess() []OptimizationRuleToEntityMappingSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []OptimizationRuleToEntityMappingSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkDisassociationsOptimizationRuleResponse) GetSuccessOk() ([]OptimizationRuleToEntityMappingSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BulkDisassociationsOptimizationRuleResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []OptimizationRuleToEntityMappingSuccessResponseItem and assigns it to the Success field.
func (o *BulkDisassociationsOptimizationRuleResponse) SetSuccess(v []OptimizationRuleToEntityMappingSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BulkDisassociationsOptimizationRuleResponse) GetError() []OptimizationRuleFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []OptimizationRuleFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkDisassociationsOptimizationRuleResponse) GetErrorOk() ([]OptimizationRuleFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BulkDisassociationsOptimizationRuleResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []OptimizationRuleFailureResponseItem and assigns it to the Error field.
func (o *BulkDisassociationsOptimizationRuleResponse) SetError(v []OptimizationRuleFailureResponseItem) {
	o.Error = v
}

func (o BulkDisassociationsOptimizationRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableBulkDisassociationsOptimizationRuleResponse struct {
	value *BulkDisassociationsOptimizationRuleResponse
	isSet bool
}

func (v NullableBulkDisassociationsOptimizationRuleResponse) Get() *BulkDisassociationsOptimizationRuleResponse {
	return v.value
}

func (v *NullableBulkDisassociationsOptimizationRuleResponse) Set(val *BulkDisassociationsOptimizationRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkDisassociationsOptimizationRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkDisassociationsOptimizationRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkDisassociationsOptimizationRuleResponse(val *BulkDisassociationsOptimizationRuleResponse) *NullableBulkDisassociationsOptimizationRuleResponse {
	return &NullableBulkDisassociationsOptimizationRuleResponse{value: val, isSet: true}
}

func (v NullableBulkDisassociationsOptimizationRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBulkDisassociationsOptimizationRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
