package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent struct for SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent
type SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent struct {
	// An array of campaignNegativeKeywords with updated values.
	CampaignNegativeKeywords []SponsoredProductsUpdateCampaignNegativeKeyword `json:"campaignNegativeKeywords"`
}

type _SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent(campaignNegativeKeywords []SponsoredProductsUpdateCampaignNegativeKeyword) *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent{}
	this.CampaignNegativeKeywords = campaignNegativeKeywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent{}
	return &this
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywords() []SponsoredProductsUpdateCampaignNegativeKeyword {
	if o == nil {
		var ret []SponsoredProductsUpdateCampaignNegativeKeyword
		return ret
	}

	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordsOk() ([]SponsoredProductsUpdateCampaignNegativeKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignNegativeKeywords, true
}

// SetCampaignNegativeKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) SetCampaignNegativeKeywords(v []SponsoredProductsUpdateCampaignNegativeKeyword) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent(val *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
