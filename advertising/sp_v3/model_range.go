package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the Range type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Range{}

// Range struct for Range
type Range struct {
	Min *float64 `json:"min,omitempty"`
	Max *float64 `json:"max,omitempty"`
}

// NewRange instantiates a new Range object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRange() *Range {
	this := Range{}
	return &this
}

// NewRangeWithDefaults instantiates a new Range object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeWithDefaults() *Range {
	this := Range{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Range) GetMin() float64 {
	if o == nil || IsNil(o.Min) {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetMinOk() (*float64, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Range) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *Range) SetMin(v float64) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Range) GetMax() float64 {
	if o == nil || IsNil(o.Max) {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetMaxOk() (*float64, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Range) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *Range) SetMax(v float64) {
	o.Max = &v
}

func (o Range) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableRange struct {
	value *Range
	isSet bool
}

func (v NullableRange) Get() *Range {
	return v.value
}

func (v *NullableRange) Set(val *Range) {
	v.value = val
	v.isSet = true
}

func (v NullableRange) IsSet() bool {
	return v.isSet
}

func (v *NullableRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRange(val *Range) *NullableRange {
	return &NullableRange{value: val, isSet: true}
}

func (v NullableRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
