package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalProductAttributeTargetingUnauthorizedException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalProductAttributeTargetingUnauthorizedException{}

// GlobalProductAttributeTargetingUnauthorizedException struct for GlobalProductAttributeTargetingUnauthorizedException
type GlobalProductAttributeTargetingUnauthorizedException struct {
	Errors []GlobalProductAttributeTargetingErrorModel `json:"errors,omitempty"`
}

// NewGlobalProductAttributeTargetingUnauthorizedException instantiates a new GlobalProductAttributeTargetingUnauthorizedException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalProductAttributeTargetingUnauthorizedException() *GlobalProductAttributeTargetingUnauthorizedException {
	this := GlobalProductAttributeTargetingUnauthorizedException{}
	return &this
}

// NewGlobalProductAttributeTargetingUnauthorizedExceptionWithDefaults instantiates a new GlobalProductAttributeTargetingUnauthorizedException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalProductAttributeTargetingUnauthorizedExceptionWithDefaults() *GlobalProductAttributeTargetingUnauthorizedException {
	this := GlobalProductAttributeTargetingUnauthorizedException{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GlobalProductAttributeTargetingUnauthorizedException) GetErrors() []GlobalProductAttributeTargetingErrorModel {
	if o == nil || IsNil(o.Errors) {
		var ret []GlobalProductAttributeTargetingErrorModel
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductAttributeTargetingUnauthorizedException) GetErrorsOk() ([]GlobalProductAttributeTargetingErrorModel, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GlobalProductAttributeTargetingUnauthorizedException) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GlobalProductAttributeTargetingErrorModel and assigns it to the Errors field.
func (o *GlobalProductAttributeTargetingUnauthorizedException) SetErrors(v []GlobalProductAttributeTargetingErrorModel) {
	o.Errors = v
}

func (o GlobalProductAttributeTargetingUnauthorizedException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGlobalProductAttributeTargetingUnauthorizedException struct {
	value *GlobalProductAttributeTargetingUnauthorizedException
	isSet bool
}

func (v NullableGlobalProductAttributeTargetingUnauthorizedException) Get() *GlobalProductAttributeTargetingUnauthorizedException {
	return v.value
}

func (v *NullableGlobalProductAttributeTargetingUnauthorizedException) Set(val *GlobalProductAttributeTargetingUnauthorizedException) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalProductAttributeTargetingUnauthorizedException) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalProductAttributeTargetingUnauthorizedException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalProductAttributeTargetingUnauthorizedException(val *GlobalProductAttributeTargetingUnauthorizedException) *NullableGlobalProductAttributeTargetingUnauthorizedException {
	return &NullableGlobalProductAttributeTargetingUnauthorizedException{value: val, isSet: true}
}

func (v NullableGlobalProductAttributeTargetingUnauthorizedException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalProductAttributeTargetingUnauthorizedException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
