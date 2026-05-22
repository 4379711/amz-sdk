package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RatingRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatingRange{}

// RatingRange Rating range is restricted to integers between 0 and 5, inclusive. Min must be less than or equal to max. We use this to retrieve the number of targetable ASINs that falls within this rating range.
type RatingRange struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewRatingRange instantiates a new RatingRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatingRange() *RatingRange {
	this := RatingRange{}
	return &this
}

// NewRatingRangeWithDefaults instantiates a new RatingRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatingRangeWithDefaults() *RatingRange {
	this := RatingRange{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *RatingRange) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingRange) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *RatingRange) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *RatingRange) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *RatingRange) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingRange) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *RatingRange) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *RatingRange) SetMax(v int32) {
	o.Max = &v
}

func (o RatingRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableRatingRange struct {
	value *RatingRange
	isSet bool
}

func (v NullableRatingRange) Get() *RatingRange {
	return v.value
}

func (v *NullableRatingRange) Set(val *RatingRange) {
	v.value = val
	v.isSet = true
}

func (v NullableRatingRange) IsSet() bool {
	return v.isSet
}

func (v *NullableRatingRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatingRange(val *RatingRange) *NullableRatingRange {
	return &NullableRatingRange{value: val, isSet: true}
}

func (v NullableRatingRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRatingRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
