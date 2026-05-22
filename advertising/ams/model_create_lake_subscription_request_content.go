package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateLakeSubscriptionRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLakeSubscriptionRequestContent{}

// CreateLakeSubscriptionRequestContent struct for CreateLakeSubscriptionRequestContent
type CreateLakeSubscriptionRequestContent struct {
	// AWS ARN of the principal associated with the subscription. Supported principal types: - AWS Account Principal
	Principal string `json:"principal"`
	// Additional details associated with the subscription
	Notes *string `json:"notes,omitempty"`
	// Unique value supplied by the caller used to track identical API requests. Should request be re-tried, the caller should supply the same value. We recommend using GUID.
	ClientRequestToken string `json:"clientRequestToken"`
	// Identifier of data set, callers can be subscribed to. Please refer to https://advertising.amazon.com/API/docs/en-us/amazon-marketing-stream/data-guide for the list of all data sets.
	DataSetId string `json:"dataSetId"`
}

type _CreateLakeSubscriptionRequestContent CreateLakeSubscriptionRequestContent

// NewCreateLakeSubscriptionRequestContent instantiates a new CreateLakeSubscriptionRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLakeSubscriptionRequestContent(principal string, clientRequestToken string, dataSetId string) *CreateLakeSubscriptionRequestContent {
	this := CreateLakeSubscriptionRequestContent{}
	this.Principal = principal
	this.ClientRequestToken = clientRequestToken
	this.DataSetId = dataSetId
	return &this
}

// NewCreateLakeSubscriptionRequestContentWithDefaults instantiates a new CreateLakeSubscriptionRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLakeSubscriptionRequestContentWithDefaults() *CreateLakeSubscriptionRequestContent {
	this := CreateLakeSubscriptionRequestContent{}
	return &this
}

// GetPrincipal returns the Principal field value
func (o *CreateLakeSubscriptionRequestContent) GetPrincipal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *CreateLakeSubscriptionRequestContent) GetPrincipalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *CreateLakeSubscriptionRequestContent) SetPrincipal(v string) {
	o.Principal = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CreateLakeSubscriptionRequestContent) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLakeSubscriptionRequestContent) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CreateLakeSubscriptionRequestContent) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CreateLakeSubscriptionRequestContent) SetNotes(v string) {
	o.Notes = &v
}

// GetClientRequestToken returns the ClientRequestToken field value
func (o *CreateLakeSubscriptionRequestContent) GetClientRequestToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientRequestToken
}

// GetClientRequestTokenOk returns a tuple with the ClientRequestToken field value
// and a boolean to check if the value has been set.
func (o *CreateLakeSubscriptionRequestContent) GetClientRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientRequestToken, true
}

// SetClientRequestToken sets field value
func (o *CreateLakeSubscriptionRequestContent) SetClientRequestToken(v string) {
	o.ClientRequestToken = v
}

// GetDataSetId returns the DataSetId field value
func (o *CreateLakeSubscriptionRequestContent) GetDataSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSetId
}

// GetDataSetIdOk returns a tuple with the DataSetId field value
// and a boolean to check if the value has been set.
func (o *CreateLakeSubscriptionRequestContent) GetDataSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSetId, true
}

// SetDataSetId sets field value
func (o *CreateLakeSubscriptionRequestContent) SetDataSetId(v string) {
	o.DataSetId = v
}

func (o CreateLakeSubscriptionRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal"] = o.Principal
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["clientRequestToken"] = o.ClientRequestToken
	toSerialize["dataSetId"] = o.DataSetId
	return toSerialize, nil
}

type NullableCreateLakeSubscriptionRequestContent struct {
	value *CreateLakeSubscriptionRequestContent
	isSet bool
}

func (v NullableCreateLakeSubscriptionRequestContent) Get() *CreateLakeSubscriptionRequestContent {
	return v.value
}

func (v *NullableCreateLakeSubscriptionRequestContent) Set(val *CreateLakeSubscriptionRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLakeSubscriptionRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLakeSubscriptionRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLakeSubscriptionRequestContent(val *CreateLakeSubscriptionRequestContent) *NullableCreateLakeSubscriptionRequestContent {
	return &NullableCreateLakeSubscriptionRequestContent{value: val, isSet: true}
}

func (v NullableCreateLakeSubscriptionRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateLakeSubscriptionRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
