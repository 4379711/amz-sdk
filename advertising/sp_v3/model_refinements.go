package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the Refinements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Refinements{}

// Refinements Response object for the GetRefinementsForCategory API, containing information on Brand Nodes, Age Range Nodes, and Genre Nodes.
type Refinements struct {
	// List of Age Ranges. Use the GetRefinementsForCategory to retrieve Age Ranges. Age Ranges are only available for categories related to children's toys and games.
	AgeRanges []AgeRange `json:"ageRanges,omitempty"`
	// List of Brands.
	Brands []Brand `json:"brands,omitempty"`
	// List of Genres. Use the GetRefinementsForCategory to retrieve Genre Node IDs. Genres are only available for categories related to books.
	Genres []Genre `json:"genres,omitempty"`
}

// NewRefinements instantiates a new Refinements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefinements() *Refinements {
	this := Refinements{}
	return &this
}

// NewRefinementsWithDefaults instantiates a new Refinements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefinementsWithDefaults() *Refinements {
	this := Refinements{}
	return &this
}

// GetAgeRanges returns the AgeRanges field value if set, zero value otherwise.
func (o *Refinements) GetAgeRanges() []AgeRange {
	if o == nil || IsNil(o.AgeRanges) {
		var ret []AgeRange
		return ret
	}
	return o.AgeRanges
}

// GetAgeRangesOk returns a tuple with the AgeRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refinements) GetAgeRangesOk() ([]AgeRange, bool) {
	if o == nil || IsNil(o.AgeRanges) {
		return nil, false
	}
	return o.AgeRanges, true
}

// HasAgeRanges returns a boolean if a field has been set.
func (o *Refinements) HasAgeRanges() bool {
	if o != nil && !IsNil(o.AgeRanges) {
		return true
	}

	return false
}

// SetAgeRanges gets a reference to the given []AgeRange and assigns it to the AgeRanges field.
func (o *Refinements) SetAgeRanges(v []AgeRange) {
	o.AgeRanges = v
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *Refinements) GetBrands() []Brand {
	if o == nil || IsNil(o.Brands) {
		var ret []Brand
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refinements) GetBrandsOk() ([]Brand, bool) {
	if o == nil || IsNil(o.Brands) {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *Refinements) HasBrands() bool {
	if o != nil && !IsNil(o.Brands) {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []Brand and assigns it to the Brands field.
func (o *Refinements) SetBrands(v []Brand) {
	o.Brands = v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *Refinements) GetGenres() []Genre {
	if o == nil || IsNil(o.Genres) {
		var ret []Genre
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refinements) GetGenresOk() ([]Genre, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *Refinements) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []Genre and assigns it to the Genres field.
func (o *Refinements) SetGenres(v []Genre) {
	o.Genres = v
}

func (o Refinements) ToMap() (map[string]interface{}, error) {
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
