package shipping_v2

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the GetTrackingResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTrackingResult{}

// GetTrackingResult The payload for the getTracking operation.
type GetTrackingResult struct {
	// The carrier generated identifier for a package in a purchased shipment.
	TrackingId string `json:"trackingId"`
	// The carrier generated reverse identifier for a returned package in a purchased shipment.
	AlternateLegTrackingId string `json:"alternateLegTrackingId"`
	// A list of tracking events.
	EventHistory []Event `json:"eventHistory"`
	// The date and time by which the shipment is promised to be delivered.
	PromisedDeliveryDate time.Time       `json:"promisedDeliveryDate"`
	Summary              TrackingSummary `json:"summary"`
}

type _GetTrackingResult GetTrackingResult

// NewGetTrackingResult instantiates a new GetTrackingResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTrackingResult(trackingId string, alternateLegTrackingId string, eventHistory []Event, promisedDeliveryDate time.Time, summary TrackingSummary) *GetTrackingResult {
	this := GetTrackingResult{}
	this.TrackingId = trackingId
	this.AlternateLegTrackingId = alternateLegTrackingId
	this.EventHistory = eventHistory
	this.PromisedDeliveryDate = promisedDeliveryDate
	this.Summary = summary
	return &this
}

// NewGetTrackingResultWithDefaults instantiates a new GetTrackingResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTrackingResultWithDefaults() *GetTrackingResult {
	this := GetTrackingResult{}
	return &this
}

// GetTrackingId returns the TrackingId field value
func (o *GetTrackingResult) GetTrackingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value
// and a boolean to check if the value has been set.
func (o *GetTrackingResult) GetTrackingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingId, true
}

// SetTrackingId sets field value
func (o *GetTrackingResult) SetTrackingId(v string) {
	o.TrackingId = v
}

// GetAlternateLegTrackingId returns the AlternateLegTrackingId field value
func (o *GetTrackingResult) GetAlternateLegTrackingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateLegTrackingId
}

// GetAlternateLegTrackingIdOk returns a tuple with the AlternateLegTrackingId field value
// and a boolean to check if the value has been set.
func (o *GetTrackingResult) GetAlternateLegTrackingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateLegTrackingId, true
}

// SetAlternateLegTrackingId sets field value
func (o *GetTrackingResult) SetAlternateLegTrackingId(v string) {
	o.AlternateLegTrackingId = v
}

// GetEventHistory returns the EventHistory field value
func (o *GetTrackingResult) GetEventHistory() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.EventHistory
}

// GetEventHistoryOk returns a tuple with the EventHistory field value
// and a boolean to check if the value has been set.
func (o *GetTrackingResult) GetEventHistoryOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventHistory, true
}

// SetEventHistory sets field value
func (o *GetTrackingResult) SetEventHistory(v []Event) {
	o.EventHistory = v
}

// GetPromisedDeliveryDate returns the PromisedDeliveryDate field value
func (o *GetTrackingResult) GetPromisedDeliveryDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PromisedDeliveryDate
}

// GetPromisedDeliveryDateOk returns a tuple with the PromisedDeliveryDate field value
// and a boolean to check if the value has been set.
func (o *GetTrackingResult) GetPromisedDeliveryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromisedDeliveryDate, true
}

// SetPromisedDeliveryDate sets field value
func (o *GetTrackingResult) SetPromisedDeliveryDate(v time.Time) {
	o.PromisedDeliveryDate = v
}

// GetSummary returns the Summary field value
func (o *GetTrackingResult) GetSummary() TrackingSummary {
	if o == nil {
		var ret TrackingSummary
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *GetTrackingResult) GetSummaryOk() (*TrackingSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *GetTrackingResult) SetSummary(v TrackingSummary) {
	o.Summary = v
}

func (o GetTrackingResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trackingId"] = o.TrackingId
	toSerialize["alternateLegTrackingId"] = o.AlternateLegTrackingId
	toSerialize["eventHistory"] = o.EventHistory
	toSerialize["promisedDeliveryDate"] = o.PromisedDeliveryDate
	toSerialize["summary"] = o.Summary
	return toSerialize, nil
}

type NullableGetTrackingResult struct {
	value *GetTrackingResult
	isSet bool
}

func (v NullableGetTrackingResult) Get() *GetTrackingResult {
	return v.value
}

func (v *NullableGetTrackingResult) Set(val *GetTrackingResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTrackingResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTrackingResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTrackingResult(val *GetTrackingResult) *NullableGetTrackingResult {
	return &NullableGetTrackingResult{value: val, isSet: true}
}

func (v NullableGetTrackingResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetTrackingResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
