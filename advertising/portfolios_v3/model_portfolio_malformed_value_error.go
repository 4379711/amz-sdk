package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioMalformedValueError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioMalformedValueError{}

// PortfolioMalformedValueError Errors being used to represent malformed values e.g. containing not allowed characters, not following patterns etc
type PortfolioMalformedValueError struct {
	Reason PortfolioMalformedValueErrorReason `json:"reason"`
	// fragment of the value which is wrong
	Fragment    *string    `json:"fragment,omitempty"`
	Marketplace *string    `json:"marketplace,omitempty"`
	Cause       ErrorCause `json:"cause"`
	// Human readable error message
	Message string `json:"message"`
}

type _PortfolioMalformedValueError PortfolioMalformedValueError

// NewPortfolioMalformedValueError instantiates a new PortfolioMalformedValueError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMalformedValueError(reason PortfolioMalformedValueErrorReason, cause ErrorCause, message string) *PortfolioMalformedValueError {
	this := PortfolioMalformedValueError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewPortfolioMalformedValueErrorWithDefaults instantiates a new PortfolioMalformedValueError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMalformedValueErrorWithDefaults() *PortfolioMalformedValueError {
	this := PortfolioMalformedValueError{}
	return &this
}

// GetReason returns the Reason field value
func (o *PortfolioMalformedValueError) GetReason() PortfolioMalformedValueErrorReason {
	if o == nil {
		var ret PortfolioMalformedValueErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PortfolioMalformedValueError) GetReasonOk() (*PortfolioMalformedValueErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PortfolioMalformedValueError) SetReason(v PortfolioMalformedValueErrorReason) {
	o.Reason = v
}

// GetFragment returns the Fragment field value if set, zero value otherwise.
func (o *PortfolioMalformedValueError) GetFragment() string {
	if o == nil || IsNil(o.Fragment) {
		var ret string
		return ret
	}
	return *o.Fragment
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMalformedValueError) GetFragmentOk() (*string, bool) {
	if o == nil || IsNil(o.Fragment) {
		return nil, false
	}
	return o.Fragment, true
}

// HasFragment returns a boolean if a field has been set.
func (o *PortfolioMalformedValueError) HasFragment() bool {
	if o != nil && !IsNil(o.Fragment) {
		return true
	}

	return false
}

// SetFragment gets a reference to the given string and assigns it to the Fragment field.
func (o *PortfolioMalformedValueError) SetFragment(v string) {
	o.Fragment = &v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *PortfolioMalformedValueError) GetMarketplace() string {
	if o == nil || IsNil(o.Marketplace) {
		var ret string
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMalformedValueError) GetMarketplaceOk() (*string, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *PortfolioMalformedValueError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given string and assigns it to the Marketplace field.
func (o *PortfolioMalformedValueError) SetMarketplace(v string) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value
func (o *PortfolioMalformedValueError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *PortfolioMalformedValueError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *PortfolioMalformedValueError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetMessage returns the Message field value
func (o *PortfolioMalformedValueError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PortfolioMalformedValueError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PortfolioMalformedValueError) SetMessage(v string) {
	o.Message = v
}

func (o PortfolioMalformedValueError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Fragment) {
		toSerialize["fragment"] = o.Fragment
	}
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["cause"] = o.Cause
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullablePortfolioMalformedValueError struct {
	value *PortfolioMalformedValueError
	isSet bool
}

func (v NullablePortfolioMalformedValueError) Get() *PortfolioMalformedValueError {
	return v.value
}

func (v *NullablePortfolioMalformedValueError) Set(val *PortfolioMalformedValueError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMalformedValueError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMalformedValueError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMalformedValueError(val *PortfolioMalformedValueError) *NullablePortfolioMalformedValueError {
	return &NullablePortfolioMalformedValueError{value: val, isSet: true}
}

func (v NullablePortfolioMalformedValueError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioMalformedValueError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
