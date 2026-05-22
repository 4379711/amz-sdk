package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetableCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetableCategories{}

// TargetableCategories Response object containing all targetable categories for the advertiser's marketplace. ID is the category ID. NA is the name. CH is the list of child categories. TA is if the category is targetable. AsinCountRange is the AsinCounts of the node. Version 4 adds the number of targetable ASINs to each category.
type TargetableCategories struct {
	CategoryTree *string `json:"categoryTree,omitempty"`
}

// NewTargetableCategories instantiates a new TargetableCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetableCategories() *TargetableCategories {
	this := TargetableCategories{}
	return &this
}

// NewTargetableCategoriesWithDefaults instantiates a new TargetableCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetableCategoriesWithDefaults() *TargetableCategories {
	this := TargetableCategories{}
	return &this
}

// GetCategoryTree returns the CategoryTree field value if set, zero value otherwise.
func (o *TargetableCategories) GetCategoryTree() string {
	if o == nil || IsNil(o.CategoryTree) {
		var ret string
		return ret
	}
	return *o.CategoryTree
}

// GetCategoryTreeOk returns a tuple with the CategoryTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetableCategories) GetCategoryTreeOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryTree) {
		return nil, false
	}
	return o.CategoryTree, true
}

// HasCategoryTree returns a boolean if a field has been set.
func (o *TargetableCategories) HasCategoryTree() bool {
	if o != nil && !IsNil(o.CategoryTree) {
		return true
	}

	return false
}

// SetCategoryTree gets a reference to the given string and assigns it to the CategoryTree field.
func (o *TargetableCategories) SetCategoryTree(v string) {
	o.CategoryTree = &v
}

func (o TargetableCategories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryTree) {
		toSerialize["categoryTree"] = o.CategoryTree
	}
	return toSerialize, nil
}

type NullableTargetableCategories struct {
	value *TargetableCategories
	isSet bool
}

func (v NullableTargetableCategories) Get() *TargetableCategories {
	return v.value
}

func (v *NullableTargetableCategories) Set(val *TargetableCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetableCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetableCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetableCategories(val *TargetableCategories) *NullableTargetableCategories {
	return &NullableTargetableCategories{value: val, isSet: true}
}

func (v NullableTargetableCategories) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetableCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
