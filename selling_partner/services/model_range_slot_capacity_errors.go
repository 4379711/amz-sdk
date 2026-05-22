package services

import (
	"github.com/bytedance/sonic"
)

// checks if the RangeSlotCapacityErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeSlotCapacityErrors{}

// RangeSlotCapacityErrors The error response schema for the `getRangeSlotCapacity` operation.
type RangeSlotCapacityErrors struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewRangeSlotCapacityErrors instantiates a new RangeSlotCapacityErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeSlotCapacityErrors() *RangeSlotCapacityErrors {
	this := RangeSlotCapacityErrors{}
	return &this
}

// NewRangeSlotCapacityErrorsWithDefaults instantiates a new RangeSlotCapacityErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeSlotCapacityErrorsWithDefaults() *RangeSlotCapacityErrors {
	this := RangeSlotCapacityErrors{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *RangeSlotCapacityErrors) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeSlotCapacityErrors) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *RangeSlotCapacityErrors) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *RangeSlotCapacityErrors) SetErrors(v []Error) {
	o.Errors = v
}

func (o RangeSlotCapacityErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableRangeSlotCapacityErrors struct {
	value *RangeSlotCapacityErrors
	isSet bool
}

func (v NullableRangeSlotCapacityErrors) Get() *RangeSlotCapacityErrors {
	return v.value
}

func (v *NullableRangeSlotCapacityErrors) Set(val *RangeSlotCapacityErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeSlotCapacityErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeSlotCapacityErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeSlotCapacityErrors(val *RangeSlotCapacityErrors) *NullableRangeSlotCapacityErrors {
	return &NullableRangeSlotCapacityErrors{value: val, isSet: true}
}

func (v NullableRangeSlotCapacityErrors) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRangeSlotCapacityErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
