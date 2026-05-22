package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetTargetableAsinCountsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTargetableAsinCountsRequest{}

// GetTargetableAsinCountsRequest struct for GetTargetableAsinCountsRequest
type GetTargetableAsinCountsRequest struct {
	// List of Age Ranges. Use the GetRefinementsForCategory to retrieve Age Ranges. Age Ranges are only available for categories related to children's toys and games.
	AgeRanges []AgeRange `json:"ageRanges,omitempty"`
	// List of Brands.
	Brands []Brand `json:"brands,omitempty"`
	// List of Genres. Use the GetRefinementsForCategory to retrieve Genre Node IDs. Genres are only available for categories related to books.
	Genres []Genre `json:"genres,omitempty"`
	// Indicates if products have prime shipping
	IsPrimeShipping *bool        `json:"isPrimeShipping,omitempty"`
	RatingRange     *RatingRange `json:"ratingRange,omitempty"`
	// The category node id. Please use the GetTargetableCategories API or GetCategoryRecommendationsForASINs API to retrieve category IDs.
	Category   string      `json:"category"`
	PriceRange *PriceRange `json:"priceRange,omitempty"`
}

type _GetTargetableAsinCountsRequest GetTargetableAsinCountsRequest

// NewGetTargetableAsinCountsRequest instantiates a new GetTargetableAsinCountsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTargetableAsinCountsRequest(category string) *GetTargetableAsinCountsRequest {
	this := GetTargetableAsinCountsRequest{}
	this.Category = category
	return &this
}

// NewGetTargetableAsinCountsRequestWithDefaults instantiates a new GetTargetableAsinCountsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTargetableAsinCountsRequestWithDefaults() *GetTargetableAsinCountsRequest {
	this := GetTargetableAsinCountsRequest{}
	return &this
}

// GetAgeRanges returns the AgeRanges field value if set, zero value otherwise.
func (o *GetTargetableAsinCountsRequest) GetAgeRanges() []AgeRange {
	if o == nil || IsNil(o.AgeRanges) {
		var ret []AgeRange
		return ret
	}
	return o.AgeRanges
}

// GetAgeRangesOk returns a tuple with the AgeRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetableAsinCountsRequest) GetAgeRangesOk() ([]AgeRange, bool) {
	if o == nil || IsNil(o.AgeRanges) {
		return nil, false
	}
	return o.AgeRanges, true
}

// HasAgeRanges returns a boolean if a field has been set.
func (o *GetTargetableAsinCountsRequest) HasAgeRanges() bool {
	if o != nil && !IsNil(o.AgeRanges) {
		return true
	}

	return false
}

// SetAgeRanges gets a reference to the given []AgeRange and assigns it to the AgeRanges field.
func (o *GetTargetableAsinCountsRequest) SetAgeRanges(v []AgeRange) {
	o.AgeRanges = v
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *GetTargetableAsinCountsRequest) GetBrands() []Brand {
	if o == nil || IsNil(o.Brands) {
		var ret []Brand
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetableAsinCountsRequest) GetBrandsOk() ([]Brand, bool) {
	if o == nil || IsNil(o.Brands) {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *GetTargetableAsinCountsRequest) HasBrands() bool {
	if o != nil && !IsNil(o.Brands) {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []Brand and assigns it to the Brands field.
func (o *GetTargetableAsinCountsRequest) SetBrands(v []Brand) {
	o.Brands = v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *GetTargetableAsinCountsRequest) GetGenres() []Genre {
	if o == nil || IsNil(o.Genres) {
		var ret []Genre
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetableAsinCountsRequest) GetGenresOk() ([]Genre, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *GetTargetableAsinCountsRequest) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []Genre and assigns it to the Genres field.
func (o *GetTargetableAsinCountsRequest) SetGenres(v []Genre) {
	o.Genres = v
}

// GetIsPrimeShipping returns the IsPrimeShipping field value if set, zero value otherwise.
func (o *GetTargetableAsinCountsRequest) GetIsPrimeShipping() bool {
	if o == nil || IsNil(o.IsPrimeShipping) {
		var ret bool
		return ret
	}
	return *o.IsPrimeShipping
}

// GetIsPrimeShippingOk returns a tuple with the IsPrimeShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetableAsinCountsRequest) GetIsPrimeShippingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimeShipping) {
		return nil, false
	}
	return o.IsPrimeShipping, true
}

// HasIsPrimeShipping returns a boolean if a field has been set.
func (o *GetTargetableAsinCountsRequest) HasIsPrimeShipping() bool {
	if o != nil && !IsNil(o.IsPrimeShipping) {
		return true
	}

	return false
}

// SetIsPrimeShipping gets a reference to the given bool and assigns it to the IsPrimeShipping field.
func (o *GetTargetableAsinCountsRequest) SetIsPrimeShipping(v bool) {
	o.IsPrimeShipping = &v
}

// GetRatingRange returns the RatingRange field value if set, zero value otherwise.
func (o *GetTargetableAsinCountsRequest) GetRatingRange() RatingRange {
	if o == nil || IsNil(o.RatingRange) {
		var ret RatingRange
		return ret
	}
	return *o.RatingRange
}

// GetRatingRangeOk returns a tuple with the RatingRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetableAsinCountsRequest) GetRatingRangeOk() (*RatingRange, bool) {
	if o == nil || IsNil(o.RatingRange) {
		return nil, false
	}
	return o.RatingRange, true
}

// HasRatingRange returns a boolean if a field has been set.
func (o *GetTargetableAsinCountsRequest) HasRatingRange() bool {
	if o != nil && !IsNil(o.RatingRange) {
		return true
	}

	return false
}

// SetRatingRange gets a reference to the given RatingRange and assigns it to the RatingRange field.
func (o *GetTargetableAsinCountsRequest) SetRatingRange(v RatingRange) {
	o.RatingRange = &v
}

// GetCategory returns the Category field value
func (o *GetTargetableAsinCountsRequest) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *GetTargetableAsinCountsRequest) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *GetTargetableAsinCountsRequest) SetCategory(v string) {
	o.Category = v
}

// GetPriceRange returns the PriceRange field value if set, zero value otherwise.
func (o *GetTargetableAsinCountsRequest) GetPriceRange() PriceRange {
	if o == nil || IsNil(o.PriceRange) {
		var ret PriceRange
		return ret
	}
	return *o.PriceRange
}

// GetPriceRangeOk returns a tuple with the PriceRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetableAsinCountsRequest) GetPriceRangeOk() (*PriceRange, bool) {
	if o == nil || IsNil(o.PriceRange) {
		return nil, false
	}
	return o.PriceRange, true
}

// HasPriceRange returns a boolean if a field has been set.
func (o *GetTargetableAsinCountsRequest) HasPriceRange() bool {
	if o != nil && !IsNil(o.PriceRange) {
		return true
	}

	return false
}

// SetPriceRange gets a reference to the given PriceRange and assigns it to the PriceRange field.
func (o *GetTargetableAsinCountsRequest) SetPriceRange(v PriceRange) {
	o.PriceRange = &v
}

func (o GetTargetableAsinCountsRequest) ToMap() (map[string]interface{}, error) {
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

type NullableGetTargetableAsinCountsRequest struct {
	value *GetTargetableAsinCountsRequest
	isSet bool
}

func (v NullableGetTargetableAsinCountsRequest) Get() *GetTargetableAsinCountsRequest {
	return v.value
}

func (v *NullableGetTargetableAsinCountsRequest) Set(val *GetTargetableAsinCountsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTargetableAsinCountsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTargetableAsinCountsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTargetableAsinCountsRequest(val *GetTargetableAsinCountsRequest) *NullableGetTargetableAsinCountsRequest {
	return &NullableGetTargetableAsinCountsRequest{value: val, isSet: true}
}

func (v NullableGetTargetableAsinCountsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetTargetableAsinCountsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
