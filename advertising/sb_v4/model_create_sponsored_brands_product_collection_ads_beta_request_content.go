package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsProductCollectionAdsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsProductCollectionAdsBetaRequestContent{}

// CreateSponsoredBrandsProductCollectionAdsBetaRequestContent struct for CreateSponsoredBrandsProductCollectionAdsBetaRequestContent
type CreateSponsoredBrandsProductCollectionAdsBetaRequestContent struct {
	Ads []CreateProductCollectionAd `json:"ads"`
}

type _CreateSponsoredBrandsProductCollectionAdsBetaRequestContent CreateSponsoredBrandsProductCollectionAdsBetaRequestContent

// NewCreateSponsoredBrandsProductCollectionAdsBetaRequestContent instantiates a new CreateSponsoredBrandsProductCollectionAdsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsProductCollectionAdsBetaRequestContent(ads []CreateProductCollectionAd) *CreateSponsoredBrandsProductCollectionAdsBetaRequestContent {
	this := CreateSponsoredBrandsProductCollectionAdsBetaRequestContent{}
	this.Ads = ads
	return &this
}

// NewCreateSponsoredBrandsProductCollectionAdsBetaRequestContentWithDefaults instantiates a new CreateSponsoredBrandsProductCollectionAdsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsProductCollectionAdsBetaRequestContentWithDefaults() *CreateSponsoredBrandsProductCollectionAdsBetaRequestContent {
	this := CreateSponsoredBrandsProductCollectionAdsBetaRequestContent{}
	return &this
}

// GetAds returns the Ads field value
func (o *CreateSponsoredBrandsProductCollectionAdsBetaRequestContent) GetAds() []CreateProductCollectionAd {
	if o == nil {
		var ret []CreateProductCollectionAd
		return ret
	}

	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsProductCollectionAdsBetaRequestContent) GetAdsOk() ([]CreateProductCollectionAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads, true
}

// SetAds sets field value
func (o *CreateSponsoredBrandsProductCollectionAdsBetaRequestContent) SetAds(v []CreateProductCollectionAd) {
	o.Ads = v
}

func (o CreateSponsoredBrandsProductCollectionAdsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ads"] = o.Ads
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsProductCollectionAdsBetaRequestContent struct {
	value *CreateSponsoredBrandsProductCollectionAdsBetaRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsBetaRequestContent) Get() *CreateSponsoredBrandsProductCollectionAdsBetaRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsBetaRequestContent) Set(val *CreateSponsoredBrandsProductCollectionAdsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsProductCollectionAdsBetaRequestContent(val *CreateSponsoredBrandsProductCollectionAdsBetaRequestContent) *NullableCreateSponsoredBrandsProductCollectionAdsBetaRequestContent {
	return &NullableCreateSponsoredBrandsProductCollectionAdsBetaRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
