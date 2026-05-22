package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioRangeError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioRangeError{}

// PortfolioRangeError Errors related to range constraints violations
type PortfolioRangeError struct {
	Reason      PortfolioValueLimitErrorReason `json:"reason"`
	Marketplace *string                        `json:"marketplace,omitempty"`
	// allowed values
	Allowed []string   `json:"allowed,omitempty"`
	Cause   ErrorCause `json:"cause"`
	// optional upper limit
	UpperLimit *string `json:"upperLimit,omitempty"`
	// optional lower limit
	LowerLimit *string `json:"lowerLimit,omitempty"`
	// Human readable error message
	Message string `json:"message"`
}

type _PortfolioRangeError PortfolioRangeError

// NewPortfolioRangeError instantiates a new PortfolioRangeError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioRangeError(reason PortfolioValueLimitErrorReason, cause ErrorCause, message string) *PortfolioRangeError {
	this := PortfolioRangeError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewPortfolioRangeErrorWithDefaults instantiates a new PortfolioRangeError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioRangeErrorWithDefaults() *PortfolioRangeError {
	this := PortfolioRangeError{}
	return &this
}

// GetReason returns the Reason field value
func (o *PortfolioRangeError) GetReason() PortfolioValueLimitErrorReason {
	if o == nil {
		var ret PortfolioValueLimitErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PortfolioRangeError) GetReasonOk() (*PortfolioValueLimitErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PortfolioRangeError) SetReason(v PortfolioValueLimitErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *PortfolioRangeError) GetMarketplace() string {
	if o == nil || IsNil(o.Marketplace) {
		var ret string
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioRangeError) GetMarketplaceOk() (*string, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *PortfolioRangeError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given string and assigns it to the Marketplace field.
func (o *PortfolioRangeError) SetMarketplace(v string) {
	o.Marketplace = &v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *PortfolioRangeError) GetAllowed() []string {
	if o == nil || IsNil(o.Allowed) {
		var ret []string
		return ret
	}
	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioRangeError) GetAllowedOk() ([]string, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *PortfolioRangeError) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given []string and assigns it to the Allowed field.
func (o *PortfolioRangeError) SetAllowed(v []string) {
	o.Allowed = v
}

// GetCause returns the Cause field value
func (o *PortfolioRangeError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *PortfolioRangeError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *PortfolioRangeError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetUpperLimit returns the UpperLimit field value if set, zero value otherwise.
func (o *PortfolioRangeError) GetUpperLimit() string {
	if o == nil || IsNil(o.UpperLimit) {
		var ret string
		return ret
	}
	return *o.UpperLimit
}

// GetUpperLimitOk returns a tuple with the UpperLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioRangeError) GetUpperLimitOk() (*string, bool) {
	if o == nil || IsNil(o.UpperLimit) {
		return nil, false
	}
	return o.UpperLimit, true
}

// HasUpperLimit returns a boolean if a field has been set.
func (o *PortfolioRangeError) HasUpperLimit() bool {
	if o != nil && !IsNil(o.UpperLimit) {
		return true
	}

	return false
}

// SetUpperLimit gets a reference to the given string and assigns it to the UpperLimit field.
func (o *PortfolioRangeError) SetUpperLimit(v string) {
	o.UpperLimit = &v
}

// GetLowerLimit returns the LowerLimit field value if set, zero value otherwise.
func (o *PortfolioRangeError) GetLowerLimit() string {
	if o == nil || IsNil(o.LowerLimit) {
		var ret string
		return ret
	}
	return *o.LowerLimit
}

// GetLowerLimitOk returns a tuple with the LowerLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioRangeError) GetLowerLimitOk() (*string, bool) {
	if o == nil || IsNil(o.LowerLimit) {
		return nil, false
	}
	return o.LowerLimit, true
}

// HasLowerLimit returns a boolean if a field has been set.
func (o *PortfolioRangeError) HasLowerLimit() bool {
	if o != nil && !IsNil(o.LowerLimit) {
		return true
	}

	return false
}

// SetLowerLimit gets a reference to the given string and assigns it to the LowerLimit field.
func (o *PortfolioRangeError) SetLowerLimit(v string) {
	o.LowerLimit = &v
}

// GetMessage returns the Message field value
func (o *PortfolioRangeError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PortfolioRangeError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PortfolioRangeError) SetMessage(v string) {
	o.Message = v
}

func (o PortfolioRangeError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
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

type NullablePortfolioRangeError struct {
	value *PortfolioRangeError
	isSet bool
}

func (v NullablePortfolioRangeError) Get() *PortfolioRangeError {
	return v.value
}

func (v *NullablePortfolioRangeError) Set(val *PortfolioRangeError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioRangeError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioRangeError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioRangeError(val *PortfolioRangeError) *NullablePortfolioRangeError {
	return &NullablePortfolioRangeError{value: val, isSet: true}
}

func (v NullablePortfolioRangeError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioRangeError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
