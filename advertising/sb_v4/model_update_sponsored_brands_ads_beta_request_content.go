package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsAdsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsAdsBetaRequestContent{}

// UpdateSponsoredBrandsAdsBetaRequestContent struct for UpdateSponsoredBrandsAdsBetaRequestContent
type UpdateSponsoredBrandsAdsBetaRequestContent struct {
	Ads []UpdateAd `json:"ads"`
}

type _UpdateSponsoredBrandsAdsBetaRequestContent UpdateSponsoredBrandsAdsBetaRequestContent

// NewUpdateSponsoredBrandsAdsBetaRequestContent instantiates a new UpdateSponsoredBrandsAdsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsAdsBetaRequestContent(ads []UpdateAd) *UpdateSponsoredBrandsAdsBetaRequestContent {
	this := UpdateSponsoredBrandsAdsBetaRequestContent{}
	this.Ads = ads
	return &this
}

// NewUpdateSponsoredBrandsAdsBetaRequestContentWithDefaults instantiates a new UpdateSponsoredBrandsAdsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsAdsBetaRequestContentWithDefaults() *UpdateSponsoredBrandsAdsBetaRequestContent {
	this := UpdateSponsoredBrandsAdsBetaRequestContent{}
	return &this
}

// GetAds returns the Ads field value
func (o *UpdateSponsoredBrandsAdsBetaRequestContent) GetAds() []UpdateAd {
	if o == nil {
		var ret []UpdateAd
		return ret
	}

	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsAdsBetaRequestContent) GetAdsOk() ([]UpdateAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads, true
}

// SetAds sets field value
func (o *UpdateSponsoredBrandsAdsBetaRequestContent) SetAds(v []UpdateAd) {
	o.Ads = v
}

func (o UpdateSponsoredBrandsAdsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ads"] = o.Ads
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsAdsBetaRequestContent struct {
	value *UpdateSponsoredBrandsAdsBetaRequestContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsAdsBetaRequestContent) Get() *UpdateSponsoredBrandsAdsBetaRequestContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsAdsBetaRequestContent) Set(val *UpdateSponsoredBrandsAdsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsAdsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsAdsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsAdsBetaRequestContent(val *UpdateSponsoredBrandsAdsBetaRequestContent) *NullableUpdateSponsoredBrandsAdsBetaRequestContent {
	return &NullableUpdateSponsoredBrandsAdsBetaRequestContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsAdsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsAdsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
