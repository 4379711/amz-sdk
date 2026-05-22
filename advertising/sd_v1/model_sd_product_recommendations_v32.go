package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDProductRecommendationsV32 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDProductRecommendationsV32{}

// SDProductRecommendationsV32 struct for SDProductRecommendationsV32
type SDProductRecommendationsV32 struct {
	// List of recommended product targets
	Products []SDProductRecommendationV32 `json:"products,omitempty"`
}

// NewSDProductRecommendationsV32 instantiates a new SDProductRecommendationsV32 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDProductRecommendationsV32() *SDProductRecommendationsV32 {
	this := SDProductRecommendationsV32{}
	return &this
}

// NewSDProductRecommendationsV32WithDefaults instantiates a new SDProductRecommendationsV32 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDProductRecommendationsV32WithDefaults() *SDProductRecommendationsV32 {
	this := SDProductRecommendationsV32{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDProductRecommendationsV32) GetProducts() []SDProductRecommendationV32 {
	if o == nil || IsNil(o.Products) {
		var ret []SDProductRecommendationV32
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductRecommendationsV32) GetProductsOk() ([]SDProductRecommendationV32, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDProductRecommendationsV32) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDProductRecommendationV32 and assigns it to the Products field.
func (o *SDProductRecommendationsV32) SetProducts(v []SDProductRecommendationV32) {
	o.Products = v
}

func (o SDProductRecommendationsV32) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableSDProductRecommendationsV32 struct {
	value *SDProductRecommendationsV32
	isSet bool
}

func (v NullableSDProductRecommendationsV32) Get() *SDProductRecommendationsV32 {
	return v.value
}

func (v *NullableSDProductRecommendationsV32) Set(val *SDProductRecommendationsV32) {
	v.value = val
	v.isSet = true
}

func (v NullableSDProductRecommendationsV32) IsSet() bool {
	return v.isSet
}

func (v *NullableSDProductRecommendationsV32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDProductRecommendationsV32(val *SDProductRecommendationsV32) *NullableSDProductRecommendationsV32 {
	return &NullableSDProductRecommendationsV32{value: val, isSet: true}
}

func (v NullableSDProductRecommendationsV32) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDProductRecommendationsV32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
