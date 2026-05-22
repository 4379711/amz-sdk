package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent struct for SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent
type SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent struct {
	CampaignNegativeKeywords SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse `json:"campaignNegativeKeywords"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent(campaignNegativeKeywords SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent{}
	this.CampaignNegativeKeywords = campaignNegativeKeywords
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent{}
	return &this
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywords() SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse
		return ret
	}

	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywordsOk() (*SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeKeywords, true
}

// SetCampaignNegativeKeywords sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) SetCampaignNegativeKeywords(v SponsoredProductsBulkDraftCampaignNegativeKeywordOperationResponse) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent(val *SponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
