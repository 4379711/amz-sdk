package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the GetLandingPageMetadataRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLandingPageMetadataRequestContent{}

// GetLandingPageMetadataRequestContent struct for GetLandingPageMetadataRequestContent
type GetLandingPageMetadataRequestContent struct {
	// An ad product is a top level offering from amazon ads as defined in our marketing, with a given feature set, and business rules and logic applied consistently across the product. Currently the only supported ad product is SPONSORED_BRANDS. | Program Type       | |--------------------| | SPONSORED_BRANDS   |
	AdProduct string `json:"adProduct"`
	// The URL of the landing page.
	LandingPageUrl string `json:"landingPageUrl"`
}

type _GetLandingPageMetadataRequestContent GetLandingPageMetadataRequestContent

// NewGetLandingPageMetadataRequestContent instantiates a new GetLandingPageMetadataRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLandingPageMetadataRequestContent(adProduct string, landingPageUrl string) *GetLandingPageMetadataRequestContent {
	this := GetLandingPageMetadataRequestContent{}
	this.AdProduct = adProduct
	this.LandingPageUrl = landingPageUrl
	return &this
}

// NewGetLandingPageMetadataRequestContentWithDefaults instantiates a new GetLandingPageMetadataRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLandingPageMetadataRequestContentWithDefaults() *GetLandingPageMetadataRequestContent {
	this := GetLandingPageMetadataRequestContent{}
	return &this
}

// GetAdProduct returns the AdProduct field value
func (o *GetLandingPageMetadataRequestContent) GetAdProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdProduct
}

// GetAdProductOk returns a tuple with the AdProduct field value
// and a boolean to check if the value has been set.
func (o *GetLandingPageMetadataRequestContent) GetAdProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdProduct, true
}

// SetAdProduct sets field value
func (o *GetLandingPageMetadataRequestContent) SetAdProduct(v string) {
	o.AdProduct = v
}

// GetLandingPageUrl returns the LandingPageUrl field value
func (o *GetLandingPageMetadataRequestContent) GetLandingPageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LandingPageUrl
}

// GetLandingPageUrlOk returns a tuple with the LandingPageUrl field value
// and a boolean to check if the value has been set.
func (o *GetLandingPageMetadataRequestContent) GetLandingPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LandingPageUrl, true
}

// SetLandingPageUrl sets field value
func (o *GetLandingPageMetadataRequestContent) SetLandingPageUrl(v string) {
	o.LandingPageUrl = v
}

func (o GetLandingPageMetadataRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adProduct"] = o.AdProduct
	toSerialize["landingPageUrl"] = o.LandingPageUrl
	return toSerialize, nil
}

type NullableGetLandingPageMetadataRequestContent struct {
	value *GetLandingPageMetadataRequestContent
	isSet bool
}

func (v NullableGetLandingPageMetadataRequestContent) Get() *GetLandingPageMetadataRequestContent {
	return v.value
}

func (v *NullableGetLandingPageMetadataRequestContent) Set(val *GetLandingPageMetadataRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLandingPageMetadataRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLandingPageMetadataRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLandingPageMetadataRequestContent(val *GetLandingPageMetadataRequestContent) *NullableGetLandingPageMetadataRequestContent {
	return &NullableGetLandingPageMetadataRequestContent{value: val, isSet: true}
}

func (v NullableGetLandingPageMetadataRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetLandingPageMetadataRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
