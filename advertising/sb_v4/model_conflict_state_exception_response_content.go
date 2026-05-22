package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ConflictStateExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConflictStateExceptionResponseContent{}

// ConflictStateExceptionResponseContent struct for ConflictStateExceptionResponseContent
type ConflictStateExceptionResponseContent struct {
	Code    ConflictStateErrorCode `json:"code"`
	Message string                 `json:"message"`
}

type _ConflictStateExceptionResponseContent ConflictStateExceptionResponseContent

// NewConflictStateExceptionResponseContent instantiates a new ConflictStateExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConflictStateExceptionResponseContent(code ConflictStateErrorCode, message string) *ConflictStateExceptionResponseContent {
	this := ConflictStateExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewConflictStateExceptionResponseContentWithDefaults instantiates a new ConflictStateExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConflictStateExceptionResponseContentWithDefaults() *ConflictStateExceptionResponseContent {
	this := ConflictStateExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *ConflictStateExceptionResponseContent) GetCode() ConflictStateErrorCode {
	if o == nil {
		var ret ConflictStateErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ConflictStateExceptionResponseContent) GetCodeOk() (*ConflictStateErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ConflictStateExceptionResponseContent) SetCode(v ConflictStateErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ConflictStateExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ConflictStateExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ConflictStateExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o ConflictStateExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableConflictStateExceptionResponseContent struct {
	value *ConflictStateExceptionResponseContent
	isSet bool
}

func (v NullableConflictStateExceptionResponseContent) Get() *ConflictStateExceptionResponseContent {
	return v.value
}

func (v *NullableConflictStateExceptionResponseContent) Set(val *ConflictStateExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableConflictStateExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableConflictStateExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflictStateExceptionResponseContent(val *ConflictStateExceptionResponseContent) *NullableConflictStateExceptionResponseContent {
	return &NullableConflictStateExceptionResponseContent{value: val, isSet: true}
}

func (v NullableConflictStateExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConflictStateExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
