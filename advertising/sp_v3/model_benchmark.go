package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the Benchmark type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Benchmark{}

// Benchmark Forecasted impact metrics for next 7 days or during special days.
type Benchmark struct {
	// Specifies the processing status of the benchmark. Success - If all fields in values property (impressions, clicks, conversions) have all non-null values. Failed - If all fields in values property have all null values. Partial - If some of the fields (impressions, clicks, or conversions) in values property have null values.
	BenchmarkStatus *string `json:"benchmarkStatus,omitempty"`
	Values          *Values `json:"values,omitempty"`
}

// NewBenchmark instantiates a new Benchmark object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBenchmark() *Benchmark {
	this := Benchmark{}
	return &this
}

// NewBenchmarkWithDefaults instantiates a new Benchmark object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBenchmarkWithDefaults() *Benchmark {
	this := Benchmark{}
	return &this
}

// GetBenchmarkStatus returns the BenchmarkStatus field value if set, zero value otherwise.
func (o *Benchmark) GetBenchmarkStatus() string {
	if o == nil || IsNil(o.BenchmarkStatus) {
		var ret string
		return ret
	}
	return *o.BenchmarkStatus
}

// GetBenchmarkStatusOk returns a tuple with the BenchmarkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetBenchmarkStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BenchmarkStatus) {
		return nil, false
	}
	return o.BenchmarkStatus, true
}

// HasBenchmarkStatus returns a boolean if a field has been set.
func (o *Benchmark) HasBenchmarkStatus() bool {
	if o != nil && !IsNil(o.BenchmarkStatus) {
		return true
	}

	return false
}

// SetBenchmarkStatus gets a reference to the given string and assigns it to the BenchmarkStatus field.
func (o *Benchmark) SetBenchmarkStatus(v string) {
	o.BenchmarkStatus = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Benchmark) GetValues() Values {
	if o == nil || IsNil(o.Values) {
		var ret Values
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetValuesOk() (*Values, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Benchmark) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given Values and assigns it to the Values field.
func (o *Benchmark) SetValues(v Values) {
	o.Values = &v
}

func (o Benchmark) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BenchmarkStatus) {
		toSerialize["benchmarkStatus"] = o.BenchmarkStatus
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableBenchmark struct {
	value *Benchmark
	isSet bool
}

func (v NullableBenchmark) Get() *Benchmark {
	return v.value
}

func (v *NullableBenchmark) Set(val *Benchmark) {
	v.value = val
	v.isSet = true
}

func (v NullableBenchmark) IsSet() bool {
	return v.isSet
}

func (v *NullableBenchmark) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBenchmark(val *Benchmark) *NullableBenchmark {
	return &NullableBenchmark{value: val, isSet: true}
}

func (v NullableBenchmark) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBenchmark) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
