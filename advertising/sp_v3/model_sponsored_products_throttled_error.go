package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsThrottledError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsThrottledError{}

// SponsoredProductsThrottledError Error that represents failure due to API caller exceeding allowed service limits.
type SponsoredProductsThrottledError struct {
	Reason SponsoredProductsThrottledErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause          `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsThrottledError SponsoredProductsThrottledError

// NewSponsoredProductsThrottledError instantiates a new SponsoredProductsThrottledError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsThrottledError(reason SponsoredProductsThrottledErrorReason, message string) *SponsoredProductsThrottledError {
	this := SponsoredProductsThrottledError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsThrottledErrorWithDefaults instantiates a new SponsoredProductsThrottledError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsThrottledErrorWithDefaults() *SponsoredProductsThrottledError {
	this := SponsoredProductsThrottledError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsThrottledError) GetReason() SponsoredProductsThrottledErrorReason {
	if o == nil {
		var ret SponsoredProductsThrottledErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsThrottledError) GetReasonOk() (*SponsoredProductsThrottledErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsThrottledError) SetReason(v SponsoredProductsThrottledErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsThrottledError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsThrottledError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsThrottledError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsThrottledError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsThrottledError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsThrottledError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsThrottledError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsThrottledError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsThrottledError struct {
	value *SponsoredProductsThrottledError
	isSet bool
}

func (v NullableSponsoredProductsThrottledError) Get() *SponsoredProductsThrottledError {
	return v.value
}

func (v *NullableSponsoredProductsThrottledError) Set(val *SponsoredProductsThrottledError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsThrottledError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsThrottledError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsThrottledError(val *SponsoredProductsThrottledError) *NullableSponsoredProductsThrottledError {
	return &NullableSponsoredProductsThrottledError{value: val, isSet: true}
}

func (v NullableSponsoredProductsThrottledError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsThrottledError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
