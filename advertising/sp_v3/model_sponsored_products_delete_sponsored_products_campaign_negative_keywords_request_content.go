package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent struct for SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent
type SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent struct {
	CampaignNegativeKeywordIdFilter SponsoredProductsObjectIdFilter `json:"campaignNegativeKeywordIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent(campaignNegativeKeywordIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent{}
	this.CampaignNegativeKeywordIdFilter = campaignNegativeKeywordIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent{}
	return &this
}

// GetCampaignNegativeKeywordIdFilter returns the CampaignNegativeKeywordIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.CampaignNegativeKeywordIdFilter
}

// GetCampaignNegativeKeywordIdFilterOk returns a tuple with the CampaignNegativeKeywordIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeKeywordIdFilter, true
}

// SetCampaignNegativeKeywordIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) SetCampaignNegativeKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignNegativeKeywordIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywordIdFilter"] = o.CampaignNegativeKeywordIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent(val *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
