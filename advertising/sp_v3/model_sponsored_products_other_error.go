package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsOtherError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsOtherError{}

// SponsoredProductsOtherError Errors not related to any of the other error types
type SponsoredProductsOtherError struct {
	Reason      SponsoredProductsOtherErrorReason `json:"reason"`
	Marketplace *SponsoredProductsMarketplace     `json:"marketplace,omitempty"`
	Cause       *SponsoredProductsErrorCause      `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsOtherError SponsoredProductsOtherError

// NewSponsoredProductsOtherError instantiates a new SponsoredProductsOtherError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsOtherError(reason SponsoredProductsOtherErrorReason, message string) *SponsoredProductsOtherError {
	this := SponsoredProductsOtherError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsOtherErrorWithDefaults instantiates a new SponsoredProductsOtherError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsOtherErrorWithDefaults() *SponsoredProductsOtherError {
	this := SponsoredProductsOtherError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsOtherError) GetReason() SponsoredProductsOtherErrorReason {
	if o == nil {
		var ret SponsoredProductsOtherErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsOtherError) GetReasonOk() (*SponsoredProductsOtherErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsOtherError) SetReason(v SponsoredProductsOtherErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsOtherError) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsOtherError) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsOtherError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsOtherError) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsOtherError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsOtherError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsOtherError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsOtherError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsOtherError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsOtherError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsOtherError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsOtherError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsOtherError struct {
	value *SponsoredProductsOtherError
	isSet bool
}

func (v NullableSponsoredProductsOtherError) Get() *SponsoredProductsOtherError {
	return v.value
}

func (v *NullableSponsoredProductsOtherError) Set(val *SponsoredProductsOtherError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsOtherError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsOtherError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsOtherError(val *SponsoredProductsOtherError) *NullableSponsoredProductsOtherError {
	return &NullableSponsoredProductsOtherError{value: val, isSet: true}
}

func (v NullableSponsoredProductsOtherError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsOtherError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
