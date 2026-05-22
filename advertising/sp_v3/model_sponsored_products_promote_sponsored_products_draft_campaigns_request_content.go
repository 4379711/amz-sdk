package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent{}

// SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent struct for SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent
type SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent struct {
	CampaignIdFilter SponsoredProductsObjectIdFilter `json:"campaignIdFilter"`
}

type _SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent

// NewSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent instantiates a new SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent(campaignIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent {
	this := SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent{}
	this.CampaignIdFilter = campaignIdFilter
	return &this
}

// NewSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContentWithDefaults() *SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent {
	this := SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value
func (o *SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) GetCampaignIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignIdFilter, true
}

// SetCampaignIdFilter sets field value
func (o *SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) SetCampaignIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignIdFilter = v
}

func (o SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent struct {
	value *SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) Get() *SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) Set(val *SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent(val *SponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) *NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent {
	return &NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
