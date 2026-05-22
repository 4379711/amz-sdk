package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the ResourceNotFoundErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceNotFoundErrorResponseContent{}

// ResourceNotFoundErrorResponseContent struct for ResourceNotFoundErrorResponseContent
type ResourceNotFoundErrorResponseContent struct {
	Code    *string `json:"code,omitempty"`
	Message string  `json:"message"`
}

type _ResourceNotFoundErrorResponseContent ResourceNotFoundErrorResponseContent

// NewResourceNotFoundErrorResponseContent instantiates a new ResourceNotFoundErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceNotFoundErrorResponseContent(message string) *ResourceNotFoundErrorResponseContent {
	this := ResourceNotFoundErrorResponseContent{}
	this.Message = message
	return &this
}

// NewResourceNotFoundErrorResponseContentWithDefaults instantiates a new ResourceNotFoundErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceNotFoundErrorResponseContentWithDefaults() *ResourceNotFoundErrorResponseContent {
	this := ResourceNotFoundErrorResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ResourceNotFoundErrorResponseContent) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceNotFoundErrorResponseContent) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ResourceNotFoundErrorResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ResourceNotFoundErrorResponseContent) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value
func (o *ResourceNotFoundErrorResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ResourceNotFoundErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ResourceNotFoundErrorResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o ResourceNotFoundErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableResourceNotFoundErrorResponseContent struct {
	value *ResourceNotFoundErrorResponseContent
	isSet bool
}

func (v NullableResourceNotFoundErrorResponseContent) Get() *ResourceNotFoundErrorResponseContent {
	return v.value
}

func (v *NullableResourceNotFoundErrorResponseContent) Set(val *ResourceNotFoundErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceNotFoundErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceNotFoundErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceNotFoundErrorResponseContent(val *ResourceNotFoundErrorResponseContent) *NullableResourceNotFoundErrorResponseContent {
	return &NullableResourceNotFoundErrorResponseContent{value: val, isSet: true}
}

func (v NullableResourceNotFoundErrorResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableResourceNotFoundErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
