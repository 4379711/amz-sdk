package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BillingError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingError{}

// BillingError Errors related to billing.
type BillingError struct {
	// Exact error reason.
	Reason string     `json:"reason"`
	Cause  ErrorCause `json:"cause"`
	// Human readable error message.
	Message string `json:"message"`
}

type _BillingError BillingError

// NewBillingError instantiates a new BillingError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingError(reason string, cause ErrorCause, message string) *BillingError {
	this := BillingError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewBillingErrorWithDefaults instantiates a new BillingError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingErrorWithDefaults() *BillingError {
	this := BillingError{}
	return &this
}

// GetReason returns the Reason field value
func (o *BillingError) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *BillingError) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *BillingError) SetReason(v string) {
	o.Reason = v
}

// GetCause returns the Cause field value
func (o *BillingError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *BillingError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *BillingError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetMessage returns the Message field value
func (o *BillingError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BillingError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BillingError) SetMessage(v string) {
	o.Message = v
}

func (o BillingError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	toSerialize["cause"] = o.Cause
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableBillingError struct {
	value *BillingError
	isSet bool
}

func (v NullableBillingError) Get() *BillingError {
	return v.value
}

func (v *NullableBillingError) Set(val *BillingError) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingError) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingError(val *BillingError) *NullableBillingError {
	return &NullableBillingError{value: val, isSet: true}
}

func (v NullableBillingError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBillingError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
