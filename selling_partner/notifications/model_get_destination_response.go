package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the GetDestinationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDestinationResponse{}

// GetDestinationResponse The response schema for the `getDestination` operation.
type GetDestinationResponse struct {
	Payload *Destination `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewGetDestinationResponse instantiates a new GetDestinationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDestinationResponse() *GetDestinationResponse {
	this := GetDestinationResponse{}
	return &this
}

// NewGetDestinationResponseWithDefaults instantiates a new GetDestinationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDestinationResponseWithDefaults() *GetDestinationResponse {
	this := GetDestinationResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GetDestinationResponse) GetPayload() Destination {
	if o == nil || IsNil(o.Payload) {
		var ret Destination
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDestinationResponse) GetPayloadOk() (*Destination, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GetDestinationResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given Destination and assigns it to the Payload field.
func (o *GetDestinationResponse) SetPayload(v Destination) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetDestinationResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDestinationResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetDestinationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *GetDestinationResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetDestinationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetDestinationResponse struct {
	value *GetDestinationResponse
	isSet bool
}

func (v NullableGetDestinationResponse) Get() *GetDestinationResponse {
	return v.value
}

func (v *NullableGetDestinationResponse) Set(val *GetDestinationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDestinationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDestinationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDestinationResponse(val *GetDestinationResponse) *NullableGetDestinationResponse {
	return &NullableGetDestinationResponse{value: val, isSet: true}
}

func (v NullableGetDestinationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetDestinationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
