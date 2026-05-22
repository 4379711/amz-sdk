package services

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateReservationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReservationResponse{}

// CreateReservationResponse Response schema for the `createReservation` operation.
type CreateReservationResponse struct {
	Payload *CreateReservationRecord `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateReservationResponse instantiates a new CreateReservationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReservationResponse() *CreateReservationResponse {
	this := CreateReservationResponse{}
	return &this
}

// NewCreateReservationResponseWithDefaults instantiates a new CreateReservationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReservationResponseWithDefaults() *CreateReservationResponse {
	this := CreateReservationResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CreateReservationResponse) GetPayload() CreateReservationRecord {
	if o == nil || IsNil(o.Payload) {
		var ret CreateReservationRecord
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReservationResponse) GetPayloadOk() (*CreateReservationRecord, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CreateReservationResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given CreateReservationRecord and assigns it to the Payload field.
func (o *CreateReservationResponse) SetPayload(v CreateReservationRecord) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateReservationResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReservationResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateReservationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateReservationResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateReservationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateReservationResponse struct {
	value *CreateReservationResponse
	isSet bool
}

func (v NullableCreateReservationResponse) Get() *CreateReservationResponse {
	return v.value
}

func (v *NullableCreateReservationResponse) Set(val *CreateReservationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReservationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReservationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReservationResponse(val *CreateReservationResponse) *NullableCreateReservationResponse {
	return &NullableCreateReservationResponse{value: val, isSet: true}
}

func (v NullableCreateReservationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateReservationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
