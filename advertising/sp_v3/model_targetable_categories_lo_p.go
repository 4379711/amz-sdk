package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetableCategoriesLoP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetableCategoriesLoP{}

// TargetableCategoriesLoP Response object containing all targetable categories for the advertiser's marketplace in a language of preference (LoP) provide by the locale query parameter. ID is the category ID. NA is the name. TN is the translated name in the language of preference. CH is the list of child categories. TA is if the category is targetable. AsinCountRange is the AsinCounts of the node. Version 4 adds the number of targetable ASINs to each category.
type TargetableCategoriesLoP struct {
	CategoryTree *string `json:"categoryTree,omitempty"`
}

// NewTargetableCategoriesLoP instantiates a new TargetableCategoriesLoP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetableCategoriesLoP() *TargetableCategoriesLoP {
	this := TargetableCategoriesLoP{}
	return &this
}

// NewTargetableCategoriesLoPWithDefaults instantiates a new TargetableCategoriesLoP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetableCategoriesLoPWithDefaults() *TargetableCategoriesLoP {
	this := TargetableCategoriesLoP{}
	return &this
}

// GetCategoryTree returns the CategoryTree field value if set, zero value otherwise.
func (o *TargetableCategoriesLoP) GetCategoryTree() string {
	if o == nil || IsNil(o.CategoryTree) {
		var ret string
		return ret
	}
	return *o.CategoryTree
}

// GetCategoryTreeOk returns a tuple with the CategoryTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetableCategoriesLoP) GetCategoryTreeOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryTree) {
		return nil, false
	}
	return o.CategoryTree, true
}

// HasCategoryTree returns a boolean if a field has been set.
func (o *TargetableCategoriesLoP) HasCategoryTree() bool {
	if o != nil && !IsNil(o.CategoryTree) {
		return true
	}

	return false
}

// SetCategoryTree gets a reference to the given string and assigns it to the CategoryTree field.
func (o *TargetableCategoriesLoP) SetCategoryTree(v string) {
	o.CategoryTree = &v
}

func (o TargetableCategoriesLoP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryTree) {
		toSerialize["categoryTree"] = o.CategoryTree
	}
	return toSerialize, nil
}

type NullableTargetableCategoriesLoP struct {
	value *TargetableCategoriesLoP
	isSet bool
}

func (v NullableTargetableCategoriesLoP) Get() *TargetableCategoriesLoP {
	return v.value
}

func (v *NullableTargetableCategoriesLoP) Set(val *TargetableCategoriesLoP) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetableCategoriesLoP) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetableCategoriesLoP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetableCategoriesLoP(val *TargetableCategoriesLoP) *NullableTargetableCategoriesLoP {
	return &NullableTargetableCategoriesLoP{value: val, isSet: true}
}

func (v NullableTargetableCategoriesLoP) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetableCategoriesLoP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
