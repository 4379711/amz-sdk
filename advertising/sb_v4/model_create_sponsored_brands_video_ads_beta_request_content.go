package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsVideoAdsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsVideoAdsBetaRequestContent{}

// CreateSponsoredBrandsVideoAdsBetaRequestContent struct for CreateSponsoredBrandsVideoAdsBetaRequestContent
type CreateSponsoredBrandsVideoAdsBetaRequestContent struct {
	Ads []CreateVideoAd `json:"ads"`
}

type _CreateSponsoredBrandsVideoAdsBetaRequestContent CreateSponsoredBrandsVideoAdsBetaRequestContent

// NewCreateSponsoredBrandsVideoAdsBetaRequestContent instantiates a new CreateSponsoredBrandsVideoAdsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsVideoAdsBetaRequestContent(ads []CreateVideoAd) *CreateSponsoredBrandsVideoAdsBetaRequestContent {
	this := CreateSponsoredBrandsVideoAdsBetaRequestContent{}
	this.Ads = ads
	return &this
}

// NewCreateSponsoredBrandsVideoAdsBetaRequestContentWithDefaults instantiates a new CreateSponsoredBrandsVideoAdsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsVideoAdsBetaRequestContentWithDefaults() *CreateSponsoredBrandsVideoAdsBetaRequestContent {
	this := CreateSponsoredBrandsVideoAdsBetaRequestContent{}
	return &this
}

// GetAds returns the Ads field value
func (o *CreateSponsoredBrandsVideoAdsBetaRequestContent) GetAds() []CreateVideoAd {
	if o == nil {
		var ret []CreateVideoAd
		return ret
	}

	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsVideoAdsBetaRequestContent) GetAdsOk() ([]CreateVideoAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads, true
}

// SetAds sets field value
func (o *CreateSponsoredBrandsVideoAdsBetaRequestContent) SetAds(v []CreateVideoAd) {
	o.Ads = v
}

func (o CreateSponsoredBrandsVideoAdsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ads"] = o.Ads
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsVideoAdsBetaRequestContent struct {
	value *CreateSponsoredBrandsVideoAdsBetaRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsVideoAdsBetaRequestContent) Get() *CreateSponsoredBrandsVideoAdsBetaRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsVideoAdsBetaRequestContent) Set(val *CreateSponsoredBrandsVideoAdsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsVideoAdsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsVideoAdsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsVideoAdsBetaRequestContent(val *CreateSponsoredBrandsVideoAdsBetaRequestContent) *NullableCreateSponsoredBrandsVideoAdsBetaRequestContent {
	return &NullableCreateSponsoredBrandsVideoAdsBetaRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsVideoAdsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsVideoAdsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
