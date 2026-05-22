package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioBillingError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioBillingError{}

// PortfolioBillingError Errors related to bids
type PortfolioBillingError struct {
	Reason      PortfolioBillingErrorReason `json:"reason"`
	Marketplace *string                     `json:"marketplace,omitempty"`
	Cause       ErrorCause                  `json:"cause"`
	// Human readable error message
	Message string `json:"message"`
}

type _PortfolioBillingError PortfolioBillingError

// NewPortfolioBillingError instantiates a new PortfolioBillingError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioBillingError(reason PortfolioBillingErrorReason, cause ErrorCause, message string) *PortfolioBillingError {
	this := PortfolioBillingError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewPortfolioBillingErrorWithDefaults instantiates a new PortfolioBillingError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioBillingErrorWithDefaults() *PortfolioBillingError {
	this := PortfolioBillingError{}
	return &this
}

// GetReason returns the Reason field value
func (o *PortfolioBillingError) GetReason() PortfolioBillingErrorReason {
	if o == nil {
		var ret PortfolioBillingErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PortfolioBillingError) GetReasonOk() (*PortfolioBillingErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PortfolioBillingError) SetReason(v PortfolioBillingErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *PortfolioBillingError) GetMarketplace() string {
	if o == nil || IsNil(o.Marketplace) {
		var ret string
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioBillingError) GetMarketplaceOk() (*string, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *PortfolioBillingError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given string and assigns it to the Marketplace field.
func (o *PortfolioBillingError) SetMarketplace(v string) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value
func (o *PortfolioBillingError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *PortfolioBillingError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *PortfolioBillingError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetMessage returns the Message field value
func (o *PortfolioBillingError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PortfolioBillingError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PortfolioBillingError) SetMessage(v string) {
	o.Message = v
}

func (o PortfolioBillingError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["cause"] = o.Cause
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullablePortfolioBillingError struct {
	value *PortfolioBillingError
	isSet bool
}

func (v NullablePortfolioBillingError) Get() *PortfolioBillingError {
	return v.value
}

func (v *NullablePortfolioBillingError) Set(val *PortfolioBillingError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioBillingError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioBillingError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioBillingError(val *PortfolioBillingError) *NullablePortfolioBillingError {
	return &NullablePortfolioBillingError{value: val, isSet: true}
}

func (v NullablePortfolioBillingError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioBillingError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
