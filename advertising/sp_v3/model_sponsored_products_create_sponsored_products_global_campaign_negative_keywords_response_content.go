package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{}

// SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent struct for SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent
type SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent struct {
	CampaignNegativeKeywords SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse `json:"campaignNegativeKeywords"`
}

type _SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent

// NewSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent(campaignNegativeKeywords SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse) *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{}
	this.CampaignNegativeKeywords = campaignNegativeKeywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{}
	return &this
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywords() SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse
		return ret
	}

	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywordsOk() (*SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeKeywords, true
}

// SetCampaignNegativeKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) SetCampaignNegativeKeywords(v SponsoredProductsBulkGlobalCampaignNegativeKeywordOperationResponse) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent(val *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
