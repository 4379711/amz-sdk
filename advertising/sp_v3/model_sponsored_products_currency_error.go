package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCurrencyError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCurrencyError{}

// SponsoredProductsCurrencyError Errors related to currency
type SponsoredProductsCurrencyError struct {
	Reason SponsoredProductsCurrencyErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause         `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsCurrencyError SponsoredProductsCurrencyError

// NewSponsoredProductsCurrencyError instantiates a new SponsoredProductsCurrencyError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCurrencyError(reason SponsoredProductsCurrencyErrorReason, message string) *SponsoredProductsCurrencyError {
	this := SponsoredProductsCurrencyError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsCurrencyErrorWithDefaults instantiates a new SponsoredProductsCurrencyError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCurrencyErrorWithDefaults() *SponsoredProductsCurrencyError {
	this := SponsoredProductsCurrencyError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsCurrencyError) GetReason() SponsoredProductsCurrencyErrorReason {
	if o == nil {
		var ret SponsoredProductsCurrencyErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCurrencyError) GetReasonOk() (*SponsoredProductsCurrencyErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsCurrencyError) SetReason(v SponsoredProductsCurrencyErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsCurrencyError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCurrencyError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsCurrencyError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsCurrencyError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsCurrencyError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCurrencyError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsCurrencyError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsCurrencyError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsCurrencyError struct {
	value *SponsoredProductsCurrencyError
	isSet bool
}

func (v NullableSponsoredProductsCurrencyError) Get() *SponsoredProductsCurrencyError {
	return v.value
}

func (v *NullableSponsoredProductsCurrencyError) Set(val *SponsoredProductsCurrencyError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCurrencyError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCurrencyError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCurrencyError(val *SponsoredProductsCurrencyError) *NullableSponsoredProductsCurrencyError {
	return &NullableSponsoredProductsCurrencyError{value: val, isSet: true}
}

func (v NullableSponsoredProductsCurrencyError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCurrencyError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
