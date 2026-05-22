package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkGlobalTargetingClauseOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkGlobalTargetingClauseOperationResponse{}

// SponsoredProductsBulkGlobalTargetingClauseOperationResponse struct for SponsoredProductsBulkGlobalTargetingClauseOperationResponse
type SponsoredProductsBulkGlobalTargetingClauseOperationResponse struct {
	Success []SponsoredProductsGlobalTargetingClauseSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsGlobalTargetingClauseFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkGlobalTargetingClauseOperationResponse instantiates a new SponsoredProductsBulkGlobalTargetingClauseOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkGlobalTargetingClauseOperationResponse() *SponsoredProductsBulkGlobalTargetingClauseOperationResponse {
	this := SponsoredProductsBulkGlobalTargetingClauseOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkGlobalTargetingClauseOperationResponseWithDefaults instantiates a new SponsoredProductsBulkGlobalTargetingClauseOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkGlobalTargetingClauseOperationResponseWithDefaults() *SponsoredProductsBulkGlobalTargetingClauseOperationResponse {
	this := SponsoredProductsBulkGlobalTargetingClauseOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalTargetingClauseOperationResponse) GetSuccess() []SponsoredProductsGlobalTargetingClauseSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsGlobalTargetingClauseSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalTargetingClauseOperationResponse) GetSuccessOk() ([]SponsoredProductsGlobalTargetingClauseSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalTargetingClauseOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsGlobalTargetingClauseSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkGlobalTargetingClauseOperationResponse) SetSuccess(v []SponsoredProductsGlobalTargetingClauseSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalTargetingClauseOperationResponse) GetError() []SponsoredProductsGlobalTargetingClauseFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsGlobalTargetingClauseFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalTargetingClauseOperationResponse) GetErrorOk() ([]SponsoredProductsGlobalTargetingClauseFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalTargetingClauseOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsGlobalTargetingClauseFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkGlobalTargetingClauseOperationResponse) SetError(v []SponsoredProductsGlobalTargetingClauseFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkGlobalTargetingClauseOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkGlobalTargetingClauseOperationResponse struct {
	value *SponsoredProductsBulkGlobalTargetingClauseOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkGlobalTargetingClauseOperationResponse) Get() *SponsoredProductsBulkGlobalTargetingClauseOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkGlobalTargetingClauseOperationResponse) Set(val *SponsoredProductsBulkGlobalTargetingClauseOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkGlobalTargetingClauseOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkGlobalTargetingClauseOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkGlobalTargetingClauseOperationResponse(val *SponsoredProductsBulkGlobalTargetingClauseOperationResponse) *NullableSponsoredProductsBulkGlobalTargetingClauseOperationResponse {
	return &NullableSponsoredProductsBulkGlobalTargetingClauseOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkGlobalTargetingClauseOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkGlobalTargetingClauseOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
