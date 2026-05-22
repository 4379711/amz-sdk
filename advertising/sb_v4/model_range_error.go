package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the RangeError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeError{}

// RangeError Errors related to range constraints violations.
type RangeError struct {
	Reason string `json:"reason"`
	// Allowed values.
	Allowed []string   `json:"allowed,omitempty"`
	Cause   ErrorCause `json:"cause"`
	// Optional upper limit.
	UpperLimit *string `json:"upperLimit,omitempty"`
	// Optional lower limit.
	LowerLimit *string `json:"lowerLimit,omitempty"`
	// Human readable error message.
	Message string `json:"message"`
}

type _RangeError RangeError

// NewRangeError instantiates a new RangeError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeError(reason string, cause ErrorCause, message string) *RangeError {
	this := RangeError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewRangeErrorWithDefaults instantiates a new RangeError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeErrorWithDefaults() *RangeError {
	this := RangeError{}
	return &this
}

// GetReason returns the Reason field value
func (o *RangeError) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *RangeError) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *RangeError) SetReason(v string) {
	o.Reason = v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *RangeError) GetAllowed() []string {
	if o == nil || IsNil(o.Allowed) {
		var ret []string
		return ret
	}
	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeError) GetAllowedOk() ([]string, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *RangeError) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given []string and assigns it to the Allowed field.
func (o *RangeError) SetAllowed(v []string) {
	o.Allowed = v
}

// GetCause returns the Cause field value
func (o *RangeError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *RangeError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *RangeError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetUpperLimit returns the UpperLimit field value if set, zero value otherwise.
func (o *RangeError) GetUpperLimit() string {
	if o == nil || IsNil(o.UpperLimit) {
		var ret string
		return ret
	}
	return *o.UpperLimit
}

// GetUpperLimitOk returns a tuple with the UpperLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeError) GetUpperLimitOk() (*string, bool) {
	if o == nil || IsNil(o.UpperLimit) {
		return nil, false
	}
	return o.UpperLimit, true
}

// HasUpperLimit returns a boolean if a field has been set.
func (o *RangeError) HasUpperLimit() bool {
	if o != nil && !IsNil(o.UpperLimit) {
		return true
	}

	return false
}

// SetUpperLimit gets a reference to the given string and assigns it to the UpperLimit field.
func (o *RangeError) SetUpperLimit(v string) {
	o.UpperLimit = &v
}

// GetLowerLimit returns the LowerLimit field value if set, zero value otherwise.
func (o *RangeError) GetLowerLimit() string {
	if o == nil || IsNil(o.LowerLimit) {
		var ret string
		return ret
	}
	return *o.LowerLimit
}

// GetLowerLimitOk returns a tuple with the LowerLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeError) GetLowerLimitOk() (*string, bool) {
	if o == nil || IsNil(o.LowerLimit) {
		return nil, false
	}
	return o.LowerLimit, true
}

// HasLowerLimit returns a boolean if a field has been set.
func (o *RangeError) HasLowerLimit() bool {
	if o != nil && !IsNil(o.LowerLimit) {
		return true
	}

	return false
}

// SetLowerLimit gets a reference to the given string and assigns it to the LowerLimit field.
func (o *RangeError) SetLowerLimit(v string) {
	o.LowerLimit = &v
}

// GetMessage returns the Message field value
func (o *RangeError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RangeError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RangeError) SetMessage(v string) {
	o.Message = v
}

func (o RangeError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	toSerialize["cause"] = o.Cause
	if !IsNil(o.UpperLimit) {
		toSerialize["upperLimit"] = o.UpperLimit
	}
	if !IsNil(o.LowerLimit) {
		toSerialize["lowerLimit"] = o.LowerLimit
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableRangeError struct {
	value *RangeError
	isSet bool
}

func (v NullableRangeError) Get() *RangeError {
	return v.value
}

func (v *NullableRangeError) Set(val *RangeError) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeError) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeError(val *RangeError) *NullableRangeError {
	return &NullableRangeError{value: val, isSet: true}
}

func (v NullableRangeError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRangeError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
