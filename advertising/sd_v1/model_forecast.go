package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the Forecast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Forecast{}

// Forecast Forecast impressions, clicks, reach, or conversions.
type Forecast struct {
	// Describes which metric is forecasted. |Name|Description| |-----------|------------------------| |IMPRESSIONS| Available impressions| |REACH      | Delivered viewable impressions| |CLICKS     | Delivered page visits| |CONVERSIONS| [Preview only] Delivered conversions|
	Metric *string        `json:"metric,omitempty"`
	Value  *ForecastRange `json:"value,omitempty"`
}

// NewForecast instantiates a new Forecast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForecast() *Forecast {
	this := Forecast{}
	return &this
}

// NewForecastWithDefaults instantiates a new Forecast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForecastWithDefaults() *Forecast {
	this := Forecast{}
	return &this
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *Forecast) GetMetric() string {
	if o == nil || IsNil(o.Metric) {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forecast) GetMetricOk() (*string, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *Forecast) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *Forecast) SetMetric(v string) {
	o.Metric = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Forecast) GetValue() ForecastRange {
	if o == nil || IsNil(o.Value) {
		var ret ForecastRange
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forecast) GetValueOk() (*ForecastRange, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Forecast) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ForecastRange and assigns it to the Value field.
func (o *Forecast) SetValue(v ForecastRange) {
	o.Value = &v
}

func (o Forecast) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableForecast struct {
	value *Forecast
	isSet bool
}

func (v NullableForecast) Get() *Forecast {
	return v.value
}

func (v *NullableForecast) Set(val *Forecast) {
	v.value = val
	v.isSet = true
}

func (v NullableForecast) IsSet() bool {
	return v.isSet
}

func (v *NullableForecast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecast(val *Forecast) *NullableForecast {
	return &NullableForecast{value: val, isSet: true}
}

func (v NullableForecast) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableForecast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
