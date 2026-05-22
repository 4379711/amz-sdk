package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AdGroupMutationErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroupMutationErrorSelector{}

// AdGroupMutationErrorSelector struct for AdGroupMutationErrorSelector
type AdGroupMutationErrorSelector struct {
	DateError    *DateError    `json:"dateError,omitempty"`
	BiddingError *BiddingError `json:"biddingError,omitempty"`
	RangeError   *RangeError   `json:"rangeError,omitempty"`
	OtherError   *OtherError   `json:"otherError,omitempty"`
}

// NewAdGroupMutationErrorSelector instantiates a new AdGroupMutationErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroupMutationErrorSelector() *AdGroupMutationErrorSelector {
	this := AdGroupMutationErrorSelector{}
	return &this
}

// NewAdGroupMutationErrorSelectorWithDefaults instantiates a new AdGroupMutationErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupMutationErrorSelectorWithDefaults() *AdGroupMutationErrorSelector {
	this := AdGroupMutationErrorSelector{}
	return &this
}

// GetDateError returns the DateError field value if set, zero value otherwise.
func (o *AdGroupMutationErrorSelector) GetDateError() DateError {
	if o == nil || IsNil(o.DateError) {
		var ret DateError
		return ret
	}
	return *o.DateError
}

// GetDateErrorOk returns a tuple with the DateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupMutationErrorSelector) GetDateErrorOk() (*DateError, bool) {
	if o == nil || IsNil(o.DateError) {
		return nil, false
	}
	return o.DateError, true
}

// HasDateError returns a boolean if a field has been set.
func (o *AdGroupMutationErrorSelector) HasDateError() bool {
	if o != nil && !IsNil(o.DateError) {
		return true
	}

	return false
}

// SetDateError gets a reference to the given DateError and assigns it to the DateError field.
func (o *AdGroupMutationErrorSelector) SetDateError(v DateError) {
	o.DateError = &v
}

// GetBiddingError returns the BiddingError field value if set, zero value otherwise.
func (o *AdGroupMutationErrorSelector) GetBiddingError() BiddingError {
	if o == nil || IsNil(o.BiddingError) {
		var ret BiddingError
		return ret
	}
	return *o.BiddingError
}

// GetBiddingErrorOk returns a tuple with the BiddingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupMutationErrorSelector) GetBiddingErrorOk() (*BiddingError, bool) {
	if o == nil || IsNil(o.BiddingError) {
		return nil, false
	}
	return o.BiddingError, true
}

// HasBiddingError returns a boolean if a field has been set.
func (o *AdGroupMutationErrorSelector) HasBiddingError() bool {
	if o != nil && !IsNil(o.BiddingError) {
		return true
	}

	return false
}

// SetBiddingError gets a reference to the given BiddingError and assigns it to the BiddingError field.
func (o *AdGroupMutationErrorSelector) SetBiddingError(v BiddingError) {
	o.BiddingError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *AdGroupMutationErrorSelector) GetRangeError() RangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret RangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupMutationErrorSelector) GetRangeErrorOk() (*RangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *AdGroupMutationErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given RangeError and assigns it to the RangeError field.
func (o *AdGroupMutationErrorSelector) SetRangeError(v RangeError) {
	o.RangeError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *AdGroupMutationErrorSelector) GetOtherError() OtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret OtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupMutationErrorSelector) GetOtherErrorOk() (*OtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *AdGroupMutationErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given OtherError and assigns it to the OtherError field.
func (o *AdGroupMutationErrorSelector) SetOtherError(v OtherError) {
	o.OtherError = &v
}

func (o AdGroupMutationErrorSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateError) {
		toSerialize["dateError"] = o.DateError
	}
	if !IsNil(o.BiddingError) {
		toSerialize["biddingError"] = o.BiddingError
	}
	if !IsNil(o.RangeError) {
		toSerialize["rangeError"] = o.RangeError
	}
	if !IsNil(o.OtherError) {
		toSerialize["otherError"] = o.OtherError
	}
	return toSerialize, nil
}

type NullableAdGroupMutationErrorSelector struct {
	value *AdGroupMutationErrorSelector
	isSet bool
}

func (v NullableAdGroupMutationErrorSelector) Get() *AdGroupMutationErrorSelector {
	return v.value
}

func (v *NullableAdGroupMutationErrorSelector) Set(val *AdGroupMutationErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroupMutationErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroupMutationErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroupMutationErrorSelector(val *AdGroupMutationErrorSelector) *NullableAdGroupMutationErrorSelector {
	return &NullableAdGroupMutationErrorSelector{value: val, isSet: true}
}

func (v NullableAdGroupMutationErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroupMutationErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
