package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the RuleDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleDuration{}

// RuleDuration struct for RuleDuration
type RuleDuration struct {
	EventTypeRuleDuration     *EventTypeRuleDuration     `json:"eventTypeRuleDuration,omitempty"`
	DateRangeTypeRuleDuration *DateRangeTypeRuleDuration `json:"dateRangeTypeRuleDuration,omitempty"`
}

// NewRuleDuration instantiates a new RuleDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleDuration() *RuleDuration {
	this := RuleDuration{}
	return &this
}

// NewRuleDurationWithDefaults instantiates a new RuleDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleDurationWithDefaults() *RuleDuration {
	this := RuleDuration{}
	return &this
}

// GetEventTypeRuleDuration returns the EventTypeRuleDuration field value if set, zero value otherwise.
func (o *RuleDuration) GetEventTypeRuleDuration() EventTypeRuleDuration {
	if o == nil || IsNil(o.EventTypeRuleDuration) {
		var ret EventTypeRuleDuration
		return ret
	}
	return *o.EventTypeRuleDuration
}

// GetEventTypeRuleDurationOk returns a tuple with the EventTypeRuleDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDuration) GetEventTypeRuleDurationOk() (*EventTypeRuleDuration, bool) {
	if o == nil || IsNil(o.EventTypeRuleDuration) {
		return nil, false
	}
	return o.EventTypeRuleDuration, true
}

// HasEventTypeRuleDuration returns a boolean if a field has been set.
func (o *RuleDuration) HasEventTypeRuleDuration() bool {
	if o != nil && !IsNil(o.EventTypeRuleDuration) {
		return true
	}

	return false
}

// SetEventTypeRuleDuration gets a reference to the given EventTypeRuleDuration and assigns it to the EventTypeRuleDuration field.
func (o *RuleDuration) SetEventTypeRuleDuration(v EventTypeRuleDuration) {
	o.EventTypeRuleDuration = &v
}

// GetDateRangeTypeRuleDuration returns the DateRangeTypeRuleDuration field value if set, zero value otherwise.
func (o *RuleDuration) GetDateRangeTypeRuleDuration() DateRangeTypeRuleDuration {
	if o == nil || IsNil(o.DateRangeTypeRuleDuration) {
		var ret DateRangeTypeRuleDuration
		return ret
	}
	return *o.DateRangeTypeRuleDuration
}

// GetDateRangeTypeRuleDurationOk returns a tuple with the DateRangeTypeRuleDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDuration) GetDateRangeTypeRuleDurationOk() (*DateRangeTypeRuleDuration, bool) {
	if o == nil || IsNil(o.DateRangeTypeRuleDuration) {
		return nil, false
	}
	return o.DateRangeTypeRuleDuration, true
}

// HasDateRangeTypeRuleDuration returns a boolean if a field has been set.
func (o *RuleDuration) HasDateRangeTypeRuleDuration() bool {
	if o != nil && !IsNil(o.DateRangeTypeRuleDuration) {
		return true
	}

	return false
}

// SetDateRangeTypeRuleDuration gets a reference to the given DateRangeTypeRuleDuration and assigns it to the DateRangeTypeRuleDuration field.
func (o *RuleDuration) SetDateRangeTypeRuleDuration(v DateRangeTypeRuleDuration) {
	o.DateRangeTypeRuleDuration = &v
}

func (o RuleDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventTypeRuleDuration) {
		toSerialize["eventTypeRuleDuration"] = o.EventTypeRuleDuration
	}
	if !IsNil(o.DateRangeTypeRuleDuration) {
		toSerialize["dateRangeTypeRuleDuration"] = o.DateRangeTypeRuleDuration
	}
	return toSerialize, nil
}

type NullableRuleDuration struct {
	value *RuleDuration
	isSet bool
}

func (v NullableRuleDuration) Get() *RuleDuration {
	return v.value
}

func (v *NullableRuleDuration) Set(val *RuleDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleDuration(val *RuleDuration) *NullableRuleDuration {
	return &NullableRuleDuration{value: val, isSet: true}
}

func (v NullableRuleDuration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
