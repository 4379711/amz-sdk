package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the ExcessiveBatchSizeExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExcessiveBatchSizeExceptionResponseContent{}

// ExcessiveBatchSizeExceptionResponseContent struct for ExcessiveBatchSizeExceptionResponseContent
type ExcessiveBatchSizeExceptionResponseContent struct {
	// Programmatic status code.
	Code *float32 `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewExcessiveBatchSizeExceptionResponseContent instantiates a new ExcessiveBatchSizeExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExcessiveBatchSizeExceptionResponseContent() *ExcessiveBatchSizeExceptionResponseContent {
	this := ExcessiveBatchSizeExceptionResponseContent{}
	return &this
}

// NewExcessiveBatchSizeExceptionResponseContentWithDefaults instantiates a new ExcessiveBatchSizeExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExcessiveBatchSizeExceptionResponseContentWithDefaults() *ExcessiveBatchSizeExceptionResponseContent {
	this := ExcessiveBatchSizeExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ExcessiveBatchSizeExceptionResponseContent) GetCode() float32 {
	if o == nil || IsNil(o.Code) {
		var ret float32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExcessiveBatchSizeExceptionResponseContent) GetCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ExcessiveBatchSizeExceptionResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given float32 and assigns it to the Code field.
func (o *ExcessiveBatchSizeExceptionResponseContent) SetCode(v float32) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ExcessiveBatchSizeExceptionResponseContent) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExcessiveBatchSizeExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ExcessiveBatchSizeExceptionResponseContent) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ExcessiveBatchSizeExceptionResponseContent) SetDetails(v string) {
	o.Details = &v
}

func (o ExcessiveBatchSizeExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableExcessiveBatchSizeExceptionResponseContent struct {
	value *ExcessiveBatchSizeExceptionResponseContent
	isSet bool
}

func (v NullableExcessiveBatchSizeExceptionResponseContent) Get() *ExcessiveBatchSizeExceptionResponseContent {
	return v.value
}

func (v *NullableExcessiveBatchSizeExceptionResponseContent) Set(val *ExcessiveBatchSizeExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableExcessiveBatchSizeExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableExcessiveBatchSizeExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExcessiveBatchSizeExceptionResponseContent(val *ExcessiveBatchSizeExceptionResponseContent) *NullableExcessiveBatchSizeExceptionResponseContent {
	return &NullableExcessiveBatchSizeExceptionResponseContent{value: val, isSet: true}
}

func (v NullableExcessiveBatchSizeExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableExcessiveBatchSizeExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
