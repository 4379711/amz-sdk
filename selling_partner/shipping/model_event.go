package shipping

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the Event type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Event{}

// Event An event of a shipment
type Event struct {
	// The event code of a shipment, such as Departed, Received, and ReadyForReceive.
	EventCode string `json:"eventCode"`
	// The date and time of an event for a shipment.
	EventTime time.Time `json:"eventTime"`
	Location  *Location `json:"location,omitempty"`
}

type _Event Event

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(eventCode string, eventTime time.Time) *Event {
	this := Event{}
	this.EventCode = eventCode
	this.EventTime = eventTime
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetEventCode returns the EventCode field value
func (o *Event) GetEventCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventCode
}

// GetEventCodeOk returns a tuple with the EventCode field value
// and a boolean to check if the value has been set.
func (o *Event) GetEventCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCode, true
}

// SetEventCode sets field value
func (o *Event) SetEventCode(v string) {
	o.EventCode = v
}

// GetEventTime returns the EventTime field value
func (o *Event) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *Event) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *Event) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Event) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Event) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *Event) SetLocation(v Location) {
	o.Location = &v
}

func (o Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventCode"] = o.EventCode
	toSerialize["eventTime"] = o.EventTime
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	return toSerialize, nil
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
