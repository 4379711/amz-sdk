package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeProperties{}

// CreativeProperties Creative properties
type CreativeProperties struct {
	BrandLogoCrop *AssetCrop `json:"brandLogoCrop,omitempty"`
	// The displayed brand name in the ad headline. Maximum length is 30 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements.
	BrandName *string `json:"brandName,omitempty"`
	// The identifier of image/video asset from the store's asset library
	CustomImageAssetId *string              `json:"customImageAssetId,omitempty"`
	LandingPage        *CreativeLandingPage `json:"landingPage,omitempty"`
	// An array of customImages associated with the creative.
	CustomImages []CustomImage `json:"customImages,omitempty"`
	// If set to true and the headline and/or video are not in the marketplace's default language, Amazon will attempt to translate them to the marketplace's default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information.
	ConsentToTranslate *bool      `json:"consentToTranslate,omitempty"`
	CustomImageCrop    *AssetCrop `json:"customImageCrop,omitempty"`
	CustomImageUrl     *string    `json:"customImageUrl,omitempty"`
	// The assetIds of the original videos submitted by the advertiser. If 'consentToTranslate' is set to true and translation is SUCCESSFUL then `originalVideoAssetIds` will return the original video assetId whereas `videoAssetIds` will return translated video assetId. In all other cases, 'originalVideoAssetIds' and `videoAssetIds` both will return original video assetId.
	OriginalVideoAssetIds []string `json:"originalVideoAssetIds,omitempty"`
	// ----------------------------------------------- List types ----------------------------------------------- A list of ASINs
	Asins        []string `json:"asins,omitempty"`
	BrandLogoUrl *string  `json:"brandLogoUrl,omitempty"`
	// An array of subpages
	Subpages []Subpage `json:"subpages,omitempty"`
	// The original headline submitted by the advertiser.
	OriginalHeadline *string `json:"originalHeadline,omitempty"`
	// The assetIds of the original videos submitted by the advertiser. If 'consentToTranslate' is set to true and translation is SUCCESSFUL then 'videoAssetIds' will return translated video assetId whereas `originalVideoAssetIds` will return the original video assetId. In all other cases, `videoAssetIds` will return original video assetId.
	VideoAssetIds []string `json:"videoAssetIds,omitempty"`
	// The identifier of image/video asset from the store's asset library
	BrandLogoAssetId *string `json:"brandLogoAssetId,omitempty"`
	// If 'consentToTranslate' is set to true and translation is SUCCESSFUL then `headline` will return the translated headline whereas `originalHeadline` will return the original headline. In all other cases, 'originalHeadline' and `headline` both will return the original headline.
	Headline *string `json:"headline,omitempty"`
}

// NewCreativeProperties instantiates a new CreativeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeProperties() *CreativeProperties {
	this := CreativeProperties{}
	return &this
}

// NewCreativePropertiesWithDefaults instantiates a new CreativeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativePropertiesWithDefaults() *CreativeProperties {
	this := CreativeProperties{}
	return &this
}

// GetBrandLogoCrop returns the BrandLogoCrop field value if set, zero value otherwise.
func (o *CreativeProperties) GetBrandLogoCrop() AssetCrop {
	if o == nil || IsNil(o.BrandLogoCrop) {
		var ret AssetCrop
		return ret
	}
	return *o.BrandLogoCrop
}

// GetBrandLogoCropOk returns a tuple with the BrandLogoCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetBrandLogoCropOk() (*AssetCrop, bool) {
	if o == nil || IsNil(o.BrandLogoCrop) {
		return nil, false
	}
	return o.BrandLogoCrop, true
}

// HasBrandLogoCrop returns a boolean if a field has been set.
func (o *CreativeProperties) HasBrandLogoCrop() bool {
	if o != nil && !IsNil(o.BrandLogoCrop) {
		return true
	}

	return false
}

// SetBrandLogoCrop gets a reference to the given AssetCrop and assigns it to the BrandLogoCrop field.
func (o *CreativeProperties) SetBrandLogoCrop(v AssetCrop) {
	o.BrandLogoCrop = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *CreativeProperties) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *CreativeProperties) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *CreativeProperties) SetBrandName(v string) {
	o.BrandName = &v
}

// GetCustomImageAssetId returns the CustomImageAssetId field value if set, zero value otherwise.
func (o *CreativeProperties) GetCustomImageAssetId() string {
	if o == nil || IsNil(o.CustomImageAssetId) {
		var ret string
		return ret
	}
	return *o.CustomImageAssetId
}

// GetCustomImageAssetIdOk returns a tuple with the CustomImageAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetCustomImageAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomImageAssetId) {
		return nil, false
	}
	return o.CustomImageAssetId, true
}

// HasCustomImageAssetId returns a boolean if a field has been set.
func (o *CreativeProperties) HasCustomImageAssetId() bool {
	if o != nil && !IsNil(o.CustomImageAssetId) {
		return true
	}

	return false
}

// SetCustomImageAssetId gets a reference to the given string and assigns it to the CustomImageAssetId field.
func (o *CreativeProperties) SetCustomImageAssetId(v string) {
	o.CustomImageAssetId = &v
}

// GetLandingPage returns the LandingPage field value if set, zero value otherwise.
func (o *CreativeProperties) GetLandingPage() CreativeLandingPage {
	if o == nil || IsNil(o.LandingPage) {
		var ret CreativeLandingPage
		return ret
	}
	return *o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetLandingPageOk() (*CreativeLandingPage, bool) {
	if o == nil || IsNil(o.LandingPage) {
		return nil, false
	}
	return o.LandingPage, true
}

// HasLandingPage returns a boolean if a field has been set.
func (o *CreativeProperties) HasLandingPage() bool {
	if o != nil && !IsNil(o.LandingPage) {
		return true
	}

	return false
}

// SetLandingPage gets a reference to the given CreativeLandingPage and assigns it to the LandingPage field.
func (o *CreativeProperties) SetLandingPage(v CreativeLandingPage) {
	o.LandingPage = &v
}

// GetCustomImages returns the CustomImages field value if set, zero value otherwise.
func (o *CreativeProperties) GetCustomImages() []CustomImage {
	if o == nil || IsNil(o.CustomImages) {
		var ret []CustomImage
		return ret
	}
	return o.CustomImages
}

// GetCustomImagesOk returns a tuple with the CustomImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetCustomImagesOk() ([]CustomImage, bool) {
	if o == nil || IsNil(o.CustomImages) {
		return nil, false
	}
	return o.CustomImages, true
}

// HasCustomImages returns a boolean if a field has been set.
func (o *CreativeProperties) HasCustomImages() bool {
	if o != nil && !IsNil(o.CustomImages) {
		return true
	}

	return false
}

// SetCustomImages gets a reference to the given []CustomImage and assigns it to the CustomImages field.
func (o *CreativeProperties) SetCustomImages(v []CustomImage) {
	o.CustomImages = v
}

// GetConsentToTranslate returns the ConsentToTranslate field value if set, zero value otherwise.
func (o *CreativeProperties) GetConsentToTranslate() bool {
	if o == nil || IsNil(o.ConsentToTranslate) {
		var ret bool
		return ret
	}
	return *o.ConsentToTranslate
}

// GetConsentToTranslateOk returns a tuple with the ConsentToTranslate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetConsentToTranslateOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsentToTranslate) {
		return nil, false
	}
	return o.ConsentToTranslate, true
}

// HasConsentToTranslate returns a boolean if a field has been set.
func (o *CreativeProperties) HasConsentToTranslate() bool {
	if o != nil && !IsNil(o.ConsentToTranslate) {
		return true
	}

	return false
}

// SetConsentToTranslate gets a reference to the given bool and assigns it to the ConsentToTranslate field.
func (o *CreativeProperties) SetConsentToTranslate(v bool) {
	o.ConsentToTranslate = &v
}

// GetCustomImageCrop returns the CustomImageCrop field value if set, zero value otherwise.
func (o *CreativeProperties) GetCustomImageCrop() AssetCrop {
	if o == nil || IsNil(o.CustomImageCrop) {
		var ret AssetCrop
		return ret
	}
	return *o.CustomImageCrop
}

// GetCustomImageCropOk returns a tuple with the CustomImageCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetCustomImageCropOk() (*AssetCrop, bool) {
	if o == nil || IsNil(o.CustomImageCrop) {
		return nil, false
	}
	return o.CustomImageCrop, true
}

// HasCustomImageCrop returns a boolean if a field has been set.
func (o *CreativeProperties) HasCustomImageCrop() bool {
	if o != nil && !IsNil(o.CustomImageCrop) {
		return true
	}

	return false
}

// SetCustomImageCrop gets a reference to the given AssetCrop and assigns it to the CustomImageCrop field.
func (o *CreativeProperties) SetCustomImageCrop(v AssetCrop) {
	o.CustomImageCrop = &v
}

// GetCustomImageUrl returns the CustomImageUrl field value if set, zero value otherwise.
func (o *CreativeProperties) GetCustomImageUrl() string {
	if o == nil || IsNil(o.CustomImageUrl) {
		var ret string
		return ret
	}
	return *o.CustomImageUrl
}

// GetCustomImageUrlOk returns a tuple with the CustomImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetCustomImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomImageUrl) {
		return nil, false
	}
	return o.CustomImageUrl, true
}

// HasCustomImageUrl returns a boolean if a field has been set.
func (o *CreativeProperties) HasCustomImageUrl() bool {
	if o != nil && !IsNil(o.CustomImageUrl) {
		return true
	}

	return false
}

// SetCustomImageUrl gets a reference to the given string and assigns it to the CustomImageUrl field.
func (o *CreativeProperties) SetCustomImageUrl(v string) {
	o.CustomImageUrl = &v
}

// GetOriginalVideoAssetIds returns the OriginalVideoAssetIds field value if set, zero value otherwise.
func (o *CreativeProperties) GetOriginalVideoAssetIds() []string {
	if o == nil || IsNil(o.OriginalVideoAssetIds) {
		var ret []string
		return ret
	}
	return o.OriginalVideoAssetIds
}

// GetOriginalVideoAssetIdsOk returns a tuple with the OriginalVideoAssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetOriginalVideoAssetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OriginalVideoAssetIds) {
		return nil, false
	}
	return o.OriginalVideoAssetIds, true
}

// HasOriginalVideoAssetIds returns a boolean if a field has been set.
func (o *CreativeProperties) HasOriginalVideoAssetIds() bool {
	if o != nil && !IsNil(o.OriginalVideoAssetIds) {
		return true
	}

	return false
}

// SetOriginalVideoAssetIds gets a reference to the given []string and assigns it to the OriginalVideoAssetIds field.
func (o *CreativeProperties) SetOriginalVideoAssetIds(v []string) {
	o.OriginalVideoAssetIds = v
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *CreativeProperties) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *CreativeProperties) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *CreativeProperties) SetAsins(v []string) {
	o.Asins = v
}

// GetBrandLogoUrl returns the BrandLogoUrl field value if set, zero value otherwise.
func (o *CreativeProperties) GetBrandLogoUrl() string {
	if o == nil || IsNil(o.BrandLogoUrl) {
		var ret string
		return ret
	}
	return *o.BrandLogoUrl
}

// GetBrandLogoUrlOk returns a tuple with the BrandLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetBrandLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogoUrl) {
		return nil, false
	}
	return o.BrandLogoUrl, true
}

// HasBrandLogoUrl returns a boolean if a field has been set.
func (o *CreativeProperties) HasBrandLogoUrl() bool {
	if o != nil && !IsNil(o.BrandLogoUrl) {
		return true
	}

	return false
}

// SetBrandLogoUrl gets a reference to the given string and assigns it to the BrandLogoUrl field.
func (o *CreativeProperties) SetBrandLogoUrl(v string) {
	o.BrandLogoUrl = &v
}

// GetSubpages returns the Subpages field value if set, zero value otherwise.
func (o *CreativeProperties) GetSubpages() []Subpage {
	if o == nil || IsNil(o.Subpages) {
		var ret []Subpage
		return ret
	}
	return o.Subpages
}

// GetSubpagesOk returns a tuple with the Subpages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetSubpagesOk() ([]Subpage, bool) {
	if o == nil || IsNil(o.Subpages) {
		return nil, false
	}
	return o.Subpages, true
}

// HasSubpages returns a boolean if a field has been set.
func (o *CreativeProperties) HasSubpages() bool {
	if o != nil && !IsNil(o.Subpages) {
		return true
	}

	return false
}

// SetSubpages gets a reference to the given []Subpage and assigns it to the Subpages field.
func (o *CreativeProperties) SetSubpages(v []Subpage) {
	o.Subpages = v
}

// GetOriginalHeadline returns the OriginalHeadline field value if set, zero value otherwise.
func (o *CreativeProperties) GetOriginalHeadline() string {
	if o == nil || IsNil(o.OriginalHeadline) {
		var ret string
		return ret
	}
	return *o.OriginalHeadline
}

// GetOriginalHeadlineOk returns a tuple with the OriginalHeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetOriginalHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalHeadline) {
		return nil, false
	}
	return o.OriginalHeadline, true
}

// HasOriginalHeadline returns a boolean if a field has been set.
func (o *CreativeProperties) HasOriginalHeadline() bool {
	if o != nil && !IsNil(o.OriginalHeadline) {
		return true
	}

	return false
}

// SetOriginalHeadline gets a reference to the given string and assigns it to the OriginalHeadline field.
func (o *CreativeProperties) SetOriginalHeadline(v string) {
	o.OriginalHeadline = &v
}

// GetVideoAssetIds returns the VideoAssetIds field value if set, zero value otherwise.
func (o *CreativeProperties) GetVideoAssetIds() []string {
	if o == nil || IsNil(o.VideoAssetIds) {
		var ret []string
		return ret
	}
	return o.VideoAssetIds
}

// GetVideoAssetIdsOk returns a tuple with the VideoAssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetVideoAssetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VideoAssetIds) {
		return nil, false
	}
	return o.VideoAssetIds, true
}

// HasVideoAssetIds returns a boolean if a field has been set.
func (o *CreativeProperties) HasVideoAssetIds() bool {
	if o != nil && !IsNil(o.VideoAssetIds) {
		return true
	}

	return false
}

// SetVideoAssetIds gets a reference to the given []string and assigns it to the VideoAssetIds field.
func (o *CreativeProperties) SetVideoAssetIds(v []string) {
	o.VideoAssetIds = v
}

// GetBrandLogoAssetId returns the BrandLogoAssetId field value if set, zero value otherwise.
func (o *CreativeProperties) GetBrandLogoAssetId() string {
	if o == nil || IsNil(o.BrandLogoAssetId) {
		var ret string
		return ret
	}
	return *o.BrandLogoAssetId
}

// GetBrandLogoAssetIdOk returns a tuple with the BrandLogoAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetBrandLogoAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogoAssetId) {
		return nil, false
	}
	return o.BrandLogoAssetId, true
}

// HasBrandLogoAssetId returns a boolean if a field has been set.
func (o *CreativeProperties) HasBrandLogoAssetId() bool {
	if o != nil && !IsNil(o.BrandLogoAssetId) {
		return true
	}

	return false
}

// SetBrandLogoAssetId gets a reference to the given string and assigns it to the BrandLogoAssetId field.
func (o *CreativeProperties) SetBrandLogoAssetId(v string) {
	o.BrandLogoAssetId = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *CreativeProperties) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeProperties) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *CreativeProperties) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *CreativeProperties) SetHeadline(v string) {
	o.Headline = &v
}

func (o CreativeProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandLogoCrop) {
		toSerialize["brandLogoCrop"] = o.BrandLogoCrop
	}
	if !IsNil(o.BrandName) {
		toSerialize["brandName"] = o.BrandName
	}
	if !IsNil(o.CustomImageAssetId) {
		toSerialize["customImageAssetId"] = o.CustomImageAssetId
	}
	if !IsNil(o.LandingPage) {
		toSerialize["landingPage"] = o.LandingPage
	}
	if !IsNil(o.CustomImages) {
		toSerialize["customImages"] = o.CustomImages
	}
	if !IsNil(o.ConsentToTranslate) {
		toSerialize["consentToTranslate"] = o.ConsentToTranslate
	}
	if !IsNil(o.CustomImageCrop) {
		toSerialize["customImageCrop"] = o.CustomImageCrop
	}
	if !IsNil(o.CustomImageUrl) {
		toSerialize["customImageUrl"] = o.CustomImageUrl
	}
	if !IsNil(o.OriginalVideoAssetIds) {
		toSerialize["originalVideoAssetIds"] = o.OriginalVideoAssetIds
	}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.BrandLogoUrl) {
		toSerialize["brandLogoUrl"] = o.BrandLogoUrl
	}
	if !IsNil(o.Subpages) {
		toSerialize["subpages"] = o.Subpages
	}
	if !IsNil(o.OriginalHeadline) {
		toSerialize["originalHeadline"] = o.OriginalHeadline
	}
	if !IsNil(o.VideoAssetIds) {
		toSerialize["videoAssetIds"] = o.VideoAssetIds
	}
	if !IsNil(o.BrandLogoAssetId) {
		toSerialize["brandLogoAssetId"] = o.BrandLogoAssetId
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	return toSerialize, nil
}

type NullableCreativeProperties struct {
	value *CreativeProperties
	isSet bool
}

func (v NullableCreativeProperties) Get() *CreativeProperties {
	return v.value
}

func (v *NullableCreativeProperties) Set(val *CreativeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeProperties(val *CreativeProperties) *NullableCreativeProperties {
	return &NullableCreativeProperties{value: val, isSet: true}
}

func (v NullableCreativeProperties) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
