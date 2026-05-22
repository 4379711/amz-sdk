package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the DeliveryPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryPolicy{}

// DeliveryPolicy The policy for a delivery offering.
type DeliveryPolicy struct {
	Message *DeliveryMessage `json:"message,omitempty"`
}

// NewDeliveryPolicy instantiates a new DeliveryPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryPolicy() *DeliveryPolicy {
	this := DeliveryPolicy{}
	return &this
}

// NewDeliveryPolicyWithDefaults instantiates a new DeliveryPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryPolicyWithDefaults() *DeliveryPolicy {
	this := DeliveryPolicy{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeliveryPolicy) GetMessage() DeliveryMessage {
	if o == nil || IsNil(o.Message) {
		var ret DeliveryMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryPolicy) GetMessageOk() (*DeliveryMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeliveryPolicy) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given DeliveryMessage and assigns it to the Message field.
func (o *DeliveryPolicy) SetMessage(v DeliveryMessage) {
	o.Message = &v
}

func (o DeliveryPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableDeliveryPolicy struct {
	value *DeliveryPolicy
	isSet bool
}

func (v NullableDeliveryPolicy) Get() *DeliveryPolicy {
	return v.value
}

func (v *NullableDeliveryPolicy) Set(val *DeliveryPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryPolicy(val *DeliveryPolicy) *NullableDeliveryPolicy {
	return &NullableDeliveryPolicy{value: val, isSet: true}
}

func (v NullableDeliveryPolicy) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeliveryPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
