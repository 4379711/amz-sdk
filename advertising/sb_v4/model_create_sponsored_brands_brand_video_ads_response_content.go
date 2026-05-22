package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsBrandVideoAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsBrandVideoAdsResponseContent{}

// CreateSponsoredBrandsBrandVideoAdsResponseContent struct for CreateSponsoredBrandsBrandVideoAdsResponseContent
type CreateSponsoredBrandsBrandVideoAdsResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewCreateSponsoredBrandsBrandVideoAdsResponseContent instantiates a new CreateSponsoredBrandsBrandVideoAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsBrandVideoAdsResponseContent() *CreateSponsoredBrandsBrandVideoAdsResponseContent {
	this := CreateSponsoredBrandsBrandVideoAdsResponseContent{}
	return &this
}

// NewCreateSponsoredBrandsBrandVideoAdsResponseContentWithDefaults instantiates a new CreateSponsoredBrandsBrandVideoAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsBrandVideoAdsResponseContentWithDefaults() *CreateSponsoredBrandsBrandVideoAdsResponseContent {
	this := CreateSponsoredBrandsBrandVideoAdsResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *CreateSponsoredBrandsBrandVideoAdsResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsBrandVideoAdsResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *CreateSponsoredBrandsBrandVideoAdsResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *CreateSponsoredBrandsBrandVideoAdsResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o CreateSponsoredBrandsBrandVideoAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsBrandVideoAdsResponseContent struct {
	value *CreateSponsoredBrandsBrandVideoAdsResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsResponseContent) Get() *CreateSponsoredBrandsBrandVideoAdsResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsResponseContent) Set(val *CreateSponsoredBrandsBrandVideoAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsBrandVideoAdsResponseContent(val *CreateSponsoredBrandsBrandVideoAdsResponseContent) *NullableCreateSponsoredBrandsBrandVideoAdsResponseContent {
	return &NullableCreateSponsoredBrandsBrandVideoAdsResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
