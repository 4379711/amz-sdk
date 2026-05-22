package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDateError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDateError{}

// SponsoredProductsDateError struct for SponsoredProductsDateError
type SponsoredProductsDateError struct {
	Reason SponsoredProductsDateErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause     `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsDateError SponsoredProductsDateError

// NewSponsoredProductsDateError instantiates a new SponsoredProductsDateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDateError(reason SponsoredProductsDateErrorReason, message string) *SponsoredProductsDateError {
	this := SponsoredProductsDateError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsDateErrorWithDefaults instantiates a new SponsoredProductsDateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDateErrorWithDefaults() *SponsoredProductsDateError {
	this := SponsoredProductsDateError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsDateError) GetReason() SponsoredProductsDateErrorReason {
	if o == nil {
		var ret SponsoredProductsDateErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDateError) GetReasonOk() (*SponsoredProductsDateErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsDateError) SetReason(v SponsoredProductsDateErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsDateError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDateError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsDateError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsDateError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsDateError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDateError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsDateError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsDateError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsDateError struct {
	value *SponsoredProductsDateError
	isSet bool
}

func (v NullableSponsoredProductsDateError) Get() *SponsoredProductsDateError {
	return v.value
}

func (v *NullableSponsoredProductsDateError) Set(val *SponsoredProductsDateError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDateError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDateError(val *SponsoredProductsDateError) *NullableSponsoredProductsDateError {
	return &NullableSponsoredProductsDateError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDateError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
