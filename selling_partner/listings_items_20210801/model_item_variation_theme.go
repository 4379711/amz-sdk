package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemVariationTheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemVariationTheme{}

// ItemVariationTheme A variation theme that indicates the combination of listing item attributes that define the variation family.
type ItemVariationTheme struct {
	// The names of the listing item attributes that are associated with the variation theme.
	Attributes []string `json:"attributes"`
	// The variation theme that indicates the combination of listing item attributes that define the variation family.
	Theme string `json:"theme"`
}

type _ItemVariationTheme ItemVariationTheme

// NewItemVariationTheme instantiates a new ItemVariationTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemVariationTheme(attributes []string, theme string) *ItemVariationTheme {
	this := ItemVariationTheme{}
	this.Attributes = attributes
	this.Theme = theme
	return &this
}

// NewItemVariationThemeWithDefaults instantiates a new ItemVariationTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemVariationThemeWithDefaults() *ItemVariationTheme {
	this := ItemVariationTheme{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *ItemVariationTheme) GetAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ItemVariationTheme) GetAttributesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *ItemVariationTheme) SetAttributes(v []string) {
	o.Attributes = v
}

// GetTheme returns the Theme field value
func (o *ItemVariationTheme) GetTheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *ItemVariationTheme) GetThemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Theme, true
}

// SetTheme sets field value
func (o *ItemVariationTheme) SetTheme(v string) {
	o.Theme = v
}

func (o ItemVariationTheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributes"] = o.Attributes
	toSerialize["theme"] = o.Theme
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
