package services

import (
	"github.com/bytedance/sonic"
)

// checks if the Warning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Warning{}

// Warning Warning returned when the request is successful, but there are important callouts based on which API clients should take defined actions.
type Warning struct {
	// An warning code that identifies the type of warning that occurred.
	Code string `json:"code"`
	// A message that describes the warning condition in a human-readable form.
	Message string `json:"message"`
	// Additional details that can help the caller understand or address the warning.
	Details *string `json:"details,omitempty"`
}

type _Warning Warning

// NewWarning instantiates a new Warning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarning(code string, message string) *Warning {
	this := Warning{}
	this.Code = code
	this.Message = message
	return &this
}

// NewWarningWithDefaults instantiates a new Warning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarningWithDefaults() *Warning {
	this := Warning{}
	return &this
}

// GetCode returns the Code field value
func (o *Warning) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Warning) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Warning) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *Warning) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Warning) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Warning) SetMessage(v string) {
	o.Message = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Warning) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warning) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Warning) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *Warning) SetDetails(v string) {
	o.Details = &v
}

func (o Warning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableWarning struct {
	value *Warning
	isSet bool
}

func (v NullableWarning) Get() *Warning {
	return v.value
}

func (v *NullableWarning) Set(val *Warning) {
	v.value = val
	v.isSet = true
}

func (v NullableWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarning(val *Warning) *NullableWarning {
	return &NullableWarning{value: val, isSet: true}
}

func (v NullableWarning) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
