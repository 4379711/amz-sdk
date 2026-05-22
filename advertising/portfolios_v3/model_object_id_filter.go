package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ObjectIdFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectIdFilter{}

// ObjectIdFilter Filter entities by the list of objectIds
type ObjectIdFilter struct {
	Include []string `json:"include,omitempty"`
}

// NewObjectIdFilter instantiates a new ObjectIdFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectIdFilter() *ObjectIdFilter {
	this := ObjectIdFilter{}
	return &this
}

// NewObjectIdFilterWithDefaults instantiates a new ObjectIdFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectIdFilterWithDefaults() *ObjectIdFilter {
	this := ObjectIdFilter{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *ObjectIdFilter) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectIdFilter) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *ObjectIdFilter) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *ObjectIdFilter) SetInclude(v []string) {
	o.Include = v
}

func (o ObjectIdFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	return toSerialize, nil
}

type NullableObjectIdFilter struct {
	value *ObjectIdFilter
	isSet bool
}

func (v NullableObjectIdFilter) Get() *ObjectIdFilter {
	return v.value
}

func (v *NullableObjectIdFilter) Set(val *ObjectIdFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectIdFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectIdFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectIdFilter(val *ObjectIdFilter) *NullableObjectIdFilter {
	return &NullableObjectIdFilter{value: val, isSet: true}
}

func (v NullableObjectIdFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableObjectIdFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
