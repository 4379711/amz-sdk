package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the IntegerRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegerRange{}

// IntegerRange struct for IntegerRange
type IntegerRange struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewIntegerRange instantiates a new IntegerRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegerRange() *IntegerRange {
	this := IntegerRange{}
	return &this
}

// NewIntegerRangeWithDefaults instantiates a new IntegerRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegerRangeWithDefaults() *IntegerRange {
	this := IntegerRange{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *IntegerRange) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerRange) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *IntegerRange) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *IntegerRange) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *IntegerRange) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerRange) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *IntegerRange) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *IntegerRange) SetMax(v int32) {
	o.Max = &v
}

func (o IntegerRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableIntegerRange struct {
	value *IntegerRange
	isSet bool
}

func (v NullableIntegerRange) Get() *IntegerRange {
	return v.value
}

func (v *NullableIntegerRange) Set(val *IntegerRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegerRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegerRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegerRange(val *IntegerRange) *NullableIntegerRange {
	return &NullableIntegerRange{value: val, isSet: true}
}

func (v NullableIntegerRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIntegerRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
