package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAdEligibilityError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAdEligibilityError{}

// SponsoredProductsAdEligibilityError Errors related to ad eligibility
type SponsoredProductsAdEligibilityError struct {
	Reason      SponsoredProductsAdEligibilityErrorReason `json:"reason"`
	Marketplace *SponsoredProductsMarketplace             `json:"marketplace,omitempty"`
	Cause       *SponsoredProductsErrorCause              `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsAdEligibilityError SponsoredProductsAdEligibilityError

// NewSponsoredProductsAdEligibilityError instantiates a new SponsoredProductsAdEligibilityError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAdEligibilityError(reason SponsoredProductsAdEligibilityErrorReason, message string) *SponsoredProductsAdEligibilityError {
	this := SponsoredProductsAdEligibilityError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsAdEligibilityErrorWithDefaults instantiates a new SponsoredProductsAdEligibilityError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAdEligibilityErrorWithDefaults() *SponsoredProductsAdEligibilityError {
	this := SponsoredProductsAdEligibilityError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsAdEligibilityError) GetReason() SponsoredProductsAdEligibilityErrorReason {
	if o == nil {
		var ret SponsoredProductsAdEligibilityErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdEligibilityError) GetReasonOk() (*SponsoredProductsAdEligibilityErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsAdEligibilityError) SetReason(v SponsoredProductsAdEligibilityErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsAdEligibilityError) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdEligibilityError) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsAdEligibilityError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsAdEligibilityError) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsAdEligibilityError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdEligibilityError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsAdEligibilityError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsAdEligibilityError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsAdEligibilityError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdEligibilityError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsAdEligibilityError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsAdEligibilityError) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsAdEligibilityError struct {
	value *SponsoredProductsAdEligibilityError
	isSet bool
}

func (v NullableSponsoredProductsAdEligibilityError) Get() *SponsoredProductsAdEligibilityError {
	return v.value
}

func (v *NullableSponsoredProductsAdEligibilityError) Set(val *SponsoredProductsAdEligibilityError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdEligibilityError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdEligibilityError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdEligibilityError(val *SponsoredProductsAdEligibilityError) *NullableSponsoredProductsAdEligibilityError {
	return &NullableSponsoredProductsAdEligibilityError{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdEligibilityError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdEligibilityError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
