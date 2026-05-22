package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandStoreSpotlightAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandStoreSpotlightAdsRequestContent{}

// CreateSponsoredBrandStoreSpotlightAdsRequestContent struct for CreateSponsoredBrandStoreSpotlightAdsRequestContent
type CreateSponsoredBrandStoreSpotlightAdsRequestContent struct {
	Ads []CreateStoreSpotlightAd `json:"ads"`
}

type _CreateSponsoredBrandStoreSpotlightAdsRequestContent CreateSponsoredBrandStoreSpotlightAdsRequestContent

// NewCreateSponsoredBrandStoreSpotlightAdsRequestContent instantiates a new CreateSponsoredBrandStoreSpotlightAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandStoreSpotlightAdsRequestContent(ads []CreateStoreSpotlightAd) *CreateSponsoredBrandStoreSpotlightAdsRequestContent {
	this := CreateSponsoredBrandStoreSpotlightAdsRequestContent{}
	this.Ads = ads
	return &this
}

// NewCreateSponsoredBrandStoreSpotlightAdsRequestContentWithDefaults instantiates a new CreateSponsoredBrandStoreSpotlightAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandStoreSpotlightAdsRequestContentWithDefaults() *CreateSponsoredBrandStoreSpotlightAdsRequestContent {
	this := CreateSponsoredBrandStoreSpotlightAdsRequestContent{}
	return &this
}

// GetAds returns the Ads field value
func (o *CreateSponsoredBrandStoreSpotlightAdsRequestContent) GetAds() []CreateStoreSpotlightAd {
	if o == nil {
		var ret []CreateStoreSpotlightAd
		return ret
	}

	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandStoreSpotlightAdsRequestContent) GetAdsOk() ([]CreateStoreSpotlightAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads, true
}

// SetAds sets field value
func (o *CreateSponsoredBrandStoreSpotlightAdsRequestContent) SetAds(v []CreateStoreSpotlightAd) {
	o.Ads = v
}

func (o CreateSponsoredBrandStoreSpotlightAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ads"] = o.Ads
	return toSerialize, nil
}

type NullableCreateSponsoredBrandStoreSpotlightAdsRequestContent struct {
	value *CreateSponsoredBrandStoreSpotlightAdsRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsRequestContent) Get() *CreateSponsoredBrandStoreSpotlightAdsRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsRequestContent) Set(val *CreateSponsoredBrandStoreSpotlightAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandStoreSpotlightAdsRequestContent(val *CreateSponsoredBrandStoreSpotlightAdsRequestContent) *NullableCreateSponsoredBrandStoreSpotlightAdsRequestContent {
	return &NullableCreateSponsoredBrandStoreSpotlightAdsRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
