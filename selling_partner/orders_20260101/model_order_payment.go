package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderPayment{}

// OrderPayment Payment information about the order.
type OrderPayment struct {
	// A list of payment executions for the order.
	PaymentExecutions []PaymentExecution `json:"paymentExecutions,omitempty"`
}

// NewOrderPayment instantiates a new OrderPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderPayment() *OrderPayment {
	this := OrderPayment{}
	return &this
}

// NewOrderPaymentWithDefaults instantiates a new OrderPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderPaymentWithDefaults() *OrderPayment {
	this := OrderPayment{}
	return &this
}

// GetPaymentExecutions returns the PaymentExecutions field value if set, zero value otherwise.
func (o *OrderPayment) GetPaymentExecutions() []PaymentExecution {
	if o == nil || IsNil(o.PaymentExecutions) {
		var ret []PaymentExecution
		return ret
	}
	return o.PaymentExecutions
}

// GetPaymentExecutionsOk returns a tuple with the PaymentExecutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPayment) GetPaymentExecutionsOk() ([]PaymentExecution, bool) {
	if o == nil || IsNil(o.PaymentExecutions) {
		return nil, false
	}
	return o.PaymentExecutions, true
}

// HasPaymentExecutions returns a boolean if a field has been set.
func (o *OrderPayment) HasPaymentExecutions() bool {
	if o != nil && !IsNil(o.PaymentExecutions) {
		return true
	}

	return false
}

// SetPaymentExecutions gets a reference to the given []PaymentExecution and assigns it to the PaymentExecutions field.
func (o *OrderPayment) SetPaymentExecutions(v []PaymentExecution) {
	o.PaymentExecutions = v
}

func (o OrderPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentExecutions) {
		toSerialize["paymentExecutions"] = o.PaymentExecutions
	}
	return toSerialize, nil
}

type NullableOrderPayment struct {
	value *OrderPayment
	isSet bool
}

func (v NullableOrderPayment) Get() *OrderPayment {
	return v.value
}

func (v *NullableOrderPayment) Set(val *OrderPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPayment(val *OrderPayment) *NullableOrderPayment {
	return &NullableOrderPayment{value: val, isSet: true}
}

func (v NullableOrderPayment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
