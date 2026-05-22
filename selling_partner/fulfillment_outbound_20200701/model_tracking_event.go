package fulfillment_outbound_20200701

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the TrackingEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingEvent{}

// TrackingEvent Information for tracking package deliveries.
type TrackingEvent struct {
	// Date timestamp
	EventDate    time.Time       `json:"eventDate"`
	EventAddress TrackingAddress `json:"eventAddress"`
	EventCode    EventCode       `json:"eventCode"`
	// A description for the corresponding event code.
	EventDescription string `json:"eventDescription"`
}

type _TrackingEvent TrackingEvent

// NewTrackingEvent instantiates a new TrackingEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingEvent(eventDate time.Time, eventAddress TrackingAddress, eventCode EventCode, eventDescription string) *TrackingEvent {
	this := TrackingEvent{}
	this.EventDate = eventDate
	this.EventAddress = eventAddress
	this.EventCode = eventCode
	this.EventDescription = eventDescription
	return &this
}

// NewTrackingEventWithDefaults instantiates a new TrackingEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingEventWithDefaults() *TrackingEvent {
	this := TrackingEvent{}
	return &this
}

// GetEventDate returns the EventDate field value
func (o *TrackingEvent) GetEventDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventDate
}

// GetEventDateOk returns a tuple with the EventDate field value
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetEventDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDate, true
}

// SetEventDate sets field value
func (o *TrackingEvent) SetEventDate(v time.Time) {
	o.EventDate = v
}

// GetEventAddress returns the EventAddress field value
func (o *TrackingEvent) GetEventAddress() TrackingAddress {
	if o == nil {
		var ret TrackingAddress
		return ret
	}

	return o.EventAddress
}

// GetEventAddressOk returns a tuple with the EventAddress field value
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetEventAddressOk() (*TrackingAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventAddress, true
}

// SetEventAddress sets field value
func (o *TrackingEvent) SetEventAddress(v TrackingAddress) {
	o.EventAddress = v
}

// GetEventCode returns the EventCode field value
func (o *TrackingEvent) GetEventCode() EventCode {
	if o == nil {
		var ret EventCode
		return ret
	}

	return o.EventCode
}

// GetEventCodeOk returns a tuple with the EventCode field value
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetEventCodeOk() (*EventCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCode, true
}

// SetEventCode sets field value
func (o *TrackingEvent) SetEventCode(v EventCode) {
	o.EventCode = v
}

// GetEventDescription returns the EventDescription field value
func (o *TrackingEvent) GetEventDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventDescription
}

// GetEventDescriptionOk returns a tuple with the EventDescription field value
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetEventDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDescription, true
}

// SetEventDescription sets field value
func (o *TrackingEvent) SetEventDescription(v string) {
	o.EventDescription = v
}

func (o TrackingEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventDate"] = o.EventDate
	toSerialize["eventAddress"] = o.EventAddress
	toSerialize["eventCode"] = o.EventCode
	toSerialize["eventDescription"] = o.EventDescription
	return toSerialize, nil
}

type NullableTrackingEvent struct {
	value *TrackingEvent
	isSet bool
}

func (v NullableTrackingEvent) Get() *TrackingEvent {
	return v.value
}

func (v *NullableTrackingEvent) Set(val *TrackingEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingEvent(val *TrackingEvent) *NullableTrackingEvent {
	return &NullableTrackingEvent{value: val, isSet: true}
}

func (v NullableTrackingEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTrackingEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
