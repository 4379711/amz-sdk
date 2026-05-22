package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioOtherError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioOtherError{}

// PortfolioOtherError Errors not related to any of the other error types
type PortfolioOtherError struct {
	Reason      PortfolioOtherErrorReason `json:"reason"`
	Marketplace *string                   `json:"marketplace,omitempty"`
	Cause       ErrorCause                `json:"cause"`
	// Human readable error message
	Message string `json:"message"`
}

type _PortfolioOtherError PortfolioOtherError

// NewPortfolioOtherError instantiates a new PortfolioOtherError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioOtherError(reason PortfolioOtherErrorReason, cause ErrorCause, message string) *PortfolioOtherError {
	this := PortfolioOtherError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewPortfolioOtherErrorWithDefaults instantiates a new PortfolioOtherError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioOtherErrorWithDefaults() *PortfolioOtherError {
	this := PortfolioOtherError{}
	return &this
}

// GetReason returns the Reason field value
func (o *PortfolioOtherError) GetReason() PortfolioOtherErrorReason {
	if o == nil {
		var ret PortfolioOtherErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PortfolioOtherError) GetReasonOk() (*PortfolioOtherErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PortfolioOtherError) SetReason(v PortfolioOtherErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *PortfolioOtherError) GetMarketplace() string {
	if o == nil || IsNil(o.Marketplace) {
		var ret string
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioOtherError) GetMarketplaceOk() (*string, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *PortfolioOtherError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given string and assigns it to the Marketplace field.
func (o *PortfolioOtherError) SetMarketplace(v string) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value
func (o *PortfolioOtherError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *PortfolioOtherError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *PortfolioOtherError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetMessage returns the Message field value
func (o *PortfolioOtherError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PortfolioOtherError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PortfolioOtherError) SetMessage(v string) {
	o.Message = v
}

func (o PortfolioOtherError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["cause"] = o.Cause
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullablePortfolioOtherError struct {
	value *PortfolioOtherError
	isSet bool
}

func (v NullablePortfolioOtherError) Get() *PortfolioOtherError {
	return v.value
}

func (v *NullablePortfolioOtherError) Set(val *PortfolioOtherError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioOtherError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioOtherError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioOtherError(val *PortfolioOtherError) *NullablePortfolioOtherError {
	return &NullablePortfolioOtherError{value: val, isSet: true}
}

func (v NullablePortfolioOtherError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioOtherError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
