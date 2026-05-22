package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsV35 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsV35{}

// SDTargetingRecommendationsV35 For v3.5 the service will continue to return the recommendations returned for v3.4, return Entertainment targeting recommendations if requested and return asin-less recommendations if a landing page URL was passed in the request
type SDTargetingRecommendationsV35 struct {
	// List of recommended product targets
	Products []SDProductRecommendationsV32 `json:"products,omitempty"`
	// List of recommended category targets
	Categories []SDCategoryRecommendationV33 `json:"categories,omitempty"`
	// List of recommended audience targets, broken down by audience category
	Audiences []SDAudienceCategoryRecommendations `json:"audiences,omitempty"`
	// List of recommended entertainment targets
	ContentCategories []SDContentCategoryRecommendations `json:"contentCategories,omitempty"`
	Themes            *SDThemeRecommendationsV34         `json:"themes,omitempty"`
}

// NewSDTargetingRecommendationsV35 instantiates a new SDTargetingRecommendationsV35 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsV35() *SDTargetingRecommendationsV35 {
	this := SDTargetingRecommendationsV35{}
	return &this
}

// NewSDTargetingRecommendationsV35WithDefaults instantiates a new SDTargetingRecommendationsV35 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsV35WithDefaults() *SDTargetingRecommendationsV35 {
	this := SDTargetingRecommendationsV35{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV35) GetProducts() []SDProductRecommendationsV32 {
	if o == nil || IsNil(o.Products) {
		var ret []SDProductRecommendationsV32
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV35) GetProductsOk() ([]SDProductRecommendationsV32, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV35) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDProductRecommendationsV32 and assigns it to the Products field.
func (o *SDTargetingRecommendationsV35) SetProducts(v []SDProductRecommendationsV32) {
	o.Products = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV35) GetCategories() []SDCategoryRecommendationV33 {
	if o == nil || IsNil(o.Categories) {
		var ret []SDCategoryRecommendationV33
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV35) GetCategoriesOk() ([]SDCategoryRecommendationV33, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV35) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []SDCategoryRecommendationV33 and assigns it to the Categories field.
func (o *SDTargetingRecommendationsV35) SetCategories(v []SDCategoryRecommendationV33) {
	o.Categories = v
}

// GetAudiences returns the Audiences field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV35) GetAudiences() []SDAudienceCategoryRecommendations {
	if o == nil || IsNil(o.Audiences) {
		var ret []SDAudienceCategoryRecommendations
		return ret
	}
	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV35) GetAudiencesOk() ([]SDAudienceCategoryRecommendations, bool) {
	if o == nil || IsNil(o.Audiences) {
		return nil, false
	}
	return o.Audiences, true
}

// HasAudiences returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV35) HasAudiences() bool {
	if o != nil && !IsNil(o.Audiences) {
		return true
	}

	return false
}

// SetAudiences gets a reference to the given []SDAudienceCategoryRecommendations and assigns it to the Audiences field.
func (o *SDTargetingRecommendationsV35) SetAudiences(v []SDAudienceCategoryRecommendations) {
	o.Audiences = v
}

// GetContentCategories returns the ContentCategories field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV35) GetContentCategories() []SDContentCategoryRecommendations {
	if o == nil || IsNil(o.ContentCategories) {
		var ret []SDContentCategoryRecommendations
		return ret
	}
	return o.ContentCategories
}

// GetContentCategoriesOk returns a tuple with the ContentCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV35) GetContentCategoriesOk() ([]SDContentCategoryRecommendations, bool) {
	if o == nil || IsNil(o.ContentCategories) {
		return nil, false
	}
	return o.ContentCategories, true
}

// HasContentCategories returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV35) HasContentCategories() bool {
	if o != nil && !IsNil(o.ContentCategories) {
		return true
	}

	return false
}

// SetContentCategories gets a reference to the given []SDContentCategoryRecommendations and assigns it to the ContentCategories field.
func (o *SDTargetingRecommendationsV35) SetContentCategories(v []SDContentCategoryRecommendations) {
	o.ContentCategories = v
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV35) GetThemes() SDThemeRecommendationsV34 {
	if o == nil || IsNil(o.Themes) {
		var ret SDThemeRecommendationsV34
		return ret
	}
	return *o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV35) GetThemesOk() (*SDThemeRecommendationsV34, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV35) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given SDThemeRecommendationsV34 and assigns it to the Themes field.
func (o *SDTargetingRecommendationsV35) SetThemes(v SDThemeRecommendationsV34) {
	o.Themes = &v
}

func (o SDTargetingRecommendationsV35) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ContentCategories) {
		toSerialize["contentCategories"] = o.ContentCategories
	}
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsV35 struct {
	value *SDTargetingRecommendationsV35
	isSet bool
}

func (v NullableSDTargetingRecommendationsV35) Get() *SDTargetingRecommendationsV35 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsV35) Set(val *SDTargetingRecommendationsV35) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsV35) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsV35) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsV35(val *SDTargetingRecommendationsV35) *NullableSDTargetingRecommendationsV35 {
	return &NullableSDTargetingRecommendationsV35{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsV35) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsV35) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
