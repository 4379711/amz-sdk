package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetingClauseSetupError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetingClauseSetupError{}

// SponsoredProductsTargetingClauseSetupError Errors related to targeting clause setup
type SponsoredProductsTargetingClauseSetupError struct {
	Reason      SponsoredProductsTargetingClauseSetupErrorReason `json:"reason"`
	Marketplace *SponsoredProductsMarketplace                    `json:"marketplace,omitempty"`
	Cause       *SponsoredProductsErrorCause                     `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsTargetingClauseSetupError SponsoredProductsTargetingClauseSetupError

// NewSponsoredProductsTargetingClauseSetupError instantiates a new SponsoredProductsTargetingClauseSetupError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetingClauseSetupError(reason SponsoredProductsTargetingClauseSetupErrorReason, message string) *SponsoredProductsTargetingClauseSetupError {
	this := SponsoredProductsTargetingClauseSetupError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsTargetingClauseSetupErrorWithDefaults instantiates a new SponsoredProductsTargetingClauseSetupError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetingClauseSetupErrorWithDefaults() *SponsoredProductsTargetingClauseSetupError {
	this := SponsoredProductsTargetingClauseSetupError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsTargetingClauseSetupError) GetReason() SponsoredProductsTargetingClauseSetupErrorReason {
	if o == nil {
		var ret SponsoredProductsTargetingClauseSetupErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClauseSetupError) GetReasonOk() (*SponsoredProductsTargetingClauseSetupErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsTargetingClauseSetupError) SetReason(v SponsoredProductsTargetingClauseSetupErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingClauseSetupError) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClauseSetupError) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingClauseSetupError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsTargetingClauseSetupError) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingClauseSetupError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClauseSetupError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingClauseSetupError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsTargetingClauseSetupError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsTargetingClauseSetupError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClauseSetupError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsTargetingClauseSetupError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsTargetingClauseSetupError) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsTargetingClauseSetupError struct {
	value *SponsoredProductsTargetingClauseSetupError
	isSet bool
}

func (v NullableSponsoredProductsTargetingClauseSetupError) Get() *SponsoredProductsTargetingClauseSetupError {
	return v.value
}

func (v *NullableSponsoredProductsTargetingClauseSetupError) Set(val *SponsoredProductsTargetingClauseSetupError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetingClauseSetupError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetingClauseSetupError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetingClauseSetupError(val *SponsoredProductsTargetingClauseSetupError) *NullableSponsoredProductsTargetingClauseSetupError {
	return &NullableSponsoredProductsTargetingClauseSetupError{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetingClauseSetupError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetingClauseSetupError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
