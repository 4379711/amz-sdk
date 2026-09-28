package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the AmountThreshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmountThreshold{}

// AmountThreshold Amount-based purchase conditions for basket building promotions. Customers must spend a minimum dollar amount.
type AmountThreshold struct {
	// The type of amount requirement. Only `AT_LEAST` is supported for spend-based conditions.
	Type     string   `json:"type"`
	Currency Currency `json:"currency"`
}

// NewAmountThreshold instantiates a new AmountThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmountThreshold(type_ string, currency Currency) *AmountThreshold {
	this := AmountThreshold{}
	this.Type = type_
	this.Currency = currency
	return &this
}

// NewAmountThresholdWithDefaults instantiates a new AmountThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmountThresholdWithDefaults() *AmountThreshold {
	this := AmountThreshold{}
	return &this
}

// GetType returns the Type field value
func (o *AmountThreshold) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AmountThreshold) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AmountThreshold) SetType(v string) {
	o.Type = v
}

// GetCurrency returns the Currency field value
func (o *AmountThreshold) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *AmountThreshold) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *AmountThreshold) SetCurrency(v Currency) {
	o.Currency = v
}

func (o AmountThreshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

type NullableAmountThreshold struct {
	value *AmountThreshold
	isSet bool
}

func (v NullableAmountThreshold) Get() *AmountThreshold {
	return v.value
}

func (v *NullableAmountThreshold) Set(val *AmountThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableAmountThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableAmountThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmountThreshold(val *AmountThreshold) *NullableAmountThreshold {
	return &NullableAmountThreshold{value: val, isSet: true}
}

func (v NullableAmountThreshold) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAmountThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
