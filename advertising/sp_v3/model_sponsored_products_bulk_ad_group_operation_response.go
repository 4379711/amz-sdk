package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkAdGroupOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkAdGroupOperationResponse{}

// SponsoredProductsBulkAdGroupOperationResponse struct for SponsoredProductsBulkAdGroupOperationResponse
type SponsoredProductsBulkAdGroupOperationResponse struct {
	Success []SponsoredProductsAdGroupSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsAdGroupFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkAdGroupOperationResponse instantiates a new SponsoredProductsBulkAdGroupOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkAdGroupOperationResponse() *SponsoredProductsBulkAdGroupOperationResponse {
	this := SponsoredProductsBulkAdGroupOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkAdGroupOperationResponseWithDefaults instantiates a new SponsoredProductsBulkAdGroupOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkAdGroupOperationResponseWithDefaults() *SponsoredProductsBulkAdGroupOperationResponse {
	this := SponsoredProductsBulkAdGroupOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkAdGroupOperationResponse) GetSuccess() []SponsoredProductsAdGroupSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsAdGroupSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkAdGroupOperationResponse) GetSuccessOk() ([]SponsoredProductsAdGroupSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkAdGroupOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsAdGroupSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkAdGroupOperationResponse) SetSuccess(v []SponsoredProductsAdGroupSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkAdGroupOperationResponse) GetError() []SponsoredProductsAdGroupFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsAdGroupFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkAdGroupOperationResponse) GetErrorOk() ([]SponsoredProductsAdGroupFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkAdGroupOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsAdGroupFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkAdGroupOperationResponse) SetError(v []SponsoredProductsAdGroupFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkAdGroupOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkAdGroupOperationResponse struct {
	value *SponsoredProductsBulkAdGroupOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkAdGroupOperationResponse) Get() *SponsoredProductsBulkAdGroupOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkAdGroupOperationResponse) Set(val *SponsoredProductsBulkAdGroupOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkAdGroupOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkAdGroupOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkAdGroupOperationResponse(val *SponsoredProductsBulkAdGroupOperationResponse) *NullableSponsoredProductsBulkAdGroupOperationResponse {
	return &NullableSponsoredProductsBulkAdGroupOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkAdGroupOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkAdGroupOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
