package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the InvalidArgumentExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidArgumentExceptionResponseContent{}

// InvalidArgumentExceptionResponseContent struct for InvalidArgumentExceptionResponseContent
type InvalidArgumentExceptionResponseContent struct {
	Code InvalidArgumentErrorCode `json:"code"`
	// Human readable error message.
	Message string                 `json:"message"`
	Errors  []InvalidArgumentError `json:"errors,omitempty"`
}

type _InvalidArgumentExceptionResponseContent InvalidArgumentExceptionResponseContent

// NewInvalidArgumentExceptionResponseContent instantiates a new InvalidArgumentExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidArgumentExceptionResponseContent(code InvalidArgumentErrorCode, message string) *InvalidArgumentExceptionResponseContent {
	this := InvalidArgumentExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewInvalidArgumentExceptionResponseContentWithDefaults instantiates a new InvalidArgumentExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidArgumentExceptionResponseContentWithDefaults() *InvalidArgumentExceptionResponseContent {
	this := InvalidArgumentExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *InvalidArgumentExceptionResponseContent) GetCode() InvalidArgumentErrorCode {
	if o == nil {
		var ret InvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *InvalidArgumentExceptionResponseContent) GetCodeOk() (*InvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *InvalidArgumentExceptionResponseContent) SetCode(v InvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *InvalidArgumentExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InvalidArgumentExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InvalidArgumentExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InvalidArgumentExceptionResponseContent) GetErrors() []InvalidArgumentError {
	if o == nil || IsNil(o.Errors) {
		var ret []InvalidArgumentError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidArgumentExceptionResponseContent) GetErrorsOk() ([]InvalidArgumentError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InvalidArgumentExceptionResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []InvalidArgumentError and assigns it to the Errors field.
func (o *InvalidArgumentExceptionResponseContent) SetErrors(v []InvalidArgumentError) {
	o.Errors = v
}

func (o InvalidArgumentExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableInvalidArgumentExceptionResponseContent struct {
	value *InvalidArgumentExceptionResponseContent
	isSet bool
}

func (v NullableInvalidArgumentExceptionResponseContent) Get() *InvalidArgumentExceptionResponseContent {
	return v.value
}

func (v *NullableInvalidArgumentExceptionResponseContent) Set(val *InvalidArgumentExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidArgumentExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidArgumentExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidArgumentExceptionResponseContent(val *InvalidArgumentExceptionResponseContent) *NullableInvalidArgumentExceptionResponseContent {
	return &NullableInvalidArgumentExceptionResponseContent{value: val, isSet: true}
}

func (v NullableInvalidArgumentExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvalidArgumentExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
