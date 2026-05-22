package ams

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the MessagesSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessagesSubscription{}

// MessagesSubscription struct for MessagesSubscription
type MessagesSubscription struct {
	// ISO8601 Timestamp
	CreatedDate time.Time `json:"createdDate"`
	// the IAM policy which is required on the destination in order for Amazon being able to send data to it
	RequiredIamPolicy string  `json:"requiredIamPolicy"`
	Notes             string  `json:"notes"`
	DestinationUri    *string `json:"destinationUri,omitempty"`
	// Identifier of data set, callers can be subscribed to. Please refer to https://advertising.amazon.com/API/docs/en-us/amazon-marketing-stream/data-guide for the list of all data sets.
	DataSetId   string       `json:"dataSetId"`
	Destination *Destination `json:"destination,omitempty"`
	// Generic resource identifier which will work in most of the cases
	MessagesSubscriptionId string `json:"messagesSubscriptionId"`
	// ISO8601 Timestamp
	UpdatedDate   time.Time                `json:"updatedDate"`
	Version       float32                  `json:"version"`
	StatusMessage *string                  `json:"statusMessage,omitempty"`
	Status        SubscriptionEntityStatus `json:"status"`
}

type _MessagesSubscription MessagesSubscription

// NewMessagesSubscription instantiates a new MessagesSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessagesSubscription(createdDate time.Time, requiredIamPolicy string, notes string, dataSetId string, messagesSubscriptionId string, updatedDate time.Time, version float32, status SubscriptionEntityStatus) *MessagesSubscription {
	this := MessagesSubscription{}
	this.CreatedDate = createdDate
	this.RequiredIamPolicy = requiredIamPolicy
	this.Notes = notes
	this.DataSetId = dataSetId
	this.MessagesSubscriptionId = messagesSubscriptionId
	this.UpdatedDate = updatedDate
	this.Version = version
	this.Status = status
	return &this
}

// NewMessagesSubscriptionWithDefaults instantiates a new MessagesSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagesSubscriptionWithDefaults() *MessagesSubscription {
	this := MessagesSubscription{}
	return &this
}

// GetCreatedDate returns the CreatedDate field value
func (o *MessagesSubscription) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *MessagesSubscription) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *MessagesSubscription) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetRequiredIamPolicy returns the RequiredIamPolicy field value
func (o *MessagesSubscription) GetRequiredIamPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequiredIamPolicy
}

// GetRequiredIamPolicyOk returns a tuple with the RequiredIamPolicy field value
// and a boolean to check if the value has been set.
func (o *MessagesSubscription) GetRequiredIamPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiredIamPolicy, true
}

// SetRequiredIamPolicy sets field value
func (o *MessagesSubscription) SetRequiredIamPolicy(v string) {
	o.RequiredIamPolicy = v
}

// GetNotes returns the Notes field value
func (o *MessagesSubscription) GetNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *MessagesSubscription) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notes, true
}

// SetNotes sets field value
func (o *MessagesSubscription) SetNotes(v string) {
	o.Notes = v
}

// GetDestinationUri returns the DestinationUri field value if set, zero value otherwise.
func (o *MessagesSubscription) GetDestinationUri() string {
	if o == nil || IsNil(o.DestinationUri) {
		var ret string
		return ret
	}
	return *o.DestinationUri
}

// GetDestinationUriOk returns a tuple with the DestinationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesSubscription) GetDestinationUriOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationUri) {
		return nil, false
	}
	return o.DestinationUri, true
}

// HasDestinationUri returns a boolean if a field has been set.
func (o *MessagesSubscription) HasDestinationUri() bool {
	if o != nil && !IsNil(o.DestinationUri) {
		return true
	}

	return false
}

// SetDestinationUri gets a reference to the given string and assigns it to the DestinationUri field.
func (o *MessagesSubscription) SetDestinationUri(v string) {
	o.DestinationUri = &v
}

// GetDataSetId returns the DataSetId field value
func (o *MessagesSubscription) GetDataSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSetId
}

// GetDataSetIdOk returns a tuple with the DataSetId field value
// and a boolean to check if the value has been set.
func (o *MessagesSubscription) GetDataSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSetId, true
}

// SetDataSetId sets field value
func (o *MessagesSubscription) SetDataSetId(v string) {
	o.DataSetId = v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *MessagesSubscription) GetDestination() Destination {
	if o == nil || IsNil(o.Destination) {
		var ret Destination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesSubscription) GetDestinationOk() (*Destination, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *MessagesSubscription) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given Destination and assigns it to the Destination field.
func (o *MessagesSubscription) SetDestination(v Destination) {
	o.Destination = &v
}

// GetMessagesSubscriptionId returns the MessagesSubscriptionId field value
func (o *MessagesSubscription) GetMessagesSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessagesSubscriptionId
}

// GetMessagesSubscriptionIdOk returns a tuple with the MessagesSubscriptionId field value
// and a boolean to check if the value has been set.
func (o *MessagesSubscription) GetMessagesSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessagesSubscriptionId, true
}

// SetMessagesSubscriptionId sets field value
func (o *MessagesSubscription) SetMessagesSubscriptionId(v string) {
	o.MessagesSubscriptionId = v
}

// GetUpdatedDate returns the UpdatedDate field value
func (o *MessagesSubscription) GetUpdatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value
// and a boolean to check if the value has been set.
func (o *MessagesSubscription) GetUpdatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedDate, true
}

// SetUpdatedDate sets field value
func (o *MessagesSubscription) SetUpdatedDate(v time.Time) {
	o.UpdatedDate = v
}

// GetVersion returns the Version field value
func (o *MessagesSubscription) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *MessagesSubscription) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *MessagesSubscription) SetVersion(v float32) {
	o.Version = v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *MessagesSubscription) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesSubscription) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *MessagesSubscription) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *MessagesSubscription) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetStatus returns the Status field value
func (o *MessagesSubscription) GetStatus() SubscriptionEntityStatus {
	if o == nil {
		var ret SubscriptionEntityStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MessagesSubscription) GetStatusOk() (*SubscriptionEntityStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MessagesSubscription) SetStatus(v SubscriptionEntityStatus) {
	o.Status = v
}

func (o MessagesSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdDate"] = o.CreatedDate
	toSerialize["requiredIamPolicy"] = o.RequiredIamPolicy
	toSerialize["notes"] = o.Notes
	if !IsNil(o.DestinationUri) {
		toSerialize["destinationUri"] = o.DestinationUri
	}
	toSerialize["dataSetId"] = o.DataSetId
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	toSerialize["messagesSubscriptionId"] = o.MessagesSubscriptionId
	toSerialize["updatedDate"] = o.UpdatedDate
	toSerialize["version"] = o.Version
	if !IsNil(o.StatusMessage) {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableMessagesSubscription struct {
	value *MessagesSubscription
	isSet bool
}

func (v NullableMessagesSubscription) Get() *MessagesSubscription {
	return v.value
}

func (v *NullableMessagesSubscription) Set(val *MessagesSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableMessagesSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableMessagesSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessagesSubscription(val *MessagesSubscription) *NullableMessagesSubscription {
	return &NullableMessagesSubscription{value: val, isSet: true}
}

func (v NullableMessagesSubscription) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMessagesSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
