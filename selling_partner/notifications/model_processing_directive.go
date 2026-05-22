package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the ProcessingDirective type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessingDirective{}

// ProcessingDirective Additional information passed to the subscription to control the processing of notifications. For example, you can use an `eventFilter` to customize your subscription to send notifications for only the specified `marketplaceId`s, or select the aggregation time period at which to send notifications (for example: limit to one notification every five minutes for high frequency notifications). The specific features available vary depending on the `notificationType`.  This feature is currently only supported by the `ANY_OFFER_CHANGED` and `ORDER_CHANGE` `notificationType`s.
type ProcessingDirective struct {
	EventFilter *EventFilter `json:"eventFilter,omitempty"`
}

// NewProcessingDirective instantiates a new ProcessingDirective object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessingDirective() *ProcessingDirective {
	this := ProcessingDirective{}
	return &this
}

// NewProcessingDirectiveWithDefaults instantiates a new ProcessingDirective object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessingDirectiveWithDefaults() *ProcessingDirective {
	this := ProcessingDirective{}
	return &this
}

// GetEventFilter returns the EventFilter field value if set, zero value otherwise.
func (o *ProcessingDirective) GetEventFilter() EventFilter {
	if o == nil || IsNil(o.EventFilter) {
		var ret EventFilter
		return ret
	}
	return *o.EventFilter
}

// GetEventFilterOk returns a tuple with the EventFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessingDirective) GetEventFilterOk() (*EventFilter, bool) {
	if o == nil || IsNil(o.EventFilter) {
		return nil, false
	}
	return o.EventFilter, true
}

// HasEventFilter returns a boolean if a field has been set.
func (o *ProcessingDirective) HasEventFilter() bool {
	if o != nil && !IsNil(o.EventFilter) {
		return true
	}

	return false
}

// SetEventFilter gets a reference to the given EventFilter and assigns it to the EventFilter field.
func (o *ProcessingDirective) SetEventFilter(v EventFilter) {
	o.EventFilter = &v
}

func (o ProcessingDirective) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventFilter) {
		toSerialize["eventFilter"] = o.EventFilter
	}
	return toSerialize, nil
}

type NullableProcessingDirective struct {
	value *ProcessingDirective
	isSet bool
}

func (v NullableProcessingDirective) Get() *ProcessingDirective {
	return v.value
}

func (v *NullableProcessingDirective) Set(val *ProcessingDirective) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessingDirective) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessingDirective) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessingDirective(val *ProcessingDirective) *NullableProcessingDirective {
	return &NullableProcessingDirective{value: val, isSet: true}
}

func (v NullableProcessingDirective) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProcessingDirective) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
