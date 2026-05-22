package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CategoryRecommendations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryRecommendations{}

// CategoryRecommendations Response object for the GetCategoryRecommendationsForAsins API.
type CategoryRecommendations struct {
	// List of category recommendations
	Categories []CategoryItem `json:"categories,omitempty"`
}

// NewCategoryRecommendations instantiates a new CategoryRecommendations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryRecommendations() *CategoryRecommendations {
	this := CategoryRecommendations{}
	return &this
}

// NewCategoryRecommendationsWithDefaults instantiates a new CategoryRecommendations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryRecommendationsWithDefaults() *CategoryRecommendations {
	this := CategoryRecommendations{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *CategoryRecommendations) GetCategories() []CategoryItem {
	if o == nil || IsNil(o.Categories) {
		var ret []CategoryItem
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryRecommendations) GetCategoriesOk() ([]CategoryItem, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *CategoryRecommendations) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []CategoryItem and assigns it to the Categories field.
func (o *CategoryRecommendations) SetCategories(v []CategoryItem) {
	o.Categories = v
}

func (o CategoryRecommendations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableCategoryRecommendations struct {
	value *CategoryRecommendations
	isSet bool
}

func (v NullableCategoryRecommendations) Get() *CategoryRecommendations {
	return v.value
}

func (v *NullableCategoryRecommendations) Set(val *CategoryRecommendations) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryRecommendations) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryRecommendations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryRecommendations(val *CategoryRecommendations) *NullableCategoryRecommendations {
	return &NullableCategoryRecommendations{value: val, isSet: true}
}

func (v NullableCategoryRecommendations) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCategoryRecommendations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
