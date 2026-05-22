package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsProductIdentifierError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsProductIdentifierError{}

// SponsoredProductsProductIdentifierError Errors related to product identifiers
type SponsoredProductsProductIdentifierError struct {
	Reason      SponsoredProductsProductIdentifierErrorReason `json:"reason"`
	Marketplace *SponsoredProductsMarketplace                 `json:"marketplace,omitempty"`
	Cause       *SponsoredProductsErrorCause                  `json:"cause,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsProductIdentifierError SponsoredProductsProductIdentifierError

// NewSponsoredProductsProductIdentifierError instantiates a new SponsoredProductsProductIdentifierError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsProductIdentifierError(reason SponsoredProductsProductIdentifierErrorReason, message string) *SponsoredProductsProductIdentifierError {
	this := SponsoredProductsProductIdentifierError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsProductIdentifierErrorWithDefaults instantiates a new SponsoredProductsProductIdentifierError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsProductIdentifierErrorWithDefaults() *SponsoredProductsProductIdentifierError {
	this := SponsoredProductsProductIdentifierError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsProductIdentifierError) GetReason() SponsoredProductsProductIdentifierErrorReason {
	if o == nil {
		var ret SponsoredProductsProductIdentifierErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductIdentifierError) GetReasonOk() (*SponsoredProductsProductIdentifierErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsProductIdentifierError) SetReason(v SponsoredProductsProductIdentifierErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsProductIdentifierError) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductIdentifierError) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsProductIdentifierError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsProductIdentifierError) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsProductIdentifierError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductIdentifierError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsProductIdentifierError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsProductIdentifierError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsProductIdentifierError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductIdentifierError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsProductIdentifierError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsProductIdentifierError) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsProductIdentifierError struct {
	value *SponsoredProductsProductIdentifierError
	isSet bool
}

func (v NullableSponsoredProductsProductIdentifierError) Get() *SponsoredProductsProductIdentifierError {
	return v.value
}

func (v *NullableSponsoredProductsProductIdentifierError) Set(val *SponsoredProductsProductIdentifierError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsProductIdentifierError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsProductIdentifierError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsProductIdentifierError(val *SponsoredProductsProductIdentifierError) *NullableSponsoredProductsProductIdentifierError {
	return &NullableSponsoredProductsProductIdentifierError{value: val, isSet: true}
}

func (v NullableSponsoredProductsProductIdentifierError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsProductIdentifierError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
