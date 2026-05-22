package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the BadRequestExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BadRequestExceptionResponseContent{}

// BadRequestExceptionResponseContent struct for BadRequestExceptionResponseContent
type BadRequestExceptionResponseContent struct {
	// Programmatic status code.
	Code *float32 `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewBadRequestExceptionResponseContent instantiates a new BadRequestExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestExceptionResponseContent() *BadRequestExceptionResponseContent {
	this := BadRequestExceptionResponseContent{}
	return &this
}

// NewBadRequestExceptionResponseContentWithDefaults instantiates a new BadRequestExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestExceptionResponseContentWithDefaults() *BadRequestExceptionResponseContent {
	this := BadRequestExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BadRequestExceptionResponseContent) GetCode() float32 {
	if o == nil || IsNil(o.Code) {
		var ret float32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestExceptionResponseContent) GetCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BadRequestExceptionResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given float32 and assigns it to the Code field.
func (o *BadRequestExceptionResponseContent) SetCode(v float32) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *BadRequestExceptionResponseContent) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *BadRequestExceptionResponseContent) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *BadRequestExceptionResponseContent) SetDetails(v string) {
	o.Details = &v
}

func (o BadRequestExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableBadRequestExceptionResponseContent struct {
	value *BadRequestExceptionResponseContent
	isSet bool
}

func (v NullableBadRequestExceptionResponseContent) Get() *BadRequestExceptionResponseContent {
	return v.value
}

func (v *NullableBadRequestExceptionResponseContent) Set(val *BadRequestExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestExceptionResponseContent(val *BadRequestExceptionResponseContent) *NullableBadRequestExceptionResponseContent {
	return &NullableBadRequestExceptionResponseContent{value: val, isSet: true}
}

func (v NullableBadRequestExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBadRequestExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
