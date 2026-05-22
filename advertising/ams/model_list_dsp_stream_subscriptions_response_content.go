package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the ListDspStreamSubscriptionsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDspStreamSubscriptionsResponseContent{}

// ListDspStreamSubscriptionsResponseContent struct for ListDspStreamSubscriptionsResponseContent
type ListDspStreamSubscriptionsResponseContent struct {
	Subscriptions []StreamSubscription `json:"subscriptions"`
	// Token which can be used to get the next page of results, if more entries exist
	NextToken *string `json:"nextToken,omitempty"`
}

type _ListDspStreamSubscriptionsResponseContent ListDspStreamSubscriptionsResponseContent

// NewListDspStreamSubscriptionsResponseContent instantiates a new ListDspStreamSubscriptionsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDspStreamSubscriptionsResponseContent(subscriptions []StreamSubscription) *ListDspStreamSubscriptionsResponseContent {
	this := ListDspStreamSubscriptionsResponseContent{}
	this.Subscriptions = subscriptions
	return &this
}

// NewListDspStreamSubscriptionsResponseContentWithDefaults instantiates a new ListDspStreamSubscriptionsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDspStreamSubscriptionsResponseContentWithDefaults() *ListDspStreamSubscriptionsResponseContent {
	this := ListDspStreamSubscriptionsResponseContent{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value
func (o *ListDspStreamSubscriptionsResponseContent) GetSubscriptions() []StreamSubscription {
	if o == nil {
		var ret []StreamSubscription
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *ListDspStreamSubscriptionsResponseContent) GetSubscriptionsOk() ([]StreamSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *ListDspStreamSubscriptionsResponseContent) SetSubscriptions(v []StreamSubscription) {
	o.Subscriptions = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListDspStreamSubscriptionsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDspStreamSubscriptionsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListDspStreamSubscriptionsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListDspStreamSubscriptionsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListDspStreamSubscriptionsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptions"] = o.Subscriptions
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableListDspStreamSubscriptionsResponseContent struct {
	value *ListDspStreamSubscriptionsResponseContent
	isSet bool
}

func (v NullableListDspStreamSubscriptionsResponseContent) Get() *ListDspStreamSubscriptionsResponseContent {
	return v.value
}

func (v *NullableListDspStreamSubscriptionsResponseContent) Set(val *ListDspStreamSubscriptionsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListDspStreamSubscriptionsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListDspStreamSubscriptionsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDspStreamSubscriptionsResponseContent(val *ListDspStreamSubscriptionsResponseContent) *NullableListDspStreamSubscriptionsResponseContent {
	return &NullableListDspStreamSubscriptionsResponseContent{value: val, isSet: true}
}

func (v NullableListDspStreamSubscriptionsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListDspStreamSubscriptionsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
