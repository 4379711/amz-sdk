package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
)

// checks if the PaymentPreference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentPreference{}

// PaymentPreference The type of payment preference in which the transfer is being scheduled.
type PaymentPreference struct {
	PaymentPreferencePaymentType PaymentPreferencePaymentType `json:"paymentPreferencePaymentType"`
	// A decimal number, such as an amount or FX rate.
	Value float32 `json:"value"`
}

type _PaymentPreference PaymentPreference

// NewPaymentPreference instantiates a new PaymentPreference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentPreference(paymentPreferencePaymentType PaymentPreferencePaymentType, value float32) *PaymentPreference {
	this := PaymentPreference{}
	this.PaymentPreferencePaymentType = paymentPreferencePaymentType
	this.Value = value
	return &this
}

// NewPaymentPreferenceWithDefaults instantiates a new PaymentPreference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentPreferenceWithDefaults() *PaymentPreference {
	this := PaymentPreference{}
	return &this
}

// GetPaymentPreferencePaymentType returns the PaymentPreferencePaymentType field value
func (o *PaymentPreference) GetPaymentPreferencePaymentType() PaymentPreferencePaymentType {
	if o == nil {
		var ret PaymentPreferencePaymentType
		return ret
	}

	return o.PaymentPreferencePaymentType
}

// GetPaymentPreferencePaymentTypeOk returns a tuple with the PaymentPreferencePaymentType field value
// and a boolean to check if the value has been set.
func (o *PaymentPreference) GetPaymentPreferencePaymentTypeOk() (*PaymentPreferencePaymentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentPreferencePaymentType, true
}

// SetPaymentPreferencePaymentType sets field value
func (o *PaymentPreference) SetPaymentPreferencePaymentType(v PaymentPreferencePaymentType) {
	o.PaymentPreferencePaymentType = v
}

// GetValue returns the Value field value
func (o *PaymentPreference) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PaymentPreference) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PaymentPreference) SetValue(v float32) {
	o.Value = v
}

func (o PaymentPreference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paymentPreferencePaymentType"] = o.PaymentPreferencePaymentType
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullablePaymentPreference struct {
	value *PaymentPreference
	isSet bool
}

func (v NullablePaymentPreference) Get() *PaymentPreference {
	return v.value
}

func (v *NullablePaymentPreference) Set(val *PaymentPreference) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentPreference) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentPreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentPreference(val *PaymentPreference) *NullablePaymentPreference {
	return &NullablePaymentPreference{value: val, isSet: true}
}

func (v NullablePaymentPreference) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePaymentPreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
