package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ErrorDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorDetails{}

// ErrorDetails struct for ErrorDetails
type ErrorDetails struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The index of the image task in the array from the request body
	Index     *float32 `json:"index,omitempty"`
	ErrorCode *string  `json:"errorCode,omitempty"`
}

// NewErrorDetails instantiates a new ErrorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDetails() *ErrorDetails {
	this := ErrorDetails{}
	return &this
}

// NewErrorDetailsWithDefaults instantiates a new ErrorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDetailsWithDefaults() *ErrorDetails {
	this := ErrorDetails{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ErrorDetails) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetails) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ErrorDetails) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ErrorDetails) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ErrorDetails) GetIndex() float32 {
	if o == nil || IsNil(o.Index) {
		var ret float32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetails) GetIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ErrorDetails) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given float32 and assigns it to the Index field.
func (o *ErrorDetails) SetIndex(v float32) {
	o.Index = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ErrorDetails) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetails) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ErrorDetails) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ErrorDetails) SetErrorCode(v string) {
	o.ErrorCode = &v
}

func (o ErrorDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	return toSerialize, nil
}

type NullableErrorDetails struct {
	value *ErrorDetails
	isSet bool
}

func (v NullableErrorDetails) Get() *ErrorDetails {
	return v.value
}

func (v *NullableErrorDetails) Set(val *ErrorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetails(val *ErrorDetails) *NullableErrorDetails {
	return &NullableErrorDetails{value: val, isSet: true}
}

func (v NullableErrorDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableErrorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
