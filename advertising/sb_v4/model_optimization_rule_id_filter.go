package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRuleIdFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRuleIdFilter{}

// OptimizationRuleIdFilter Filter optimization rules by the list of optimization rule ids.
type OptimizationRuleIdFilter struct {
	Include []string `json:"include,omitempty"`
}

// NewOptimizationRuleIdFilter instantiates a new OptimizationRuleIdFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRuleIdFilter() *OptimizationRuleIdFilter {
	this := OptimizationRuleIdFilter{}
	return &this
}

// NewOptimizationRuleIdFilterWithDefaults instantiates a new OptimizationRuleIdFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRuleIdFilterWithDefaults() *OptimizationRuleIdFilter {
	this := OptimizationRuleIdFilter{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *OptimizationRuleIdFilter) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRuleIdFilter) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *OptimizationRuleIdFilter) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *OptimizationRuleIdFilter) SetInclude(v []string) {
	o.Include = v
}

func (o OptimizationRuleIdFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	return toSerialize, nil
}

type NullableOptimizationRuleIdFilter struct {
	value *OptimizationRuleIdFilter
	isSet bool
}

func (v NullableOptimizationRuleIdFilter) Get() *OptimizationRuleIdFilter {
	return v.value
}

func (v *NullableOptimizationRuleIdFilter) Set(val *OptimizationRuleIdFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRuleIdFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRuleIdFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRuleIdFilter(val *OptimizationRuleIdFilter) *NullableOptimizationRuleIdFilter {
	return &NullableOptimizationRuleIdFilter{value: val, isSet: true}
}

func (v NullableOptimizationRuleIdFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRuleIdFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
