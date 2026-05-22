package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the GetLakeSubscriptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLakeSubscriptionResponseContent{}

// GetLakeSubscriptionResponseContent struct for GetLakeSubscriptionResponseContent
type GetLakeSubscriptionResponseContent struct {
	Subscription LakeSubscription `json:"subscription"`
}

type _GetLakeSubscriptionResponseContent GetLakeSubscriptionResponseContent

// NewGetLakeSubscriptionResponseContent instantiates a new GetLakeSubscriptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLakeSubscriptionResponseContent(subscription LakeSubscription) *GetLakeSubscriptionResponseContent {
	this := GetLakeSubscriptionResponseContent{}
	this.Subscription = subscription
	return &this
}

// NewGetLakeSubscriptionResponseContentWithDefaults instantiates a new GetLakeSubscriptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLakeSubscriptionResponseContentWithDefaults() *GetLakeSubscriptionResponseContent {
	this := GetLakeSubscriptionResponseContent{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *GetLakeSubscriptionResponseContent) GetSubscription() LakeSubscription {
	if o == nil {
		var ret LakeSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *GetLakeSubscriptionResponseContent) GetSubscriptionOk() (*LakeSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *GetLakeSubscriptionResponseContent) SetSubscription(v LakeSubscription) {
	o.Subscription = v
}

func (o GetLakeSubscriptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	return toSerialize, nil
}

type NullableGetLakeSubscriptionResponseContent struct {
	value *GetLakeSubscriptionResponseContent
	isSet bool
}

func (v NullableGetLakeSubscriptionResponseContent) Get() *GetLakeSubscriptionResponseContent {
	return v.value
}

func (v *NullableGetLakeSubscriptionResponseContent) Set(val *GetLakeSubscriptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLakeSubscriptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLakeSubscriptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLakeSubscriptionResponseContent(val *GetLakeSubscriptionResponseContent) *NullableGetLakeSubscriptionResponseContent {
	return &NullableGetLakeSubscriptionResponseContent{value: val, isSet: true}
}

func (v NullableGetLakeSubscriptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetLakeSubscriptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
