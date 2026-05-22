package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateStreamSubscriptionRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStreamSubscriptionRequestContent{}

// CreateStreamSubscriptionRequestContent struct for CreateStreamSubscriptionRequestContent
type CreateStreamSubscriptionRequestContent struct {
	// Additional details associated with the subscription
	Notes *string `json:"notes,omitempty"`
	// Unique value supplied by the caller used to track identical API requests. Should request be re-tried, the caller should supply the same value. We recommend using GUID.
	ClientRequestToken string `json:"clientRequestToken"`
	// Identifier of data set, callers can be subscribed to. Please refer to https://advertising.amazon.com/API/docs/en-us/amazon-marketing-stream/data-guide for the list of all data sets.
	DataSetId string `json:"dataSetId"`
	// AWS ARN of the destination endpoint associated with the subscription. Supported destination types: - SQS
	DestinationArn *string      `json:"destinationArn,omitempty"`
	Destination    *Destination `json:"destination,omitempty"`
}

type _CreateStreamSubscriptionRequestContent CreateStreamSubscriptionRequestContent

// NewCreateStreamSubscriptionRequestContent instantiates a new CreateStreamSubscriptionRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStreamSubscriptionRequestContent(clientRequestToken string, dataSetId string) *CreateStreamSubscriptionRequestContent {
	this := CreateStreamSubscriptionRequestContent{}
	this.ClientRequestToken = clientRequestToken
	this.DataSetId = dataSetId
	return &this
}

// NewCreateStreamSubscriptionRequestContentWithDefaults instantiates a new CreateStreamSubscriptionRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStreamSubscriptionRequestContentWithDefaults() *CreateStreamSubscriptionRequestContent {
	this := CreateStreamSubscriptionRequestContent{}
	return &this
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CreateStreamSubscriptionRequestContent) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStreamSubscriptionRequestContent) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CreateStreamSubscriptionRequestContent) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CreateStreamSubscriptionRequestContent) SetNotes(v string) {
	o.Notes = &v
}

// GetClientRequestToken returns the ClientRequestToken field value
func (o *CreateStreamSubscriptionRequestContent) GetClientRequestToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientRequestToken
}

// GetClientRequestTokenOk returns a tuple with the ClientRequestToken field value
// and a boolean to check if the value has been set.
func (o *CreateStreamSubscriptionRequestContent) GetClientRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientRequestToken, true
}

// SetClientRequestToken sets field value
func (o *CreateStreamSubscriptionRequestContent) SetClientRequestToken(v string) {
	o.ClientRequestToken = v
}

// GetDataSetId returns the DataSetId field value
func (o *CreateStreamSubscriptionRequestContent) GetDataSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSetId
}

// GetDataSetIdOk returns a tuple with the DataSetId field value
// and a boolean to check if the value has been set.
func (o *CreateStreamSubscriptionRequestContent) GetDataSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSetId, true
}

// SetDataSetId sets field value
func (o *CreateStreamSubscriptionRequestContent) SetDataSetId(v string) {
	o.DataSetId = v
}

// GetDestinationArn returns the DestinationArn field value if set, zero value otherwise.
func (o *CreateStreamSubscriptionRequestContent) GetDestinationArn() string {
	if o == nil || IsNil(o.DestinationArn) {
		var ret string
		return ret
	}
	return *o.DestinationArn
}

// GetDestinationArnOk returns a tuple with the DestinationArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStreamSubscriptionRequestContent) GetDestinationArnOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationArn) {
		return nil, false
	}
	return o.DestinationArn, true
}

// HasDestinationArn returns a boolean if a field has been set.
func (o *CreateStreamSubscriptionRequestContent) HasDestinationArn() bool {
	if o != nil && !IsNil(o.DestinationArn) {
		return true
	}

	return false
}

// SetDestinationArn gets a reference to the given string and assigns it to the DestinationArn field.
func (o *CreateStreamSubscriptionRequestContent) SetDestinationArn(v string) {
	o.DestinationArn = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *CreateStreamSubscriptionRequestContent) GetDestination() Destination {
	if o == nil || IsNil(o.Destination) {
		var ret Destination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStreamSubscriptionRequestContent) GetDestinationOk() (*Destination, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *CreateStreamSubscriptionRequestContent) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given Destination and assigns it to the Destination field.
func (o *CreateStreamSubscriptionRequestContent) SetDestination(v Destination) {
	o.Destination = &v
}

func (o CreateStreamSubscriptionRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["clientRequestToken"] = o.ClientRequestToken
	toSerialize["dataSetId"] = o.DataSetId
	if !IsNil(o.DestinationArn) {
		toSerialize["destinationArn"] = o.DestinationArn
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	return toSerialize, nil
}

type NullableCreateStreamSubscriptionRequestContent struct {
	value *CreateStreamSubscriptionRequestContent
	isSet bool
}

func (v NullableCreateStreamSubscriptionRequestContent) Get() *CreateStreamSubscriptionRequestContent {
	return v.value
}

func (v *NullableCreateStreamSubscriptionRequestContent) Set(val *CreateStreamSubscriptionRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStreamSubscriptionRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStreamSubscriptionRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStreamSubscriptionRequestContent(val *CreateStreamSubscriptionRequestContent) *NullableCreateStreamSubscriptionRequestContent {
	return &NullableCreateStreamSubscriptionRequestContent{value: val, isSet: true}
}

func (v NullableCreateStreamSubscriptionRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateStreamSubscriptionRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
