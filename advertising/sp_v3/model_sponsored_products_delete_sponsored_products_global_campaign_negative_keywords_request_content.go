package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct for SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent
type SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct {
	CampaignNegativeKeywordIdFilter *SponsoredProductsObjectIdFilter `json:"campaignNegativeKeywordIdFilter,omitempty"`
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent() *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{}
	return &this
}

// GetCampaignNegativeKeywordIdFilter returns the CampaignNegativeKeywordIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.CampaignNegativeKeywordIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.CampaignNegativeKeywordIdFilter
}

// GetCampaignNegativeKeywordIdFilterOk returns a tuple with the CampaignNegativeKeywordIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) GetCampaignNegativeKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignNegativeKeywordIdFilter) {
		return nil, false
	}
	return o.CampaignNegativeKeywordIdFilter, true
}

// HasCampaignNegativeKeywordIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) HasCampaignNegativeKeywordIdFilter() bool {
	if o != nil && !IsNil(o.CampaignNegativeKeywordIdFilter) {
		return true
	}

	return false
}

// SetCampaignNegativeKeywordIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the CampaignNegativeKeywordIdFilter field.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) SetCampaignNegativeKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignNegativeKeywordIdFilter = &v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignNegativeKeywordIdFilter) {
		toSerialize["campaignNegativeKeywordIdFilter"] = o.CampaignNegativeKeywordIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent(val *SponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
