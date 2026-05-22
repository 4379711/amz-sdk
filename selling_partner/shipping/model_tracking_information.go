package shipping

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the TrackingInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingInformation{}

// TrackingInformation The payload schema for the getTrackingInformation operation.
type TrackingInformation struct {
	// The tracking id generated to each shipment. It contains a series of letters or digits or both.
	TrackingId string          `json:"trackingId"`
	Summary    TrackingSummary `json:"summary"`
	// The promised delivery date and time of a shipment.
	PromisedDeliveryDate time.Time `json:"promisedDeliveryDate"`
	// A list of events of a shipment.
	EventHistory []Event `json:"eventHistory"`
}

type _TrackingInformation TrackingInformation

// NewTrackingInformation instantiates a new TrackingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingInformation(trackingId string, summary TrackingSummary, promisedDeliveryDate time.Time, eventHistory []Event) *TrackingInformation {
	this := TrackingInformation{}
	this.TrackingId = trackingId
	this.Summary = summary
	this.PromisedDeliveryDate = promisedDeliveryDate
	this.EventHistory = eventHistory
	return &this
}

// NewTrackingInformationWithDefaults instantiates a new TrackingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingInformationWithDefaults() *TrackingInformation {
	this := TrackingInformation{}
	return &this
}

// GetTrackingId returns the TrackingId field value
func (o *TrackingInformation) GetTrackingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value
// and a boolean to check if the value has been set.
func (o *TrackingInformation) GetTrackingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingId, true
}

// SetTrackingId sets field value
func (o *TrackingInformation) SetTrackingId(v string) {
	o.TrackingId = v
}

// GetSummary returns the Summary field value
func (o *TrackingInformation) GetSummary() TrackingSummary {
	if o == nil {
		var ret TrackingSummary
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *TrackingInformation) GetSummaryOk() (*TrackingSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *TrackingInformation) SetSummary(v TrackingSummary) {
	o.Summary = v
}

// GetPromisedDeliveryDate returns the PromisedDeliveryDate field value
func (o *TrackingInformation) GetPromisedDeliveryDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PromisedDeliveryDate
}

// GetPromisedDeliveryDateOk returns a tuple with the PromisedDeliveryDate field value
// and a boolean to check if the value has been set.
func (o *TrackingInformation) GetPromisedDeliveryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromisedDeliveryDate, true
}

// SetPromisedDeliveryDate sets field value
func (o *TrackingInformation) SetPromisedDeliveryDate(v time.Time) {
	o.PromisedDeliveryDate = v
}

// GetEventHistory returns the EventHistory field value
func (o *TrackingInformation) GetEventHistory() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.EventHistory
}

// GetEventHistoryOk returns a tuple with the EventHistory field value
// and a boolean to check if the value has been set.
func (o *TrackingInformation) GetEventHistoryOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventHistory, true
}

// SetEventHistory sets field value
func (o *TrackingInformation) SetEventHistory(v []Event) {
	o.EventHistory = v
}

func (o TrackingInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trackingId"] = o.TrackingId
	toSerialize["summary"] = o.Summary
	toSerialize["promisedDeliveryDate"] = o.PromisedDeliveryDate
	toSerialize["eventHistory"] = o.EventHistory
	return toSerialize, nil
}

type NullableTrackingInformation struct {
	value *TrackingInformation
	isSet bool
}

func (v NullableTrackingInformation) Get() *TrackingInformation {
	return v.value
}

func (v *NullableTrackingInformation) Set(val *TrackingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingInformation(val *TrackingInformation) *NullableTrackingInformation {
	return &NullableTrackingInformation{value: val, isSet: true}
}

func (v NullableTrackingInformation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTrackingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
