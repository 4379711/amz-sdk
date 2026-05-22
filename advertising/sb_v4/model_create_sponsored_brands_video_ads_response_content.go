package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsVideoAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsVideoAdsResponseContent{}

// CreateSponsoredBrandsVideoAdsResponseContent struct for CreateSponsoredBrandsVideoAdsResponseContent
type CreateSponsoredBrandsVideoAdsResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewCreateSponsoredBrandsVideoAdsResponseContent instantiates a new CreateSponsoredBrandsVideoAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsVideoAdsResponseContent() *CreateSponsoredBrandsVideoAdsResponseContent {
	this := CreateSponsoredBrandsVideoAdsResponseContent{}
	return &this
}

// NewCreateSponsoredBrandsVideoAdsResponseContentWithDefaults instantiates a new CreateSponsoredBrandsVideoAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsVideoAdsResponseContentWithDefaults() *CreateSponsoredBrandsVideoAdsResponseContent {
	this := CreateSponsoredBrandsVideoAdsResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *CreateSponsoredBrandsVideoAdsResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsVideoAdsResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *CreateSponsoredBrandsVideoAdsResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *CreateSponsoredBrandsVideoAdsResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o CreateSponsoredBrandsVideoAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsVideoAdsResponseContent struct {
	value *CreateSponsoredBrandsVideoAdsResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsVideoAdsResponseContent) Get() *CreateSponsoredBrandsVideoAdsResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsVideoAdsResponseContent) Set(val *CreateSponsoredBrandsVideoAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsVideoAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsVideoAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsVideoAdsResponseContent(val *CreateSponsoredBrandsVideoAdsResponseContent) *NullableCreateSponsoredBrandsVideoAdsResponseContent {
	return &NullableCreateSponsoredBrandsVideoAdsResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsVideoAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsVideoAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
