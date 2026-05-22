package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RangeMetricValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeMetricValue{}

// RangeMetricValue Describes lower and upper bounds of the range. <br> Note: This object is nullable
type RangeMetricValue struct {
	Lower *int32 `json:"lower,omitempty"`
	Upper *int32 `json:"upper,omitempty"`
}

// NewRangeMetricValue instantiates a new RangeMetricValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeMetricValue() *RangeMetricValue {
	this := RangeMetricValue{}
	return &this
}

// NewRangeMetricValueWithDefaults instantiates a new RangeMetricValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeMetricValueWithDefaults() *RangeMetricValue {
	this := RangeMetricValue{}
	return &this
}

// GetLower returns the Lower field value if set, zero value otherwise.
func (o *RangeMetricValue) GetLower() int32 {
	if o == nil || IsNil(o.Lower) {
		var ret int32
		return ret
	}
	return *o.Lower
}

// GetLowerOk returns a tuple with the Lower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeMetricValue) GetLowerOk() (*int32, bool) {
	if o == nil || IsNil(o.Lower) {
		return nil, false
	}
	return o.Lower, true
}

// HasLower returns a boolean if a field has been set.
func (o *RangeMetricValue) HasLower() bool {
	if o != nil && !IsNil(o.Lower) {
		return true
	}

	return false
}

// SetLower gets a reference to the given int32 and assigns it to the Lower field.
func (o *RangeMetricValue) SetLower(v int32) {
	o.Lower = &v
}

// GetUpper returns the Upper field value if set, zero value otherwise.
func (o *RangeMetricValue) GetUpper() int32 {
	if o == nil || IsNil(o.Upper) {
		var ret int32
		return ret
	}
	return *o.Upper
}

// GetUpperOk returns a tuple with the Upper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeMetricValue) GetUpperOk() (*int32, bool) {
	if o == nil || IsNil(o.Upper) {
		return nil, false
	}
	return o.Upper, true
}

// HasUpper returns a boolean if a field has been set.
func (o *RangeMetricValue) HasUpper() bool {
	if o != nil && !IsNil(o.Upper) {
		return true
	}

	return false
}

// SetUpper gets a reference to the given int32 and assigns it to the Upper field.
func (o *RangeMetricValue) SetUpper(v int32) {
	o.Upper = &v
}

func (o RangeMetricValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lower) {
		toSerialize["lower"] = o.Lower
	}
	if !IsNil(o.Upper) {
		toSerialize["upper"] = o.Upper
	}
	return toSerialize, nil
}

type NullableRangeMetricValue struct {
	value *RangeMetricValue
	isSet bool
}

func (v NullableRangeMetricValue) Get() *RangeMetricValue {
	return v.value
}

func (v *NullableRangeMetricValue) Set(val *RangeMetricValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeMetricValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeMetricValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeMetricValue(val *RangeMetricValue) *NullableRangeMetricValue {
	return &NullableRangeMetricValue{value: val, isSet: true}
}

func (v NullableRangeMetricValue) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRangeMetricValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
