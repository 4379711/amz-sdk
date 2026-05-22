package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMissingValueError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMissingValueError{}

// SponsoredProductsMissingValueError Error describing missing values in API payloads
type SponsoredProductsMissingValueError struct {
	Reason      SponsoredProductsMissingValueErrorReason `json:"reason"`
	Marketplace *SponsoredProductsMarketplace            `json:"marketplace,omitempty"`
	Cause       *SponsoredProductsErrorCause             `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsMissingValueError SponsoredProductsMissingValueError

// NewSponsoredProductsMissingValueError instantiates a new SponsoredProductsMissingValueError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMissingValueError(reason SponsoredProductsMissingValueErrorReason, message string) *SponsoredProductsMissingValueError {
	this := SponsoredProductsMissingValueError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsMissingValueErrorWithDefaults instantiates a new SponsoredProductsMissingValueError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMissingValueErrorWithDefaults() *SponsoredProductsMissingValueError {
	this := SponsoredProductsMissingValueError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsMissingValueError) GetReason() SponsoredProductsMissingValueErrorReason {
	if o == nil {
		var ret SponsoredProductsMissingValueErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMissingValueError) GetReasonOk() (*SponsoredProductsMissingValueErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsMissingValueError) SetReason(v SponsoredProductsMissingValueErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsMissingValueError) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMissingValueError) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsMissingValueError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsMissingValueError) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsMissingValueError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMissingValueError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsMissingValueError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsMissingValueError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsMissingValueError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMissingValueError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsMissingValueError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsMissingValueError) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsMissingValueError struct {
	value *SponsoredProductsMissingValueError
	isSet bool
}

func (v NullableSponsoredProductsMissingValueError) Get() *SponsoredProductsMissingValueError {
	return v.value
}

func (v *NullableSponsoredProductsMissingValueError) Set(val *SponsoredProductsMissingValueError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMissingValueError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMissingValueError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMissingValueError(val *SponsoredProductsMissingValueError) *NullableSponsoredProductsMissingValueError {
	return &NullableSponsoredProductsMissingValueError{value: val, isSet: true}
}

func (v NullableSponsoredProductsMissingValueError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMissingValueError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
