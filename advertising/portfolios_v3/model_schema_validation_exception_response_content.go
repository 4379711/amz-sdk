package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SchemaValidationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaValidationExceptionResponseContent{}

// SchemaValidationExceptionResponseContent struct for SchemaValidationExceptionResponseContent
type SchemaValidationExceptionResponseContent struct {
	Code InvalidArgumentErrorCode `json:"code"`
	// Human readable error message
	Message string `json:"message"`
}

type _SchemaValidationExceptionResponseContent SchemaValidationExceptionResponseContent

// NewSchemaValidationExceptionResponseContent instantiates a new SchemaValidationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaValidationExceptionResponseContent(code InvalidArgumentErrorCode, message string) *SchemaValidationExceptionResponseContent {
	this := SchemaValidationExceptionResponseContent{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSchemaValidationExceptionResponseContentWithDefaults instantiates a new SchemaValidationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaValidationExceptionResponseContentWithDefaults() *SchemaValidationExceptionResponseContent {
	this := SchemaValidationExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *SchemaValidationExceptionResponseContent) GetCode() InvalidArgumentErrorCode {
	if o == nil {
		var ret InvalidArgumentErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SchemaValidationExceptionResponseContent) GetCodeOk() (*InvalidArgumentErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SchemaValidationExceptionResponseContent) SetCode(v InvalidArgumentErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SchemaValidationExceptionResponseContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SchemaValidationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SchemaValidationExceptionResponseContent) SetMessage(v string) {
	o.Message = v
}

func (o SchemaValidationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSchemaValidationExceptionResponseContent struct {
	value *SchemaValidationExceptionResponseContent
	isSet bool
}

func (v NullableSchemaValidationExceptionResponseContent) Get() *SchemaValidationExceptionResponseContent {
	return v.value
}

func (v *NullableSchemaValidationExceptionResponseContent) Set(val *SchemaValidationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaValidationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaValidationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaValidationExceptionResponseContent(val *SchemaValidationExceptionResponseContent) *NullableSchemaValidationExceptionResponseContent {
	return &NullableSchemaValidationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableSchemaValidationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSchemaValidationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
