package shipping_v2

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the Event type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Event{}

// Event A tracking event.
type Event struct {
	EventCode EventCode `json:"eventCode"`
	Location  *Location `json:"location,omitempty"`
	// The ISO 8601 formatted timestamp of the event.
	EventTime    time.Time     `json:"eventTime"`
	ShipmentType *ShipmentType `json:"shipmentType,omitempty"`
}

type _Event Event

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(eventCode EventCode, eventTime time.Time) *Event {
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
func (o *Event) GetEventCode() EventCode {
	if o == nil {
		var ret EventCode
		return ret
	}

	return o.EventCode
}

// GetEventCodeOk returns a tuple with the EventCode field value
// and a boolean to check if the value has been set.
func (o *Event) GetEventCodeOk() (*EventCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCode, true
}

// SetEventCode sets field value
func (o *Event) SetEventCode(v EventCode) {
	o.EventCode = v
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

// GetShipmentType returns the ShipmentType field value if set, zero value otherwise.
func (o *Event) GetShipmentType() ShipmentType {
	if o == nil || IsNil(o.ShipmentType) {
		var ret ShipmentType
		return ret
	}
	return *o.ShipmentType
}

// GetShipmentTypeOk returns a tuple with the ShipmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetShipmentTypeOk() (*ShipmentType, bool) {
	if o == nil || IsNil(o.ShipmentType) {
		return nil, false
	}
	return o.ShipmentType, true
}

// HasShipmentType returns a boolean if a field has been set.
func (o *Event) HasShipmentType() bool {
	if o != nil && !IsNil(o.ShipmentType) {
		return true
	}

	return false
}

// SetShipmentType gets a reference to the given ShipmentType and assigns it to the ShipmentType field.
func (o *Event) SetShipmentType(v ShipmentType) {
	o.ShipmentType = &v
}

func (o Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventCode"] = o.EventCode
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	toSerialize["eventTime"] = o.EventTime
	if !IsNil(o.ShipmentType) {
		toSerialize["shipmentType"] = o.ShipmentType
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
