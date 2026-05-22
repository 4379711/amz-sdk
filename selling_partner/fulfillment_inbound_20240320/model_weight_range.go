package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the WeightRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WeightRange{}

// WeightRange The range of weights that are allowed for a package.
type WeightRange struct {
	// Maximum allowed weight.
	Maximum float32 `json:"maximum"`
	// Minimum allowed weight.
	Minimum float32      `json:"minimum"`
	Unit    UnitOfWeight `json:"unit"`
}

type _WeightRange WeightRange

// NewWeightRange instantiates a new WeightRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeightRange(maximum float32, minimum float32, unit UnitOfWeight) *WeightRange {
	this := WeightRange{}
	this.Maximum = maximum
	this.Minimum = minimum
	this.Unit = unit
	return &this
}

// NewWeightRangeWithDefaults instantiates a new WeightRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeightRangeWithDefaults() *WeightRange {
	this := WeightRange{}
	return &this
}

// GetMaximum returns the Maximum field value
func (o *WeightRange) GetMaximum() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value
// and a boolean to check if the value has been set.
func (o *WeightRange) GetMaximumOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maximum, true
}

// SetMaximum sets field value
func (o *WeightRange) SetMaximum(v float32) {
	o.Maximum = v
}

// GetMinimum returns the Minimum field value
func (o *WeightRange) GetMinimum() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value
// and a boolean to check if the value has been set.
func (o *WeightRange) GetMinimumOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minimum, true
}

// SetMinimum sets field value
func (o *WeightRange) SetMinimum(v float32) {
	o.Minimum = v
}

// GetUnit returns the Unit field value
func (o *WeightRange) GetUnit() UnitOfWeight {
	if o == nil {
		var ret UnitOfWeight
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *WeightRange) GetUnitOk() (*UnitOfWeight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *WeightRange) SetUnit(v UnitOfWeight) {
	o.Unit = v
}

func (o WeightRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maximum"] = o.Maximum
	toSerialize["minimum"] = o.Minimum
	toSerialize["unit"] = o.Unit
	return toSerialize, nil
}

type NullableWeightRange struct {
	value *WeightRange
	isSet bool
}

func (v NullableWeightRange) Get() *WeightRange {
	return v.value
}

func (v *NullableWeightRange) Set(val *WeightRange) {
	v.value = val
	v.isSet = true
}

func (v NullableWeightRange) IsSet() bool {
	return v.isSet
}

func (v *NullableWeightRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeightRange(val *WeightRange) *NullableWeightRange {
	return &NullableWeightRange{value: val, isSet: true}
}

func (v NullableWeightRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableWeightRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
