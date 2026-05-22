package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkCampaignNegativeKeywordOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkCampaignNegativeKeywordOperationResponse{}

// SponsoredProductsBulkCampaignNegativeKeywordOperationResponse struct for SponsoredProductsBulkCampaignNegativeKeywordOperationResponse
type SponsoredProductsBulkCampaignNegativeKeywordOperationResponse struct {
	Success []SponsoredProductsCampaignNegativeKeywordSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsCampaignNegativeKeywordFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkCampaignNegativeKeywordOperationResponse instantiates a new SponsoredProductsBulkCampaignNegativeKeywordOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkCampaignNegativeKeywordOperationResponse() *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse {
	this := SponsoredProductsBulkCampaignNegativeKeywordOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkCampaignNegativeKeywordOperationResponseWithDefaults instantiates a new SponsoredProductsBulkCampaignNegativeKeywordOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkCampaignNegativeKeywordOperationResponseWithDefaults() *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse {
	this := SponsoredProductsBulkCampaignNegativeKeywordOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) GetSuccess() []SponsoredProductsCampaignNegativeKeywordSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsCampaignNegativeKeywordSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) GetSuccessOk() ([]SponsoredProductsCampaignNegativeKeywordSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsCampaignNegativeKeywordSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) SetSuccess(v []SponsoredProductsCampaignNegativeKeywordSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) GetError() []SponsoredProductsCampaignNegativeKeywordFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsCampaignNegativeKeywordFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) GetErrorOk() ([]SponsoredProductsCampaignNegativeKeywordFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsCampaignNegativeKeywordFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) SetError(v []SponsoredProductsCampaignNegativeKeywordFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkCampaignNegativeKeywordOperationResponse struct {
	value *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkCampaignNegativeKeywordOperationResponse) Get() *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkCampaignNegativeKeywordOperationResponse) Set(val *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkCampaignNegativeKeywordOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkCampaignNegativeKeywordOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkCampaignNegativeKeywordOperationResponse(val *SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) *NullableSponsoredProductsBulkCampaignNegativeKeywordOperationResponse {
	return &NullableSponsoredProductsBulkCampaignNegativeKeywordOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkCampaignNegativeKeywordOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkCampaignNegativeKeywordOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
