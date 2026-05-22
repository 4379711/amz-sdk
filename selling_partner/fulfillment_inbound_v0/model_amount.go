package fulfillment_inbound_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Amount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Amount{}

// Amount The monetary value.
type Amount struct {
	CurrencyCode CurrencyCode `json:"CurrencyCode"`
	// Number format that supports decimal.
	Value float64 `json:"Value"`
}

type _Amount Amount

// NewAmount instantiates a new Amount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmount(currencyCode CurrencyCode, value float64) *Amount {
	this := Amount{}
	this.CurrencyCode = currencyCode
	this.Value = value
	return &this
}

// NewAmountWithDefaults instantiates a new Amount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmountWithDefaults() *Amount {
	this := Amount{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *Amount) GetCurrencyCode() CurrencyCode {
	if o == nil {
		var ret CurrencyCode
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *Amount) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *Amount) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = v
}

// GetValue returns the Value field value
func (o *Amount) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Amount) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Amount) SetValue(v float64) {
	o.Value = v
}

func (o Amount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CurrencyCode"] = o.CurrencyCode
	toSerialize["Value"] = o.Value
	return toSerialize, nil
}

type NullableAmount struct {
	value *Amount
	isSet bool
}

func (v NullableAmount) Get() *Amount {
	return v.value
}

func (v *NullableAmount) Set(val *Amount) {
	v.value = val
	v.isSet = true
}

func (v NullableAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmount(val *Amount) *NullableAmount {
	return &NullableAmount{value: val, isSet: true}
}

func (v NullableAmount) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
