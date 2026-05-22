package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the NotFoundErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotFoundErrorResponseContent{}

// NotFoundErrorResponseContent struct for NotFoundErrorResponseContent
type NotFoundErrorResponseContent struct {
	Code      NotFoundErrorCode `json:"code"`
	RequestId string            `json:"requestId"`
	Message   string            `json:"message"`
}

type _NotFoundErrorResponseContent NotFoundErrorResponseContent

// NewNotFoundErrorResponseContent instantiates a new NotFoundErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotFoundErrorResponseContent(code NotFoundErrorCode, requestId string, message string) *NotFoundErrorResponseContent {
	this := NotFoundErrorResponseContent{}
	this.Code = code
	this.RequestId = requestId
	this.Message = message
	return &this
}

// NewNotFoundErrorResponseContentWithDefaults instantiates a new NotFoundErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotFoundErrorResponseContentWithDefaults() *NotFoundErrorResponseContent {
	this := NotFoundErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *NotFoundErrorResponseContent) GetCode() NotFoundErrorCode {
	if o == nil {
		var ret NotFoundErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *NotFoundErrorResponseContent) GetCodeOk() (*NotFoundErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *NotFoundErrorResponseContent) SetCode(v NotFoundErrorCode) {
	o.Code = v
}

// GetRequestId returns the RequestId field value
func (o *NotFoundErrorResponseContent) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *NotFoundErrorResponseContent) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *NotFoundErrorResponseContent) SetRequestId(v string) {
	o.RequestId = v
}

// GetMessage returns the Message field value
func (o *NotFoundErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NotFoundErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NotFoundErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o NotFoundErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["requestId"] = o.RequestId
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableNotFoundErrorResponseContent struct {
	value *NotFoundErrorResponseContent
	isSet bool
}

func (v NullableNotFoundErrorResponseContent) Get() *NotFoundErrorResponseContent {
	return v.value
}

func (v *NullableNotFoundErrorResponseContent) Set(val *NotFoundErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundErrorResponseContent(val *NotFoundErrorResponseContent) *NullableNotFoundErrorResponseContent {
	return &NullableNotFoundErrorResponseContent{value: val, isSet: true}
}

func (v NullableNotFoundErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableNotFoundErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
