package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BaseNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseNegativeTargetingClause{}

// BaseNegativeTargetingClause struct for BaseNegativeTargetingClause
type BaseNegativeTargetingClause struct {
	State *string `json:"state,omitempty"`
}

// NewBaseNegativeTargetingClause instantiates a new BaseNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseNegativeTargetingClause() *BaseNegativeTargetingClause {
	this := BaseNegativeTargetingClause{}
	return &this
}

// NewBaseNegativeTargetingClauseWithDefaults instantiates a new BaseNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseNegativeTargetingClauseWithDefaults() *BaseNegativeTargetingClause {
	this := BaseNegativeTargetingClause{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BaseNegativeTargetingClause) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseNegativeTargetingClause) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BaseNegativeTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BaseNegativeTargetingClause) SetState(v string) {
	o.State = &v
}

func (o BaseNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableBaseNegativeTargetingClause struct {
	value *BaseNegativeTargetingClause
	isSet bool
}

func (v NullableBaseNegativeTargetingClause) Get() *BaseNegativeTargetingClause {
	return v.value
}

func (v *NullableBaseNegativeTargetingClause) Set(val *BaseNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseNegativeTargetingClause(val *BaseNegativeTargetingClause) *NullableBaseNegativeTargetingClause {
	return &NullableBaseNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableBaseNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBaseNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
