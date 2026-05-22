package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsAdsResponseContent{}

// DeleteSponsoredBrandsAdsResponseContent struct for DeleteSponsoredBrandsAdsResponseContent
type DeleteSponsoredBrandsAdsResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewDeleteSponsoredBrandsAdsResponseContent instantiates a new DeleteSponsoredBrandsAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsAdsResponseContent() *DeleteSponsoredBrandsAdsResponseContent {
	this := DeleteSponsoredBrandsAdsResponseContent{}
	return &this
}

// NewDeleteSponsoredBrandsAdsResponseContentWithDefaults instantiates a new DeleteSponsoredBrandsAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsAdsResponseContentWithDefaults() *DeleteSponsoredBrandsAdsResponseContent {
	this := DeleteSponsoredBrandsAdsResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsAdsResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsAdsResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsAdsResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *DeleteSponsoredBrandsAdsResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o DeleteSponsoredBrandsAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsAdsResponseContent struct {
	value *DeleteSponsoredBrandsAdsResponseContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsAdsResponseContent) Get() *DeleteSponsoredBrandsAdsResponseContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsAdsResponseContent) Set(val *DeleteSponsoredBrandsAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsAdsResponseContent(val *DeleteSponsoredBrandsAdsResponseContent) *NullableDeleteSponsoredBrandsAdsResponseContent {
	return &NullableDeleteSponsoredBrandsAdsResponseContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
