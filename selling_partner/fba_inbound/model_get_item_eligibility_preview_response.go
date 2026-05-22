package fba_inbound

import (
	"github.com/bytedance/sonic"
)

// checks if the GetItemEligibilityPreviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetItemEligibilityPreviewResponse{}

// GetItemEligibilityPreviewResponse The response schema for the getItemEligibilityPreview operation.
type GetItemEligibilityPreviewResponse struct {
	Payload *ItemEligibilityPreview `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetItemEligibilityPreviewResponse instantiates a new GetItemEligibilityPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetItemEligibilityPreviewResponse() *GetItemEligibilityPreviewResponse {
	this := GetItemEligibilityPreviewResponse{}
	return &this
}

// NewGetItemEligibilityPreviewResponseWithDefaults instantiates a new GetItemEligibilityPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetItemEligibilityPreviewResponseWithDefaults() *GetItemEligibilityPreviewResponse {
	this := GetItemEligibilityPreviewResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetItemEligibilityPreviewResponse) GetPayload() ItemEligibilityPreview {
	if o == nil || IsNil(o.Payload) {
		var ret ItemEligibilityPreview
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetItemEligibilityPreviewResponse) GetPayloadOk() (*ItemEligibilityPreview, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetItemEligibilityPreviewResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ItemEligibilityPreview and assigns it to the Payload field.
func (o *GetItemEligibilityPreviewResponse) SetPayload(v ItemEligibilityPreview) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetItemEligibilityPreviewResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetItemEligibilityPreviewResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetItemEligibilityPreviewResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetItemEligibilityPreviewResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetItemEligibilityPreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetItemEligibilityPreviewResponse struct {
	value *GetItemEligibilityPreviewResponse
	isSet bool
}

func (v NullableGetItemEligibilityPreviewResponse) Get() *GetItemEligibilityPreviewResponse {
	return v.value
}

func (v *NullableGetItemEligibilityPreviewResponse) Set(val *GetItemEligibilityPreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetItemEligibilityPreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetItemEligibilityPreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetItemEligibilityPreviewResponse(val *GetItemEligibilityPreviewResponse) *NullableGetItemEligibilityPreviewResponse {
	return &NullableGetItemEligibilityPreviewResponse{value: val, isSet: true}
}

func (v NullableGetItemEligibilityPreviewResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetItemEligibilityPreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
