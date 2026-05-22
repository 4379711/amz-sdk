package services

import (
	"github.com/bytedance/sonic"
)

// checks if the RescheduleAppointmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RescheduleAppointmentRequest{}

// RescheduleAppointmentRequest Input for rescheduled appointment operation.
type RescheduleAppointmentRequest struct {
	AppointmentTime AppointmentTimeInput `json:"appointmentTime"`
	// The appointment reschedule reason code.
	RescheduleReasonCode string `json:"rescheduleReasonCode"`
}

type _RescheduleAppointmentRequest RescheduleAppointmentRequest

// NewRescheduleAppointmentRequest instantiates a new RescheduleAppointmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRescheduleAppointmentRequest(appointmentTime AppointmentTimeInput, rescheduleReasonCode string) *RescheduleAppointmentRequest {
	this := RescheduleAppointmentRequest{}
	this.AppointmentTime = appointmentTime
	this.RescheduleReasonCode = rescheduleReasonCode
	return &this
}

// NewRescheduleAppointmentRequestWithDefaults instantiates a new RescheduleAppointmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRescheduleAppointmentRequestWithDefaults() *RescheduleAppointmentRequest {
	this := RescheduleAppointmentRequest{}
	return &this
}

// GetAppointmentTime returns the AppointmentTime field value
func (o *RescheduleAppointmentRequest) GetAppointmentTime() AppointmentTimeInput {
	if o == nil {
		var ret AppointmentTimeInput
		return ret
	}

	return o.AppointmentTime
}

// GetAppointmentTimeOk returns a tuple with the AppointmentTime field value
// and a boolean to check if the value has been set.
func (o *RescheduleAppointmentRequest) GetAppointmentTimeOk() (*AppointmentTimeInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppointmentTime, true
}

// SetAppointmentTime sets field value
func (o *RescheduleAppointmentRequest) SetAppointmentTime(v AppointmentTimeInput) {
	o.AppointmentTime = v
}

// GetRescheduleReasonCode returns the RescheduleReasonCode field value
func (o *RescheduleAppointmentRequest) GetRescheduleReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RescheduleReasonCode
}

// GetRescheduleReasonCodeOk returns a tuple with the RescheduleReasonCode field value
// and a boolean to check if the value has been set.
func (o *RescheduleAppointmentRequest) GetRescheduleReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RescheduleReasonCode, true
}

// SetRescheduleReasonCode sets field value
func (o *RescheduleAppointmentRequest) SetRescheduleReasonCode(v string) {
	o.RescheduleReasonCode = v
}

func (o RescheduleAppointmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appointmentTime"] = o.AppointmentTime
	toSerialize["rescheduleReasonCode"] = o.RescheduleReasonCode
	return toSerialize, nil
}

type NullableRescheduleAppointmentRequest struct {
	value *RescheduleAppointmentRequest
	isSet bool
}

func (v NullableRescheduleAppointmentRequest) Get() *RescheduleAppointmentRequest {
	return v.value
}

func (v *NullableRescheduleAppointmentRequest) Set(val *RescheduleAppointmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRescheduleAppointmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRescheduleAppointmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRescheduleAppointmentRequest(val *RescheduleAppointmentRequest) *NullableRescheduleAppointmentRequest {
	return &NullableRescheduleAppointmentRequest{value: val, isSet: true}
}

func (v NullableRescheduleAppointmentRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRescheduleAppointmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
