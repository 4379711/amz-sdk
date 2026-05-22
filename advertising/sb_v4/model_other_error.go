package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the OtherError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OtherError{}

// OtherError Errors not related to any of the other error types.
type OtherError struct {
	Reason string     `json:"reason"`
	Cause  ErrorCause `json:"cause"`
	// Human readable error message.
	Message string `json:"message"`
}

type _OtherError OtherError

// NewOtherError instantiates a new OtherError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtherError(reason string, cause ErrorCause, message string) *OtherError {
	this := OtherError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewOtherErrorWithDefaults instantiates a new OtherError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtherErrorWithDefaults() *OtherError {
	this := OtherError{}
	return &this
}

// GetReason returns the Reason field value
func (o *OtherError) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *OtherError) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *OtherError) SetReason(v string) {
	o.Reason = v
}

// GetCause returns the Cause field value
func (o *OtherError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *OtherError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *OtherError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetMessage returns the Message field value
func (o *OtherError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OtherError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OtherError) SetMessage(v string) {
	o.Message = v
}

func (o OtherError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	toSerialize["cause"] = o.Cause
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableOtherError struct {
	value *OtherError
	isSet bool
}

func (v NullableOtherError) Get() *OtherError {
	return v.value
}

func (v *NullableOtherError) Set(val *OtherError) {
	v.value = val
	v.isSet = true
}

func (v NullableOtherError) IsSet() bool {
	return v.isSet
}

func (v *NullableOtherError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtherError(val *OtherError) *NullableOtherError {
	return &NullableOtherError{value: val, isSet: true}
}

func (v NullableOtherError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOtherError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
