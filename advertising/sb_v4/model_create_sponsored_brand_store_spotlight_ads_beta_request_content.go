package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent{}

// CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent struct for CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent
type CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent struct {
	Ads []CreateStoreSpotlightAd `json:"ads"`
}

type _CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent

// NewCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent instantiates a new CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent(ads []CreateStoreSpotlightAd) *CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent {
	this := CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent{}
	this.Ads = ads
	return &this
}

// NewCreateSponsoredBrandStoreSpotlightAdsBetaRequestContentWithDefaults instantiates a new CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandStoreSpotlightAdsBetaRequestContentWithDefaults() *CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent {
	this := CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent{}
	return &this
}

// GetAds returns the Ads field value
func (o *CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) GetAds() []CreateStoreSpotlightAd {
	if o == nil {
		var ret []CreateStoreSpotlightAd
		return ret
	}

	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) GetAdsOk() ([]CreateStoreSpotlightAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads, true
}

// SetAds sets field value
func (o *CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) SetAds(v []CreateStoreSpotlightAd) {
	o.Ads = v
}

func (o CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ads"] = o.Ads
	return toSerialize, nil
}

type NullableCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent struct {
	value *CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) Get() *CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) Set(val *CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent(val *CreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) *NullableCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent {
	return &NullableCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
