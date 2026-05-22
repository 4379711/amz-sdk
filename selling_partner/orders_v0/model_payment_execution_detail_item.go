package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the PaymentExecutionDetailItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentExecutionDetailItem{}

// PaymentExecutionDetailItem Information about a sub-payment method used to pay for a COD order.
type PaymentExecutionDetailItem struct {
	Payment Money `json:"Payment"`
	// A sub-payment method for a COD order.  **Possible values**: * `COD`: Cash on delivery  * `GC`: Gift card  * `PointsAccount`: Amazon Points * `Invoice`: Invoice
	PaymentMethod string `json:"PaymentMethod"`
}

type _PaymentExecutionDetailItem PaymentExecutionDetailItem

// NewPaymentExecutionDetailItem instantiates a new PaymentExecutionDetailItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentExecutionDetailItem(payment Money, paymentMethod string) *PaymentExecutionDetailItem {
	this := PaymentExecutionDetailItem{}
	this.Payment = payment
	this.PaymentMethod = paymentMethod
	return &this
}

// NewPaymentExecutionDetailItemWithDefaults instantiates a new PaymentExecutionDetailItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentExecutionDetailItemWithDefaults() *PaymentExecutionDetailItem {
	this := PaymentExecutionDetailItem{}
	return &this
}

// GetPayment returns the Payment field value
func (o *PaymentExecutionDetailItem) GetPayment() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value
// and a boolean to check if the value has been set.
func (o *PaymentExecutionDetailItem) GetPaymentOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payment, true
}

// SetPayment sets field value
func (o *PaymentExecutionDetailItem) SetPayment(v Money) {
	o.Payment = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *PaymentExecutionDetailItem) GetPaymentMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *PaymentExecutionDetailItem) GetPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *PaymentExecutionDetailItem) SetPaymentMethod(v string) {
	o.PaymentMethod = v
}

func (o PaymentExecutionDetailItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Payment"] = o.Payment
	toSerialize["PaymentMethod"] = o.PaymentMethod
	return toSerialize, nil
}

type NullablePaymentExecutionDetailItem struct {
	value *PaymentExecutionDetailItem
	isSet bool
}

func (v NullablePaymentExecutionDetailItem) Get() *PaymentExecutionDetailItem {
	return v.value
}

func (v *NullablePaymentExecutionDetailItem) Set(val *PaymentExecutionDetailItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentExecutionDetailItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentExecutionDetailItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentExecutionDetailItem(val *PaymentExecutionDetailItem) *NullablePaymentExecutionDetailItem {
	return &NullablePaymentExecutionDetailItem{value: val, isSet: true}
}

func (v NullablePaymentExecutionDetailItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePaymentExecutionDetailItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
