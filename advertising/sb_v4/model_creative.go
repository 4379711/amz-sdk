package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the Creative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Creative{}

// Creative struct for Creative
type Creative struct {
	BrandLogoCrop      *BrandLogoCrop `json:"brandLogoCrop,omitempty"`
	BrandName          *string        `json:"brandName,omitempty"`
	CustomImageAssetId *string        `json:"customImageAssetId,omitempty"`
	// If set to true and the headline and/or video asset are not in the marketplace's default language, Amazon will attempt to translate them to the marketplace's default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information.
	ConsentToTranslate *bool `json:"consentToTranslate,omitempty"`
	// Requires minimum one custom image. You can add an optional collection of custom images that can be displayed on the ad as slideshow. Learn more about slideshow here https://advertising.amazon.com/resources/whats-new/slideshow-ads-creative-for-sponsored-brands/
	CustomImages    []CustomImage    `json:"customImages,omitempty"`
	CustomImageCrop *CustomImageCrop `json:"customImageCrop,omitempty"`
	CustomImageUrl  *string          `json:"customImageUrl,omitempty"`
	Type            *CreativeType    `json:"type,omitempty"`
	// The assetIds of the original videos submitted by the advertiser. If 'consentToTranslate' is set to true and translation is SUCCESSFUL then `originalVideoAssetIds` will return the original video assetId whereas `videoAssetIds` will return translated video assetId. In all other cases, 'originalVideoAssetIds' and `videoAssetIds` both will return original video assetId.
	OriginalVideoAssetIds []string  `json:"originalVideoAssetIds,omitempty"`
	Asins                 []string  `json:"asins,omitempty"`
	BrandLogoUrl          *string   `json:"brandLogoUrl,omitempty"`
	Subpages              []Subpage `json:"subpages,omitempty"`
	// The original headline submitted by the advertiser.
	OriginalHeadline *string `json:"originalHeadline,omitempty"`
	// In SB API V4, `videoMediaIds` is replaced by `videoAssetIds`. `videoAssetIds` will only allow Asset Library identifiers for ad creation, but responses can include mediaIds for v1 campaigns and API V3 operations. At a future state, existing mediaIds will be added to Asset library for use in SB campaigns.
	VideoAssetIds    []string `json:"videoAssetIds,omitempty"`
	BrandLogoAssetID *string  `json:"brandLogoAssetID,omitempty"`
	// The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters.
	Headline       *string         `json:"headline,omitempty"`
	CreativeStatus *CreativeStatus `json:"creativeStatus,omitempty"`
}

// NewCreative instantiates a new Creative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreative() *Creative {
	this := Creative{}
	return &this
}

// NewCreativeWithDefaults instantiates a new Creative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeWithDefaults() *Creative {
	this := Creative{}
	return &this
}

// GetBrandLogoCrop returns the BrandLogoCrop field value if set, zero value otherwise.
func (o *Creative) GetBrandLogoCrop() BrandLogoCrop {
	if o == nil || IsNil(o.BrandLogoCrop) {
		var ret BrandLogoCrop
		return ret
	}
	return *o.BrandLogoCrop
}

// GetBrandLogoCropOk returns a tuple with the BrandLogoCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetBrandLogoCropOk() (*BrandLogoCrop, bool) {
	if o == nil || IsNil(o.BrandLogoCrop) {
		return nil, false
	}
	return o.BrandLogoCrop, true
}

// HasBrandLogoCrop returns a boolean if a field has been set.
func (o *Creative) HasBrandLogoCrop() bool {
	if o != nil && !IsNil(o.BrandLogoCrop) {
		return true
	}

	return false
}

// SetBrandLogoCrop gets a reference to the given BrandLogoCrop and assigns it to the BrandLogoCrop field.
func (o *Creative) SetBrandLogoCrop(v BrandLogoCrop) {
	o.BrandLogoCrop = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *Creative) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *Creative) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *Creative) SetBrandName(v string) {
	o.BrandName = &v
}

// GetCustomImageAssetId returns the CustomImageAssetId field value if set, zero value otherwise.
func (o *Creative) GetCustomImageAssetId() string {
	if o == nil || IsNil(o.CustomImageAssetId) {
		var ret string
		return ret
	}
	return *o.CustomImageAssetId
}

// GetCustomImageAssetIdOk returns a tuple with the CustomImageAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetCustomImageAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomImageAssetId) {
		return nil, false
	}
	return o.CustomImageAssetId, true
}

// HasCustomImageAssetId returns a boolean if a field has been set.
func (o *Creative) HasCustomImageAssetId() bool {
	if o != nil && !IsNil(o.CustomImageAssetId) {
		return true
	}

	return false
}

// SetCustomImageAssetId gets a reference to the given string and assigns it to the CustomImageAssetId field.
func (o *Creative) SetCustomImageAssetId(v string) {
	o.CustomImageAssetId = &v
}

// GetConsentToTranslate returns the ConsentToTranslate field value if set, zero value otherwise.
func (o *Creative) GetConsentToTranslate() bool {
	if o == nil || IsNil(o.ConsentToTranslate) {
		var ret bool
		return ret
	}
	return *o.ConsentToTranslate
}

// GetConsentToTranslateOk returns a tuple with the ConsentToTranslate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetConsentToTranslateOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsentToTranslate) {
		return nil, false
	}
	return o.ConsentToTranslate, true
}

// HasConsentToTranslate returns a boolean if a field has been set.
func (o *Creative) HasConsentToTranslate() bool {
	if o != nil && !IsNil(o.ConsentToTranslate) {
		return true
	}

	return false
}

// SetConsentToTranslate gets a reference to the given bool and assigns it to the ConsentToTranslate field.
func (o *Creative) SetConsentToTranslate(v bool) {
	o.ConsentToTranslate = &v
}

// GetCustomImages returns the CustomImages field value if set, zero value otherwise.
func (o *Creative) GetCustomImages() []CustomImage {
	if o == nil || IsNil(o.CustomImages) {
		var ret []CustomImage
		return ret
	}
	return o.CustomImages
}

// GetCustomImagesOk returns a tuple with the CustomImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetCustomImagesOk() ([]CustomImage, bool) {
	if o == nil || IsNil(o.CustomImages) {
		return nil, false
	}
	return o.CustomImages, true
}

// HasCustomImages returns a boolean if a field has been set.
func (o *Creative) HasCustomImages() bool {
	if o != nil && !IsNil(o.CustomImages) {
		return true
	}

	return false
}

// SetCustomImages gets a reference to the given []CustomImage and assigns it to the CustomImages field.
func (o *Creative) SetCustomImages(v []CustomImage) {
	o.CustomImages = v
}

// GetCustomImageCrop returns the CustomImageCrop field value if set, zero value otherwise.
func (o *Creative) GetCustomImageCrop() CustomImageCrop {
	if o == nil || IsNil(o.CustomImageCrop) {
		var ret CustomImageCrop
		return ret
	}
	return *o.CustomImageCrop
}

// GetCustomImageCropOk returns a tuple with the CustomImageCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetCustomImageCropOk() (*CustomImageCrop, bool) {
	if o == nil || IsNil(o.CustomImageCrop) {
		return nil, false
	}
	return o.CustomImageCrop, true
}

// HasCustomImageCrop returns a boolean if a field has been set.
func (o *Creative) HasCustomImageCrop() bool {
	if o != nil && !IsNil(o.CustomImageCrop) {
		return true
	}

	return false
}

// SetCustomImageCrop gets a reference to the given CustomImageCrop and assigns it to the CustomImageCrop field.
func (o *Creative) SetCustomImageCrop(v CustomImageCrop) {
	o.CustomImageCrop = &v
}

// GetCustomImageUrl returns the CustomImageUrl field value if set, zero value otherwise.
func (o *Creative) GetCustomImageUrl() string {
	if o == nil || IsNil(o.CustomImageUrl) {
		var ret string
		return ret
	}
	return *o.CustomImageUrl
}

// GetCustomImageUrlOk returns a tuple with the CustomImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetCustomImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomImageUrl) {
		return nil, false
	}
	return o.CustomImageUrl, true
}

// HasCustomImageUrl returns a boolean if a field has been set.
func (o *Creative) HasCustomImageUrl() bool {
	if o != nil && !IsNil(o.CustomImageUrl) {
		return true
	}

	return false
}

// SetCustomImageUrl gets a reference to the given string and assigns it to the CustomImageUrl field.
func (o *Creative) SetCustomImageUrl(v string) {
	o.CustomImageUrl = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Creative) GetType() CreativeType {
	if o == nil || IsNil(o.Type) {
		var ret CreativeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetTypeOk() (*CreativeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Creative) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CreativeType and assigns it to the Type field.
func (o *Creative) SetType(v CreativeType) {
	o.Type = &v
}

// GetOriginalVideoAssetIds returns the OriginalVideoAssetIds field value if set, zero value otherwise.
func (o *Creative) GetOriginalVideoAssetIds() []string {
	if o == nil || IsNil(o.OriginalVideoAssetIds) {
		var ret []string
		return ret
	}
	return o.OriginalVideoAssetIds
}

// GetOriginalVideoAssetIdsOk returns a tuple with the OriginalVideoAssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetOriginalVideoAssetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OriginalVideoAssetIds) {
		return nil, false
	}
	return o.OriginalVideoAssetIds, true
}

// HasOriginalVideoAssetIds returns a boolean if a field has been set.
func (o *Creative) HasOriginalVideoAssetIds() bool {
	if o != nil && !IsNil(o.OriginalVideoAssetIds) {
		return true
	}

	return false
}

// SetOriginalVideoAssetIds gets a reference to the given []string and assigns it to the OriginalVideoAssetIds field.
func (o *Creative) SetOriginalVideoAssetIds(v []string) {
	o.OriginalVideoAssetIds = v
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *Creative) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *Creative) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *Creative) SetAsins(v []string) {
	o.Asins = v
}

// GetBrandLogoUrl returns the BrandLogoUrl field value if set, zero value otherwise.
func (o *Creative) GetBrandLogoUrl() string {
	if o == nil || IsNil(o.BrandLogoUrl) {
		var ret string
		return ret
	}
	return *o.BrandLogoUrl
}

// GetBrandLogoUrlOk returns a tuple with the BrandLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetBrandLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogoUrl) {
		return nil, false
	}
	return o.BrandLogoUrl, true
}

// HasBrandLogoUrl returns a boolean if a field has been set.
func (o *Creative) HasBrandLogoUrl() bool {
	if o != nil && !IsNil(o.BrandLogoUrl) {
		return true
	}

	return false
}

// SetBrandLogoUrl gets a reference to the given string and assigns it to the BrandLogoUrl field.
func (o *Creative) SetBrandLogoUrl(v string) {
	o.BrandLogoUrl = &v
}

// GetSubpages returns the Subpages field value if set, zero value otherwise.
func (o *Creative) GetSubpages() []Subpage {
	if o == nil || IsNil(o.Subpages) {
		var ret []Subpage
		return ret
	}
	return o.Subpages
}

// GetSubpagesOk returns a tuple with the Subpages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetSubpagesOk() ([]Subpage, bool) {
	if o == nil || IsNil(o.Subpages) {
		return nil, false
	}
	return o.Subpages, true
}

// HasSubpages returns a boolean if a field has been set.
func (o *Creative) HasSubpages() bool {
	if o != nil && !IsNil(o.Subpages) {
		return true
	}

	return false
}

// SetSubpages gets a reference to the given []Subpage and assigns it to the Subpages field.
func (o *Creative) SetSubpages(v []Subpage) {
	o.Subpages = v
}

// GetOriginalHeadline returns the OriginalHeadline field value if set, zero value otherwise.
func (o *Creative) GetOriginalHeadline() string {
	if o == nil || IsNil(o.OriginalHeadline) {
		var ret string
		return ret
	}
	return *o.OriginalHeadline
}

// GetOriginalHeadlineOk returns a tuple with the OriginalHeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetOriginalHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalHeadline) {
		return nil, false
	}
	return o.OriginalHeadline, true
}

// HasOriginalHeadline returns a boolean if a field has been set.
func (o *Creative) HasOriginalHeadline() bool {
	if o != nil && !IsNil(o.OriginalHeadline) {
		return true
	}

	return false
}

// SetOriginalHeadline gets a reference to the given string and assigns it to the OriginalHeadline field.
func (o *Creative) SetOriginalHeadline(v string) {
	o.OriginalHeadline = &v
}

// GetVideoAssetIds returns the VideoAssetIds field value if set, zero value otherwise.
func (o *Creative) GetVideoAssetIds() []string {
	if o == nil || IsNil(o.VideoAssetIds) {
		var ret []string
		return ret
	}
	return o.VideoAssetIds
}

// GetVideoAssetIdsOk returns a tuple with the VideoAssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetVideoAssetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VideoAssetIds) {
		return nil, false
	}
	return o.VideoAssetIds, true
}

// HasVideoAssetIds returns a boolean if a field has been set.
func (o *Creative) HasVideoAssetIds() bool {
	if o != nil && !IsNil(o.VideoAssetIds) {
		return true
	}

	return false
}

// SetVideoAssetIds gets a reference to the given []string and assigns it to the VideoAssetIds field.
func (o *Creative) SetVideoAssetIds(v []string) {
	o.VideoAssetIds = v
}

// GetBrandLogoAssetID returns the BrandLogoAssetID field value if set, zero value otherwise.
func (o *Creative) GetBrandLogoAssetID() string {
	if o == nil || IsNil(o.BrandLogoAssetID) {
		var ret string
		return ret
	}
	return *o.BrandLogoAssetID
}

// GetBrandLogoAssetIDOk returns a tuple with the BrandLogoAssetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetBrandLogoAssetIDOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogoAssetID) {
		return nil, false
	}
	return o.BrandLogoAssetID, true
}

// HasBrandLogoAssetID returns a boolean if a field has been set.
func (o *Creative) HasBrandLogoAssetID() bool {
	if o != nil && !IsNil(o.BrandLogoAssetID) {
		return true
	}

	return false
}

// SetBrandLogoAssetID gets a reference to the given string and assigns it to the BrandLogoAssetID field.
func (o *Creative) SetBrandLogoAssetID(v string) {
	o.BrandLogoAssetID = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *Creative) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *Creative) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *Creative) SetHeadline(v string) {
	o.Headline = &v
}

// GetCreativeStatus returns the CreativeStatus field value if set, zero value otherwise.
func (o *Creative) GetCreativeStatus() CreativeStatus {
	if o == nil || IsNil(o.CreativeStatus) {
		var ret CreativeStatus
		return ret
	}
	return *o.CreativeStatus
}

// GetCreativeStatusOk returns a tuple with the CreativeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creative) GetCreativeStatusOk() (*CreativeStatus, bool) {
	if o == nil || IsNil(o.CreativeStatus) {
		return nil, false
	}
	return o.CreativeStatus, true
}

// HasCreativeStatus returns a boolean if a field has been set.
func (o *Creative) HasCreativeStatus() bool {
	if o != nil && !IsNil(o.CreativeStatus) {
		return true
	}

	return false
}

// SetCreativeStatus gets a reference to the given CreativeStatus and assigns it to the CreativeStatus field.
func (o *Creative) SetCreativeStatus(v CreativeStatus) {
	o.CreativeStatus = &v
}

func (o Creative) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ConsentToTranslate) {
		toSerialize["consentToTranslate"] = o.ConsentToTranslate
	}
	if !IsNil(o.CustomImages) {
		toSerialize["customImages"] = o.CustomImages
	}
	if !IsNil(o.CustomImageCrop) {
		toSerialize["customImageCrop"] = o.CustomImageCrop
	}
	if !IsNil(o.CustomImageUrl) {
		toSerialize["customImageUrl"] = o.CustomImageUrl
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
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
	if !IsNil(o.BrandLogoAssetID) {
		toSerialize["brandLogoAssetID"] = o.BrandLogoAssetID
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	if !IsNil(o.CreativeStatus) {
		toSerialize["creativeStatus"] = o.CreativeStatus
	}
	return toSerialize, nil
}

type NullableCreative struct {
	value *Creative
	isSet bool
}

func (v NullableCreative) Get() *Creative {
	return v.value
}

func (v *NullableCreative) Set(val *Creative) {
	v.value = val
	v.isSet = true
}

func (v NullableCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreative(val *Creative) *NullableCreative {
	return &NullableCreative{value: val, isSet: true}
}

func (v NullableCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
