package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemRelationship{}

// ItemRelationship the relationship details for a listing item.
type ItemRelationship struct {
	// Identifiers (SKUs) of the related items that are children of this listing item.
	ChildSkus []string `json:"childSkus,omitempty"`
	// Identifiers (SKUs) of the related items that are parents of this listing item.
	ParentSkus     []string            `json:"parentSkus,omitempty"`
	VariationTheme *ItemVariationTheme `json:"variationTheme,omitempty"`
	// The type of relationship.
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

// GetChildSkus returns the ChildSkus field value if set, zero value otherwise.
func (o *ItemRelationship) GetChildSkus() []string {
	if o == nil || IsNil(o.ChildSkus) {
		var ret []string
		return ret
	}
	return o.ChildSkus
}

// GetChildSkusOk returns a tuple with the ChildSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemRelationship) GetChildSkusOk() ([]string, bool) {
	if o == nil || IsNil(o.ChildSkus) {
		return nil, false
	}
	return o.ChildSkus, true
}

// HasChildSkus returns a boolean if a field has been set.
func (o *ItemRelationship) HasChildSkus() bool {
	if o != nil && !IsNil(o.ChildSkus) {
		return true
	}

	return false
}

// SetChildSkus gets a reference to the given []string and assigns it to the ChildSkus field.
func (o *ItemRelationship) SetChildSkus(v []string) {
	o.ChildSkus = v
}

// GetParentSkus returns the ParentSkus field value if set, zero value otherwise.
func (o *ItemRelationship) GetParentSkus() []string {
	if o == nil || IsNil(o.ParentSkus) {
		var ret []string
		return ret
	}
	return o.ParentSkus
}

// GetParentSkusOk returns a tuple with the ParentSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemRelationship) GetParentSkusOk() ([]string, bool) {
	if o == nil || IsNil(o.ParentSkus) {
		return nil, false
	}
	return o.ParentSkus, true
}

// HasParentSkus returns a boolean if a field has been set.
func (o *ItemRelationship) HasParentSkus() bool {
	if o != nil && !IsNil(o.ParentSkus) {
		return true
	}

	return false
}

// SetParentSkus gets a reference to the given []string and assigns it to the ParentSkus field.
func (o *ItemRelationship) SetParentSkus(v []string) {
	o.ParentSkus = v
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
	if !IsNil(o.ChildSkus) {
		toSerialize["childSkus"] = o.ChildSkus
	}
	if !IsNil(o.ParentSkus) {
		toSerialize["parentSkus"] = o.ParentSkus
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
