package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BulkUpdateOptimizationRuleOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkUpdateOptimizationRuleOperationResponse{}

// BulkUpdateOptimizationRuleOperationResponse struct for BulkUpdateOptimizationRuleOperationResponse
type BulkUpdateOptimizationRuleOperationResponse struct {
	Success []UpdateOptimizationRuleSuccessResponseItem `json:"success,omitempty"`
	Error   []OptimizationRuleFailureResponseItem       `json:"error,omitempty"`
}

// NewBulkUpdateOptimizationRuleOperationResponse instantiates a new BulkUpdateOptimizationRuleOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateOptimizationRuleOperationResponse() *BulkUpdateOptimizationRuleOperationResponse {
	this := BulkUpdateOptimizationRuleOperationResponse{}
	return &this
}

// NewBulkUpdateOptimizationRuleOperationResponseWithDefaults instantiates a new BulkUpdateOptimizationRuleOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateOptimizationRuleOperationResponseWithDefaults() *BulkUpdateOptimizationRuleOperationResponse {
	this := BulkUpdateOptimizationRuleOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BulkUpdateOptimizationRuleOperationResponse) GetSuccess() []UpdateOptimizationRuleSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []UpdateOptimizationRuleSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateOptimizationRuleOperationResponse) GetSuccessOk() ([]UpdateOptimizationRuleSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BulkUpdateOptimizationRuleOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []UpdateOptimizationRuleSuccessResponseItem and assigns it to the Success field.
func (o *BulkUpdateOptimizationRuleOperationResponse) SetSuccess(v []UpdateOptimizationRuleSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BulkUpdateOptimizationRuleOperationResponse) GetError() []OptimizationRuleFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []OptimizationRuleFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateOptimizationRuleOperationResponse) GetErrorOk() ([]OptimizationRuleFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BulkUpdateOptimizationRuleOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []OptimizationRuleFailureResponseItem and assigns it to the Error field.
func (o *BulkUpdateOptimizationRuleOperationResponse) SetError(v []OptimizationRuleFailureResponseItem) {
	o.Error = v
}

func (o BulkUpdateOptimizationRuleOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableBulkUpdateOptimizationRuleOperationResponse struct {
	value *BulkUpdateOptimizationRuleOperationResponse
	isSet bool
}

func (v NullableBulkUpdateOptimizationRuleOperationResponse) Get() *BulkUpdateOptimizationRuleOperationResponse {
	return v.value
}

func (v *NullableBulkUpdateOptimizationRuleOperationResponse) Set(val *BulkUpdateOptimizationRuleOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdateOptimizationRuleOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdateOptimizationRuleOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdateOptimizationRuleOperationResponse(val *BulkUpdateOptimizationRuleOperationResponse) *NullableBulkUpdateOptimizationRuleOperationResponse {
	return &NullableBulkUpdateOptimizationRuleOperationResponse{value: val, isSet: true}
}

func (v NullableBulkUpdateOptimizationRuleOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBulkUpdateOptimizationRuleOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
