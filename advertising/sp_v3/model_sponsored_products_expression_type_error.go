package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsExpressionTypeError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsExpressionTypeError{}

// SponsoredProductsExpressionTypeError struct for SponsoredProductsExpressionTypeError
type SponsoredProductsExpressionTypeError struct {
	Reason SponsoredProductsExpressionTypeErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause               `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsExpressionTypeError SponsoredProductsExpressionTypeError

// NewSponsoredProductsExpressionTypeError instantiates a new SponsoredProductsExpressionTypeError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsExpressionTypeError(reason SponsoredProductsExpressionTypeErrorReason, message string) *SponsoredProductsExpressionTypeError {
	this := SponsoredProductsExpressionTypeError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsExpressionTypeErrorWithDefaults instantiates a new SponsoredProductsExpressionTypeError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsExpressionTypeErrorWithDefaults() *SponsoredProductsExpressionTypeError {
	this := SponsoredProductsExpressionTypeError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsExpressionTypeError) GetReason() SponsoredProductsExpressionTypeErrorReason {
	if o == nil {
		var ret SponsoredProductsExpressionTypeErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsExpressionTypeError) GetReasonOk() (*SponsoredProductsExpressionTypeErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsExpressionTypeError) SetReason(v SponsoredProductsExpressionTypeErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsExpressionTypeError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsExpressionTypeError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsExpressionTypeError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsExpressionTypeError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsExpressionTypeError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsExpressionTypeError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsExpressionTypeError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsExpressionTypeError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsExpressionTypeError struct {
	value *SponsoredProductsExpressionTypeError
	isSet bool
}

func (v NullableSponsoredProductsExpressionTypeError) Get() *SponsoredProductsExpressionTypeError {
	return v.value
}

func (v *NullableSponsoredProductsExpressionTypeError) Set(val *SponsoredProductsExpressionTypeError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsExpressionTypeError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsExpressionTypeError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsExpressionTypeError(val *SponsoredProductsExpressionTypeError) *NullableSponsoredProductsExpressionTypeError {
	return &NullableSponsoredProductsExpressionTypeError{value: val, isSet: true}
}

func (v NullableSponsoredProductsExpressionTypeError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsExpressionTypeError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
