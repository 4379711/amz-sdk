package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the Money type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Money{}

// Money An amount of money, including units in the form of currency.
type Money struct {
	// Three digit currency code in ISO 4217 format.
	CurrencyCode string `json:"currencyCode"`
	// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation.
	Value string `json:"value"`
}

type _Money Money

// NewMoney instantiates a new Money object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoney(currencyCode string, value string) *Money {
	this := Money{}
	this.CurrencyCode = currencyCode
	this.Value = value
	return &this
}

// NewMoneyWithDefaults instantiates a new Money object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyWithDefaults() *Money {
	this := Money{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *Money) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *Money) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *Money) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetValue returns the Value field value
func (o *Money) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Money) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Money) SetValue(v string) {
	o.Value = v
}

func (o Money) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currencyCode"] = o.CurrencyCode
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableMoney struct {
	value *Money
	isSet bool
}

func (v NullableMoney) Get() *Money {
	return v.value
}

func (v *NullableMoney) Set(val *Money) {
	v.value = val
	v.isSet = true
}

func (v NullableMoney) IsSet() bool {
	return v.isSet
}

func (v *NullableMoney) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoney(val *Money) *NullableMoney {
	return &NullableMoney{value: val, isSet: true}
}

func (v NullableMoney) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMoney) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
