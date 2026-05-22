package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeRecommendationProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeRecommendationProperties{}

// CreativeRecommendationProperties Nested Creative Properties Structure for fetching Creative Recommendations.
type CreativeRecommendationProperties struct {
	// ----------------------------------------------- List types ----------------------------------------------- A list of ASINs
	Asins []string `json:"asins,omitempty"`
	// The displayed brand name in the ad headline. Maximum length is 30 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements.
	BrandName *string `json:"brandName,omitempty"`
	// An array of subpages
	Subpages    []Subpage              `json:"subpages,omitempty"`
	LandingPage *CreativeLandingPageV2 `json:"landingPage,omitempty"`
	// An array of customImages associated with the creative.
	CustomImages []CustomImage `json:"customImages,omitempty"`
	// An array of videoAssetIds associated with the creative. Advertisers can get video assetIds from Asset Library /assets/search API.
	VideoAssetIds []string `json:"videoAssetIds,omitempty"`
	// a Unique Id identifying the creative Recommendation
	RecommendedCreativeId *string    `json:"recommendedCreativeId,omitempty"`
	BrandLogo             *BrandLogo `json:"brandLogo,omitempty"`
	// The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements.
	Headline *string `json:"headline,omitempty"`
}

// NewCreativeRecommendationProperties instantiates a new CreativeRecommendationProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeRecommendationProperties() *CreativeRecommendationProperties {
	this := CreativeRecommendationProperties{}
	return &this
}

// NewCreativeRecommendationPropertiesWithDefaults instantiates a new CreativeRecommendationProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeRecommendationPropertiesWithDefaults() *CreativeRecommendationProperties {
	this := CreativeRecommendationProperties{}
	return &this
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *CreativeRecommendationProperties) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationProperties) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *CreativeRecommendationProperties) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *CreativeRecommendationProperties) SetAsins(v []string) {
	o.Asins = v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *CreativeRecommendationProperties) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationProperties) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *CreativeRecommendationProperties) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *CreativeRecommendationProperties) SetBrandName(v string) {
	o.BrandName = &v
}

// GetSubpages returns the Subpages field value if set, zero value otherwise.
func (o *CreativeRecommendationProperties) GetSubpages() []Subpage {
	if o == nil || IsNil(o.Subpages) {
		var ret []Subpage
		return ret
	}
	return o.Subpages
}

// GetSubpagesOk returns a tuple with the Subpages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationProperties) GetSubpagesOk() ([]Subpage, bool) {
	if o == nil || IsNil(o.Subpages) {
		return nil, false
	}
	return o.Subpages, true
}

// HasSubpages returns a boolean if a field has been set.
func (o *CreativeRecommendationProperties) HasSubpages() bool {
	if o != nil && !IsNil(o.Subpages) {
		return true
	}

	return false
}

// SetSubpages gets a reference to the given []Subpage and assigns it to the Subpages field.
func (o *CreativeRecommendationProperties) SetSubpages(v []Subpage) {
	o.Subpages = v
}

// GetLandingPage returns the LandingPage field value if set, zero value otherwise.
func (o *CreativeRecommendationProperties) GetLandingPage() CreativeLandingPageV2 {
	if o == nil || IsNil(o.LandingPage) {
		var ret CreativeLandingPageV2
		return ret
	}
	return *o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationProperties) GetLandingPageOk() (*CreativeLandingPageV2, bool) {
	if o == nil || IsNil(o.LandingPage) {
		return nil, false
	}
	return o.LandingPage, true
}

// HasLandingPage returns a boolean if a field has been set.
func (o *CreativeRecommendationProperties) HasLandingPage() bool {
	if o != nil && !IsNil(o.LandingPage) {
		return true
	}

	return false
}

// SetLandingPage gets a reference to the given CreativeLandingPageV2 and assigns it to the LandingPage field.
func (o *CreativeRecommendationProperties) SetLandingPage(v CreativeLandingPageV2) {
	o.LandingPage = &v
}

// GetCustomImages returns the CustomImages field value if set, zero value otherwise.
func (o *CreativeRecommendationProperties) GetCustomImages() []CustomImage {
	if o == nil || IsNil(o.CustomImages) {
		var ret []CustomImage
		return ret
	}
	return o.CustomImages
}

// GetCustomImagesOk returns a tuple with the CustomImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationProperties) GetCustomImagesOk() ([]CustomImage, bool) {
	if o == nil || IsNil(o.CustomImages) {
		return nil, false
	}
	return o.CustomImages, true
}

// HasCustomImages returns a boolean if a field has been set.
func (o *CreativeRecommendationProperties) HasCustomImages() bool {
	if o != nil && !IsNil(o.CustomImages) {
		return true
	}

	return false
}

// SetCustomImages gets a reference to the given []CustomImage and assigns it to the CustomImages field.
func (o *CreativeRecommendationProperties) SetCustomImages(v []CustomImage) {
	o.CustomImages = v
}

// GetVideoAssetIds returns the VideoAssetIds field value if set, zero value otherwise.
func (o *CreativeRecommendationProperties) GetVideoAssetIds() []string {
	if o == nil || IsNil(o.VideoAssetIds) {
		var ret []string
		return ret
	}
	return o.VideoAssetIds
}

// GetVideoAssetIdsOk returns a tuple with the VideoAssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationProperties) GetVideoAssetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VideoAssetIds) {
		return nil, false
	}
	return o.VideoAssetIds, true
}

// HasVideoAssetIds returns a boolean if a field has been set.
func (o *CreativeRecommendationProperties) HasVideoAssetIds() bool {
	if o != nil && !IsNil(o.VideoAssetIds) {
		return true
	}

	return false
}

// SetVideoAssetIds gets a reference to the given []string and assigns it to the VideoAssetIds field.
func (o *CreativeRecommendationProperties) SetVideoAssetIds(v []string) {
	o.VideoAssetIds = v
}

// GetRecommendedCreativeId returns the RecommendedCreativeId field value if set, zero value otherwise.
func (o *CreativeRecommendationProperties) GetRecommendedCreativeId() string {
	if o == nil || IsNil(o.RecommendedCreativeId) {
		var ret string
		return ret
	}
	return *o.RecommendedCreativeId
}

// GetRecommendedCreativeIdOk returns a tuple with the RecommendedCreativeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationProperties) GetRecommendedCreativeIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendedCreativeId) {
		return nil, false
	}
	return o.RecommendedCreativeId, true
}

// HasRecommendedCreativeId returns a boolean if a field has been set.
func (o *CreativeRecommendationProperties) HasRecommendedCreativeId() bool {
	if o != nil && !IsNil(o.RecommendedCreativeId) {
		return true
	}

	return false
}

// SetRecommendedCreativeId gets a reference to the given string and assigns it to the RecommendedCreativeId field.
func (o *CreativeRecommendationProperties) SetRecommendedCreativeId(v string) {
	o.RecommendedCreativeId = &v
}

// GetBrandLogo returns the BrandLogo field value if set, zero value otherwise.
func (o *CreativeRecommendationProperties) GetBrandLogo() BrandLogo {
	if o == nil || IsNil(o.BrandLogo) {
		var ret BrandLogo
		return ret
	}
	return *o.BrandLogo
}

// GetBrandLogoOk returns a tuple with the BrandLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationProperties) GetBrandLogoOk() (*BrandLogo, bool) {
	if o == nil || IsNil(o.BrandLogo) {
		return nil, false
	}
	return o.BrandLogo, true
}

// HasBrandLogo returns a boolean if a field has been set.
func (o *CreativeRecommendationProperties) HasBrandLogo() bool {
	if o != nil && !IsNil(o.BrandLogo) {
		return true
	}

	return false
}

// SetBrandLogo gets a reference to the given BrandLogo and assigns it to the BrandLogo field.
func (o *CreativeRecommendationProperties) SetBrandLogo(v BrandLogo) {
	o.BrandLogo = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *CreativeRecommendationProperties) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationProperties) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *CreativeRecommendationProperties) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *CreativeRecommendationProperties) SetHeadline(v string) {
	o.Headline = &v
}

func (o CreativeRecommendationProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.BrandName) {
		toSerialize["brandName"] = o.BrandName
	}
	if !IsNil(o.Subpages) {
		toSerialize["subpages"] = o.Subpages
	}
	if !IsNil(o.LandingPage) {
		toSerialize["landingPage"] = o.LandingPage
	}
	if !IsNil(o.CustomImages) {
		toSerialize["customImages"] = o.CustomImages
	}
	if !IsNil(o.VideoAssetIds) {
		toSerialize["videoAssetIds"] = o.VideoAssetIds
	}
	if !IsNil(o.RecommendedCreativeId) {
		toSerialize["recommendedCreativeId"] = o.RecommendedCreativeId
	}
	if !IsNil(o.BrandLogo) {
		toSerialize["brandLogo"] = o.BrandLogo
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	return toSerialize, nil
}

type NullableCreativeRecommendationProperties struct {
	value *CreativeRecommendationProperties
	isSet bool
}

func (v NullableCreativeRecommendationProperties) Get() *CreativeRecommendationProperties {
	return v.value
}

func (v *NullableCreativeRecommendationProperties) Set(val *CreativeRecommendationProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeRecommendationProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeRecommendationProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeRecommendationProperties(val *CreativeRecommendationProperties) *NullableCreativeRecommendationProperties {
	return &NullableCreativeRecommendationProperties{value: val, isSet: true}
}

func (v NullableCreativeRecommendationProperties) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeRecommendationProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
