package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalProductAttributeTargetingTooManyRequestsException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalProductAttributeTargetingTooManyRequestsException{}

// GlobalProductAttributeTargetingTooManyRequestsException struct for GlobalProductAttributeTargetingTooManyRequestsException
type GlobalProductAttributeTargetingTooManyRequestsException struct {
	Errors []GlobalProductAttributeTargetingErrorModel `json:"errors,omitempty"`
}

// NewGlobalProductAttributeTargetingTooManyRequestsException instantiates a new GlobalProductAttributeTargetingTooManyRequestsException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalProductAttributeTargetingTooManyRequestsException() *GlobalProductAttributeTargetingTooManyRequestsException {
	this := GlobalProductAttributeTargetingTooManyRequestsException{}
	return &this
}

// NewGlobalProductAttributeTargetingTooManyRequestsExceptionWithDefaults instantiates a new GlobalProductAttributeTargetingTooManyRequestsException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalProductAttributeTargetingTooManyRequestsExceptionWithDefaults() *GlobalProductAttributeTargetingTooManyRequestsException {
	this := GlobalProductAttributeTargetingTooManyRequestsException{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GlobalProductAttributeTargetingTooManyRequestsException) GetErrors() []GlobalProductAttributeTargetingErrorModel {
	if o == nil || IsNil(o.Errors) {
		var ret []GlobalProductAttributeTargetingErrorModel
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductAttributeTargetingTooManyRequestsException) GetErrorsOk() ([]GlobalProductAttributeTargetingErrorModel, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GlobalProductAttributeTargetingTooManyRequestsException) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GlobalProductAttributeTargetingErrorModel and assigns it to the Errors field.
func (o *GlobalProductAttributeTargetingTooManyRequestsException) SetErrors(v []GlobalProductAttributeTargetingErrorModel) {
	o.Errors = v
}

func (o GlobalProductAttributeTargetingTooManyRequestsException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGlobalProductAttributeTargetingTooManyRequestsException struct {
	value *GlobalProductAttributeTargetingTooManyRequestsException
	isSet bool
}

func (v NullableGlobalProductAttributeTargetingTooManyRequestsException) Get() *GlobalProductAttributeTargetingTooManyRequestsException {
	return v.value
}

func (v *NullableGlobalProductAttributeTargetingTooManyRequestsException) Set(val *GlobalProductAttributeTargetingTooManyRequestsException) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalProductAttributeTargetingTooManyRequestsException) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalProductAttributeTargetingTooManyRequestsException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalProductAttributeTargetingTooManyRequestsException(val *GlobalProductAttributeTargetingTooManyRequestsException) *NullableGlobalProductAttributeTargetingTooManyRequestsException {
	return &NullableGlobalProductAttributeTargetingTooManyRequestsException{value: val, isSet: true}
}

func (v NullableGlobalProductAttributeTargetingTooManyRequestsException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalProductAttributeTargetingTooManyRequestsException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
