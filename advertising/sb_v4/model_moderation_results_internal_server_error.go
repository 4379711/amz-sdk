package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ModerationResultsInternalServerError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModerationResultsInternalServerError{}

// ModerationResultsInternalServerError struct for ModerationResultsInternalServerError
type ModerationResultsInternalServerError struct {
	// Internal error code.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// NewModerationResultsInternalServerError instantiates a new ModerationResultsInternalServerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModerationResultsInternalServerError() *ModerationResultsInternalServerError {
	this := ModerationResultsInternalServerError{}
	return &this
}

// NewModerationResultsInternalServerErrorWithDefaults instantiates a new ModerationResultsInternalServerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModerationResultsInternalServerErrorWithDefaults() *ModerationResultsInternalServerError {
	this := ModerationResultsInternalServerError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ModerationResultsInternalServerError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResultsInternalServerError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ModerationResultsInternalServerError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ModerationResultsInternalServerError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ModerationResultsInternalServerError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResultsInternalServerError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ModerationResultsInternalServerError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ModerationResultsInternalServerError) SetDetails(v string) {
	o.Details = &v
}

func (o ModerationResultsInternalServerError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableModerationResultsInternalServerError struct {
	value *ModerationResultsInternalServerError
	isSet bool
}

func (v NullableModerationResultsInternalServerError) Get() *ModerationResultsInternalServerError {
	return v.value
}

func (v *NullableModerationResultsInternalServerError) Set(val *ModerationResultsInternalServerError) {
	v.value = val
	v.isSet = true
}

func (v NullableModerationResultsInternalServerError) IsSet() bool {
	return v.isSet
}

func (v *NullableModerationResultsInternalServerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModerationResultsInternalServerError(val *ModerationResultsInternalServerError) *NullableModerationResultsInternalServerError {
	return &NullableModerationResultsInternalServerError{value: val, isSet: true}
}

func (v NullableModerationResultsInternalServerError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableModerationResultsInternalServerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
