package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent struct for SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent
type SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent struct {
	CampaignNegativeKeywords SponsoredProductsBulkCampaignNegativeKeywordOperationResponse `json:"campaignNegativeKeywords"`
}

type _SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent(campaignNegativeKeywords SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent{}
	this.CampaignNegativeKeywords = campaignNegativeKeywords
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent{}
	return &this
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywords() SponsoredProductsBulkCampaignNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkCampaignNegativeKeywordOperationResponse
		return ret
	}

	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywordsOk() (*SponsoredProductsBulkCampaignNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeKeywords, true
}

// SetCampaignNegativeKeywords sets field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) SetCampaignNegativeKeywords(v SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent(val *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
