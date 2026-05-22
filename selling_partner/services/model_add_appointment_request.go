package services

import (
	"github.com/bytedance/sonic"
)

// checks if the AddAppointmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddAppointmentRequest{}

// AddAppointmentRequest Input for add appointment operation.
type AddAppointmentRequest struct {
	AppointmentTime AppointmentTimeInput `json:"appointmentTime"`
}

type _AddAppointmentRequest AddAppointmentRequest

// NewAddAppointmentRequest instantiates a new AddAppointmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAppointmentRequest(appointmentTime AppointmentTimeInput) *AddAppointmentRequest {
	this := AddAppointmentRequest{}
	this.AppointmentTime = appointmentTime
	return &this
}

// NewAddAppointmentRequestWithDefaults instantiates a new AddAppointmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAppointmentRequestWithDefaults() *AddAppointmentRequest {
	this := AddAppointmentRequest{}
	return &this
}

// GetAppointmentTime returns the AppointmentTime field value
func (o *AddAppointmentRequest) GetAppointmentTime() AppointmentTimeInput {
	if o == nil {
		var ret AppointmentTimeInput
		return ret
	}

	return o.AppointmentTime
}

// GetAppointmentTimeOk returns a tuple with the AppointmentTime field value
// and a boolean to check if the value has been set.
func (o *AddAppointmentRequest) GetAppointmentTimeOk() (*AppointmentTimeInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppointmentTime, true
}

// SetAppointmentTime sets field value
func (o *AddAppointmentRequest) SetAppointmentTime(v AppointmentTimeInput) {
	o.AppointmentTime = v
}

func (o AddAppointmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appointmentTime"] = o.AppointmentTime
	return toSerialize, nil
}

type NullableAddAppointmentRequest struct {
	value *AddAppointmentRequest
	isSet bool
}

func (v NullableAddAppointmentRequest) Get() *AddAppointmentRequest {
	return v.value
}

func (v *NullableAddAppointmentRequest) Set(val *AddAppointmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAppointmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAppointmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAppointmentRequest(val *AddAppointmentRequest) *NullableAddAppointmentRequest {
	return &NullableAddAppointmentRequest{value: val, isSet: true}
}

func (v NullableAddAppointmentRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAddAppointmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
