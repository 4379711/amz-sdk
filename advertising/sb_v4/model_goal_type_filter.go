package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the GoalTypeFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoalTypeFilter{}

// GoalTypeFilter Filter entities by goal type.
type GoalTypeFilter struct {
	Include []string `json:"include,omitempty"`
}

// NewGoalTypeFilter instantiates a new GoalTypeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoalTypeFilter() *GoalTypeFilter {
	this := GoalTypeFilter{}
	return &this
}

// NewGoalTypeFilterWithDefaults instantiates a new GoalTypeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoalTypeFilterWithDefaults() *GoalTypeFilter {
	this := GoalTypeFilter{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *GoalTypeFilter) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalTypeFilter) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *GoalTypeFilter) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *GoalTypeFilter) SetInclude(v []string) {
	o.Include = v
}

func (o GoalTypeFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	return toSerialize, nil
}

type NullableGoalTypeFilter struct {
	value *GoalTypeFilter
	isSet bool
}

func (v NullableGoalTypeFilter) Get() *GoalTypeFilter {
	return v.value
}

func (v *NullableGoalTypeFilter) Set(val *GoalTypeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableGoalTypeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableGoalTypeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoalTypeFilter(val *GoalTypeFilter) *NullableGoalTypeFilter {
	return &NullableGoalTypeFilter{value: val, isSet: true}
}

func (v NullableGoalTypeFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGoalTypeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
