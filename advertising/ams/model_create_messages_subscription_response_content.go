package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateMessagesSubscriptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMessagesSubscriptionResponseContent{}

// CreateMessagesSubscriptionResponseContent struct for CreateMessagesSubscriptionResponseContent
type CreateMessagesSubscriptionResponseContent struct {
	// Unique value supplied by the caller used to track identical API requests. Should request be re-tried, the caller should supply the same value. We recommend using GUID.
	ClientToken string `json:"clientToken"`
	// Subscription identifier. Ex: amzn1.fead.cs1.sample
	MessagesSubscriptionId string `json:"messagesSubscriptionId"`
	// version for the  Subscription. Defaults to 1 upon creation. This is used for tracking updates to the subscription. Every update increments the version by 1.
	Version float32 `json:"version"`
}

type _CreateMessagesSubscriptionResponseContent CreateMessagesSubscriptionResponseContent

// NewCreateMessagesSubscriptionResponseContent instantiates a new CreateMessagesSubscriptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessagesSubscriptionResponseContent(clientToken string, messagesSubscriptionId string, version float32) *CreateMessagesSubscriptionResponseContent {
	this := CreateMessagesSubscriptionResponseContent{}
	this.ClientToken = clientToken
	this.MessagesSubscriptionId = messagesSubscriptionId
	this.Version = version
	return &this
}

// NewCreateMessagesSubscriptionResponseContentWithDefaults instantiates a new CreateMessagesSubscriptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessagesSubscriptionResponseContentWithDefaults() *CreateMessagesSubscriptionResponseContent {
	this := CreateMessagesSubscriptionResponseContent{}
	return &this
}

// GetClientToken returns the ClientToken field value
func (o *CreateMessagesSubscriptionResponseContent) GetClientToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientToken
}

// GetClientTokenOk returns a tuple with the ClientToken field value
// and a boolean to check if the value has been set.
func (o *CreateMessagesSubscriptionResponseContent) GetClientTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientToken, true
}

// SetClientToken sets field value
func (o *CreateMessagesSubscriptionResponseContent) SetClientToken(v string) {
	o.ClientToken = v
}

// GetMessagesSubscriptionId returns the MessagesSubscriptionId field value
func (o *CreateMessagesSubscriptionResponseContent) GetMessagesSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessagesSubscriptionId
}

// GetMessagesSubscriptionIdOk returns a tuple with the MessagesSubscriptionId field value
// and a boolean to check if the value has been set.
func (o *CreateMessagesSubscriptionResponseContent) GetMessagesSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessagesSubscriptionId, true
}

// SetMessagesSubscriptionId sets field value
func (o *CreateMessagesSubscriptionResponseContent) SetMessagesSubscriptionId(v string) {
	o.MessagesSubscriptionId = v
}

// GetVersion returns the Version field value
func (o *CreateMessagesSubscriptionResponseContent) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CreateMessagesSubscriptionResponseContent) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CreateMessagesSubscriptionResponseContent) SetVersion(v float32) {
	o.Version = v
}

func (o CreateMessagesSubscriptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientToken"] = o.ClientToken
	toSerialize["messagesSubscriptionId"] = o.MessagesSubscriptionId
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableCreateMessagesSubscriptionResponseContent struct {
	value *CreateMessagesSubscriptionResponseContent
	isSet bool
}

func (v NullableCreateMessagesSubscriptionResponseContent) Get() *CreateMessagesSubscriptionResponseContent {
	return v.value
}

func (v *NullableCreateMessagesSubscriptionResponseContent) Set(val *CreateMessagesSubscriptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessagesSubscriptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessagesSubscriptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessagesSubscriptionResponseContent(val *CreateMessagesSubscriptionResponseContent) *NullableCreateMessagesSubscriptionResponseContent {
	return &NullableCreateMessagesSubscriptionResponseContent{value: val, isSet: true}
}

func (v NullableCreateMessagesSubscriptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateMessagesSubscriptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
