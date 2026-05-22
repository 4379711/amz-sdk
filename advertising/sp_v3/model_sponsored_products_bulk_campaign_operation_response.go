package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkCampaignOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkCampaignOperationResponse{}

// SponsoredProductsBulkCampaignOperationResponse struct for SponsoredProductsBulkCampaignOperationResponse
type SponsoredProductsBulkCampaignOperationResponse struct {
	Success []SponsoredProductsCampaignMutationSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsCampaignMutationFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkCampaignOperationResponse instantiates a new SponsoredProductsBulkCampaignOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkCampaignOperationResponse() *SponsoredProductsBulkCampaignOperationResponse {
	this := SponsoredProductsBulkCampaignOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkCampaignOperationResponseWithDefaults instantiates a new SponsoredProductsBulkCampaignOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkCampaignOperationResponseWithDefaults() *SponsoredProductsBulkCampaignOperationResponse {
	this := SponsoredProductsBulkCampaignOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkCampaignOperationResponse) GetSuccess() []SponsoredProductsCampaignMutationSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsCampaignMutationSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkCampaignOperationResponse) GetSuccessOk() ([]SponsoredProductsCampaignMutationSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkCampaignOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsCampaignMutationSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkCampaignOperationResponse) SetSuccess(v []SponsoredProductsCampaignMutationSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkCampaignOperationResponse) GetError() []SponsoredProductsCampaignMutationFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsCampaignMutationFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkCampaignOperationResponse) GetErrorOk() ([]SponsoredProductsCampaignMutationFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkCampaignOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsCampaignMutationFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkCampaignOperationResponse) SetError(v []SponsoredProductsCampaignMutationFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkCampaignOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkCampaignOperationResponse struct {
	value *SponsoredProductsBulkCampaignOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkCampaignOperationResponse) Get() *SponsoredProductsBulkCampaignOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkCampaignOperationResponse) Set(val *SponsoredProductsBulkCampaignOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkCampaignOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkCampaignOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkCampaignOperationResponse(val *SponsoredProductsBulkCampaignOperationResponse) *NullableSponsoredProductsBulkCampaignOperationResponse {
	return &NullableSponsoredProductsBulkCampaignOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkCampaignOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkCampaignOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
