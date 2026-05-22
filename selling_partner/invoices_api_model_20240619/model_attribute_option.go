package invoices_api_model_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the AttributeOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeOption{}

// AttributeOption The definition of the attribute option.
type AttributeOption struct {
	// The description of the attribute value.
	Description *string `json:"description,omitempty"`
	// The possible values for the attribute option.
	Value *string `json:"value,omitempty"`
}

// NewAttributeOption instantiates a new AttributeOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeOption() *AttributeOption {
	this := AttributeOption{}
	return &this
}

// NewAttributeOptionWithDefaults instantiates a new AttributeOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeOptionWithDefaults() *AttributeOption {
	this := AttributeOption{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AttributeOption) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeOption) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AttributeOption) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AttributeOption) SetDescription(v string) {
	o.Description = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AttributeOption) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeOption) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AttributeOption) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AttributeOption) SetValue(v string) {
	o.Value = &v
}

func (o AttributeOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAttributeOption struct {
	value *AttributeOption
	isSet bool
}

func (v NullableAttributeOption) Get() *AttributeOption {
	return v.value
}

func (v *NullableAttributeOption) Set(val *AttributeOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeOption(val *AttributeOption) *NullableAttributeOption {
	return &NullableAttributeOption{value: val, isSet: true}
}

func (v NullableAttributeOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAttributeOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
