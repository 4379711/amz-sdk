package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingMetricValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingMetricValue{}

// SBForecastingMetricValue The forecast min and max value.
type SBForecastingMetricValue struct {
	// The forecast min value.
	Min *float32 `json:"min,omitempty"`
	// The forecast max value.
	Max *float32 `json:"max,omitempty"`
}

// NewSBForecastingMetricValue instantiates a new SBForecastingMetricValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingMetricValue() *SBForecastingMetricValue {
	this := SBForecastingMetricValue{}
	return &this
}

// NewSBForecastingMetricValueWithDefaults instantiates a new SBForecastingMetricValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingMetricValueWithDefaults() *SBForecastingMetricValue {
	this := SBForecastingMetricValue{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *SBForecastingMetricValue) GetMin() float32 {
	if o == nil || IsNil(o.Min) {
		var ret float32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingMetricValue) GetMinOk() (*float32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *SBForecastingMetricValue) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float32 and assigns it to the Min field.
func (o *SBForecastingMetricValue) SetMin(v float32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *SBForecastingMetricValue) GetMax() float32 {
	if o == nil || IsNil(o.Max) {
		var ret float32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingMetricValue) GetMaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *SBForecastingMetricValue) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float32 and assigns it to the Max field.
func (o *SBForecastingMetricValue) SetMax(v float32) {
	o.Max = &v
}

func (o SBForecastingMetricValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableSBForecastingMetricValue struct {
	value *SBForecastingMetricValue
	isSet bool
}

func (v NullableSBForecastingMetricValue) Get() *SBForecastingMetricValue {
	return v.value
}

func (v *NullableSBForecastingMetricValue) Set(val *SBForecastingMetricValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingMetricValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingMetricValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingMetricValue(val *SBForecastingMetricValue) *NullableSBForecastingMetricValue {
	return &NullableSBForecastingMetricValue{value: val, isSet: true}
}

func (v NullableSBForecastingMetricValue) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingMetricValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
