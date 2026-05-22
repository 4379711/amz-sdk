package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingMetric{}

// SBForecastingMetric The forecast metric.
type SBForecastingMetric struct {
	// The forecast metric name. Currently supported metrics are IMPRESSION and CLICK.
	Metric *string                   `json:"metric,omitempty"`
	Value  *SBForecastingMetricValue `json:"value,omitempty"`
}

// NewSBForecastingMetric instantiates a new SBForecastingMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingMetric() *SBForecastingMetric {
	this := SBForecastingMetric{}
	return &this
}

// NewSBForecastingMetricWithDefaults instantiates a new SBForecastingMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingMetricWithDefaults() *SBForecastingMetric {
	this := SBForecastingMetric{}
	return &this
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *SBForecastingMetric) GetMetric() string {
	if o == nil || IsNil(o.Metric) {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingMetric) GetMetricOk() (*string, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *SBForecastingMetric) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *SBForecastingMetric) SetMetric(v string) {
	o.Metric = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SBForecastingMetric) GetValue() SBForecastingMetricValue {
	if o == nil || IsNil(o.Value) {
		var ret SBForecastingMetricValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingMetric) GetValueOk() (*SBForecastingMetricValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SBForecastingMetric) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given SBForecastingMetricValue and assigns it to the Value field.
func (o *SBForecastingMetric) SetValue(v SBForecastingMetricValue) {
	o.Value = &v
}

func (o SBForecastingMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSBForecastingMetric struct {
	value *SBForecastingMetric
	isSet bool
}

func (v NullableSBForecastingMetric) Get() *SBForecastingMetric {
	return v.value
}

func (v *NullableSBForecastingMetric) Set(val *SBForecastingMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingMetric(val *SBForecastingMetric) *NullableSBForecastingMetric {
	return &NullableSBForecastingMetric{value: val, isSet: true}
}

func (v NullableSBForecastingMetric) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
