package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBudgetError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBudgetError{}

// SponsoredProductsBudgetError struct for SponsoredProductsBudgetError
type SponsoredProductsBudgetError struct {
	Reason     SponsoredProductsBudgetErrorReason `json:"reason"`
	Cause      *SponsoredProductsErrorCause       `json:"cause,omitempty"`
	UpperLimit *string                            `json:"upperLimit,omitempty"`
	LowerLimit *string                            `json:"lowerLimit,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsBudgetError SponsoredProductsBudgetError

// NewSponsoredProductsBudgetError instantiates a new SponsoredProductsBudgetError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBudgetError(reason SponsoredProductsBudgetErrorReason, message string) *SponsoredProductsBudgetError {
	this := SponsoredProductsBudgetError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsBudgetErrorWithDefaults instantiates a new SponsoredProductsBudgetError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBudgetErrorWithDefaults() *SponsoredProductsBudgetError {
	this := SponsoredProductsBudgetError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsBudgetError) GetReason() SponsoredProductsBudgetErrorReason {
	if o == nil {
		var ret SponsoredProductsBudgetErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBudgetError) GetReasonOk() (*SponsoredProductsBudgetErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsBudgetError) SetReason(v SponsoredProductsBudgetErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsBudgetError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBudgetError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsBudgetError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsBudgetError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetUpperLimit returns the UpperLimit field value if set, zero value otherwise.
func (o *SponsoredProductsBudgetError) GetUpperLimit() string {
	if o == nil || IsNil(o.UpperLimit) {
		var ret string
		return ret
	}
	return *o.UpperLimit
}

// GetUpperLimitOk returns a tuple with the UpperLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBudgetError) GetUpperLimitOk() (*string, bool) {
	if o == nil || IsNil(o.UpperLimit) {
		return nil, false
	}
	return o.UpperLimit, true
}

// HasUpperLimit returns a boolean if a field has been set.
func (o *SponsoredProductsBudgetError) HasUpperLimit() bool {
	if o != nil && !IsNil(o.UpperLimit) {
		return true
	}

	return false
}

// SetUpperLimit gets a reference to the given string and assigns it to the UpperLimit field.
func (o *SponsoredProductsBudgetError) SetUpperLimit(v string) {
	o.UpperLimit = &v
}

// GetLowerLimit returns the LowerLimit field value if set, zero value otherwise.
func (o *SponsoredProductsBudgetError) GetLowerLimit() string {
	if o == nil || IsNil(o.LowerLimit) {
		var ret string
		return ret
	}
	return *o.LowerLimit
}

// GetLowerLimitOk returns a tuple with the LowerLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBudgetError) GetLowerLimitOk() (*string, bool) {
	if o == nil || IsNil(o.LowerLimit) {
		return nil, false
	}
	return o.LowerLimit, true
}

// HasLowerLimit returns a boolean if a field has been set.
func (o *SponsoredProductsBudgetError) HasLowerLimit() bool {
	if o != nil && !IsNil(o.LowerLimit) {
		return true
	}

	return false
}

// SetLowerLimit gets a reference to the given string and assigns it to the LowerLimit field.
func (o *SponsoredProductsBudgetError) SetLowerLimit(v string) {
	o.LowerLimit = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsBudgetError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBudgetError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsBudgetError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsBudgetError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !IsNil(o.UpperLimit) {
		toSerialize["upperLimit"] = o.UpperLimit
	}
	if !IsNil(o.LowerLimit) {
		toSerialize["lowerLimit"] = o.LowerLimit
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsBudgetError struct {
	value *SponsoredProductsBudgetError
	isSet bool
}

func (v NullableSponsoredProductsBudgetError) Get() *SponsoredProductsBudgetError {
	return v.value
}

func (v *NullableSponsoredProductsBudgetError) Set(val *SponsoredProductsBudgetError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBudgetError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBudgetError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBudgetError(val *SponsoredProductsBudgetError) *NullableSponsoredProductsBudgetError {
	return &NullableSponsoredProductsBudgetError{value: val, isSet: true}
}

func (v NullableSponsoredProductsBudgetError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBudgetError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
