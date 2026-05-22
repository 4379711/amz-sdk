package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsAdsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsAdsBetaResponseContent{}

// UpdateSponsoredBrandsAdsBetaResponseContent struct for UpdateSponsoredBrandsAdsBetaResponseContent
type UpdateSponsoredBrandsAdsBetaResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewUpdateSponsoredBrandsAdsBetaResponseContent instantiates a new UpdateSponsoredBrandsAdsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsAdsBetaResponseContent() *UpdateSponsoredBrandsAdsBetaResponseContent {
	this := UpdateSponsoredBrandsAdsBetaResponseContent{}
	return &this
}

// NewUpdateSponsoredBrandsAdsBetaResponseContentWithDefaults instantiates a new UpdateSponsoredBrandsAdsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsAdsBetaResponseContentWithDefaults() *UpdateSponsoredBrandsAdsBetaResponseContent {
	this := UpdateSponsoredBrandsAdsBetaResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *UpdateSponsoredBrandsAdsBetaResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsAdsBetaResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *UpdateSponsoredBrandsAdsBetaResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *UpdateSponsoredBrandsAdsBetaResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o UpdateSponsoredBrandsAdsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsAdsBetaResponseContent struct {
	value *UpdateSponsoredBrandsAdsBetaResponseContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsAdsBetaResponseContent) Get() *UpdateSponsoredBrandsAdsBetaResponseContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsAdsBetaResponseContent) Set(val *UpdateSponsoredBrandsAdsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsAdsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsAdsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsAdsBetaResponseContent(val *UpdateSponsoredBrandsAdsBetaResponseContent) *NullableUpdateSponsoredBrandsAdsBetaResponseContent {
	return &NullableUpdateSponsoredBrandsAdsBetaResponseContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsAdsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsAdsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
