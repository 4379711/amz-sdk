package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateExtendedProductCollectionCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateExtendedProductCollectionCreative{}

// CreateExtendedProductCollectionCreative struct for CreateExtendedProductCollectionCreative
type CreateExtendedProductCollectionCreative struct {
	BrandLogoCrop *BrandLogoCrop `json:"brandLogoCrop,omitempty"`
	Asins         []string       `json:"asins,omitempty"`
	BrandName     *string        `json:"brandName,omitempty"`
	// Requires minimum one custom image. You can add an optional collection of custom images that can be displayed on the ad as slideshow. Learn more about slideshow here https://advertising.amazon.com/resources/whats-new/slideshow-ads-creative-for-sponsored-brands/
	CustomImages []CustomImage `json:"customImages,omitempty"`
	// If set to true and the headline and/or video are not in the marketplace's default language, Amazon will attempt to translate them to the marketplace's default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information.
	ConsentToTranslate *bool   `json:"consentToTranslate,omitempty"`
	BrandLogoAssetID   *string `json:"brandLogoAssetID,omitempty"`
	Headline           *string `json:"headline,omitempty"`
}

// NewCreateExtendedProductCollectionCreative instantiates a new CreateExtendedProductCollectionCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateExtendedProductCollectionCreative() *CreateExtendedProductCollectionCreative {
	this := CreateExtendedProductCollectionCreative{}
	return &this
}

// NewCreateExtendedProductCollectionCreativeWithDefaults instantiates a new CreateExtendedProductCollectionCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateExtendedProductCollectionCreativeWithDefaults() *CreateExtendedProductCollectionCreative {
	this := CreateExtendedProductCollectionCreative{}
	return &this
}

// GetBrandLogoCrop returns the BrandLogoCrop field value if set, zero value otherwise.
func (o *CreateExtendedProductCollectionCreative) GetBrandLogoCrop() BrandLogoCrop {
	if o == nil || IsNil(o.BrandLogoCrop) {
		var ret BrandLogoCrop
		return ret
	}
	return *o.BrandLogoCrop
}

// GetBrandLogoCropOk returns a tuple with the BrandLogoCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionCreative) GetBrandLogoCropOk() (*BrandLogoCrop, bool) {
	if o == nil || IsNil(o.BrandLogoCrop) {
		return nil, false
	}
	return o.BrandLogoCrop, true
}

// HasBrandLogoCrop returns a boolean if a field has been set.
func (o *CreateExtendedProductCollectionCreative) HasBrandLogoCrop() bool {
	if o != nil && !IsNil(o.BrandLogoCrop) {
		return true
	}

	return false
}

// SetBrandLogoCrop gets a reference to the given BrandLogoCrop and assigns it to the BrandLogoCrop field.
func (o *CreateExtendedProductCollectionCreative) SetBrandLogoCrop(v BrandLogoCrop) {
	o.BrandLogoCrop = &v
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *CreateExtendedProductCollectionCreative) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionCreative) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *CreateExtendedProductCollectionCreative) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *CreateExtendedProductCollectionCreative) SetAsins(v []string) {
	o.Asins = v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *CreateExtendedProductCollectionCreative) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionCreative) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *CreateExtendedProductCollectionCreative) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *CreateExtendedProductCollectionCreative) SetBrandName(v string) {
	o.BrandName = &v
}

// GetCustomImages returns the CustomImages field value if set, zero value otherwise.
func (o *CreateExtendedProductCollectionCreative) GetCustomImages() []CustomImage {
	if o == nil || IsNil(o.CustomImages) {
		var ret []CustomImage
		return ret
	}
	return o.CustomImages
}

// GetCustomImagesOk returns a tuple with the CustomImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionCreative) GetCustomImagesOk() ([]CustomImage, bool) {
	if o == nil || IsNil(o.CustomImages) {
		return nil, false
	}
	return o.CustomImages, true
}

// HasCustomImages returns a boolean if a field has been set.
func (o *CreateExtendedProductCollectionCreative) HasCustomImages() bool {
	if o != nil && !IsNil(o.CustomImages) {
		return true
	}

	return false
}

// SetCustomImages gets a reference to the given []CustomImage and assigns it to the CustomImages field.
func (o *CreateExtendedProductCollectionCreative) SetCustomImages(v []CustomImage) {
	o.CustomImages = v
}

// GetConsentToTranslate returns the ConsentToTranslate field value if set, zero value otherwise.
func (o *CreateExtendedProductCollectionCreative) GetConsentToTranslate() bool {
	if o == nil || IsNil(o.ConsentToTranslate) {
		var ret bool
		return ret
	}
	return *o.ConsentToTranslate
}

// GetConsentToTranslateOk returns a tuple with the ConsentToTranslate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionCreative) GetConsentToTranslateOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsentToTranslate) {
		return nil, false
	}
	return o.ConsentToTranslate, true
}

// HasConsentToTranslate returns a boolean if a field has been set.
func (o *CreateExtendedProductCollectionCreative) HasConsentToTranslate() bool {
	if o != nil && !IsNil(o.ConsentToTranslate) {
		return true
	}

	return false
}

// SetConsentToTranslate gets a reference to the given bool and assigns it to the ConsentToTranslate field.
func (o *CreateExtendedProductCollectionCreative) SetConsentToTranslate(v bool) {
	o.ConsentToTranslate = &v
}

// GetBrandLogoAssetID returns the BrandLogoAssetID field value if set, zero value otherwise.
func (o *CreateExtendedProductCollectionCreative) GetBrandLogoAssetID() string {
	if o == nil || IsNil(o.BrandLogoAssetID) {
		var ret string
		return ret
	}
	return *o.BrandLogoAssetID
}

// GetBrandLogoAssetIDOk returns a tuple with the BrandLogoAssetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionCreative) GetBrandLogoAssetIDOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogoAssetID) {
		return nil, false
	}
	return o.BrandLogoAssetID, true
}

// HasBrandLogoAssetID returns a boolean if a field has been set.
func (o *CreateExtendedProductCollectionCreative) HasBrandLogoAssetID() bool {
	if o != nil && !IsNil(o.BrandLogoAssetID) {
		return true
	}

	return false
}

// SetBrandLogoAssetID gets a reference to the given string and assigns it to the BrandLogoAssetID field.
func (o *CreateExtendedProductCollectionCreative) SetBrandLogoAssetID(v string) {
	o.BrandLogoAssetID = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *CreateExtendedProductCollectionCreative) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionCreative) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *CreateExtendedProductCollectionCreative) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *CreateExtendedProductCollectionCreative) SetHeadline(v string) {
	o.Headline = &v
}

func (o CreateExtendedProductCollectionCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandLogoCrop) {
		toSerialize["brandLogoCrop"] = o.BrandLogoCrop
	}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.BrandName) {
		toSerialize["brandName"] = o.BrandName
	}
	if !IsNil(o.CustomImages) {
		toSerialize["customImages"] = o.CustomImages
	}
	if !IsNil(o.ConsentToTranslate) {
		toSerialize["consentToTranslate"] = o.ConsentToTranslate
	}
	if !IsNil(o.BrandLogoAssetID) {
		toSerialize["brandLogoAssetID"] = o.BrandLogoAssetID
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	return toSerialize, nil
}

type NullableCreateExtendedProductCollectionCreative struct {
	value *CreateExtendedProductCollectionCreative
	isSet bool
}

func (v NullableCreateExtendedProductCollectionCreative) Get() *CreateExtendedProductCollectionCreative {
	return v.value
}

func (v *NullableCreateExtendedProductCollectionCreative) Set(val *CreateExtendedProductCollectionCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateExtendedProductCollectionCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateExtendedProductCollectionCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateExtendedProductCollectionCreative(val *CreateExtendedProductCollectionCreative) *NullableCreateExtendedProductCollectionCreative {
	return &NullableCreateExtendedProductCollectionCreative{value: val, isSet: true}
}

func (v NullableCreateExtendedProductCollectionCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateExtendedProductCollectionCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
