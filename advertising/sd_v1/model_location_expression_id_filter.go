package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the LocationExpressionIdFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationExpressionIdFilter{}

// LocationExpressionIdFilter Filter entities by the list of objectIds
type LocationExpressionIdFilter struct {
	// Array of Location Expression Ids
	Include []int64 `json:"include"`
}

type _LocationExpressionIdFilter LocationExpressionIdFilter

// NewLocationExpressionIdFilter instantiates a new LocationExpressionIdFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationExpressionIdFilter(include []int64) *LocationExpressionIdFilter {
	this := LocationExpressionIdFilter{}
	this.Include = include
	return &this
}

// NewLocationExpressionIdFilterWithDefaults instantiates a new LocationExpressionIdFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationExpressionIdFilterWithDefaults() *LocationExpressionIdFilter {
	this := LocationExpressionIdFilter{}
	return &this
}

// GetInclude returns the Include field value
func (o *LocationExpressionIdFilter) GetInclude() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *LocationExpressionIdFilter) GetIncludeOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Include, true
}

// SetInclude sets field value
func (o *LocationExpressionIdFilter) SetInclude(v []int64) {
	o.Include = v
}

func (o LocationExpressionIdFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["include"] = o.Include
	return toSerialize, nil
}

type NullableLocationExpressionIdFilter struct {
	value *LocationExpressionIdFilter
	isSet bool
}

func (v NullableLocationExpressionIdFilter) Get() *LocationExpressionIdFilter {
	return v.value
}

func (v *NullableLocationExpressionIdFilter) Set(val *LocationExpressionIdFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationExpressionIdFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationExpressionIdFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationExpressionIdFilter(val *LocationExpressionIdFilter) *NullableLocationExpressionIdFilter {
	return &NullableLocationExpressionIdFilter{value: val, isSet: true}
}

func (v NullableLocationExpressionIdFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLocationExpressionIdFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
