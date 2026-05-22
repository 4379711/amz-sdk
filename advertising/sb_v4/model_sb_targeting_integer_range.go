package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingIntegerRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingIntegerRange{}

// SBTargetingIntegerRange struct for SBTargetingIntegerRange
type SBTargetingIntegerRange struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewSBTargetingIntegerRange instantiates a new SBTargetingIntegerRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingIntegerRange() *SBTargetingIntegerRange {
	this := SBTargetingIntegerRange{}
	return &this
}

// NewSBTargetingIntegerRangeWithDefaults instantiates a new SBTargetingIntegerRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingIntegerRangeWithDefaults() *SBTargetingIntegerRange {
	this := SBTargetingIntegerRange{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *SBTargetingIntegerRange) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingIntegerRange) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *SBTargetingIntegerRange) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *SBTargetingIntegerRange) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *SBTargetingIntegerRange) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingIntegerRange) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *SBTargetingIntegerRange) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *SBTargetingIntegerRange) SetMax(v int32) {
	o.Max = &v
}

func (o SBTargetingIntegerRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableSBTargetingIntegerRange struct {
	value *SBTargetingIntegerRange
	isSet bool
}

func (v NullableSBTargetingIntegerRange) Get() *SBTargetingIntegerRange {
	return v.value
}

func (v *NullableSBTargetingIntegerRange) Set(val *SBTargetingIntegerRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingIntegerRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingIntegerRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingIntegerRange(val *SBTargetingIntegerRange) *NullableSBTargetingIntegerRange {
	return &NullableSBTargetingIntegerRange{value: val, isSet: true}
}

func (v NullableSBTargetingIntegerRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingIntegerRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
