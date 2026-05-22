package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse{}

// SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse struct for SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse
type SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse struct {
	Success []SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse instantiates a new SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse() *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse {
	this := SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponseWithDefaults instantiates a new SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponseWithDefaults() *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse {
	this := SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) GetSuccess() []SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) GetSuccessOk() ([]SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) SetSuccess(v []SponsoredProductsDraftCampaignNegativeKeywordSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) GetError() []SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) GetErrorOk() ([]SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) SetError(v []SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse struct {
	value *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) Get() *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) Set(val *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse(val *SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) *NullableSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse {
	return &NullableSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
