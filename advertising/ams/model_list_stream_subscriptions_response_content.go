package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the ListStreamSubscriptionsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListStreamSubscriptionsResponseContent{}

// ListStreamSubscriptionsResponseContent struct for ListStreamSubscriptionsResponseContent
type ListStreamSubscriptionsResponseContent struct {
	Subscriptions []StreamSubscription `json:"subscriptions"`
	// Token which can be used to get the next page of results, if more entries exist
	NextToken *string `json:"nextToken,omitempty"`
}

type _ListStreamSubscriptionsResponseContent ListStreamSubscriptionsResponseContent

// NewListStreamSubscriptionsResponseContent instantiates a new ListStreamSubscriptionsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStreamSubscriptionsResponseContent(subscriptions []StreamSubscription) *ListStreamSubscriptionsResponseContent {
	this := ListStreamSubscriptionsResponseContent{}
	this.Subscriptions = subscriptions
	return &this
}

// NewListStreamSubscriptionsResponseContentWithDefaults instantiates a new ListStreamSubscriptionsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStreamSubscriptionsResponseContentWithDefaults() *ListStreamSubscriptionsResponseContent {
	this := ListStreamSubscriptionsResponseContent{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value
func (o *ListStreamSubscriptionsResponseContent) GetSubscriptions() []StreamSubscription {
	if o == nil {
		var ret []StreamSubscription
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *ListStreamSubscriptionsResponseContent) GetSubscriptionsOk() ([]StreamSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *ListStreamSubscriptionsResponseContent) SetSubscriptions(v []StreamSubscription) {
	o.Subscriptions = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListStreamSubscriptionsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamSubscriptionsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListStreamSubscriptionsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListStreamSubscriptionsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListStreamSubscriptionsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptions"] = o.Subscriptions
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableListStreamSubscriptionsResponseContent struct {
	value *ListStreamSubscriptionsResponseContent
	isSet bool
}

func (v NullableListStreamSubscriptionsResponseContent) Get() *ListStreamSubscriptionsResponseContent {
	return v.value
}

func (v *NullableListStreamSubscriptionsResponseContent) Set(val *ListStreamSubscriptionsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListStreamSubscriptionsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListStreamSubscriptionsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStreamSubscriptionsResponseContent(val *ListStreamSubscriptionsResponseContent) *NullableListStreamSubscriptionsResponseContent {
	return &NullableListStreamSubscriptionsResponseContent{value: val, isSet: true}
}

func (v NullableListStreamSubscriptionsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListStreamSubscriptionsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
