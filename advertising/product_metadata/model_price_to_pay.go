package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the PriceToPay type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceToPay{}

// PriceToPay The price customer would pay for the buying option
type PriceToPay struct {
	// Price amount
	Amount *float32 `json:"amount,omitempty"`
	// Currency of the price
	Currency *string `json:"currency,omitempty"`
}

// NewPriceToPay instantiates a new PriceToPay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceToPay() *PriceToPay {
	this := PriceToPay{}
	return &this
}

// NewPriceToPayWithDefaults instantiates a new PriceToPay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceToPayWithDefaults() *PriceToPay {
	this := PriceToPay{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PriceToPay) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceToPay) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PriceToPay) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *PriceToPay) SetAmount(v float32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PriceToPay) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceToPay) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PriceToPay) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PriceToPay) SetCurrency(v string) {
	o.Currency = &v
}

func (o PriceToPay) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullablePriceToPay struct {
	value *PriceToPay
	isSet bool
}

func (v NullablePriceToPay) Get() *PriceToPay {
	return v.value
}

func (v *NullablePriceToPay) Set(val *PriceToPay) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceToPay) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceToPay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceToPay(val *PriceToPay) *NullablePriceToPay {
	return &NullablePriceToPay{value: val, isSet: true}
}

func (v NullablePriceToPay) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePriceToPay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
