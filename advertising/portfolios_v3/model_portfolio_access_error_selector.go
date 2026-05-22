package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioAccessErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioAccessErrorSelector{}

// PortfolioAccessErrorSelector struct for PortfolioAccessErrorSelector
type PortfolioAccessErrorSelector struct {
	EntityNotFoundError *PortfolioEntityNotFoundError `json:"entityNotFoundError,omitempty"`
	MissingValueError   *PortfolioMissingValueError   `json:"missingValueError,omitempty"`
	DateError           *PortfolioDateError           `json:"dateError,omitempty"`
	MalformedValueError *PortfolioMalformedValueError `json:"malformedValueError,omitempty"`
	RangeError          *PortfolioRangeError          `json:"rangeError,omitempty"`
	OtherError          *PortfolioOtherError          `json:"otherError,omitempty"`
}

// NewPortfolioAccessErrorSelector instantiates a new PortfolioAccessErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioAccessErrorSelector() *PortfolioAccessErrorSelector {
	this := PortfolioAccessErrorSelector{}
	return &this
}

// NewPortfolioAccessErrorSelectorWithDefaults instantiates a new PortfolioAccessErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioAccessErrorSelectorWithDefaults() *PortfolioAccessErrorSelector {
	this := PortfolioAccessErrorSelector{}
	return &this
}

// GetEntityNotFoundError returns the EntityNotFoundError field value if set, zero value otherwise.
func (o *PortfolioAccessErrorSelector) GetEntityNotFoundError() PortfolioEntityNotFoundError {
	if o == nil || IsNil(o.EntityNotFoundError) {
		var ret PortfolioEntityNotFoundError
		return ret
	}
	return *o.EntityNotFoundError
}

// GetEntityNotFoundErrorOk returns a tuple with the EntityNotFoundError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioAccessErrorSelector) GetEntityNotFoundErrorOk() (*PortfolioEntityNotFoundError, bool) {
	if o == nil || IsNil(o.EntityNotFoundError) {
		return nil, false
	}
	return o.EntityNotFoundError, true
}

// HasEntityNotFoundError returns a boolean if a field has been set.
func (o *PortfolioAccessErrorSelector) HasEntityNotFoundError() bool {
	if o != nil && !IsNil(o.EntityNotFoundError) {
		return true
	}

	return false
}

// SetEntityNotFoundError gets a reference to the given PortfolioEntityNotFoundError and assigns it to the EntityNotFoundError field.
func (o *PortfolioAccessErrorSelector) SetEntityNotFoundError(v PortfolioEntityNotFoundError) {
	o.EntityNotFoundError = &v
}

// GetMissingValueError returns the MissingValueError field value if set, zero value otherwise.
func (o *PortfolioAccessErrorSelector) GetMissingValueError() PortfolioMissingValueError {
	if o == nil || IsNil(o.MissingValueError) {
		var ret PortfolioMissingValueError
		return ret
	}
	return *o.MissingValueError
}

// GetMissingValueErrorOk returns a tuple with the MissingValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioAccessErrorSelector) GetMissingValueErrorOk() (*PortfolioMissingValueError, bool) {
	if o == nil || IsNil(o.MissingValueError) {
		return nil, false
	}
	return o.MissingValueError, true
}

// HasMissingValueError returns a boolean if a field has been set.
func (o *PortfolioAccessErrorSelector) HasMissingValueError() bool {
	if o != nil && !IsNil(o.MissingValueError) {
		return true
	}

	return false
}

// SetMissingValueError gets a reference to the given PortfolioMissingValueError and assigns it to the MissingValueError field.
func (o *PortfolioAccessErrorSelector) SetMissingValueError(v PortfolioMissingValueError) {
	o.MissingValueError = &v
}

// GetDateError returns the DateError field value if set, zero value otherwise.
func (o *PortfolioAccessErrorSelector) GetDateError() PortfolioDateError {
	if o == nil || IsNil(o.DateError) {
		var ret PortfolioDateError
		return ret
	}
	return *o.DateError
}

// GetDateErrorOk returns a tuple with the DateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioAccessErrorSelector) GetDateErrorOk() (*PortfolioDateError, bool) {
	if o == nil || IsNil(o.DateError) {
		return nil, false
	}
	return o.DateError, true
}

// HasDateError returns a boolean if a field has been set.
func (o *PortfolioAccessErrorSelector) HasDateError() bool {
	if o != nil && !IsNil(o.DateError) {
		return true
	}

	return false
}

// SetDateError gets a reference to the given PortfolioDateError and assigns it to the DateError field.
func (o *PortfolioAccessErrorSelector) SetDateError(v PortfolioDateError) {
	o.DateError = &v
}

// GetMalformedValueError returns the MalformedValueError field value if set, zero value otherwise.
func (o *PortfolioAccessErrorSelector) GetMalformedValueError() PortfolioMalformedValueError {
	if o == nil || IsNil(o.MalformedValueError) {
		var ret PortfolioMalformedValueError
		return ret
	}
	return *o.MalformedValueError
}

// GetMalformedValueErrorOk returns a tuple with the MalformedValueError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioAccessErrorSelector) GetMalformedValueErrorOk() (*PortfolioMalformedValueError, bool) {
	if o == nil || IsNil(o.MalformedValueError) {
		return nil, false
	}
	return o.MalformedValueError, true
}

// HasMalformedValueError returns a boolean if a field has been set.
func (o *PortfolioAccessErrorSelector) HasMalformedValueError() bool {
	if o != nil && !IsNil(o.MalformedValueError) {
		return true
	}

	return false
}

// SetMalformedValueError gets a reference to the given PortfolioMalformedValueError and assigns it to the MalformedValueError field.
func (o *PortfolioAccessErrorSelector) SetMalformedValueError(v PortfolioMalformedValueError) {
	o.MalformedValueError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *PortfolioAccessErrorSelector) GetRangeError() PortfolioRangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret PortfolioRangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioAccessErrorSelector) GetRangeErrorOk() (*PortfolioRangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *PortfolioAccessErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given PortfolioRangeError and assigns it to the RangeError field.
func (o *PortfolioAccessErrorSelector) SetRangeError(v PortfolioRangeError) {
	o.RangeError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *PortfolioAccessErrorSelector) GetOtherError() PortfolioOtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret PortfolioOtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioAccessErrorSelector) GetOtherErrorOk() (*PortfolioOtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *PortfolioAccessErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given PortfolioOtherError and assigns it to the OtherError field.
func (o *PortfolioAccessErrorSelector) SetOtherError(v PortfolioOtherError) {
	o.OtherError = &v
}

func (o PortfolioAccessErrorSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityNotFoundError) {
		toSerialize["entityNotFoundError"] = o.EntityNotFoundError
	}
	if !IsNil(o.MissingValueError) {
		toSerialize["missingValueError"] = o.MissingValueError
	}
	if !IsNil(o.DateError) {
		toSerialize["dateError"] = o.DateError
	}
	if !IsNil(o.MalformedValueError) {
		toSerialize["malformedValueError"] = o.MalformedValueError
	}
	if !IsNil(o.RangeError) {
		toSerialize["rangeError"] = o.RangeError
	}
	if !IsNil(o.OtherError) {
		toSerialize["otherError"] = o.OtherError
	}
	return toSerialize, nil
}

type NullablePortfolioAccessErrorSelector struct {
	value *PortfolioAccessErrorSelector
	isSet bool
}

func (v NullablePortfolioAccessErrorSelector) Get() *PortfolioAccessErrorSelector {
	return v.value
}

func (v *NullablePortfolioAccessErrorSelector) Set(val *PortfolioAccessErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioAccessErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioAccessErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioAccessErrorSelector(val *PortfolioAccessErrorSelector) *NullablePortfolioAccessErrorSelector {
	return &NullablePortfolioAccessErrorSelector{value: val, isSet: true}
}

func (v NullablePortfolioAccessErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioAccessErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
