package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BulkCreateOptimizationRuleOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkCreateOptimizationRuleOperationResponse{}

// BulkCreateOptimizationRuleOperationResponse struct for BulkCreateOptimizationRuleOperationResponse
type BulkCreateOptimizationRuleOperationResponse struct {
	Success []CreateOptimizationRuleSuccessResponseItem `json:"success,omitempty"`
	Error   []OptimizationRuleFailureResponseItem       `json:"error,omitempty"`
}

// NewBulkCreateOptimizationRuleOperationResponse instantiates a new BulkCreateOptimizationRuleOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCreateOptimizationRuleOperationResponse() *BulkCreateOptimizationRuleOperationResponse {
	this := BulkCreateOptimizationRuleOperationResponse{}
	return &this
}

// NewBulkCreateOptimizationRuleOperationResponseWithDefaults instantiates a new BulkCreateOptimizationRuleOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCreateOptimizationRuleOperationResponseWithDefaults() *BulkCreateOptimizationRuleOperationResponse {
	this := BulkCreateOptimizationRuleOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BulkCreateOptimizationRuleOperationResponse) GetSuccess() []CreateOptimizationRuleSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []CreateOptimizationRuleSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreateOptimizationRuleOperationResponse) GetSuccessOk() ([]CreateOptimizationRuleSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BulkCreateOptimizationRuleOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []CreateOptimizationRuleSuccessResponseItem and assigns it to the Success field.
func (o *BulkCreateOptimizationRuleOperationResponse) SetSuccess(v []CreateOptimizationRuleSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BulkCreateOptimizationRuleOperationResponse) GetError() []OptimizationRuleFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []OptimizationRuleFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreateOptimizationRuleOperationResponse) GetErrorOk() ([]OptimizationRuleFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BulkCreateOptimizationRuleOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []OptimizationRuleFailureResponseItem and assigns it to the Error field.
func (o *BulkCreateOptimizationRuleOperationResponse) SetError(v []OptimizationRuleFailureResponseItem) {
	o.Error = v
}

func (o BulkCreateOptimizationRuleOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableBulkCreateOptimizationRuleOperationResponse struct {
	value *BulkCreateOptimizationRuleOperationResponse
	isSet bool
}

func (v NullableBulkCreateOptimizationRuleOperationResponse) Get() *BulkCreateOptimizationRuleOperationResponse {
	return v.value
}

func (v *NullableBulkCreateOptimizationRuleOperationResponse) Set(val *BulkCreateOptimizationRuleOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCreateOptimizationRuleOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCreateOptimizationRuleOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCreateOptimizationRuleOperationResponse(val *BulkCreateOptimizationRuleOperationResponse) *NullableBulkCreateOptimizationRuleOperationResponse {
	return &NullableBulkCreateOptimizationRuleOperationResponse{value: val, isSet: true}
}

func (v NullableBulkCreateOptimizationRuleOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBulkCreateOptimizationRuleOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
