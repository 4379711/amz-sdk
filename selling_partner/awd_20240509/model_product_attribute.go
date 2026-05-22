package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAttribute{}

// ProductAttribute Product instance attribute that is not described at the SKU level in the catalog.
type ProductAttribute struct {
	// Product attribute name.
	Name *string `json:"name,omitempty"`
	// Product attribute value.
	Value *string `json:"value,omitempty"`
}

// NewProductAttribute instantiates a new ProductAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAttribute() *ProductAttribute {
	this := ProductAttribute{}
	return &this
}

// NewProductAttributeWithDefaults instantiates a new ProductAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAttributeWithDefaults() *ProductAttribute {
	this := ProductAttribute{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAttribute) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttribute) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAttribute) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAttribute) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProductAttribute) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttribute) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProductAttribute) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ProductAttribute) SetValue(v string) {
	o.Value = &v
}

func (o ProductAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableProductAttribute struct {
	value *ProductAttribute
	isSet bool
}

func (v NullableProductAttribute) Get() *ProductAttribute {
	return v.value
}

func (v *NullableProductAttribute) Set(val *ProductAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAttribute(val *ProductAttribute) *NullableProductAttribute {
	return &NullableProductAttribute{value: val, isSet: true}
}

func (v NullableProductAttribute) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
