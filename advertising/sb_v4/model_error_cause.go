package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ErrorCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorCause{}

// ErrorCause Structure describing error cause - location in the payload and data causing error.
type ErrorCause struct {
	// Error location, JSON Path expression specifying element of API payload causing error.
	Location string `json:"location"`
	// Optional value causing error.
	Trigger *string `json:"trigger,omitempty"`
}

type _ErrorCause ErrorCause

// NewErrorCause instantiates a new ErrorCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorCause(location string) *ErrorCause {
	this := ErrorCause{}
	this.Location = location
	return &this
}

// NewErrorCauseWithDefaults instantiates a new ErrorCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorCauseWithDefaults() *ErrorCause {
	this := ErrorCause{}
	return &this
}

// GetLocation returns the Location field value
func (o *ErrorCause) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *ErrorCause) SetLocation(v string) {
	o.Location = v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *ErrorCause) GetTrigger() string {
	if o == nil || IsNil(o.Trigger) {
		var ret string
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetTriggerOk() (*string, bool) {
	if o == nil || IsNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *ErrorCause) HasTrigger() bool {
	if o != nil && !IsNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given string and assigns it to the Trigger field.
func (o *ErrorCause) SetTrigger(v string) {
	o.Trigger = &v
}

func (o ErrorCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	if !IsNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}
	return toSerialize, nil
}

type NullableErrorCause struct {
	value *ErrorCause
	isSet bool
}

func (v NullableErrorCause) Get() *ErrorCause {
	return v.value
}

func (v *NullableErrorCause) Set(val *ErrorCause) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCause) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCause(val *ErrorCause) *NullableErrorCause {
	return &NullableErrorCause{value: val, isSet: true}
}

func (v NullableErrorCause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableErrorCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
