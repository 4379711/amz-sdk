package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the SelfShipAppointmentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfShipAppointmentDetails{}

// SelfShipAppointmentDetails Appointment details for carrier pickup or fulfillment center appointments.
type SelfShipAppointmentDetails struct {
	// Identifier for appointment.
	AppointmentId       *float32             `json:"appointmentId,omitempty"`
	AppointmentSlotTime *AppointmentSlotTime `json:"appointmentSlotTime,omitempty"`
	// Status of the appointment.
	AppointmentStatus *string `json:"appointmentStatus,omitempty"`
}

// NewSelfShipAppointmentDetails instantiates a new SelfShipAppointmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfShipAppointmentDetails() *SelfShipAppointmentDetails {
	this := SelfShipAppointmentDetails{}
	return &this
}

// NewSelfShipAppointmentDetailsWithDefaults instantiates a new SelfShipAppointmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfShipAppointmentDetailsWithDefaults() *SelfShipAppointmentDetails {
	this := SelfShipAppointmentDetails{}
	return &this
}

// GetAppointmentId returns the AppointmentId field value if set, zero value otherwise.
func (o *SelfShipAppointmentDetails) GetAppointmentId() float32 {
	if o == nil || IsNil(o.AppointmentId) {
		var ret float32
		return ret
	}
	return *o.AppointmentId
}

// GetAppointmentIdOk returns a tuple with the AppointmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfShipAppointmentDetails) GetAppointmentIdOk() (*float32, bool) {
	if o == nil || IsNil(o.AppointmentId) {
		return nil, false
	}
	return o.AppointmentId, true
}

// HasAppointmentId returns a boolean if a field has been set.
func (o *SelfShipAppointmentDetails) HasAppointmentId() bool {
	if o != nil && !IsNil(o.AppointmentId) {
		return true
	}

	return false
}

// SetAppointmentId gets a reference to the given float32 and assigns it to the AppointmentId field.
func (o *SelfShipAppointmentDetails) SetAppointmentId(v float32) {
	o.AppointmentId = &v
}

// GetAppointmentSlotTime returns the AppointmentSlotTime field value if set, zero value otherwise.
func (o *SelfShipAppointmentDetails) GetAppointmentSlotTime() AppointmentSlotTime {
	if o == nil || IsNil(o.AppointmentSlotTime) {
		var ret AppointmentSlotTime
		return ret
	}
	return *o.AppointmentSlotTime
}

// GetAppointmentSlotTimeOk returns a tuple with the AppointmentSlotTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfShipAppointmentDetails) GetAppointmentSlotTimeOk() (*AppointmentSlotTime, bool) {
	if o == nil || IsNil(o.AppointmentSlotTime) {
		return nil, false
	}
	return o.AppointmentSlotTime, true
}

// HasAppointmentSlotTime returns a boolean if a field has been set.
func (o *SelfShipAppointmentDetails) HasAppointmentSlotTime() bool {
	if o != nil && !IsNil(o.AppointmentSlotTime) {
		return true
	}

	return false
}

// SetAppointmentSlotTime gets a reference to the given AppointmentSlotTime and assigns it to the AppointmentSlotTime field.
func (o *SelfShipAppointmentDetails) SetAppointmentSlotTime(v AppointmentSlotTime) {
	o.AppointmentSlotTime = &v
}

// GetAppointmentStatus returns the AppointmentStatus field value if set, zero value otherwise.
func (o *SelfShipAppointmentDetails) GetAppointmentStatus() string {
	if o == nil || IsNil(o.AppointmentStatus) {
		var ret string
		return ret
	}
	return *o.AppointmentStatus
}

// GetAppointmentStatusOk returns a tuple with the AppointmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfShipAppointmentDetails) GetAppointmentStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AppointmentStatus) {
		return nil, false
	}
	return o.AppointmentStatus, true
}

// HasAppointmentStatus returns a boolean if a field has been set.
func (o *SelfShipAppointmentDetails) HasAppointmentStatus() bool {
	if o != nil && !IsNil(o.AppointmentStatus) {
		return true
	}

	return false
}

// SetAppointmentStatus gets a reference to the given string and assigns it to the AppointmentStatus field.
func (o *SelfShipAppointmentDetails) SetAppointmentStatus(v string) {
	o.AppointmentStatus = &v
}

func (o SelfShipAppointmentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppointmentId) {
		toSerialize["appointmentId"] = o.AppointmentId
	}
	if !IsNil(o.AppointmentSlotTime) {
		toSerialize["appointmentSlotTime"] = o.AppointmentSlotTime
	}
	if !IsNil(o.AppointmentStatus) {
		toSerialize["appointmentStatus"] = o.AppointmentStatus
	}
	return toSerialize, nil
}

type NullableSelfShipAppointmentDetails struct {
	value *SelfShipAppointmentDetails
	isSet bool
}

func (v NullableSelfShipAppointmentDetails) Get() *SelfShipAppointmentDetails {
	return v.value
}

func (v *NullableSelfShipAppointmentDetails) Set(val *SelfShipAppointmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfShipAppointmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfShipAppointmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfShipAppointmentDetails(val *SelfShipAppointmentDetails) *NullableSelfShipAppointmentDetails {
	return &NullableSelfShipAppointmentDetails{value: val, isSet: true}
}

func (v NullableSelfShipAppointmentDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSelfShipAppointmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
