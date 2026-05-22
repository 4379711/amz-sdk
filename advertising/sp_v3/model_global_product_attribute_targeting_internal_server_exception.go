package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalProductAttributeTargetingInternalServerException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalProductAttributeTargetingInternalServerException{}

// GlobalProductAttributeTargetingInternalServerException struct for GlobalProductAttributeTargetingInternalServerException
type GlobalProductAttributeTargetingInternalServerException struct {
	Errors []GlobalProductAttributeTargetingErrorModel `json:"errors,omitempty"`
}

// NewGlobalProductAttributeTargetingInternalServerException instantiates a new GlobalProductAttributeTargetingInternalServerException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalProductAttributeTargetingInternalServerException() *GlobalProductAttributeTargetingInternalServerException {
	this := GlobalProductAttributeTargetingInternalServerException{}
	return &this
}

// NewGlobalProductAttributeTargetingInternalServerExceptionWithDefaults instantiates a new GlobalProductAttributeTargetingInternalServerException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalProductAttributeTargetingInternalServerExceptionWithDefaults() *GlobalProductAttributeTargetingInternalServerException {
	this := GlobalProductAttributeTargetingInternalServerException{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GlobalProductAttributeTargetingInternalServerException) GetErrors() []GlobalProductAttributeTargetingErrorModel {
	if o == nil || IsNil(o.Errors) {
		var ret []GlobalProductAttributeTargetingErrorModel
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductAttributeTargetingInternalServerException) GetErrorsOk() ([]GlobalProductAttributeTargetingErrorModel, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GlobalProductAttributeTargetingInternalServerException) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GlobalProductAttributeTargetingErrorModel and assigns it to the Errors field.
func (o *GlobalProductAttributeTargetingInternalServerException) SetErrors(v []GlobalProductAttributeTargetingErrorModel) {
	o.Errors = v
}

func (o GlobalProductAttributeTargetingInternalServerException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGlobalProductAttributeTargetingInternalServerException struct {
	value *GlobalProductAttributeTargetingInternalServerException
	isSet bool
}

func (v NullableGlobalProductAttributeTargetingInternalServerException) Get() *GlobalProductAttributeTargetingInternalServerException {
	return v.value
}

func (v *NullableGlobalProductAttributeTargetingInternalServerException) Set(val *GlobalProductAttributeTargetingInternalServerException) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalProductAttributeTargetingInternalServerException) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalProductAttributeTargetingInternalServerException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalProductAttributeTargetingInternalServerException(val *GlobalProductAttributeTargetingInternalServerException) *NullableGlobalProductAttributeTargetingInternalServerException {
	return &NullableGlobalProductAttributeTargetingInternalServerException{value: val, isSet: true}
}

func (v NullableGlobalProductAttributeTargetingInternalServerException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalProductAttributeTargetingInternalServerException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
