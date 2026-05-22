package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingPriceRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingPriceRange{}

// SBTargetingPriceRange A range of prices. We use this to retrieve the number of targetable ASINs that falls within this price range.
type SBTargetingPriceRange struct {
	Min *float64 `json:"min,omitempty"`
	Max *float64 `json:"max,omitempty"`
}

// NewSBTargetingPriceRange instantiates a new SBTargetingPriceRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingPriceRange() *SBTargetingPriceRange {
	this := SBTargetingPriceRange{}
	return &this
}

// NewSBTargetingPriceRangeWithDefaults instantiates a new SBTargetingPriceRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingPriceRangeWithDefaults() *SBTargetingPriceRange {
	this := SBTargetingPriceRange{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *SBTargetingPriceRange) GetMin() float64 {
	if o == nil || IsNil(o.Min) {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingPriceRange) GetMinOk() (*float64, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *SBTargetingPriceRange) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *SBTargetingPriceRange) SetMin(v float64) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *SBTargetingPriceRange) GetMax() float64 {
	if o == nil || IsNil(o.Max) {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingPriceRange) GetMaxOk() (*float64, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *SBTargetingPriceRange) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *SBTargetingPriceRange) SetMax(v float64) {
	o.Max = &v
}

func (o SBTargetingPriceRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableSBTargetingPriceRange struct {
	value *SBTargetingPriceRange
	isSet bool
}

func (v NullableSBTargetingPriceRange) Get() *SBTargetingPriceRange {
	return v.value
}

func (v *NullableSBTargetingPriceRange) Set(val *SBTargetingPriceRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingPriceRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingPriceRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingPriceRange(val *SBTargetingPriceRange) *NullableSBTargetingPriceRange {
	return &NullableSBTargetingPriceRange{value: val, isSet: true}
}

func (v NullableSBTargetingPriceRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingPriceRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
