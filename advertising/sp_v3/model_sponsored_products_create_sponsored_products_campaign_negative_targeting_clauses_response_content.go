package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}

// SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct for SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent
type SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct {
	CampaignNegativeTargetingClauses SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse `json:"campaignNegativeTargetingClauses"`
}

type _SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent

// NewSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent(campaignNegativeTargetingClauses SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}
	this.CampaignNegativeTargetingClauses = campaignNegativeTargetingClauses
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent{}
	return &this
}

// GetCampaignNegativeTargetingClauses returns the CampaignNegativeTargetingClauses field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetCampaignNegativeTargetingClauses() SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse
		return ret
	}

	return o.CampaignNegativeTargetingClauses
}

// GetCampaignNegativeTargetingClausesOk returns a tuple with the CampaignNegativeTargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) GetCampaignNegativeTargetingClausesOk() (*SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignNegativeTargetingClauses, true
}

// SetCampaignNegativeTargetingClauses sets field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) SetCampaignNegativeTargetingClauses(v SponsoredProductsBulkCampaignNegativeTargetingClauseOperationResponse) {
	o.CampaignNegativeTargetingClauses = v
}

func (o SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignNegativeTargetingClauses"] = o.CampaignNegativeTargetingClauses
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Get() *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent(val *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
