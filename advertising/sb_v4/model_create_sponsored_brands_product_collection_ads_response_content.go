package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsProductCollectionAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsProductCollectionAdsResponseContent{}

// CreateSponsoredBrandsProductCollectionAdsResponseContent struct for CreateSponsoredBrandsProductCollectionAdsResponseContent
type CreateSponsoredBrandsProductCollectionAdsResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewCreateSponsoredBrandsProductCollectionAdsResponseContent instantiates a new CreateSponsoredBrandsProductCollectionAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsProductCollectionAdsResponseContent() *CreateSponsoredBrandsProductCollectionAdsResponseContent {
	this := CreateSponsoredBrandsProductCollectionAdsResponseContent{}
	return &this
}

// NewCreateSponsoredBrandsProductCollectionAdsResponseContentWithDefaults instantiates a new CreateSponsoredBrandsProductCollectionAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsProductCollectionAdsResponseContentWithDefaults() *CreateSponsoredBrandsProductCollectionAdsResponseContent {
	this := CreateSponsoredBrandsProductCollectionAdsResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *CreateSponsoredBrandsProductCollectionAdsResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsProductCollectionAdsResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *CreateSponsoredBrandsProductCollectionAdsResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *CreateSponsoredBrandsProductCollectionAdsResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o CreateSponsoredBrandsProductCollectionAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsProductCollectionAdsResponseContent struct {
	value *CreateSponsoredBrandsProductCollectionAdsResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsResponseContent) Get() *CreateSponsoredBrandsProductCollectionAdsResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsResponseContent) Set(val *CreateSponsoredBrandsProductCollectionAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsProductCollectionAdsResponseContent(val *CreateSponsoredBrandsProductCollectionAdsResponseContent) *NullableCreateSponsoredBrandsProductCollectionAdsResponseContent {
	return &NullableCreateSponsoredBrandsProductCollectionAdsResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
