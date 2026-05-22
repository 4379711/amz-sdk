package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMalformedValueError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMalformedValueError{}

// SponsoredProductsMalformedValueError Errors being used to represent malformed values e.g. containing not allowed characters, not following patters etc
type SponsoredProductsMalformedValueError struct {
	Reason SponsoredProductsMalformedValueErrorReason `json:"reason"`
	// fragment of the value which is wrong
	Fragment    *string                       `json:"fragment,omitempty"`
	Marketplace *SponsoredProductsMarketplace `json:"marketplace,omitempty"`
	Cause       *SponsoredProductsErrorCause  `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsMalformedValueError SponsoredProductsMalformedValueError

// NewSponsoredProductsMalformedValueError instantiates a new SponsoredProductsMalformedValueError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMalformedValueError(reason SponsoredProductsMalformedValueErrorReason, message string) *SponsoredProductsMalformedValueError {
	this := SponsoredProductsMalformedValueError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsMalformedValueErrorWithDefaults instantiates a new SponsoredProductsMalformedValueError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMalformedValueErrorWithDefaults() *SponsoredProductsMalformedValueError {
	this := SponsoredProductsMalformedValueError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsMalformedValueError) GetReason() SponsoredProductsMalformedValueErrorReason {
	if o == nil {
		var ret SponsoredProductsMalformedValueErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMalformedValueError) GetReasonOk() (*SponsoredProductsMalformedValueErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsMalformedValueError) SetReason(v SponsoredProductsMalformedValueErrorReason) {
	o.Reason = v
}

// GetFragment returns the Fragment field value if set, zero value otherwise.
func (o *SponsoredProductsMalformedValueError) GetFragment() string {
	if o == nil || IsNil(o.Fragment) {
		var ret string
		return ret
	}
	return *o.Fragment
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMalformedValueError) GetFragmentOk() (*string, bool) {
	if o == nil || IsNil(o.Fragment) {
		return nil, false
	}
	return o.Fragment, true
}

// HasFragment returns a boolean if a field has been set.
func (o *SponsoredProductsMalformedValueError) HasFragment() bool {
	if o != nil && !IsNil(o.Fragment) {
		return true
	}

	return false
}

// SetFragment gets a reference to the given string and assigns it to the Fragment field.
func (o *SponsoredProductsMalformedValueError) SetFragment(v string) {
	o.Fragment = &v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsMalformedValueError) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMalformedValueError) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsMalformedValueError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsMalformedValueError) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsMalformedValueError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMalformedValueError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsMalformedValueError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsMalformedValueError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsMalformedValueError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMalformedValueError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsMalformedValueError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsMalformedValueError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Fragment) {
		toSerialize["fragment"] = o.Fragment
	}
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsMalformedValueError struct {
	value *SponsoredProductsMalformedValueError
	isSet bool
}

func (v NullableSponsoredProductsMalformedValueError) Get() *SponsoredProductsMalformedValueError {
	return v.value
}

func (v *NullableSponsoredProductsMalformedValueError) Set(val *SponsoredProductsMalformedValueError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMalformedValueError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMalformedValueError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMalformedValueError(val *SponsoredProductsMalformedValueError) *NullableSponsoredProductsMalformedValueError {
	return &NullableSponsoredProductsMalformedValueError{value: val, isSet: true}
}

func (v NullableSponsoredProductsMalformedValueError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMalformedValueError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
