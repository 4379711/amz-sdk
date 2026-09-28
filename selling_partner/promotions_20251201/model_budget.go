package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the Budget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Budget{}

// Budget Budget configuration for spending limits. For `COUPON` and `BASKET_BUILDING` promotion types, the budget is set at the promotion level and applies to the entire promotion. For `DEAL` and `PRICE_DISCOUNT` promotion types, the budget is set at the item level within each item in the selection.
type Budget struct {
	// The budget type, either monetary or unit-based.
	Type string `json:"type"`
	// Budget value (amount or unit count).
	Value float32 `json:"value"`
	// The currency code in ISO 4217 format. Required when `type` is `AMOUNT`.
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// NewBudget instantiates a new Budget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudget(type_ string, value float32) *Budget {
	this := Budget{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewBudgetWithDefaults instantiates a new Budget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetWithDefaults() *Budget {
	this := Budget{}
	return &this
}

// GetType returns the Type field value
func (o *Budget) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Budget) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Budget) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *Budget) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Budget) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Budget) SetValue(v float32) {
	o.Value = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *Budget) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *Budget) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *Budget) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

func (o Budget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	return toSerialize, nil
}

type NullableBudget struct {
	value *Budget
	isSet bool
}

func (v NullableBudget) Get() *Budget {
	return v.value
}

func (v *NullableBudget) Set(val *Budget) {
	v.value = val
	v.isSet = true
}

func (v NullableBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudget(val *Budget) *NullableBudget {
	return &NullableBudget{value: val, isSet: true}
}

func (v NullableBudget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
