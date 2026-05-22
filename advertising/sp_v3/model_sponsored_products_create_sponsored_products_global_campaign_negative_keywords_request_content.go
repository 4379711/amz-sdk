package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}

// SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct for SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent
type SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct {
	// An array of campaignNegativeKeywords.
	CampaignNegativeKeywords []SponsoredProductsCreateGlobalCampaignNegativeKeyword `json:"campaignNegativeKeywords,omitempty"`
}

// NewSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent() *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}
	return &this
}

// GetCampaignNegativeKeywords returns the CampaignNegativeKeywords field value if set, zero value otherwise.
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywords() []SponsoredProductsCreateGlobalCampaignNegativeKeyword {
	if o == nil || IsNil(o.CampaignNegativeKeywords) {
		var ret []SponsoredProductsCreateGlobalCampaignNegativeKeyword
		return ret
	}
	return o.CampaignNegativeKeywords
}

// GetCampaignNegativeKeywordsOk returns a tuple with the CampaignNegativeKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordsOk() ([]SponsoredProductsCreateGlobalCampaignNegativeKeyword, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeywords) {
		return nil, false
	}
	return o.CampaignNegativeKeywords, true
}

// HasCampaignNegativeKeywords returns a boolean if a field has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) HasCampaignNegativeKeywords() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeywords) {
		return true
	}

	return false
}

// SetCampaignNegativeKeywords gets a reference to the given []SponsoredProductsCreateGlobalCampaignNegativeKeyword and assigns it to the CampaignNegativeKeywords field.
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) SetCampaignNegativeKeywords(v []SponsoredProductsCreateGlobalCampaignNegativeKeyword) {
	o.CampaignNegativeKeywords = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignNegativeKeywords) {
		toSerialize["campaignNegativeKeywords"] = o.CampaignNegativeKeywords
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent(val *SponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
