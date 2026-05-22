package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingGetRefinementsForCategoryResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingGetRefinementsForCategoryResponseContent{}

// SBTargetingGetRefinementsForCategoryResponseContent Response object for /sb/targets/categories/{categoryRefinementId}/refinements containing information on Brand Nodes, Age Range Nodes, and Genre Nodes.     Response is paginated with pagination occurring for all three arrays at once.     Example: If there are 800 brands, 5 age ranges, and 600 genres, the first response will return 500 brands, 5 age ranges, and 500 genres. The next paginated response will return 300 brands, 0 age ranges, and 100 genres.
type SBTargetingGetRefinementsForCategoryResponseContent struct {
	// List of Age Ranges. Use /sb/targets/categories/{categoryRefinementId}/refinements to retrieve Age Ranges. Age Ranges are only available for categories related to children's toys and games.
	AgeRanges []SBTargetingAgeRange `json:"ageRanges,omitempty"`
	// List of Brands.
	Brands []SBTargetingBrand `json:"brands,omitempty"`
	// List of Genres. Use /sb/targets/categories/{categoryRefinementId}/refinements to retrieve Genre Node IDs. Genres are only available for categories related to books.
	Genres []SBTargetingGenre `json:"genres,omitempty"`
	// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSBTargetingGetRefinementsForCategoryResponseContent instantiates a new SBTargetingGetRefinementsForCategoryResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingGetRefinementsForCategoryResponseContent() *SBTargetingGetRefinementsForCategoryResponseContent {
	this := SBTargetingGetRefinementsForCategoryResponseContent{}
	return &this
}

// NewSBTargetingGetRefinementsForCategoryResponseContentWithDefaults instantiates a new SBTargetingGetRefinementsForCategoryResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingGetRefinementsForCategoryResponseContentWithDefaults() *SBTargetingGetRefinementsForCategoryResponseContent {
	this := SBTargetingGetRefinementsForCategoryResponseContent{}
	return &this
}

// GetAgeRanges returns the AgeRanges field value if set, zero value otherwise.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) GetAgeRanges() []SBTargetingAgeRange {
	if o == nil || IsNil(o.AgeRanges) {
		var ret []SBTargetingAgeRange
		return ret
	}
	return o.AgeRanges
}

// GetAgeRangesOk returns a tuple with the AgeRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) GetAgeRangesOk() ([]SBTargetingAgeRange, bool) {
	if o == nil || IsNil(o.AgeRanges) {
		return nil, false
	}
	return o.AgeRanges, true
}

// HasAgeRanges returns a boolean if a field has been set.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) HasAgeRanges() bool {
	if o != nil && !IsNil(o.AgeRanges) {
		return true
	}

	return false
}

// SetAgeRanges gets a reference to the given []SBTargetingAgeRange and assigns it to the AgeRanges field.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) SetAgeRanges(v []SBTargetingAgeRange) {
	o.AgeRanges = v
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) GetBrands() []SBTargetingBrand {
	if o == nil || IsNil(o.Brands) {
		var ret []SBTargetingBrand
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) GetBrandsOk() ([]SBTargetingBrand, bool) {
	if o == nil || IsNil(o.Brands) {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) HasBrands() bool {
	if o != nil && !IsNil(o.Brands) {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []SBTargetingBrand and assigns it to the Brands field.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) SetBrands(v []SBTargetingBrand) {
	o.Brands = v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) GetGenres() []SBTargetingGenre {
	if o == nil || IsNil(o.Genres) {
		var ret []SBTargetingGenre
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) GetGenresOk() ([]SBTargetingGenre, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []SBTargetingGenre and assigns it to the Genres field.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) SetGenres(v []SBTargetingGenre) {
	o.Genres = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SBTargetingGetRefinementsForCategoryResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SBTargetingGetRefinementsForCategoryResponseContent) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableSBTargetingGetRefinementsForCategoryResponseContent struct {
	value *SBTargetingGetRefinementsForCategoryResponseContent
	isSet bool
}

func (v NullableSBTargetingGetRefinementsForCategoryResponseContent) Get() *SBTargetingGetRefinementsForCategoryResponseContent {
	return v.value
}

func (v *NullableSBTargetingGetRefinementsForCategoryResponseContent) Set(val *SBTargetingGetRefinementsForCategoryResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingGetRefinementsForCategoryResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingGetRefinementsForCategoryResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingGetRefinementsForCategoryResponseContent(val *SBTargetingGetRefinementsForCategoryResponseContent) *NullableSBTargetingGetRefinementsForCategoryResponseContent {
	return &NullableSBTargetingGetRefinementsForCategoryResponseContent{value: val, isSet: true}
}

func (v NullableSBTargetingGetRefinementsForCategoryResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingGetRefinementsForCategoryResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
