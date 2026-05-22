package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BulkCampaignOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkCampaignOperationResponse{}

// BulkCampaignOperationResponse struct for BulkCampaignOperationResponse
type BulkCampaignOperationResponse struct {
	Success []CampaignMutationSuccessResponseItem `json:"success,omitempty"`
	Error   []CampaignMutationFailureResponseItem `json:"error,omitempty"`
}

// NewBulkCampaignOperationResponse instantiates a new BulkCampaignOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCampaignOperationResponse() *BulkCampaignOperationResponse {
	this := BulkCampaignOperationResponse{}
	return &this
}

// NewBulkCampaignOperationResponseWithDefaults instantiates a new BulkCampaignOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCampaignOperationResponseWithDefaults() *BulkCampaignOperationResponse {
	this := BulkCampaignOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BulkCampaignOperationResponse) GetSuccess() []CampaignMutationSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []CampaignMutationSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCampaignOperationResponse) GetSuccessOk() ([]CampaignMutationSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BulkCampaignOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []CampaignMutationSuccessResponseItem and assigns it to the Success field.
func (o *BulkCampaignOperationResponse) SetSuccess(v []CampaignMutationSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BulkCampaignOperationResponse) GetError() []CampaignMutationFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []CampaignMutationFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCampaignOperationResponse) GetErrorOk() ([]CampaignMutationFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BulkCampaignOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []CampaignMutationFailureResponseItem and assigns it to the Error field.
func (o *BulkCampaignOperationResponse) SetError(v []CampaignMutationFailureResponseItem) {
	o.Error = v
}

func (o BulkCampaignOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableBulkCampaignOperationResponse struct {
	value *BulkCampaignOperationResponse
	isSet bool
}

func (v NullableBulkCampaignOperationResponse) Get() *BulkCampaignOperationResponse {
	return v.value
}

func (v *NullableBulkCampaignOperationResponse) Set(val *BulkCampaignOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCampaignOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCampaignOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCampaignOperationResponse(val *BulkCampaignOperationResponse) *NullableBulkCampaignOperationResponse {
	return &NullableBulkCampaignOperationResponse{value: val, isSet: true}
}

func (v NullableBulkCampaignOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBulkCampaignOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
