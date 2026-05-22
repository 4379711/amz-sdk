package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the PerformanceMetricValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceMetricValue{}

// PerformanceMetricValue An object giving the name of the performance metric and its value when the rule was evaluated
type PerformanceMetricValue struct {
	// Name of the performance metric
	Name *string `json:"name,omitempty"`
	// Value of the performance metric
	Value *float64 `json:"value,omitempty"`
}

// NewPerformanceMetricValue instantiates a new PerformanceMetricValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceMetricValue() *PerformanceMetricValue {
	this := PerformanceMetricValue{}
	return &this
}

// NewPerformanceMetricValueWithDefaults instantiates a new PerformanceMetricValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceMetricValueWithDefaults() *PerformanceMetricValue {
	this := PerformanceMetricValue{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PerformanceMetricValue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceMetricValue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PerformanceMetricValue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PerformanceMetricValue) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PerformanceMetricValue) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceMetricValue) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PerformanceMetricValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *PerformanceMetricValue) SetValue(v float64) {
	o.Value = &v
}

func (o PerformanceMetricValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePerformanceMetricValue struct {
	value *PerformanceMetricValue
	isSet bool
}

func (v NullablePerformanceMetricValue) Get() *PerformanceMetricValue {
	return v.value
}

func (v *NullablePerformanceMetricValue) Set(val *PerformanceMetricValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceMetricValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceMetricValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceMetricValue(val *PerformanceMetricValue) *NullablePerformanceMetricValue {
	return &NullablePerformanceMetricValue{value: val, isSet: true}
}

func (v NullablePerformanceMetricValue) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePerformanceMetricValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
