package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the InternalServerErrorExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalServerErrorExceptionResponseContent{}

// InternalServerErrorExceptionResponseContent struct for InternalServerErrorExceptionResponseContent
type InternalServerErrorExceptionResponseContent struct {
	// Programmatic status code.
	Code *float32 `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewInternalServerErrorExceptionResponseContent instantiates a new InternalServerErrorExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalServerErrorExceptionResponseContent() *InternalServerErrorExceptionResponseContent {
	this := InternalServerErrorExceptionResponseContent{}
	return &this
}

// NewInternalServerErrorExceptionResponseContentWithDefaults instantiates a new InternalServerErrorExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalServerErrorExceptionResponseContentWithDefaults() *InternalServerErrorExceptionResponseContent {
	this := InternalServerErrorExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InternalServerErrorExceptionResponseContent) GetCode() float32 {
	if o == nil || IsNil(o.Code) {
		var ret float32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalServerErrorExceptionResponseContent) GetCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InternalServerErrorExceptionResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given float32 and assigns it to the Code field.
func (o *InternalServerErrorExceptionResponseContent) SetCode(v float32) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InternalServerErrorExceptionResponseContent) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalServerErrorExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InternalServerErrorExceptionResponseContent) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *InternalServerErrorExceptionResponseContent) SetDetails(v string) {
	o.Details = &v
}

func (o InternalServerErrorExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableInternalServerErrorExceptionResponseContent struct {
	value *InternalServerErrorExceptionResponseContent
	isSet bool
}

func (v NullableInternalServerErrorExceptionResponseContent) Get() *InternalServerErrorExceptionResponseContent {
	return v.value
}

func (v *NullableInternalServerErrorExceptionResponseContent) Set(val *InternalServerErrorExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalServerErrorExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalServerErrorExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalServerErrorExceptionResponseContent(val *InternalServerErrorExceptionResponseContent) *NullableInternalServerErrorExceptionResponseContent {
	return &NullableInternalServerErrorExceptionResponseContent{value: val, isSet: true}
}

func (v NullableInternalServerErrorExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInternalServerErrorExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
