package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the ErrorList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorList{}

// ErrorList A list of error responses returned when a request is unsuccessful.
type ErrorList struct {
	Errors []Error `json:"errors"`
}

type _ErrorList ErrorList

// NewErrorList instantiates a new ErrorList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorList(errors []Error) *ErrorList {
	this := ErrorList{}
	this.Errors = errors
	return &this
}

// NewErrorListWithDefaults instantiates a new ErrorList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorListWithDefaults() *ErrorList {
	this := ErrorList{}
	return &this
}

// GetErrors returns the Errors field value
func (o *ErrorList) GetErrors() []Error {
	if o == nil {
		var ret []Error
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ErrorList) GetErrorsOk() ([]Error, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *ErrorList) SetErrors(v []Error) {
	o.Errors = v
}

func (o ErrorList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

type NullableErrorList struct {
	value *ErrorList
	isSet bool
}

func (v NullableErrorList) Get() *ErrorList {
	return v.value
}

func (v *NullableErrorList) Set(val *ErrorList) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorList) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorList(val *ErrorList) *NullableErrorList {
	return &NullableErrorList{value: val, isSet: true}
}

func (v NullableErrorList) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableErrorList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
