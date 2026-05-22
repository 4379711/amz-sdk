package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the AplusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AplusResponse{}

// AplusResponse The base response data for all A+ Content operations when a request is successful or partially successful. Individual operations can extend this with additional data.
type AplusResponse struct {
	// A set of messages to the user, such as warnings or comments.
	Warnings []Error `json:"warnings,omitempty"`
}

// NewAplusResponse instantiates a new AplusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAplusResponse() *AplusResponse {
	this := AplusResponse{}
	return &this
}

// NewAplusResponseWithDefaults instantiates a new AplusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAplusResponseWithDefaults() *AplusResponse {
	this := AplusResponse{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AplusResponse) GetWarnings() []Error {
	if o == nil || IsNil(o.Warnings) {
		var ret []Error
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AplusResponse) GetWarningsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AplusResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Error and assigns it to the Warnings field.
func (o *AplusResponse) SetWarnings(v []Error) {
	o.Warnings = v
}

func (o AplusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableAplusResponse struct {
	value *AplusResponse
	isSet bool
}

func (v NullableAplusResponse) Get() *AplusResponse {
	return v.value
}

func (v *NullableAplusResponse) Set(val *AplusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAplusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAplusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAplusResponse(val *AplusResponse) *NullableAplusResponse {
	return &NullableAplusResponse{value: val, isSet: true}
}

func (v NullableAplusResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAplusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
