package shipment_invoicing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Money type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Money{}

// Money The currency type and amount.
type Money struct {
	// Three-digit currency code in ISO 4217 format.
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
	// The currency amount.
	Amount *string `json:"Amount,omitempty"`
}

// NewMoney instantiates a new Money object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoney() *Money {
	this := Money{}
	return &this
}

// NewMoneyWithDefaults instantiates a new Money object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyWithDefaults() *Money {
	this := Money{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *Money) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Money) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *Money) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *Money) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Money) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Money) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Money) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *Money) SetAmount(v string) {
	o.Amount = &v
}

func (o Money) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyCode) {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.Amount) {
		toSerialize["Amount"] = o.Amount
	}
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
