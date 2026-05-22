package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the InvalidArgumentErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidArgumentErrorResponseContent{}

// InvalidArgumentErrorResponseContent struct for InvalidArgumentErrorResponseContent
type InvalidArgumentErrorResponseContent struct {
	Code      InvalidArgumentErrorCode `json:"code"`
	RequestId string                   `json:"requestId"`
	Message   string                   `json:"message"`
}

type _InvalidArgumentErrorResponseContent InvalidArgumentErrorResponseContent

// NewInvalidArgumentErrorResponseContent instantiates a new InvalidArgumentErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidArgumentErrorResponseContent(code InvalidArgumentErrorCode, requestId string, message string) *InvalidArgumentErrorResponseContent {
	this := InvalidArgumentErrorResponseContent{}
	this.Code = code
	this.RequestId = requestId
	this.Message = message
	return &this
}

// NewInvalidArgumentErrorResponseContentWithDefaults instantiates a new InvalidArgumentErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidArgumentErrorResponseContentWithDefaults() *InvalidArgumentErrorResponseContent {
	this := InvalidArgumentErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *InvalidArgumentErrorResponseContent) GetCode() InvalidArgumentErrorCode {
	if o == nil {
		var ret InvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *InvalidArgumentErrorResponseContent) GetCodeOk() (*InvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *InvalidArgumentErrorResponseContent) SetCode(v InvalidArgumentErrorCode) {
	o.Code = v
}

// GetRequestId returns the RequestId field value
func (o *InvalidArgumentErrorResponseContent) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *InvalidArgumentErrorResponseContent) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *InvalidArgumentErrorResponseContent) SetRequestId(v string) {
	o.RequestId = v
}

// GetMessage returns the Message field value
func (o *InvalidArgumentErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InvalidArgumentErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InvalidArgumentErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o InvalidArgumentErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["requestId"] = o.RequestId
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableInvalidArgumentErrorResponseContent struct {
	value *InvalidArgumentErrorResponseContent
	isSet bool
}

func (v NullableInvalidArgumentErrorResponseContent) Get() *InvalidArgumentErrorResponseContent {
	return v.value
}

func (v *NullableInvalidArgumentErrorResponseContent) Set(val *InvalidArgumentErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidArgumentErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidArgumentErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidArgumentErrorResponseContent(val *InvalidArgumentErrorResponseContent) *NullableInvalidArgumentErrorResponseContent {
	return &NullableInvalidArgumentErrorResponseContent{value: val, isSet: true}
}

func (v NullableInvalidArgumentErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvalidArgumentErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
