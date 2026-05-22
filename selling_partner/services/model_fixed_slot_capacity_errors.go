package services

import (
	"github.com/bytedance/sonic"
)

// checks if the FixedSlotCapacityErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FixedSlotCapacityErrors{}

// FixedSlotCapacityErrors The error response schema for the `getFixedSlotCapacity` operation.
type FixedSlotCapacityErrors struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewFixedSlotCapacityErrors instantiates a new FixedSlotCapacityErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedSlotCapacityErrors() *FixedSlotCapacityErrors {
	this := FixedSlotCapacityErrors{}
	return &this
}

// NewFixedSlotCapacityErrorsWithDefaults instantiates a new FixedSlotCapacityErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedSlotCapacityErrorsWithDefaults() *FixedSlotCapacityErrors {
	this := FixedSlotCapacityErrors{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *FixedSlotCapacityErrors) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedSlotCapacityErrors) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *FixedSlotCapacityErrors) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *FixedSlotCapacityErrors) SetErrors(v []Error) {
	o.Errors = v
}

func (o FixedSlotCapacityErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableFixedSlotCapacityErrors struct {
	value *FixedSlotCapacityErrors
	isSet bool
}

func (v NullableFixedSlotCapacityErrors) Get() *FixedSlotCapacityErrors {
	return v.value
}

func (v *NullableFixedSlotCapacityErrors) Set(val *FixedSlotCapacityErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedSlotCapacityErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedSlotCapacityErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedSlotCapacityErrors(val *FixedSlotCapacityErrors) *NullableFixedSlotCapacityErrors {
	return &NullableFixedSlotCapacityErrors{value: val, isSet: true}
}

func (v NullableFixedSlotCapacityErrors) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFixedSlotCapacityErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
