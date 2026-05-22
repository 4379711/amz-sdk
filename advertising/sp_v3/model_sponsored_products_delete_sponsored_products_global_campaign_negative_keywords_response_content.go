package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent struct for SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent
type SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent struct {
	CampaignNegativeKeywords SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse `json:"campaignNegativeKeywords"`
}

type _SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent(campaignNegativeKeywords SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse) *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{}
	this.CampaignNegativeKeywords = campaignNegativeKeywords
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{}
	return &this
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywords() SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse
		return ret
	}

	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywordsOk() (*SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeKeywords, true
}

// SetCampaignNegativeKeywords sets field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) SetCampaignNegativeKeywords(v SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent(val *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
