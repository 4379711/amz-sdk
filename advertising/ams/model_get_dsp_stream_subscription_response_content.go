package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the GetDspStreamSubscriptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDspStreamSubscriptionResponseContent{}

// GetDspStreamSubscriptionResponseContent struct for GetDspStreamSubscriptionResponseContent
type GetDspStreamSubscriptionResponseContent struct {
	Subscription StreamSubscription `json:"subscription"`
}

type _GetDspStreamSubscriptionResponseContent GetDspStreamSubscriptionResponseContent

// NewGetDspStreamSubscriptionResponseContent instantiates a new GetDspStreamSubscriptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDspStreamSubscriptionResponseContent(subscription StreamSubscription) *GetDspStreamSubscriptionResponseContent {
	this := GetDspStreamSubscriptionResponseContent{}
	this.Subscription = subscription
	return &this
}

// NewGetDspStreamSubscriptionResponseContentWithDefaults instantiates a new GetDspStreamSubscriptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDspStreamSubscriptionResponseContentWithDefaults() *GetDspStreamSubscriptionResponseContent {
	this := GetDspStreamSubscriptionResponseContent{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *GetDspStreamSubscriptionResponseContent) GetSubscription() StreamSubscription {
	if o == nil {
		var ret StreamSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *GetDspStreamSubscriptionResponseContent) GetSubscriptionOk() (*StreamSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *GetDspStreamSubscriptionResponseContent) SetSubscription(v StreamSubscription) {
	o.Subscription = v
}

func (o GetDspStreamSubscriptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	return toSerialize, nil
}

type NullableGetDspStreamSubscriptionResponseContent struct {
	value *GetDspStreamSubscriptionResponseContent
	isSet bool
}

func (v NullableGetDspStreamSubscriptionResponseContent) Get() *GetDspStreamSubscriptionResponseContent {
	return v.value
}

func (v *NullableGetDspStreamSubscriptionResponseContent) Set(val *GetDspStreamSubscriptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDspStreamSubscriptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDspStreamSubscriptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDspStreamSubscriptionResponseContent(val *GetDspStreamSubscriptionResponseContent) *NullableGetDspStreamSubscriptionResponseContent {
	return &NullableGetDspStreamSubscriptionResponseContent{value: val, isSet: true}
}

func (v NullableGetDspStreamSubscriptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetDspStreamSubscriptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
