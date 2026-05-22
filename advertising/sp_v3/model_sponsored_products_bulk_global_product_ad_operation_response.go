package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkGlobalProductAdOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkGlobalProductAdOperationResponse{}

// SponsoredProductsBulkGlobalProductAdOperationResponse struct for SponsoredProductsBulkGlobalProductAdOperationResponse
type SponsoredProductsBulkGlobalProductAdOperationResponse struct {
	Success []SponsoredProductsGlobalProductAdSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsGlobalProductAdFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkGlobalProductAdOperationResponse instantiates a new SponsoredProductsBulkGlobalProductAdOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkGlobalProductAdOperationResponse() *SponsoredProductsBulkGlobalProductAdOperationResponse {
	this := SponsoredProductsBulkGlobalProductAdOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkGlobalProductAdOperationResponseWithDefaults instantiates a new SponsoredProductsBulkGlobalProductAdOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkGlobalProductAdOperationResponseWithDefaults() *SponsoredProductsBulkGlobalProductAdOperationResponse {
	this := SponsoredProductsBulkGlobalProductAdOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalProductAdOperationResponse) GetSuccess() []SponsoredProductsGlobalProductAdSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsGlobalProductAdSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalProductAdOperationResponse) GetSuccessOk() ([]SponsoredProductsGlobalProductAdSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalProductAdOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsGlobalProductAdSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkGlobalProductAdOperationResponse) SetSuccess(v []SponsoredProductsGlobalProductAdSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalProductAdOperationResponse) GetError() []SponsoredProductsGlobalProductAdFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsGlobalProductAdFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalProductAdOperationResponse) GetErrorOk() ([]SponsoredProductsGlobalProductAdFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalProductAdOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsGlobalProductAdFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkGlobalProductAdOperationResponse) SetError(v []SponsoredProductsGlobalProductAdFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkGlobalProductAdOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkGlobalProductAdOperationResponse struct {
	value *SponsoredProductsBulkGlobalProductAdOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkGlobalProductAdOperationResponse) Get() *SponsoredProductsBulkGlobalProductAdOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkGlobalProductAdOperationResponse) Set(val *SponsoredProductsBulkGlobalProductAdOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkGlobalProductAdOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkGlobalProductAdOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkGlobalProductAdOperationResponse(val *SponsoredProductsBulkGlobalProductAdOperationResponse) *NullableSponsoredProductsBulkGlobalProductAdOperationResponse {
	return &NullableSponsoredProductsBulkGlobalProductAdOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkGlobalProductAdOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkGlobalProductAdOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
