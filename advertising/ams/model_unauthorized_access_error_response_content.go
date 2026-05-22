package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the UnauthorizedAccessErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnauthorizedAccessErrorResponseContent{}

// UnauthorizedAccessErrorResponseContent struct for UnauthorizedAccessErrorResponseContent
type UnauthorizedAccessErrorResponseContent struct {
	Code    *string `json:"code,omitempty"`
	Message string  `json:"message"`
}

type _UnauthorizedAccessErrorResponseContent UnauthorizedAccessErrorResponseContent

// NewUnauthorizedAccessErrorResponseContent instantiates a new UnauthorizedAccessErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnauthorizedAccessErrorResponseContent(message string) *UnauthorizedAccessErrorResponseContent {
	this := UnauthorizedAccessErrorResponseContent{}
	this.Message = message
	return &this
}

// NewUnauthorizedAccessErrorResponseContentWithDefaults instantiates a new UnauthorizedAccessErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnauthorizedAccessErrorResponseContentWithDefaults() *UnauthorizedAccessErrorResponseContent {
	this := UnauthorizedAccessErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UnauthorizedAccessErrorResponseContent) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnauthorizedAccessErrorResponseContent) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UnauthorizedAccessErrorResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UnauthorizedAccessErrorResponseContent) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value
func (o *UnauthorizedAccessErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedAccessErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *UnauthorizedAccessErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o UnauthorizedAccessErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableUnauthorizedAccessErrorResponseContent struct {
	value *UnauthorizedAccessErrorResponseContent
	isSet bool
}

func (v NullableUnauthorizedAccessErrorResponseContent) Get() *UnauthorizedAccessErrorResponseContent {
	return v.value
}

func (v *NullableUnauthorizedAccessErrorResponseContent) Set(val *UnauthorizedAccessErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUnauthorizedAccessErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUnauthorizedAccessErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnauthorizedAccessErrorResponseContent(val *UnauthorizedAccessErrorResponseContent) *NullableUnauthorizedAccessErrorResponseContent {
	return &NullableUnauthorizedAccessErrorResponseContent{value: val, isSet: true}
}

func (v NullableUnauthorizedAccessErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnauthorizedAccessErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
