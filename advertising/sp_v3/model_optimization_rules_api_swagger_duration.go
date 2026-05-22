package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDuration{}

// OptimizationRulesAPIDuration The duration of an optimization rule.
type OptimizationRulesAPIDuration struct {
	// Identifier for the event during which the rule is applied.
	EventId *string `json:"eventId,omitempty"`
	// Name of the event during which the rule is applied.
	EventName *string `json:"eventName,omitempty"`
	// Time of optimization rule creation in ISO 8061.
	StartTime string `json:"startTime"`
	// Time of optimization rule completion in ISO 8061.
	EndTime *string `json:"endTime,omitempty"`
}

type _OptimizationRulesAPIDuration OptimizationRulesAPIDuration

// NewOptimizationRulesAPIDuration instantiates a new OptimizationRulesAPIDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDuration(startTime string) *OptimizationRulesAPIDuration {
	this := OptimizationRulesAPIDuration{}
	this.StartTime = startTime
	return &this
}

// NewOptimizationRulesAPIDurationWithDefaults instantiates a new OptimizationRulesAPIDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDurationWithDefaults() *OptimizationRulesAPIDuration {
	this := OptimizationRulesAPIDuration{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *OptimizationRulesAPIDuration) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDuration) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *OptimizationRulesAPIDuration) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *OptimizationRulesAPIDuration) SetEventId(v string) {
	o.EventId = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *OptimizationRulesAPIDuration) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDuration) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *OptimizationRulesAPIDuration) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *OptimizationRulesAPIDuration) SetEventName(v string) {
	o.EventName = &v
}

// GetStartTime returns the StartTime field value
func (o *OptimizationRulesAPIDuration) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDuration) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *OptimizationRulesAPIDuration) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *OptimizationRulesAPIDuration) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDuration) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *OptimizationRulesAPIDuration) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *OptimizationRulesAPIDuration) SetEndTime(v string) {
	o.EndTime = &v
}

func (o OptimizationRulesAPIDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	toSerialize["startTime"] = o.StartTime
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDuration struct {
	value *OptimizationRulesAPIDuration
	isSet bool
}

func (v NullableOptimizationRulesAPIDuration) Get() *OptimizationRulesAPIDuration {
	return v.value
}

func (v *NullableOptimizationRulesAPIDuration) Set(val *OptimizationRulesAPIDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDuration(val *OptimizationRulesAPIDuration) *NullableOptimizationRulesAPIDuration {
	return &NullableOptimizationRulesAPIDuration{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDuration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
