package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the ListLakeSubscriptionsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListLakeSubscriptionsResponseContent{}

// ListLakeSubscriptionsResponseContent struct for ListLakeSubscriptionsResponseContent
type ListLakeSubscriptionsResponseContent struct {
	Subscriptions []LakeSubscription `json:"subscriptions"`
	// Token which can be used to get the next page of results, if more entries exist
	NextToken *string `json:"nextToken,omitempty"`
}

type _ListLakeSubscriptionsResponseContent ListLakeSubscriptionsResponseContent

// NewListLakeSubscriptionsResponseContent instantiates a new ListLakeSubscriptionsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLakeSubscriptionsResponseContent(subscriptions []LakeSubscription) *ListLakeSubscriptionsResponseContent {
	this := ListLakeSubscriptionsResponseContent{}
	this.Subscriptions = subscriptions
	return &this
}

// NewListLakeSubscriptionsResponseContentWithDefaults instantiates a new ListLakeSubscriptionsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLakeSubscriptionsResponseContentWithDefaults() *ListLakeSubscriptionsResponseContent {
	this := ListLakeSubscriptionsResponseContent{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value
func (o *ListLakeSubscriptionsResponseContent) GetSubscriptions() []LakeSubscription {
	if o == nil {
		var ret []LakeSubscription
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *ListLakeSubscriptionsResponseContent) GetSubscriptionsOk() ([]LakeSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *ListLakeSubscriptionsResponseContent) SetSubscriptions(v []LakeSubscription) {
	o.Subscriptions = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListLakeSubscriptionsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLakeSubscriptionsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListLakeSubscriptionsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListLakeSubscriptionsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListLakeSubscriptionsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptions"] = o.Subscriptions
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableListLakeSubscriptionsResponseContent struct {
	value *ListLakeSubscriptionsResponseContent
	isSet bool
}

func (v NullableListLakeSubscriptionsResponseContent) Get() *ListLakeSubscriptionsResponseContent {
	return v.value
}

func (v *NullableListLakeSubscriptionsResponseContent) Set(val *ListLakeSubscriptionsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListLakeSubscriptionsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListLakeSubscriptionsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLakeSubscriptionsResponseContent(val *ListLakeSubscriptionsResponseContent) *NullableListLakeSubscriptionsResponseContent {
	return &NullableListLakeSubscriptionsResponseContent{value: val, isSet: true}
}

func (v NullableListLakeSubscriptionsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListLakeSubscriptionsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
