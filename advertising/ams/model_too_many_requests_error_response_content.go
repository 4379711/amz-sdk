package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the TooManyRequestsErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TooManyRequestsErrorResponseContent{}

// TooManyRequestsErrorResponseContent struct for TooManyRequestsErrorResponseContent
type TooManyRequestsErrorResponseContent struct {
	Code    *string `json:"code,omitempty"`
	Message string  `json:"message"`
}

type _TooManyRequestsErrorResponseContent TooManyRequestsErrorResponseContent

// NewTooManyRequestsErrorResponseContent instantiates a new TooManyRequestsErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTooManyRequestsErrorResponseContent(message string) *TooManyRequestsErrorResponseContent {
	this := TooManyRequestsErrorResponseContent{}
	this.Message = message
	return &this
}

// NewTooManyRequestsErrorResponseContentWithDefaults instantiates a new TooManyRequestsErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTooManyRequestsErrorResponseContentWithDefaults() *TooManyRequestsErrorResponseContent {
	this := TooManyRequestsErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TooManyRequestsErrorResponseContent) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TooManyRequestsErrorResponseContent) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TooManyRequestsErrorResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TooManyRequestsErrorResponseContent) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value
func (o *TooManyRequestsErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TooManyRequestsErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TooManyRequestsErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o TooManyRequestsErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableTooManyRequestsErrorResponseContent struct {
	value *TooManyRequestsErrorResponseContent
	isSet bool
}

func (v NullableTooManyRequestsErrorResponseContent) Get() *TooManyRequestsErrorResponseContent {
	return v.value
}

func (v *NullableTooManyRequestsErrorResponseContent) Set(val *TooManyRequestsErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableTooManyRequestsErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableTooManyRequestsErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTooManyRequestsErrorResponseContent(val *TooManyRequestsErrorResponseContent) *NullableTooManyRequestsErrorResponseContent {
	return &NullableTooManyRequestsErrorResponseContent{value: val, isSet: true}
}

func (v NullableTooManyRequestsErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTooManyRequestsErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
