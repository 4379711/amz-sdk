package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}

// SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct for SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent
type SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct {
	CampaignNegativeTargetIdFilter SponsoredProductsObjectIdFilter `json:"campaignNegativeTargetIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent

// NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent(campaignNegativeTargetIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}
	this.CampaignNegativeTargetIdFilter = campaignNegativeTargetIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}
	return &this
}

// GetCampaignNegativeTargetIdFilter returns the CampaignNegativeTargetIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetCampaignNegativeTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.CampaignNegativeTargetIdFilter
}

// GetCampaignNegativeTargetIdFilterOk returns a tuple with the CampaignNegativeTargetIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetCampaignNegativeTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeTargetIdFilter, true
}

// SetCampaignNegativeTargetIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) SetCampaignNegativeTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignNegativeTargetIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeTargetIdFilter"] = o.CampaignNegativeTargetIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent(val *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
