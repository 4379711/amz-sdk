package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsV31{}

// SDTargetingRecommendationsV31 struct for SDTargetingRecommendationsV31
type SDTargetingRecommendationsV31 struct {
	// List of recommended product targets
	Products []SDProductRecommendation `json:"products,omitempty"`
	// List of recommended category targets
	Categories []SDCategoryRecommendation `json:"categories,omitempty"`
}

// NewSDTargetingRecommendationsV31 instantiates a new SDTargetingRecommendationsV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsV31() *SDTargetingRecommendationsV31 {
	this := SDTargetingRecommendationsV31{}
	return &this
}

// NewSDTargetingRecommendationsV31WithDefaults instantiates a new SDTargetingRecommendationsV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsV31WithDefaults() *SDTargetingRecommendationsV31 {
	this := SDTargetingRecommendationsV31{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV31) GetProducts() []SDProductRecommendation {
	if o == nil || IsNil(o.Products) {
		var ret []SDProductRecommendation
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV31) GetProductsOk() ([]SDProductRecommendation, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV31) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDProductRecommendation and assigns it to the Products field.
func (o *SDTargetingRecommendationsV31) SetProducts(v []SDProductRecommendation) {
	o.Products = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsV31) GetCategories() []SDCategoryRecommendation {
	if o == nil || IsNil(o.Categories) {
		var ret []SDCategoryRecommendation
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsV31) GetCategoriesOk() ([]SDCategoryRecommendation, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsV31) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []SDCategoryRecommendation and assigns it to the Categories field.
func (o *SDTargetingRecommendationsV31) SetCategories(v []SDCategoryRecommendation) {
	o.Categories = v
}

func (o SDTargetingRecommendationsV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsV31 struct {
	value *SDTargetingRecommendationsV31
	isSet bool
}

func (v NullableSDTargetingRecommendationsV31) Get() *SDTargetingRecommendationsV31 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsV31) Set(val *SDTargetingRecommendationsV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsV31(val *SDTargetingRecommendationsV31) *NullableSDTargetingRecommendationsV31 {
	return &NullableSDTargetingRecommendationsV31{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
