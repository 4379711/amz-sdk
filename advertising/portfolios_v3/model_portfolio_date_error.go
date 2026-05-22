package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioDateError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioDateError{}

// PortfolioDateError struct for PortfolioDateError
type PortfolioDateError struct {
	Reason      PortfolioDateErrorReason `json:"reason"`
	Marketplace *string                  `json:"marketplace,omitempty"`
	Cause       ErrorCause               `json:"cause"`
	// Human readable error message
	Message string `json:"message"`
}

type _PortfolioDateError PortfolioDateError

// NewPortfolioDateError instantiates a new PortfolioDateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioDateError(reason PortfolioDateErrorReason, cause ErrorCause, message string) *PortfolioDateError {
	this := PortfolioDateError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewPortfolioDateErrorWithDefaults instantiates a new PortfolioDateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioDateErrorWithDefaults() *PortfolioDateError {
	this := PortfolioDateError{}
	return &this
}

// GetReason returns the Reason field value
func (o *PortfolioDateError) GetReason() PortfolioDateErrorReason {
	if o == nil {
		var ret PortfolioDateErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PortfolioDateError) GetReasonOk() (*PortfolioDateErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PortfolioDateError) SetReason(v PortfolioDateErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *PortfolioDateError) GetMarketplace() string {
	if o == nil || IsNil(o.Marketplace) {
		var ret string
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioDateError) GetMarketplaceOk() (*string, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *PortfolioDateError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given string and assigns it to the Marketplace field.
func (o *PortfolioDateError) SetMarketplace(v string) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value
func (o *PortfolioDateError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *PortfolioDateError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *PortfolioDateError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetMessage returns the Message field value
func (o *PortfolioDateError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PortfolioDateError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PortfolioDateError) SetMessage(v string) {
	o.Message = v
}

func (o PortfolioDateError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["cause"] = o.Cause
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullablePortfolioDateError struct {
	value *PortfolioDateError
	isSet bool
}

func (v NullablePortfolioDateError) Get() *PortfolioDateError {
	return v.value
}

func (v *NullablePortfolioDateError) Set(val *PortfolioDateError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioDateError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioDateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioDateError(val *PortfolioDateError) *NullablePortfolioDateError {
	return &NullablePortfolioDateError{value: val, isSet: true}
}

func (v NullablePortfolioDateError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioDateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
