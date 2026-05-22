package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingEstimatedReachRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingEstimatedReachRange{}

// SBTargetingEstimatedReachRange struct for SBTargetingEstimatedReachRange
type SBTargetingEstimatedReachRange struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewSBTargetingEstimatedReachRange instantiates a new SBTargetingEstimatedReachRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingEstimatedReachRange() *SBTargetingEstimatedReachRange {
	this := SBTargetingEstimatedReachRange{}
	return &this
}

// NewSBTargetingEstimatedReachRangeWithDefaults instantiates a new SBTargetingEstimatedReachRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingEstimatedReachRangeWithDefaults() *SBTargetingEstimatedReachRange {
	this := SBTargetingEstimatedReachRange{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *SBTargetingEstimatedReachRange) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingEstimatedReachRange) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *SBTargetingEstimatedReachRange) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *SBTargetingEstimatedReachRange) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *SBTargetingEstimatedReachRange) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingEstimatedReachRange) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *SBTargetingEstimatedReachRange) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *SBTargetingEstimatedReachRange) SetMax(v int32) {
	o.Max = &v
}

func (o SBTargetingEstimatedReachRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableSBTargetingEstimatedReachRange struct {
	value *SBTargetingEstimatedReachRange
	isSet bool
}

func (v NullableSBTargetingEstimatedReachRange) Get() *SBTargetingEstimatedReachRange {
	return v.value
}

func (v *NullableSBTargetingEstimatedReachRange) Set(val *SBTargetingEstimatedReachRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingEstimatedReachRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingEstimatedReachRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingEstimatedReachRange(val *SBTargetingEstimatedReachRange) *NullableSBTargetingEstimatedReachRange {
	return &NullableSBTargetingEstimatedReachRange{value: val, isSet: true}
}

func (v NullableSBTargetingEstimatedReachRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingEstimatedReachRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
