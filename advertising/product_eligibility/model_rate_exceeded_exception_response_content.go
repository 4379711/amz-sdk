package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the RateExceededExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateExceededExceptionResponseContent{}

// RateExceededExceptionResponseContent struct for RateExceededExceptionResponseContent
type RateExceededExceptionResponseContent struct {
	// Programmatic status code.
	Code *float32 `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewRateExceededExceptionResponseContent instantiates a new RateExceededExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateExceededExceptionResponseContent() *RateExceededExceptionResponseContent {
	this := RateExceededExceptionResponseContent{}
	return &this
}

// NewRateExceededExceptionResponseContentWithDefaults instantiates a new RateExceededExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateExceededExceptionResponseContentWithDefaults() *RateExceededExceptionResponseContent {
	this := RateExceededExceptionResponseContent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RateExceededExceptionResponseContent) GetCode() float32 {
	if o == nil || IsNil(o.Code) {
		var ret float32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateExceededExceptionResponseContent) GetCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RateExceededExceptionResponseContent) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given float32 and assigns it to the Code field.
func (o *RateExceededExceptionResponseContent) SetCode(v float32) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RateExceededExceptionResponseContent) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateExceededExceptionResponseContent) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RateExceededExceptionResponseContent) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *RateExceededExceptionResponseContent) SetDetails(v string) {
	o.Details = &v
}

func (o RateExceededExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableRateExceededExceptionResponseContent struct {
	value *RateExceededExceptionResponseContent
	isSet bool
}

func (v NullableRateExceededExceptionResponseContent) Get() *RateExceededExceptionResponseContent {
	return v.value
}

func (v *NullableRateExceededExceptionResponseContent) Set(val *RateExceededExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableRateExceededExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableRateExceededExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateExceededExceptionResponseContent(val *RateExceededExceptionResponseContent) *NullableRateExceededExceptionResponseContent {
	return &NullableRateExceededExceptionResponseContent{value: val, isSet: true}
}

func (v NullableRateExceededExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRateExceededExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
