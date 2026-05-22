package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the StoreSpotlightCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreSpotlightCreative{}

// StoreSpotlightCreative struct for StoreSpotlightCreative
type StoreSpotlightCreative struct {
	BrandLogoCrop *AssetCrop `json:"brandLogoCrop,omitempty"`
	// The displayed brand name in the ad headline. Maximum length is 30 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements.
	BrandName string `json:"brandName"`
	// An array of subpages
	Subpages    []Subpage              `json:"subpages"`
	LandingPage *CreativeLandingPageV2 `json:"landingPage,omitempty"`
	// If set to true and the headline and/or video are not in the marketplace's default language, Amazon will attempt to translate them to the marketplace's default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information.
	ConsentToTranslate *bool `json:"consentToTranslate,omitempty"`
	// The identifier of the [brand logo](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#brandlogo) image from the brand store's asset library. Note that for campaigns created in the Amazon Advertising console prior to release of the brand store's assets library, responses will not include a value for this field.
	BrandLogoAssetId string `json:"brandLogoAssetId"`
	// The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements.
	Headline string `json:"headline"`
}

type _StoreSpotlightCreative StoreSpotlightCreative

// NewStoreSpotlightCreative instantiates a new StoreSpotlightCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreSpotlightCreative(brandName string, subpages []Subpage, brandLogoAssetId string, headline string) *StoreSpotlightCreative {
	this := StoreSpotlightCreative{}
	this.BrandName = brandName
	this.Subpages = subpages
	this.BrandLogoAssetId = brandLogoAssetId
	this.Headline = headline
	return &this
}

// NewStoreSpotlightCreativeWithDefaults instantiates a new StoreSpotlightCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreSpotlightCreativeWithDefaults() *StoreSpotlightCreative {
	this := StoreSpotlightCreative{}
	return &this
}

// GetBrandLogoCrop returns the BrandLogoCrop field value if set, zero value otherwise.
func (o *StoreSpotlightCreative) GetBrandLogoCrop() AssetCrop {
	if o == nil || IsNil(o.BrandLogoCrop) {
		var ret AssetCrop
		return ret
	}
	return *o.BrandLogoCrop
}

// GetBrandLogoCropOk returns a tuple with the BrandLogoCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreSpotlightCreative) GetBrandLogoCropOk() (*AssetCrop, bool) {
	if o == nil || IsNil(o.BrandLogoCrop) {
		return nil, false
	}
	return o.BrandLogoCrop, true
}

// HasBrandLogoCrop returns a boolean if a field has been set.
func (o *StoreSpotlightCreative) HasBrandLogoCrop() bool {
	if o != nil && !IsNil(o.BrandLogoCrop) {
		return true
	}

	return false
}

// SetBrandLogoCrop gets a reference to the given AssetCrop and assigns it to the BrandLogoCrop field.
func (o *StoreSpotlightCreative) SetBrandLogoCrop(v AssetCrop) {
	o.BrandLogoCrop = &v
}

// GetBrandName returns the BrandName field value
func (o *StoreSpotlightCreative) GetBrandName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value
// and a boolean to check if the value has been set.
func (o *StoreSpotlightCreative) GetBrandNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandName, true
}

// SetBrandName sets field value
func (o *StoreSpotlightCreative) SetBrandName(v string) {
	o.BrandName = v
}

// GetSubpages returns the Subpages field value
func (o *StoreSpotlightCreative) GetSubpages() []Subpage {
	if o == nil {
		var ret []Subpage
		return ret
	}

	return o.Subpages
}

// GetSubpagesOk returns a tuple with the Subpages field value
// and a boolean to check if the value has been set.
func (o *StoreSpotlightCreative) GetSubpagesOk() ([]Subpage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subpages, true
}

// SetSubpages sets field value
func (o *StoreSpotlightCreative) SetSubpages(v []Subpage) {
	o.Subpages = v
}

// GetLandingPage returns the LandingPage field value if set, zero value otherwise.
func (o *StoreSpotlightCreative) GetLandingPage() CreativeLandingPageV2 {
	if o == nil || IsNil(o.LandingPage) {
		var ret CreativeLandingPageV2
		return ret
	}
	return *o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreSpotlightCreative) GetLandingPageOk() (*CreativeLandingPageV2, bool) {
	if o == nil || IsNil(o.LandingPage) {
		return nil, false
	}
	return o.LandingPage, true
}

// HasLandingPage returns a boolean if a field has been set.
func (o *StoreSpotlightCreative) HasLandingPage() bool {
	if o != nil && !IsNil(o.LandingPage) {
		return true
	}

	return false
}

// SetLandingPage gets a reference to the given CreativeLandingPageV2 and assigns it to the LandingPage field.
func (o *StoreSpotlightCreative) SetLandingPage(v CreativeLandingPageV2) {
	o.LandingPage = &v
}

// GetConsentToTranslate returns the ConsentToTranslate field value if set, zero value otherwise.
func (o *StoreSpotlightCreative) GetConsentToTranslate() bool {
	if o == nil || IsNil(o.ConsentToTranslate) {
		var ret bool
		return ret
	}
	return *o.ConsentToTranslate
}

// GetConsentToTranslateOk returns a tuple with the ConsentToTranslate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreSpotlightCreative) GetConsentToTranslateOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsentToTranslate) {
		return nil, false
	}
	return o.ConsentToTranslate, true
}

// HasConsentToTranslate returns a boolean if a field has been set.
func (o *StoreSpotlightCreative) HasConsentToTranslate() bool {
	if o != nil && !IsNil(o.ConsentToTranslate) {
		return true
	}

	return false
}

// SetConsentToTranslate gets a reference to the given bool and assigns it to the ConsentToTranslate field.
func (o *StoreSpotlightCreative) SetConsentToTranslate(v bool) {
	o.ConsentToTranslate = &v
}

// GetBrandLogoAssetId returns the BrandLogoAssetId field value
func (o *StoreSpotlightCreative) GetBrandLogoAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandLogoAssetId
}

// GetBrandLogoAssetIdOk returns a tuple with the BrandLogoAssetId field value
// and a boolean to check if the value has been set.
func (o *StoreSpotlightCreative) GetBrandLogoAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandLogoAssetId, true
}

// SetBrandLogoAssetId sets field value
func (o *StoreSpotlightCreative) SetBrandLogoAssetId(v string) {
	o.BrandLogoAssetId = v
}

// GetHeadline returns the Headline field value
func (o *StoreSpotlightCreative) GetHeadline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value
// and a boolean to check if the value has been set.
func (o *StoreSpotlightCreative) GetHeadlineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headline, true
}

// SetHeadline sets field value
func (o *StoreSpotlightCreative) SetHeadline(v string) {
	o.Headline = v
}

func (o StoreSpotlightCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandLogoCrop) {
		toSerialize["brandLogoCrop"] = o.BrandLogoCrop
	}
	toSerialize["brandName"] = o.BrandName
	toSerialize["subpages"] = o.Subpages
	if !IsNil(o.LandingPage) {
		toSerialize["landingPage"] = o.LandingPage
	}
	if !IsNil(o.ConsentToTranslate) {
		toSerialize["consentToTranslate"] = o.ConsentToTranslate
	}
	toSerialize["brandLogoAssetId"] = o.BrandLogoAssetId
	toSerialize["headline"] = o.Headline
	return toSerialize, nil
}

type NullableStoreSpotlightCreative struct {
	value *StoreSpotlightCreative
	isSet bool
}

func (v NullableStoreSpotlightCreative) Get() *StoreSpotlightCreative {
	return v.value
}

func (v *NullableStoreSpotlightCreative) Set(val *StoreSpotlightCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreSpotlightCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreSpotlightCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreSpotlightCreative(val *StoreSpotlightCreative) *NullableStoreSpotlightCreative {
	return &NullableStoreSpotlightCreative{value: val, isSet: true}
}

func (v NullableStoreSpotlightCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStoreSpotlightCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
