package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsApplicableMarketplacesError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsApplicableMarketplacesError{}

// SponsoredProductsApplicableMarketplacesError Errors related to ad eligibility
type SponsoredProductsApplicableMarketplacesError struct {
	Reason SponsoredProductsApplicableMarketplacesErrorReason `json:"reason"`
	Cause  *SponsoredProductsErrorCause                       `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsApplicableMarketplacesError SponsoredProductsApplicableMarketplacesError

// NewSponsoredProductsApplicableMarketplacesError instantiates a new SponsoredProductsApplicableMarketplacesError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsApplicableMarketplacesError(reason SponsoredProductsApplicableMarketplacesErrorReason, message string) *SponsoredProductsApplicableMarketplacesError {
	this := SponsoredProductsApplicableMarketplacesError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsApplicableMarketplacesErrorWithDefaults instantiates a new SponsoredProductsApplicableMarketplacesError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsApplicableMarketplacesErrorWithDefaults() *SponsoredProductsApplicableMarketplacesError {
	this := SponsoredProductsApplicableMarketplacesError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsApplicableMarketplacesError) GetReason() SponsoredProductsApplicableMarketplacesErrorReason {
	if o == nil {
		var ret SponsoredProductsApplicableMarketplacesErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsApplicableMarketplacesError) GetReasonOk() (*SponsoredProductsApplicableMarketplacesErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsApplicableMarketplacesError) SetReason(v SponsoredProductsApplicableMarketplacesErrorReason) {
	o.Reason = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsApplicableMarketplacesError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsApplicableMarketplacesError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsApplicableMarketplacesError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsApplicableMarketplacesError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsApplicableMarketplacesError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsApplicableMarketplacesError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsApplicableMarketplacesError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsApplicableMarketplacesError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsApplicableMarketplacesError struct {
	value *SponsoredProductsApplicableMarketplacesError
	isSet bool
}

func (v NullableSponsoredProductsApplicableMarketplacesError) Get() *SponsoredProductsApplicableMarketplacesError {
	return v.value
}

func (v *NullableSponsoredProductsApplicableMarketplacesError) Set(val *SponsoredProductsApplicableMarketplacesError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsApplicableMarketplacesError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsApplicableMarketplacesError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsApplicableMarketplacesError(val *SponsoredProductsApplicableMarketplacesError) *NullableSponsoredProductsApplicableMarketplacesError {
	return &NullableSponsoredProductsApplicableMarketplacesError{value: val, isSet: true}
}

func (v NullableSponsoredProductsApplicableMarketplacesError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsApplicableMarketplacesError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
