package product_fees_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the MoneyType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoneyType{}

// MoneyType struct for MoneyType
type MoneyType struct {
	// The currency code in ISO 4217 format.
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
	// The monetary value.
	Amount *float32 `json:"Amount,omitempty"`
}

// NewMoneyType instantiates a new MoneyType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyType() *MoneyType {
	this := MoneyType{}
	return &this
}

// NewMoneyTypeWithDefaults instantiates a new MoneyType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyTypeWithDefaults() *MoneyType {
	this := MoneyType{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *MoneyType) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyType) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *MoneyType) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *MoneyType) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MoneyType) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyType) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MoneyType) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *MoneyType) SetAmount(v float32) {
	o.Amount = &v
}

func (o MoneyType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyCode) {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.Amount) {
		toSerialize["Amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableMoneyType struct {
	value *MoneyType
	isSet bool
}

func (v NullableMoneyType) Get() *MoneyType {
	return v.value
}

func (v *NullableMoneyType) Set(val *MoneyType) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyType) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyType(val *MoneyType) *NullableMoneyType {
	return &NullableMoneyType{value: val, isSet: true}
}

func (v NullableMoneyType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMoneyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
