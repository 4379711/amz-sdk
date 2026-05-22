package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemRelationship{}

// ItemRelationship Relationship details for an Amazon catalog item.
type ItemRelationship struct {
	// ASINs of the related items that are children of this item.
	ChildAsins []string `json:"childAsins,omitempty"`
	// ASINs of the related items that are parents of this item.
	ParentAsins    []string            `json:"parentAsins,omitempty"`
	VariationTheme *ItemVariationTheme `json:"variationTheme,omitempty"`
	// Type of relationship.
	Type string `json:"type"`
}

type _ItemRelationship ItemRelationship

// NewItemRelationship instantiates a new ItemRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemRelationship(type_ string) *ItemRelationship {
	this := ItemRelationship{}
	this.Type = type_
	return &this
}

// NewItemRelationshipWithDefaults instantiates a new ItemRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemRelationshipWithDefaults() *ItemRelationship {
	this := ItemRelationship{}
	return &this
}

// GetChildAsins returns the ChildAsins field value if set, zero value otherwise.
func (o *ItemRelationship) GetChildAsins() []string {
	if o == nil || IsNil(o.ChildAsins) {
		var ret []string
		return ret
	}
	return o.ChildAsins
}

// GetChildAsinsOk returns a tuple with the ChildAsins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemRelationship) GetChildAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.ChildAsins) {
		return nil, false
	}
	return o.ChildAsins, true
}

// HasChildAsins returns a boolean if a field has been set.
func (o *ItemRelationship) HasChildAsins() bool {
	if o != nil && !IsNil(o.ChildAsins) {
		return true
	}

	return false
}

// SetChildAsins gets a reference to the given []string and assigns it to the ChildAsins field.
func (o *ItemRelationship) SetChildAsins(v []string) {
	o.ChildAsins = v
}

// GetParentAsins returns the ParentAsins field value if set, zero value otherwise.
func (o *ItemRelationship) GetParentAsins() []string {
	if o == nil || IsNil(o.ParentAsins) {
		var ret []string
		return ret
	}
	return o.ParentAsins
}

// GetParentAsinsOk returns a tuple with the ParentAsins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemRelationship) GetParentAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.ParentAsins) {
		return nil, false
	}
	return o.ParentAsins, true
}

// HasParentAsins returns a boolean if a field has been set.
func (o *ItemRelationship) HasParentAsins() bool {
	if o != nil && !IsNil(o.ParentAsins) {
		return true
	}

	return false
}

// SetParentAsins gets a reference to the given []string and assigns it to the ParentAsins field.
func (o *ItemRelationship) SetParentAsins(v []string) {
	o.ParentAsins = v
}

// GetVariationTheme returns the VariationTheme field value if set, zero value otherwise.
func (o *ItemRelationship) GetVariationTheme() ItemVariationTheme {
	if o == nil || IsNil(o.VariationTheme) {
		var ret ItemVariationTheme
		return ret
	}
	return *o.VariationTheme
}

// GetVariationThemeOk returns a tuple with the VariationTheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemRelationship) GetVariationThemeOk() (*ItemVariationTheme, bool) {
	if o == nil || IsNil(o.VariationTheme) {
		return nil, false
	}
	return o.VariationTheme, true
}

// HasVariationTheme returns a boolean if a field has been set.
func (o *ItemRelationship) HasVariationTheme() bool {
	if o != nil && !IsNil(o.VariationTheme) {
		return true
	}

	return false
}

// SetVariationTheme gets a reference to the given ItemVariationTheme and assigns it to the VariationTheme field.
func (o *ItemRelationship) SetVariationTheme(v ItemVariationTheme) {
	o.VariationTheme = &v
}

// GetType returns the Type field value
func (o *ItemRelationship) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ItemRelationship) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ItemRelationship) SetType(v string) {
	o.Type = v
}

func (o ItemRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChildAsins) {
		toSerialize["childAsins"] = o.ChildAsins
	}
	if !IsNil(o.ParentAsins) {
		toSerialize["parentAsins"] = o.ParentAsins
	}
	if !IsNil(o.VariationTheme) {
		toSerialize["variationTheme"] = o.VariationTheme
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableItemRelationship struct {
	value *ItemRelationship
	isSet bool
}

func (v NullableItemRelationship) Get() *ItemRelationship {
	return v.value
}

func (v *NullableItemRelationship) Set(val *ItemRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableItemRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableItemRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemRelationship(val *ItemRelationship) *NullableItemRelationship {
	return &NullableItemRelationship{value: val, isSet: true}
}

func (v NullableItemRelationship) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
