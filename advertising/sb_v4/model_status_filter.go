package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the StatusFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusFilter{}

// StatusFilter struct for StatusFilter
type StatusFilter struct {
	Include []string `json:"include,omitempty"`
}

// NewStatusFilter instantiates a new StatusFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusFilter() *StatusFilter {
	this := StatusFilter{}
	return &this
}

// NewStatusFilterWithDefaults instantiates a new StatusFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusFilterWithDefaults() *StatusFilter {
	this := StatusFilter{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *StatusFilter) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusFilter) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *StatusFilter) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *StatusFilter) SetInclude(v []string) {
	o.Include = v
}

func (o StatusFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	return toSerialize, nil
}

type NullableStatusFilter struct {
	value *StatusFilter
	isSet bool
}

func (v NullableStatusFilter) Get() *StatusFilter {
	return v.value
}

func (v *NullableStatusFilter) Set(val *StatusFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusFilter(val *StatusFilter) *NullableStatusFilter {
	return &NullableStatusFilter{value: val, isSet: true}
}

func (v NullableStatusFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStatusFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
