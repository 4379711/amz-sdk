package fulfillment_outbound_20200701

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the PaymentInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentInformation{}

// PaymentInformation The attributes related to the payment made from customer to seller for this order.
type PaymentInformation struct {
	// The transaction identifier of this payment.
	PaymentTransactionId string `json:"paymentTransactionId"`
	// The transaction mode of this payment.
	PaymentMode string `json:"paymentMode"`
	// Date timestamp
	PaymentDate time.Time `json:"paymentDate"`
}

type _PaymentInformation PaymentInformation

// NewPaymentInformation instantiates a new PaymentInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInformation(paymentTransactionId string, paymentMode string, paymentDate time.Time) *PaymentInformation {
	this := PaymentInformation{}
	this.PaymentTransactionId = paymentTransactionId
	this.PaymentMode = paymentMode
	this.PaymentDate = paymentDate
	return &this
}

// NewPaymentInformationWithDefaults instantiates a new PaymentInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInformationWithDefaults() *PaymentInformation {
	this := PaymentInformation{}
	return &this
}

// GetPaymentTransactionId returns the PaymentTransactionId field value
func (o *PaymentInformation) GetPaymentTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentTransactionId
}

// GetPaymentTransactionIdOk returns a tuple with the PaymentTransactionId field value
// and a boolean to check if the value has been set.
func (o *PaymentInformation) GetPaymentTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentTransactionId, true
}

// SetPaymentTransactionId sets field value
func (o *PaymentInformation) SetPaymentTransactionId(v string) {
	o.PaymentTransactionId = v
}

// GetPaymentMode returns the PaymentMode field value
func (o *PaymentInformation) GetPaymentMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMode
}

// GetPaymentModeOk returns a tuple with the PaymentMode field value
// and a boolean to check if the value has been set.
func (o *PaymentInformation) GetPaymentModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMode, true
}

// SetPaymentMode sets field value
func (o *PaymentInformation) SetPaymentMode(v string) {
	o.PaymentMode = v
}

// GetPaymentDate returns the PaymentDate field value
func (o *PaymentInformation) GetPaymentDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PaymentDate
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value
// and a boolean to check if the value has been set.
func (o *PaymentInformation) GetPaymentDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentDate, true
}

// SetPaymentDate sets field value
func (o *PaymentInformation) SetPaymentDate(v time.Time) {
	o.PaymentDate = v
}

func (o PaymentInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paymentTransactionId"] = o.PaymentTransactionId
	toSerialize["paymentMode"] = o.PaymentMode
	toSerialize["paymentDate"] = o.PaymentDate
	return toSerialize, nil
}

type NullablePaymentInformation struct {
	value *PaymentInformation
	isSet bool
}

func (v NullablePaymentInformation) Get() *PaymentInformation {
	return v.value
}

func (v *NullablePaymentInformation) Set(val *PaymentInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInformation(val *PaymentInformation) *NullablePaymentInformation {
	return &NullablePaymentInformation{value: val, isSet: true}
}

func (v NullablePaymentInformation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePaymentInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
