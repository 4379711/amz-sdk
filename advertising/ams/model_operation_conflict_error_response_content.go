package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the OperationConflictErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationConflictErrorResponseContent{}

// OperationConflictErrorResponseContent struct for OperationConflictErrorResponseContent
type OperationConflictErrorResponseContent struct {
	Code    *string `json:"code,omitempty"`
	Message string  `json:"message"`
}

type _OperationConflictErrorResponseContent OperationConflictErrorResponseContent

// NewOperationConflictErrorResponseContent instantiates a new OperationConflictErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationConflictErrorResponseContent(message string) *OperationConflictErrorResponseContent {
	this := OperationConflictErrorResponseContent{}
	this.Message = message
	return &this
}

// NewOperationConflictErrorResponseContentWithDefaults instantiates a new OperationConflictErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationConflictErrorResponseContentWithDefaults() *OperationConflictErrorResponseContent {
	this := OperationConflictErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OperationConflictErrorResponseContent) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationConflictErrorResponseContent) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OperationConflictErrorResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OperationConflictErrorResponseContent) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value
func (o *OperationConflictErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OperationConflictErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OperationConflictErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o OperationConflictErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableOperationConflictErrorResponseContent struct {
	value *OperationConflictErrorResponseContent
	isSet bool
}

func (v NullableOperationConflictErrorResponseContent) Get() *OperationConflictErrorResponseContent {
	return v.value
}

func (v *NullableOperationConflictErrorResponseContent) Set(val *OperationConflictErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationConflictErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationConflictErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationConflictErrorResponseContent(val *OperationConflictErrorResponseContent) *NullableOperationConflictErrorResponseContent {
	return &NullableOperationConflictErrorResponseContent{value: val, isSet: true}
}

func (v NullableOperationConflictErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOperationConflictErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
