package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the ListMessagesSubscriptionsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMessagesSubscriptionsResponseContent{}

// ListMessagesSubscriptionsResponseContent struct for ListMessagesSubscriptionsResponseContent
type ListMessagesSubscriptionsResponseContent struct {
	// Token which can be used to get the next page of results, if more entries exist.
	NextToken             *string                `json:"nextToken,omitempty"`
	MessagesSubscriptions []MessagesSubscription `json:"messagesSubscriptions"`
}

type _ListMessagesSubscriptionsResponseContent ListMessagesSubscriptionsResponseContent

// NewListMessagesSubscriptionsResponseContent instantiates a new ListMessagesSubscriptionsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMessagesSubscriptionsResponseContent(messagesSubscriptions []MessagesSubscription) *ListMessagesSubscriptionsResponseContent {
	this := ListMessagesSubscriptionsResponseContent{}
	this.MessagesSubscriptions = messagesSubscriptions
	return &this
}

// NewListMessagesSubscriptionsResponseContentWithDefaults instantiates a new ListMessagesSubscriptionsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMessagesSubscriptionsResponseContentWithDefaults() *ListMessagesSubscriptionsResponseContent {
	this := ListMessagesSubscriptionsResponseContent{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListMessagesSubscriptionsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessagesSubscriptionsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListMessagesSubscriptionsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListMessagesSubscriptionsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetMessagesSubscriptions returns the MessagesSubscriptions field value
func (o *ListMessagesSubscriptionsResponseContent) GetMessagesSubscriptions() []MessagesSubscription {
	if o == nil {
		var ret []MessagesSubscription
		return ret
	}

	return o.MessagesSubscriptions
}

// GetMessagesSubscriptionsOk returns a tuple with the MessagesSubscriptions field value
// and a boolean to check if the value has been set.
func (o *ListMessagesSubscriptionsResponseContent) GetMessagesSubscriptionsOk() ([]MessagesSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessagesSubscriptions, true
}

// SetMessagesSubscriptions sets field value
func (o *ListMessagesSubscriptionsResponseContent) SetMessagesSubscriptions(v []MessagesSubscription) {
	o.MessagesSubscriptions = v
}

func (o ListMessagesSubscriptionsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	toSerialize["messagesSubscriptions"] = o.MessagesSubscriptions
	return toSerialize, nil
}

type NullableListMessagesSubscriptionsResponseContent struct {
	value *ListMessagesSubscriptionsResponseContent
	isSet bool
}

func (v NullableListMessagesSubscriptionsResponseContent) Get() *ListMessagesSubscriptionsResponseContent {
	return v.value
}

func (v *NullableListMessagesSubscriptionsResponseContent) Set(val *ListMessagesSubscriptionsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListMessagesSubscriptionsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListMessagesSubscriptionsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMessagesSubscriptionsResponseContent(val *ListMessagesSubscriptionsResponseContent) *NullableListMessagesSubscriptionsResponseContent {
	return &NullableListMessagesSubscriptionsResponseContent{value: val, isSet: true}
}

func (v NullableListMessagesSubscriptionsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListMessagesSubscriptionsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
