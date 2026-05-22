package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}

// SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct for SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent
type SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct {
	// An array of Campaign Negative TargetingClauses with updated values.
	CampaignNegativeTargetingClauses []SponsoredProductsUpdateCampaignNegativeTargetingClause `json:"campaignNegativeTargetingClauses"`
}

type _SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent

// NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent(campaignNegativeTargetingClauses []SponsoredProductsUpdateCampaignNegativeTargetingClause) *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}
	this.CampaignNegativeTargetingClauses = campaignNegativeTargetingClauses
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}
	return &this
}

// GetCampaignNegativeTargetingClauses returns the CampaignNegativeTargetingClauses field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetCampaignNegativeTargetingClauses() []SponsoredProductsUpdateCampaignNegativeTargetingClause {
	if o == nil {
		var ret []SponsoredProductsUpdateCampaignNegativeTargetingClause
		return ret
	}

	return o.CampaignNegativeTargetingClauses
}

// GetCampaignNegativeTargetingClausesOk returns a tuple with the CampaignNegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetCampaignNegativeTargetingClausesOk() ([]SponsoredProductsUpdateCampaignNegativeTargetingClause, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignNegativeTargetingClauses, true
}

// SetCampaignNegativeTargetingClauses sets field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) SetCampaignNegativeTargetingClauses(v []SponsoredProductsUpdateCampaignNegativeTargetingClause) {
	o.CampaignNegativeTargetingClauses = v
}

func (o SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeTargetingClauses"] = o.CampaignNegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent(val *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
