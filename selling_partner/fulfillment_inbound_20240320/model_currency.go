package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the Currency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Currency{}

// Currency The type and amount of currency.
type Currency struct {
	// Decimal value of the currency.
	Amount float32 `json:"amount"`
	// ISO 4217 standard of a currency code.
	Code string `json:"code"`
}

type _Currency Currency

// NewCurrency instantiates a new Currency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrency(amount float32, code string) *Currency {
	this := Currency{}
	this.Amount = amount
	this.Code = code
	return &this
}

// NewCurrencyWithDefaults instantiates a new Currency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyWithDefaults() *Currency {
	this := Currency{}
	return &this
}

// GetAmount returns the Amount field value
func (o *Currency) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Currency) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Currency) SetAmount(v float32) {
	o.Amount = v
}

// GetCode returns the Code field value
func (o *Currency) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Currency) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Currency) SetCode(v string) {
	o.Code = v
}

func (o Currency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

type NullableCurrency struct {
	value *Currency
	isSet bool
}

func (v NullableCurrency) Get() *Currency {
	return v.value
}

func (v *NullableCurrency) Set(val *Currency) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrency(val *Currency) *NullableCurrency {
	return &NullableCurrency{value: val, isSet: true}
}

func (v NullableCurrency) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
