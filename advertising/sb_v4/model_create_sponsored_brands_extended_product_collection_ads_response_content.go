package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent{}

// CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent struct for CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent
type CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent instantiates a new CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent() *CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent {
	this := CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent{}
	return &this
}

// NewCreateSponsoredBrandsExtendedProductCollectionAdsResponseContentWithDefaults instantiates a new CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsExtendedProductCollectionAdsResponseContentWithDefaults() *CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent {
	this := CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent struct {
	value *CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) Get() *CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) Set(val *CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent(val *CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) *NullableCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent {
	return &NullableCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsExtendedProductCollectionAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
