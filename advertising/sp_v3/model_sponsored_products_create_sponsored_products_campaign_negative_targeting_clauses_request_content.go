package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}

// SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct for SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent
type SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct {
	// An array of Campaign Negative TargetingClauses.
	CampaignNegativeTargetingClauses []SponsoredProductsCreateCampaignNegativeTargetingClause `json:"campaignNegativeTargetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent

// NewSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent(campaignNegativeTargetingClauses []SponsoredProductsCreateCampaignNegativeTargetingClause) *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}
	this.CampaignNegativeTargetingClauses = campaignNegativeTargetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent{}
	return &this
}

// GetCampaignNegativeTargetingClauses returns the CampaignNegativeTargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetCampaignNegativeTargetingClauses() []SponsoredProductsCreateCampaignNegativeTargetingClause {
	if o == nil {
		var ret []SponsoredProductsCreateCampaignNegativeTargetingClause
		return ret
	}

	return o.CampaignNegativeTargetingClauses
}

// GetCampaignNegativeTargetingClausesOk returns a tuple with the CampaignNegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) GetCampaignNegativeTargetingClausesOk() ([]SponsoredProductsCreateCampaignNegativeTargetingClause, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignNegativeTargetingClauses, true
}

// SetCampaignNegativeTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) SetCampaignNegativeTargetingClauses(v []SponsoredProductsCreateCampaignNegativeTargetingClause) {
	o.CampaignNegativeTargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeTargetingClauses"] = o.CampaignNegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Get() *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent(val *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
