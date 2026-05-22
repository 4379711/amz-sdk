package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateDspStreamSubscriptionRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDspStreamSubscriptionRequestContent{}

// CreateDspStreamSubscriptionRequestContent struct for CreateDspStreamSubscriptionRequestContent
type CreateDspStreamSubscriptionRequestContent struct {
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

type _CreateDspStreamSubscriptionRequestContent CreateDspStreamSubscriptionRequestContent

// NewCreateDspStreamSubscriptionRequestContent instantiates a new CreateDspStreamSubscriptionRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDspStreamSubscriptionRequestContent(clientRequestToken string, dataSetId string) *CreateDspStreamSubscriptionRequestContent {
	this := CreateDspStreamSubscriptionRequestContent{}
	this.ClientRequestToken = clientRequestToken
	this.DataSetId = dataSetId
	return &this
}

// NewCreateDspStreamSubscriptionRequestContentWithDefaults instantiates a new CreateDspStreamSubscriptionRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDspStreamSubscriptionRequestContentWithDefaults() *CreateDspStreamSubscriptionRequestContent {
	this := CreateDspStreamSubscriptionRequestContent{}
	return &this
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CreateDspStreamSubscriptionRequestContent) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDspStreamSubscriptionRequestContent) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CreateDspStreamSubscriptionRequestContent) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CreateDspStreamSubscriptionRequestContent) SetNotes(v string) {
	o.Notes = &v
}

// GetClientRequestToken returns the ClientRequestToken field value
func (o *CreateDspStreamSubscriptionRequestContent) GetClientRequestToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientRequestToken
}

// GetClientRequestTokenOk returns a tuple with the ClientRequestToken field value
// and a boolean to check if the value has been set.
func (o *CreateDspStreamSubscriptionRequestContent) GetClientRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientRequestToken, true
}

// SetClientRequestToken sets field value
func (o *CreateDspStreamSubscriptionRequestContent) SetClientRequestToken(v string) {
	o.ClientRequestToken = v
}

// GetDataSetId returns the DataSetId field value
func (o *CreateDspStreamSubscriptionRequestContent) GetDataSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSetId
}

// GetDataSetIdOk returns a tuple with the DataSetId field value
// and a boolean to check if the value has been set.
func (o *CreateDspStreamSubscriptionRequestContent) GetDataSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSetId, true
}

// SetDataSetId sets field value
func (o *CreateDspStreamSubscriptionRequestContent) SetDataSetId(v string) {
	o.DataSetId = v
}

// GetDestinationArn returns the DestinationArn field value if set, zero value otherwise.
func (o *CreateDspStreamSubscriptionRequestContent) GetDestinationArn() string {
	if o == nil || IsNil(o.DestinationArn) {
		var ret string
		return ret
	}
	return *o.DestinationArn
}

// GetDestinationArnOk returns a tuple with the DestinationArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDspStreamSubscriptionRequestContent) GetDestinationArnOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationArn) {
		return nil, false
	}
	return o.DestinationArn, true
}

// HasDestinationArn returns a boolean if a field has been set.
func (o *CreateDspStreamSubscriptionRequestContent) HasDestinationArn() bool {
	if o != nil && !IsNil(o.DestinationArn) {
		return true
	}

	return false
}

// SetDestinationArn gets a reference to the given string and assigns it to the DestinationArn field.
func (o *CreateDspStreamSubscriptionRequestContent) SetDestinationArn(v string) {
	o.DestinationArn = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *CreateDspStreamSubscriptionRequestContent) GetDestination() Destination {
	if o == nil || IsNil(o.Destination) {
		var ret Destination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDspStreamSubscriptionRequestContent) GetDestinationOk() (*Destination, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *CreateDspStreamSubscriptionRequestContent) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given Destination and assigns it to the Destination field.
func (o *CreateDspStreamSubscriptionRequestContent) SetDestination(v Destination) {
	o.Destination = &v
}

func (o CreateDspStreamSubscriptionRequestContent) ToMap() (map[string]interface{}, error) {
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

type NullableCreateDspStreamSubscriptionRequestContent struct {
	value *CreateDspStreamSubscriptionRequestContent
	isSet bool
}

func (v NullableCreateDspStreamSubscriptionRequestContent) Get() *CreateDspStreamSubscriptionRequestContent {
	return v.value
}

func (v *NullableCreateDspStreamSubscriptionRequestContent) Set(val *CreateDspStreamSubscriptionRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDspStreamSubscriptionRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDspStreamSubscriptionRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDspStreamSubscriptionRequestContent(val *CreateDspStreamSubscriptionRequestContent) *NullableCreateDspStreamSubscriptionRequestContent {
	return &NullableCreateDspStreamSubscriptionRequestContent{value: val, isSet: true}
}

func (v NullableCreateDspStreamSubscriptionRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateDspStreamSubscriptionRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
