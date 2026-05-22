package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse{}

// OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse Response object for deleting campaign to optimization rules association.
type OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse struct {
	Success []OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem `json:"success,omitempty"`
	Error   []OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem   `json:"error,omitempty"`
}

// NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse() *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse {
	this := OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse{}
	return &this
}

// NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseWithDefaults instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseWithDefaults() *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse {
	this := OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) GetSuccess() []OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem {
	if o == nil || IsNil(o.Success) {
		var ret []OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) GetSuccessOk() ([]OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem and assigns it to the Success field.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) SetSuccess(v []OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) GetError() []OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem {
	if o == nil || IsNil(o.Error) {
		var ret []OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) GetErrorOk() ([]OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem and assigns it to the Error field.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) SetError(v []OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) {
	o.Error = v
}

func (o OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse struct {
	value *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse
	isSet bool
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) Get() *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) Set(val *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse(val *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse {
	return &NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
