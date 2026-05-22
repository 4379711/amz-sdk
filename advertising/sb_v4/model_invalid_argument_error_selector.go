package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the InvalidArgumentErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidArgumentErrorSelector{}

// InvalidArgumentErrorSelector struct for InvalidArgumentErrorSelector
type InvalidArgumentErrorSelector struct {
	RangeError *RangeError `json:"rangeError,omitempty"`
	OtherError *OtherError `json:"otherError,omitempty"`
}

// NewInvalidArgumentErrorSelector instantiates a new InvalidArgumentErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidArgumentErrorSelector() *InvalidArgumentErrorSelector {
	this := InvalidArgumentErrorSelector{}
	return &this
}

// NewInvalidArgumentErrorSelectorWithDefaults instantiates a new InvalidArgumentErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidArgumentErrorSelectorWithDefaults() *InvalidArgumentErrorSelector {
	this := InvalidArgumentErrorSelector{}
	return &this
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *InvalidArgumentErrorSelector) GetRangeError() RangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret RangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidArgumentErrorSelector) GetRangeErrorOk() (*RangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *InvalidArgumentErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given RangeError and assigns it to the RangeError field.
func (o *InvalidArgumentErrorSelector) SetRangeError(v RangeError) {
	o.RangeError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *InvalidArgumentErrorSelector) GetOtherError() OtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret OtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidArgumentErrorSelector) GetOtherErrorOk() (*OtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *InvalidArgumentErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given OtherError and assigns it to the OtherError field.
func (o *InvalidArgumentErrorSelector) SetOtherError(v OtherError) {
	o.OtherError = &v
}

func (o InvalidArgumentErrorSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RangeError) {
		toSerialize["rangeError"] = o.RangeError
	}
	if !IsNil(o.OtherError) {
		toSerialize["otherError"] = o.OtherError
	}
	return toSerialize, nil
}

type NullableInvalidArgumentErrorSelector struct {
	value *InvalidArgumentErrorSelector
	isSet bool
}

func (v NullableInvalidArgumentErrorSelector) Get() *InvalidArgumentErrorSelector {
	return v.value
}

func (v *NullableInvalidArgumentErrorSelector) Set(val *InvalidArgumentErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidArgumentErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidArgumentErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidArgumentErrorSelector(val *InvalidArgumentErrorSelector) *NullableInvalidArgumentErrorSelector {
	return &NullableInvalidArgumentErrorSelector{value: val, isSet: true}
}

func (v NullableInvalidArgumentErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvalidArgumentErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
