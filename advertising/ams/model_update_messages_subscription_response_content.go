package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateMessagesSubscriptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMessagesSubscriptionResponseContent{}

// UpdateMessagesSubscriptionResponseContent struct for UpdateMessagesSubscriptionResponseContent
type UpdateMessagesSubscriptionResponseContent struct {
	// Subscription identifier. Ex: amzn1.fead.cd1.sample
	MessagesSubscriptionId string `json:"messagesSubscriptionId"`
	// Updated version of the  Subscription.
	Version float32 `json:"version"`
}

type _UpdateMessagesSubscriptionResponseContent UpdateMessagesSubscriptionResponseContent

// NewUpdateMessagesSubscriptionResponseContent instantiates a new UpdateMessagesSubscriptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMessagesSubscriptionResponseContent(messagesSubscriptionId string, version float32) *UpdateMessagesSubscriptionResponseContent {
	this := UpdateMessagesSubscriptionResponseContent{}
	this.MessagesSubscriptionId = messagesSubscriptionId
	this.Version = version
	return &this
}

// NewUpdateMessagesSubscriptionResponseContentWithDefaults instantiates a new UpdateMessagesSubscriptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMessagesSubscriptionResponseContentWithDefaults() *UpdateMessagesSubscriptionResponseContent {
	this := UpdateMessagesSubscriptionResponseContent{}
	return &this
}

// GetMessagesSubscriptionId returns the MessagesSubscriptionId field value
func (o *UpdateMessagesSubscriptionResponseContent) GetMessagesSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessagesSubscriptionId
}

// GetMessagesSubscriptionIdOk returns a tuple with the MessagesSubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UpdateMessagesSubscriptionResponseContent) GetMessagesSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessagesSubscriptionId, true
}

// SetMessagesSubscriptionId sets field value
func (o *UpdateMessagesSubscriptionResponseContent) SetMessagesSubscriptionId(v string) {
	o.MessagesSubscriptionId = v
}

// GetVersion returns the Version field value
func (o *UpdateMessagesSubscriptionResponseContent) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UpdateMessagesSubscriptionResponseContent) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UpdateMessagesSubscriptionResponseContent) SetVersion(v float32) {
	o.Version = v
}

func (o UpdateMessagesSubscriptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messagesSubscriptionId"] = o.MessagesSubscriptionId
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableUpdateMessagesSubscriptionResponseContent struct {
	value *UpdateMessagesSubscriptionResponseContent
	isSet bool
}

func (v NullableUpdateMessagesSubscriptionResponseContent) Get() *UpdateMessagesSubscriptionResponseContent {
	return v.value
}

func (v *NullableUpdateMessagesSubscriptionResponseContent) Set(val *UpdateMessagesSubscriptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMessagesSubscriptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMessagesSubscriptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMessagesSubscriptionResponseContent(val *UpdateMessagesSubscriptionResponseContent) *NullableUpdateMessagesSubscriptionResponseContent {
	return &NullableUpdateMessagesSubscriptionResponseContent{value: val, isSet: true}
}

func (v NullableUpdateMessagesSubscriptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateMessagesSubscriptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
