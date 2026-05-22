package supply_sources_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the ThroughputCap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThroughputCap{}

// ThroughputCap The throughput capacity
type ThroughputCap struct {
	// An unsigned integer that can be only positive or zero.
	Value    *int32    `json:"value,omitempty"`
	TimeUnit *TimeUnit `json:"timeUnit,omitempty"`
}

// NewThroughputCap instantiates a new ThroughputCap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThroughputCap() *ThroughputCap {
	this := ThroughputCap{}
	return &this
}

// NewThroughputCapWithDefaults instantiates a new ThroughputCap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThroughputCapWithDefaults() *ThroughputCap {
	this := ThroughputCap{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ThroughputCap) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThroughputCap) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ThroughputCap) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *ThroughputCap) SetValue(v int32) {
	o.Value = &v
}

// GetTimeUnit returns the TimeUnit field value if set, zero value otherwise.
func (o *ThroughputCap) GetTimeUnit() TimeUnit {
	if o == nil || IsNil(o.TimeUnit) {
		var ret TimeUnit
		return ret
	}
	return *o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThroughputCap) GetTimeUnitOk() (*TimeUnit, bool) {
	if o == nil || IsNil(o.TimeUnit) {
		return nil, false
	}
	return o.TimeUnit, true
}

// HasTimeUnit returns a boolean if a field has been set.
func (o *ThroughputCap) HasTimeUnit() bool {
	if o != nil && !IsNil(o.TimeUnit) {
		return true
	}

	return false
}

// SetTimeUnit gets a reference to the given TimeUnit and assigns it to the TimeUnit field.
func (o *ThroughputCap) SetTimeUnit(v TimeUnit) {
	o.TimeUnit = &v
}

func (o ThroughputCap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.TimeUnit) {
		toSerialize["timeUnit"] = o.TimeUnit
	}
	return toSerialize, nil
}

type NullableThroughputCap struct {
	value *ThroughputCap
	isSet bool
}

func (v NullableThroughputCap) Get() *ThroughputCap {
	return v.value
}

func (v *NullableThroughputCap) Set(val *ThroughputCap) {
	v.value = val
	v.isSet = true
}

func (v NullableThroughputCap) IsSet() bool {
	return v.isSet
}

func (v *NullableThroughputCap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThroughputCap(val *ThroughputCap) *NullableThroughputCap {
	return &NullableThroughputCap{value: val, isSet: true}
}

func (v NullableThroughputCap) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThroughputCap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
