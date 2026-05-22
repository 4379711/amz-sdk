package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BiddingError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BiddingError{}

// BiddingError Errors related to bids.
type BiddingError struct {
	// Exact error reason.
	Reason     string     `json:"reason"`
	Cause      ErrorCause `json:"cause"`
	UpperLimit *string    `json:"upperLimit,omitempty"`
	LowerLimit *string    `json:"lowerLimit,omitempty"`
	// Human readable error message.
	Message string `json:"message"`
}

type _BiddingError BiddingError

// NewBiddingError instantiates a new BiddingError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiddingError(reason string, cause ErrorCause, message string) *BiddingError {
	this := BiddingError{}
	this.Reason = reason
	this.Cause = cause
	this.Message = message
	return &this
}

// NewBiddingErrorWithDefaults instantiates a new BiddingError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiddingErrorWithDefaults() *BiddingError {
	this := BiddingError{}
	return &this
}

// GetReason returns the Reason field value
func (o *BiddingError) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *BiddingError) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *BiddingError) SetReason(v string) {
	o.Reason = v
}

// GetCause returns the Cause field value
func (o *BiddingError) GetCause() ErrorCause {
	if o == nil {
		var ret ErrorCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *BiddingError) GetCauseOk() (*ErrorCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *BiddingError) SetCause(v ErrorCause) {
	o.Cause = v
}

// GetUpperLimit returns the UpperLimit field value if set, zero value otherwise.
func (o *BiddingError) GetUpperLimit() string {
	if o == nil || IsNil(o.UpperLimit) {
		var ret string
		return ret
	}
	return *o.UpperLimit
}

// GetUpperLimitOk returns a tuple with the UpperLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiddingError) GetUpperLimitOk() (*string, bool) {
	if o == nil || IsNil(o.UpperLimit) {
		return nil, false
	}
	return o.UpperLimit, true
}

// HasUpperLimit returns a boolean if a field has been set.
func (o *BiddingError) HasUpperLimit() bool {
	if o != nil && !IsNil(o.UpperLimit) {
		return true
	}

	return false
}

// SetUpperLimit gets a reference to the given string and assigns it to the UpperLimit field.
func (o *BiddingError) SetUpperLimit(v string) {
	o.UpperLimit = &v
}

// GetLowerLimit returns the LowerLimit field value if set, zero value otherwise.
func (o *BiddingError) GetLowerLimit() string {
	if o == nil || IsNil(o.LowerLimit) {
		var ret string
		return ret
	}
	return *o.LowerLimit
}

// GetLowerLimitOk returns a tuple with the LowerLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiddingError) GetLowerLimitOk() (*string, bool) {
	if o == nil || IsNil(o.LowerLimit) {
		return nil, false
	}
	return o.LowerLimit, true
}

// HasLowerLimit returns a boolean if a field has been set.
func (o *BiddingError) HasLowerLimit() bool {
	if o != nil && !IsNil(o.LowerLimit) {
		return true
	}

	return false
}

// SetLowerLimit gets a reference to the given string and assigns it to the LowerLimit field.
func (o *BiddingError) SetLowerLimit(v string) {
	o.LowerLimit = &v
}

// GetMessage returns the Message field value
func (o *BiddingError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BiddingError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BiddingError) SetMessage(v string) {
	o.Message = v
}

func (o BiddingError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
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

type NullableBiddingError struct {
	value *BiddingError
	isSet bool
}

func (v NullableBiddingError) Get() *BiddingError {
	return v.value
}

func (v *NullableBiddingError) Set(val *BiddingError) {
	v.value = val
	v.isSet = true
}

func (v NullableBiddingError) IsSet() bool {
	return v.isSet
}

func (v *NullableBiddingError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiddingError(val *BiddingError) *NullableBiddingError {
	return &NullableBiddingError{value: val, isSet: true}
}

func (v NullableBiddingError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBiddingError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
