package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetError{}

// BudgetError struct for BudgetError
type BudgetError struct {
	Reason     string     `json:"reason"`
	Cause      ErrorCause `json:"cause"`
	UpperLimit *string    `json:"upperLimit,omitempty"`
	LowerLimit *string    `json:"lowerLimit,omitempty"`
	// Human readable error message.
	Message string `json:"message"`
}

type _BudgetError BudgetError

// NewBudgetError instantiates a new BudgetError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetError(reason string, cause ErrorCause, message string) *BudgetError {
	this := BudgetError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewBudgetErrorWithDefaults instantiates a new BudgetError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetErrorWithDefaults() *BudgetError {
	this := BudgetError{}
	return &this
}

// GetReason returns the Reason field value
func (o *BudgetError) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *BudgetError) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *BudgetError) SetReason(v string) {
	o.Reason = v
}

// GetCause returns the Cause field value
func (o *BudgetError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *BudgetError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *BudgetError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetUpperLimit returns the UpperLimit field value if set, zero value otherwise.
func (o *BudgetError) GetUpperLimit() string {
	if o == nil || IsNil(o.UpperLimit) {
		var ret string
		return ret
	}
	return *o.UpperLimit
}

// GetUpperLimitOk returns a tuple with the UpperLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetError) GetUpperLimitOk() (*string, bool) {
	if o == nil || IsNil(o.UpperLimit) {
		return nil, false
	}
	return o.UpperLimit, true
}

// HasUpperLimit returns a boolean if a field has been set.
func (o *BudgetError) HasUpperLimit() bool {
	if o != nil && !IsNil(o.UpperLimit) {
		return true
	}

	return false
}

// SetUpperLimit gets a reference to the given string and assigns it to the UpperLimit field.
func (o *BudgetError) SetUpperLimit(v string) {
	o.UpperLimit = &v
}

// GetLowerLimit returns the LowerLimit field value if set, zero value otherwise.
func (o *BudgetError) GetLowerLimit() string {
	if o == nil || IsNil(o.LowerLimit) {
		var ret string
		return ret
	}
	return *o.LowerLimit
}

// GetLowerLimitOk returns a tuple with the LowerLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetError) GetLowerLimitOk() (*string, bool) {
	if o == nil || IsNil(o.LowerLimit) {
		return nil, false
	}
	return o.LowerLimit, true
}

// HasLowerLimit returns a boolean if a field has been set.
func (o *BudgetError) HasLowerLimit() bool {
	if o != nil && !IsNil(o.LowerLimit) {
		return true
	}

	return false
}

// SetLowerLimit gets a reference to the given string and assigns it to the LowerLimit field.
func (o *BudgetError) SetLowerLimit(v string) {
	o.LowerLimit = &v
}

// GetMessage returns the Message field value
func (o *BudgetError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BudgetError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BudgetError) SetMessage(v string) {
	o.Message = v
}

func (o BudgetError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
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

type NullableBudgetError struct {
	value *BudgetError
	isSet bool
}

func (v NullableBudgetError) Get() *BudgetError {
	return v.value
}

func (v *NullableBudgetError) Set(val *BudgetError) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetError) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetError(val *BudgetError) *NullableBudgetError {
	return &NullableBudgetError{value: val, isSet: true}
}

func (v NullableBudgetError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
