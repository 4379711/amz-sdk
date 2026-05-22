package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkDraftNegativeKeywordOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkDraftNegativeKeywordOperationResponse{}

// SponsoredProductsBulkDraftNegativeKeywordOperationResponse struct for SponsoredProductsBulkDraftNegativeKeywordOperationResponse
type SponsoredProductsBulkDraftNegativeKeywordOperationResponse struct {
	Success []SponsoredProductsDraftNegativeKeywordSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsDraftNegativeKeywordFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkDraftNegativeKeywordOperationResponse instantiates a new SponsoredProductsBulkDraftNegativeKeywordOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkDraftNegativeKeywordOperationResponse() *SponsoredProductsBulkDraftNegativeKeywordOperationResponse {
	this := SponsoredProductsBulkDraftNegativeKeywordOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkDraftNegativeKeywordOperationResponseWithDefaults instantiates a new SponsoredProductsBulkDraftNegativeKeywordOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkDraftNegativeKeywordOperationResponseWithDefaults() *SponsoredProductsBulkDraftNegativeKeywordOperationResponse {
	this := SponsoredProductsBulkDraftNegativeKeywordOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftNegativeKeywordOperationResponse) GetSuccess() []SponsoredProductsDraftNegativeKeywordSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsDraftNegativeKeywordSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftNegativeKeywordOperationResponse) GetSuccessOk() ([]SponsoredProductsDraftNegativeKeywordSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftNegativeKeywordOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsDraftNegativeKeywordSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkDraftNegativeKeywordOperationResponse) SetSuccess(v []SponsoredProductsDraftNegativeKeywordSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftNegativeKeywordOperationResponse) GetError() []SponsoredProductsDraftNegativeKeywordFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsDraftNegativeKeywordFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftNegativeKeywordOperationResponse) GetErrorOk() ([]SponsoredProductsDraftNegativeKeywordFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftNegativeKeywordOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsDraftNegativeKeywordFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkDraftNegativeKeywordOperationResponse) SetError(v []SponsoredProductsDraftNegativeKeywordFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkDraftNegativeKeywordOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkDraftNegativeKeywordOperationResponse struct {
	value *SponsoredProductsBulkDraftNegativeKeywordOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkDraftNegativeKeywordOperationResponse) Get() *SponsoredProductsBulkDraftNegativeKeywordOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkDraftNegativeKeywordOperationResponse) Set(val *SponsoredProductsBulkDraftNegativeKeywordOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkDraftNegativeKeywordOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkDraftNegativeKeywordOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkDraftNegativeKeywordOperationResponse(val *SponsoredProductsBulkDraftNegativeKeywordOperationResponse) *NullableSponsoredProductsBulkDraftNegativeKeywordOperationResponse {
	return &NullableSponsoredProductsBulkDraftNegativeKeywordOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkDraftNegativeKeywordOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkDraftNegativeKeywordOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
