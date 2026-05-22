package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkDraftKeywordOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkDraftKeywordOperationResponse{}

// SponsoredProductsBulkDraftKeywordOperationResponse struct for SponsoredProductsBulkDraftKeywordOperationResponse
type SponsoredProductsBulkDraftKeywordOperationResponse struct {
	Success []SponsoredProductsDraftKeywordSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsDraftKeywordFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkDraftKeywordOperationResponse instantiates a new SponsoredProductsBulkDraftKeywordOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkDraftKeywordOperationResponse() *SponsoredProductsBulkDraftKeywordOperationResponse {
	this := SponsoredProductsBulkDraftKeywordOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkDraftKeywordOperationResponseWithDefaults instantiates a new SponsoredProductsBulkDraftKeywordOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkDraftKeywordOperationResponseWithDefaults() *SponsoredProductsBulkDraftKeywordOperationResponse {
	this := SponsoredProductsBulkDraftKeywordOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftKeywordOperationResponse) GetSuccess() []SponsoredProductsDraftKeywordSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsDraftKeywordSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftKeywordOperationResponse) GetSuccessOk() ([]SponsoredProductsDraftKeywordSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftKeywordOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsDraftKeywordSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkDraftKeywordOperationResponse) SetSuccess(v []SponsoredProductsDraftKeywordSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftKeywordOperationResponse) GetError() []SponsoredProductsDraftKeywordFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsDraftKeywordFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftKeywordOperationResponse) GetErrorOk() ([]SponsoredProductsDraftKeywordFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftKeywordOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsDraftKeywordFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkDraftKeywordOperationResponse) SetError(v []SponsoredProductsDraftKeywordFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkDraftKeywordOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkDraftKeywordOperationResponse struct {
	value *SponsoredProductsBulkDraftKeywordOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkDraftKeywordOperationResponse) Get() *SponsoredProductsBulkDraftKeywordOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkDraftKeywordOperationResponse) Set(val *SponsoredProductsBulkDraftKeywordOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkDraftKeywordOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkDraftKeywordOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkDraftKeywordOperationResponse(val *SponsoredProductsBulkDraftKeywordOperationResponse) *NullableSponsoredProductsBulkDraftKeywordOperationResponse {
	return &NullableSponsoredProductsBulkDraftKeywordOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkDraftKeywordOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkDraftKeywordOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
