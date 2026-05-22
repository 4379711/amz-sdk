package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the Refinements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Refinements{}

// Refinements Optional fields that you can use to refine your search results.
type Refinements struct {
	// A list of brands you can use to refine your search.
	Brands []BrandRefinement `json:"brands"`
	// A list of classifications you can use to refine your search.
	Classifications []ClassificationRefinement `json:"classifications"`
}

type _Refinements Refinements

// NewRefinements instantiates a new Refinements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefinements(brands []BrandRefinement, classifications []ClassificationRefinement) *Refinements {
	this := Refinements{}
	this.Brands = brands
	this.Classifications = classifications
	return &this
}

// NewRefinementsWithDefaults instantiates a new Refinements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefinementsWithDefaults() *Refinements {
	this := Refinements{}
	return &this
}

// GetBrands returns the Brands field value
func (o *Refinements) GetBrands() []BrandRefinement {
	if o == nil {
		var ret []BrandRefinement
		return ret
	}

	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value
// and a boolean to check if the value has been set.
func (o *Refinements) GetBrandsOk() ([]BrandRefinement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Brands, true
}

// SetBrands sets field value
func (o *Refinements) SetBrands(v []BrandRefinement) {
	o.Brands = v
}

// GetClassifications returns the Classifications field value
func (o *Refinements) GetClassifications() []ClassificationRefinement {
	if o == nil {
		var ret []ClassificationRefinement
		return ret
	}

	return o.Classifications
}

// GetClassificationsOk returns a tuple with the Classifications field value
// and a boolean to check if the value has been set.
func (o *Refinements) GetClassificationsOk() ([]ClassificationRefinement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Classifications, true
}

// SetClassifications sets field value
func (o *Refinements) SetClassifications(v []ClassificationRefinement) {
	o.Classifications = v
}

func (o Refinements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["brands"] = o.Brands
	toSerialize["classifications"] = o.Classifications
	return toSerialize, nil
}

type NullableRefinements struct {
	value *Refinements
	isSet bool
}

func (v NullableRefinements) Get() *Refinements {
	return v.value
}

func (v *NullableRefinements) Set(val *Refinements) {
	v.value = val
	v.isSet = true
}

func (v NullableRefinements) IsSet() bool {
	return v.isSet
}

func (v *NullableRefinements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefinements(val *Refinements) *NullableRefinements {
	return &NullableRefinements{value: val, isSet: true}
}

func (v NullableRefinements) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRefinements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
