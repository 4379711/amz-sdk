package services

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateReservationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReservationRequest{}

// CreateReservationRequest Request schema for the `createReservation` operation.
type CreateReservationRequest struct {
	// Resource (store) identifier.
	ResourceId  string      `json:"resourceId"`
	Reservation Reservation `json:"reservation"`
}

type _CreateReservationRequest CreateReservationRequest

// NewCreateReservationRequest instantiates a new CreateReservationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReservationRequest(resourceId string, reservation Reservation) *CreateReservationRequest {
	this := CreateReservationRequest{}
	this.ResourceId = resourceId
	this.Reservation = reservation
	return &this
}

// NewCreateReservationRequestWithDefaults instantiates a new CreateReservationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReservationRequestWithDefaults() *CreateReservationRequest {
	this := CreateReservationRequest{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *CreateReservationRequest) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *CreateReservationRequest) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *CreateReservationRequest) SetResourceId(v string) {
	o.ResourceId = v
}

// GetReservation returns the Reservation field value
func (o *CreateReservationRequest) GetReservation() Reservation {
	if o == nil {
		var ret Reservation
		return ret
	}

	return o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value
// and a boolean to check if the value has been set.
func (o *CreateReservationRequest) GetReservationOk() (*Reservation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reservation, true
}

// SetReservation sets field value
func (o *CreateReservationRequest) SetReservation(v Reservation) {
	o.Reservation = v
}

func (o CreateReservationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["reservation"] = o.Reservation
	return toSerialize, nil
}

type NullableCreateReservationRequest struct {
	value *CreateReservationRequest
	isSet bool
}

func (v NullableCreateReservationRequest) Get() *CreateReservationRequest {
	return v.value
}

func (v *NullableCreateReservationRequest) Set(val *CreateReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReservationRequest(val *CreateReservationRequest) *NullableCreateReservationRequest {
	return &NullableCreateReservationRequest{value: val, isSet: true}
}

func (v NullableCreateReservationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
