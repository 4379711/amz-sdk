package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent struct for SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent
type SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent struct {
	CampaignNegativeKeywords SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse `json:"campaignNegativeKeywords"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent(campaignNegativeKeywords SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse) *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{}
	this.CampaignNegativeKeywords = campaignNegativeKeywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{}
	return &this
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywords() SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse
		return ret
	}

	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywordsOk() (*SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeKeywords, true
}

// SetCampaignNegativeKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) SetCampaignNegativeKeywords(v SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent(val *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
