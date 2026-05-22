package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the InboundOrderReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundOrderReference{}

// InboundOrderReference A response that contains the reference identifiers for the newly created or updated inbound order. Consists of an order ID and version.
type InboundOrderReference struct {
	// Order ID of the inbound order.
	OrderId string `json:"orderId"`
}

type _InboundOrderReference InboundOrderReference

// NewInboundOrderReference instantiates a new InboundOrderReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundOrderReference(orderId string) *InboundOrderReference {
	this := InboundOrderReference{}
	this.OrderId = orderId
	return &this
}

// NewInboundOrderReferenceWithDefaults instantiates a new InboundOrderReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundOrderReferenceWithDefaults() *InboundOrderReference {
	this := InboundOrderReference{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *InboundOrderReference) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *InboundOrderReference) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *InboundOrderReference) SetOrderId(v string) {
	o.OrderId = v
}

func (o InboundOrderReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderId"] = o.OrderId
	return toSerialize, nil
}

type NullableInboundOrderReference struct {
	value *InboundOrderReference
	isSet bool
}

func (v NullableInboundOrderReference) Get() *InboundOrderReference {
	return v.value
}

func (v *NullableInboundOrderReference) Set(val *InboundOrderReference) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundOrderReference) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundOrderReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundOrderReference(val *InboundOrderReference) *NullableInboundOrderReference {
	return &NullableInboundOrderReference{value: val, isSet: true}
}

func (v NullableInboundOrderReference) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInboundOrderReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
