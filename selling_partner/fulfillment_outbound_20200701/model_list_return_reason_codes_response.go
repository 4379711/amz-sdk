package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the ListReturnReasonCodesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListReturnReasonCodesResponse{}

// ListReturnReasonCodesResponse The response schema for the `listReturnReasonCodes` operation.
type ListReturnReasonCodesResponse struct {
	Payload *ListReturnReasonCodesResult `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewListReturnReasonCodesResponse instantiates a new ListReturnReasonCodesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListReturnReasonCodesResponse() *ListReturnReasonCodesResponse {
	this := ListReturnReasonCodesResponse{}
	return &this
}

// NewListReturnReasonCodesResponseWithDefaults instantiates a new ListReturnReasonCodesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListReturnReasonCodesResponseWithDefaults() *ListReturnReasonCodesResponse {
	this := ListReturnReasonCodesResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ListReturnReasonCodesResponse) GetPayload() ListReturnReasonCodesResult {
	if o == nil || IsNil(o.Payload) {
		var ret ListReturnReasonCodesResult
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnReasonCodesResponse) GetPayloadOk() (*ListReturnReasonCodesResult, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ListReturnReasonCodesResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ListReturnReasonCodesResult and assigns it to the Payload field.
func (o *ListReturnReasonCodesResponse) SetPayload(v ListReturnReasonCodesResult) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ListReturnReasonCodesResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnReasonCodesResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ListReturnReasonCodesResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ListReturnReasonCodesResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o ListReturnReasonCodesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableListReturnReasonCodesResponse struct {
	value *ListReturnReasonCodesResponse
	isSet bool
}

func (v NullableListReturnReasonCodesResponse) Get() *ListReturnReasonCodesResponse {
	return v.value
}

func (v *NullableListReturnReasonCodesResponse) Set(val *ListReturnReasonCodesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListReturnReasonCodesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListReturnReasonCodesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListReturnReasonCodesResponse(val *ListReturnReasonCodesResponse) *NullableListReturnReasonCodesResponse {
	return &NullableListReturnReasonCodesResponse{value: val, isSet: true}
}

func (v NullableListReturnReasonCodesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListReturnReasonCodesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
