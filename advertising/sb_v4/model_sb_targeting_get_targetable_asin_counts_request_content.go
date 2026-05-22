package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingGetTargetableASINCountsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingGetTargetableASINCountsRequestContent{}

// SBTargetingGetTargetableASINCountsRequestContent struct for SBTargetingGetTargetableASINCountsRequestContent
type SBTargetingGetTargetableASINCountsRequestContent struct {
	// List of Age Range Refinement Ids.
	AgeRanges []string `json:"ageRanges,omitempty"`
	// List of Brand Refinement Ids.
	Brands []string `json:"brands,omitempty"`
	// List of Genre Refinement Ids.
	Genres []string `json:"genres,omitempty"`
	// Indicates if products have prime shipping. Leave empty to include both prime shipping and non-prime shipping products.
	IsPrimeShipping *bool                   `json:"isPrimeShipping,omitempty"`
	RatingRange     *SBTargetingRatingRange `json:"ratingRange,omitempty"`
	// The category refinement id. Please use /sb/targets/categories or /sb/recommendations/targets/category to retrieve category IDs.
	Category   string                 `json:"category"`
	PriceRange *SBTargetingPriceRange `json:"priceRange,omitempty"`
}

type _SBTargetingGetTargetableASINCountsRequestContent SBTargetingGetTargetableASINCountsRequestContent

// NewSBTargetingGetTargetableASINCountsRequestContent instantiates a new SBTargetingGetTargetableASINCountsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingGetTargetableASINCountsRequestContent(category string) *SBTargetingGetTargetableASINCountsRequestContent {
	this := SBTargetingGetTargetableASINCountsRequestContent{}
	this.Category = category
	return &this
}

// NewSBTargetingGetTargetableASINCountsRequestContentWithDefaults instantiates a new SBTargetingGetTargetableASINCountsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingGetTargetableASINCountsRequestContentWithDefaults() *SBTargetingGetTargetableASINCountsRequestContent {
	this := SBTargetingGetTargetableASINCountsRequestContent{}
	return &this
}

// GetAgeRanges returns the AgeRanges field value if set, zero value otherwise.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetAgeRanges() []string {
	if o == nil || IsNil(o.AgeRanges) {
		var ret []string
		return ret
	}
	return o.AgeRanges
}

// GetAgeRangesOk returns a tuple with the AgeRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetAgeRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.AgeRanges) {
		return nil, false
	}
	return o.AgeRanges, true
}

// HasAgeRanges returns a boolean if a field has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) HasAgeRanges() bool {
	if o != nil && !IsNil(o.AgeRanges) {
		return true
	}

	return false
}

// SetAgeRanges gets a reference to the given []string and assigns it to the AgeRanges field.
func (o *SBTargetingGetTargetableASINCountsRequestContent) SetAgeRanges(v []string) {
	o.AgeRanges = v
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetBrands() []string {
	if o == nil || IsNil(o.Brands) {
		var ret []string
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetBrandsOk() ([]string, bool) {
	if o == nil || IsNil(o.Brands) {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) HasBrands() bool {
	if o != nil && !IsNil(o.Brands) {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []string and assigns it to the Brands field.
func (o *SBTargetingGetTargetableASINCountsRequestContent) SetBrands(v []string) {
	o.Brands = v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetGenres() []string {
	if o == nil || IsNil(o.Genres) {
		var ret []string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetGenresOk() ([]string, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *SBTargetingGetTargetableASINCountsRequestContent) SetGenres(v []string) {
	o.Genres = v
}

// GetIsPrimeShipping returns the IsPrimeShipping field value if set, zero value otherwise.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetIsPrimeShipping() bool {
	if o == nil || IsNil(o.IsPrimeShipping) {
		var ret bool
		return ret
	}
	return *o.IsPrimeShipping
}

// GetIsPrimeShippingOk returns a tuple with the IsPrimeShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetIsPrimeShippingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimeShipping) {
		return nil, false
	}
	return o.IsPrimeShipping, true
}

// HasIsPrimeShipping returns a boolean if a field has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) HasIsPrimeShipping() bool {
	if o != nil && !IsNil(o.IsPrimeShipping) {
		return true
	}

	return false
}

// SetIsPrimeShipping gets a reference to the given bool and assigns it to the IsPrimeShipping field.
func (o *SBTargetingGetTargetableASINCountsRequestContent) SetIsPrimeShipping(v bool) {
	o.IsPrimeShipping = &v
}

// GetRatingRange returns the RatingRange field value if set, zero value otherwise.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetRatingRange() SBTargetingRatingRange {
	if o == nil || IsNil(o.RatingRange) {
		var ret SBTargetingRatingRange
		return ret
	}
	return *o.RatingRange
}

// GetRatingRangeOk returns a tuple with the RatingRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetRatingRangeOk() (*SBTargetingRatingRange, bool) {
	if o == nil || IsNil(o.RatingRange) {
		return nil, false
	}
	return o.RatingRange, true
}

// HasRatingRange returns a boolean if a field has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) HasRatingRange() bool {
	if o != nil && !IsNil(o.RatingRange) {
		return true
	}

	return false
}

// SetRatingRange gets a reference to the given SBTargetingRatingRange and assigns it to the RatingRange field.
func (o *SBTargetingGetTargetableASINCountsRequestContent) SetRatingRange(v SBTargetingRatingRange) {
	o.RatingRange = &v
}

// GetCategory returns the Category field value
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *SBTargetingGetTargetableASINCountsRequestContent) SetCategory(v string) {
	o.Category = v
}

// GetPriceRange returns the PriceRange field value if set, zero value otherwise.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetPriceRange() SBTargetingPriceRange {
	if o == nil || IsNil(o.PriceRange) {
		var ret SBTargetingPriceRange
		return ret
	}
	return *o.PriceRange
}

// GetPriceRangeOk returns a tuple with the PriceRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) GetPriceRangeOk() (*SBTargetingPriceRange, bool) {
	if o == nil || IsNil(o.PriceRange) {
		return nil, false
	}
	return o.PriceRange, true
}

// HasPriceRange returns a boolean if a field has been set.
func (o *SBTargetingGetTargetableASINCountsRequestContent) HasPriceRange() bool {
	if o != nil && !IsNil(o.PriceRange) {
		return true
	}

	return false
}

// SetPriceRange gets a reference to the given SBTargetingPriceRange and assigns it to the PriceRange field.
func (o *SBTargetingGetTargetableASINCountsRequestContent) SetPriceRange(v SBTargetingPriceRange) {
	o.PriceRange = &v
}

func (o SBTargetingGetTargetableASINCountsRequestContent) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsPrimeShipping) {
		toSerialize["isPrimeShipping"] = o.IsPrimeShipping
	}
	if !IsNil(o.RatingRange) {
		toSerialize["ratingRange"] = o.RatingRange
	}
	toSerialize["category"] = o.Category
	if !IsNil(o.PriceRange) {
		toSerialize["priceRange"] = o.PriceRange
	}
	return toSerialize, nil
}

type NullableSBTargetingGetTargetableASINCountsRequestContent struct {
	value *SBTargetingGetTargetableASINCountsRequestContent
	isSet bool
}

func (v NullableSBTargetingGetTargetableASINCountsRequestContent) Get() *SBTargetingGetTargetableASINCountsRequestContent {
	return v.value
}

func (v *NullableSBTargetingGetTargetableASINCountsRequestContent) Set(val *SBTargetingGetTargetableASINCountsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingGetTargetableASINCountsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingGetTargetableASINCountsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingGetTargetableASINCountsRequestContent(val *SBTargetingGetTargetableASINCountsRequestContent) *NullableSBTargetingGetTargetableASINCountsRequestContent {
	return &NullableSBTargetingGetTargetableASINCountsRequestContent{value: val, isSet: true}
}

func (v NullableSBTargetingGetTargetableASINCountsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingGetTargetableASINCountsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
