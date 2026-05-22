package services

import (
	"github.com/bytedance/sonic"
)

// checks if the CancelReservationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelReservationResponse{}

// CancelReservationResponse Response schema for the `cancelReservation` operation.
type CancelReservationResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCancelReservationResponse instantiates a new CancelReservationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelReservationResponse() *CancelReservationResponse {
	this := CancelReservationResponse{}
	return &this
}

// NewCancelReservationResponseWithDefaults instantiates a new CancelReservationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelReservationResponseWithDefaults() *CancelReservationResponse {
	this := CancelReservationResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CancelReservationResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelReservationResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CancelReservationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CancelReservationResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CancelReservationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCancelReservationResponse struct {
	value *CancelReservationResponse
	isSet bool
}

func (v NullableCancelReservationResponse) Get() *CancelReservationResponse {
	return v.value
}

func (v *NullableCancelReservationResponse) Set(val *CancelReservationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelReservationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelReservationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelReservationResponse(val *CancelReservationResponse) *NullableCancelReservationResponse {
	return &NullableCancelReservationResponse{value: val, isSet: true}
}

func (v NullableCancelReservationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCancelReservationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
