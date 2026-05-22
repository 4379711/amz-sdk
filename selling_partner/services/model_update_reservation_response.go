package services

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateReservationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateReservationResponse{}

// UpdateReservationResponse Response schema for the `updateReservation` operation.
type UpdateReservationResponse struct {
	Payload *UpdateReservationRecord `json:"payload,omitempty"`
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewUpdateReservationResponse instantiates a new UpdateReservationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReservationResponse() *UpdateReservationResponse {
	this := UpdateReservationResponse{}
	return &this
}

// NewUpdateReservationResponseWithDefaults instantiates a new UpdateReservationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReservationResponseWithDefaults() *UpdateReservationResponse {
	this := UpdateReservationResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *UpdateReservationResponse) GetPayload() UpdateReservationRecord {
	if o == nil || IsNil(o.Payload) {
		var ret UpdateReservationRecord
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReservationResponse) GetPayloadOk() (*UpdateReservationRecord, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *UpdateReservationResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given UpdateReservationRecord and assigns it to the Payload field.
func (o *UpdateReservationResponse) SetPayload(v UpdateReservationRecord) {
	o.Payload = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *UpdateReservationResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReservationResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *UpdateReservationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *UpdateReservationResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o UpdateReservationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableUpdateReservationResponse struct {
	value *UpdateReservationResponse
	isSet bool
}

func (v NullableUpdateReservationResponse) Get() *UpdateReservationResponse {
	return v.value
}

func (v *NullableUpdateReservationResponse) Set(val *UpdateReservationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReservationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReservationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReservationResponse(val *UpdateReservationResponse) *NullableUpdateReservationResponse {
	return &NullableUpdateReservationResponse{value: val, isSet: true}
}

func (v NullableUpdateReservationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateReservationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
