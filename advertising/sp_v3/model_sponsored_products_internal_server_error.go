package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsInternalServerError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsInternalServerError{}

// SponsoredProductsInternalServerError Error that represents non-retryable API service error. Sending the same request will result in another error.
type SponsoredProductsInternalServerError struct {
	Reason SponsoredProductsInternalServerErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause               `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsInternalServerError SponsoredProductsInternalServerError

// NewSponsoredProductsInternalServerError instantiates a new SponsoredProductsInternalServerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsInternalServerError(reason SponsoredProductsInternalServerErrorReason, message string) *SponsoredProductsInternalServerError {
	this := SponsoredProductsInternalServerError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsInternalServerErrorWithDefaults instantiates a new SponsoredProductsInternalServerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsInternalServerErrorWithDefaults() *SponsoredProductsInternalServerError {
	this := SponsoredProductsInternalServerError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsInternalServerError) GetReason() SponsoredProductsInternalServerErrorReason {
	if o == nil {
		var ret SponsoredProductsInternalServerErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsInternalServerError) GetReasonOk() (*SponsoredProductsInternalServerErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsInternalServerError) SetReason(v SponsoredProductsInternalServerErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsInternalServerError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsInternalServerError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsInternalServerError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsInternalServerError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsInternalServerError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsInternalServerError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsInternalServerError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsInternalServerError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsInternalServerError struct {
	value *SponsoredProductsInternalServerError
	isSet bool
}

func (v NullableSponsoredProductsInternalServerError) Get() *SponsoredProductsInternalServerError {
	return v.value
}

func (v *NullableSponsoredProductsInternalServerError) Set(val *SponsoredProductsInternalServerError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsInternalServerError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsInternalServerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsInternalServerError(val *SponsoredProductsInternalServerError) *NullableSponsoredProductsInternalServerError {
	return &NullableSponsoredProductsInternalServerError{value: val, isSet: true}
}

func (v NullableSponsoredProductsInternalServerError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsInternalServerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
