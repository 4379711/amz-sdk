package supply_sources_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the ThroughputConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThroughputConfig{}

// ThroughputConfig The throughput configuration.
type ThroughputConfig struct {
	ThroughputCap  *ThroughputCap `json:"throughputCap,omitempty"`
	ThroughputUnit ThroughputUnit `json:"throughputUnit"`
}

type _ThroughputConfig ThroughputConfig

// NewThroughputConfig instantiates a new ThroughputConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThroughputConfig(throughputUnit ThroughputUnit) *ThroughputConfig {
	this := ThroughputConfig{}
	this.ThroughputUnit = throughputUnit
	return &this
}

// NewThroughputConfigWithDefaults instantiates a new ThroughputConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThroughputConfigWithDefaults() *ThroughputConfig {
	this := ThroughputConfig{}
	return &this
}

// GetThroughputCap returns the ThroughputCap field value if set, zero value otherwise.
func (o *ThroughputConfig) GetThroughputCap() ThroughputCap {
	if o == nil || IsNil(o.ThroughputCap) {
		var ret ThroughputCap
		return ret
	}
	return *o.ThroughputCap
}

// GetThroughputCapOk returns a tuple with the ThroughputCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThroughputConfig) GetThroughputCapOk() (*ThroughputCap, bool) {
	if o == nil || IsNil(o.ThroughputCap) {
		return nil, false
	}
	return o.ThroughputCap, true
}

// HasThroughputCap returns a boolean if a field has been set.
func (o *ThroughputConfig) HasThroughputCap() bool {
	if o != nil && !IsNil(o.ThroughputCap) {
		return true
	}

	return false
}

// SetThroughputCap gets a reference to the given ThroughputCap and assigns it to the ThroughputCap field.
func (o *ThroughputConfig) SetThroughputCap(v ThroughputCap) {
	o.ThroughputCap = &v
}

// GetThroughputUnit returns the ThroughputUnit field value
func (o *ThroughputConfig) GetThroughputUnit() ThroughputUnit {
	if o == nil {
		var ret ThroughputUnit
		return ret
	}

	return o.ThroughputUnit
}

// GetThroughputUnitOk returns a tuple with the ThroughputUnit field value
// and a boolean to check if the value has been set.
func (o *ThroughputConfig) GetThroughputUnitOk() (*ThroughputUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThroughputUnit, true
}

// SetThroughputUnit sets field value
func (o *ThroughputConfig) SetThroughputUnit(v ThroughputUnit) {
	o.ThroughputUnit = v
}

func (o ThroughputConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThroughputCap) {
		toSerialize["throughputCap"] = o.ThroughputCap
	}
	toSerialize["throughputUnit"] = o.ThroughputUnit
	return toSerialize, nil
}

type NullableThroughputConfig struct {
	value *ThroughputConfig
	isSet bool
}

func (v NullableThroughputConfig) Get() *ThroughputConfig {
	return v.value
}

func (v *NullableThroughputConfig) Set(val *ThroughputConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableThroughputConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableThroughputConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThroughputConfig(val *ThroughputConfig) *NullableThroughputConfig {
	return &NullableThroughputConfig{value: val, isSet: true}
}

func (v NullableThroughputConfig) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThroughputConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
