package product_fees_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the FeesEstimateError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeesEstimateError{}

// FeesEstimateError An unexpected error occurred during this operation.
type FeesEstimateError struct {
	// An error type, identifying either the receiver or the sender as the originator of the error.
	Type string `json:"Type"`
	// An error code that identifies the type of error that occurred.
	Code string `json:"Code"`
	// A message that describes the error condition.
	Message string `json:"Message"`
	// Additional information that can help the caller understand or fix the issue.
	Detail []map[string]interface{} `json:"Detail"`
}

type _FeesEstimateError FeesEstimateError

// NewFeesEstimateError instantiates a new FeesEstimateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeesEstimateError(type_ string, code string, message string, detail []map[string]interface{}) *FeesEstimateError {
	this := FeesEstimateError{}
	this.Type = type_
	this.Code = code
	this.Message = message
	this.Detail = detail
	return &this
}

// NewFeesEstimateErrorWithDefaults instantiates a new FeesEstimateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeesEstimateErrorWithDefaults() *FeesEstimateError {
	this := FeesEstimateError{}
	return &this
}

// GetType returns the Type field value
func (o *FeesEstimateError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FeesEstimateError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FeesEstimateError) SetType(v string) {
	o.Type = v
}

// GetCode returns the Code field value
func (o *FeesEstimateError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *FeesEstimateError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *FeesEstimateError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *FeesEstimateError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *FeesEstimateError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *FeesEstimateError) SetMessage(v string) {
	o.Message = v
}

// GetDetail returns the Detail field value
func (o *FeesEstimateError) GetDetail() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *FeesEstimateError) GetDetailOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail, true
}

// SetDetail sets field value
func (o *FeesEstimateError) SetDetail(v []map[string]interface{}) {
	o.Detail = v
}

func (o FeesEstimateError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Type"] = o.Type
	toSerialize["Code"] = o.Code
	toSerialize["Message"] = o.Message
	toSerialize["Detail"] = o.Detail
	return toSerialize, nil
}

type NullableFeesEstimateError struct {
	value *FeesEstimateError
	isSet bool
}

func (v NullableFeesEstimateError) Get() *FeesEstimateError {
	return v.value
}

func (v *NullableFeesEstimateError) Set(val *FeesEstimateError) {
	v.value = val
	v.isSet = true
}

func (v NullableFeesEstimateError) IsSet() bool {
	return v.isSet
}

func (v *NullableFeesEstimateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeesEstimateError(val *FeesEstimateError) *NullableFeesEstimateError {
	return &NullableFeesEstimateError{value: val, isSet: true}
}

func (v NullableFeesEstimateError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeesEstimateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
