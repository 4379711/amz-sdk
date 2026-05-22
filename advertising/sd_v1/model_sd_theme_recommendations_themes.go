package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDThemeRecommendationsThemes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDThemeRecommendationsThemes{}

// SDThemeRecommendationsThemes struct for SDThemeRecommendationsThemes
type SDThemeRecommendationsThemes struct {
	// A list of contextual targeting theme recommendations.
	Products []SDThemeRecommendationsThemesProductsInner `json:"products,omitempty"`
}

// NewSDThemeRecommendationsThemes instantiates a new SDThemeRecommendationsThemes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDThemeRecommendationsThemes() *SDThemeRecommendationsThemes {
	this := SDThemeRecommendationsThemes{}
	return &this
}

// NewSDThemeRecommendationsThemesWithDefaults instantiates a new SDThemeRecommendationsThemes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDThemeRecommendationsThemesWithDefaults() *SDThemeRecommendationsThemes {
	this := SDThemeRecommendationsThemes{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDThemeRecommendationsThemes) GetProducts() []SDThemeRecommendationsThemesProductsInner {
	if o == nil || IsNil(o.Products) {
		var ret []SDThemeRecommendationsThemesProductsInner
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDThemeRecommendationsThemes) GetProductsOk() ([]SDThemeRecommendationsThemesProductsInner, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDThemeRecommendationsThemes) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDThemeRecommendationsThemesProductsInner and assigns it to the Products field.
func (o *SDThemeRecommendationsThemes) SetProducts(v []SDThemeRecommendationsThemesProductsInner) {
	o.Products = v
}

func (o SDThemeRecommendationsThemes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableSDThemeRecommendationsThemes struct {
	value *SDThemeRecommendationsThemes
	isSet bool
}

func (v NullableSDThemeRecommendationsThemes) Get() *SDThemeRecommendationsThemes {
	return v.value
}

func (v *NullableSDThemeRecommendationsThemes) Set(val *SDThemeRecommendationsThemes) {
	v.value = val
	v.isSet = true
}

func (v NullableSDThemeRecommendationsThemes) IsSet() bool {
	return v.isSet
}

func (v *NullableSDThemeRecommendationsThemes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDThemeRecommendationsThemes(val *SDThemeRecommendationsThemes) *NullableSDThemeRecommendationsThemes {
	return &NullableSDThemeRecommendationsThemes{value: val, isSet: true}
}

func (v NullableSDThemeRecommendationsThemes) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDThemeRecommendationsThemes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
