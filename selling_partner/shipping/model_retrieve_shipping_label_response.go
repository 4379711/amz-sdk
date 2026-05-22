package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the RetrieveShippingLabelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetrieveShippingLabelResponse{}

// RetrieveShippingLabelResponse The response schema for the retrieveShippingLabel operation.
type RetrieveShippingLabelResponse struct {
	Payload *RetrieveShippingLabelResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewRetrieveShippingLabelResponse instantiates a new RetrieveShippingLabelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetrieveShippingLabelResponse() *RetrieveShippingLabelResponse {
	this := RetrieveShippingLabelResponse{}
	return &this
}

// NewRetrieveShippingLabelResponseWithDefaults instantiates a new RetrieveShippingLabelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetrieveShippingLabelResponseWithDefaults() *RetrieveShippingLabelResponse {
	this := RetrieveShippingLabelResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *RetrieveShippingLabelResponse) GetPayload() RetrieveShippingLabelResult {
	if o == nil || IsNil(o.Payload) {
		var ret RetrieveShippingLabelResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveShippingLabelResponse) GetPayloadOk() (*RetrieveShippingLabelResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *RetrieveShippingLabelResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given RetrieveShippingLabelResult and assigns it to the Payload field.
func (o *RetrieveShippingLabelResponse) SetPayload(v RetrieveShippingLabelResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *RetrieveShippingLabelResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveShippingLabelResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *RetrieveShippingLabelResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *RetrieveShippingLabelResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o RetrieveShippingLabelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableRetrieveShippingLabelResponse struct {
	value *RetrieveShippingLabelResponse
	isSet bool
}

func (v NullableRetrieveShippingLabelResponse) Get() *RetrieveShippingLabelResponse {
	return v.value
}

func (v *NullableRetrieveShippingLabelResponse) Set(val *RetrieveShippingLabelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveShippingLabelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveShippingLabelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveShippingLabelResponse(val *RetrieveShippingLabelResponse) *NullableRetrieveShippingLabelResponse {
	return &NullableRetrieveShippingLabelResponse{value: val, isSet: true}
}

func (v NullableRetrieveShippingLabelResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRetrieveShippingLabelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
