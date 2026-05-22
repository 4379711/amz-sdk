package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent struct for SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent
type SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent struct {
	CampaignIdFilter SponsoredProductsObjectIdFilter `json:"campaignIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent(campaignIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent{}
	this.CampaignIdFilter = campaignIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) GetCampaignIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignIdFilter, true
}

// SetCampaignIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) SetCampaignIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent(val *SponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
