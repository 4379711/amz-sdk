package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandRefinement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandRefinement{}

// BrandRefinement A brand that you can use to refine your search.
type BrandRefinement struct {
	// The estimated number of results that would be returned if you refine your search by the specified `brandName`.
	NumberOfResults int32 `json:"numberOfResults"`
	// The brand name that you can use to refine your search.
	BrandName string `json:"brandName"`
}

type _BrandRefinement BrandRefinement

// NewBrandRefinement instantiates a new BrandRefinement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandRefinement(numberOfResults int32, brandName string) *BrandRefinement {
	this := BrandRefinement{}
	this.NumberOfResults = numberOfResults
	this.BrandName = brandName
	return &this
}

// NewBrandRefinementWithDefaults instantiates a new BrandRefinement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandRefinementWithDefaults() *BrandRefinement {
	this := BrandRefinement{}
	return &this
}

// GetNumberOfResults returns the NumberOfResults field value
func (o *BrandRefinement) GetNumberOfResults() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfResults
}

// GetNumberOfResultsOk returns a tuple with the NumberOfResults field value
// and a boolean to check if the value has been set.
func (o *BrandRefinement) GetNumberOfResultsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfResults, true
}

// SetNumberOfResults sets field value
func (o *BrandRefinement) SetNumberOfResults(v int32) {
	o.NumberOfResults = v
}

// GetBrandName returns the BrandName field value
func (o *BrandRefinement) GetBrandName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value
// and a boolean to check if the value has been set.
func (o *BrandRefinement) GetBrandNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandName, true
}

// SetBrandName sets field value
func (o *BrandRefinement) SetBrandName(v string) {
	o.BrandName = v
}

func (o BrandRefinement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["numberOfResults"] = o.NumberOfResults
	toSerialize["brandName"] = o.BrandName
	return toSerialize, nil
}

type NullableBrandRefinement struct {
	value *BrandRefinement
	isSet bool
}

func (v NullableBrandRefinement) Get() *BrandRefinement {
	return v.value
}

func (v *NullableBrandRefinement) Set(val *BrandRefinement) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandRefinement) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandRefinement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandRefinement(val *BrandRefinement) *NullableBrandRefinement {
	return &NullableBrandRefinement{value: val, isSet: true}
}

func (v NullableBrandRefinement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandRefinement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
