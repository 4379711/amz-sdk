package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsThemes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsThemes{}

// SDTargetingRecommendationsThemes The themes used to refine the recommendations. Currently only contextual targeting themes are supported.
type SDTargetingRecommendationsThemes struct {
	// A list of themes for product targeting recommendations. If this list is empty, the service will return all the current available theme recommendations. Recommendations will be returned for each theme. If specified, each theme should only include unique expressions.
	Product []SDProductTargetingTheme `json:"product,omitempty"`
}

// NewSDTargetingRecommendationsThemes instantiates a new SDTargetingRecommendationsThemes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsThemes() *SDTargetingRecommendationsThemes {
	this := SDTargetingRecommendationsThemes{}
	return &this
}

// NewSDTargetingRecommendationsThemesWithDefaults instantiates a new SDTargetingRecommendationsThemes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsThemesWithDefaults() *SDTargetingRecommendationsThemes {
	this := SDTargetingRecommendationsThemes{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsThemes) GetProduct() []SDProductTargetingTheme {
	if o == nil || IsNil(o.Product) {
		var ret []SDProductTargetingTheme
		return ret
	}
	return o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsThemes) GetProductOk() ([]SDProductTargetingTheme, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsThemes) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given []SDProductTargetingTheme and assigns it to the Product field.
func (o *SDTargetingRecommendationsThemes) SetProduct(v []SDProductTargetingTheme) {
	o.Product = v
}

func (o SDTargetingRecommendationsThemes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsThemes struct {
	value *SDTargetingRecommendationsThemes
	isSet bool
}

func (v NullableSDTargetingRecommendationsThemes) Get() *SDTargetingRecommendationsThemes {
	return v.value
}

func (v *NullableSDTargetingRecommendationsThemes) Set(val *SDTargetingRecommendationsThemes) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsThemes) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsThemes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsThemes(val *SDTargetingRecommendationsThemes) *NullableSDTargetingRecommendationsThemes {
	return &NullableSDTargetingRecommendationsThemes{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsThemes) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsThemes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
