package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateNegativeTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNegativeTargetingClause{}

// UpdateNegativeTargetingClause struct for UpdateNegativeTargetingClause
type UpdateNegativeTargetingClause struct {
	State    *string `json:"state,omitempty"`
	TargetId int64   `json:"targetId"`
}

type _UpdateNegativeTargetingClause UpdateNegativeTargetingClause

// NewUpdateNegativeTargetingClause instantiates a new UpdateNegativeTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNegativeTargetingClause(targetId int64) *UpdateNegativeTargetingClause {
	this := UpdateNegativeTargetingClause{}
	this.TargetId = targetId
	return &this
}

// NewUpdateNegativeTargetingClauseWithDefaults instantiates a new UpdateNegativeTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNegativeTargetingClauseWithDefaults() *UpdateNegativeTargetingClause {
	this := UpdateNegativeTargetingClause{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateNegativeTargetingClause) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNegativeTargetingClause) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateNegativeTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UpdateNegativeTargetingClause) SetState(v string) {
	o.State = &v
}

// GetTargetId returns the TargetId field value
func (o *UpdateNegativeTargetingClause) GetTargetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *UpdateNegativeTargetingClause) GetTargetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *UpdateNegativeTargetingClause) SetTargetId(v int64) {
	o.TargetId = v
}

func (o UpdateNegativeTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["targetId"] = o.TargetId
	return toSerialize, nil
}

type NullableUpdateNegativeTargetingClause struct {
	value *UpdateNegativeTargetingClause
	isSet bool
}

func (v NullableUpdateNegativeTargetingClause) Get() *UpdateNegativeTargetingClause {
	return v.value
}

func (v *NullableUpdateNegativeTargetingClause) Set(val *UpdateNegativeTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNegativeTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNegativeTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNegativeTargetingClause(val *UpdateNegativeTargetingClause) *NullableUpdateNegativeTargetingClause {
	return &NullableUpdateNegativeTargetingClause{value: val, isSet: true}
}

func (v NullableUpdateNegativeTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateNegativeTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
