package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDThemeRecommendationsV34Themes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDThemeRecommendationsV34Themes{}

// SDThemeRecommendationsV34Themes struct for SDThemeRecommendationsV34Themes
type SDThemeRecommendationsV34Themes struct {
	// A list of contextual targeting theme recommendations.
	Products []SDThemeRecommendationsV34ThemesProductsInner `json:"products,omitempty"`
}

// NewSDThemeRecommendationsV34Themes instantiates a new SDThemeRecommendationsV34Themes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDThemeRecommendationsV34Themes() *SDThemeRecommendationsV34Themes {
	this := SDThemeRecommendationsV34Themes{}
	return &this
}

// NewSDThemeRecommendationsV34ThemesWithDefaults instantiates a new SDThemeRecommendationsV34Themes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDThemeRecommendationsV34ThemesWithDefaults() *SDThemeRecommendationsV34Themes {
	this := SDThemeRecommendationsV34Themes{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDThemeRecommendationsV34Themes) GetProducts() []SDThemeRecommendationsV34ThemesProductsInner {
	if o == nil || IsNil(o.Products) {
		var ret []SDThemeRecommendationsV34ThemesProductsInner
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDThemeRecommendationsV34Themes) GetProductsOk() ([]SDThemeRecommendationsV34ThemesProductsInner, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDThemeRecommendationsV34Themes) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDThemeRecommendationsV34ThemesProductsInner and assigns it to the Products field.
func (o *SDThemeRecommendationsV34Themes) SetProducts(v []SDThemeRecommendationsV34ThemesProductsInner) {
	o.Products = v
}

func (o SDThemeRecommendationsV34Themes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableSDThemeRecommendationsV34Themes struct {
	value *SDThemeRecommendationsV34Themes
	isSet bool
}

func (v NullableSDThemeRecommendationsV34Themes) Get() *SDThemeRecommendationsV34Themes {
	return v.value
}

func (v *NullableSDThemeRecommendationsV34Themes) Set(val *SDThemeRecommendationsV34Themes) {
	v.value = val
	v.isSet = true
}

func (v NullableSDThemeRecommendationsV34Themes) IsSet() bool {
	return v.isSet
}

func (v *NullableSDThemeRecommendationsV34Themes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDThemeRecommendationsV34Themes(val *SDThemeRecommendationsV34Themes) *NullableSDThemeRecommendationsV34Themes {
	return &NullableSDThemeRecommendationsV34Themes{value: val, isSet: true}
}

func (v NullableSDThemeRecommendationsV34Themes) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDThemeRecommendationsV34Themes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
