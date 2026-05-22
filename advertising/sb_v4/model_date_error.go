package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DateError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateError{}

// DateError Errors related to dates.
type DateError struct {
	// Exact error reason..
	Reason string     `json:"reason"`
	Cause  ErrorCause `json:"cause"`
	// Human readable error message.
	Message string `json:"message"`
}

type _DateError DateError

// NewDateError instantiates a new DateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateError(reason string, cause ErrorCause, message string) *DateError {
	this := DateError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewDateErrorWithDefaults instantiates a new DateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateErrorWithDefaults() *DateError {
	this := DateError{}
	return &this
}

// GetReason returns the Reason field value
func (o *DateError) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *DateError) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *DateError) SetReason(v string) {
	o.Reason = v
}

// GetCause returns the Cause field value
func (o *DateError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *DateError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *DateError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetMessage returns the Message field value
func (o *DateError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DateError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *DateError) SetMessage(v string) {
	o.Message = v
}

func (o DateError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	toSerialize["cause"] = o.Cause
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableDateError struct {
	value *DateError
	isSet bool
}

func (v NullableDateError) Get() *DateError {
	return v.value
}

func (v *NullableDateError) Set(val *DateError) {
	v.value = val
	v.isSet = true
}

func (v NullableDateError) IsSet() bool {
	return v.isSet
}

func (v *NullableDateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateError(val *DateError) *NullableDateError {
	return &NullableDateError{value: val, isSet: true}
}

func (v NullableDateError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
