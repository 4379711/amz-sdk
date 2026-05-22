package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ThrottlingErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThrottlingErrorResponseContent{}

// ThrottlingErrorResponseContent struct for ThrottlingErrorResponseContent
type ThrottlingErrorResponseContent struct {
	Code      ThrottlingErrorCode `json:"code"`
	RequestId string              `json:"requestId"`
	Message   string              `json:"message"`
}

type _ThrottlingErrorResponseContent ThrottlingErrorResponseContent

// NewThrottlingErrorResponseContent instantiates a new ThrottlingErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThrottlingErrorResponseContent(code ThrottlingErrorCode, requestId string, message string) *ThrottlingErrorResponseContent {
	this := ThrottlingErrorResponseContent{}
	this.Code = code
	this.RequestId = requestId
	this.Message = message
	return &this
}

// NewThrottlingErrorResponseContentWithDefaults instantiates a new ThrottlingErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThrottlingErrorResponseContentWithDefaults() *ThrottlingErrorResponseContent {
	this := ThrottlingErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *ThrottlingErrorResponseContent) GetCode() ThrottlingErrorCode {
	if o == nil {
		var ret ThrottlingErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ThrottlingErrorResponseContent) GetCodeOk() (*ThrottlingErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ThrottlingErrorResponseContent) SetCode(v ThrottlingErrorCode) {
	o.Code = v
}

// GetRequestId returns the RequestId field value
func (o *ThrottlingErrorResponseContent) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ThrottlingErrorResponseContent) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ThrottlingErrorResponseContent) SetRequestId(v string) {
	o.RequestId = v
}

// GetMessage returns the Message field value
func (o *ThrottlingErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ThrottlingErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ThrottlingErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o ThrottlingErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["requestId"] = o.RequestId
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableThrottlingErrorResponseContent struct {
	value *ThrottlingErrorResponseContent
	isSet bool
}

func (v NullableThrottlingErrorResponseContent) Get() *ThrottlingErrorResponseContent {
	return v.value
}

func (v *NullableThrottlingErrorResponseContent) Set(val *ThrottlingErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableThrottlingErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableThrottlingErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThrottlingErrorResponseContent(val *ThrottlingErrorResponseContent) *NullableThrottlingErrorResponseContent {
	return &NullableThrottlingErrorResponseContent{value: val, isSet: true}
}

func (v NullableThrottlingErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThrottlingErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
