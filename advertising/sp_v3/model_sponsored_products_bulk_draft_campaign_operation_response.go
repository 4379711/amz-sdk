package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkDraftCampaignOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkDraftCampaignOperationResponse{}

// SponsoredProductsBulkDraftCampaignOperationResponse struct for SponsoredProductsBulkDraftCampaignOperationResponse
type SponsoredProductsBulkDraftCampaignOperationResponse struct {
	Success []SponsoredProductsDraftCampaignMutationSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsDraftCampaignMutationFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkDraftCampaignOperationResponse instantiates a new SponsoredProductsBulkDraftCampaignOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkDraftCampaignOperationResponse() *SponsoredProductsBulkDraftCampaignOperationResponse {
	this := SponsoredProductsBulkDraftCampaignOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkDraftCampaignOperationResponseWithDefaults instantiates a new SponsoredProductsBulkDraftCampaignOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkDraftCampaignOperationResponseWithDefaults() *SponsoredProductsBulkDraftCampaignOperationResponse {
	this := SponsoredProductsBulkDraftCampaignOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftCampaignOperationResponse) GetSuccess() []SponsoredProductsDraftCampaignMutationSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsDraftCampaignMutationSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftCampaignOperationResponse) GetSuccessOk() ([]SponsoredProductsDraftCampaignMutationSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftCampaignOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsDraftCampaignMutationSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkDraftCampaignOperationResponse) SetSuccess(v []SponsoredProductsDraftCampaignMutationSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftCampaignOperationResponse) GetError() []SponsoredProductsDraftCampaignMutationFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsDraftCampaignMutationFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftCampaignOperationResponse) GetErrorOk() ([]SponsoredProductsDraftCampaignMutationFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftCampaignOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsDraftCampaignMutationFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkDraftCampaignOperationResponse) SetError(v []SponsoredProductsDraftCampaignMutationFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkDraftCampaignOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkDraftCampaignOperationResponse struct {
	value *SponsoredProductsBulkDraftCampaignOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkDraftCampaignOperationResponse) Get() *SponsoredProductsBulkDraftCampaignOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkDraftCampaignOperationResponse) Set(val *SponsoredProductsBulkDraftCampaignOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkDraftCampaignOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkDraftCampaignOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkDraftCampaignOperationResponse(val *SponsoredProductsBulkDraftCampaignOperationResponse) *NullableSponsoredProductsBulkDraftCampaignOperationResponse {
	return &NullableSponsoredProductsBulkDraftCampaignOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkDraftCampaignOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkDraftCampaignOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
