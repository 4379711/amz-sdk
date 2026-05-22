package services

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateReservationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateReservationRequest{}

// UpdateReservationRequest Request schema for the `updateReservation` operation.
type UpdateReservationRequest struct {
	// Resource (store) identifier.
	ResourceId  string      `json:"resourceId"`
	Reservation Reservation `json:"reservation"`
}

type _UpdateReservationRequest UpdateReservationRequest

// NewUpdateReservationRequest instantiates a new UpdateReservationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReservationRequest(resourceId string, reservation Reservation) *UpdateReservationRequest {
	this := UpdateReservationRequest{}
	this.ResourceId = resourceId
	this.Reservation = reservation
	return &this
}

// NewUpdateReservationRequestWithDefaults instantiates a new UpdateReservationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReservationRequestWithDefaults() *UpdateReservationRequest {
	this := UpdateReservationRequest{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *UpdateReservationRequest) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *UpdateReservationRequest) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *UpdateReservationRequest) SetResourceId(v string) {
	o.ResourceId = v
}

// GetReservation returns the Reservation field value
func (o *UpdateReservationRequest) GetReservation() Reservation {
	if o == nil {
		var ret Reservation
		return ret
	}

	return o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value
// and a boolean to check if the value has been set.
func (o *UpdateReservationRequest) GetReservationOk() (*Reservation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reservation, true
}

// SetReservation sets field value
func (o *UpdateReservationRequest) SetReservation(v Reservation) {
	o.Reservation = v
}

func (o UpdateReservationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["reservation"] = o.Reservation
	return toSerialize, nil
}

type NullableUpdateReservationRequest struct {
	value *UpdateReservationRequest
	isSet bool
}

func (v NullableUpdateReservationRequest) Get() *UpdateReservationRequest {
	return v.value
}

func (v *NullableUpdateReservationRequest) Set(val *UpdateReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReservationRequest(val *UpdateReservationRequest) *NullableUpdateReservationRequest {
	return &NullableUpdateReservationRequest{value: val, isSet: true}
}

func (v NullableUpdateReservationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
