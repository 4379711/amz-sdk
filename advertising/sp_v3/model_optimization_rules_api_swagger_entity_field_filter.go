package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIEntityFieldFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIEntityFieldFilter{}

// OptimizationRulesAPIEntityFieldFilter Filter type and value pair.
type OptimizationRulesAPIEntityFieldFilter struct {
	Values     []string                        `json:"values,omitempty"`
	FilterType *OptimizationRulesAPIFilterType `json:"filterType,omitempty"`
}

// NewOptimizationRulesAPIEntityFieldFilter instantiates a new OptimizationRulesAPIEntityFieldFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIEntityFieldFilter() *OptimizationRulesAPIEntityFieldFilter {
	this := OptimizationRulesAPIEntityFieldFilter{}
	return &this
}

// NewOptimizationRulesAPIEntityFieldFilterWithDefaults instantiates a new OptimizationRulesAPIEntityFieldFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIEntityFieldFilterWithDefaults() *OptimizationRulesAPIEntityFieldFilter {
	this := OptimizationRulesAPIEntityFieldFilter{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *OptimizationRulesAPIEntityFieldFilter) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIEntityFieldFilter) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *OptimizationRulesAPIEntityFieldFilter) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *OptimizationRulesAPIEntityFieldFilter) SetValues(v []string) {
	o.Values = v
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *OptimizationRulesAPIEntityFieldFilter) GetFilterType() OptimizationRulesAPIFilterType {
	if o == nil || IsNil(o.FilterType) {
		var ret OptimizationRulesAPIFilterType
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIEntityFieldFilter) GetFilterTypeOk() (*OptimizationRulesAPIFilterType, bool) {
	if o == nil || IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *OptimizationRulesAPIEntityFieldFilter) HasFilterType() bool {
	if o != nil && !IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given OptimizationRulesAPIFilterType and assigns it to the FilterType field.
func (o *OptimizationRulesAPIEntityFieldFilter) SetFilterType(v OptimizationRulesAPIFilterType) {
	o.FilterType = &v
}

func (o OptimizationRulesAPIEntityFieldFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIEntityFieldFilter struct {
	value *OptimizationRulesAPIEntityFieldFilter
	isSet bool
}

func (v NullableOptimizationRulesAPIEntityFieldFilter) Get() *OptimizationRulesAPIEntityFieldFilter {
	return v.value
}

func (v *NullableOptimizationRulesAPIEntityFieldFilter) Set(val *OptimizationRulesAPIEntityFieldFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIEntityFieldFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIEntityFieldFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIEntityFieldFilter(val *OptimizationRulesAPIEntityFieldFilter) *NullableOptimizationRulesAPIEntityFieldFilter {
	return &NullableOptimizationRulesAPIEntityFieldFilter{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIEntityFieldFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIEntityFieldFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
