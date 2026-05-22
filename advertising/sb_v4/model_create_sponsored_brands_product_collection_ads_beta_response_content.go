package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsProductCollectionAdsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsProductCollectionAdsBetaResponseContent{}

// CreateSponsoredBrandsProductCollectionAdsBetaResponseContent struct for CreateSponsoredBrandsProductCollectionAdsBetaResponseContent
type CreateSponsoredBrandsProductCollectionAdsBetaResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewCreateSponsoredBrandsProductCollectionAdsBetaResponseContent instantiates a new CreateSponsoredBrandsProductCollectionAdsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsProductCollectionAdsBetaResponseContent() *CreateSponsoredBrandsProductCollectionAdsBetaResponseContent {
	this := CreateSponsoredBrandsProductCollectionAdsBetaResponseContent{}
	return &this
}

// NewCreateSponsoredBrandsProductCollectionAdsBetaResponseContentWithDefaults instantiates a new CreateSponsoredBrandsProductCollectionAdsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsProductCollectionAdsBetaResponseContentWithDefaults() *CreateSponsoredBrandsProductCollectionAdsBetaResponseContent {
	this := CreateSponsoredBrandsProductCollectionAdsBetaResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *CreateSponsoredBrandsProductCollectionAdsBetaResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsProductCollectionAdsBetaResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *CreateSponsoredBrandsProductCollectionAdsBetaResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *CreateSponsoredBrandsProductCollectionAdsBetaResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o CreateSponsoredBrandsProductCollectionAdsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsProductCollectionAdsBetaResponseContent struct {
	value *CreateSponsoredBrandsProductCollectionAdsBetaResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsBetaResponseContent) Get() *CreateSponsoredBrandsProductCollectionAdsBetaResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsBetaResponseContent) Set(val *CreateSponsoredBrandsProductCollectionAdsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsProductCollectionAdsBetaResponseContent(val *CreateSponsoredBrandsProductCollectionAdsBetaResponseContent) *NullableCreateSponsoredBrandsProductCollectionAdsBetaResponseContent {
	return &NullableCreateSponsoredBrandsProductCollectionAdsBetaResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsProductCollectionAdsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsProductCollectionAdsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
