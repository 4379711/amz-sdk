package sp_v3

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the SPIndividualEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPIndividualEvent{}

// SPIndividualEvent struct for SPIndividualEvent
type SPIndividualEvent struct {
	// The event identifier.
	EventId *string `json:"eventId,omitempty"`
	// The end date in ISO-8601 format.
	EndDate *time.Time `json:"endDate,omitempty"`
	// The event name.
	EventName *string `json:"eventName,omitempty"`
	// The start date in ISO-8601 format.
	StartDate *time.Time `json:"startDate,omitempty"`
}

// NewSPIndividualEvent instantiates a new SPIndividualEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPIndividualEvent() *SPIndividualEvent {
	this := SPIndividualEvent{}
	return &this
}

// NewSPIndividualEventWithDefaults instantiates a new SPIndividualEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPIndividualEventWithDefaults() *SPIndividualEvent {
	this := SPIndividualEvent{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *SPIndividualEvent) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPIndividualEvent) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *SPIndividualEvent) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *SPIndividualEvent) SetEventId(v string) {
	o.EventId = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SPIndividualEvent) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPIndividualEvent) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SPIndividualEvent) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *SPIndividualEvent) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *SPIndividualEvent) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPIndividualEvent) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *SPIndividualEvent) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *SPIndividualEvent) SetEventName(v string) {
	o.EventName = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SPIndividualEvent) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPIndividualEvent) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SPIndividualEvent) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *SPIndividualEvent) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o SPIndividualEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
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

type NullableSPIndividualEvent struct {
	value *SPIndividualEvent
	isSet bool
}

func (v NullableSPIndividualEvent) Get() *SPIndividualEvent {
	return v.value
}

func (v *NullableSPIndividualEvent) Set(val *SPIndividualEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSPIndividualEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSPIndividualEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPIndividualEvent(val *SPIndividualEvent) *NullableSPIndividualEvent {
	return &NullableSPIndividualEvent{value: val, isSet: true}
}

func (v NullableSPIndividualEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPIndividualEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
