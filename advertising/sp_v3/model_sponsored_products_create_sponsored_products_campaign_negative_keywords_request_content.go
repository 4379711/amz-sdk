package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent{}

// SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent struct for SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent
type SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent struct {
	// An array of campaignNegativeKeywords.
	CampaignNegativeKeywords []SponsoredProductsCreateCampaignNegativeKeyword `json:"campaignNegativeKeywords"`
}

type _SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent

// NewSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent(campaignNegativeKeywords []SponsoredProductsCreateCampaignNegativeKeyword) *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent{}
	this.CampaignNegativeKeywords = campaignNegativeKeywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent{}
	return &this
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywords() []SponsoredProductsCreateCampaignNegativeKeyword {
	if o == nil {
		var ret []SponsoredProductsCreateCampaignNegativeKeyword
		return ret
	}

	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordsOk() ([]SponsoredProductsCreateCampaignNegativeKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignNegativeKeywords, true
}

// SetCampaignNegativeKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) SetCampaignNegativeKeywords(v []SponsoredProductsCreateCampaignNegativeKeyword) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent(val *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
