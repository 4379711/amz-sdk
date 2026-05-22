package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUnsupportedOperationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUnsupportedOperationError{}

// SponsoredProductsUnsupportedOperationError Errors being used to represent an unsupported operation e.g. Seller are not supported to create custom text product ads.
type SponsoredProductsUnsupportedOperationError struct {
	Reason SponsoredProductsUnsupportedOperationErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause                     `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsUnsupportedOperationError SponsoredProductsUnsupportedOperationError

// NewSponsoredProductsUnsupportedOperationError instantiates a new SponsoredProductsUnsupportedOperationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUnsupportedOperationError(reason SponsoredProductsUnsupportedOperationErrorReason, message string) *SponsoredProductsUnsupportedOperationError {
	this := SponsoredProductsUnsupportedOperationError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsUnsupportedOperationErrorWithDefaults instantiates a new SponsoredProductsUnsupportedOperationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUnsupportedOperationErrorWithDefaults() *SponsoredProductsUnsupportedOperationError {
	this := SponsoredProductsUnsupportedOperationError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsUnsupportedOperationError) GetReason() SponsoredProductsUnsupportedOperationErrorReason {
	if o == nil {
		var ret SponsoredProductsUnsupportedOperationErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUnsupportedOperationError) GetReasonOk() (*SponsoredProductsUnsupportedOperationErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsUnsupportedOperationError) SetReason(v SponsoredProductsUnsupportedOperationErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsUnsupportedOperationError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUnsupportedOperationError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsUnsupportedOperationError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsUnsupportedOperationError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsUnsupportedOperationError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUnsupportedOperationError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsUnsupportedOperationError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsUnsupportedOperationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsUnsupportedOperationError struct {
	value *SponsoredProductsUnsupportedOperationError
	isSet bool
}

func (v NullableSponsoredProductsUnsupportedOperationError) Get() *SponsoredProductsUnsupportedOperationError {
	return v.value
}

func (v *NullableSponsoredProductsUnsupportedOperationError) Set(val *SponsoredProductsUnsupportedOperationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUnsupportedOperationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUnsupportedOperationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUnsupportedOperationError(val *SponsoredProductsUnsupportedOperationError) *NullableSponsoredProductsUnsupportedOperationError {
	return &NullableSponsoredProductsUnsupportedOperationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsUnsupportedOperationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUnsupportedOperationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
