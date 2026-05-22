package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsParentEntityError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsParentEntityError{}

// SponsoredProductsParentEntityError Errors related to parent entity
type SponsoredProductsParentEntityError struct {
	Reason SponsoredProductsParentEntityErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause             `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsParentEntityError SponsoredProductsParentEntityError

// NewSponsoredProductsParentEntityError instantiates a new SponsoredProductsParentEntityError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsParentEntityError(reason SponsoredProductsParentEntityErrorReason, message string) *SponsoredProductsParentEntityError {
	this := SponsoredProductsParentEntityError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsParentEntityErrorWithDefaults instantiates a new SponsoredProductsParentEntityError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsParentEntityErrorWithDefaults() *SponsoredProductsParentEntityError {
	this := SponsoredProductsParentEntityError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsParentEntityError) GetReason() SponsoredProductsParentEntityErrorReason {
	if o == nil {
		var ret SponsoredProductsParentEntityErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsParentEntityError) GetReasonOk() (*SponsoredProductsParentEntityErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsParentEntityError) SetReason(v SponsoredProductsParentEntityErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsParentEntityError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsParentEntityError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsParentEntityError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsParentEntityError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsParentEntityError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsParentEntityError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsParentEntityError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsParentEntityError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsParentEntityError struct {
	value *SponsoredProductsParentEntityError
	isSet bool
}

func (v NullableSponsoredProductsParentEntityError) Get() *SponsoredProductsParentEntityError {
	return v.value
}

func (v *NullableSponsoredProductsParentEntityError) Set(val *SponsoredProductsParentEntityError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsParentEntityError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsParentEntityError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsParentEntityError(val *SponsoredProductsParentEntityError) *NullableSponsoredProductsParentEntityError {
	return &NullableSponsoredProductsParentEntityError{value: val, isSet: true}
}

func (v NullableSponsoredProductsParentEntityError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsParentEntityError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
