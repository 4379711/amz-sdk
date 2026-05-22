package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsLocaleError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsLocaleError{}

// SponsoredProductsLocaleError struct for SponsoredProductsLocaleError
type SponsoredProductsLocaleError struct {
	Reason SponsoredProductsLocaleErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause       `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsLocaleError SponsoredProductsLocaleError

// NewSponsoredProductsLocaleError instantiates a new SponsoredProductsLocaleError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsLocaleError(reason SponsoredProductsLocaleErrorReason, message string) *SponsoredProductsLocaleError {
	this := SponsoredProductsLocaleError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsLocaleErrorWithDefaults instantiates a new SponsoredProductsLocaleError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsLocaleErrorWithDefaults() *SponsoredProductsLocaleError {
	this := SponsoredProductsLocaleError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsLocaleError) GetReason() SponsoredProductsLocaleErrorReason {
	if o == nil {
		var ret SponsoredProductsLocaleErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsLocaleError) GetReasonOk() (*SponsoredProductsLocaleErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsLocaleError) SetReason(v SponsoredProductsLocaleErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsLocaleError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsLocaleError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsLocaleError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsLocaleError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsLocaleError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsLocaleError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsLocaleError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsLocaleError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsLocaleError struct {
	value *SponsoredProductsLocaleError
	isSet bool
}

func (v NullableSponsoredProductsLocaleError) Get() *SponsoredProductsLocaleError {
	return v.value
}

func (v *NullableSponsoredProductsLocaleError) Set(val *SponsoredProductsLocaleError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsLocaleError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsLocaleError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsLocaleError(val *SponsoredProductsLocaleError) *NullableSponsoredProductsLocaleError {
	return &NullableSponsoredProductsLocaleError{value: val, isSet: true}
}

func (v NullableSponsoredProductsLocaleError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsLocaleError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
