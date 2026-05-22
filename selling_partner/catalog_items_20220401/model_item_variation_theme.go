package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemVariationTheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemVariationTheme{}

// ItemVariationTheme The variation theme is a list of Amazon catalog item attributes that define the variation family.
type ItemVariationTheme struct {
	// Names of the Amazon catalog item attributes that are associated with the variation theme.
	Attributes []string `json:"attributes,omitempty"`
	// Variation theme that indicates the combination of Amazon catalog item attributes that define the variation family.
	Theme *string `json:"theme,omitempty"`
}

// NewItemVariationTheme instantiates a new ItemVariationTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemVariationTheme() *ItemVariationTheme {
	this := ItemVariationTheme{}
	return &this
}

// NewItemVariationThemeWithDefaults instantiates a new ItemVariationTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemVariationThemeWithDefaults() *ItemVariationTheme {
	this := ItemVariationTheme{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ItemVariationTheme) GetAttributes() []string {
	if o == nil || IsNil(o.Attributes) {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemVariationTheme) GetAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ItemVariationTheme) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *ItemVariationTheme) SetAttributes(v []string) {
	o.Attributes = v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *ItemVariationTheme) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemVariationTheme) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *ItemVariationTheme) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *ItemVariationTheme) SetTheme(v string) {
	o.Theme = &v
}

func (o ItemVariationTheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	return toSerialize, nil
}

type NullableItemVariationTheme struct {
	value *ItemVariationTheme
	isSet bool
}

func (v NullableItemVariationTheme) Get() *ItemVariationTheme {
	return v.value
}

func (v *NullableItemVariationTheme) Set(val *ItemVariationTheme) {
	v.value = val
	v.isSet = true
}

func (v NullableItemVariationTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableItemVariationTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemVariationTheme(val *ItemVariationTheme) *NullableItemVariationTheme {
	return &NullableItemVariationTheme{value: val, isSet: true}
}

func (v NullableItemVariationTheme) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemVariationTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
