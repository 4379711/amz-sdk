package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPGetAllRuleEventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPGetAllRuleEventResponse{}

// SPGetAllRuleEventResponse All Special individual and grouped events with date range.
type SPGetAllRuleEventResponse struct {
	// A list of grouped events with date range.
	GroupedEvents []SPGroupedEvent `json:"groupedEvents,omitempty"`
	// A list of individual events with date range.
	Events []SPIndividualEvent `json:"events,omitempty"`
}

// NewSPGetAllRuleEventResponse instantiates a new SPGetAllRuleEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPGetAllRuleEventResponse() *SPGetAllRuleEventResponse {
	this := SPGetAllRuleEventResponse{}
	return &this
}

// NewSPGetAllRuleEventResponseWithDefaults instantiates a new SPGetAllRuleEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPGetAllRuleEventResponseWithDefaults() *SPGetAllRuleEventResponse {
	this := SPGetAllRuleEventResponse{}
	return &this
}

// GetGroupedEvents returns the GroupedEvents field value if set, zero value otherwise.
func (o *SPGetAllRuleEventResponse) GetGroupedEvents() []SPGroupedEvent {
	if o == nil || IsNil(o.GroupedEvents) {
		var ret []SPGroupedEvent
		return ret
	}
	return o.GroupedEvents
}

// GetGroupedEventsOk returns a tuple with the GroupedEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPGetAllRuleEventResponse) GetGroupedEventsOk() ([]SPGroupedEvent, bool) {
	if o == nil || IsNil(o.GroupedEvents) {
		return nil, false
	}
	return o.GroupedEvents, true
}

// HasGroupedEvents returns a boolean if a field has been set.
func (o *SPGetAllRuleEventResponse) HasGroupedEvents() bool {
	if o != nil && !IsNil(o.GroupedEvents) {
		return true
	}

	return false
}

// SetGroupedEvents gets a reference to the given []SPGroupedEvent and assigns it to the GroupedEvents field.
func (o *SPGetAllRuleEventResponse) SetGroupedEvents(v []SPGroupedEvent) {
	o.GroupedEvents = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *SPGetAllRuleEventResponse) GetEvents() []SPIndividualEvent {
	if o == nil || IsNil(o.Events) {
		var ret []SPIndividualEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPGetAllRuleEventResponse) GetEventsOk() ([]SPIndividualEvent, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *SPGetAllRuleEventResponse) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []SPIndividualEvent and assigns it to the Events field.
func (o *SPGetAllRuleEventResponse) SetEvents(v []SPIndividualEvent) {
	o.Events = v
}

func (o SPGetAllRuleEventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupedEvents) {
		toSerialize["groupedEvents"] = o.GroupedEvents
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableSPGetAllRuleEventResponse struct {
	value *SPGetAllRuleEventResponse
	isSet bool
}

func (v NullableSPGetAllRuleEventResponse) Get() *SPGetAllRuleEventResponse {
	return v.value
}

func (v *NullableSPGetAllRuleEventResponse) Set(val *SPGetAllRuleEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSPGetAllRuleEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSPGetAllRuleEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPGetAllRuleEventResponse(val *SPGetAllRuleEventResponse) *NullableSPGetAllRuleEventResponse {
	return &NullableSPGetAllRuleEventResponse{value: val, isSet: true}
}

func (v NullableSPGetAllRuleEventResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPGetAllRuleEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
