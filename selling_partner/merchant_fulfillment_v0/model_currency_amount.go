package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the CurrencyAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrencyAmount{}

// CurrencyAmount Currency type and amount.
type CurrencyAmount struct {
	// Three-digit currency code in ISO 4217 format.
	CurrencyCode string `json:"CurrencyCode"`
	// The currency amount.
	Amount float64 `json:"Amount"`
}

type _CurrencyAmount CurrencyAmount

// NewCurrencyAmount instantiates a new CurrencyAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencyAmount(currencyCode string, amount float64) *CurrencyAmount {
	this := CurrencyAmount{}
	this.CurrencyCode = currencyCode
	this.Amount = amount
	return &this
}

// NewCurrencyAmountWithDefaults instantiates a new CurrencyAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyAmountWithDefaults() *CurrencyAmount {
	this := CurrencyAmount{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *CurrencyAmount) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *CurrencyAmount) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *CurrencyAmount) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetAmount returns the Amount field value
func (o *CurrencyAmount) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CurrencyAmount) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CurrencyAmount) SetAmount(v float64) {
	o.Amount = v
}

func (o CurrencyAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CurrencyCode"] = o.CurrencyCode
	toSerialize["Amount"] = o.Amount
	return toSerialize, nil
}

type NullableCurrencyAmount struct {
	value *CurrencyAmount
	isSet bool
}

func (v NullableCurrencyAmount) Get() *CurrencyAmount {
	return v.value
}

func (v *NullableCurrencyAmount) Set(val *CurrencyAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyAmount(val *CurrencyAmount) *NullableCurrencyAmount {
	return &NullableCurrencyAmount{value: val, isSet: true}
}

func (v NullableCurrencyAmount) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCurrencyAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
