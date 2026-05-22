package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AccessDeniedExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessDeniedExceptionResponseContent{}

// AccessDeniedExceptionResponseContent struct for AccessDeniedExceptionResponseContent
type AccessDeniedExceptionResponseContent struct {
	Code AccessDeniedErrorCode `json:"code"`
	// Human readable error message
	Message string `json:"message"`
}

type _AccessDeniedExceptionResponseContent AccessDeniedExceptionResponseContent

// NewAccessDeniedExceptionResponseContent instantiates a new AccessDeniedExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessDeniedExceptionResponseContent(code AccessDeniedErrorCode, message string) *AccessDeniedExceptionResponseContent {
	this := AccessDeniedExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAccessDeniedExceptionResponseContentWithDefaults instantiates a new AccessDeniedExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessDeniedExceptionResponseContentWithDefaults() *AccessDeniedExceptionResponseContent {
	this := AccessDeniedExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *AccessDeniedExceptionResponseContent) GetCode() AccessDeniedErrorCode {
	if o == nil {
		var ret AccessDeniedErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AccessDeniedExceptionResponseContent) GetCodeOk() (*AccessDeniedErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AccessDeniedExceptionResponseContent) SetCode(v AccessDeniedErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *AccessDeniedExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AccessDeniedExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AccessDeniedExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o AccessDeniedExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableAccessDeniedExceptionResponseContent struct {
	value *AccessDeniedExceptionResponseContent
	isSet bool
}

func (v NullableAccessDeniedExceptionResponseContent) Get() *AccessDeniedExceptionResponseContent {
	return v.value
}

func (v *NullableAccessDeniedExceptionResponseContent) Set(val *AccessDeniedExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDeniedExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDeniedExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDeniedExceptionResponseContent(val *AccessDeniedExceptionResponseContent) *NullableAccessDeniedExceptionResponseContent {
	return &NullableAccessDeniedExceptionResponseContent{value: val, isSet: true}
}

func (v NullableAccessDeniedExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccessDeniedExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
