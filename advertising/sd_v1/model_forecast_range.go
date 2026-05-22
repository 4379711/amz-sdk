package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ForecastRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForecastRange{}

// ForecastRange Forecast range values.
type ForecastRange struct {
	Min *int64 `json:"min,omitempty"`
	Max *int64 `json:"max,omitempty"`
}

// NewForecastRange instantiates a new ForecastRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForecastRange() *ForecastRange {
	this := ForecastRange{}
	return &this
}

// NewForecastRangeWithDefaults instantiates a new ForecastRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForecastRangeWithDefaults() *ForecastRange {
	this := ForecastRange{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ForecastRange) GetMin() int64 {
	if o == nil || IsNil(o.Min) {
		var ret int64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastRange) GetMinOk() (*int64, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ForecastRange) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int64 and assigns it to the Min field.
func (o *ForecastRange) SetMin(v int64) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ForecastRange) GetMax() int64 {
	if o == nil || IsNil(o.Max) {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastRange) GetMaxOk() (*int64, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ForecastRange) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *ForecastRange) SetMax(v int64) {
	o.Max = &v
}

func (o ForecastRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableForecastRange struct {
	value *ForecastRange
	isSet bool
}

func (v NullableForecastRange) Get() *ForecastRange {
	return v.value
}

func (v *NullableForecastRange) Set(val *ForecastRange) {
	v.value = val
	v.isSet = true
}

func (v NullableForecastRange) IsSet() bool {
	return v.isSet
}

func (v *NullableForecastRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecastRange(val *ForecastRange) *NullableForecastRange {
	return &NullableForecastRange{value: val, isSet: true}
}

func (v NullableForecastRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableForecastRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
