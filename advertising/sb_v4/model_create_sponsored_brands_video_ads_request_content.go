package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsVideoAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsVideoAdsRequestContent{}

// CreateSponsoredBrandsVideoAdsRequestContent struct for CreateSponsoredBrandsVideoAdsRequestContent
type CreateSponsoredBrandsVideoAdsRequestContent struct {
	Ads []CreateVideoAd `json:"ads"`
}

type _CreateSponsoredBrandsVideoAdsRequestContent CreateSponsoredBrandsVideoAdsRequestContent

// NewCreateSponsoredBrandsVideoAdsRequestContent instantiates a new CreateSponsoredBrandsVideoAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsVideoAdsRequestContent(ads []CreateVideoAd) *CreateSponsoredBrandsVideoAdsRequestContent {
	this := CreateSponsoredBrandsVideoAdsRequestContent{}
	this.Ads = ads
	return &this
}

// NewCreateSponsoredBrandsVideoAdsRequestContentWithDefaults instantiates a new CreateSponsoredBrandsVideoAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsVideoAdsRequestContentWithDefaults() *CreateSponsoredBrandsVideoAdsRequestContent {
	this := CreateSponsoredBrandsVideoAdsRequestContent{}
	return &this
}

// GetAds returns the Ads field value
func (o *CreateSponsoredBrandsVideoAdsRequestContent) GetAds() []CreateVideoAd {
	if o == nil {
		var ret []CreateVideoAd
		return ret
	}

	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsVideoAdsRequestContent) GetAdsOk() ([]CreateVideoAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads, true
}

// SetAds sets field value
func (o *CreateSponsoredBrandsVideoAdsRequestContent) SetAds(v []CreateVideoAd) {
	o.Ads = v
}

func (o CreateSponsoredBrandsVideoAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ads"] = o.Ads
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsVideoAdsRequestContent struct {
	value *CreateSponsoredBrandsVideoAdsRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsVideoAdsRequestContent) Get() *CreateSponsoredBrandsVideoAdsRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsVideoAdsRequestContent) Set(val *CreateSponsoredBrandsVideoAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsVideoAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsVideoAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsVideoAdsRequestContent(val *CreateSponsoredBrandsVideoAdsRequestContent) *NullableCreateSponsoredBrandsVideoAdsRequestContent {
	return &NullableCreateSponsoredBrandsVideoAdsRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsVideoAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsVideoAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
