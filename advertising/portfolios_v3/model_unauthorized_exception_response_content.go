package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the UnauthorizedExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnauthorizedExceptionResponseContent{}

// UnauthorizedExceptionResponseContent struct for UnauthorizedExceptionResponseContent
type UnauthorizedExceptionResponseContent struct {
	Code UnauthorizedErrorCode `json:"code"`
	// Human readable error message
	Message string `json:"message"`
}

type _UnauthorizedExceptionResponseContent UnauthorizedExceptionResponseContent

// NewUnauthorizedExceptionResponseContent instantiates a new UnauthorizedExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnauthorizedExceptionResponseContent(code UnauthorizedErrorCode, message string) *UnauthorizedExceptionResponseContent {
	this := UnauthorizedExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewUnauthorizedExceptionResponseContentWithDefaults instantiates a new UnauthorizedExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnauthorizedExceptionResponseContentWithDefaults() *UnauthorizedExceptionResponseContent {
	this := UnauthorizedExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *UnauthorizedExceptionResponseContent) GetCode() UnauthorizedErrorCode {
	if o == nil {
		var ret UnauthorizedErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedExceptionResponseContent) GetCodeOk() (*UnauthorizedErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *UnauthorizedExceptionResponseContent) SetCode(v UnauthorizedErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *UnauthorizedExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *UnauthorizedExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o UnauthorizedExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableUnauthorizedExceptionResponseContent struct {
	value *UnauthorizedExceptionResponseContent
	isSet bool
}

func (v NullableUnauthorizedExceptionResponseContent) Get() *UnauthorizedExceptionResponseContent {
	return v.value
}

func (v *NullableUnauthorizedExceptionResponseContent) Set(val *UnauthorizedExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUnauthorizedExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUnauthorizedExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnauthorizedExceptionResponseContent(val *UnauthorizedExceptionResponseContent) *NullableUnauthorizedExceptionResponseContent {
	return &NullableUnauthorizedExceptionResponseContent{value: val, isSet: true}
}

func (v NullableUnauthorizedExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnauthorizedExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
