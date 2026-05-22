package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CategoryRecommendationsWithAsinCountsLoP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryRecommendationsWithAsinCountsLoP{}

// CategoryRecommendationsWithAsinCountsLoP Response object for the GetCategoryRecommendationsForAsins API.
type CategoryRecommendationsWithAsinCountsLoP struct {
	// List of category recommendations
	Categories []CategoryItemWithAsinCountsLoP `json:"categories,omitempty"`
}

// NewCategoryRecommendationsWithAsinCountsLoP instantiates a new CategoryRecommendationsWithAsinCountsLoP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryRecommendationsWithAsinCountsLoP() *CategoryRecommendationsWithAsinCountsLoP {
	this := CategoryRecommendationsWithAsinCountsLoP{}
	return &this
}

// NewCategoryRecommendationsWithAsinCountsLoPWithDefaults instantiates a new CategoryRecommendationsWithAsinCountsLoP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryRecommendationsWithAsinCountsLoPWithDefaults() *CategoryRecommendationsWithAsinCountsLoP {
	this := CategoryRecommendationsWithAsinCountsLoP{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *CategoryRecommendationsWithAsinCountsLoP) GetCategories() []CategoryItemWithAsinCountsLoP {
	if o == nil || IsNil(o.Categories) {
		var ret []CategoryItemWithAsinCountsLoP
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryRecommendationsWithAsinCountsLoP) GetCategoriesOk() ([]CategoryItemWithAsinCountsLoP, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *CategoryRecommendationsWithAsinCountsLoP) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []CategoryItemWithAsinCountsLoP and assigns it to the Categories field.
func (o *CategoryRecommendationsWithAsinCountsLoP) SetCategories(v []CategoryItemWithAsinCountsLoP) {
	o.Categories = v
}

func (o CategoryRecommendationsWithAsinCountsLoP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableCategoryRecommendationsWithAsinCountsLoP struct {
	value *CategoryRecommendationsWithAsinCountsLoP
	isSet bool
}

func (v NullableCategoryRecommendationsWithAsinCountsLoP) Get() *CategoryRecommendationsWithAsinCountsLoP {
	return v.value
}

func (v *NullableCategoryRecommendationsWithAsinCountsLoP) Set(val *CategoryRecommendationsWithAsinCountsLoP) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryRecommendationsWithAsinCountsLoP) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryRecommendationsWithAsinCountsLoP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryRecommendationsWithAsinCountsLoP(val *CategoryRecommendationsWithAsinCountsLoP) *NullableCategoryRecommendationsWithAsinCountsLoP {
	return &NullableCategoryRecommendationsWithAsinCountsLoP{value: val, isSet: true}
}

func (v NullableCategoryRecommendationsWithAsinCountsLoP) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCategoryRecommendationsWithAsinCountsLoP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
