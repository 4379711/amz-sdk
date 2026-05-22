package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandStoreSpotlightAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandStoreSpotlightAdsResponseContent{}

// CreateSponsoredBrandStoreSpotlightAdsResponseContent struct for CreateSponsoredBrandStoreSpotlightAdsResponseContent
type CreateSponsoredBrandStoreSpotlightAdsResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewCreateSponsoredBrandStoreSpotlightAdsResponseContent instantiates a new CreateSponsoredBrandStoreSpotlightAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandStoreSpotlightAdsResponseContent() *CreateSponsoredBrandStoreSpotlightAdsResponseContent {
	this := CreateSponsoredBrandStoreSpotlightAdsResponseContent{}
	return &this
}

// NewCreateSponsoredBrandStoreSpotlightAdsResponseContentWithDefaults instantiates a new CreateSponsoredBrandStoreSpotlightAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandStoreSpotlightAdsResponseContentWithDefaults() *CreateSponsoredBrandStoreSpotlightAdsResponseContent {
	this := CreateSponsoredBrandStoreSpotlightAdsResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *CreateSponsoredBrandStoreSpotlightAdsResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandStoreSpotlightAdsResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *CreateSponsoredBrandStoreSpotlightAdsResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *CreateSponsoredBrandStoreSpotlightAdsResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o CreateSponsoredBrandStoreSpotlightAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandStoreSpotlightAdsResponseContent struct {
	value *CreateSponsoredBrandStoreSpotlightAdsResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsResponseContent) Get() *CreateSponsoredBrandStoreSpotlightAdsResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsResponseContent) Set(val *CreateSponsoredBrandStoreSpotlightAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandStoreSpotlightAdsResponseContent(val *CreateSponsoredBrandStoreSpotlightAdsResponseContent) *NullableCreateSponsoredBrandStoreSpotlightAdsResponseContent {
	return &NullableCreateSponsoredBrandStoreSpotlightAdsResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
