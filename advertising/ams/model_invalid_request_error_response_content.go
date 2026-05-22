package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the InvalidRequestErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidRequestErrorResponseContent{}

// InvalidRequestErrorResponseContent struct for InvalidRequestErrorResponseContent
type InvalidRequestErrorResponseContent struct {
	Code    *string `json:"code,omitempty"`
	Message string  `json:"message"`
}

type _InvalidRequestErrorResponseContent InvalidRequestErrorResponseContent

// NewInvalidRequestErrorResponseContent instantiates a new InvalidRequestErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidRequestErrorResponseContent(message string) *InvalidRequestErrorResponseContent {
	this := InvalidRequestErrorResponseContent{}
	this.Message = message
	return &this
}

// NewInvalidRequestErrorResponseContentWithDefaults instantiates a new InvalidRequestErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidRequestErrorResponseContentWithDefaults() *InvalidRequestErrorResponseContent {
	this := InvalidRequestErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InvalidRequestErrorResponseContent) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidRequestErrorResponseContent) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InvalidRequestErrorResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InvalidRequestErrorResponseContent) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value
func (o *InvalidRequestErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InvalidRequestErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InvalidRequestErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o InvalidRequestErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableInvalidRequestErrorResponseContent struct {
	value *InvalidRequestErrorResponseContent
	isSet bool
}

func (v NullableInvalidRequestErrorResponseContent) Get() *InvalidRequestErrorResponseContent {
	return v.value
}

func (v *NullableInvalidRequestErrorResponseContent) Set(val *InvalidRequestErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidRequestErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidRequestErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidRequestErrorResponseContent(val *InvalidRequestErrorResponseContent) *NullableInvalidRequestErrorResponseContent {
	return &NullableInvalidRequestErrorResponseContent{value: val, isSet: true}
}

func (v NullableInvalidRequestErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvalidRequestErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
