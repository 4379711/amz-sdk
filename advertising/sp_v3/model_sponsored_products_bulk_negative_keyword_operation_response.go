package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkNegativeKeywordOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkNegativeKeywordOperationResponse{}

// SponsoredProductsBulkNegativeKeywordOperationResponse struct for SponsoredProductsBulkNegativeKeywordOperationResponse
type SponsoredProductsBulkNegativeKeywordOperationResponse struct {
	Success []SponsoredProductsNegativeKeywordSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsNegativeKeywordFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkNegativeKeywordOperationResponse instantiates a new SponsoredProductsBulkNegativeKeywordOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkNegativeKeywordOperationResponse() *SponsoredProductsBulkNegativeKeywordOperationResponse {
	this := SponsoredProductsBulkNegativeKeywordOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkNegativeKeywordOperationResponseWithDefaults instantiates a new SponsoredProductsBulkNegativeKeywordOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkNegativeKeywordOperationResponseWithDefaults() *SponsoredProductsBulkNegativeKeywordOperationResponse {
	this := SponsoredProductsBulkNegativeKeywordOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkNegativeKeywordOperationResponse) GetSuccess() []SponsoredProductsNegativeKeywordSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsNegativeKeywordSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkNegativeKeywordOperationResponse) GetSuccessOk() ([]SponsoredProductsNegativeKeywordSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkNegativeKeywordOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsNegativeKeywordSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkNegativeKeywordOperationResponse) SetSuccess(v []SponsoredProductsNegativeKeywordSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkNegativeKeywordOperationResponse) GetError() []SponsoredProductsNegativeKeywordFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsNegativeKeywordFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkNegativeKeywordOperationResponse) GetErrorOk() ([]SponsoredProductsNegativeKeywordFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkNegativeKeywordOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsNegativeKeywordFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkNegativeKeywordOperationResponse) SetError(v []SponsoredProductsNegativeKeywordFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkNegativeKeywordOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkNegativeKeywordOperationResponse struct {
	value *SponsoredProductsBulkNegativeKeywordOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkNegativeKeywordOperationResponse) Get() *SponsoredProductsBulkNegativeKeywordOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkNegativeKeywordOperationResponse) Set(val *SponsoredProductsBulkNegativeKeywordOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkNegativeKeywordOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkNegativeKeywordOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkNegativeKeywordOperationResponse(val *SponsoredProductsBulkNegativeKeywordOperationResponse) *NullableSponsoredProductsBulkNegativeKeywordOperationResponse {
	return &NullableSponsoredProductsBulkNegativeKeywordOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkNegativeKeywordOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkNegativeKeywordOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
