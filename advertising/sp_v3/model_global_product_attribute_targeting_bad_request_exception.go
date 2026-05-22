package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalProductAttributeTargetingBadRequestException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalProductAttributeTargetingBadRequestException{}

// GlobalProductAttributeTargetingBadRequestException struct for GlobalProductAttributeTargetingBadRequestException
type GlobalProductAttributeTargetingBadRequestException struct {
	Errors []GlobalProductAttributeTargetingErrorModel `json:"errors,omitempty"`
}

// NewGlobalProductAttributeTargetingBadRequestException instantiates a new GlobalProductAttributeTargetingBadRequestException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalProductAttributeTargetingBadRequestException() *GlobalProductAttributeTargetingBadRequestException {
	this := GlobalProductAttributeTargetingBadRequestException{}
	return &this
}

// NewGlobalProductAttributeTargetingBadRequestExceptionWithDefaults instantiates a new GlobalProductAttributeTargetingBadRequestException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalProductAttributeTargetingBadRequestExceptionWithDefaults() *GlobalProductAttributeTargetingBadRequestException {
	this := GlobalProductAttributeTargetingBadRequestException{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GlobalProductAttributeTargetingBadRequestException) GetErrors() []GlobalProductAttributeTargetingErrorModel {
	if o == nil || IsNil(o.Errors) {
		var ret []GlobalProductAttributeTargetingErrorModel
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductAttributeTargetingBadRequestException) GetErrorsOk() ([]GlobalProductAttributeTargetingErrorModel, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GlobalProductAttributeTargetingBadRequestException) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GlobalProductAttributeTargetingErrorModel and assigns it to the Errors field.
func (o *GlobalProductAttributeTargetingBadRequestException) SetErrors(v []GlobalProductAttributeTargetingErrorModel) {
	o.Errors = v
}

func (o GlobalProductAttributeTargetingBadRequestException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGlobalProductAttributeTargetingBadRequestException struct {
	value *GlobalProductAttributeTargetingBadRequestException
	isSet bool
}

func (v NullableGlobalProductAttributeTargetingBadRequestException) Get() *GlobalProductAttributeTargetingBadRequestException {
	return v.value
}

func (v *NullableGlobalProductAttributeTargetingBadRequestException) Set(val *GlobalProductAttributeTargetingBadRequestException) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalProductAttributeTargetingBadRequestException) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalProductAttributeTargetingBadRequestException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalProductAttributeTargetingBadRequestException(val *GlobalProductAttributeTargetingBadRequestException) *NullableGlobalProductAttributeTargetingBadRequestException {
	return &NullableGlobalProductAttributeTargetingBadRequestException{value: val, isSet: true}
}

func (v NullableGlobalProductAttributeTargetingBadRequestException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalProductAttributeTargetingBadRequestException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
