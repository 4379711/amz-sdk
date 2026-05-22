package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RefinementsLoP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefinementsLoP{}

// RefinementsLoP Response object for the POST /sp/targets/category/{categoryId}/refinements endpoint, containing information on Brand Nodes, Age Range Nodes, and Genre Nodes.
type RefinementsLoP struct {
	// List of Age Ranges in a language of preference (LoP). Use the POST /sp/targets/category/{categoryId}/refinements endpoint to retrieve Age Ranges. Age Ranges are only available for categories related to children's toys and games.
	AgeRanges []AgeRangeLoP `json:"ageRanges,omitempty"`
	// List of Brands.
	Brands []BrandLoP `json:"brands,omitempty"`
	// List of Genres in a language of preference (LoP). Use the POST /sp/targets/category/{categoryId}/refinements endpoint to retrieve Genre Node IDs. Genres are only available for categories related to books.
	Genres []GenreLoP `json:"genres,omitempty"`
}

// NewRefinementsLoP instantiates a new RefinementsLoP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefinementsLoP() *RefinementsLoP {
	this := RefinementsLoP{}
	return &this
}

// NewRefinementsLoPWithDefaults instantiates a new RefinementsLoP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefinementsLoPWithDefaults() *RefinementsLoP {
	this := RefinementsLoP{}
	return &this
}

// GetAgeRanges returns the AgeRanges field value if set, zero value otherwise.
func (o *RefinementsLoP) GetAgeRanges() []AgeRangeLoP {
	if o == nil || IsNil(o.AgeRanges) {
		var ret []AgeRangeLoP
		return ret
	}
	return o.AgeRanges
}

// GetAgeRangesOk returns a tuple with the AgeRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefinementsLoP) GetAgeRangesOk() ([]AgeRangeLoP, bool) {
	if o == nil || IsNil(o.AgeRanges) {
		return nil, false
	}
	return o.AgeRanges, true
}

// HasAgeRanges returns a boolean if a field has been set.
func (o *RefinementsLoP) HasAgeRanges() bool {
	if o != nil && !IsNil(o.AgeRanges) {
		return true
	}

	return false
}

// SetAgeRanges gets a reference to the given []AgeRangeLoP and assigns it to the AgeRanges field.
func (o *RefinementsLoP) SetAgeRanges(v []AgeRangeLoP) {
	o.AgeRanges = v
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *RefinementsLoP) GetBrands() []BrandLoP {
	if o == nil || IsNil(o.Brands) {
		var ret []BrandLoP
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefinementsLoP) GetBrandsOk() ([]BrandLoP, bool) {
	if o == nil || IsNil(o.Brands) {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *RefinementsLoP) HasBrands() bool {
	if o != nil && !IsNil(o.Brands) {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []BrandLoP and assigns it to the Brands field.
func (o *RefinementsLoP) SetBrands(v []BrandLoP) {
	o.Brands = v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *RefinementsLoP) GetGenres() []GenreLoP {
	if o == nil || IsNil(o.Genres) {
		var ret []GenreLoP
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefinementsLoP) GetGenresOk() ([]GenreLoP, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *RefinementsLoP) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []GenreLoP and assigns it to the Genres field.
func (o *RefinementsLoP) SetGenres(v []GenreLoP) {
	o.Genres = v
}

func (o RefinementsLoP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgeRanges) {
		toSerialize["ageRanges"] = o.AgeRanges
	}
	if !IsNil(o.Brands) {
		toSerialize["brands"] = o.Brands
	}
	if !IsNil(o.Genres) {
		toSerialize["genres"] = o.Genres
	}
	return toSerialize, nil
}

type NullableRefinementsLoP struct {
	value *RefinementsLoP
	isSet bool
}

func (v NullableRefinementsLoP) Get() *RefinementsLoP {
	return v.value
}

func (v *NullableRefinementsLoP) Set(val *RefinementsLoP) {
	v.value = val
	v.isSet = true
}

func (v NullableRefinementsLoP) IsSet() bool {
	return v.isSet
}

func (v *NullableRefinementsLoP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefinementsLoP(val *RefinementsLoP) *NullableRefinementsLoP {
	return &NullableRefinementsLoP{value: val, isSet: true}
}

func (v NullableRefinementsLoP) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRefinementsLoP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
