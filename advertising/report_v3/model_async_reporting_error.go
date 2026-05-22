package report_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AsyncReportingError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncReportingError{}

// AsyncReportingError The Error Response.
type AsyncReportingError struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Detail *string `json:"detail,omitempty"`
}

// NewAsyncReportingError instantiates a new AsyncReportingError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncReportingError() *AsyncReportingError {
	this := AsyncReportingError{}
	return &this
}

// NewAsyncReportingErrorWithDefaults instantiates a new AsyncReportingError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncReportingErrorWithDefaults() *AsyncReportingError {
	this := AsyncReportingError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AsyncReportingError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncReportingError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AsyncReportingError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AsyncReportingError) SetCode(v string) {
	o.Code = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *AsyncReportingError) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncReportingError) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *AsyncReportingError) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *AsyncReportingError) SetDetail(v string) {
	o.Detail = &v
}

func (o AsyncReportingError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableAsyncReportingError struct {
	value *AsyncReportingError
	isSet bool
}

func (v NullableAsyncReportingError) Get() *AsyncReportingError {
	return v.value
}

func (v *NullableAsyncReportingError) Set(val *AsyncReportingError) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncReportingError) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncReportingError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncReportingError(val *AsyncReportingError) *NullableAsyncReportingError {
	return &NullableAsyncReportingError{value: val, isSet: true}
}

func (v NullableAsyncReportingError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsyncReportingError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
