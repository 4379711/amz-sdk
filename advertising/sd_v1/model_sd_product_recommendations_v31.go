package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDProductRecommendationsV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDProductRecommendationsV31{}

// SDProductRecommendationsV31 struct for SDProductRecommendationsV31
type SDProductRecommendationsV31 struct {
	// List of recommended product targets
	Products []SDProductRecommendation `json:"products,omitempty"`
}

// NewSDProductRecommendationsV31 instantiates a new SDProductRecommendationsV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDProductRecommendationsV31() *SDProductRecommendationsV31 {
	this := SDProductRecommendationsV31{}
	return &this
}

// NewSDProductRecommendationsV31WithDefaults instantiates a new SDProductRecommendationsV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDProductRecommendationsV31WithDefaults() *SDProductRecommendationsV31 {
	this := SDProductRecommendationsV31{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDProductRecommendationsV31) GetProducts() []SDProductRecommendation {
	if o == nil || IsNil(o.Products) {
		var ret []SDProductRecommendation
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductRecommendationsV31) GetProductsOk() ([]SDProductRecommendation, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDProductRecommendationsV31) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDProductRecommendation and assigns it to the Products field.
func (o *SDProductRecommendationsV31) SetProducts(v []SDProductRecommendation) {
	o.Products = v
}

func (o SDProductRecommendationsV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableSDProductRecommendationsV31 struct {
	value *SDProductRecommendationsV31
	isSet bool
}

func (v NullableSDProductRecommendationsV31) Get() *SDProductRecommendationsV31 {
	return v.value
}

func (v *NullableSDProductRecommendationsV31) Set(val *SDProductRecommendationsV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDProductRecommendationsV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDProductRecommendationsV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDProductRecommendationsV31(val *SDProductRecommendationsV31) *NullableSDProductRecommendationsV31 {
	return &NullableSDProductRecommendationsV31{value: val, isSet: true}
}

func (v NullableSDProductRecommendationsV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDProductRecommendationsV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
