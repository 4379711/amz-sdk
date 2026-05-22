package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the IntegerWithUnits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegerWithUnits{}

// IntegerWithUnits A whole number dimension and its unit of measurement. For example, this can represent 100 pixels.
type IntegerWithUnits struct {
	// The dimension value.
	Value int32 `json:"value"`
	// The unit of measurement.
	Units string `json:"units"`
}

type _IntegerWithUnits IntegerWithUnits

// NewIntegerWithUnits instantiates a new IntegerWithUnits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegerWithUnits(value int32, units string) *IntegerWithUnits {
	this := IntegerWithUnits{}
	this.Value = value
	this.Units = units
	return &this
}

// NewIntegerWithUnitsWithDefaults instantiates a new IntegerWithUnits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegerWithUnitsWithDefaults() *IntegerWithUnits {
	this := IntegerWithUnits{}
	return &this
}

// GetValue returns the Value field value
func (o *IntegerWithUnits) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IntegerWithUnits) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *IntegerWithUnits) SetValue(v int32) {
	o.Value = v
}

// GetUnits returns the Units field value
func (o *IntegerWithUnits) GetUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *IntegerWithUnits) GetUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *IntegerWithUnits) SetUnits(v string) {
	o.Units = v
}

func (o IntegerWithUnits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["units"] = o.Units
	return toSerialize, nil
}

type NullableIntegerWithUnits struct {
	value *IntegerWithUnits
	isSet bool
}

func (v NullableIntegerWithUnits) Get() *IntegerWithUnits {
	return v.value
}

func (v *NullableIntegerWithUnits) Set(val *IntegerWithUnits) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegerWithUnits) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegerWithUnits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegerWithUnits(val *IntegerWithUnits) *NullableIntegerWithUnits {
	return &NullableIntegerWithUnits{value: val, isSet: true}
}

func (v NullableIntegerWithUnits) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIntegerWithUnits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
