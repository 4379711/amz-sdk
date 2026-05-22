package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsRangeError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsRangeError{}

// SponsoredProductsRangeError Errors related to range constraints violations
type SponsoredProductsRangeError struct {
	Reason      SponsoredProductsValueLimitErrorReason `json:"reason"`
	Marketplace *SponsoredProductsMarketplace          `json:"marketplace,omitempty"`
	// allowed values
	Allowed []string                     `json:"allowed,omitempty"`
	Cause   *SponsoredProductsErrorCause `json:"cause,omitempty"`
	// optional upper limit
	UpperLimit *string `json:"upperLimit,omitempty"`
	// optional lower limit
	LowerLimit *string `json:"lowerLimit,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _SponsoredProductsRangeError SponsoredProductsRangeError

// NewSponsoredProductsRangeError instantiates a new SponsoredProductsRangeError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsRangeError(reason SponsoredProductsValueLimitErrorReason, message string) *SponsoredProductsRangeError {
	this := SponsoredProductsRangeError{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSponsoredProductsRangeErrorWithDefaults instantiates a new SponsoredProductsRangeError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsRangeErrorWithDefaults() *SponsoredProductsRangeError {
	this := SponsoredProductsRangeError{}
	return &this
}

// GetReason returns the Reason field value
func (o *SponsoredProductsRangeError) GetReason() SponsoredProductsValueLimitErrorReason {
	if o == nil {
		var ret SponsoredProductsValueLimitErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRangeError) GetReasonOk() (*SponsoredProductsValueLimitErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SponsoredProductsRangeError) SetReason(v SponsoredProductsValueLimitErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsRangeError) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRangeError) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsRangeError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsRangeError) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *SponsoredProductsRangeError) GetAllowed() []string {
	if o == nil || IsNil(o.Allowed) {
		var ret []string
		return ret
	}
	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRangeError) GetAllowedOk() ([]string, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *SponsoredProductsRangeError) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given []string and assigns it to the Allowed field.
func (o *SponsoredProductsRangeError) SetAllowed(v []string) {
	o.Allowed = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *SponsoredProductsRangeError) GetCause() SponsoredProductsErrorCause {
	if o == nil || IsNil(o.Cause) {
		var ret SponsoredProductsErrorCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRangeError) GetCauseOk() (*SponsoredProductsErrorCause, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *SponsoredProductsRangeError) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given SponsoredProductsErrorCause and assigns it to the Cause field.
func (o *SponsoredProductsRangeError) SetCause(v SponsoredProductsErrorCause) {
	o.Cause = &v
}

// GetUpperLimit returns the UpperLimit field value if set, zero value otherwise.
func (o *SponsoredProductsRangeError) GetUpperLimit() string {
	if o == nil || IsNil(o.UpperLimit) {
		var ret string
		return ret
	}
	return *o.UpperLimit
}

// GetUpperLimitOk returns a tuple with the UpperLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRangeError) GetUpperLimitOk() (*string, bool) {
	if o == nil || IsNil(o.UpperLimit) {
		return nil, false
	}
	return o.UpperLimit, true
}

// HasUpperLimit returns a boolean if a field has been set.
func (o *SponsoredProductsRangeError) HasUpperLimit() bool {
	if o != nil && !IsNil(o.UpperLimit) {
		return true
	}

	return false
}

// SetUpperLimit gets a reference to the given string and assigns it to the UpperLimit field.
func (o *SponsoredProductsRangeError) SetUpperLimit(v string) {
	o.UpperLimit = &v
}

// GetLowerLimit returns the LowerLimit field value if set, zero value otherwise.
func (o *SponsoredProductsRangeError) GetLowerLimit() string {
	if o == nil || IsNil(o.LowerLimit) {
		var ret string
		return ret
	}
	return *o.LowerLimit
}

// GetLowerLimitOk returns a tuple with the LowerLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRangeError) GetLowerLimitOk() (*string, bool) {
	if o == nil || IsNil(o.LowerLimit) {
		return nil, false
	}
	return o.LowerLimit, true
}

// HasLowerLimit returns a boolean if a field has been set.
func (o *SponsoredProductsRangeError) HasLowerLimit() bool {
	if o != nil && !IsNil(o.LowerLimit) {
		return true
	}

	return false
}

// SetLowerLimit gets a reference to the given string and assigns it to the LowerLimit field.
func (o *SponsoredProductsRangeError) SetLowerLimit(v string) {
	o.LowerLimit = &v
}

// GetMessage returns the Message field value
func (o *SponsoredProductsRangeError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRangeError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SponsoredProductsRangeError) SetMessage(v string) {
	o.Message = v
}

func (o SponsoredProductsRangeError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
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

type NullableSponsoredProductsRangeError struct {
	value *SponsoredProductsRangeError
	isSet bool
}

func (v NullableSponsoredProductsRangeError) Get() *SponsoredProductsRangeError {
	return v.value
}

func (v *NullableSponsoredProductsRangeError) Set(val *SponsoredProductsRangeError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsRangeError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsRangeError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsRangeError(val *SponsoredProductsRangeError) *NullableSponsoredProductsRangeError {
	return &NullableSponsoredProductsRangeError{value: val, isSet: true}
}

func (v NullableSponsoredProductsRangeError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsRangeError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
