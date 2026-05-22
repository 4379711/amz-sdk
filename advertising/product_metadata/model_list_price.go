package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the ListPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPrice{}

// ListPrice struct for ListPrice
type ListPrice struct {
	Min           *float32 `json:"min,omitempty"`
	DisplayString *string  `json:"displayString,omitempty"`
	Max           *float32 `json:"max,omitempty"`
	ValueType     *string  `json:"valueType,omitempty"`
	// Currency of the price
	Currency *string `json:"currency,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// NewListPrice instantiates a new ListPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPrice() *ListPrice {
	this := ListPrice{}
	return &this
}

// NewListPriceWithDefaults instantiates a new ListPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPriceWithDefaults() *ListPrice {
	this := ListPrice{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ListPrice) GetMin() float32 {
	if o == nil || IsNil(o.Min) {
		var ret float32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPrice) GetMinOk() (*float32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ListPrice) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float32 and assigns it to the Min field.
func (o *ListPrice) SetMin(v float32) {
	o.Min = &v
}

// GetDisplayString returns the DisplayString field value if set, zero value otherwise.
func (o *ListPrice) GetDisplayString() string {
	if o == nil || IsNil(o.DisplayString) {
		var ret string
		return ret
	}
	return *o.DisplayString
}

// GetDisplayStringOk returns a tuple with the DisplayString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPrice) GetDisplayStringOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayString) {
		return nil, false
	}
	return o.DisplayString, true
}

// HasDisplayString returns a boolean if a field has been set.
func (o *ListPrice) HasDisplayString() bool {
	if o != nil && !IsNil(o.DisplayString) {
		return true
	}

	return false
}

// SetDisplayString gets a reference to the given string and assigns it to the DisplayString field.
func (o *ListPrice) SetDisplayString(v string) {
	o.DisplayString = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ListPrice) GetMax() float32 {
	if o == nil || IsNil(o.Max) {
		var ret float32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPrice) GetMaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ListPrice) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float32 and assigns it to the Max field.
func (o *ListPrice) SetMax(v float32) {
	o.Max = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *ListPrice) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPrice) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *ListPrice) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *ListPrice) SetValueType(v string) {
	o.ValueType = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ListPrice) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPrice) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ListPrice) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ListPrice) SetCurrency(v string) {
	o.Currency = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListPrice) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPrice) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListPrice) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListPrice) SetType(v string) {
	o.Type = &v
}

func (o ListPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.DisplayString) {
		toSerialize["displayString"] = o.DisplayString
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableListPrice struct {
	value *ListPrice
	isSet bool
}

func (v NullableListPrice) Get() *ListPrice {
	return v.value
}

func (v *NullableListPrice) Set(val *ListPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableListPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableListPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPrice(val *ListPrice) *NullableListPrice {
	return &NullableListPrice{value: val, isSet: true}
}

func (v NullableListPrice) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
