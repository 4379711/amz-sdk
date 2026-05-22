package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the NotFoundExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotFoundExceptionResponseContent{}

// NotFoundExceptionResponseContent struct for NotFoundExceptionResponseContent
type NotFoundExceptionResponseContent struct {
	Code    NotFoundErrorCode `json:"code"`
	Message string            `json:"message"`
}

type _NotFoundExceptionResponseContent NotFoundExceptionResponseContent

// NewNotFoundExceptionResponseContent instantiates a new NotFoundExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotFoundExceptionResponseContent(code NotFoundErrorCode, message string) *NotFoundExceptionResponseContent {
	this := NotFoundExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewNotFoundExceptionResponseContentWithDefaults instantiates a new NotFoundExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotFoundExceptionResponseContentWithDefaults() *NotFoundExceptionResponseContent {
	this := NotFoundExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *NotFoundExceptionResponseContent) GetCode() NotFoundErrorCode {
	if o == nil {
		var ret NotFoundErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *NotFoundExceptionResponseContent) GetCodeOk() (*NotFoundErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *NotFoundExceptionResponseContent) SetCode(v NotFoundErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *NotFoundExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NotFoundExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NotFoundExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o NotFoundExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableNotFoundExceptionResponseContent struct {
	value *NotFoundExceptionResponseContent
	isSet bool
}

func (v NullableNotFoundExceptionResponseContent) Get() *NotFoundExceptionResponseContent {
	return v.value
}

func (v *NullableNotFoundExceptionResponseContent) Set(val *NotFoundExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundExceptionResponseContent(val *NotFoundExceptionResponseContent) *NullableNotFoundExceptionResponseContent {
	return &NullableNotFoundExceptionResponseContent{value: val, isSet: true}
}

func (v NullableNotFoundExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableNotFoundExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
