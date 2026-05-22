package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAsinOwnershipError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAsinOwnershipError{}

// SponsoredProductsAsinOwnershipError Errors related to author asin ownership
type SponsoredProductsAsinOwnershipError struct {
	Reason SponsoredProductsAsinOwnershipErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause              `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsAsinOwnershipError SponsoredProductsAsinOwnershipError

// NewSponsoredProductsAsinOwnershipError instantiates a new SponsoredProductsAsinOwnershipError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAsinOwnershipError(reason SponsoredProductsAsinOwnershipErrorReason, message string) *SponsoredProductsAsinOwnershipError {
	this := SponsoredProductsAsinOwnershipError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsAsinOwnershipErrorWithDefaults instantiates a new SponsoredProductsAsinOwnershipError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAsinOwnershipErrorWithDefaults() *SponsoredProductsAsinOwnershipError {
	this := SponsoredProductsAsinOwnershipError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsAsinOwnershipError) GetReason() SponsoredProductsAsinOwnershipErrorReason {
	if o == nil {
		var ret SponsoredProductsAsinOwnershipErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAsinOwnershipError) GetReasonOk() (*SponsoredProductsAsinOwnershipErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsAsinOwnershipError) SetReason(v SponsoredProductsAsinOwnershipErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsAsinOwnershipError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAsinOwnershipError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsAsinOwnershipError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsAsinOwnershipError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsAsinOwnershipError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAsinOwnershipError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsAsinOwnershipError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsAsinOwnershipError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsAsinOwnershipError struct {
	value *SponsoredProductsAsinOwnershipError
	isSet bool
}

func (v NullableSponsoredProductsAsinOwnershipError) Get() *SponsoredProductsAsinOwnershipError {
	return v.value
}

func (v *NullableSponsoredProductsAsinOwnershipError) Set(val *SponsoredProductsAsinOwnershipError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAsinOwnershipError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAsinOwnershipError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAsinOwnershipError(val *SponsoredProductsAsinOwnershipError) *NullableSponsoredProductsAsinOwnershipError {
	return &NullableSponsoredProductsAsinOwnershipError{value: val, isSet: true}
}

func (v NullableSponsoredProductsAsinOwnershipError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAsinOwnershipError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
