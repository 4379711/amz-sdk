package sp_v3

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the SPGroupedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPGroupedEvent{}

// SPGroupedEvent struct for SPGroupedEvent
type SPGroupedEvent struct {
	// The grouped event identifier.
	GroupedEventId *string `json:"groupedEventId,omitempty"`
	// The grouped event name.
	GroupedEventName *string `json:"groupedEventName,omitempty"`
	// The end date in ISO-8601 format.
	EndDate *time.Time `json:"endDate,omitempty"`
	// The start date in ISO-8601 format.
	StartDate *time.Time `json:"startDate,omitempty"`
}

// NewSPGroupedEvent instantiates a new SPGroupedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPGroupedEvent() *SPGroupedEvent {
	this := SPGroupedEvent{}
	return &this
}

// NewSPGroupedEventWithDefaults instantiates a new SPGroupedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPGroupedEventWithDefaults() *SPGroupedEvent {
	this := SPGroupedEvent{}
	return &this
}

// GetGroupedEventId returns the GroupedEventId field value if set, zero value otherwise.
func (o *SPGroupedEvent) GetGroupedEventId() string {
	if o == nil || IsNil(o.GroupedEventId) {
		var ret string
		return ret
	}
	return *o.GroupedEventId
}

// GetGroupedEventIdOk returns a tuple with the GroupedEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPGroupedEvent) GetGroupedEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupedEventId) {
		return nil, false
	}
	return o.GroupedEventId, true
}

// HasGroupedEventId returns a boolean if a field has been set.
func (o *SPGroupedEvent) HasGroupedEventId() bool {
	if o != nil && !IsNil(o.GroupedEventId) {
		return true
	}

	return false
}

// SetGroupedEventId gets a reference to the given string and assigns it to the GroupedEventId field.
func (o *SPGroupedEvent) SetGroupedEventId(v string) {
	o.GroupedEventId = &v
}

// GetGroupedEventName returns the GroupedEventName field value if set, zero value otherwise.
func (o *SPGroupedEvent) GetGroupedEventName() string {
	if o == nil || IsNil(o.GroupedEventName) {
		var ret string
		return ret
	}
	return *o.GroupedEventName
}

// GetGroupedEventNameOk returns a tuple with the GroupedEventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPGroupedEvent) GetGroupedEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupedEventName) {
		return nil, false
	}
	return o.GroupedEventName, true
}

// HasGroupedEventName returns a boolean if a field has been set.
func (o *SPGroupedEvent) HasGroupedEventName() bool {
	if o != nil && !IsNil(o.GroupedEventName) {
		return true
	}

	return false
}

// SetGroupedEventName gets a reference to the given string and assigns it to the GroupedEventName field.
func (o *SPGroupedEvent) SetGroupedEventName(v string) {
	o.GroupedEventName = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SPGroupedEvent) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPGroupedEvent) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SPGroupedEvent) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *SPGroupedEvent) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SPGroupedEvent) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPGroupedEvent) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SPGroupedEvent) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *SPGroupedEvent) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o SPGroupedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupedEventId) {
		toSerialize["groupedEventId"] = o.GroupedEventId
	}
	if !IsNil(o.GroupedEventName) {
		toSerialize["groupedEventName"] = o.GroupedEventName
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableSPGroupedEvent struct {
	value *SPGroupedEvent
	isSet bool
}

func (v NullableSPGroupedEvent) Get() *SPGroupedEvent {
	return v.value
}

func (v *NullableSPGroupedEvent) Set(val *SPGroupedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSPGroupedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSPGroupedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPGroupedEvent(val *SPGroupedEvent) *NullableSPGroupedEvent {
	return &NullableSPGroupedEvent{value: val, isSet: true}
}

func (v NullableSPGroupedEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPGroupedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
