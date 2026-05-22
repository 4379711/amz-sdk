package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent struct for SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent
type SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent struct {
	CampaignNegativeKeywords SponsoredProductsBulkCampaignNegativeKeywordOperationResponse `json:"campaignNegativeKeywords"`
}

type _SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent(campaignNegativeKeywords SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent{}
	this.CampaignNegativeKeywords = campaignNegativeKeywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent{}
	return &this
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywords() SponsoredProductsBulkCampaignNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkCampaignNegativeKeywordOperationResponse
		return ret
	}

	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywordsOk() (*SponsoredProductsBulkCampaignNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeKeywords, true
}

// SetCampaignNegativeKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) SetCampaignNegativeKeywords(v SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent(val *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
