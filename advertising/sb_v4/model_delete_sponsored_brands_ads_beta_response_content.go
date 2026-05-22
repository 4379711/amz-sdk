package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsAdsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsAdsBetaResponseContent{}

// DeleteSponsoredBrandsAdsBetaResponseContent struct for DeleteSponsoredBrandsAdsBetaResponseContent
type DeleteSponsoredBrandsAdsBetaResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewDeleteSponsoredBrandsAdsBetaResponseContent instantiates a new DeleteSponsoredBrandsAdsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsAdsBetaResponseContent() *DeleteSponsoredBrandsAdsBetaResponseContent {
	this := DeleteSponsoredBrandsAdsBetaResponseContent{}
	return &this
}

// NewDeleteSponsoredBrandsAdsBetaResponseContentWithDefaults instantiates a new DeleteSponsoredBrandsAdsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsAdsBetaResponseContentWithDefaults() *DeleteSponsoredBrandsAdsBetaResponseContent {
	this := DeleteSponsoredBrandsAdsBetaResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsAdsBetaResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsAdsBetaResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsAdsBetaResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *DeleteSponsoredBrandsAdsBetaResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o DeleteSponsoredBrandsAdsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsAdsBetaResponseContent struct {
	value *DeleteSponsoredBrandsAdsBetaResponseContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsAdsBetaResponseContent) Get() *DeleteSponsoredBrandsAdsBetaResponseContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsAdsBetaResponseContent) Set(val *DeleteSponsoredBrandsAdsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsAdsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsAdsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsAdsBetaResponseContent(val *DeleteSponsoredBrandsAdsBetaResponseContent) *NullableDeleteSponsoredBrandsAdsBetaResponseContent {
	return &NullableDeleteSponsoredBrandsAdsBetaResponseContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsAdsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsAdsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
