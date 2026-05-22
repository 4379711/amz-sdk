package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsVideoAdsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsVideoAdsBetaResponseContent{}

// CreateSponsoredBrandsVideoAdsBetaResponseContent struct for CreateSponsoredBrandsVideoAdsBetaResponseContent
type CreateSponsoredBrandsVideoAdsBetaResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewCreateSponsoredBrandsVideoAdsBetaResponseContent instantiates a new CreateSponsoredBrandsVideoAdsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsVideoAdsBetaResponseContent() *CreateSponsoredBrandsVideoAdsBetaResponseContent {
	this := CreateSponsoredBrandsVideoAdsBetaResponseContent{}
	return &this
}

// NewCreateSponsoredBrandsVideoAdsBetaResponseContentWithDefaults instantiates a new CreateSponsoredBrandsVideoAdsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsVideoAdsBetaResponseContentWithDefaults() *CreateSponsoredBrandsVideoAdsBetaResponseContent {
	this := CreateSponsoredBrandsVideoAdsBetaResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *CreateSponsoredBrandsVideoAdsBetaResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsVideoAdsBetaResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *CreateSponsoredBrandsVideoAdsBetaResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *CreateSponsoredBrandsVideoAdsBetaResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o CreateSponsoredBrandsVideoAdsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsVideoAdsBetaResponseContent struct {
	value *CreateSponsoredBrandsVideoAdsBetaResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsVideoAdsBetaResponseContent) Get() *CreateSponsoredBrandsVideoAdsBetaResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsVideoAdsBetaResponseContent) Set(val *CreateSponsoredBrandsVideoAdsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsVideoAdsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsVideoAdsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsVideoAdsBetaResponseContent(val *CreateSponsoredBrandsVideoAdsBetaResponseContent) *NullableCreateSponsoredBrandsVideoAdsBetaResponseContent {
	return &NullableCreateSponsoredBrandsVideoAdsBetaResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsVideoAdsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsVideoAdsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
