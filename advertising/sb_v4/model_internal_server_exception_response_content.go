package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the InternalServerExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalServerExceptionResponseContent{}

// InternalServerExceptionResponseContent struct for InternalServerExceptionResponseContent
type InternalServerExceptionResponseContent struct {
	Code InternalErrorErrorCode `json:"code"`
	// Human readable error message.
	Message string `json:"message"`
}

type _InternalServerExceptionResponseContent InternalServerExceptionResponseContent

// NewInternalServerExceptionResponseContent instantiates a new InternalServerExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalServerExceptionResponseContent(code InternalErrorErrorCode, message string) *InternalServerExceptionResponseContent {
	this := InternalServerExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewInternalServerExceptionResponseContentWithDefaults instantiates a new InternalServerExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalServerExceptionResponseContentWithDefaults() *InternalServerExceptionResponseContent {
	this := InternalServerExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *InternalServerExceptionResponseContent) GetCode() InternalErrorErrorCode {
	if o == nil {
		var ret InternalErrorErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *InternalServerExceptionResponseContent) GetCodeOk() (*InternalErrorErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *InternalServerExceptionResponseContent) SetCode(v InternalErrorErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *InternalServerExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InternalServerExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InternalServerExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o InternalServerExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableInternalServerExceptionResponseContent struct {
	value *InternalServerExceptionResponseContent
	isSet bool
}

func (v NullableInternalServerExceptionResponseContent) Get() *InternalServerExceptionResponseContent {
	return v.value
}

func (v *NullableInternalServerExceptionResponseContent) Set(val *InternalServerExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalServerExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalServerExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalServerExceptionResponseContent(val *InternalServerExceptionResponseContent) *NullableInternalServerExceptionResponseContent {
	return &NullableInternalServerExceptionResponseContent{value: val, isSet: true}
}

func (v NullableInternalServerExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInternalServerExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
