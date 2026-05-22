package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsAdsResponseContent{}

// UpdateSponsoredBrandsAdsResponseContent struct for UpdateSponsoredBrandsAdsResponseContent
type UpdateSponsoredBrandsAdsResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewUpdateSponsoredBrandsAdsResponseContent instantiates a new UpdateSponsoredBrandsAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsAdsResponseContent() *UpdateSponsoredBrandsAdsResponseContent {
	this := UpdateSponsoredBrandsAdsResponseContent{}
	return &this
}

// NewUpdateSponsoredBrandsAdsResponseContentWithDefaults instantiates a new UpdateSponsoredBrandsAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsAdsResponseContentWithDefaults() *UpdateSponsoredBrandsAdsResponseContent {
	this := UpdateSponsoredBrandsAdsResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *UpdateSponsoredBrandsAdsResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsAdsResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *UpdateSponsoredBrandsAdsResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *UpdateSponsoredBrandsAdsResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o UpdateSponsoredBrandsAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsAdsResponseContent struct {
	value *UpdateSponsoredBrandsAdsResponseContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsAdsResponseContent) Get() *UpdateSponsoredBrandsAdsResponseContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsAdsResponseContent) Set(val *UpdateSponsoredBrandsAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsAdsResponseContent(val *UpdateSponsoredBrandsAdsResponseContent) *NullableUpdateSponsoredBrandsAdsResponseContent {
	return &NullableUpdateSponsoredBrandsAdsResponseContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
