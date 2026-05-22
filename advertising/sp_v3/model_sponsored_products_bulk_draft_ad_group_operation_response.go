package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkDraftAdGroupOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkDraftAdGroupOperationResponse{}

// SponsoredProductsBulkDraftAdGroupOperationResponse struct for SponsoredProductsBulkDraftAdGroupOperationResponse
type SponsoredProductsBulkDraftAdGroupOperationResponse struct {
	Success []SponsoredProductsDraftAdGroupSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsDraftAdGroupFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkDraftAdGroupOperationResponse instantiates a new SponsoredProductsBulkDraftAdGroupOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkDraftAdGroupOperationResponse() *SponsoredProductsBulkDraftAdGroupOperationResponse {
	this := SponsoredProductsBulkDraftAdGroupOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkDraftAdGroupOperationResponseWithDefaults instantiates a new SponsoredProductsBulkDraftAdGroupOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkDraftAdGroupOperationResponseWithDefaults() *SponsoredProductsBulkDraftAdGroupOperationResponse {
	this := SponsoredProductsBulkDraftAdGroupOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftAdGroupOperationResponse) GetSuccess() []SponsoredProductsDraftAdGroupSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsDraftAdGroupSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftAdGroupOperationResponse) GetSuccessOk() ([]SponsoredProductsDraftAdGroupSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftAdGroupOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsDraftAdGroupSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkDraftAdGroupOperationResponse) SetSuccess(v []SponsoredProductsDraftAdGroupSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftAdGroupOperationResponse) GetError() []SponsoredProductsDraftAdGroupFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsDraftAdGroupFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftAdGroupOperationResponse) GetErrorOk() ([]SponsoredProductsDraftAdGroupFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftAdGroupOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsDraftAdGroupFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkDraftAdGroupOperationResponse) SetError(v []SponsoredProductsDraftAdGroupFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkDraftAdGroupOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkDraftAdGroupOperationResponse struct {
	value *SponsoredProductsBulkDraftAdGroupOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkDraftAdGroupOperationResponse) Get() *SponsoredProductsBulkDraftAdGroupOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkDraftAdGroupOperationResponse) Set(val *SponsoredProductsBulkDraftAdGroupOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkDraftAdGroupOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkDraftAdGroupOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkDraftAdGroupOperationResponse(val *SponsoredProductsBulkDraftAdGroupOperationResponse) *NullableSponsoredProductsBulkDraftAdGroupOperationResponse {
	return &NullableSponsoredProductsBulkDraftAdGroupOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkDraftAdGroupOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkDraftAdGroupOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
