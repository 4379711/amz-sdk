package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the EntityStateFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityStateFilter{}

// EntityStateFilter Filter entities by state.
type EntityStateFilter struct {
	Include []EntityState `json:"include,omitempty"`
}

// NewEntityStateFilter instantiates a new EntityStateFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityStateFilter() *EntityStateFilter {
	this := EntityStateFilter{}
	return &this
}

// NewEntityStateFilterWithDefaults instantiates a new EntityStateFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityStateFilterWithDefaults() *EntityStateFilter {
	this := EntityStateFilter{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *EntityStateFilter) GetInclude() []EntityState {
	if o == nil || IsNil(o.Include) {
		var ret []EntityState
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityStateFilter) GetIncludeOk() ([]EntityState, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *EntityStateFilter) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []EntityState and assigns it to the Include field.
func (o *EntityStateFilter) SetInclude(v []EntityState) {
	o.Include = v
}

func (o EntityStateFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	return toSerialize, nil
}

type NullableEntityStateFilter struct {
	value *EntityStateFilter
	isSet bool
}

func (v NullableEntityStateFilter) Get() *EntityStateFilter {
	return v.value
}

func (v *NullableEntityStateFilter) Set(val *EntityStateFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityStateFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityStateFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityStateFilter(val *EntityStateFilter) *NullableEntityStateFilter {
	return &NullableEntityStateFilter{value: val, isSet: true}
}

func (v NullableEntityStateFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEntityStateFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
