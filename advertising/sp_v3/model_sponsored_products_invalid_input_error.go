package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsInvalidInputError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsInvalidInputError{}

// SponsoredProductsInvalidInputError Errors related to ad eligibility
type SponsoredProductsInvalidInputError struct {
	Reason SponsoredProductsInvalidInputErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause             `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsInvalidInputError SponsoredProductsInvalidInputError

// NewSponsoredProductsInvalidInputError instantiates a new SponsoredProductsInvalidInputError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsInvalidInputError(reason SponsoredProductsInvalidInputErrorReason, message string) *SponsoredProductsInvalidInputError {
	this := SponsoredProductsInvalidInputError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsInvalidInputErrorWithDefaults instantiates a new SponsoredProductsInvalidInputError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsInvalidInputErrorWithDefaults() *SponsoredProductsInvalidInputError {
	this := SponsoredProductsInvalidInputError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsInvalidInputError) GetReason() SponsoredProductsInvalidInputErrorReason {
	if o == nil {
		var ret SponsoredProductsInvalidInputErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsInvalidInputError) GetReasonOk() (*SponsoredProductsInvalidInputErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsInvalidInputError) SetReason(v SponsoredProductsInvalidInputErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsInvalidInputError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsInvalidInputError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsInvalidInputError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsInvalidInputError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsInvalidInputError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsInvalidInputError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsInvalidInputError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsInvalidInputError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsInvalidInputError struct {
	value *SponsoredProductsInvalidInputError
	isSet bool
}

func (v NullableSponsoredProductsInvalidInputError) Get() *SponsoredProductsInvalidInputError {
	return v.value
}

func (v *NullableSponsoredProductsInvalidInputError) Set(val *SponsoredProductsInvalidInputError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsInvalidInputError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsInvalidInputError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsInvalidInputError(val *SponsoredProductsInvalidInputError) *NullableSponsoredProductsInvalidInputError {
	return &NullableSponsoredProductsInvalidInputError{value: val, isSet: true}
}

func (v NullableSponsoredProductsInvalidInputError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsInvalidInputError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
