package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ImpactMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImpactMetric{}

// ImpactMetric The impact metrics are given in the same order of suggested bids. <br> Note: This object is nullable
type ImpactMetric struct {
	Values []RangeMetricValue `json:"values,omitempty"`
}

// NewImpactMetric instantiates a new ImpactMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImpactMetric() *ImpactMetric {
	this := ImpactMetric{}
	return &this
}

// NewImpactMetricWithDefaults instantiates a new ImpactMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImpactMetricWithDefaults() *ImpactMetric {
	this := ImpactMetric{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ImpactMetric) GetValues() []RangeMetricValue {
	if o == nil || IsNil(o.Values) {
		var ret []RangeMetricValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImpactMetric) GetValuesOk() ([]RangeMetricValue, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ImpactMetric) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []RangeMetricValue and assigns it to the Values field.
func (o *ImpactMetric) SetValues(v []RangeMetricValue) {
	o.Values = v
}

func (o ImpactMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableImpactMetric struct {
	value *ImpactMetric
	isSet bool
}

func (v NullableImpactMetric) Get() *ImpactMetric {
	return v.value
}

func (v *NullableImpactMetric) Set(val *ImpactMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableImpactMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableImpactMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImpactMetric(val *ImpactMetric) *NullableImpactMetric {
	return &NullableImpactMetric{value: val, isSet: true}
}

func (v NullableImpactMetric) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImpactMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
