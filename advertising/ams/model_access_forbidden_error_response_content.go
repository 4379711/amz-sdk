package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the AccessForbiddenErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessForbiddenErrorResponseContent{}

// AccessForbiddenErrorResponseContent struct for AccessForbiddenErrorResponseContent
type AccessForbiddenErrorResponseContent struct {
	Code    *string `json:"code,omitempty"`
	Message string  `json:"message"`
}

type _AccessForbiddenErrorResponseContent AccessForbiddenErrorResponseContent

// NewAccessForbiddenErrorResponseContent instantiates a new AccessForbiddenErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessForbiddenErrorResponseContent(message string) *AccessForbiddenErrorResponseContent {
	this := AccessForbiddenErrorResponseContent{}
	this.Message = message
	return &this
}

// NewAccessForbiddenErrorResponseContentWithDefaults instantiates a new AccessForbiddenErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessForbiddenErrorResponseContentWithDefaults() *AccessForbiddenErrorResponseContent {
	this := AccessForbiddenErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AccessForbiddenErrorResponseContent) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessForbiddenErrorResponseContent) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AccessForbiddenErrorResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AccessForbiddenErrorResponseContent) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value
func (o *AccessForbiddenErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AccessForbiddenErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AccessForbiddenErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o AccessForbiddenErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableAccessForbiddenErrorResponseContent struct {
	value *AccessForbiddenErrorResponseContent
	isSet bool
}

func (v NullableAccessForbiddenErrorResponseContent) Get() *AccessForbiddenErrorResponseContent {
	return v.value
}

func (v *NullableAccessForbiddenErrorResponseContent) Set(val *AccessForbiddenErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessForbiddenErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessForbiddenErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessForbiddenErrorResponseContent(val *AccessForbiddenErrorResponseContent) *NullableAccessForbiddenErrorResponseContent {
	return &NullableAccessForbiddenErrorResponseContent{value: val, isSet: true}
}

func (v NullableAccessForbiddenErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccessForbiddenErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
