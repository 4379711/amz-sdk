package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ModerationResultsAccessDeniedError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModerationResultsAccessDeniedError{}

// ModerationResultsAccessDeniedError struct for ModerationResultsAccessDeniedError
type ModerationResultsAccessDeniedError struct {
	// Access denied error code.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// NewModerationResultsAccessDeniedError instantiates a new ModerationResultsAccessDeniedError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModerationResultsAccessDeniedError() *ModerationResultsAccessDeniedError {
	this := ModerationResultsAccessDeniedError{}
	return &this
}

// NewModerationResultsAccessDeniedErrorWithDefaults instantiates a new ModerationResultsAccessDeniedError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModerationResultsAccessDeniedErrorWithDefaults() *ModerationResultsAccessDeniedError {
	this := ModerationResultsAccessDeniedError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ModerationResultsAccessDeniedError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResultsAccessDeniedError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ModerationResultsAccessDeniedError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ModerationResultsAccessDeniedError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ModerationResultsAccessDeniedError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResultsAccessDeniedError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ModerationResultsAccessDeniedError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ModerationResultsAccessDeniedError) SetDetails(v string) {
	o.Details = &v
}

func (o ModerationResultsAccessDeniedError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableModerationResultsAccessDeniedError struct {
	value *ModerationResultsAccessDeniedError
	isSet bool
}

func (v NullableModerationResultsAccessDeniedError) Get() *ModerationResultsAccessDeniedError {
	return v.value
}

func (v *NullableModerationResultsAccessDeniedError) Set(val *ModerationResultsAccessDeniedError) {
	v.value = val
	v.isSet = true
}

func (v NullableModerationResultsAccessDeniedError) IsSet() bool {
	return v.isSet
}

func (v *NullableModerationResultsAccessDeniedError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModerationResultsAccessDeniedError(val *ModerationResultsAccessDeniedError) *NullableModerationResultsAccessDeniedError {
	return &NullableModerationResultsAccessDeniedError{value: val, isSet: true}
}

func (v NullableModerationResultsAccessDeniedError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableModerationResultsAccessDeniedError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
