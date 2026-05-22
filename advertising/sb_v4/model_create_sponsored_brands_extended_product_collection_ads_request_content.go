package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent{}

// CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent struct for CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent
type CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent struct {
	// An array of Product Collection ad objects to create. Maximum length of the array is 10 objects.
	Ads []CreateExtendedProductCollectionAd `json:"ads"`
}

type _CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent

// NewCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent instantiates a new CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent(ads []CreateExtendedProductCollectionAd) *CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent {
	this := CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent{}
	this.Ads = ads
	return &this
}

// NewCreateSponsoredBrandsExtendedProductCollectionAdsRequestContentWithDefaults instantiates a new CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsExtendedProductCollectionAdsRequestContentWithDefaults() *CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent {
	this := CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent{}
	return &this
}

// GetAds returns the Ads field value
func (o *CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) GetAds() []CreateExtendedProductCollectionAd {
	if o == nil {
		var ret []CreateExtendedProductCollectionAd
		return ret
	}

	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) GetAdsOk() ([]CreateExtendedProductCollectionAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads, true
}

// SetAds sets field value
func (o *CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) SetAds(v []CreateExtendedProductCollectionAd) {
	o.Ads = v
}

func (o CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ads"] = o.Ads
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent struct {
	value *CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) Get() *CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) Set(val *CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent(val *CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) *NullableCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent {
	return &NullableCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
