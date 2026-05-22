package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the Subscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subscription{}

// Subscription Information about the subscription.
type Subscription struct {
	// The subscription identifier generated when the subscription is created.
	SubscriptionId string `json:"subscriptionId"`
	// The version of the payload object to be used in the notification.
	PayloadVersion string `json:"payloadVersion"`
	// The identifier for the destination where notifications will be delivered.
	DestinationId       string               `json:"destinationId"`
	ProcessingDirective *ProcessingDirective `json:"processingDirective,omitempty"`
}

type _Subscription Subscription

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription(subscriptionId string, payloadVersion string, destinationId string) *Subscription {
	this := Subscription{}
	this.SubscriptionId = subscriptionId
	this.PayloadVersion = payloadVersion
	this.DestinationId = destinationId
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *Subscription) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *Subscription) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetPayloadVersion returns the PayloadVersion field value
func (o *Subscription) GetPayloadVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayloadVersion
}

// GetPayloadVersionOk returns a tuple with the PayloadVersion field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetPayloadVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadVersion, true
}

// SetPayloadVersion sets field value
func (o *Subscription) SetPayloadVersion(v string) {
	o.PayloadVersion = v
}

// GetDestinationId returns the DestinationId field value
func (o *Subscription) GetDestinationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetDestinationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationId, true
}

// SetDestinationId sets field value
func (o *Subscription) SetDestinationId(v string) {
	o.DestinationId = v
}

// GetProcessingDirective returns the ProcessingDirective field value if set, zero value otherwise.
func (o *Subscription) GetProcessingDirective() ProcessingDirective {
	if o == nil || IsNil(o.ProcessingDirective) {
		var ret ProcessingDirective
		return ret
	}
	return *o.ProcessingDirective
}

// GetProcessingDirectiveOk returns a tuple with the ProcessingDirective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetProcessingDirectiveOk() (*ProcessingDirective, bool) {
	if o == nil || IsNil(o.ProcessingDirective) {
		return nil, false
	}
	return o.ProcessingDirective, true
}

// HasProcessingDirective returns a boolean if a field has been set.
func (o *Subscription) HasProcessingDirective() bool {
	if o != nil && !IsNil(o.ProcessingDirective) {
		return true
	}

	return false
}

// SetProcessingDirective gets a reference to the given ProcessingDirective and assigns it to the ProcessingDirective field.
func (o *Subscription) SetProcessingDirective(v ProcessingDirective) {
	o.ProcessingDirective = &v
}

func (o Subscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["payloadVersion"] = o.PayloadVersion
	toSerialize["destinationId"] = o.DestinationId
	if !IsNil(o.ProcessingDirective) {
		toSerialize["processingDirective"] = o.ProcessingDirective
	}
	return toSerialize, nil
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
