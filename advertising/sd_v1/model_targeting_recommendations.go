package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetingRecommendations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingRecommendations{}

// TargetingRecommendations A collection of targeting recommendations. Results will be sorted with strongest recommendations in the beginning.
type TargetingRecommendations struct {
	// List of recommended product targets
	Products []ProductRecommendation `json:"products,omitempty"`
}

// NewTargetingRecommendations instantiates a new TargetingRecommendations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingRecommendations() *TargetingRecommendations {
	this := TargetingRecommendations{}
	return &this
}

// NewTargetingRecommendationsWithDefaults instantiates a new TargetingRecommendations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingRecommendationsWithDefaults() *TargetingRecommendations {
	this := TargetingRecommendations{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *TargetingRecommendations) GetProducts() []ProductRecommendation {
	if o == nil || IsNil(o.Products) {
		var ret []ProductRecommendation
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingRecommendations) GetProductsOk() ([]ProductRecommendation, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *TargetingRecommendations) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []ProductRecommendation and assigns it to the Products field.
func (o *TargetingRecommendations) SetProducts(v []ProductRecommendation) {
	o.Products = v
}

func (o TargetingRecommendations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableTargetingRecommendations struct {
	value *TargetingRecommendations
	isSet bool
}

func (v NullableTargetingRecommendations) Get() *TargetingRecommendations {
	return v.value
}

func (v *NullableTargetingRecommendations) Set(val *TargetingRecommendations) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingRecommendations) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingRecommendations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingRecommendations(val *TargetingRecommendations) *NullableTargetingRecommendations {
	return &NullableTargetingRecommendations{value: val, isSet: true}
}

func (v NullableTargetingRecommendations) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingRecommendations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
