package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the EventTypeRuleDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventTypeRuleDuration{}

// EventTypeRuleDuration Object representing event type rule duration.
type EventTypeRuleDuration struct {
	// The event identifier. This value is available from the budget rules recommendation API.
	EventId string `json:"eventId"`
	// The event end date in YYYYMMDD format. Read-only.
	EndDate *string `json:"endDate,omitempty"`
	// The event name. Read-only.
	EventName *string `json:"eventName,omitempty"`
	// The event start date in YYYYMMDD format. Read-only. Note that this field is present only for announced events.
	StartDate *string `json:"startDate,omitempty"`
}

type _EventTypeRuleDuration EventTypeRuleDuration

// NewEventTypeRuleDuration instantiates a new EventTypeRuleDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypeRuleDuration(eventId string) *EventTypeRuleDuration {
	this := EventTypeRuleDuration{}
	this.EventId = eventId
	return &this
}

// NewEventTypeRuleDurationWithDefaults instantiates a new EventTypeRuleDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeRuleDurationWithDefaults() *EventTypeRuleDuration {
	this := EventTypeRuleDuration{}
	return &this
}

// GetEventId returns the EventId field value
func (o *EventTypeRuleDuration) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *EventTypeRuleDuration) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *EventTypeRuleDuration) SetEventId(v string) {
	o.EventId = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *EventTypeRuleDuration) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeRuleDuration) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *EventTypeRuleDuration) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *EventTypeRuleDuration) SetEndDate(v string) {
	o.EndDate = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *EventTypeRuleDuration) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeRuleDuration) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *EventTypeRuleDuration) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *EventTypeRuleDuration) SetEventName(v string) {
	o.EventName = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *EventTypeRuleDuration) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeRuleDuration) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *EventTypeRuleDuration) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *EventTypeRuleDuration) SetStartDate(v string) {
	o.StartDate = &v
}

func (o EventTypeRuleDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventId"] = o.EventId
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableEventTypeRuleDuration struct {
	value *EventTypeRuleDuration
	isSet bool
}

func (v NullableEventTypeRuleDuration) Get() *EventTypeRuleDuration {
	return v.value
}

func (v *NullableEventTypeRuleDuration) Set(val *EventTypeRuleDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeRuleDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeRuleDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeRuleDuration(val *EventTypeRuleDuration) *NullableEventTypeRuleDuration {
	return &NullableEventTypeRuleDuration{value: val, isSet: true}
}

func (v NullableEventTypeRuleDuration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEventTypeRuleDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
