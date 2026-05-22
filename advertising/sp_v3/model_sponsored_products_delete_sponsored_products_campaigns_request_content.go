package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent struct for SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent
type SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent struct {
	CampaignIdFilter SponsoredProductsObjectIdFilter `json:"campaignIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent(campaignIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent{}
	this.CampaignIdFilter = campaignIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsCampaignsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) GetCampaignIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignIdFilter, true
}

// SetCampaignIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) SetCampaignIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent(val *SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
