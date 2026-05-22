package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsBrandVideoAdsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsBrandVideoAdsBetaResponseContent{}

// CreateSponsoredBrandsBrandVideoAdsBetaResponseContent struct for CreateSponsoredBrandsBrandVideoAdsBetaResponseContent
type CreateSponsoredBrandsBrandVideoAdsBetaResponseContent struct {
	Ads *BulkAdOperationResponse `json:"ads,omitempty"`
}

// NewCreateSponsoredBrandsBrandVideoAdsBetaResponseContent instantiates a new CreateSponsoredBrandsBrandVideoAdsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsBrandVideoAdsBetaResponseContent() *CreateSponsoredBrandsBrandVideoAdsBetaResponseContent {
	this := CreateSponsoredBrandsBrandVideoAdsBetaResponseContent{}
	return &this
}

// NewCreateSponsoredBrandsBrandVideoAdsBetaResponseContentWithDefaults instantiates a new CreateSponsoredBrandsBrandVideoAdsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsBrandVideoAdsBetaResponseContentWithDefaults() *CreateSponsoredBrandsBrandVideoAdsBetaResponseContent {
	this := CreateSponsoredBrandsBrandVideoAdsBetaResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *CreateSponsoredBrandsBrandVideoAdsBetaResponseContent) GetAds() BulkAdOperationResponse {
	if o == nil || IsNil(o.Ads) {
		var ret BulkAdOperationResponse
		return ret
	}
	return *o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsBrandVideoAdsBetaResponseContent) GetAdsOk() (*BulkAdOperationResponse, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *CreateSponsoredBrandsBrandVideoAdsBetaResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given BulkAdOperationResponse and assigns it to the Ads field.
func (o *CreateSponsoredBrandsBrandVideoAdsBetaResponseContent) SetAds(v BulkAdOperationResponse) {
	o.Ads = &v
}

func (o CreateSponsoredBrandsBrandVideoAdsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsBrandVideoAdsBetaResponseContent struct {
	value *CreateSponsoredBrandsBrandVideoAdsBetaResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsBetaResponseContent) Get() *CreateSponsoredBrandsBrandVideoAdsBetaResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsBetaResponseContent) Set(val *CreateSponsoredBrandsBrandVideoAdsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsBrandVideoAdsBetaResponseContent(val *CreateSponsoredBrandsBrandVideoAdsBetaResponseContent) *NullableCreateSponsoredBrandsBrandVideoAdsBetaResponseContent {
	return &NullableCreateSponsoredBrandsBrandVideoAdsBetaResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsBrandVideoAdsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsBrandVideoAdsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
