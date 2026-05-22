package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsV32 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsV32{}

// SDTargetingRecommendationsV32 For v3.2 the service will continue to return the recommendations returned for v3.1 in products field, and return recommendations for contextual targeting themes in themes field.
type SDTargetingRecommendationsV32 struct {
	// List of recommended product targets
	Products []SDProductRecommendationV32 `json:"products,omitempty"`
	// List of recommended category targets
	Categories []SDCategoryRecommendation    `json:"categories,omitempty"`
	Themes     *SDThemeRecommendationsThemes `json:"themes,omitempty"`
}

// NewSDTargetingRecommendationsV32 instantiates a new SDTargetingRecommendationsV32 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsV32() *SDTargetingRecommendationsV32 {
	this := SDTargetingRecommendationsV32{}
	return &this
}

// NewSDTargetingRecommendationsV32WithDefaults instantiates a new SDTargetingRecommendationsV32 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsV32WithDefaults() *SDTargetingRecommendationsV32 {
	this := SDTargetingRecommendationsV32{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV32) GetProducts() []SDProductRecommendationV32 {
	if o == nil || IsNil(o.Products) {
		var ret []SDProductRecommendationV32
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV32) GetProductsOk() ([]SDProductRecommendationV32, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV32) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDProductRecommendationV32 and assigns it to the Products field.
func (o *SDTargetingRecommendationsV32) SetProducts(v []SDProductRecommendationV32) {
	o.Products = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV32) GetCategories() []SDCategoryRecommendation {
	if o == nil || IsNil(o.Categories) {
		var ret []SDCategoryRecommendation
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV32) GetCategoriesOk() ([]SDCategoryRecommendation, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV32) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []SDCategoryRecommendation and assigns it to the Categories field.
func (o *SDTargetingRecommendationsV32) SetCategories(v []SDCategoryRecommendation) {
	o.Categories = v
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV32) GetThemes() SDThemeRecommendationsThemes {
	if o == nil || IsNil(o.Themes) {
		var ret SDThemeRecommendationsThemes
		return ret
	}
	return *o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV32) GetThemesOk() (*SDThemeRecommendationsThemes, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV32) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given SDThemeRecommendationsThemes and assigns it to the Themes field.
func (o *SDTargetingRecommendationsV32) SetThemes(v SDThemeRecommendationsThemes) {
	o.Themes = &v
}

func (o SDTargetingRecommendationsV32) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsV32 struct {
	value *SDTargetingRecommendationsV32
	isSet bool
}

func (v NullableSDTargetingRecommendationsV32) Get() *SDTargetingRecommendationsV32 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsV32) Set(val *SDTargetingRecommendationsV32) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsV32) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsV32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsV32(val *SDTargetingRecommendationsV32) *NullableSDTargetingRecommendationsV32 {
	return &NullableSDTargetingRecommendationsV32{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsV32) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsV32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
