package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ThrottlingExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThrottlingExceptionResponseContent{}

// ThrottlingExceptionResponseContent struct for ThrottlingExceptionResponseContent
type ThrottlingExceptionResponseContent struct {
	Code ThrottledErrorCode `json:"code"`
	// Human readable error message.
	Message string `json:"message"`
}

type _ThrottlingExceptionResponseContent ThrottlingExceptionResponseContent

// NewThrottlingExceptionResponseContent instantiates a new ThrottlingExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThrottlingExceptionResponseContent(code ThrottledErrorCode, message string) *ThrottlingExceptionResponseContent {
	this := ThrottlingExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewThrottlingExceptionResponseContentWithDefaults instantiates a new ThrottlingExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThrottlingExceptionResponseContentWithDefaults() *ThrottlingExceptionResponseContent {
	this := ThrottlingExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *ThrottlingExceptionResponseContent) GetCode() ThrottledErrorCode {
	if o == nil {
		var ret ThrottledErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ThrottlingExceptionResponseContent) GetCodeOk() (*ThrottledErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ThrottlingExceptionResponseContent) SetCode(v ThrottledErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ThrottlingExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ThrottlingExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ThrottlingExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o ThrottlingExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableThrottlingExceptionResponseContent struct {
	value *ThrottlingExceptionResponseContent
	isSet bool
}

func (v NullableThrottlingExceptionResponseContent) Get() *ThrottlingExceptionResponseContent {
	return v.value
}

func (v *NullableThrottlingExceptionResponseContent) Set(val *ThrottlingExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableThrottlingExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableThrottlingExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThrottlingExceptionResponseContent(val *ThrottlingExceptionResponseContent) *NullableThrottlingExceptionResponseContent {
	return &NullableThrottlingExceptionResponseContent{value: val, isSet: true}
}

func (v NullableThrottlingExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThrottlingExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
