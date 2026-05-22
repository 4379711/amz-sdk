package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIUpdateRuleRecommendationStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIUpdateRuleRecommendationStatusResponse{}

// OptimizationRulesAPIUpdateRuleRecommendationStatusResponse Response object for UpdateRuleRecommendationStatus API.
type OptimizationRulesAPIUpdateRuleRecommendationStatusResponse struct {
	Success []OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem `json:"success,omitempty"`
	Error   []OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem   `json:"error,omitempty"`
}

// NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponse instantiates a new OptimizationRulesAPIUpdateRuleRecommendationStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponse() *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse {
	this := OptimizationRulesAPIUpdateRuleRecommendationStatusResponse{}
	return &this
}

// NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponseWithDefaults instantiates a new OptimizationRulesAPIUpdateRuleRecommendationStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponseWithDefaults() *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse {
	this := OptimizationRulesAPIUpdateRuleRecommendationStatusResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse) GetSuccess() []OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem {
	if o == nil || IsNil(o.Success) {
		var ret []OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse) GetSuccessOk() ([]OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem and assigns it to the Success field.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse) SetSuccess(v []OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse) GetError() []OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem {
	if o == nil || IsNil(o.Error) {
		var ret []OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse) GetErrorOk() ([]OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem and assigns it to the Error field.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse) SetError(v []OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) {
	o.Error = v
}

func (o OptimizationRulesAPIUpdateRuleRecommendationStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponse struct {
	value *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse
	isSet bool
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponse) Get() *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponse) Set(val *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponse(val *OptimizationRulesAPIUpdateRuleRecommendationStatusResponse) *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponse {
	return &NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
