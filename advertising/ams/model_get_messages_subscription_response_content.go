package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the GetMessagesSubscriptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMessagesSubscriptionResponseContent{}

// GetMessagesSubscriptionResponseContent struct for GetMessagesSubscriptionResponseContent
type GetMessagesSubscriptionResponseContent struct {
	MessagesSubscription MessagesSubscription `json:"messagesSubscription"`
}

type _GetMessagesSubscriptionResponseContent GetMessagesSubscriptionResponseContent

// NewGetMessagesSubscriptionResponseContent instantiates a new GetMessagesSubscriptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMessagesSubscriptionResponseContent(messagesSubscription MessagesSubscription) *GetMessagesSubscriptionResponseContent {
	this := GetMessagesSubscriptionResponseContent{}
	this.MessagesSubscription = messagesSubscription
	return &this
}

// NewGetMessagesSubscriptionResponseContentWithDefaults instantiates a new GetMessagesSubscriptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMessagesSubscriptionResponseContentWithDefaults() *GetMessagesSubscriptionResponseContent {
	this := GetMessagesSubscriptionResponseContent{}
	return &this
}

// GetMessagesSubscription returns the MessagesSubscription field value
func (o *GetMessagesSubscriptionResponseContent) GetMessagesSubscription() MessagesSubscription {
	if o == nil {
		var ret MessagesSubscription
		return ret
	}

	return o.MessagesSubscription
}

// GetMessagesSubscriptionOk returns a tuple with the MessagesSubscription field value
// and a boolean to check if the value has been set.
func (o *GetMessagesSubscriptionResponseContent) GetMessagesSubscriptionOk() (*MessagesSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessagesSubscription, true
}

// SetMessagesSubscription sets field value
func (o *GetMessagesSubscriptionResponseContent) SetMessagesSubscription(v MessagesSubscription) {
	o.MessagesSubscription = v
}

func (o GetMessagesSubscriptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messagesSubscription"] = o.MessagesSubscription
	return toSerialize, nil
}

type NullableGetMessagesSubscriptionResponseContent struct {
	value *GetMessagesSubscriptionResponseContent
	isSet bool
}

func (v NullableGetMessagesSubscriptionResponseContent) Get() *GetMessagesSubscriptionResponseContent {
	return v.value
}

func (v *NullableGetMessagesSubscriptionResponseContent) Set(val *GetMessagesSubscriptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMessagesSubscriptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMessagesSubscriptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMessagesSubscriptionResponseContent(val *GetMessagesSubscriptionResponseContent) *NullableGetMessagesSubscriptionResponseContent {
	return &NullableGetMessagesSubscriptionResponseContent{value: val, isSet: true}
}

func (v NullableGetMessagesSubscriptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetMessagesSubscriptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
