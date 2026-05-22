package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the TaskIdFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskIdFilter{}

// TaskIdFilter struct for TaskIdFilter
type TaskIdFilter struct {
	Include []string `json:"include,omitempty"`
}

// NewTaskIdFilter instantiates a new TaskIdFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskIdFilter() *TaskIdFilter {
	this := TaskIdFilter{}
	return &this
}

// NewTaskIdFilterWithDefaults instantiates a new TaskIdFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskIdFilterWithDefaults() *TaskIdFilter {
	this := TaskIdFilter{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *TaskIdFilter) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskIdFilter) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *TaskIdFilter) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *TaskIdFilter) SetInclude(v []string) {
	o.Include = v
}

func (o TaskIdFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	return toSerialize, nil
}

type NullableTaskIdFilter struct {
	value *TaskIdFilter
	isSet bool
}

func (v NullableTaskIdFilter) Get() *TaskIdFilter {
	return v.value
}

func (v *NullableTaskIdFilter) Set(val *TaskIdFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskIdFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskIdFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskIdFilter(val *TaskIdFilter) *NullableTaskIdFilter {
	return &NullableTaskIdFilter{value: val, isSet: true}
}

func (v NullableTaskIdFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTaskIdFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
