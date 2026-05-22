package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CategoryRecommendationsWithAsinCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryRecommendationsWithAsinCounts{}

// CategoryRecommendationsWithAsinCounts Response object for the GetCategoryRecommendationsForAsins API.
type CategoryRecommendationsWithAsinCounts struct {
	// List of category recommendations
	Categories []CategoryItemWithAsinCounts `json:"categories,omitempty"`
}

// NewCategoryRecommendationsWithAsinCounts instantiates a new CategoryRecommendationsWithAsinCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryRecommendationsWithAsinCounts() *CategoryRecommendationsWithAsinCounts {
	this := CategoryRecommendationsWithAsinCounts{}
	return &this
}

// NewCategoryRecommendationsWithAsinCountsWithDefaults instantiates a new CategoryRecommendationsWithAsinCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryRecommendationsWithAsinCountsWithDefaults() *CategoryRecommendationsWithAsinCounts {
	this := CategoryRecommendationsWithAsinCounts{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *CategoryRecommendationsWithAsinCounts) GetCategories() []CategoryItemWithAsinCounts {
	if o == nil || IsNil(o.Categories) {
		var ret []CategoryItemWithAsinCounts
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryRecommendationsWithAsinCounts) GetCategoriesOk() ([]CategoryItemWithAsinCounts, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *CategoryRecommendationsWithAsinCounts) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []CategoryItemWithAsinCounts and assigns it to the Categories field.
func (o *CategoryRecommendationsWithAsinCounts) SetCategories(v []CategoryItemWithAsinCounts) {
	o.Categories = v
}

func (o CategoryRecommendationsWithAsinCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableCategoryRecommendationsWithAsinCounts struct {
	value *CategoryRecommendationsWithAsinCounts
	isSet bool
}

func (v NullableCategoryRecommendationsWithAsinCounts) Get() *CategoryRecommendationsWithAsinCounts {
	return v.value
}

func (v *NullableCategoryRecommendationsWithAsinCounts) Set(val *CategoryRecommendationsWithAsinCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryRecommendationsWithAsinCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryRecommendationsWithAsinCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryRecommendationsWithAsinCounts(val *CategoryRecommendationsWithAsinCounts) *NullableCategoryRecommendationsWithAsinCounts {
	return &NullableCategoryRecommendationsWithAsinCounts{value: val, isSet: true}
}

func (v NullableCategoryRecommendationsWithAsinCounts) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCategoryRecommendationsWithAsinCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
