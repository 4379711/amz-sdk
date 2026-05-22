package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioBudgetError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioBudgetError{}

// PortfolioBudgetError struct for PortfolioBudgetError
type PortfolioBudgetError struct {
	Reason      PortfolioBudgetErrorReason `json:"reason"`
	Marketplace *string                    `json:"marketplace,omitempty"`
	Cause       ErrorCause                 `json:"cause"`
	UpperLimit  *string                    `json:"upperLimit,omitempty"`
	LowerLimit  *string                    `json:"lowerLimit,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _PortfolioBudgetError PortfolioBudgetError

// NewPortfolioBudgetError instantiates a new PortfolioBudgetError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioBudgetError(reason PortfolioBudgetErrorReason, cause ErrorCause, message string) *PortfolioBudgetError {
	this := PortfolioBudgetError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewPortfolioBudgetErrorWithDefaults instantiates a new PortfolioBudgetError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioBudgetErrorWithDefaults() *PortfolioBudgetError {
	this := PortfolioBudgetError{}
	return &this
}

// GetReason returns the Reason field value
func (o *PortfolioBudgetError) GetReason() PortfolioBudgetErrorReason {
	if o == nil {
		var ret PortfolioBudgetErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PortfolioBudgetError) GetReasonOk() (*PortfolioBudgetErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PortfolioBudgetError) SetReason(v PortfolioBudgetErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *PortfolioBudgetError) GetMarketplace() string {
	if o == nil || IsNil(o.Marketplace) {
		var ret string
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioBudgetError) GetMarketplaceOk() (*string, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *PortfolioBudgetError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given string and assigns it to the Marketplace field.
func (o *PortfolioBudgetError) SetMarketplace(v string) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value
func (o *PortfolioBudgetError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *PortfolioBudgetError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *PortfolioBudgetError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetUpperLimit returns the UpperLimit field value if set, zero value otherwise.
func (o *PortfolioBudgetError) GetUpperLimit() string {
	if o == nil || IsNil(o.UpperLimit) {
		var ret string
		return ret
	}
	return *o.UpperLimit
}

// GetUpperLimitOk returns a tuple with the UpperLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioBudgetError) GetUpperLimitOk() (*string, bool) {
	if o == nil || IsNil(o.UpperLimit) {
		return nil, false
	}
	return o.UpperLimit, true
}

// HasUpperLimit returns a boolean if a field has been set.
func (o *PortfolioBudgetError) HasUpperLimit() bool {
	if o != nil && !IsNil(o.UpperLimit) {
		return true
	}

	return false
}

// SetUpperLimit gets a reference to the given string and assigns it to the UpperLimit field.
func (o *PortfolioBudgetError) SetUpperLimit(v string) {
	o.UpperLimit = &v
}

// GetLowerLimit returns the LowerLimit field value if set, zero value otherwise.
func (o *PortfolioBudgetError) GetLowerLimit() string {
	if o == nil || IsNil(o.LowerLimit) {
		var ret string
		return ret
	}
	return *o.LowerLimit
}

// GetLowerLimitOk returns a tuple with the LowerLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioBudgetError) GetLowerLimitOk() (*string, bool) {
	if o == nil || IsNil(o.LowerLimit) {
		return nil, false
	}
	return o.LowerLimit, true
}

// HasLowerLimit returns a boolean if a field has been set.
func (o *PortfolioBudgetError) HasLowerLimit() bool {
	if o != nil && !IsNil(o.LowerLimit) {
		return true
	}

	return false
}

// SetLowerLimit gets a reference to the given string and assigns it to the LowerLimit field.
func (o *PortfolioBudgetError) SetLowerLimit(v string) {
	o.LowerLimit = &v
}

// GetMessage returns the Message field value
func (o *PortfolioBudgetError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PortfolioBudgetError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PortfolioBudgetError) SetMessage(v string) {
	o.Message = v
}

func (o PortfolioBudgetError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["cause"] = o.Cause
	if !IsNil(o.UpperLimit) {
		toSerialize["upperLimit"] = o.UpperLimit
	}
	if !IsNil(o.LowerLimit) {
		toSerialize["lowerLimit"] = o.LowerLimit
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullablePortfolioBudgetError struct {
	value *PortfolioBudgetError
	isSet bool
}

func (v NullablePortfolioBudgetError) Get() *PortfolioBudgetError {
	return v.value
}

func (v *NullablePortfolioBudgetError) Set(val *PortfolioBudgetError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioBudgetError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioBudgetError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioBudgetError(val *PortfolioBudgetError) *NullablePortfolioBudgetError {
	return &NullablePortfolioBudgetError{value: val, isSet: true}
}

func (v NullablePortfolioBudgetError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioBudgetError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
