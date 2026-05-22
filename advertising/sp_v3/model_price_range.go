package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PriceRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceRange{}

// PriceRange A range of prices. We use this to retrieve the number of targetable ASINs that falls within this price range.
type PriceRange struct {
	Min *float64 `json:"min,omitempty"`
	Max *float64 `json:"max,omitempty"`
}

// NewPriceRange instantiates a new PriceRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceRange() *PriceRange {
	this := PriceRange{}
	return &this
}

// NewPriceRangeWithDefaults instantiates a new PriceRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceRangeWithDefaults() *PriceRange {
	this := PriceRange{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *PriceRange) GetMin() float64 {
	if o == nil || IsNil(o.Min) {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceRange) GetMinOk() (*float64, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *PriceRange) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *PriceRange) SetMin(v float64) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *PriceRange) GetMax() float64 {
	if o == nil || IsNil(o.Max) {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceRange) GetMaxOk() (*float64, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *PriceRange) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *PriceRange) SetMax(v float64) {
	o.Max = &v
}

func (o PriceRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullablePriceRange struct {
	value *PriceRange
	isSet bool
}

func (v NullablePriceRange) Get() *PriceRange {
	return v.value
}

func (v *NullablePriceRange) Set(val *PriceRange) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceRange) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceRange(val *PriceRange) *NullablePriceRange {
	return &NullablePriceRange{value: val, isSet: true}
}

func (v NullablePriceRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePriceRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
