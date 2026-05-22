package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the InternalServerErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalServerErrorResponseContent{}

// InternalServerErrorResponseContent struct for InternalServerErrorResponseContent
type InternalServerErrorResponseContent struct {
	Code    *string `json:"code,omitempty"`
	Message string  `json:"message"`
}

type _InternalServerErrorResponseContent InternalServerErrorResponseContent

// NewInternalServerErrorResponseContent instantiates a new InternalServerErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalServerErrorResponseContent(message string) *InternalServerErrorResponseContent {
	this := InternalServerErrorResponseContent{}
	this.Message = message
	return &this
}

// NewInternalServerErrorResponseContentWithDefaults instantiates a new InternalServerErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalServerErrorResponseContentWithDefaults() *InternalServerErrorResponseContent {
	this := InternalServerErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InternalServerErrorResponseContent) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalServerErrorResponseContent) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InternalServerErrorResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InternalServerErrorResponseContent) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value
func (o *InternalServerErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InternalServerErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InternalServerErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o InternalServerErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableInternalServerErrorResponseContent struct {
	value *InternalServerErrorResponseContent
	isSet bool
}

func (v NullableInternalServerErrorResponseContent) Get() *InternalServerErrorResponseContent {
	return v.value
}

func (v *NullableInternalServerErrorResponseContent) Set(val *InternalServerErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalServerErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalServerErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalServerErrorResponseContent(val *InternalServerErrorResponseContent) *NullableInternalServerErrorResponseContent {
	return &NullableInternalServerErrorResponseContent{value: val, isSet: true}
}

func (v NullableInternalServerErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInternalServerErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
