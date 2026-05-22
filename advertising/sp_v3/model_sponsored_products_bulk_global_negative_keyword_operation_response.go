package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkGlobalNegativeKeywordOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkGlobalNegativeKeywordOperationResponse{}

// SponsoredProductsBulkGlobalNegativeKeywordOperationResponse struct for SponsoredProductsBulkGlobalNegativeKeywordOperationResponse
type SponsoredProductsBulkGlobalNegativeKeywordOperationResponse struct {
	Success []SponsoredProductsGlobalNegativeKeywordSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsGlobalNegativeKeywordFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkGlobalNegativeKeywordOperationResponse instantiates a new SponsoredProductsBulkGlobalNegativeKeywordOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkGlobalNegativeKeywordOperationResponse() *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse {
	this := SponsoredProductsBulkGlobalNegativeKeywordOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkGlobalNegativeKeywordOperationResponseWithDefaults instantiates a new SponsoredProductsBulkGlobalNegativeKeywordOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkGlobalNegativeKeywordOperationResponseWithDefaults() *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse {
	this := SponsoredProductsBulkGlobalNegativeKeywordOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) GetSuccess() []SponsoredProductsGlobalNegativeKeywordSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsGlobalNegativeKeywordSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) GetSuccessOk() ([]SponsoredProductsGlobalNegativeKeywordSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsGlobalNegativeKeywordSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) SetSuccess(v []SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) GetError() []SponsoredProductsGlobalNegativeKeywordFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsGlobalNegativeKeywordFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) GetErrorOk() ([]SponsoredProductsGlobalNegativeKeywordFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsGlobalNegativeKeywordFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) SetError(v []SponsoredProductsGlobalNegativeKeywordFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkGlobalNegativeKeywordOperationResponse struct {
	value *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkGlobalNegativeKeywordOperationResponse) Get() *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkGlobalNegativeKeywordOperationResponse) Set(val *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkGlobalNegativeKeywordOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkGlobalNegativeKeywordOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkGlobalNegativeKeywordOperationResponse(val *SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) *NullableSponsoredProductsBulkGlobalNegativeKeywordOperationResponse {
	return &NullableSponsoredProductsBulkGlobalNegativeKeywordOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkGlobalNegativeKeywordOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkGlobalNegativeKeywordOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
