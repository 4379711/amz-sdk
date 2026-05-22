package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the ValidationMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationMetadata{}

// ValidationMetadata ValidationMetadata Details
type ValidationMetadata struct {
	// errorMessage for the error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// validationStrategy for the error.
	ValidationStrategy *string `json:"validationStrategy,omitempty"`
	// Value.
	Value *string `json:"value,omitempty"`
}

// NewValidationMetadata instantiates a new ValidationMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationMetadata() *ValidationMetadata {
	this := ValidationMetadata{}
	return &this
}

// NewValidationMetadataWithDefaults instantiates a new ValidationMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationMetadataWithDefaults() *ValidationMetadata {
	this := ValidationMetadata{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ValidationMetadata) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationMetadata) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ValidationMetadata) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ValidationMetadata) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetValidationStrategy returns the ValidationStrategy field value if set, zero value otherwise.
func (o *ValidationMetadata) GetValidationStrategy() string {
	if o == nil || IsNil(o.ValidationStrategy) {
		var ret string
		return ret
	}
	return *o.ValidationStrategy
}

// GetValidationStrategyOk returns a tuple with the ValidationStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationMetadata) GetValidationStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationStrategy) {
		return nil, false
	}
	return o.ValidationStrategy, true
}

// HasValidationStrategy returns a boolean if a field has been set.
func (o *ValidationMetadata) HasValidationStrategy() bool {
	if o != nil && !IsNil(o.ValidationStrategy) {
		return true
	}

	return false
}

// SetValidationStrategy gets a reference to the given string and assigns it to the ValidationStrategy field.
func (o *ValidationMetadata) SetValidationStrategy(v string) {
	o.ValidationStrategy = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ValidationMetadata) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationMetadata) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ValidationMetadata) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ValidationMetadata) SetValue(v string) {
	o.Value = &v
}

func (o ValidationMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.ValidationStrategy) {
		toSerialize["validationStrategy"] = o.ValidationStrategy
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableValidationMetadata struct {
	value *ValidationMetadata
	isSet bool
}

func (v NullableValidationMetadata) Get() *ValidationMetadata {
	return v.value
}

func (v *NullableValidationMetadata) Set(val *ValidationMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationMetadata(val *ValidationMetadata) *NullableValidationMetadata {
	return &NullableValidationMetadata{value: val, isSet: true}
}

func (v NullableValidationMetadata) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableValidationMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
