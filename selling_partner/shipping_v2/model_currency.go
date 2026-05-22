package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the Currency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Currency{}

// Currency The monetary value in the currency indicated, in ISO 4217 standard format.
type Currency struct {
	// The monetary value.
	Value float32 `json:"value"`
	// The ISO 4217 format 3-character currency code.
	Unit string `json:"unit"`
}

type _Currency Currency

// NewCurrency instantiates a new Currency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrency(value float32, unit string) *Currency {
	this := Currency{}
	this.Value = value
	this.Unit = unit
	return &this
}

// NewCurrencyWithDefaults instantiates a new Currency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyWithDefaults() *Currency {
	this := Currency{}
	return &this
}

// GetValue returns the Value field value
func (o *Currency) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Currency) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Currency) SetValue(v float32) {
	o.Value = v
}

// GetUnit returns the Unit field value
func (o *Currency) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *Currency) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *Currency) SetUnit(v string) {
	o.Unit = v
}

func (o Currency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["unit"] = o.Unit
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
