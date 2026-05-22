package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDCategoryRecommendations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDCategoryRecommendations{}

// SDCategoryRecommendations struct for SDCategoryRecommendations
type SDCategoryRecommendations struct {
	// List of recommended category targets
	Categories []SDCategoryRecommendation `json:"categories,omitempty"`
}

// NewSDCategoryRecommendations instantiates a new SDCategoryRecommendations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDCategoryRecommendations() *SDCategoryRecommendations {
	this := SDCategoryRecommendations{}
	return &this
}

// NewSDCategoryRecommendationsWithDefaults instantiates a new SDCategoryRecommendations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDCategoryRecommendationsWithDefaults() *SDCategoryRecommendations {
	this := SDCategoryRecommendations{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *SDCategoryRecommendations) GetCategories() []SDCategoryRecommendation {
	if o == nil || IsNil(o.Categories) {
		var ret []SDCategoryRecommendation
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendations) GetCategoriesOk() ([]SDCategoryRecommendation, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *SDCategoryRecommendations) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []SDCategoryRecommendation and assigns it to the Categories field.
func (o *SDCategoryRecommendations) SetCategories(v []SDCategoryRecommendation) {
	o.Categories = v
}

func (o SDCategoryRecommendations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableSDCategoryRecommendations struct {
	value *SDCategoryRecommendations
	isSet bool
}

func (v NullableSDCategoryRecommendations) Get() *SDCategoryRecommendations {
	return v.value
}

func (v *NullableSDCategoryRecommendations) Set(val *SDCategoryRecommendations) {
	v.value = val
	v.isSet = true
}

func (v NullableSDCategoryRecommendations) IsSet() bool {
	return v.isSet
}

func (v *NullableSDCategoryRecommendations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDCategoryRecommendations(val *SDCategoryRecommendations) *NullableSDCategoryRecommendations {
	return &NullableSDCategoryRecommendations{value: val, isSet: true}
}

func (v NullableSDCategoryRecommendations) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDCategoryRecommendations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
