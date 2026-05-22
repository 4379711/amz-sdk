package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent{}

// SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent struct for SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent
type SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent struct {
	CampaignNegativeKeywords SponsoredProductsBulkCampaignNegativeKeywordOperationResponse `json:"campaignNegativeKeywords"`
}

type _SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent

// NewSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent(campaignNegativeKeywords SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent{}
	this.CampaignNegativeKeywords = campaignNegativeKeywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent{}
	return &this
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywords() SponsoredProductsBulkCampaignNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkCampaignNegativeKeywordOperationResponse
		return ret
	}

	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) GetCampaignNegativeKeywordsOk() (*SponsoredProductsBulkCampaignNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeKeywords, true
}

// SetCampaignNegativeKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) SetCampaignNegativeKeywords(v SponsoredProductsBulkCampaignNegativeKeywordOperationResponse) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent(val *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
