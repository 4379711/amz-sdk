package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the CarrierAppointment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CarrierAppointment{}

// CarrierAppointment Contains details for a transportation carrier appointment. This appointment is vended out by Amazon and is an indicator for when a transportation carrier is accepting shipments to be picked up.
type CarrierAppointment struct {
	// The end timestamp of the appointment in UTC.
	EndTime time.Time `json:"endTime"`
	// The start timestamp of the appointment in UTC.
	StartTime time.Time `json:"startTime"`
}

type _CarrierAppointment CarrierAppointment

// NewCarrierAppointment instantiates a new CarrierAppointment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrierAppointment(endTime time.Time, startTime time.Time) *CarrierAppointment {
	this := CarrierAppointment{}
	this.EndTime = endTime
	this.StartTime = startTime
	return &this
}

// NewCarrierAppointmentWithDefaults instantiates a new CarrierAppointment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierAppointmentWithDefaults() *CarrierAppointment {
	this := CarrierAppointment{}
	return &this
}

// GetEndTime returns the EndTime field value
func (o *CarrierAppointment) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *CarrierAppointment) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *CarrierAppointment) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetStartTime returns the StartTime field value
func (o *CarrierAppointment) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *CarrierAppointment) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *CarrierAppointment) SetStartTime(v time.Time) {
	o.StartTime = v
}

func (o CarrierAppointment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endTime"] = o.EndTime
	toSerialize["startTime"] = o.StartTime
	return toSerialize, nil
}

type NullableCarrierAppointment struct {
	value *CarrierAppointment
	isSet bool
}

func (v NullableCarrierAppointment) Get() *CarrierAppointment {
	return v.value
}

func (v *NullableCarrierAppointment) Set(val *CarrierAppointment) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierAppointment) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierAppointment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierAppointment(val *CarrierAppointment) *NullableCarrierAppointment {
	return &NullableCarrierAppointment{value: val, isSet: true}
}

func (v NullableCarrierAppointment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCarrierAppointment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
