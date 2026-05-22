package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendations{}

// SDTargetingRecommendations A collection of targeting recommendations. Results will be sorted with strongest recommendations in the beginning.
type SDTargetingRecommendations struct {
	// List of recommended product targets
	Products []SDProductRecommendation `json:"products,omitempty"`
}

// NewSDTargetingRecommendations instantiates a new SDTargetingRecommendations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendations() *SDTargetingRecommendations {
	this := SDTargetingRecommendations{}
	return &this
}

// NewSDTargetingRecommendationsWithDefaults instantiates a new SDTargetingRecommendations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsWithDefaults() *SDTargetingRecommendations {
	this := SDTargetingRecommendations{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *SDTargetingRecommendations) GetProducts() []SDProductRecommendation {
	if o == nil || IsNil(o.Products) {
		var ret []SDProductRecommendation
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendations) GetProductsOk() ([]SDProductRecommendation, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *SDTargetingRecommendations) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SDProductRecommendation and assigns it to the Products field.
func (o *SDTargetingRecommendations) SetProducts(v []SDProductRecommendation) {
	o.Products = v
}

func (o SDTargetingRecommendations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendations struct {
	value *SDTargetingRecommendations
	isSet bool
}

func (v NullableSDTargetingRecommendations) Get() *SDTargetingRecommendations {
	return v.value
}

func (v *NullableSDTargetingRecommendations) Set(val *SDTargetingRecommendations) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendations) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendations(val *SDTargetingRecommendations) *NullableSDTargetingRecommendations {
	return &NullableSDTargetingRecommendations{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendations) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
