package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the BasisPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasisPrice{}

// BasisPrice The basis price before the savings are calculated
type BasisPrice struct {
	// Price amount
	Amount *float32 `json:"amount,omitempty"`
	// Currency of the price
	Currency *string `json:"currency,omitempty"`
}

// NewBasisPrice instantiates a new BasisPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasisPrice() *BasisPrice {
	this := BasisPrice{}
	return &this
}

// NewBasisPriceWithDefaults instantiates a new BasisPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasisPriceWithDefaults() *BasisPrice {
	this := BasisPrice{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BasisPrice) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasisPrice) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BasisPrice) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *BasisPrice) SetAmount(v float32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BasisPrice) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasisPrice) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BasisPrice) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BasisPrice) SetCurrency(v string) {
	o.Currency = &v
}

func (o BasisPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableBasisPrice struct {
	value *BasisPrice
	isSet bool
}

func (v NullableBasisPrice) Get() *BasisPrice {
	return v.value
}

func (v *NullableBasisPrice) Set(val *BasisPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableBasisPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableBasisPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasisPrice(val *BasisPrice) *NullableBasisPrice {
	return &NullableBasisPrice{value: val, isSet: true}
}

func (v NullableBasisPrice) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBasisPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
