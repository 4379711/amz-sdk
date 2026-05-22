package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingRatingRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingRatingRange{}

// SBTargetingRatingRange Rating range is restricted to integers between 0 and 5, inclusive. Min must be less than or equal to max. We use this to retrieve the number of targetable ASINs that falls within this rating range.
type SBTargetingRatingRange struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewSBTargetingRatingRange instantiates a new SBTargetingRatingRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingRatingRange() *SBTargetingRatingRange {
	this := SBTargetingRatingRange{}
	return &this
}

// NewSBTargetingRatingRangeWithDefaults instantiates a new SBTargetingRatingRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingRatingRangeWithDefaults() *SBTargetingRatingRange {
	this := SBTargetingRatingRange{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *SBTargetingRatingRange) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingRatingRange) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *SBTargetingRatingRange) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *SBTargetingRatingRange) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *SBTargetingRatingRange) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingRatingRange) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *SBTargetingRatingRange) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *SBTargetingRatingRange) SetMax(v int32) {
	o.Max = &v
}

func (o SBTargetingRatingRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableSBTargetingRatingRange struct {
	value *SBTargetingRatingRange
	isSet bool
}

func (v NullableSBTargetingRatingRange) Get() *SBTargetingRatingRange {
	return v.value
}

func (v *NullableSBTargetingRatingRange) Set(val *SBTargetingRatingRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingRatingRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingRatingRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingRatingRange(val *SBTargetingRatingRange) *NullableSBTargetingRatingRange {
	return &NullableSBTargetingRatingRange{value: val, isSet: true}
}

func (v NullableSBTargetingRatingRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingRatingRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
