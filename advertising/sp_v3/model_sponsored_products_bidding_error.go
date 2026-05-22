package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBiddingError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBiddingError{}

// SponsoredProductsBiddingError Errors related to bids
type SponsoredProductsBiddingError struct {
	Reason      SponsoredProductsBiddingErrorReason `json:"reason"`
	Marketplace *SponsoredProductsMarketplace       `json:"marketplace,omitempty"`
	Cause       *SponsoredProductsErrorCause        `json:"cause,omitempty"`
	UpperLimit  *string                             `json:"upperLimit,omitempty"`
	LowerLimit  *string                             `json:"lowerLimit,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsBiddingError SponsoredProductsBiddingError

// NewSponsoredProductsBiddingError instantiates a new SponsoredProductsBiddingError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBiddingError(reason SponsoredProductsBiddingErrorReason, message string) *SponsoredProductsBiddingError {
	this := SponsoredProductsBiddingError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsBiddingErrorWithDefaults instantiates a new SponsoredProductsBiddingError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBiddingErrorWithDefaults() *SponsoredProductsBiddingError {
	this := SponsoredProductsBiddingError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsBiddingError) GetReason() SponsoredProductsBiddingErrorReason {
	if o == nil {
		var ret SponsoredProductsBiddingErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBiddingError) GetReasonOk() (*SponsoredProductsBiddingErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsBiddingError) SetReason(v SponsoredProductsBiddingErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsBiddingError) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBiddingError) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsBiddingError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsBiddingError) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsBiddingError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBiddingError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsBiddingError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsBiddingError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetUpperLimit returns the UpperLimit field value if set, zero value otherwise.
func (o *SponsoredProductsBiddingError) GetUpperLimit() string {
	if o == nil || IsNil(o.UpperLimit) {
		var ret string
		return ret
	}
	return *o.UpperLimit
}

// GetUpperLimitOk returns a tuple with the UpperLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBiddingError) GetUpperLimitOk() (*string, bool) {
	if o == nil || IsNil(o.UpperLimit) {
		return nil, false
	}
	return o.UpperLimit, true
}

// HasUpperLimit returns a boolean if a field has been set.
func (o *SponsoredProductsBiddingError) HasUpperLimit() bool {
	if o != nil && !IsNil(o.UpperLimit) {
		return true
	}

	return false
}

// SetUpperLimit gets a reference to the given string and assigns it to the UpperLimit field.
func (o *SponsoredProductsBiddingError) SetUpperLimit(v string) {
	o.UpperLimit = &v
}

// GetLowerLimit returns the LowerLimit field value if set, zero value otherwise.
func (o *SponsoredProductsBiddingError) GetLowerLimit() string {
	if o == nil || IsNil(o.LowerLimit) {
		var ret string
		return ret
	}
	return *o.LowerLimit
}

// GetLowerLimitOk returns a tuple with the LowerLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBiddingError) GetLowerLimitOk() (*string, bool) {
	if o == nil || IsNil(o.LowerLimit) {
		return nil, false
	}
	return o.LowerLimit, true
}

// HasLowerLimit returns a boolean if a field has been set.
func (o *SponsoredProductsBiddingError) HasLowerLimit() bool {
	if o != nil && !IsNil(o.LowerLimit) {
		return true
	}

	return false
}

// SetLowerLimit gets a reference to the given string and assigns it to the LowerLimit field.
func (o *SponsoredProductsBiddingError) SetLowerLimit(v string) {
	o.LowerLimit = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsBiddingError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBiddingError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsBiddingError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsBiddingError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !IsNil(o.UpperLimit) {
		toSerialize["upperLimit"] = o.UpperLimit
	}
	if !IsNil(o.LowerLimit) {
		toSerialize["lowerLimit"] = o.LowerLimit
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableSponsoredProductsBiddingError struct {
	value *SponsoredProductsBiddingError
	isSet bool
}

func (v NullableSponsoredProductsBiddingError) Get() *SponsoredProductsBiddingError {
	return v.value
}

func (v *NullableSponsoredProductsBiddingError) Set(val *SponsoredProductsBiddingError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBiddingError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBiddingError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBiddingError(val *SponsoredProductsBiddingError) *NullableSponsoredProductsBiddingError {
	return &NullableSponsoredProductsBiddingError{value: val, isSet: true}
}

func (v NullableSponsoredProductsBiddingError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBiddingError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
