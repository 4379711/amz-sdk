package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UnauthorizedErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnauthorizedErrorResponseContent{}

// UnauthorizedErrorResponseContent struct for UnauthorizedErrorResponseContent
type UnauthorizedErrorResponseContent struct {
	Code      UnauthorizedErrorCode `json:"code"`
	RequestId string                `json:"requestId"`
	Message   string                `json:"message"`
}

type _UnauthorizedErrorResponseContent UnauthorizedErrorResponseContent

// NewUnauthorizedErrorResponseContent instantiates a new UnauthorizedErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnauthorizedErrorResponseContent(code UnauthorizedErrorCode, requestId string, message string) *UnauthorizedErrorResponseContent {
	this := UnauthorizedErrorResponseContent{}
	this.Code = code
	this.RequestId = requestId
	this.Message = message
	return &this
}

// NewUnauthorizedErrorResponseContentWithDefaults instantiates a new UnauthorizedErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnauthorizedErrorResponseContentWithDefaults() *UnauthorizedErrorResponseContent {
	this := UnauthorizedErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *UnauthorizedErrorResponseContent) GetCode() UnauthorizedErrorCode {
	if o == nil {
		var ret UnauthorizedErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedErrorResponseContent) GetCodeOk() (*UnauthorizedErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *UnauthorizedErrorResponseContent) SetCode(v UnauthorizedErrorCode) {
	o.Code = v
}

// GetRequestId returns the RequestId field value
func (o *UnauthorizedErrorResponseContent) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedErrorResponseContent) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *UnauthorizedErrorResponseContent) SetRequestId(v string) {
	o.RequestId = v
}

// GetMessage returns the Message field value
func (o *UnauthorizedErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *UnauthorizedErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o UnauthorizedErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["requestId"] = o.RequestId
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableUnauthorizedErrorResponseContent struct {
	value *UnauthorizedErrorResponseContent
	isSet bool
}

func (v NullableUnauthorizedErrorResponseContent) Get() *UnauthorizedErrorResponseContent {
	return v.value
}

func (v *NullableUnauthorizedErrorResponseContent) Set(val *UnauthorizedErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUnauthorizedErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUnauthorizedErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnauthorizedErrorResponseContent(val *UnauthorizedErrorResponseContent) *NullableUnauthorizedErrorResponseContent {
	return &NullableUnauthorizedErrorResponseContent{value: val, isSet: true}
}

func (v NullableUnauthorizedErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnauthorizedErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
