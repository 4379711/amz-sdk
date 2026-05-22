package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioMissingValueError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioMissingValueError{}

// PortfolioMissingValueError Error describing missing values in API payloads
type PortfolioMissingValueError struct {
	Reason      PortfolioMissingValueErrorReason `json:"reason"`
	Marketplace *string                          `json:"marketplace,omitempty"`
	Cause       ErrorCause                       `json:"cause"`
	// Human readable error message
	Message string `json:"message"`
}

type _PortfolioMissingValueError PortfolioMissingValueError

// NewPortfolioMissingValueError instantiates a new PortfolioMissingValueError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMissingValueError(reason PortfolioMissingValueErrorReason, cause ErrorCause, message string) *PortfolioMissingValueError {
	this := PortfolioMissingValueError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewPortfolioMissingValueErrorWithDefaults instantiates a new PortfolioMissingValueError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMissingValueErrorWithDefaults() *PortfolioMissingValueError {
	this := PortfolioMissingValueError{}
	return &this
}

// GetReason returns the Reason field value
func (o *PortfolioMissingValueError) GetReason() PortfolioMissingValueErrorReason {
	if o == nil {
		var ret PortfolioMissingValueErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PortfolioMissingValueError) GetReasonOk() (*PortfolioMissingValueErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PortfolioMissingValueError) SetReason(v PortfolioMissingValueErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *PortfolioMissingValueError) GetMarketplace() string {
	if o == nil || IsNil(o.Marketplace) {
		var ret string
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMissingValueError) GetMarketplaceOk() (*string, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *PortfolioMissingValueError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given string and assigns it to the Marketplace field.
func (o *PortfolioMissingValueError) SetMarketplace(v string) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value
func (o *PortfolioMissingValueError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *PortfolioMissingValueError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *PortfolioMissingValueError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetMessage returns the Message field value
func (o *PortfolioMissingValueError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PortfolioMissingValueError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PortfolioMissingValueError) SetMessage(v string) {
	o.Message = v
}

func (o PortfolioMissingValueError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["cause"] = o.Cause
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullablePortfolioMissingValueError struct {
	value *PortfolioMissingValueError
	isSet bool
}

func (v NullablePortfolioMissingValueError) Get() *PortfolioMissingValueError {
	return v.value
}

func (v *NullablePortfolioMissingValueError) Set(val *PortfolioMissingValueError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMissingValueError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMissingValueError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMissingValueError(val *PortfolioMissingValueError) *NullablePortfolioMissingValueError {
	return &NullablePortfolioMissingValueError{value: val, isSet: true}
}

func (v NullablePortfolioMissingValueError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioMissingValueError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
