package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioDuplicateValueError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioDuplicateValueError{}

// PortfolioDuplicateValueError struct for PortfolioDuplicateValueError
type PortfolioDuplicateValueError struct {
	Reason      PortfolioDuplicateValueErrorReason `json:"reason"`
	Marketplace *string                            `json:"marketplace,omitempty"`
	Cause       ErrorCause                         `json:"cause"`
	// Human readable error message
	Message string `json:"message"`
}

type _PortfolioDuplicateValueError PortfolioDuplicateValueError

// NewPortfolioDuplicateValueError instantiates a new PortfolioDuplicateValueError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioDuplicateValueError(reason PortfolioDuplicateValueErrorReason, cause ErrorCause, message string) *PortfolioDuplicateValueError {
	this := PortfolioDuplicateValueError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewPortfolioDuplicateValueErrorWithDefaults instantiates a new PortfolioDuplicateValueError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioDuplicateValueErrorWithDefaults() *PortfolioDuplicateValueError {
	this := PortfolioDuplicateValueError{}
	return &this
}

// GetReason returns the Reason field value
func (o *PortfolioDuplicateValueError) GetReason() PortfolioDuplicateValueErrorReason {
	if o == nil {
		var ret PortfolioDuplicateValueErrorReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PortfolioDuplicateValueError) GetReasonOk() (*PortfolioDuplicateValueErrorReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PortfolioDuplicateValueError) SetReason(v PortfolioDuplicateValueErrorReason) {
	o.Reason = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *PortfolioDuplicateValueError) GetMarketplace() string {
	if o == nil || IsNil(o.Marketplace) {
		var ret string
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioDuplicateValueError) GetMarketplaceOk() (*string, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *PortfolioDuplicateValueError) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given string and assigns it to the Marketplace field.
func (o *PortfolioDuplicateValueError) SetMarketplace(v string) {
	o.Marketplace = &v
}

// GetCause returns the Cause field value
func (o *PortfolioDuplicateValueError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *PortfolioDuplicateValueError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *PortfolioDuplicateValueError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetMessage returns the Message field value
func (o *PortfolioDuplicateValueError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PortfolioDuplicateValueError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PortfolioDuplicateValueError) SetMessage(v string) {
	o.Message = v
}

func (o PortfolioDuplicateValueError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["cause"] = o.Cause
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullablePortfolioDuplicateValueError struct {
	value *PortfolioDuplicateValueError
	isSet bool
}

func (v NullablePortfolioDuplicateValueError) Get() *PortfolioDuplicateValueError {
	return v.value
}

func (v *NullablePortfolioDuplicateValueError) Set(val *PortfolioDuplicateValueError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioDuplicateValueError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioDuplicateValueError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioDuplicateValueError(val *PortfolioDuplicateValueError) *NullablePortfolioDuplicateValueError {
	return &NullablePortfolioDuplicateValueError{value: val, isSet: true}
}

func (v NullablePortfolioDuplicateValueError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioDuplicateValueError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
