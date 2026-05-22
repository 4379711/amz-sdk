package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkGlobalAdGroupOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkGlobalAdGroupOperationResponse{}

// SponsoredProductsBulkGlobalAdGroupOperationResponse struct for SponsoredProductsBulkGlobalAdGroupOperationResponse
type SponsoredProductsBulkGlobalAdGroupOperationResponse struct {
	Success []SponsoredProductsGlobalAdGroupSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsGlobalAdGroupFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkGlobalAdGroupOperationResponse instantiates a new SponsoredProductsBulkGlobalAdGroupOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkGlobalAdGroupOperationResponse() *SponsoredProductsBulkGlobalAdGroupOperationResponse {
	this := SponsoredProductsBulkGlobalAdGroupOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkGlobalAdGroupOperationResponseWithDefaults instantiates a new SponsoredProductsBulkGlobalAdGroupOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkGlobalAdGroupOperationResponseWithDefaults() *SponsoredProductsBulkGlobalAdGroupOperationResponse {
	this := SponsoredProductsBulkGlobalAdGroupOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalAdGroupOperationResponse) GetSuccess() []SponsoredProductsGlobalAdGroupSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsGlobalAdGroupSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalAdGroupOperationResponse) GetSuccessOk() ([]SponsoredProductsGlobalAdGroupSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalAdGroupOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsGlobalAdGroupSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkGlobalAdGroupOperationResponse) SetSuccess(v []SponsoredProductsGlobalAdGroupSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalAdGroupOperationResponse) GetError() []SponsoredProductsGlobalAdGroupFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsGlobalAdGroupFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalAdGroupOperationResponse) GetErrorOk() ([]SponsoredProductsGlobalAdGroupFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalAdGroupOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsGlobalAdGroupFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkGlobalAdGroupOperationResponse) SetError(v []SponsoredProductsGlobalAdGroupFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkGlobalAdGroupOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkGlobalAdGroupOperationResponse struct {
	value *SponsoredProductsBulkGlobalAdGroupOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkGlobalAdGroupOperationResponse) Get() *SponsoredProductsBulkGlobalAdGroupOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkGlobalAdGroupOperationResponse) Set(val *SponsoredProductsBulkGlobalAdGroupOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkGlobalAdGroupOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkGlobalAdGroupOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkGlobalAdGroupOperationResponse(val *SponsoredProductsBulkGlobalAdGroupOperationResponse) *NullableSponsoredProductsBulkGlobalAdGroupOperationResponse {
	return &NullableSponsoredProductsBulkGlobalAdGroupOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkGlobalAdGroupOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkGlobalAdGroupOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
