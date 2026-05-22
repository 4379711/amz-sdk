package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AccessDeniedException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessDeniedException{}

// AccessDeniedException Returns information about an AccessDeniedException.
type AccessDeniedException struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewAccessDeniedException instantiates a new AccessDeniedException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessDeniedException() *AccessDeniedException {
	this := AccessDeniedException{}
	return &this
}

// NewAccessDeniedExceptionWithDefaults instantiates a new AccessDeniedException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessDeniedExceptionWithDefaults() *AccessDeniedException {
	this := AccessDeniedException{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AccessDeniedException) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessDeniedException) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AccessDeniedException) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AccessDeniedException) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *AccessDeniedException) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessDeniedException) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *AccessDeniedException) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *AccessDeniedException) SetDetails(v string) {
	o.Details = &v
}

func (o AccessDeniedException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableAccessDeniedException struct {
	value *AccessDeniedException
	isSet bool
}

func (v NullableAccessDeniedException) Get() *AccessDeniedException {
	return v.value
}

func (v *NullableAccessDeniedException) Set(val *AccessDeniedException) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDeniedException) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDeniedException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDeniedException(val *AccessDeniedException) *NullableAccessDeniedException {
	return &NullableAccessDeniedException{value: val, isSet: true}
}

func (v NullableAccessDeniedException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccessDeniedException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
