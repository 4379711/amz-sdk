package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UnsupportedMediaTypeExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnsupportedMediaTypeExceptionResponseContent{}

// UnsupportedMediaTypeExceptionResponseContent struct for UnsupportedMediaTypeExceptionResponseContent
type UnsupportedMediaTypeExceptionResponseContent struct {
	// A human-readable description of the enumerated response code in the `code` field.
	Code string `json:"code"`
	// An enumerated response code.
	Details string `json:"details"`
}

type _UnsupportedMediaTypeExceptionResponseContent UnsupportedMediaTypeExceptionResponseContent

// NewUnsupportedMediaTypeExceptionResponseContent instantiates a new UnsupportedMediaTypeExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnsupportedMediaTypeExceptionResponseContent(code string, details string) *UnsupportedMediaTypeExceptionResponseContent {
	this := UnsupportedMediaTypeExceptionResponseContent{}
	this.Code = code
	this.Details = details
	return &this
}

// NewUnsupportedMediaTypeExceptionResponseContentWithDefaults instantiates a new UnsupportedMediaTypeExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnsupportedMediaTypeExceptionResponseContentWithDefaults() *UnsupportedMediaTypeExceptionResponseContent {
	this := UnsupportedMediaTypeExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value
func (o *UnsupportedMediaTypeExceptionResponseContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *UnsupportedMediaTypeExceptionResponseContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *UnsupportedMediaTypeExceptionResponseContent) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *UnsupportedMediaTypeExceptionResponseContent) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *UnsupportedMediaTypeExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *UnsupportedMediaTypeExceptionResponseContent) SetDetails(v string) {
	o.Details = v
}

func (o UnsupportedMediaTypeExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableUnsupportedMediaTypeExceptionResponseContent struct {
	value *UnsupportedMediaTypeExceptionResponseContent
	isSet bool
}

func (v NullableUnsupportedMediaTypeExceptionResponseContent) Get() *UnsupportedMediaTypeExceptionResponseContent {
	return v.value
}

func (v *NullableUnsupportedMediaTypeExceptionResponseContent) Set(val *UnsupportedMediaTypeExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUnsupportedMediaTypeExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUnsupportedMediaTypeExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnsupportedMediaTypeExceptionResponseContent(val *UnsupportedMediaTypeExceptionResponseContent) *NullableUnsupportedMediaTypeExceptionResponseContent {
	return &NullableUnsupportedMediaTypeExceptionResponseContent{value: val, isSet: true}
}

func (v NullableUnsupportedMediaTypeExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnsupportedMediaTypeExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
