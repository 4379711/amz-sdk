package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsProductCollectionAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsProductCollectionAdsRequestContent{}

// CreateSponsoredBrandsProductCollectionAdsRequestContent struct for CreateSponsoredBrandsProductCollectionAdsRequestContent
type CreateSponsoredBrandsProductCollectionAdsRequestContent struct {
	Ads []CreateProductCollectionAd `json:"ads"`
}

type _CreateSponsoredBrandsProductCollectionAdsRequestContent CreateSponsoredBrandsProductCollectionAdsRequestContent

// NewCreateSponsoredBrandsProductCollectionAdsRequestContent instantiates a new CreateSponsoredBrandsProductCollectionAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsProductCollectionAdsRequestContent(ads []CreateProductCollectionAd) *CreateSponsoredBrandsProductCollectionAdsRequestContent {
	this := CreateSponsoredBrandsProductCollectionAdsRequestContent{}
	this.Ads = ads
	return &this
}

// NewCreateSponsoredBrandsProductCollectionAdsRequestContentWithDefaults instantiates a new CreateSponsoredBrandsProductCollectionAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsProductCollectionAdsRequestContentWithDefaults() *CreateSponsoredBrandsProductCollectionAdsRequestContent {
	this := CreateSponsoredBrandsProductCollectionAdsRequestContent{}
	return &this
}

// GetAds returns the Ads field value
func (o *CreateSponsoredBrandsProductCollectionAdsRequestContent) GetAds() []CreateProductCollectionAd {
	if o == nil {
		var ret []CreateProductCollectionAd
		return ret
	}

	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsProductCollectionAdsRequestContent) GetAdsOk() ([]CreateProductCollectionAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads, true
}

// SetAds sets field value
func (o *CreateSponsoredBrandsProductCollectionAdsRequestContent) SetAds(v []CreateProductCollectionAd) {
	o.Ads = v
}

func (o CreateSponsoredBrandsProductCollectionAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ads"] = o.Ads
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsProductCollectionAdsRequestContent struct {
	value *CreateSponsoredBrandsProductCollectionAdsRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsRequestContent) Get() *CreateSponsoredBrandsProductCollectionAdsRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsRequestContent) Set(val *CreateSponsoredBrandsProductCollectionAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsProductCollectionAdsRequestContent(val *CreateSponsoredBrandsProductCollectionAdsRequestContent) *NullableCreateSponsoredBrandsProductCollectionAdsRequestContent {
	return &NullableCreateSponsoredBrandsProductCollectionAdsRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
