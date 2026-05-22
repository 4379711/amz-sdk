package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the AppointmentSlot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppointmentSlot{}

// AppointmentSlot The fulfillment center appointment slot for the transportation option.
type AppointmentSlot struct {
	// An identifier to a self-ship appointment slot.
	SlotId   string              `json:"slotId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	SlotTime AppointmentSlotTime `json:"slotTime"`
}

type _AppointmentSlot AppointmentSlot

// NewAppointmentSlot instantiates a new AppointmentSlot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentSlot(slotId string, slotTime AppointmentSlotTime) *AppointmentSlot {
	this := AppointmentSlot{}
	this.SlotId = slotId
	this.SlotTime = slotTime
	return &this
}

// NewAppointmentSlotWithDefaults instantiates a new AppointmentSlot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentSlotWithDefaults() *AppointmentSlot {
	this := AppointmentSlot{}
	return &this
}

// GetSlotId returns the SlotId field value
func (o *AppointmentSlot) GetSlotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value
// and a boolean to check if the value has been set.
func (o *AppointmentSlot) GetSlotIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlotId, true
}

// SetSlotId sets field value
func (o *AppointmentSlot) SetSlotId(v string) {
	o.SlotId = v
}

// GetSlotTime returns the SlotTime field value
func (o *AppointmentSlot) GetSlotTime() AppointmentSlotTime {
	if o == nil {
		var ret AppointmentSlotTime
		return ret
	}

	return o.SlotTime
}

// GetSlotTimeOk returns a tuple with the SlotTime field value
// and a boolean to check if the value has been set.
func (o *AppointmentSlot) GetSlotTimeOk() (*AppointmentSlotTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlotTime, true
}

// SetSlotTime sets field value
func (o *AppointmentSlot) SetSlotTime(v AppointmentSlotTime) {
	o.SlotTime = v
}

func (o AppointmentSlot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["slotId"] = o.SlotId
	toSerialize["slotTime"] = o.SlotTime
	return toSerialize, nil
}

type NullableAppointmentSlot struct {
	value *AppointmentSlot
	isSet bool
}

func (v NullableAppointmentSlot) Get() *AppointmentSlot {
	return v.value
}

func (v *NullableAppointmentSlot) Set(val *AppointmentSlot) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentSlot) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentSlot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentSlot(val *AppointmentSlot) *NullableAppointmentSlot {
	return &NullableAppointmentSlot{value: val, isSet: true}
}

func (v NullableAppointmentSlot) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAppointmentSlot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
