package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBillingError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBillingError{}

// SponsoredProductsBillingError Errors related to bids
type SponsoredProductsBillingError struct {
	Reason SponsoredProductsBillingErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause        `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsBillingError SponsoredProductsBillingError

// NewSponsoredProductsBillingError instantiates a new SponsoredProductsBillingError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBillingError(reason SponsoredProductsBillingErrorReason, message string) *SponsoredProductsBillingError {
	this := SponsoredProductsBillingError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsBillingErrorWithDefaults instantiates a new SponsoredProductsBillingError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBillingErrorWithDefaults() *SponsoredProductsBillingError {
	this := SponsoredProductsBillingError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsBillingError) GetReason() SponsoredProductsBillingErrorReason {
	if o == nil {
		var ret SponsoredProductsBillingErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBillingError) GetReasonOk() (*SponsoredProductsBillingErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsBillingError) SetReason(v SponsoredProductsBillingErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsBillingError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBillingError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsBillingError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsBillingError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsBillingError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBillingError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsBillingError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsBillingError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsBillingError struct {
	value *SponsoredProductsBillingError
	isSet bool
}

func (v NullableSponsoredProductsBillingError) Get() *SponsoredProductsBillingError {
	return v.value
}

func (v *NullableSponsoredProductsBillingError) Set(val *SponsoredProductsBillingError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBillingError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBillingError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBillingError(val *SponsoredProductsBillingError) *NullableSponsoredProductsBillingError {
	return &NullableSponsoredProductsBillingError{value: val, isSet: true}
}

func (v NullableSponsoredProductsBillingError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBillingError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
