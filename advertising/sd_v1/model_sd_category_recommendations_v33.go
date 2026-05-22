package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDCategoryRecommendationsV33 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDCategoryRecommendationsV33{}

// SDCategoryRecommendationsV33 struct for SDCategoryRecommendationsV33
type SDCategoryRecommendationsV33 struct {
	// List of recommended category targets.
	Categories []SDCategoryRecommendationV33 `json:"categories,omitempty"`
}

// NewSDCategoryRecommendationsV33 instantiates a new SDCategoryRecommendationsV33 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDCategoryRecommendationsV33() *SDCategoryRecommendationsV33 {
	this := SDCategoryRecommendationsV33{}
	return &this
}

// NewSDCategoryRecommendationsV33WithDefaults instantiates a new SDCategoryRecommendationsV33 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDCategoryRecommendationsV33WithDefaults() *SDCategoryRecommendationsV33 {
	this := SDCategoryRecommendationsV33{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *SDCategoryRecommendationsV33) GetCategories() []SDCategoryRecommendationV33 {
	if o == nil || IsNil(o.Categories) {
		var ret []SDCategoryRecommendationV33
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationsV33) GetCategoriesOk() ([]SDCategoryRecommendationV33, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *SDCategoryRecommendationsV33) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []SDCategoryRecommendationV33 and assigns it to the Categories field.
func (o *SDCategoryRecommendationsV33) SetCategories(v []SDCategoryRecommendationV33) {
	o.Categories = v
}

func (o SDCategoryRecommendationsV33) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableSDCategoryRecommendationsV33 struct {
	value *SDCategoryRecommendationsV33
	isSet bool
}

func (v NullableSDCategoryRecommendationsV33) Get() *SDCategoryRecommendationsV33 {
	return v.value
}

func (v *NullableSDCategoryRecommendationsV33) Set(val *SDCategoryRecommendationsV33) {
	v.value = val
	v.isSet = true
}

func (v NullableSDCategoryRecommendationsV33) IsSet() bool {
	return v.isSet
}

func (v *NullableSDCategoryRecommendationsV33) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDCategoryRecommendationsV33(val *SDCategoryRecommendationsV33) *NullableSDCategoryRecommendationsV33 {
	return &NullableSDCategoryRecommendationsV33{value: val, isSet: true}
}

func (v NullableSDCategoryRecommendationsV33) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDCategoryRecommendationsV33) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
