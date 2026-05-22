package services

import (
	"github.com/bytedance/sonic"
)

// checks if the Reservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reservation{}

// Reservation Reservation object reduces the capacity of a resource.
type Reservation struct {
	// Unique identifier for a reservation. If present, it is treated as an update reservation request and will update the corresponding reservation. Otherwise, it is treated as a new create reservation request.
	ReservationId *string `json:"reservationId,omitempty"`
	// Type of reservation.
	Type         string             `json:"type"`
	Availability AvailabilityRecord `json:"availability"`
}

type _Reservation Reservation

// NewReservation instantiates a new Reservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservation(type_ string, availability AvailabilityRecord) *Reservation {
	this := Reservation{}
	this.Type = type_
	this.Availability = availability
	return &this
}

// NewReservationWithDefaults instantiates a new Reservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationWithDefaults() *Reservation {
	this := Reservation{}
	return &this
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *Reservation) GetReservationId() string {
	if o == nil || IsNil(o.ReservationId) {
		var ret string
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetReservationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *Reservation) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given string and assigns it to the ReservationId field.
func (o *Reservation) SetReservationId(v string) {
	o.ReservationId = &v
}

// GetType returns the Type field value
func (o *Reservation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Reservation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Reservation) SetType(v string) {
	o.Type = v
}

// GetAvailability returns the Availability field value
func (o *Reservation) GetAvailability() AvailabilityRecord {
	if o == nil {
		var ret AvailabilityRecord
		return ret
	}

	return o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value
// and a boolean to check if the value has been set.
func (o *Reservation) GetAvailabilityOk() (*AvailabilityRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Availability, true
}

// SetAvailability sets field value
func (o *Reservation) SetAvailability(v AvailabilityRecord) {
	o.Availability = v
}

func (o Reservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	toSerialize["type"] = o.Type
	toSerialize["availability"] = o.Availability
	return toSerialize, nil
}

type NullableReservation struct {
	value *Reservation
	isSet bool
}

func (v NullableReservation) Get() *Reservation {
	return v.value
}

func (v *NullableReservation) Set(val *Reservation) {
	v.value = val
	v.isSet = true
}

func (v NullableReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservation(val *Reservation) *NullableReservation {
	return &NullableReservation{value: val, isSet: true}
}

func (v NullableReservation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
