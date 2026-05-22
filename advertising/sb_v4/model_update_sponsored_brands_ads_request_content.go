package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsAdsRequestContent{}

// UpdateSponsoredBrandsAdsRequestContent struct for UpdateSponsoredBrandsAdsRequestContent
type UpdateSponsoredBrandsAdsRequestContent struct {
	Ads []UpdateAd `json:"ads"`
}

type _UpdateSponsoredBrandsAdsRequestContent UpdateSponsoredBrandsAdsRequestContent

// NewUpdateSponsoredBrandsAdsRequestContent instantiates a new UpdateSponsoredBrandsAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsAdsRequestContent(ads []UpdateAd) *UpdateSponsoredBrandsAdsRequestContent {
	this := UpdateSponsoredBrandsAdsRequestContent{}
	this.Ads = ads
	return &this
}

// NewUpdateSponsoredBrandsAdsRequestContentWithDefaults instantiates a new UpdateSponsoredBrandsAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsAdsRequestContentWithDefaults() *UpdateSponsoredBrandsAdsRequestContent {
	this := UpdateSponsoredBrandsAdsRequestContent{}
	return &this
}

// GetAds returns the Ads field value
func (o *UpdateSponsoredBrandsAdsRequestContent) GetAds() []UpdateAd {
	if o == nil {
		var ret []UpdateAd
		return ret
	}

	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsAdsRequestContent) GetAdsOk() ([]UpdateAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads, true
}

// SetAds sets field value
func (o *UpdateSponsoredBrandsAdsRequestContent) SetAds(v []UpdateAd) {
	o.Ads = v
}

func (o UpdateSponsoredBrandsAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ads"] = o.Ads
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsAdsRequestContent struct {
	value *UpdateSponsoredBrandsAdsRequestContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsAdsRequestContent) Get() *UpdateSponsoredBrandsAdsRequestContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsAdsRequestContent) Set(val *UpdateSponsoredBrandsAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsAdsRequestContent(val *UpdateSponsoredBrandsAdsRequestContent) *NullableUpdateSponsoredBrandsAdsRequestContent {
	return &NullableUpdateSponsoredBrandsAdsRequestContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
