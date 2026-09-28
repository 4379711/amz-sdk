package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the Discount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Discount{}

// Discount The discount configuration shared across benefit types.
type Discount struct {
	// The method used to calculate the discount amount.
	Type string `json:"type"`
	// The percentage discount value (1-100). Only valid when `type` is `PERCENTAGE_OFF`.
	PercentOff *float32 `json:"percentOff,omitempty"`
	// The amount off value. Only valid when `type` is `AMOUNT_OFF`.
	AmountOff *float32 `json:"amountOff,omitempty"`
	// The currency code in ISO 4217 format. Required when `type` is `AMOUNT_OFF`.
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// NewDiscount instantiates a new Discount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscount(type_ string) *Discount {
	this := Discount{}
	this.Type = type_
	return &this
}

// NewDiscountWithDefaults instantiates a new Discount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountWithDefaults() *Discount {
	this := Discount{}
	return &this
}

// GetType returns the Type field value
func (o *Discount) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Discount) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Discount) SetType(v string) {
	o.Type = v
}

// GetPercentOff returns the PercentOff field value if set, zero value otherwise.
func (o *Discount) GetPercentOff() float32 {
	if o == nil || IsNil(o.PercentOff) {
		var ret float32
		return ret
	}
	return *o.PercentOff
}

// GetPercentOffOk returns a tuple with the PercentOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetPercentOffOk() (*float32, bool) {
	if o == nil || IsNil(o.PercentOff) {
		return nil, false
	}
	return o.PercentOff, true
}

// HasPercentOff returns a boolean if a field has been set.
func (o *Discount) HasPercentOff() bool {
	if o != nil && !IsNil(o.PercentOff) {
		return true
	}

	return false
}

// SetPercentOff gets a reference to the given float32 and assigns it to the PercentOff field.
func (o *Discount) SetPercentOff(v float32) {
	o.PercentOff = &v
}

// GetAmountOff returns the AmountOff field value if set, zero value otherwise.
func (o *Discount) GetAmountOff() float32 {
	if o == nil || IsNil(o.AmountOff) {
		var ret float32
		return ret
	}
	return *o.AmountOff
}

// GetAmountOffOk returns a tuple with the AmountOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetAmountOffOk() (*float32, bool) {
	if o == nil || IsNil(o.AmountOff) {
		return nil, false
	}
	return o.AmountOff, true
}

// HasAmountOff returns a boolean if a field has been set.
func (o *Discount) HasAmountOff() bool {
	if o != nil && !IsNil(o.AmountOff) {
		return true
	}

	return false
}

// SetAmountOff gets a reference to the given float32 and assigns it to the AmountOff field.
func (o *Discount) SetAmountOff(v float32) {
	o.AmountOff = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *Discount) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *Discount) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *Discount) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

func (o Discount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.PercentOff) {
		toSerialize["percentOff"] = o.PercentOff
	}
	if !IsNil(o.AmountOff) {
		toSerialize["amountOff"] = o.AmountOff
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	return toSerialize, nil
}

type NullableDiscount struct {
	value *Discount
	isSet bool
}

func (v NullableDiscount) Get() *Discount {
	return v.value
}

func (v *NullableDiscount) Set(val *Discount) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscount(val *Discount) *NullableDiscount {
	return &NullableDiscount{value: val, isSet: true}
}

func (v NullableDiscount) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
