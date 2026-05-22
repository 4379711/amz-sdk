package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ForecastRangeDouble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForecastRangeDouble{}

// ForecastRangeDouble A range of value.
type ForecastRangeDouble struct {
	// Lower bound.
	Min interface{} `json:"min,omitempty"`
	// Geometric mean of the upper and lower bounds.
	Mean interface{} `json:"mean,omitempty"`
	// Upper bound.
	Max interface{} `json:"max,omitempty"`
}

// NewForecastRangeDouble instantiates a new ForecastRangeDouble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForecastRangeDouble() *ForecastRangeDouble {
	this := ForecastRangeDouble{}
	return &this
}

// NewForecastRangeDoubleWithDefaults instantiates a new ForecastRangeDouble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForecastRangeDoubleWithDefaults() *ForecastRangeDouble {
	this := ForecastRangeDouble{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastRangeDouble) GetMin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastRangeDouble) GetMinOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return &o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ForecastRangeDouble) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given interface{} and assigns it to the Min field.
func (o *ForecastRangeDouble) SetMin(v interface{}) {
	o.Min = v
}

// GetMean returns the Mean field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastRangeDouble) GetMean() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Mean
}

// GetMeanOk returns a tuple with the Mean field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastRangeDouble) GetMeanOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Mean) {
		return nil, false
	}
	return &o.Mean, true
}

// HasMean returns a boolean if a field has been set.
func (o *ForecastRangeDouble) HasMean() bool {
	if o != nil && !IsNil(o.Mean) {
		return true
	}

	return false
}

// SetMean gets a reference to the given interface{} and assigns it to the Mean field.
func (o *ForecastRangeDouble) SetMean(v interface{}) {
	o.Mean = v
}

// GetMax returns the Max field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastRangeDouble) GetMax() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastRangeDouble) GetMaxOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return &o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ForecastRangeDouble) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given interface{} and assigns it to the Max field.
func (o *ForecastRangeDouble) SetMax(v interface{}) {
	o.Max = v
}

func (o ForecastRangeDouble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Mean != nil {
		toSerialize["mean"] = o.Mean
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableForecastRangeDouble struct {
	value *ForecastRangeDouble
	isSet bool
}

func (v NullableForecastRangeDouble) Get() *ForecastRangeDouble {
	return v.value
}

func (v *NullableForecastRangeDouble) Set(val *ForecastRangeDouble) {
	v.value = val
	v.isSet = true
}

func (v NullableForecastRangeDouble) IsSet() bool {
	return v.isSet
}

func (v *NullableForecastRangeDouble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecastRangeDouble(val *ForecastRangeDouble) *NullableForecastRangeDouble {
	return &NullableForecastRangeDouble{value: val, isSet: true}
}

func (v NullableForecastRangeDouble) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableForecastRangeDouble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
