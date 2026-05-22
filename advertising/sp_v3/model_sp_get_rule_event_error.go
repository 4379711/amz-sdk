package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPGetRuleEventError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPGetRuleEventError{}

// SPGetRuleEventError The Error Response Object.
type SPGetRuleEventError struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewSPGetRuleEventError instantiates a new SPGetRuleEventError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPGetRuleEventError() *SPGetRuleEventError {
	this := SPGetRuleEventError{}
	return &this
}

// NewSPGetRuleEventErrorWithDefaults instantiates a new SPGetRuleEventError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPGetRuleEventErrorWithDefaults() *SPGetRuleEventError {
	this := SPGetRuleEventError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SPGetRuleEventError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPGetRuleEventError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SPGetRuleEventError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SPGetRuleEventError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SPGetRuleEventError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPGetRuleEventError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SPGetRuleEventError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SPGetRuleEventError) SetDetails(v string) {
	o.Details = &v
}

func (o SPGetRuleEventError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableSPGetRuleEventError struct {
	value *SPGetRuleEventError
	isSet bool
}

func (v NullableSPGetRuleEventError) Get() *SPGetRuleEventError {
	return v.value
}

func (v *NullableSPGetRuleEventError) Set(val *SPGetRuleEventError) {
	v.value = val
	v.isSet = true
}

func (v NullableSPGetRuleEventError) IsSet() bool {
	return v.isSet
}

func (v *NullableSPGetRuleEventError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPGetRuleEventError(val *SPGetRuleEventError) *NullableSPGetRuleEventError {
	return &NullableSPGetRuleEventError{value: val, isSet: true}
}

func (v NullableSPGetRuleEventError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPGetRuleEventError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
