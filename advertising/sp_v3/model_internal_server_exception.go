package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the InternalServerException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalServerException{}

// InternalServerException Returns information about an InternalServerException.
type InternalServerException struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewInternalServerException instantiates a new InternalServerException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalServerException() *InternalServerException {
	this := InternalServerException{}
	return &this
}

// NewInternalServerExceptionWithDefaults instantiates a new InternalServerException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalServerExceptionWithDefaults() *InternalServerException {
	this := InternalServerException{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InternalServerException) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalServerException) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InternalServerException) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InternalServerException) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InternalServerException) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalServerException) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InternalServerException) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *InternalServerException) SetDetails(v string) {
	o.Details = &v
}

func (o InternalServerException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableInternalServerException struct {
	value *InternalServerException
	isSet bool
}

func (v NullableInternalServerException) Get() *InternalServerException {
	return v.value
}

func (v *NullableInternalServerException) Set(val *InternalServerException) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalServerException) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalServerException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalServerException(val *InternalServerException) *NullableInternalServerException {
	return &NullableInternalServerException{value: val, isSet: true}
}

func (v NullableInternalServerException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInternalServerException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
