package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductRecommendation{}

// ProductRecommendation A recommended product to target ads on
type ProductRecommendation struct {
	// Amazon Standard Identification Number
	Asin *string `json:"asin,omitempty" validate:"regexp=[a-zA-Z0-9]{10}"`
	// A rank to signify which recommendations are weighed more heavily, with a lower rank signifying a stronger recommendation
	Rank *int32 `json:"rank,omitempty"`
}

// NewProductRecommendation instantiates a new ProductRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductRecommendation() *ProductRecommendation {
	this := ProductRecommendation{}
	return &this
}

// NewProductRecommendationWithDefaults instantiates a new ProductRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductRecommendationWithDefaults() *ProductRecommendation {
	this := ProductRecommendation{}
	return &this
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ProductRecommendation) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecommendation) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ProductRecommendation) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ProductRecommendation) SetAsin(v string) {
	o.Asin = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *ProductRecommendation) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecommendation) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *ProductRecommendation) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *ProductRecommendation) SetRank(v int32) {
	o.Rank = &v
}

func (o ProductRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	return toSerialize, nil
}

type NullableProductRecommendation struct {
	value *ProductRecommendation
	isSet bool
}

func (v NullableProductRecommendation) Get() *ProductRecommendation {
	return v.value
}

func (v *NullableProductRecommendation) Set(val *ProductRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableProductRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableProductRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductRecommendation(val *ProductRecommendation) *NullableProductRecommendation {
	return &NullableProductRecommendation{value: val, isSet: true}
}

func (v NullableProductRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
