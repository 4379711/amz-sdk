package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductCollectionCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductCollectionCreative{}

// ProductCollectionCreative struct for ProductCollectionCreative
type ProductCollectionCreative struct {
	// An array of ASINs associated with the creative.
	Asins         []string   `json:"asins"`
	BrandLogoCrop *AssetCrop `json:"brandLogoCrop,omitempty"`
	// The displayed brand name in the ad headline. Maximum length is 30 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements.
	BrandName string `json:"brandName"`
	// The identifier of the Custom image from the Store assets library. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#customimage) for more information on what constitutes a valid Custom image.
	CustomImageAssetId *string    `json:"customImageAssetId,omitempty"`
	CustomImageCrop    *AssetCrop `json:"customImageCrop,omitempty"`
	// The identifier of the [brand logo](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#brandlogo) image from the brand store's asset library. Note that for campaigns created in the Amazon Advertising console prior to release of the brand store's assets library, responses will not include a value for this field.
	BrandLogoAssetId string `json:"brandLogoAssetId"`
	// The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements.
	Headline string `json:"headline"`
}

type _ProductCollectionCreative ProductCollectionCreative

// NewProductCollectionCreative instantiates a new ProductCollectionCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductCollectionCreative(asins []string, brandName string, brandLogoAssetId string, headline string) *ProductCollectionCreative {
	this := ProductCollectionCreative{}
	this.Asins = asins
	this.BrandName = brandName
	this.BrandLogoAssetId = brandLogoAssetId
	this.Headline = headline
	return &this
}

// NewProductCollectionCreativeWithDefaults instantiates a new ProductCollectionCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductCollectionCreativeWithDefaults() *ProductCollectionCreative {
	this := ProductCollectionCreative{}
	return &this
}

// GetAsins returns the Asins field value
func (o *ProductCollectionCreative) GetAsins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value
// and a boolean to check if the value has been set.
func (o *ProductCollectionCreative) GetAsinsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asins, true
}

// SetAsins sets field value
func (o *ProductCollectionCreative) SetAsins(v []string) {
	o.Asins = v
}

// GetBrandLogoCrop returns the BrandLogoCrop field value if set, zero value otherwise.
func (o *ProductCollectionCreative) GetBrandLogoCrop() AssetCrop {
	if o == nil || IsNil(o.BrandLogoCrop) {
		var ret AssetCrop
		return ret
	}
	return *o.BrandLogoCrop
}

// GetBrandLogoCropOk returns a tuple with the BrandLogoCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCollectionCreative) GetBrandLogoCropOk() (*AssetCrop, bool) {
	if o == nil || IsNil(o.BrandLogoCrop) {
		return nil, false
	}
	return o.BrandLogoCrop, true
}

// HasBrandLogoCrop returns a boolean if a field has been set.
func (o *ProductCollectionCreative) HasBrandLogoCrop() bool {
	if o != nil && !IsNil(o.BrandLogoCrop) {
		return true
	}

	return false
}

// SetBrandLogoCrop gets a reference to the given AssetCrop and assigns it to the BrandLogoCrop field.
func (o *ProductCollectionCreative) SetBrandLogoCrop(v AssetCrop) {
	o.BrandLogoCrop = &v
}

// GetBrandName returns the BrandName field value
func (o *ProductCollectionCreative) GetBrandName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value
// and a boolean to check if the value has been set.
func (o *ProductCollectionCreative) GetBrandNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandName, true
}

// SetBrandName sets field value
func (o *ProductCollectionCreative) SetBrandName(v string) {
	o.BrandName = v
}

// GetCustomImageAssetId returns the CustomImageAssetId field value if set, zero value otherwise.
func (o *ProductCollectionCreative) GetCustomImageAssetId() string {
	if o == nil || IsNil(o.CustomImageAssetId) {
		var ret string
		return ret
	}
	return *o.CustomImageAssetId
}

// GetCustomImageAssetIdOk returns a tuple with the CustomImageAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCollectionCreative) GetCustomImageAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomImageAssetId) {
		return nil, false
	}
	return o.CustomImageAssetId, true
}

// HasCustomImageAssetId returns a boolean if a field has been set.
func (o *ProductCollectionCreative) HasCustomImageAssetId() bool {
	if o != nil && !IsNil(o.CustomImageAssetId) {
		return true
	}

	return false
}

// SetCustomImageAssetId gets a reference to the given string and assigns it to the CustomImageAssetId field.
func (o *ProductCollectionCreative) SetCustomImageAssetId(v string) {
	o.CustomImageAssetId = &v
}

// GetCustomImageCrop returns the CustomImageCrop field value if set, zero value otherwise.
func (o *ProductCollectionCreative) GetCustomImageCrop() AssetCrop {
	if o == nil || IsNil(o.CustomImageCrop) {
		var ret AssetCrop
		return ret
	}
	return *o.CustomImageCrop
}

// GetCustomImageCropOk returns a tuple with the CustomImageCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCollectionCreative) GetCustomImageCropOk() (*AssetCrop, bool) {
	if o == nil || IsNil(o.CustomImageCrop) {
		return nil, false
	}
	return o.CustomImageCrop, true
}

// HasCustomImageCrop returns a boolean if a field has been set.
func (o *ProductCollectionCreative) HasCustomImageCrop() bool {
	if o != nil && !IsNil(o.CustomImageCrop) {
		return true
	}

	return false
}

// SetCustomImageCrop gets a reference to the given AssetCrop and assigns it to the CustomImageCrop field.
func (o *ProductCollectionCreative) SetCustomImageCrop(v AssetCrop) {
	o.CustomImageCrop = &v
}

// GetBrandLogoAssetId returns the BrandLogoAssetId field value
func (o *ProductCollectionCreative) GetBrandLogoAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandLogoAssetId
}

// GetBrandLogoAssetIdOk returns a tuple with the BrandLogoAssetId field value
// and a boolean to check if the value has been set.
func (o *ProductCollectionCreative) GetBrandLogoAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandLogoAssetId, true
}

// SetBrandLogoAssetId sets field value
func (o *ProductCollectionCreative) SetBrandLogoAssetId(v string) {
	o.BrandLogoAssetId = v
}

// GetHeadline returns the Headline field value
func (o *ProductCollectionCreative) GetHeadline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value
// and a boolean to check if the value has been set.
func (o *ProductCollectionCreative) GetHeadlineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headline, true
}

// SetHeadline sets field value
func (o *ProductCollectionCreative) SetHeadline(v string) {
	o.Headline = v
}

func (o ProductCollectionCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asins"] = o.Asins
	if !IsNil(o.BrandLogoCrop) {
		toSerialize["brandLogoCrop"] = o.BrandLogoCrop
	}
	toSerialize["brandName"] = o.BrandName
	if !IsNil(o.CustomImageAssetId) {
		toSerialize["customImageAssetId"] = o.CustomImageAssetId
	}
	if !IsNil(o.CustomImageCrop) {
		toSerialize["customImageCrop"] = o.CustomImageCrop
	}
	toSerialize["brandLogoAssetId"] = o.BrandLogoAssetId
	toSerialize["headline"] = o.Headline
	return toSerialize, nil
}

type NullableProductCollectionCreative struct {
	value *ProductCollectionCreative
	isSet bool
}

func (v NullableProductCollectionCreative) Get() *ProductCollectionCreative {
	return v.value
}

func (v *NullableProductCollectionCreative) Set(val *ProductCollectionCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableProductCollectionCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableProductCollectionCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductCollectionCreative(val *ProductCollectionCreative) *NullableProductCollectionCreative {
	return &NullableProductCollectionCreative{value: val, isSet: true}
}

func (v NullableProductCollectionCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductCollectionCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
