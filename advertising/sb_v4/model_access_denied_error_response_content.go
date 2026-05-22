package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AccessDeniedErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessDeniedErrorResponseContent{}

// AccessDeniedErrorResponseContent struct for AccessDeniedErrorResponseContent
type AccessDeniedErrorResponseContent struct {
	Code      AccessDeniedErrorCode `json:"code"`
	RequestId string                `json:"requestId"`
	Message   string                `json:"message"`
}

type _AccessDeniedErrorResponseContent AccessDeniedErrorResponseContent

// NewAccessDeniedErrorResponseContent instantiates a new AccessDeniedErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessDeniedErrorResponseContent(code AccessDeniedErrorCode, requestId string, message string) *AccessDeniedErrorResponseContent {
	this := AccessDeniedErrorResponseContent{}
	this.Code = code
	this.RequestId = requestId
	this.Message = message
	return &this
}

// NewAccessDeniedErrorResponseContentWithDefaults instantiates a new AccessDeniedErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessDeniedErrorResponseContentWithDefaults() *AccessDeniedErrorResponseContent {
	this := AccessDeniedErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *AccessDeniedErrorResponseContent) GetCode() AccessDeniedErrorCode {
	if o == nil {
		var ret AccessDeniedErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AccessDeniedErrorResponseContent) GetCodeOk() (*AccessDeniedErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AccessDeniedErrorResponseContent) SetCode(v AccessDeniedErrorCode) {
	o.Code = v
}

// GetRequestId returns the RequestId field value
func (o *AccessDeniedErrorResponseContent) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AccessDeniedErrorResponseContent) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AccessDeniedErrorResponseContent) SetRequestId(v string) {
	o.RequestId = v
}

// GetMessage returns the Message field value
func (o *AccessDeniedErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AccessDeniedErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AccessDeniedErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o AccessDeniedErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["requestId"] = o.RequestId
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableAccessDeniedErrorResponseContent struct {
	value *AccessDeniedErrorResponseContent
	isSet bool
}

func (v NullableAccessDeniedErrorResponseContent) Get() *AccessDeniedErrorResponseContent {
	return v.value
}

func (v *NullableAccessDeniedErrorResponseContent) Set(val *AccessDeniedErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDeniedErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDeniedErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDeniedErrorResponseContent(val *AccessDeniedErrorResponseContent) *NullableAccessDeniedErrorResponseContent {
	return &NullableAccessDeniedErrorResponseContent{value: val, isSet: true}
}

func (v NullableAccessDeniedErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccessDeniedErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
