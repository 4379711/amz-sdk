package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent{}

// CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent struct for CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent
type CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent instantiates a new CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent() *CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent {
	this := CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent{}
	return &this
}

// NewCreateSponsoredBrandStoreSpotlightAdsBetaResponseContentWithDefaults instantiates a new CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandStoreSpotlightAdsBetaResponseContentWithDefaults() *CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent {
	this := CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent struct {
	value *CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) Get() *CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) Set(val *CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent(val *CreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) *NullableCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent {
	return &NullableCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandStoreSpotlightAdsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
