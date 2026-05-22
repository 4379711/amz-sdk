package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Errors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Errors{}

// Errors A list of error responses returned when a request is unsuccessful.
type Errors struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors"`
}

type _Errors Errors

// NewErrors instantiates a new Errors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrors(errors []Error) *Errors {
	this := Errors{}
	this.Errors = errors
	return &this
}

// NewErrorsWithDefaults instantiates a new Errors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorsWithDefaults() *Errors {
	this := Errors{}
	return &this
}

// GetErrors returns the Errors field value
func (o *Errors) GetErrors() []Error {
	if o == nil {
		var ret []Error
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *Errors) GetErrorsOk() ([]Error, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *Errors) SetErrors(v []Error) {
	o.Errors = v
}

func (o Errors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

type NullableErrors struct {
	value *Errors
	isSet bool
}

func (v NullableErrors) Get() *Errors {
	return v.value
}

func (v *NullableErrors) Set(val *Errors) {
	v.value = val
	v.isSet = true
}

func (v NullableErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrors(val *Errors) *NullableErrors {
	return &NullableErrors{value: val, isSet: true}
}

func (v NullableErrors) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
