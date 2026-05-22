package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent struct for SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent
type SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent struct {
	CampaignNegativeKeywordIdFilter SponsoredProductsObjectIdFilter `json:"campaignNegativeKeywordIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent(campaignNegativeKeywordIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent{}
	this.CampaignNegativeKeywordIdFilter = campaignNegativeKeywordIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent{}
	return &this
}

// GetCampaignNegativeKeywordIdFilter returns the CampaignNegativeKeywordIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.CampaignNegativeKeywordIdFilter
}

// GetCampaignNegativeKeywordIdFilterOk returns a tuple with the CampaignNegativeKeywordIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeKeywordIdFilter, true
}

// SetCampaignNegativeKeywordIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) SetCampaignNegativeKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignNegativeKeywordIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywordIdFilter"] = o.CampaignNegativeKeywordIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent(val *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
