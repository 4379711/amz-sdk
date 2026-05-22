package ams

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the StreamSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamSubscription{}

// StreamSubscription struct for StreamSubscription
type StreamSubscription struct {
	// ISO8601 Timestamp
	CreatedDate time.Time `json:"createdDate"`
	// Additional details associated with the subscription
	Notes *string `json:"notes,omitempty"`
	// Identifier of data set, callers can be subscribed to. Please refer to https://advertising.amazon.com/API/docs/en-us/amazon-marketing-stream/data-guide for the list of all data sets.
	DataSetId string `json:"dataSetId"`
	// AWS ARN of the destination endpoint associated with the subscription. Supported destination types: - SQS
	DestinationArn *string      `json:"destinationArn,omitempty"`
	Destination    *Destination `json:"destination,omitempty"`
	// ISO8601 Timestamp
	UpdatedDate time.Time `json:"updatedDate"`
	// Unique subscription identifier
	SubscriptionId string                   `json:"subscriptionId"`
	Status         SubscriptionEntityStatus `json:"status"`
}

type _StreamSubscription StreamSubscription

// NewStreamSubscription instantiates a new StreamSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamSubscription(createdDate time.Time, dataSetId string, updatedDate time.Time, subscriptionId string, status SubscriptionEntityStatus) *StreamSubscription {
	this := StreamSubscription{}
	this.CreatedDate = createdDate
	this.DataSetId = dataSetId
	this.UpdatedDate = updatedDate
	this.SubscriptionId = subscriptionId
	this.Status = status
	return &this
}

// NewStreamSubscriptionWithDefaults instantiates a new StreamSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamSubscriptionWithDefaults() *StreamSubscription {
	this := StreamSubscription{}
	return &this
}

// GetCreatedDate returns the CreatedDate field value
func (o *StreamSubscription) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *StreamSubscription) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *StreamSubscription) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *StreamSubscription) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSubscription) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *StreamSubscription) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *StreamSubscription) SetNotes(v string) {
	o.Notes = &v
}

// GetDataSetId returns the DataSetId field value
func (o *StreamSubscription) GetDataSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSetId
}

// GetDataSetIdOk returns a tuple with the DataSetId field value
// and a boolean to check if the value has been set.
func (o *StreamSubscription) GetDataSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSetId, true
}

// SetDataSetId sets field value
func (o *StreamSubscription) SetDataSetId(v string) {
	o.DataSetId = v
}

// GetDestinationArn returns the DestinationArn field value if set, zero value otherwise.
func (o *StreamSubscription) GetDestinationArn() string {
	if o == nil || IsNil(o.DestinationArn) {
		var ret string
		return ret
	}
	return *o.DestinationArn
}

// GetDestinationArnOk returns a tuple with the DestinationArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSubscription) GetDestinationArnOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationArn) {
		return nil, false
	}
	return o.DestinationArn, true
}

// HasDestinationArn returns a boolean if a field has been set.
func (o *StreamSubscription) HasDestinationArn() bool {
	if o != nil && !IsNil(o.DestinationArn) {
		return true
	}

	return false
}

// SetDestinationArn gets a reference to the given string and assigns it to the DestinationArn field.
func (o *StreamSubscription) SetDestinationArn(v string) {
	o.DestinationArn = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *StreamSubscription) GetDestination() Destination {
	if o == nil || IsNil(o.Destination) {
		var ret Destination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSubscription) GetDestinationOk() (*Destination, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *StreamSubscription) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given Destination and assigns it to the Destination field.
func (o *StreamSubscription) SetDestination(v Destination) {
	o.Destination = &v
}

// GetUpdatedDate returns the UpdatedDate field value
func (o *StreamSubscription) GetUpdatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value
// and a boolean to check if the value has been set.
func (o *StreamSubscription) GetUpdatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedDate, true
}

// SetUpdatedDate sets field value
func (o *StreamSubscription) SetUpdatedDate(v time.Time) {
	o.UpdatedDate = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *StreamSubscription) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *StreamSubscription) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *StreamSubscription) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetStatus returns the Status field value
func (o *StreamSubscription) GetStatus() SubscriptionEntityStatus {
	if o == nil {
		var ret SubscriptionEntityStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StreamSubscription) GetStatusOk() (*SubscriptionEntityStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StreamSubscription) SetStatus(v SubscriptionEntityStatus) {
	o.Status = v
}

func (o StreamSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdDate"] = o.CreatedDate
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["dataSetId"] = o.DataSetId
	if !IsNil(o.DestinationArn) {
		toSerialize["destinationArn"] = o.DestinationArn
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	toSerialize["updatedDate"] = o.UpdatedDate
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableStreamSubscription struct {
	value *StreamSubscription
	isSet bool
}

func (v NullableStreamSubscription) Get() *StreamSubscription {
	return v.value
}

func (v *NullableStreamSubscription) Set(val *StreamSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamSubscription(val *StreamSubscription) *NullableStreamSubscription {
	return &NullableStreamSubscription{value: val, isSet: true}
}

func (v NullableStreamSubscription) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStreamSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
