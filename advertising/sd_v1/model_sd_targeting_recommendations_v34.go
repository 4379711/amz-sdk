package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsV34 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsV34{}

// SDTargetingRecommendationsV34 For v3.4 the service will continue to return the recommendations returned for v3.2, return audience recommendations if requested, and return the theme expression used in product targeting if requested.
type SDTargetingRecommendationsV34 struct {
	// List of recommended product targets
	Products []SDProductRecommendationsV32 `json:"products,omitempty"`
	// List of recommended category targets
	Categories []SDCategoryRecommendationV33 `json:"categories,omitempty"`
	// List of recommended audience targets, broken down by audience category
	Audiences []SDAudienceCategoryRecommendations `json:"audiences,omitempty"`
	Themes    *SDThemeRecommendationsV34          `json:"themes,omitempty"`
}

// NewSDTargetingRecommendationsV34 instantiates a new SDTargetingRecommendationsV34 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsV34() *SDTargetingRecommendationsV34 {
	this := SDTargetingRecommendationsV34{}
	return &this
}

// NewSDTargetingRecommendationsV34WithDefaults instantiates a new SDTargetingRecommendationsV34 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsV34WithDefaults() *SDTargetingRecommendationsV34 {
	this := SDTargetingRecommendationsV34{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV34) GetProducts() []SDProductRecommendationsV32 {
	if o == nil || IsNil(o.Products) {
		var ret []SDProductRecommendationsV32
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV34) GetProductsOk() ([]SDProductRecommendationsV32, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV34) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDProductRecommendationsV32 and assigns it to the Products field.
func (o *SDTargetingRecommendationsV34) SetProducts(v []SDProductRecommendationsV32) {
	o.Products = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV34) GetCategories() []SDCategoryRecommendationV33 {
	if o == nil || IsNil(o.Categories) {
		var ret []SDCategoryRecommendationV33
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV34) GetCategoriesOk() ([]SDCategoryRecommendationV33, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV34) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []SDCategoryRecommendationV33 and assigns it to the Categories field.
func (o *SDTargetingRecommendationsV34) SetCategories(v []SDCategoryRecommendationV33) {
	o.Categories = v
}

// GetAudiences returns the Audiences field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV34) GetAudiences() []SDAudienceCategoryRecommendations {
	if o == nil || IsNil(o.Audiences) {
		var ret []SDAudienceCategoryRecommendations
		return ret
	}
	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV34) GetAudiencesOk() ([]SDAudienceCategoryRecommendations, bool) {
	if o == nil || IsNil(o.Audiences) {
		return nil, false
	}
	return o.Audiences, true
}

// HasAudiences returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV34) HasAudiences() bool {
	if o != nil && !IsNil(o.Audiences) {
		return true
	}

	return false
}

// SetAudiences gets a reference to the given []SDAudienceCategoryRecommendations and assigns it to the Audiences field.
func (o *SDTargetingRecommendationsV34) SetAudiences(v []SDAudienceCategoryRecommendations) {
	o.Audiences = v
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV34) GetThemes() SDThemeRecommendationsV34 {
	if o == nil || IsNil(o.Themes) {
		var ret SDThemeRecommendationsV34
		return ret
	}
	return *o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV34) GetThemesOk() (*SDThemeRecommendationsV34, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV34) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given SDThemeRecommendationsV34 and assigns it to the Themes field.
func (o *SDTargetingRecommendationsV34) SetThemes(v SDThemeRecommendationsV34) {
	o.Themes = &v
}

func (o SDTargetingRecommendationsV34) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Audiences) {
		toSerialize["audiences"] = o.Audiences
	}
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsV34 struct {
	value *SDTargetingRecommendationsV34
	isSet bool
}

func (v NullableSDTargetingRecommendationsV34) Get() *SDTargetingRecommendationsV34 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsV34) Set(val *SDTargetingRecommendationsV34) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsV34) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsV34) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsV34(val *SDTargetingRecommendationsV34) *NullableSDTargetingRecommendationsV34 {
	return &NullableSDTargetingRecommendationsV34{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsV34) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsV34) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
