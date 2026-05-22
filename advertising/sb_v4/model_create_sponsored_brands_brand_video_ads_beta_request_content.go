package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsBrandVideoAdsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsBrandVideoAdsBetaRequestContent{}

// CreateSponsoredBrandsBrandVideoAdsBetaRequestContent struct for CreateSponsoredBrandsBrandVideoAdsBetaRequestContent
type CreateSponsoredBrandsBrandVideoAdsBetaRequestContent struct {
	Ads []CreateBrandVideoAd `json:"ads"`
}

type _CreateSponsoredBrandsBrandVideoAdsBetaRequestContent CreateSponsoredBrandsBrandVideoAdsBetaRequestContent

// NewCreateSponsoredBrandsBrandVideoAdsBetaRequestContent instantiates a new CreateSponsoredBrandsBrandVideoAdsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsBrandVideoAdsBetaRequestContent(ads []CreateBrandVideoAd) *CreateSponsoredBrandsBrandVideoAdsBetaRequestContent {
	this := CreateSponsoredBrandsBrandVideoAdsBetaRequestContent{}
	this.Ads = ads
	return &this
}

// NewCreateSponsoredBrandsBrandVideoAdsBetaRequestContentWithDefaults instantiates a new CreateSponsoredBrandsBrandVideoAdsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsBrandVideoAdsBetaRequestContentWithDefaults() *CreateSponsoredBrandsBrandVideoAdsBetaRequestContent {
	this := CreateSponsoredBrandsBrandVideoAdsBetaRequestContent{}
	return &this
}

// GetAds returns the Ads field value
func (o *CreateSponsoredBrandsBrandVideoAdsBetaRequestContent) GetAds() []CreateBrandVideoAd {
	if o == nil {
		var ret []CreateBrandVideoAd
		return ret
	}

	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsBrandVideoAdsBetaRequestContent) GetAdsOk() ([]CreateBrandVideoAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads, true
}

// SetAds sets field value
func (o *CreateSponsoredBrandsBrandVideoAdsBetaRequestContent) SetAds(v []CreateBrandVideoAd) {
	o.Ads = v
}

func (o CreateSponsoredBrandsBrandVideoAdsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ads"] = o.Ads
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsBrandVideoAdsBetaRequestContent struct {
	value *CreateSponsoredBrandsBrandVideoAdsBetaRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsBetaRequestContent) Get() *CreateSponsoredBrandsBrandVideoAdsBetaRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsBetaRequestContent) Set(val *CreateSponsoredBrandsBrandVideoAdsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsBrandVideoAdsBetaRequestContent(val *CreateSponsoredBrandsBrandVideoAdsBetaRequestContent) *NullableCreateSponsoredBrandsBrandVideoAdsBetaRequestContent {
	return &NullableCreateSponsoredBrandsBrandVideoAdsBetaRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
