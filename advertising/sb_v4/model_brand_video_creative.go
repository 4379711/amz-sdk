package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandVideoCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandVideoCreative{}

// BrandVideoCreative struct for BrandVideoCreative
type BrandVideoCreative struct {
	// An array of ASINs associated with the creative.
	Asins         []string   `json:"asins"`
	BrandLogoCrop *AssetCrop `json:"brandLogoCrop,omitempty"`
	// The displayed brand name in the ad headline. Maximum length is 30 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements.
	BrandName   string                 `json:"brandName"`
	LandingPage *CreativeLandingPageV2 `json:"landingPage,omitempty"`
	// If set to true and the headline and/or video are not in the marketplace's default language, Amazon will attempt to translate them to the marketplace's default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information.
	ConsentToTranslate *bool `json:"consentToTranslate,omitempty"`
	// The assetIds of the original videos submitted by the advertiser. If 'consentToTranslate' is set to true and translation is SUCCESSFUL then 'videoAssetIds' will return translated video assetId whereas `originalVideoAssetIds` will return the original video assetId. In all other cases, `videoAssetIds` will return original video assetId.
	VideoAssetIds []string `json:"videoAssetIds"`
	// The identifier of the [brand logo](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#brandlogo) image from the brand store's asset library. Note that for campaigns created in the Amazon Advertising console prior to release of the brand store's assets library, responses will not include a value for this field.
	BrandLogoAssetId string `json:"brandLogoAssetId"`
	// The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements.
	Headline string `json:"headline"`
}

type _BrandVideoCreative BrandVideoCreative

// NewBrandVideoCreative instantiates a new BrandVideoCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandVideoCreative(asins []string, brandName string, videoAssetIds []string, brandLogoAssetId string, headline string) *BrandVideoCreative {
	this := BrandVideoCreative{}
	this.Asins = asins
	this.BrandName = brandName
	this.VideoAssetIds = videoAssetIds
	this.BrandLogoAssetId = brandLogoAssetId
	this.Headline = headline
	return &this
}

// NewBrandVideoCreativeWithDefaults instantiates a new BrandVideoCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandVideoCreativeWithDefaults() *BrandVideoCreative {
	this := BrandVideoCreative{}
	return &this
}

// GetAsins returns the Asins field value
func (o *BrandVideoCreative) GetAsins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value
// and a boolean to check if the value has been set.
func (o *BrandVideoCreative) GetAsinsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asins, true
}

// SetAsins sets field value
func (o *BrandVideoCreative) SetAsins(v []string) {
	o.Asins = v
}

// GetBrandLogoCrop returns the BrandLogoCrop field value if set, zero value otherwise.
func (o *BrandVideoCreative) GetBrandLogoCrop() AssetCrop {
	if o == nil || IsNil(o.BrandLogoCrop) {
		var ret AssetCrop
		return ret
	}
	return *o.BrandLogoCrop
}

// GetBrandLogoCropOk returns a tuple with the BrandLogoCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandVideoCreative) GetBrandLogoCropOk() (*AssetCrop, bool) {
	if o == nil || IsNil(o.BrandLogoCrop) {
		return nil, false
	}
	return o.BrandLogoCrop, true
}

// HasBrandLogoCrop returns a boolean if a field has been set.
func (o *BrandVideoCreative) HasBrandLogoCrop() bool {
	if o != nil && !IsNil(o.BrandLogoCrop) {
		return true
	}

	return false
}

// SetBrandLogoCrop gets a reference to the given AssetCrop and assigns it to the BrandLogoCrop field.
func (o *BrandVideoCreative) SetBrandLogoCrop(v AssetCrop) {
	o.BrandLogoCrop = &v
}

// GetBrandName returns the BrandName field value
func (o *BrandVideoCreative) GetBrandName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value
// and a boolean to check if the value has been set.
func (o *BrandVideoCreative) GetBrandNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandName, true
}

// SetBrandName sets field value
func (o *BrandVideoCreative) SetBrandName(v string) {
	o.BrandName = v
}

// GetLandingPage returns the LandingPage field value if set, zero value otherwise.
func (o *BrandVideoCreative) GetLandingPage() CreativeLandingPageV2 {
	if o == nil || IsNil(o.LandingPage) {
		var ret CreativeLandingPageV2
		return ret
	}
	return *o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandVideoCreative) GetLandingPageOk() (*CreativeLandingPageV2, bool) {
	if o == nil || IsNil(o.LandingPage) {
		return nil, false
	}
	return o.LandingPage, true
}

// HasLandingPage returns a boolean if a field has been set.
func (o *BrandVideoCreative) HasLandingPage() bool {
	if o != nil && !IsNil(o.LandingPage) {
		return true
	}

	return false
}

// SetLandingPage gets a reference to the given CreativeLandingPageV2 and assigns it to the LandingPage field.
func (o *BrandVideoCreative) SetLandingPage(v CreativeLandingPageV2) {
	o.LandingPage = &v
}

// GetConsentToTranslate returns the ConsentToTranslate field value if set, zero value otherwise.
func (o *BrandVideoCreative) GetConsentToTranslate() bool {
	if o == nil || IsNil(o.ConsentToTranslate) {
		var ret bool
		return ret
	}
	return *o.ConsentToTranslate
}

// GetConsentToTranslateOk returns a tuple with the ConsentToTranslate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandVideoCreative) GetConsentToTranslateOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsentToTranslate) {
		return nil, false
	}
	return o.ConsentToTranslate, true
}

// HasConsentToTranslate returns a boolean if a field has been set.
func (o *BrandVideoCreative) HasConsentToTranslate() bool {
	if o != nil && !IsNil(o.ConsentToTranslate) {
		return true
	}

	return false
}

// SetConsentToTranslate gets a reference to the given bool and assigns it to the ConsentToTranslate field.
func (o *BrandVideoCreative) SetConsentToTranslate(v bool) {
	o.ConsentToTranslate = &v
}

// GetVideoAssetIds returns the VideoAssetIds field value
func (o *BrandVideoCreative) GetVideoAssetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.VideoAssetIds
}

// GetVideoAssetIdsOk returns a tuple with the VideoAssetIds field value
// and a boolean to check if the value has been set.
func (o *BrandVideoCreative) GetVideoAssetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoAssetIds, true
}

// SetVideoAssetIds sets field value
func (o *BrandVideoCreative) SetVideoAssetIds(v []string) {
	o.VideoAssetIds = v
}

// GetBrandLogoAssetId returns the BrandLogoAssetId field value
func (o *BrandVideoCreative) GetBrandLogoAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandLogoAssetId
}

// GetBrandLogoAssetIdOk returns a tuple with the BrandLogoAssetId field value
// and a boolean to check if the value has been set.
func (o *BrandVideoCreative) GetBrandLogoAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandLogoAssetId, true
}

// SetBrandLogoAssetId sets field value
func (o *BrandVideoCreative) SetBrandLogoAssetId(v string) {
	o.BrandLogoAssetId = v
}

// GetHeadline returns the Headline field value
func (o *BrandVideoCreative) GetHeadline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value
// and a boolean to check if the value has been set.
func (o *BrandVideoCreative) GetHeadlineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headline, true
}

// SetHeadline sets field value
func (o *BrandVideoCreative) SetHeadline(v string) {
	o.Headline = v
}

func (o BrandVideoCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asins"] = o.Asins
	if !IsNil(o.BrandLogoCrop) {
		toSerialize["brandLogoCrop"] = o.BrandLogoCrop
	}
	toSerialize["brandName"] = o.BrandName
	if !IsNil(o.LandingPage) {
		toSerialize["landingPage"] = o.LandingPage
	}
	if !IsNil(o.ConsentToTranslate) {
		toSerialize["consentToTranslate"] = o.ConsentToTranslate
	}
	toSerialize["videoAssetIds"] = o.VideoAssetIds
	toSerialize["brandLogoAssetId"] = o.BrandLogoAssetId
	toSerialize["headline"] = o.Headline
	return toSerialize, nil
}

type NullableBrandVideoCreative struct {
	value *BrandVideoCreative
	isSet bool
}

func (v NullableBrandVideoCreative) Get() *BrandVideoCreative {
	return v.value
}

func (v *NullableBrandVideoCreative) Set(val *BrandVideoCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandVideoCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandVideoCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandVideoCreative(val *BrandVideoCreative) *NullableBrandVideoCreative {
	return &NullableBrandVideoCreative{value: val, isSet: true}
}

func (v NullableBrandVideoCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandVideoCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
