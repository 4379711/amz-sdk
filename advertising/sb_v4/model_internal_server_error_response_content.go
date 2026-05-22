package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the InternalServerErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalServerErrorResponseContent{}

// InternalServerErrorResponseContent struct for InternalServerErrorResponseContent
type InternalServerErrorResponseContent struct {
	Code      InternalServerErrorCode `json:"code"`
	RequestId string                  `json:"requestId"`
	Message   string                  `json:"message"`
}

type _InternalServerErrorResponseContent InternalServerErrorResponseContent

// NewInternalServerErrorResponseContent instantiates a new InternalServerErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalServerErrorResponseContent(code InternalServerErrorCode, requestId string, message string) *InternalServerErrorResponseContent {
	this := InternalServerErrorResponseContent{}
	this.Code = code
	this.RequestId = requestId
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

// GetCode returns the Code field value
func (o *InternalServerErrorResponseContent) GetCode() InternalServerErrorCode {
	if o == nil {
		var ret InternalServerErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *InternalServerErrorResponseContent) GetCodeOk() (*InternalServerErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *InternalServerErrorResponseContent) SetCode(v InternalServerErrorCode) {
	o.Code = v
}

// GetRequestId returns the RequestId field value
func (o *InternalServerErrorResponseContent) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *InternalServerErrorResponseContent) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *InternalServerErrorResponseContent) SetRequestId(v string) {
	o.RequestId = v
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
	toSerialize["code"] = o.Code
	toSerialize["requestId"] = o.RequestId
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
