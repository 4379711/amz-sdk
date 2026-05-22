package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the GetStreamSubscriptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStreamSubscriptionResponseContent{}

// GetStreamSubscriptionResponseContent struct for GetStreamSubscriptionResponseContent
type GetStreamSubscriptionResponseContent struct {
	Subscription StreamSubscription `json:"subscription"`
}

type _GetStreamSubscriptionResponseContent GetStreamSubscriptionResponseContent

// NewGetStreamSubscriptionResponseContent instantiates a new GetStreamSubscriptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStreamSubscriptionResponseContent(subscription StreamSubscription) *GetStreamSubscriptionResponseContent {
	this := GetStreamSubscriptionResponseContent{}
	this.Subscription = subscription
	return &this
}

// NewGetStreamSubscriptionResponseContentWithDefaults instantiates a new GetStreamSubscriptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStreamSubscriptionResponseContentWithDefaults() *GetStreamSubscriptionResponseContent {
	this := GetStreamSubscriptionResponseContent{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *GetStreamSubscriptionResponseContent) GetSubscription() StreamSubscription {
	if o == nil {
		var ret StreamSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *GetStreamSubscriptionResponseContent) GetSubscriptionOk() (*StreamSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *GetStreamSubscriptionResponseContent) SetSubscription(v StreamSubscription) {
	o.Subscription = v
}

func (o GetStreamSubscriptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	return toSerialize, nil
}

type NullableGetStreamSubscriptionResponseContent struct {
	value *GetStreamSubscriptionResponseContent
	isSet bool
}

func (v NullableGetStreamSubscriptionResponseContent) Get() *GetStreamSubscriptionResponseContent {
	return v.value
}

func (v *NullableGetStreamSubscriptionResponseContent) Set(val *GetStreamSubscriptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStreamSubscriptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStreamSubscriptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStreamSubscriptionResponseContent(val *GetStreamSubscriptionResponseContent) *NullableGetStreamSubscriptionResponseContent {
	return &NullableGetStreamSubscriptionResponseContent{value: val, isSet: true}
}

func (v NullableGetStreamSubscriptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetStreamSubscriptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
