package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkGlobalKeywordOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkGlobalKeywordOperationResponse{}

// SponsoredProductsBulkGlobalKeywordOperationResponse struct for SponsoredProductsBulkGlobalKeywordOperationResponse
type SponsoredProductsBulkGlobalKeywordOperationResponse struct {
	Success []SponsoredProductsGlobalKeywordSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsGlobalKeywordFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkGlobalKeywordOperationResponse instantiates a new SponsoredProductsBulkGlobalKeywordOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkGlobalKeywordOperationResponse() *SponsoredProductsBulkGlobalKeywordOperationResponse {
	this := SponsoredProductsBulkGlobalKeywordOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkGlobalKeywordOperationResponseWithDefaults instantiates a new SponsoredProductsBulkGlobalKeywordOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkGlobalKeywordOperationResponseWithDefaults() *SponsoredProductsBulkGlobalKeywordOperationResponse {
	this := SponsoredProductsBulkGlobalKeywordOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalKeywordOperationResponse) GetSuccess() []SponsoredProductsGlobalKeywordSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsGlobalKeywordSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalKeywordOperationResponse) GetSuccessOk() ([]SponsoredProductsGlobalKeywordSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalKeywordOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsGlobalKeywordSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkGlobalKeywordOperationResponse) SetSuccess(v []SponsoredProductsGlobalKeywordSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalKeywordOperationResponse) GetError() []SponsoredProductsGlobalKeywordFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsGlobalKeywordFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalKeywordOperationResponse) GetErrorOk() ([]SponsoredProductsGlobalKeywordFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalKeywordOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsGlobalKeywordFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkGlobalKeywordOperationResponse) SetError(v []SponsoredProductsGlobalKeywordFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkGlobalKeywordOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkGlobalKeywordOperationResponse struct {
	value *SponsoredProductsBulkGlobalKeywordOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkGlobalKeywordOperationResponse) Get() *SponsoredProductsBulkGlobalKeywordOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkGlobalKeywordOperationResponse) Set(val *SponsoredProductsBulkGlobalKeywordOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkGlobalKeywordOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkGlobalKeywordOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkGlobalKeywordOperationResponse(val *SponsoredProductsBulkGlobalKeywordOperationResponse) *NullableSponsoredProductsBulkGlobalKeywordOperationResponse {
	return &NullableSponsoredProductsBulkGlobalKeywordOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkGlobalKeywordOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkGlobalKeywordOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
