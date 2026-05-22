package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the FirehoseDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirehoseDestination{}

// FirehoseDestination struct for FirehoseDestination
type FirehoseDestination struct {
	DeliveryStreamArn   string `json:"deliveryStreamArn"`
	SubscriptionRoleArn string `json:"subscriptionRoleArn"`
	SubscriberRoleArn   string `json:"subscriberRoleArn"`
}

type _FirehoseDestination FirehoseDestination

// NewFirehoseDestination instantiates a new FirehoseDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirehoseDestination(deliveryStreamArn string, subscriptionRoleArn string, subscriberRoleArn string) *FirehoseDestination {
	this := FirehoseDestination{}
	this.DeliveryStreamArn = deliveryStreamArn
	this.SubscriptionRoleArn = subscriptionRoleArn
	this.SubscriberRoleArn = subscriberRoleArn
	return &this
}

// NewFirehoseDestinationWithDefaults instantiates a new FirehoseDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirehoseDestinationWithDefaults() *FirehoseDestination {
	this := FirehoseDestination{}
	return &this
}

// GetDeliveryStreamArn returns the DeliveryStreamArn field value
func (o *FirehoseDestination) GetDeliveryStreamArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryStreamArn
}

// GetDeliveryStreamArnOk returns a tuple with the DeliveryStreamArn field value
// and a boolean to check if the value has been set.
func (o *FirehoseDestination) GetDeliveryStreamArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryStreamArn, true
}

// SetDeliveryStreamArn sets field value
func (o *FirehoseDestination) SetDeliveryStreamArn(v string) {
	o.DeliveryStreamArn = v
}

// GetSubscriptionRoleArn returns the SubscriptionRoleArn field value
func (o *FirehoseDestination) GetSubscriptionRoleArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionRoleArn
}

// GetSubscriptionRoleArnOk returns a tuple with the SubscriptionRoleArn field value
// and a boolean to check if the value has been set.
func (o *FirehoseDestination) GetSubscriptionRoleArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionRoleArn, true
}

// SetSubscriptionRoleArn sets field value
func (o *FirehoseDestination) SetSubscriptionRoleArn(v string) {
	o.SubscriptionRoleArn = v
}

// GetSubscriberRoleArn returns the SubscriberRoleArn field value
func (o *FirehoseDestination) GetSubscriberRoleArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriberRoleArn
}

// GetSubscriberRoleArnOk returns a tuple with the SubscriberRoleArn field value
// and a boolean to check if the value has been set.
func (o *FirehoseDestination) GetSubscriberRoleArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriberRoleArn, true
}

// SetSubscriberRoleArn sets field value
func (o *FirehoseDestination) SetSubscriberRoleArn(v string) {
	o.SubscriberRoleArn = v
}

func (o FirehoseDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deliveryStreamArn"] = o.DeliveryStreamArn
	toSerialize["subscriptionRoleArn"] = o.SubscriptionRoleArn
	toSerialize["subscriberRoleArn"] = o.SubscriberRoleArn
	return toSerialize, nil
}

type NullableFirehoseDestination struct {
	value *FirehoseDestination
	isSet bool
}

func (v NullableFirehoseDestination) Get() *FirehoseDestination {
	return v.value
}

func (v *NullableFirehoseDestination) Set(val *FirehoseDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableFirehoseDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableFirehoseDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirehoseDestination(val *FirehoseDestination) *NullableFirehoseDestination {
	return &NullableFirehoseDestination{value: val, isSet: true}
}

func (v NullableFirehoseDestination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFirehoseDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
