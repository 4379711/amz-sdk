package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateMessagesSubscriptionRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMessagesSubscriptionRequestContent{}

// CreateMessagesSubscriptionRequestContent struct for CreateMessagesSubscriptionRequestContent
type CreateMessagesSubscriptionRequestContent struct {
	// Additional details which can be used to identity the  Destination. (Max Size of 128 characters)
	Notes string `json:"notes"`
	// Name of the dataset to create a  Subscription for.
	DataSetId string `json:"dataSetId"`
	// AWS ARN for the recipient of all data routed to the Consumer Destination.
	DestinationUri *string `json:"destinationUri,omitempty"`
	// Unique value supplied by the caller used to track identical API requests. Should request be re-tried, the caller should supply the same value. We recommend using GUID.
	ClientToken string       `json:"clientToken"`
	Destination *Destination `json:"destination,omitempty"`
}

type _CreateMessagesSubscriptionRequestContent CreateMessagesSubscriptionRequestContent

// NewCreateMessagesSubscriptionRequestContent instantiates a new CreateMessagesSubscriptionRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessagesSubscriptionRequestContent(notes string, dataSetId string, clientToken string) *CreateMessagesSubscriptionRequestContent {
	this := CreateMessagesSubscriptionRequestContent{}
	this.Notes = notes
	this.DataSetId = dataSetId
	this.ClientToken = clientToken
	return &this
}

// NewCreateMessagesSubscriptionRequestContentWithDefaults instantiates a new CreateMessagesSubscriptionRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessagesSubscriptionRequestContentWithDefaults() *CreateMessagesSubscriptionRequestContent {
	this := CreateMessagesSubscriptionRequestContent{}
	return &this
}

// GetNotes returns the Notes field value
func (o *CreateMessagesSubscriptionRequestContent) GetNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *CreateMessagesSubscriptionRequestContent) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notes, true
}

// SetNotes sets field value
func (o *CreateMessagesSubscriptionRequestContent) SetNotes(v string) {
	o.Notes = v
}

// GetDataSetId returns the DataSetId field value
func (o *CreateMessagesSubscriptionRequestContent) GetDataSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSetId
}

// GetDataSetIdOk returns a tuple with the DataSetId field value
// and a boolean to check if the value has been set.
func (o *CreateMessagesSubscriptionRequestContent) GetDataSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSetId, true
}

// SetDataSetId sets field value
func (o *CreateMessagesSubscriptionRequestContent) SetDataSetId(v string) {
	o.DataSetId = v
}

// GetDestinationUri returns the DestinationUri field value if set, zero value otherwise.
func (o *CreateMessagesSubscriptionRequestContent) GetDestinationUri() string {
	if o == nil || IsNil(o.DestinationUri) {
		var ret string
		return ret
	}
	return *o.DestinationUri
}

// GetDestinationUriOk returns a tuple with the DestinationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessagesSubscriptionRequestContent) GetDestinationUriOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationUri) {
		return nil, false
	}
	return o.DestinationUri, true
}

// HasDestinationUri returns a boolean if a field has been set.
func (o *CreateMessagesSubscriptionRequestContent) HasDestinationUri() bool {
	if o != nil && !IsNil(o.DestinationUri) {
		return true
	}

	return false
}

// SetDestinationUri gets a reference to the given string and assigns it to the DestinationUri field.
func (o *CreateMessagesSubscriptionRequestContent) SetDestinationUri(v string) {
	o.DestinationUri = &v
}

// GetClientToken returns the ClientToken field value
func (o *CreateMessagesSubscriptionRequestContent) GetClientToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientToken
}

// GetClientTokenOk returns a tuple with the ClientToken field value
// and a boolean to check if the value has been set.
func (o *CreateMessagesSubscriptionRequestContent) GetClientTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientToken, true
}

// SetClientToken sets field value
func (o *CreateMessagesSubscriptionRequestContent) SetClientToken(v string) {
	o.ClientToken = v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *CreateMessagesSubscriptionRequestContent) GetDestination() Destination {
	if o == nil || IsNil(o.Destination) {
		var ret Destination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessagesSubscriptionRequestContent) GetDestinationOk() (*Destination, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *CreateMessagesSubscriptionRequestContent) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given Destination and assigns it to the Destination field.
func (o *CreateMessagesSubscriptionRequestContent) SetDestination(v Destination) {
	o.Destination = &v
}

func (o CreateMessagesSubscriptionRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notes"] = o.Notes
	toSerialize["dataSetId"] = o.DataSetId
	if !IsNil(o.DestinationUri) {
		toSerialize["destinationUri"] = o.DestinationUri
	}
	toSerialize["clientToken"] = o.ClientToken
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	return toSerialize, nil
}

type NullableCreateMessagesSubscriptionRequestContent struct {
	value *CreateMessagesSubscriptionRequestContent
	isSet bool
}

func (v NullableCreateMessagesSubscriptionRequestContent) Get() *CreateMessagesSubscriptionRequestContent {
	return v.value
}

func (v *NullableCreateMessagesSubscriptionRequestContent) Set(val *CreateMessagesSubscriptionRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessagesSubscriptionRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessagesSubscriptionRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessagesSubscriptionRequestContent(val *CreateMessagesSubscriptionRequestContent) *NullableCreateMessagesSubscriptionRequestContent {
	return &NullableCreateMessagesSubscriptionRequestContent{value: val, isSet: true}
}

func (v NullableCreateMessagesSubscriptionRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateMessagesSubscriptionRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
