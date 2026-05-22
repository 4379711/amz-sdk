package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsBrandVideoAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsBrandVideoAdsRequestContent{}

// CreateSponsoredBrandsBrandVideoAdsRequestContent struct for CreateSponsoredBrandsBrandVideoAdsRequestContent
type CreateSponsoredBrandsBrandVideoAdsRequestContent struct {
	Ads []CreateBrandVideoAd `json:"ads"`
}

type _CreateSponsoredBrandsBrandVideoAdsRequestContent CreateSponsoredBrandsBrandVideoAdsRequestContent

// NewCreateSponsoredBrandsBrandVideoAdsRequestContent instantiates a new CreateSponsoredBrandsBrandVideoAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsBrandVideoAdsRequestContent(ads []CreateBrandVideoAd) *CreateSponsoredBrandsBrandVideoAdsRequestContent {
	this := CreateSponsoredBrandsBrandVideoAdsRequestContent{}
	this.Ads = ads
	return &this
}

// NewCreateSponsoredBrandsBrandVideoAdsRequestContentWithDefaults instantiates a new CreateSponsoredBrandsBrandVideoAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsBrandVideoAdsRequestContentWithDefaults() *CreateSponsoredBrandsBrandVideoAdsRequestContent {
	this := CreateSponsoredBrandsBrandVideoAdsRequestContent{}
	return &this
}

// GetAds returns the Ads field value
func (o *CreateSponsoredBrandsBrandVideoAdsRequestContent) GetAds() []CreateBrandVideoAd {
	if o == nil {
		var ret []CreateBrandVideoAd
		return ret
	}

	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsBrandVideoAdsRequestContent) GetAdsOk() ([]CreateBrandVideoAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads, true
}

// SetAds sets field value
func (o *CreateSponsoredBrandsBrandVideoAdsRequestContent) SetAds(v []CreateBrandVideoAd) {
	o.Ads = v
}

func (o CreateSponsoredBrandsBrandVideoAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ads"] = o.Ads
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsBrandVideoAdsRequestContent struct {
	value *CreateSponsoredBrandsBrandVideoAdsRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsRequestContent) Get() *CreateSponsoredBrandsBrandVideoAdsRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsRequestContent) Set(val *CreateSponsoredBrandsBrandVideoAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsBrandVideoAdsRequestContent(val *CreateSponsoredBrandsBrandVideoAdsRequestContent) *NullableCreateSponsoredBrandsBrandVideoAdsRequestContent {
	return &NullableCreateSponsoredBrandsBrandVideoAdsRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
