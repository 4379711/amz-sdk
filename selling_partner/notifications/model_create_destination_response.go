package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateDestinationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDestinationResponse{}

// CreateDestinationResponse The response schema for the createDestination operation.
type CreateDestinationResponse struct {
	Payload *Destination `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateDestinationResponse instantiates a new CreateDestinationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDestinationResponse() *CreateDestinationResponse {
	this := CreateDestinationResponse{}
	return &this
}

// NewCreateDestinationResponseWithDefaults instantiates a new CreateDestinationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDestinationResponseWithDefaults() *CreateDestinationResponse {
	this := CreateDestinationResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CreateDestinationResponse) GetPayload() Destination {
	if o == nil || IsNil(o.Payload) {
		var ret Destination
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDestinationResponse) GetPayloadOk() (*Destination, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CreateDestinationResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given Destination and assigns it to the Payload field.
func (o *CreateDestinationResponse) SetPayload(v Destination) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateDestinationResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDestinationResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateDestinationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateDestinationResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateDestinationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateDestinationResponse struct {
	value *CreateDestinationResponse
	isSet bool
}

func (v NullableCreateDestinationResponse) Get() *CreateDestinationResponse {
	return v.value
}

func (v *NullableCreateDestinationResponse) Set(val *CreateDestinationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDestinationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDestinationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDestinationResponse(val *CreateDestinationResponse) *NullableCreateDestinationResponse {
	return &NullableCreateDestinationResponse{value: val, isSet: true}
}

func (v NullableCreateDestinationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateDestinationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
