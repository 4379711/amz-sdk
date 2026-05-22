package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct for SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent
type SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct {
	// An array of campaignNegativeKeywords with updated values.
	CampaignNegativeKeywords []SponsoredProductsUpdateGlobalCampaignNegativeKeyword `json:"campaignNegativeKeywords"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent(campaignNegativeKeywords []SponsoredProductsUpdateGlobalCampaignNegativeKeyword) *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}
	this.CampaignNegativeKeywords = campaignNegativeKeywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}
	return &this
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywords() []SponsoredProductsUpdateGlobalCampaignNegativeKeyword {
	if o == nil {
		var ret []SponsoredProductsUpdateGlobalCampaignNegativeKeyword
		return ret
	}

	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordsOk() ([]SponsoredProductsUpdateGlobalCampaignNegativeKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignNegativeKeywords, true
}

// SetCampaignNegativeKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) SetCampaignNegativeKeywords(v []SponsoredProductsUpdateGlobalCampaignNegativeKeyword) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent(val *SponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
