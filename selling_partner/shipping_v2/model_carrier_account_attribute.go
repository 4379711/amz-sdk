package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the CarrierAccountAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CarrierAccountAttribute{}

// CarrierAccountAttribute Attribute Properties required by carrier
type CarrierAccountAttribute struct {
	// Attribute Name .
	AttributeName *string `json:"attributeName,omitempty"`
	// Property Group.
	PropertyGroup *string `json:"propertyGroup,omitempty"`
	// Value .
	Value *string `json:"value,omitempty"`
}

// NewCarrierAccountAttribute instantiates a new CarrierAccountAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrierAccountAttribute() *CarrierAccountAttribute {
	this := CarrierAccountAttribute{}
	return &this
}

// NewCarrierAccountAttributeWithDefaults instantiates a new CarrierAccountAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierAccountAttributeWithDefaults() *CarrierAccountAttribute {
	this := CarrierAccountAttribute{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *CarrierAccountAttribute) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierAccountAttribute) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *CarrierAccountAttribute) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *CarrierAccountAttribute) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetPropertyGroup returns the PropertyGroup field value if set, zero value otherwise.
func (o *CarrierAccountAttribute) GetPropertyGroup() string {
	if o == nil || IsNil(o.PropertyGroup) {
		var ret string
		return ret
	}
	return *o.PropertyGroup
}

// GetPropertyGroupOk returns a tuple with the PropertyGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierAccountAttribute) GetPropertyGroupOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyGroup) {
		return nil, false
	}
	return o.PropertyGroup, true
}

// HasPropertyGroup returns a boolean if a field has been set.
func (o *CarrierAccountAttribute) HasPropertyGroup() bool {
	if o != nil && !IsNil(o.PropertyGroup) {
		return true
	}

	return false
}

// SetPropertyGroup gets a reference to the given string and assigns it to the PropertyGroup field.
func (o *CarrierAccountAttribute) SetPropertyGroup(v string) {
	o.PropertyGroup = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CarrierAccountAttribute) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierAccountAttribute) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CarrierAccountAttribute) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CarrierAccountAttribute) SetValue(v string) {
	o.Value = &v
}

func (o CarrierAccountAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !IsNil(o.PropertyGroup) {
		toSerialize["propertyGroup"] = o.PropertyGroup
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCarrierAccountAttribute struct {
	value *CarrierAccountAttribute
	isSet bool
}

func (v NullableCarrierAccountAttribute) Get() *CarrierAccountAttribute {
	return v.value
}

func (v *NullableCarrierAccountAttribute) Set(val *CarrierAccountAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierAccountAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierAccountAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierAccountAttribute(val *CarrierAccountAttribute) *NullableCarrierAccountAttribute {
	return &NullableCarrierAccountAttribute{value: val, isSet: true}
}

func (v NullableCarrierAccountAttribute) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCarrierAccountAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
