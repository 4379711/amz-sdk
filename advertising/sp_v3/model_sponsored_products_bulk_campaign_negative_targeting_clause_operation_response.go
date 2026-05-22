package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse{}

// SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse struct for SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse
type SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse struct {
	Success []SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse instantiates a new SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse() *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse {
	this := SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponseWithDefaults instantiates a new SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponseWithDefaults() *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse {
	this := SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) GetSuccess() []SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) GetSuccessOk() ([]SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) SetSuccess(v []SponsoredProductsCampaignNegativeTargetingClauseSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) GetError() []SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) GetErrorOk() ([]SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) SetError(v []SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse struct {
	value *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) Get() *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) Set(val *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse(val *SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) *NullableSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse {
	return &NullableSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
