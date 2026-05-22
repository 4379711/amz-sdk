package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ThrottlingException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThrottlingException{}

// ThrottlingException Returns information about a ThrottlingException.
type ThrottlingException struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewThrottlingException instantiates a new ThrottlingException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThrottlingException() *ThrottlingException {
	this := ThrottlingException{}
	return &this
}

// NewThrottlingExceptionWithDefaults instantiates a new ThrottlingException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThrottlingExceptionWithDefaults() *ThrottlingException {
	this := ThrottlingException{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ThrottlingException) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThrottlingException) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ThrottlingException) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ThrottlingException) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ThrottlingException) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThrottlingException) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ThrottlingException) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ThrottlingException) SetDetails(v string) {
	o.Details = &v
}

func (o ThrottlingException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableThrottlingException struct {
	value *ThrottlingException
	isSet bool
}

func (v NullableThrottlingException) Get() *ThrottlingException {
	return v.value
}

func (v *NullableThrottlingException) Set(val *ThrottlingException) {
	v.value = val
	v.isSet = true
}

func (v NullableThrottlingException) IsSet() bool {
	return v.isSet
}

func (v *NullableThrottlingException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThrottlingException(val *ThrottlingException) *NullableThrottlingException {
	return &NullableThrottlingException{value: val, isSet: true}
}

func (v NullableThrottlingException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThrottlingException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
