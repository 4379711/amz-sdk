package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse{}

// SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse struct for SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse
type SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse struct {
	Success []SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse instantiates a new SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse() *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse {
	this := SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponseWithDefaults instantiates a new SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponseWithDefaults() *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse {
	this := SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) GetSuccess() []SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) GetSuccessOk() ([]SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) SetSuccess(v []SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) GetError() []SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) GetErrorOk() ([]SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) SetError(v []SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse struct {
	value *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) Get() *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) Set(val *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse(val *SponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) *NullableSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse {
	return &NullableSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkDraftNegativeTargetingClauseOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
